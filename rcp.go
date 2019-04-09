package gcp

import (
	"encoding/binary"
	"errors"
	"fmt"
)

// A TLV represents a RCP TLV message.
type TLV struct {
	Type   uint8  // Type: 1 byte
	Length uint16 // Value Length: 2 bytes
	Value  []byte
}

// An RCP encodes a TLV used in the RCP protocol.
type RCP interface {
	Name() uint8
	marshal() ([]byte, error)
	unmarshal(b []byte) error
}

// Error messages
var (
	ErrUnexpectedEOF = errors.New("unexpected EOF")
)

// Name ...
func (t *TLV) Name() byte { return t.Type }

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

// parseTLVs ...
func parseTLVs(b []byte) ([]TLV, error) {
	var tlvs []TLV
	for i := 0; len(b[i:]) != 0; {
		// Three bytes: TLV type and TLV length.
		if len(b[i:]) < 3 {
			return nil, ErrUnexpectedEOF
		}

		//ty := b[i]
		l := int(binary.BigEndian.Uint16(b[i+1 : i+3]))

		// Verify that we won't advance beyond the end of the byte slice.
		if l > len(b[i+3:]) {
			return nil, ErrUnexpectedEOF
		}

		tlv := TLV{}

		// Unmarshal at the current offset, up to the expected length.
		if err := tlv.unmarshal(b[i : i+3+l]); err != nil {
			return nil, err
		}

		switch {
		// Complex TLV
		case l > 3 && tlv.isComplex():
			// Recursive call
			rectlv, err := parseTLVs(b[i+3 : i+3+l])
			if err != nil {
				return nil, err
			}
			tlvs = append(tlvs, tlv)
			tlvs = append(tlvs, rectlv...)
		case l < 3 || !tlv.isComplex():
			tlvs = append(tlvs, tlv)
		default:
			fmt.Printf("We really shouldn't get here: TLV Type %d\n", tlv.Type)
		}
		// Advance to the next TLV's type field.
		i += (l + 3)
	}

	return tlvs, nil

}

func (t TLV) isComplex() bool {
	tp := RCPType(t.Type)
	switch tp {
	// Complex TLV
	case IRA, REX, NTF, Seq, RfPortSel, RpdCap, GnrlNtf, RpdInfo:
		return true
	case SeqNmr, Oper:
		return false
	default:
		return false
	}
}

// RCPType is a RCP TLV Type.
type RCPType uint8

// RCP TLV Types
const (
	IRA       RCPType = 1   // Complex TLV - IRA Message
	REX       RCPType = 2   // Complex TLV - REX Message
	NTF       RCPType = 3   // Complex TLV - NTF Message
	Seq       RCPType = 9   // Complex TLV - Sequence
	SeqNmr    RCPType = 10  // UnsignedShort - Sequence Number
	Oper      RCPType = 11  // UnsignedByte - Operation
	RfPortSel RCPType = 13  // Complex TLV - RfPortSelector
	RpdCap    RCPType = 50  // Complex TLV - RpdCapabilities
	GnrlNtf   RCPType = 86  // Complex TLV - GeneralNotification
	RpdInfo   RCPType = 100 // Complex TLV - RpdInfo
)

// RpdIdentification (19) will need different TLV processing
