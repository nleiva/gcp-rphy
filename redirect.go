package gcp

import (
	"log"
)

// A RpdRed is a RpdRedirect TLV (Complex TLV).
type RpdRed struct {
	TLV
	IPindex int8
}

// Name returns the type name of a RpdRedirect TLV.
func (t *RpdRed) Name() string { return "RpdRedirect" }

// IsComplex returns whether a RpdRedirect TLV is Complex or not.
func (t *RpdRed) IsComplex() bool { return true }

func (t *RpdRed) newTLV(b byte) RCP {
	switch int(b) {
	case 1:
		r := new(RedIPAdd)
		r.parentMsg = t.parentMsg
		r.IPindex = t.IPindex
		return r
	default:
		log.Printf("RpdRedirect TLV type: %d not supported", int(b))
		return nil
	}
}

// parseTLVs parses RpdRedirect TLVs.
func (t *RpdRed) parseTLVs(b []byte) ([]RCP, error) {
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

// A RedIPAdd is a RedirectIpAddress TLV.
type RedIPAdd struct {
	TLV
	IPindex int8
}

// Name returns the type name of aRedirectIpAddress TLV.
func (t *RedIPAdd) Name() string { return "RedirectIpAddress" }

// Val returns the value a RedirectIpAddress TLV carries.
func (t *RedIPAdd) Val() interface{} {
	s := ipVal(t.Value)
	t.parentMsg.IRA.Sequence.RpdRedirect.RpdRedirectIPAddress = append(t.parentMsg.IRA.Sequence.RpdRedirect.RpdRedirectIPAddress,
		s,
	)
	// t.parentMsg.IRA.Sequence.RpdRedirect.RpdRedirectIPAddress[t.IPindex] = s
	return s
}

// IsComplex returns whether a RedirectIpAddress TLV is Complex or not.
func (t *RedIPAdd) IsComplex() bool { return false }
