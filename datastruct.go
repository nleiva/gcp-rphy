package gcp

// A GCP represents a GCP data structure.
type GCP struct {
	IRA *dSeq `json:"IRA,omitempty"` // Identification and Resource Advertising
	REX *dSeq `json:"REX,omitempty"` // RCP Object Exchange
	NTF *dSeq `json:"NTF,omitempty"` // Notify
}

// A dSeq represents a URA, REX or NTF data structure.
type dSeq struct {
	Sequence Sequence `json:"Sequence,omitempty"`
}

// A Sequence represents a Sequence data structure.
type Sequence struct {
	SequenceNumber  string `json:"Sequence Number,omitempty"`
	Operation       string `json:"Operation,omitempty"`
	RpdCapabilities RpdC   `json:"RPD Capabilities,omitempty"`
}

// A RpdC represents a RpdCapabilities data structure.
type RpdC struct {
	// This object represents the total number of downstream
	// RF ports supported by the RPD.
	NumDsRfPorts string `json:"NumDsRfPorts,omitempty"`
	// This object represents the total number of upstream
	// RF ports supported by the RPD.
	NumUsRfPorts string `json:"NumUsRfPorts,omitempty"`
	// This object represents the total number of 10 Gigabit
	// Ethernet ports supported by the RPD.
	NumTenGeNsPorts string `json:"NumTenGeNsPorts,omitempty"`
	// This object represents the total number of 1 Gigabit
	// Ethernet ports supported by the RPD.
	NumOneGeNsPorts string `json:"NumOneGeNsPorts,omitempty"`
	//...

	// A complex TLV through which the RPD communicates a
	// set of identifying parameters.
	RpdIdentification RpdIden `json:"RpdIdentification,omitempty"`

	// This TLV allows the RPD to inform the CCAP Core about it its location.
	DeviceLocation DeLoc `json:"Device Location,omitempty"`
}

// A RpdIden represents a RpdCapabilities data structure.
type RpdIden struct {
	// The VendorName object identifies the RPD's manufacturer.
	// The detailed format is vendor proprietary.
	VendorName string `json:"VendorName,omitempty"`
	// This TLV communicates the RPD's manufacturer's vendor id
	// as the IANA-assigned "SMI Network Management Private Enterprise
	// Codes" [RFC 1700] value.
	VendorID string `json:"VendorId,omitempty"`
	// This TLV convey the model name and number assigned to the RPD.
	// The format of the string is vendor proprietary.
	ModelNumber string `json:"ModelNumber,omitempty"`
	// This TLV convey the main MAC address of the RPD. Typically the
	// MAC address associated with the lowest numbered CIN facing Ethernet port.
	DeviceMacAddress string `json:"DeviceMacAddress,omitempty"`
	// This TLV conveys the SW version currently running on of the RPD.
	// The format of the string is vendor proprietary.
	CurrentSwVersion string `json:"CurrentSwVersion,omitempty"`
	// This TLV conveys the version of the BootRom currently installed
	// on of the RPD. The format of the string is vendor proprietary.
	BootRomVersion string `json:"BootRomVersion,omitempty"`
	// This TLV conveys a short description of the RPD in the form a string,
	// selected by the RPD's manufacturer.
	DeviceDescription string `json:"DeviceDescription,omitempty"`
	// This TLV communicates a device name assigned by the operator via
	// management interface. This object is an 'alias' name for the device
	// as specified by a network manager, and provides a non-volatile 'handle'
	// for the RPD.
	DeviceAlias string `json:"DeviceAlias,omitempty"`
	// This TLV communicates RPD's serial number. The format of the string
	// is vendor proprietary.
	SerialNumber string `json:"SerialNumber,omitempty"`
	// This TLV is used to communicate the identifier of the manufacturer of RPD's
	// US burst receiver as the IANAassigned value as defined in "SMI Network
	// Management Private Enterprise Codes"[RFC 1700].
	UsBurstReceiverVendorID string `json:"UsBurstReceiverVendorId,omitempty"`
	// This TLV is used to communicate the model number identifying RPD's US burst
	// receiver. The US burst receiver manufacturer is expected to specify the value
	// to be used in documentation provided to CCAP Core and RPD vendors.
	UsBurstReceiverModelNumber string `json:"UsBurstReceiverModelNumber,omitempty"`
	// This TLV is used to communicate the version of driver software supplied by the
	// RPD’s UsBurstReceiver vendor, if any. The US burst receiver manufacturer is
	// expected to specify the value to be used in documentation provided to CCAP
	// Core and RPD vendors.
	UsBurstReceiverDriverVersion string `json:"UsBurstReceiverDriverVersion,omitempty"`
	// This TLV is used to communicate the serial number of RPD's US burst receiver,
	// if any. Note that this value is not the RPD serial number as reported in
	// TLV 50.19.9.
	UsBurstReceiverSerialNumber string `json:"UsBurstReceiverSerialNumber,omitempty"`
	// This TLV is used to communicate the version of the RCP protocol supported
	// by the RPD.
	RpdRcpProtocolVersion string `json:"RpdRcpProtocolVersion,omitempty"`
	// This TLV is used to communicate the version of the RCP schema supported by
	// the RPD.
	RpdRcpSchemaVersion string `json:"RpdRcpSchemaVersion,omitempty"`
	// This TLV is used to communicate the revision of the RPD hardware.
	HwRevision string `json:"HwRevision,omitempty"`
	// This attribute is modeled after entPhysicalAssetID object defined in RFC 6933.
	// AssetId is used to communicate the asset tracking identifier as assigned by a
	// network manager.
	AssetID string `json:"AssetId,omitempty"`
	// The RPD advertises VspSelector (VSP stands for Vendor-Specific Pre-configuration)
	// in the form of human readable.
	VspSelector string `json:"VspSelector,omitempty"`
	// This attribute reports the date and time when the software image currently running
	// on the RPD was successfully updated. The RPD preserves the value of this attribute
	// across hardReset and softReset.
	CurrentSwImageLastUpdate string `json:"CurrentSwImageLastUpdate,omitempty"`
	// This attribute reports the name of the software image currently running on the RPD.
	// The RPD preserves the value of this attribute across reboots.
	CurrentSwImageName string `json:"CurrentSwImageName,omitempty"`
	// This attribute reports the Internet address of the server from which the software
	// image currently running on the RPD was downloaded.
	CurrentSwImageServer string `json:"CurrentSwImageServer,omitempty"`
	// This attribute reports which software image is currently running on the RPD.
	// An RPD which supports only one SW image always reports 0.
	CurrrentSwImageIndex string `json:"CurrrentSwImageIndex,omitempty"`
}

// A DeLoc represents a Device Location data structure.
type DeLoc struct {
	// This object allows the RPD to inform the CCAP Core about it its location.
	// The format of the information is specific to a cable operator.
	Description string `json:"Device Location Description,omitempty"`
	// This object allows the RPD to inform the CCAP Core about the latitude
	// portion of its geographic location.
	Latitude string `json:"Geographic Location Latitude,omitempty"`
	// This object allows the RPD to inform the CCAP Core about the longitude
	// portion of its geographic location.
	Longitude string `json:"Geographic Location Longitude,omitempty"`
}
