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
	RpdCapabilities *RpdC  `json:"RPD Capabilities,omitempty"`
	ResponseCode    string `json:"Response Code,omitempty"`
	RpdRedirect     *RpdR  `json:"RPD Redirect,omitempty"`
	GeneralNtf      *GNtf  `json:"General Notification,omitempty"`
	RpdInfo         *RpdI  `json:"RPD Info,omitempty"`
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

// A RpdR represents a RpdRedirect data structure.
// This TLV is used to communicate an ordered list of CCAP Cores to which
// the RPD is redirected.
type RpdR struct {
	// This TLV communicates an IPv4 address of CCAP Core to which the RPD
	// is redirected.
	RpdRedirectIPAddress []string `json:"IP Address,omitempty"`
}

// A GNtf represents a GeneralNotification data structure.
// GeneralNotification is a complex TLV used by the RPD to report events
// to the CCAP Core.
type GNtf struct {
	// NotificationType indicates the specific notification being sent
	// by the RPD.
	NotificationType string `json:"Type,omitempty"`
}

// A RpdI represents a RpdInfo data structure.
type RpdI struct {
	// This object provides details about the Ethernet interfaces on the RPD.
	// The attributes of this object are based on the ifTable/ifXTable specified
	// in [RFC 2863].
	IfEnet []IfEn `json:"IfEnet,omitempty"`
	// This object contains addressing information relevant to the RPD's interfaces.
	IPAddress []IPAdd `json:"IpAddress,omitempty"`
}

// A IfEn represents an IfEnet data structure.
type IfEn struct {
	// This key attribute reports a unique index for this Ethernet port interface.
	EnetPortIndex string `json:"Port Index,omitempty"`
	// This attribute reports a textual string representing a name that describes
	// the interface.
	Name string `json:"Name,omitempty"`
	// This attribute reports a textual string containing information about the
	// Ethernet interface.
	Descr string `json:"Description,omitempty"`
	// This attribute reports the type of interface. Additional values for Type
	// are assigned by the Internet Assigned Numbers Authority (IANA), through
	// updating the syntax of the IANAifType textual convention.
	Type string `json:"Type,omitempty"`
	// This attribute reports an Alias for the interface. On the first instantiation
	// of an interface, the value of Alias associated with that interface is the
	// zero-length string.
	Alias string `json:"Alias,omitempty"`
	// This attribute reports the size of the largest packet that can be sent/received
	// on the interface, specified in octets.
	MTU string `json:"MTU,omitempty"`
	// This attribute reports the interface's address at its protocol sub-layer.
	// For example, for an 802.x interface, this attribute normally
	// contains a MAC address.
	PhysAddress string `json:"Physical Address,omitempty"`
	// This attribute reports the state of the interface. The testing(3) state
	// indicates that no operational packets can be passed. When a managed system
	// initializes, all interfaces start with AdminStatus in the down(2) state.
	AdminStatus string `json:"Admin State,omitempty"`
	// This attribute reports the current operational state of the interface.
	// The testing(3) state indicates that no operational packets can be passed.
	// If AdminStatus is down(2) then OperStatus should be down(2).
	OperStatus string `json:"Operational State,omitempty"`
	// This attribute reports the value of RpdSysUpTime at the time the interface
	// entered its current operational state.
	LastChange string `json:"Last Change,omitempty"`
	// This attribute reports an estimate of the interface's current bandwidth
	// in units of 1,000,000 bits per second.
	HighSpeed string `json:"Bandwidth,omitempty"`
	// This attribute reports whether linkup/linkdown traps are generated for this
	// interface. A value of '1' indicates that traps are enabled.
	LinkUpDownTrapEnable string `json:"LinkUpDownTrapEnable,omitempty"`
	// This attribute reports a value of '2' (false) if this interface only accepts
	// packets/frames that are addressed to this interface. This attribute reports a
	// value of '1' (true) when the station accepts all packets/frames transmitted
	// on the media.
	PromiscuousMode string `json:"PromiscuousMode,omitempty"`
	// This attribute reports the value 'true' if the interface sublayer has a physical
	// connector and the value 'false' otherwise.
	ConnectorPresent bool `json:"Connector Present,omitempty"`
	// This attribute reports the network authentication status of this interface.
	NetworkAuthStatus string `json:"Network Auth Status,omitempty"`
}

// A IPAdd represents an IPAddress data structure.
type IPAdd struct {
	// This key attribute reports the IP address type of the IpAddress attribute.
	AddrType string `json:"Address Type,omitempty"`
	// This key attribute reports the IP address to which this entry's addressing
	// information pertains
	IPAddress string `json:"IP Address,omitempty"`
	// This attribute reports a unique index for this Ethernet port interface.
	EnetPortIndex string `json:"Port Index,omitempty"`
	// This attribute reports the type of traffic for which the address can be used.
	Type string `json:"Type,omitempty"`
	// This attribute reports the prefix length associated with this address.
	PrefixLen string `json:"Prefix Length,omitempty"`
	// This attribute reports the origin of this IP address.
	// 'manual' indicates an IP address that was manually configured.
	// 'dhcp' indicates an IP address that was assigned by a DHCP server.
	Origin string `json:"Origin,omitempty"`
	// This attribute reports the status of an address. Most of the states
	// correspond to states from the IPv6 Stateless Address Autoconfiguration protocol.
	Status string `json:"Status,omitempty"`
	// This attribute reports the value of RpdSysUpTime at the time this entry
	// was created. If this entry was created prior to the last re-initialization
	// of the local network management subsystem, then this attribute contains a zero value.
	Created string `json:"Created,omitempty"`
	// This attribute reports the value of RpdSysUpTime at the time this entry
	// was last updated. If this entry was updated prior to the last re-initialization
	// of the local network management subsystem, then this attribute contains a zero value.
	LastChanged string `json:"Last Changed,omitempty"`
}
