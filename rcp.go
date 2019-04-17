package gcp

import (
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"strconv"
	"time"
)

// A TLV represents a RCP TLV message.
type TLV struct {
	Type      uint8  // Type: 1 byte
	Length    uint16 // Value Length: 2 bytes
	Value     []byte
	parentMsg *GCP
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
		case l > 3 && tlv.IsComplex():
			rectlv, err := tlv.parseTLVs(b[i+3 : i+3+l])
			if err != nil {
				return nil, err
			}
			tlvs = append(tlvs, tlv)
			tlvs = append(tlvs, rectlv...)
		case l <= 3 || !tlv.IsComplex():
			tlvs = append(tlvs, tlv)
		}
		// Advance to the next TLV's type field.
		i += (l + 3)
	}

	return tlvs, nil
}

func (t *TLV) newTLV(b byte) RCP {
	switch int(b) {
	case 1:
		r := new(IRA)
		r.parentMsg = t.parentMsg
		return r
	case 2:
		r := new(REX)
		r.parentMsg = t.parentMsg
		return r
	case 3:
		r := new(NTF)
		r.parentMsg = t.parentMsg
		return r
	case 9:
		r := new(Seq)
		r.parentMsg = t.parentMsg
		return r
	case 10:
		r := new(SeqNmr)
		r.parentMsg = t.parentMsg
		return r
	case 11:
		r := new(Oper)
		r.parentMsg = t.parentMsg
		return r
	case 50:
		r := new(RpdCap)
		r.parentMsg = t.parentMsg
		return r
	case 86:
		r := new(GenrlNtf)
		r.parentMsg = t.parentMsg
		return r
	case 100:
		r := new(RpdInfo)
		r.parentMsg = t.parentMsg
		return r
	default:
		r := new(TLV)
		r.parentMsg = t.parentMsg
		return r
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
func (t *SeqNmr) Val() interface{} {
	if t.parentMsg != nil {
		// Might not be a NTF, B=but IRA or REX
		t.parentMsg.NTF.Sequence.SequenceNumber = u16Val(t.Value)
	} else {
		fmt.Printf("\n***\n***\nEMPTY SeqNmr!\n***\n***\n")
	}
	return u16Val(t.Value)
}

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
	s := "Unknown Operation"

	switch int(t.Value[0]) {
	case 1:
		s = "Read"
	case 2:
		s = "Write"
	case 3:
		s = "Delete"
	case 4:
		s = "ReadResponse"
	case 5:
		s = "WriteResponse"
	case 6:
		s = "DeleteResponse"
	case 7:
		s = "AllocateWrite"
	case 8:
		s = "AllocateWriteResponse"
	default:
	}
	if t.parentMsg != nil {
		// Might not be a NTF, B=but IRA or REX
		t.parentMsg.NTF.Sequence.Operation = s
	} else {
		fmt.Printf("\n***\n***\nEMPTY SeqNmr!\n***\n***\n")
	}
	return s
}

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

func u8Val(b []byte) string {
	if len(b) != 1 {
		return fmt.Sprintf("unexpected lenght: %v, want: 1", len(b))
	}
	if b[0] == 0 {
		return "0"
	}
	return string(b[0])
}

func u16Val(b []byte) string {
	if len(b) != 2 {
		return fmt.Sprintf("unexpected lenght: %v, want: 2", len(b))
	}
	v := binary.BigEndian.Uint16(b)
	return strconv.FormatUint(uint64(v), 10)
}

func u32Val(b []byte) interface{} {
	// TODO: Should I stick to uint32 and return "0" if there is an error?
	if len(b) != 4 {
		return fmt.Sprintf("unexpected lenght: %v, want: 4", len(b))
	}
	return binary.BigEndian.Uint32(b)
}

func timeVal(b []byte) string {
	if len(b) != 4 {
		return fmt.Sprintf("unexpected lenght: %v, want: 4", len(b))
	}
	t := binary.BigEndian.Uint32(b)
	if t == 0 {
		return "0"
	}
	// We receive a hundredths of a second
	return time.Unix(int64(t)*100, 0).String()
}

func macVal(b []byte) string {
	if len(b) != 6 {
		return fmt.Sprintf("unexpected lenght: %v, want: 6", len(b))
	}
	return net.HardwareAddr(b).String()
}

func ipVal(b []byte) string {
	l := len(b)
	if l != 4 && l != 16 {
		return fmt.Sprintf("unexpected lenght: %v, want: 4 or 16", l)
	}
	return net.IP(b).String()
}

func timeRFC2579Val(b []byte) string {
	l := len(b)
	if l != 8 && l != 11 {
		return fmt.Sprintf("unexpected lenght: %v, want: 8 or 11", l)
	}
	year := binary.BigEndian.Uint16(b[0:2])
	month := uint8(b[2])
	day := uint8(b[3])
	hour := uint8(b[4])
	min := uint8(b[5])
	sec := uint8(b[6])
	// deci-seconds 0..9
	dsec := uint8(b[7])

	var dir, offHour, offMin uint8
	if l == 11 {
		// direction from UTC '+' / '-'
		dir = uint8(b[8])
		offHour = uint8(b[9])
		offMin = uint8(b[10])
	}

	var utcOffset time.Duration
	utcOffset += time.Duration(offHour) * time.Hour
	utcOffset += time.Duration(offMin) * time.Minute
	var loc *time.Location
	if dir == '-' {
		loc = time.FixedZone("", -int(utcOffset.Seconds()))
	} else {
		loc = time.FixedZone("", int(utcOffset.Seconds()))
	}

	nsec := int(dsec) * 100 * int(time.Millisecond)
	t := time.Date(int(year), time.Month(month), int(day), int(hour), int(min), int(sec), nsec, loc)

	return t.String()
}
