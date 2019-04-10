package gcp

import (
	"encoding/binary"
	"errors"
	"fmt"
	"strconv"
)

// A TLV represents a RCP TLV message.
type TLV struct {
	Type   uint8  // Type: 1 byte
	Length uint16 // Value Length: 2 bytes
	Value  []byte
}

// An RCP encodes a TLV used in the Remote PHY System Control Plane (RCP).
type RCP interface {
	Name() string
	Len() uint16
	Val() interface{}
	IsComplex() bool
	marshal() ([]byte, error)
	unmarshal(b []byte) error
}

// Error messages
var (
	ErrUnexpectedEOF = errors.New("unexpected EOF")
)

// Name returns the name of the TLV Type.
func (t *TLV) Name() string { return strconv.Itoa(int(t.Type)) }

// Len returns the length of the TLV.
func (t *TLV) Len() uint16 { return t.Length }

func (t *TLV) marshal() ([]byte, error) {
	l := int(t.Length)
	if len(t.Value) != l {
		return nil, ErrUnexpectedEOF
	}

	b := make([]byte, t.Length)
	b[0] = t.Type
	binary.BigEndian.PutUint16(b[1:3], t.Length)

	copy(b[3:], t.Value)

	return b, nil
}

func (t *TLV) unmarshal(b []byte) error {
	if len(b) < 3 {
		return ErrUnexpectedEOF
	}

	t.Type = b[0]
	t.Length = binary.BigEndian.Uint16(b[1:3])
	l := int(t.Length)

	// Enforce a valid length value that matches the expected one.
	if lb := len(b[3:]); l != lb {
		return fmt.Errorf("TLV length should be %d, but length is %d", l, lb)
	}

	t.Value = make([]byte, l)
	copy(t.Value, b[3:])

	return nil
}

// IsComplex returns whether the TLV is Complex or not.
func (t *TLV) IsComplex() bool {
	tp := RCPType(t.Type)
	switch tp {
	// Complex TLV
	case IRA, REX, Seq, RfPortSel, RpdCap, GnrlNtf, RpdInfo:
		return true
	case Oper:
		return false
	default:
		return false
	}
}

// Val returns the value the TLV carries.
func (t *TLV) Val() interface{} {
	return t.Value
}

// A NTF is a NTF Message TLV (Complex TLV).
type NTF struct {
	Type   uint8  // Type: 1 byte
	Length uint16 // Value Length: 2 bytes
	Value  []byte
}

// Name returns the type name of a NTF Message TLV.
func (t *NTF) Name() string { return "NTF" }

// Len returns the Length of a NTF Message TLV.
func (t *NTF) Len() uint16 { return t.Length }

func (t *NTF) marshal() ([]byte, error) {
	tlv := &TLV{
		Type:   t.Type,
		Length: t.Length,
		Value:  t.Value,
	}
	return tlv.marshal()
}

func (t *NTF) unmarshal(b []byte) error {
	tlv := new(TLV)
	if err := tlv.unmarshal(b); err != nil {
		return err
	}
	t.Type = tlv.Type
	t.Length = tlv.Length
	t.Value = append(tlv.Value[:0:0], tlv.Value...)
	return nil
}

// IsComplex returns whether a NTF Message TLV is Complex or not.
func (t *NTF) IsComplex() bool {
	return true
}

// Val returns the value a NTF Message TLV carries.
func (t *NTF) Val() interface{} {
	return t.Value
}

// A SeqNmr is a SequenceNumber TLV.
type SeqNmr struct {
	Type   uint8  // Type: 1 byte
	Length uint16 // Value Length: 2 bytes
	Value  []byte
}

// Name returns the type name of a NTF Message TLV.
func (t *SeqNmr) Name() string { return "SequenceNumber" }

// Len returns the Length of a SequenceNumber TLV.
func (t *SeqNmr) Len() uint16 {
	// TODO: Verify this is always 2.
	return t.Length
}

func (t *SeqNmr) marshal() ([]byte, error) {
	tlv := &TLV{
		Type:   t.Type,
		Length: t.Length,
		Value:  t.Value,
	}
	// TODO: Validate Length and Value Length is 2.
	return tlv.marshal()
}

func (t *SeqNmr) unmarshal(b []byte) error {
	tlv := new(TLV)
	if err := tlv.unmarshal(b); err != nil {
		return err
	}
	t.Type = tlv.Type
	t.Length = tlv.Length
	// TODO: Validate Length and Value Length is 2.
	t.Value = append(tlv.Value[:0:0], tlv.Value...)
	return nil
}

// IsComplex returns whether a SequenceNumber TLV is Complex or not.
func (t *SeqNmr) IsComplex() bool {
	return false
}

