package gcp

import (
	"fmt"
	"log"
)

// A RpdCap is a RpdCapabilities TLV (Complex TLV).
type RpdCap struct{ TLV }

// Name returns the type name of a RpdCapabilities TLV.
func (t *RpdCap) Name() string { return "RpdCapabilities" }

// IsComplex returns whether a RpdCapabilities TLV is Complex or not.
func (t *RpdCap) IsComplex() bool { return true }

func (t *RpdCap) newTLV(b byte) RCP {
	switch int(b) {
	case 19:
		r := new(RpdIdf)
		r.parentMsg = t.parentMsg
		return r
	case 24:
		r := new(DevLoc)
		r.parentMsg = t.parentMsg
		return r
	default:
		log.Printf("RpdCapabilities TLV type: %d not supported", int(b))
		return nil
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
type RpdIdf struct{ TLV }

// Name returns the type name of a RpdIdentification TLV.
func (t *RpdIdf) Name() string { return "RpdIdentification" }

// IsComplex returns whether a RpdIdentification TLV is Complex or not.
func (t *RpdIdf) IsComplex() bool { return true }

func (t *RpdIdf) newTLV(b byte) RCP {
	switch int(b) {
	case 1:
		r := new(VendorName)
		r.parentMsg = t.parentMsg
		return r
	case 2:
		r := new(VendorID)
		r.parentMsg = t.parentMsg
		return r
	case 3:
		r := new(ModelNbr)
		r.parentMsg = t.parentMsg
		return r
	case 4:
		r := new(DevMacAddr)
		r.parentMsg = t.parentMsg
		return r
	case 5:
		r := new(CurSwVer)
		r.parentMsg = t.parentMsg
		return r
	case 6:
		r := new(BootVer)
		r.parentMsg = t.parentMsg
		return r
	case 7:
		r := new(DevDesc)
		r.parentMsg = t.parentMsg
		return r
	case 8:
		r := new(DevAlias)
		r.parentMsg = t.parentMsg
		return r
	case 9:
		r := new(SerialNum)
		r.parentMsg = t.parentMsg
		return r
	case 10:
		r := new(UsBurRecID)
		r.parentMsg = t.parentMsg
		return r
	case 11:
		r := new(UsBurRecMod)
		r.parentMsg = t.parentMsg
		return r
	case 12:
		r := new(UsBurRecDrv)
		r.parentMsg = t.parentMsg
		return r
	case 13:
		r := new(UsBurRecSN)
		r.parentMsg = t.parentMsg
		return r
	case 14:
		r := new(RpdRcpPrVer)
		r.parentMsg = t.parentMsg
		return r
	case 15:
		r := new(RpdRcpSchVer)
		r.parentMsg = t.parentMsg
		return r
	case 16:
		r := new(HwRev)
		r.parentMsg = t.parentMsg
		return r
	case 17:
		r := new(AsID)
		r.parentMsg = t.parentMsg
		return r
	case 18:
		r := new(VspSel)
		r.parentMsg = t.parentMsg
		return r
	case 19:
		r := new(CurSwUpd)
		r.parentMsg = t.parentMsg
		return r
	case 20:
		r := new(CurSwName)
		r.parentMsg = t.parentMsg
		return r
	case 21:
		r := new(CurSwSer)
		r.parentMsg = t.parentMsg
		return r
	case 22:
		r := new(CurSwIdx)
		r.parentMsg = t.parentMsg
		return r
	default:
		log.Printf("RpdIdentification TLV type: %d not supported", int(b))
		return nil
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
type VendorName struct{ TLV }

// Name returns the type name of a VendorName TLV.
func (t *VendorName) Name() string { return "VendorName" }

// Val returns the value a VendorName TLV carries.
// A string identifying the RPD's manufacturer.
func (t *VendorName) Val() interface{} {
	s := stringVal(t.Value)
	t.parentMsg.NTF.Sequence.RpdCapabilities.RpdIdentification.VendorName = s
	return s
}

// A VendorID is a VendorId TLV.
type VendorID struct{ TLV }

// Name returns the type name of a VendorId TLV.
func (t *VendorID) Name() string { return "VendorId" }

// Val returns the value a VendorId TLV carries.
// An unsigned short with Vendor Id of the RPD's manufacturer
func (t *VendorID) Val() interface{} {
	s := u16Val(t.Value)
	t.parentMsg.NTF.Sequence.RpdCapabilities.RpdIdentification.VendorID = s
	return s
}

// A ModelNbr is a ModelNumber TLV.
type ModelNbr struct{ TLV }

// Name returns the type name of a ModelNumber TLV.
func (t *ModelNbr) Name() string { return "ModelNumber" }

// Val returns the value a ModelNumber TLV carries.
// A string identifying the RPD's model number.
func (t *ModelNbr) Val() interface{} {
	s := stringVal(t.Value)
	t.parentMsg.NTF.Sequence.RpdCapabilities.RpdIdentification.ModelNumber = s
	return s
}

// A DevMacAddr is a DeviceMacAddress TLV.
type DevMacAddr struct{ TLV }

// Name returns the type name of a DeviceMacAddress TLV.
func (t *DevMacAddr) Name() string { return "DeviceMacAddress" }

// Val returns the value a DeviceMacAddress TLV carries.
// The MAC address used to uniquely identify the RPD.
func (t *DevMacAddr) Val() interface{} {
	s := macVal(t.Value)
	t.parentMsg.NTF.Sequence.RpdCapabilities.RpdIdentification.DeviceMacAddress = s
	return s
}

// A CurSwVer is a CurrentSwVersion TLV.
type CurSwVer struct{ TLV }

// Name returns the type name of a CurrentSwVersion TLV.
func (t *CurSwVer) Name() string { return "CurrentSwVersion" }

// Val returns the value a CurrentSwVersion TLV carries.
// A string representing the SW version currently running on of the RPD.
func (t *CurSwVer) Val() interface{} {
	s := stringVal(t.Value)
	t.parentMsg.NTF.Sequence.RpdCapabilities.RpdIdentification.CurrentSwVersion = s
	return s
}

// A BootVer is a BootRomVersion TLV.
type BootVer struct{ TLV }

// Name returns the type name of a BootRomVersion TLV.
func (t *BootVer) Name() string { return "BootRomVersion" }

// Val returns the value a BootRomVersion TLV carries.
// A string representing the BootRom version currently installed
// on of the RPD.
func (t *BootVer) Val() interface{} {
	s := stringVal(t.Value)
	t.parentMsg.NTF.Sequence.RpdCapabilities.RpdIdentification.BootRomVersion = s
	return s
}

// A DevDesc is a DeviceDescription TLV.
type DevDesc struct{ TLV }

// Name returns the type name of a DeviceDescription TLV.
func (t *DevDesc) Name() string { return "DeviceDescription" }

// Val returns the value a DeviceDescription TLV carries.
// A string selected by the RPD manufacturer.
func (t *DevDesc) Val() interface{} {
	s := stringVal(t.Value)
	t.parentMsg.NTF.Sequence.RpdCapabilities.RpdIdentification.DeviceDescription = s
	return s
}

// A DevAlias is a DeviceAlias TLV.
type DevAlias struct{ TLV }

// Name returns the type name of a DeviceAlias TLV.
func (t *DevAlias) Name() string { return "DeviceAlias" }

// Val returns the value a DeviceAlias TLV carries.
// A string communicating device's name assigned by the operator.
func (t *DevAlias) Val() interface{} {
	s := stringVal(t.Value)
	t.parentMsg.NTF.Sequence.RpdCapabilities.RpdIdentification.DeviceAlias = s
	return s
}

// A SerialNum is a SerialNumber TLV.
type SerialNum struct{ TLV }

// Name returns the type name of a SerialNumber TLV.
func (t *SerialNum) Name() string { return "SerialNumber" }

// Val returns the value a SerialNumber TLV carries.
// A string representing device's serial number.
func (t *SerialNum) Val() interface{} {
	s := stringVal(t.Value)
	t.parentMsg.NTF.Sequence.RpdCapabilities.RpdIdentification.SerialNumber = s
	return s
}

// A UsBurRecID is a UsBurstReceiverVendorId TLV.
type UsBurRecID struct{ TLV }

// Name returns the type name of a UsBurstReceiverVendorId TLV.
func (t *UsBurRecID) Name() string { return "UsBurstReceiverVendorId" }

// Val returns the value a UsBurstReceiverVendorId TLV carries.
// An unsigned 16-bit integer with the IANA Enterprise Code of
// the manufacturer of the RPD's US burst receiver.
func (t *UsBurRecID) Val() interface{} {
	s := u16Val(t.Value)
	t.parentMsg.NTF.Sequence.RpdCapabilities.RpdIdentification.UsBurstReceiverVendorID = s
	return s
}

// A UsBurRecMod is a UsBurstReceiverModelNumber TLV.
type UsBurRecMod struct{ TLV }

// Name returns the type name of a UsBurstReceiverModelNumber TLV.
func (t *UsBurRecMod) Name() string { return "UsBurstReceiverModelNumber" }

// Val returns the value a UsBurstReceiverModelNumber TLV carries.
// A string with the identifier of the model number of the RPD's US
// burst receiver. If not available from the vendor, report a zerolength string.
func (t *UsBurRecMod) Val() interface{} {
	// The length of this one is 0-16, not 0-255. Do I create a new stringVal func for this?
	s := stringVal(t.Value)
	t.parentMsg.NTF.Sequence.RpdCapabilities.RpdIdentification.UsBurstReceiverModelNumber = s
	return s
}

// A UsBurRecDrv is a UsBurstReceiverDriverVersion TLV.
type UsBurRecDrv struct{ TLV }

// Name returns the type name of a UsBurstReceiverDriverVersion TLV.
func (t *UsBurRecDrv) Name() string { return "UsBurstReceiverDriverVersion" }

// Val returns the value a UsBurstReceiverDriverVersion TLV carries.
// A string identifying the version of the driver of the RPD's US
// burst receiver. A zero-length string indicates the driver version
// is not available or not applicable.
func (t *UsBurRecDrv) Val() interface{} {
	// The length of this one is 0-16, not 0-255. Do I create a new stringVal func for this?
	s := stringVal(t.Value)
	t.parentMsg.NTF.Sequence.RpdCapabilities.RpdIdentification.UsBurstReceiverDriverVersion = s
	return s
}

// A UsBurRecSN is a UsBurstReceiverSerialNumber TLV.
type UsBurRecSN struct{ TLV }

// Name returns the type name of a UsBurstReceiverSerialNumber TLV.
func (t *UsBurRecSN) Name() string { return "UsBurstReceiverSerialNumber" }

// Val returns the value a UsBurstReceiverSerialNumber TLV carries.
// A string identifying the serial number of the RPD's US burst
// receiver. A zero-length string indicates the serial number is not
// available
func (t *UsBurRecSN) Val() interface{} {
	// The length of this one is 0-16, not 0-255. Do I create a new stringVal func for this?
	s := stringVal(t.Value)
	t.parentMsg.NTF.Sequence.RpdCapabilities.RpdIdentification.UsBurstReceiverSerialNumber = s
	return s
}

// A RpdRcpPrVer is a RpdRcpProtocolVersion TLV.
type RpdRcpPrVer struct{ TLV }

// Name returns the type name of a RpdRcpProtocolVersion TLV.
func (t *RpdRcpPrVer) Name() string { return "RpdRcpProtocolVersion" }

// Val returns the value a RpdRcpProtocolVersion TLV carries.
// A string identifying the RCP protocol version supported by the RPD.
func (t *RpdRcpPrVer) Val() interface{} {
	// The length of this one is 3-32, not 0-255. Do I create a new stringVal func for this?
	s := stringVal(t.Value)
	t.parentMsg.NTF.Sequence.RpdCapabilities.RpdIdentification.RpdRcpProtocolVersion = s
	return s
}

// A RpdRcpSchVer is a RpdRcpSchemaVersion TLV.
type RpdRcpSchVer struct{ TLV }

// Name returns the type name of a RpdRcpSchemaVersion TLV.
func (t *RpdRcpSchVer) Name() string { return "RpdRcpSchemaVersion" }

// Val returns the value a RpdRcpSchemaVersion TLV carries.
// A string identifying the RCP schema version supported by the RPD.
func (t *RpdRcpSchVer) Val() interface{} {
	// The length of this one is 5-32, not 0-255. Do I create a new stringVal func for this?
	s := stringVal(t.Value)
	t.parentMsg.NTF.Sequence.RpdCapabilities.RpdIdentification.RpdRcpSchemaVersion = s
	return s
}

// A HwRev is a HwRevision TLV.
type HwRev struct{ TLV }

// Name returns the type name of a HwRevision TLV.
func (t *HwRev) Name() string { return "HwRevision" }

// Val returns the value a HwRevision TLV carries.
// A string identifying the revision of the RPD hardware.
func (t *HwRev) Val() interface{} {
	s := stringVal(t.Value)
	t.parentMsg.NTF.Sequence.RpdCapabilities.RpdIdentification.HwRevision = s
	return s
}

// An AsID is a AssetId TLV.
type AsID struct{ TLV }

// Name returns the type name of a AssetId TLV.
func (t *AsID) Name() string { return "AssetId" }

// Val returns the value a AssetId TLV carries.
// A string containing asset identification of the RPD.
// The default value is the zero-length string or "".
func (t *AsID) Val() interface{} {
	// The length of this one is 0-32, not 0-255. Do I create a new stringVal func for this?
	s := stringVal(t.Value)
	t.parentMsg.NTF.Sequence.RpdCapabilities.RpdIdentification.AssetID = s
	return s
}

// A VspSel is a VspSelector TLV.
type VspSel struct{ TLV }

// Name returns the type name of a VspSelector TLV.
func (t *VspSel) Name() string { return "VspSelector" }

// Val returns the value a AssetId TLV carries.
// A string containing a VSP Selector. If the RPD does not support
// VSP the RPD communicates VSP as a zero-length string.
func (t *VspSel) Val() interface{} {
	// The length of this one is 0-16, not 0-255. Do I create a new stringVal func for this?
	s := stringVal(t.Value)
	t.parentMsg.NTF.Sequence.RpdCapabilities.RpdIdentification.VspSelector = s
	return s
}

// A CurSwUpd is a CurrentSwImageLastUpdate TLV.
type CurSwUpd struct{ TLV }

// Name returns the type name of a CurrentSwImageLastUpdate TLV.
func (t *CurSwUpd) Name() string { return "CurrentSwImageLastUpdate" }

// Val returns the value a CurrentSwImageLastUpdate TLV carries.
// An octet string conforming to the definition of DateAndTime from [RFC 2578].
func (t *CurSwUpd) Val() interface{} {
	s := timeRFC2579Val(t.Value)
	t.parentMsg.NTF.Sequence.RpdCapabilities.RpdIdentification.CurrentSwImageLastUpdate = s
	return s
}

// A CurSwName is a CurrentSwImageName TLV.
type CurSwName struct{ TLV }

// Name returns the type name of a CurrentSwImageName TLV.
func (t *CurSwName) Name() string { return "CurrentSwImageName" }

// Val returns the value a CurrentSwImageName TLV carries.
// A string with the name of the current SW image.
func (t *CurSwName) Val() interface{} {
	s := stringVal(t.Value)
	t.parentMsg.NTF.Sequence.RpdCapabilities.RpdIdentification.CurrentSwImageName = s
	return s
}

// A CurSwSer is a CurrentSwImageServer TLV.
type CurSwSer struct{ TLV }

// Name returns the type name of a CurrentSwImageServer TLV.
func (t *CurSwSer) Name() string { return "CurrentSwImageServer" }

// Val returns the value a CurrentSwImageServer TLV carries.
// The IP Address of the server from which the current SW image
// was downloaded.
func (t *CurSwSer) Val() interface{} {
	s := ipVal(t.Value)
	t.parentMsg.NTF.Sequence.RpdCapabilities.RpdIdentification.CurrentSwImageServer = s
	return s
}

// A CurSwIdx is a CurrrentSwImageIndex TLV.
type CurSwIdx struct{ TLV }

// Name returns the type name of a CurrrentSwImageIndex TLV.
func (t *CurSwIdx) Name() string { return "CurrrentSwImageIndex" }

// Val returns the value a CurrrentSwImageIndex TLV carries.
// An unsigned byte reporting which SW image is currently
// running on the RPD. The following range of values are
// permitted by this specification: 0..3.
// The value of zero is reserved for the main SW image.
func (t *CurSwIdx) Val() interface{} {
	s := u8Val(t.Value)
	t.parentMsg.NTF.Sequence.RpdCapabilities.RpdIdentification.CurrrentSwImageIndex = s
	return s
}

// A DevLoc is a Device Location TLV (Complex TLV).
type DevLoc struct {
	TLV
	// index identifies whether this is part of IRA(1), REX(2) or NTF(3).
	index uint8
}

// Name returns the type name of a Device Location TLV.
func (t *DevLoc) Name() string { return "Device Location" }

// IsComplex returns whether a Device Location TLV is Complex or not.
func (t *DevLoc) IsComplex() bool { return true }

func (t *DevLoc) newTLV(b byte) RCP {
	switch int(b) {
	case 1:
		r := new(DevLocDesc)
		r.parentMsg = t.parentMsg
		return r
	case 2:
		r := new(GeoLocLat)
		r.parentMsg = t.parentMsg
		return r
	case 3:
		r := new(GeoLocLon)
		r.parentMsg = t.parentMsg
		return r
	default:
		log.Printf("DeviceLocation TLV type: %d not supported", int(b))
		return nil
	}
}

// parseTLVs parses Device Location TLVs.
func (t *DevLoc) parseTLVs(b []byte) ([]RCP, error) {
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

// A DevLocDesc is a Device Location Description TLV.
type DevLocDesc struct{ TLV }

// Name returns the type name of a Device Location Description TLV.
func (t *DevLocDesc) Name() string { return "Device Location Description" }

// Val returns the value a Device Location Description TLV carries.
// A string with a short text description of where the RPD has
// been installed, such as a street address. The format is specific
// to the operator.
func (t *DevLocDesc) Val() interface{} {
	s := stringVal(t.Value)
	t.parentMsg.NTF.Sequence.RpdCapabilities.DeviceLocation.Description = s
	return s
}

// A GeoLocLat is a Geographic Location Latitude TLV.
type GeoLocLat struct{ TLV }

// Name returns the type name of a Geographic Location Latitude TLV.
func (t *GeoLocLat) Name() string { return "Geographic Location Latitude" }

// Val returns the value a Geographic Location Latitude TLV carries.
// A 9 byte long string with RPD's latitude formatted as in ISO
// 6709-2008. The RPD uses "6 digit notation" in the format deg,
// min, sec, ±DDMMSS.S. example: -750015.1
func (t *GeoLocLat) Val() interface{} {
	if len(t.Value) != 9 {
		return fmt.Sprintf("unexpected lenght: %v, want: 9", len(t.Value))
	}
	s := stringVal(t.Value)
	// TODO: Parse ISO 6709-2008
	t.parentMsg.NTF.Sequence.RpdCapabilities.DeviceLocation.Latitude = s
	return s
}

// A GeoLocLon is a Geographic Location Longitude TLV.
type GeoLocLon struct{ TLV }

// Name returns the type name of a Geographic Location Longitude TLV.
func (t *GeoLocLon) Name() string { return "Geographic Location Longitude" }

// Val returns the value a Geographic Location Longitude TLV carries.
// A 10 byte long string with RPD's latitude formatted as in ISO
// 6709-2008. The RPD uses "7 digit notation" in the format deg,
// min, sec, ±DDDMMSS.S.
// example: -0100015.1
func (t *GeoLocLon) Val() interface{} {
	if len(t.Value) != 10 {
		return fmt.Sprintf("unexpected lenght: %v, want: 9", len(t.Value))
	}
	s := stringVal(t.Value)
	// TODO: Parse ISO 6709-2008
	t.parentMsg.NTF.Sequence.RpdCapabilities.DeviceLocation.Longitude = s
	return s
}
