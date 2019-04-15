package gcp

// A RpdCap is a RpdCapabilities TLV (Complex TLV).
type RpdCap struct {
	TLV
}

// Name returns the type name of a RpdCapabilities TLV.
func (t *RpdCap) Name() string { return "RpdCapabilities" }

// IsComplex returns whether a RpdCapabilities TLV is Complex or not.
func (t *RpdCap) IsComplex() bool { return true }

func (t *RpdCap) newTLV(b byte) RCP {
	switch int(b) {
	case 19:
		return new(RpdIdf)
	default:
		return new(TLV)
	}
}

// parseTLVs parses RpdCapabilities TLVs.
func (t *RpdCap) parseTLVs(b []byte) ([]RCP, error) {
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

// A RpdIdf is a RpdIdentification TLV (Complex TLV).
type RpdIdf struct {
	TLV
}

// Name returns the type name of a RpdIdentification TLV.
func (t *RpdIdf) Name() string { return "RpdIdentification" }

// IsComplex returns whether a RpdIdentification TLV is Complex or not.
func (t *RpdIdf) IsComplex() bool { return true }

func (t *RpdIdf) newTLV(b byte) RCP {
	switch int(b) {
	case 1:
		return new(VendorName)
	case 2:
		return new(VendorID)
	case 3:
		return new(ModelNbr)
	case 4:
		return new(DevMacAddr)
	case 6:
		return new(BootVer)
	default:
		return new(TLV)
	}
}

// parseTLVs parses RpdCapabilities TLVs.
func (t *RpdIdf) parseTLVs(b []byte) ([]RCP, error) {
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

// A VendorName is a VendorName TLV.
type VendorName struct {
	TLV
}

// Name returns the type name of a VendorName TLV.
func (t *VendorName) Name() string { return "VendorName" }

// IsComplex returns whether a VendorName TLV is Complex or not.
func (t *VendorName) IsComplex() bool { return false }

// Val returns the value a VendorName TLV carries.
func (t *VendorName) Val() interface{} { return stringVal(t.Value) }

// A VendorID is a VendorId TLV.
type VendorID struct {
	TLV
}

// Name returns the type name of a VendorId TLV.
func (t *VendorID) Name() string { return "VendorId" }

// IsComplex returns whether a VendorId TLV is Complex or not.
func (t *VendorID) IsComplex() bool { return false }

// Val returns the value a VendorId TLV carries.
func (t *VendorID) Val() interface{} { return u16Val(t.Value) }

// A ModelNbr is a ModelNumber TLV.
type ModelNbr struct {
	TLV
}

// Name returns the type name of a ModelNumber TLV.
func (t *ModelNbr) Name() string { return "ModelNumber" }

// IsComplex returns whether a ModelNumber TLV is Complex or not.
func (t *ModelNbr) IsComplex() bool { return false }

// Val returns the value a ModelNumber TLV carries.
func (t *ModelNbr) Val() interface{} { return stringVal(t.Value) }

// A DevMacAddr is a DeviceMacAddress TLV.
type DevMacAddr struct {
	TLV
}

// Name returns the type name of a DeviceMacAddress TLV.
func (t *DevMacAddr) Name() string { return "DeviceMacAddress" }

// IsComplex returns whether a DeviceMacAddress TLV is Complex or not.
func (t *DevMacAddr) IsComplex() bool { return false }

// Val returns the value a DeviceMacAddress TLV carries.
func (t *DevMacAddr) Val() interface{} { return macVal(t.Value) }

// A BootVer is a BootRomVersion TLV.
type BootVer struct {
	TLV
}

// Name returns the type name of a BootRomVersion TLV.
func (t *BootVer) Name() string { return "BootRomVersion" }

// IsComplex returns whether a BootRomVersion TLV is Complex or not.
func (t *BootVer) IsComplex() bool { return false }

// Val returns the value a BootRomVersion TLV carries.
func (t *BootVer) Val() interface{} { return stringVal(t.Value) }