// Val returns the value a SequenceNumber TLV carries.
func (t *SeqNmr) Val() interface{} {
	if len(t.Value) != 2 {
		return fmt.Errorf("unexpected lenght: %v, want: 2", len(t.Value))
	}
	return binary.BigEndian.Uint16(t.Value)
}

// parseTLVs ...
func parseTLVs(b []byte) ([]RCP, error) {
	var tlvs []RCP
	for i := 0; len(b[i:]) != 0; {
		// Three bytes: TLV type and TLV length.
		if len(b[i:]) < 3 {
			return nil, ErrUnexpectedEOF
		}

		ty := int(b[i])
		l := int(binary.BigEndian.Uint16(b[i+1 : i+3]))

		// Verify that we won't advance beyond the end of the byte slice.
		if l > len(b[i+3:]) {
			return nil, ErrUnexpectedEOF
		}
		var tlv RCP
		// Select Type of TLV depending on ty
		// This will be problematic for recursive calls as TLV numbers
		// are re-used. Fix with [1]
		switch ty {
		case 3:
			tlv = new(NTF)
		case 10:
			tlv = new(SeqNmr)
		default:
			tlv = new(TLV)
		}

		// Unmarshal at the current offset, up to the expected length.
		if err := tlv.unmarshal(b[i : i+3+l]); err != nil {
			return nil, err
		}

		switch {
		// Complex TLV
		case l > 3 && tlv.IsComplex():
			// Recursive call
			// Will need to create parseTLV per type. [1]
			rectlv, err := parseTLVs(b[i+3 : i+3+l])
			if err != nil {
				return nil, err
			}
			tlvs = append(tlvs, tlv)
			tlvs = append(tlvs, rectlv...)
		case l < 3 || !tlv.IsComplex():
			tlvs = append(tlvs, tlv)
		default:
			fmt.Printf("We really shouldn't get here: TLV Type %s\n", tlv.Name())
		}
		// Advance to the next TLV's type field.
		i += (l + 3)
	}

	return tlvs, nil

}

// RCPType is a RCP TLV Type.
type RCPType uint8

// RCP TLV Types
const (
	IRA       RCPType = 1   // Complex TLV - IRA Message
	REX       RCPType = 2   // Complex TLV - REX Message
	Seq       RCPType = 9   // Complex TLV - Sequence
	Oper      RCPType = 11  // UnsignedByte - Operation
	RfPortSel RCPType = 13  // Complex TLV - RfPortSelector
	RpdCap    RCPType = 50  // Complex TLV - RpdCapabilities
	GnrlNtf   RCPType = 86  // Complex TLV - GeneralNotification
	RpdInfo   RCPType = 100 // Complex TLV - RpdInfo
)

// General purpose TLVs
//
// RfChannelSelector Complex TLV 12
// RfPortSelector Complex TLV 13
// RpdGlobal Complex TLV 15
// VendorSpecificExtension Complex TLV 21
// RpdRedirect Complex TLV 25

// RPD Capabilities TLVs
// 50
// RpdCapabilities Complex TLV 50
//   RpdIdentification Complex TLV 19
//   LcceChannelReachability Complex TLV 20
//   PilotToneCapabilities Complex TLV 21
//   AllocDsChanResources Complex TLV 22
//   AllocUsChanResources Complex TLV 23
//   DeviceLocation Complex TLV 24
//   RdtiCapabilities Complex TLV 34
//   UsPowerCapabilities Complex TLV 49
//   StaticPwCapabilities Complex TLV 50
//   DsCapabilities Complex TLV 51
//   GcpCapabilities Complex TLV 52
//   SwImageCapabilities Complex TLV 53
//   OfdmConfigurationCapabilities Complex TLV 54
//     PmapCapabilities Complex TLV 3
//   ResetCapabilities Complex TLV 55
//   RpdCoreRedundancyCapabilities Complex TLV 56
//   FdxCapabilities Complex 57
//     EcCapabilities Complex TLV 4
//   SpectrumCaptureCapabilities Complex TLV 59
//     SacCapabilities Complex TLV 2
//   RfmCapabilities Complex TLV 60
//     NodeRfPortCapabilities Complex TLV 17
//   UpstreamCapabilities Complex TLV 61
//   PmtudCapabilities Complex TLV 62

// RPD Operational Configuration TLVs
// 15
// RpdGlobal Complex TLV 15 ?
//   EvCfg Complex TLV 1
//     EvControl Complex TLV 1
//   GcpConnVerification Complex TLV 2
//   IpConfig Complex TLV 3
//   UepiControl Complex TLV 4
//   LldpConfig Complex TLV 6
// ...
//
//
//
//
//
//
//
