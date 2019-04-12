package gcp

import (
	"fmt"
)

// A GenrlNtf is a GeneralNotification TLV (Complex TLV).
type GenrlNtf struct {
	TLV
}

// Name returns the type name of a GeneralNotification TLV.
func (t *GenrlNtf) Name() string { return "GeneralNotification" }

// IsComplex returns whether a GeneralNotification TLV is Complex or not.
func (t *GenrlNtf) IsComplex() bool { return true }

func (t *GenrlNtf) newTLV(b byte) RCP {
	switch int(b) {
	case 1:
		return new(NtfType)
	default:
		return new(TLV)
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
		// Complex TLV
		case l > 3 && tlv.IsComplex():
			// Recursive call
			rectlv, err := tlv.parseTLVs(b[i+3 : i+3+l])
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

// A NtfType is a NotificationType TLV.
type NtfType struct {
	TLV
}

// Name returns the type name of a NotificationType TLV.
func (t *NtfType) Name() string { return "NotificationType" }

// Val returns the value a NotificationType TLV carries.
func (t *NtfType) Val() interface{} {
	if len(t.Value) != 1 {
		return fmt.Errorf("unexpected lenght: %v, want: 1", len(t.Value))
	}
	switch int(t.Value[0]) {
	case 1:
		return "StartUpNotification"
	case 2:
		return "RedirectResultNotification"
	case 3:
		return "PtpResultNotification"
	case 4:
		return "AuxCoreResultNotification"
	case 5:
		return "TimeOutNotification"
	case 7:
		return "ReconnectNotification"
	case 8:
		return "AuxCoreGcpStatusNotification"
	case 9:
		return "ChannelUcdRefreshRequest"
	case 10:
		return "HandoverNotification"
	case 11:
		return "SsdFailureNotification"
	default:
		return "Unknown Notification"
	}
}

// IsComplex returns whether a NotificationType TLV is Complex or not.
func (t *NtfType) IsComplex() bool { return false }
