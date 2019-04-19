package gcp

import (
	"fmt"
	"log"
)

// A GenrlNtf is a GeneralNotification TLV (Complex TLV).
type GenrlNtf struct{ TLV }

// Name returns the type name of a GeneralNotification TLV.
func (t *GenrlNtf) Name() string { return "GeneralNotification" }

// IsComplex returns whether a GeneralNotification TLV is Complex or not.
func (t *GenrlNtf) IsComplex() bool { return true }

func (t *GenrlNtf) newTLV(b byte) RCP {
	switch int(b) {
	case 1:
		r := new(NtfType)
		r.parentMsg = t.parentMsg
		return r
	default:
		log.Printf("GeneralNotification TLV type: %d not supported", int(b))
		return nil
	}
}

// parseTLVs parses GeneralNotification TLVs.
func (t *GenrlNtf) parseTLVs(b []byte) ([]RCP, error) {
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

// A NtfType is a NotificationType TLV.
type NtfType struct{ TLV }

// Name returns the type name of a NotificationType TLV.
func (t *NtfType) Name() string { return "NotificationType" }

// Val returns the value a NotificationType TLV carries.
func (t *NtfType) Val() interface{} {
	if len(t.Value) != 1 {
		return fmt.Errorf("unexpected lenght: %v, want: 1", len(t.Value))
	}
	var s string
	switch int(t.Value[0]) {
	case 1:
		s = "StartUpNotification"
	case 2:
		s = "RedirectResultNotification"
	case 3:
		s = "PtpResultNotification"
	case 4:
		s = "AuxCoreResultNotification"
	case 5:
		s = "TimeOutNotification"
	case 7:
		s = "ReconnectNotification"
	case 8:
		s = "AuxCoreGcpStatusNotification"
	case 9:
		s = "ChannelUcdRefreshRequest"
	case 10:
		s = "HandoverNotification"
	case 11:
		s = "SsdFailureNotification"
	default:
	}
	t.parentMsg.NTF.Sequence.GeneralNtf.NotificationType = s
	return s
}

// IsComplex returns whether a NotificationType TLV is Complex or not.
func (t *NtfType) IsComplex() bool { return false }
