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
	// The output of parseTLVs could change to adapt to requirements (type struct?).
	parseTLVs(b []byte) ([]RCP, error)
}

// Error messages
var (
	ErrUnexpectedEOF = errors.New("unexpected EOF")
)

// Name returns the name of the TLV Type.
func (t *TLV) Name() string { return strconv.Itoa(int(t.Type)) }

// Len returns the length of the TLV.
func (t *TLV) Len() uint16 { return t.Length }

// Val returns the value the TLV carries.
func (t *TLV) Val() interface{} { return t.Value }

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
func (t *TLV) IsComplex() bool { return false }

func (t *TLV) parseTLVs(b []byte) ([]RCP, error) {
	var tlvs []RCP
	for i := 0; len(b[i:]) != 0; {
		l, err := boundsChk(i, b)
		if err != nil {
			return nil, err
		}

		tlv := t.newTLV(b[i])

		// Unmarshal at the current offset, up to the expected length.
		if err := tlv.unmarshal(b[i : i+3+l]); err != nil {
			return nil, err
		}

		switch {
		// Complex TLV
		case l > 3 && tlv.IsComplex():
			rectlv, err := tlv.parseTLVs(b[i+3 : i+3+l])
			if err != nil {
				return nil, err
			}
			tlvs = append(tlvs, tlv)
			tlvs = append(tlvs, rectlv...)
		case l <= 3 || !tlv.IsComplex():
			tlvs = append(tlvs, tlv)
		default:
			fmt.Printf("We really shouldn't get here: TLV Type %s\n", tlv.Name())
		}
		// Advance to the next TLV's type field.
		i += (l + 3)
	}

	return tlvs, nil
}

func (t *TLV) newTLV(b byte) RCP {
	switch int(b) {
	case 1:
		return new(IRA)
	case 2:
		return new(REX)
	case 3:
		return new(NTF)
	case 9:
		return new(Seq)
	case 10:
		return new(SeqNmr)
	case 11:
		return new(Oper)
	case 50:
		return new(RpdCap)
	case 86:
		return new(GenrlNtf)
	case 100:
		return new(RpdInfo)
	default:
		return new(TLV)
	}
}

// A IRA is a IRA Message TLV (Complex TLV).
type IRA struct {
	TLV
}

// Name returns the type name of a IRA Message TLV.
func (t *IRA) Name() string { return "IRA" }

// IsComplex returns whether a IRA Message TLV is Complex or not.
func (t *IRA) IsComplex() bool { return true }

// A REX is a REX Message TLV (Complex TLV).
type REX struct {
	TLV
}

// Name returns the type name of a REX Message TLV.
func (t *REX) Name() string { return "REX" }

// IsComplex returns whether a REX Message TLV is Complex or not.
func (t *REX) IsComplex() bool { return true }

// A NTF is a Notify Message TLV (Complex TLV).
type NTF struct {
	TLV
}

// Name returns the type name of a Notify Message TLV.
func (t *NTF) Name() string { return "Notify" }

// IsComplex returns whether a Notify Message TLV is Complex or not.
func (t *NTF) IsComplex() bool { return true }

// A Seq is a Sequence TLV (Complex TLV).
type Seq struct {
	TLV
}

// Name returns the type name of a Sequence TLV.
func (t *Seq) Name() string { return "Sequence" }

// IsComplex returns whether a Sequence TLV is Complex or not.
func (t *Seq) IsComplex() bool { return true }

// A SeqNmr is a SequenceNumber TLV.
type SeqNmr struct {
	TLV
}

// Name returns the type name of a NTF Message TLV.
func (t *SeqNmr) Name() string { return "SequenceNumber" }

// Val returns the value a SequenceNumber TLV carries.
func (t *SeqNmr) Val() interface{} { return u16Val(t.Value) }

// IsComplex returns whether a SequenceNumber TLV is Complex or not.
func (t *SeqNmr) IsComplex() bool { return false }

// A Oper is a Operation TLV.
type Oper struct {
	TLV
}

// Name returns the type name of a Operation TLV.
func (t *Oper) Name() string { return "Operation" }

// Val returns the value a Operation TLV carries.
func (t *Oper) Val() interface{} {
	if len(t.Value) != 1 {
		return fmt.Errorf("unexpected lenght: %v, want: 1", len(t.Value))
	}
	switch int(t.Value[0]) {
	case 1:
		return "Read"
	case 2:
		return "Write"
	case 3:
		return "Delete"
	case 4:
		return "ReadResponse"
	case 5:
		return "WriteResponse"
	case 6:
		return "DeleteResponse"
	case 7:
		return "AllocateWrite"
	case 8:
		return "AllocateWriteResponse"
	}
	return "Unknown Operation"
}

// IsComplex returns whether a Operation TLV is Complex or not.
func (t *Oper) IsComplex() bool { return false }

// parseTLVs parses Top Level and General Purpose TLVs.
func parseTLVs(b []byte) ([]RCP, error) {
	var t TLV
	return t.parseTLVs(b)
}

func boundsChk(i int, b []byte) (int, error) {
	// Three bytes: TLV type and TLV length.
	if len(b[i:]) < 3 {
		return 0, fmt.Errorf("TLV Length is %d (less than 3). Index: %d, Type: %v", len(b[i:]), i, b[i])
	}
	l := int(binary.BigEndian.Uint16(b[i+1 : i+3]))
	// Verify that we won't advance beyond the end of the byte slice.
	if l > len(b[i+3:]) {
		return 0, fmt.Errorf("TLV Length is greater than bytes in the buffer. Index: %d, Type: %v", i, b[i])
	}
	return l, nil
}

func stringVal(b []byte) string {
	if len(b) > 255 {
		return fmt.Sprintf("unexpected lenght: %v, want: 0-255", len(b))
	}
	return string(b)
}

func u16Val(b []byte) interface{} {
	// TODO: Should I stick to uint16 and return "0" if there is an error?
	if len(b) != 2 {
		return fmt.Sprintf("unexpected lenght: %v, want: 2", len(b))
	}
	return binary.BigEndian.Uint16(b)
}

func u32Val(b []byte) interface{} {
	// TODO: Should I stick to uint32 and return "0" if there is an error?
	if len(b) != 4 {
		return fmt.Sprintf("unexpected lenght: %v, want: 4", len(b))
	}
	return binary.BigEndian.Uint32(b)
}

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
