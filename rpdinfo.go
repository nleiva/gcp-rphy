package gcp

import (
	"fmt"
	"log"
)

// A RpdInfo is a RpdInfo TLV (Complex TLV).
type RpdInfo struct {
	TLV
	Ifindex int8
	IPindex int8
}

// Name returns the type name of a RpdInfo TLV.
func (t *RpdInfo) Name() string { return "RpdInfo" }

// IsComplex returns whether a RpdInfo TLV is Complex or not.
func (t *RpdInfo) IsComplex() bool { return true }

func (t *RpdInfo) newTLV(b byte) RCP {
	switch int(b) {
	case 8:
		r := new(IfEnet)
		t.Ifindex++
		r.parentMsg = t.parentMsg
		r.portIndex = t.Ifindex
		return r
	case 15:
		r := new(IPAddress)
		t.IPindex++
		r.parentMsg = t.parentMsg
		r.portIndex = t.IPindex
		return r
	default:
		log.Printf("RpdInfo TLV type: %d not supported", int(b))
		return nil
	}
}

// parseTLVs parses RpdInfo TLVs.
func (t *RpdInfo) parseTLVs(b []byte) ([]RCP, error) {
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

// A IfEnet is an IfEnet TLV.
type IfEnet struct {
	TLV
	// Fake port index to creat an array of interfaces
	portIndex int8
}

// Name returns the type name of an IfEnet TLV.
func (t *IfEnet) Name() string { return "IfEnet" }

// IsComplex returns whether a IfEnet TLV is Complex or not.
func (t *IfEnet) IsComplex() bool { return true }

func (t *IfEnet) newTLV(b byte) RCP {
	switch int(b) {
	case 1:
		r := new(EnPortIdx)
		r.parentMsg = t.parentMsg
		r.portIndex = t.portIndex
		t.parentMsg.REX.Sequence.RpdInfo.IfEnet = append(t.parentMsg.REX.Sequence.RpdInfo.IfEnet,
			IfEn{},
		)
		return r
	case 2:
		r := new(IfName)
		r.parentMsg = t.parentMsg
		r.portIndex = t.portIndex
		return r
	case 3:
		r := new(Descr)
		r.parentMsg = t.parentMsg
		r.portIndex = t.portIndex
		return r
	case 4:
		r := new(IifType)
		r.parentMsg = t.parentMsg
		r.portIndex = t.portIndex
		return r
	case 5:
		r := new(Alias)
		r.parentMsg = t.parentMsg
		r.portIndex = t.portIndex
		return r
	case 6:
		r := new(Mtu)
		r.parentMsg = t.parentMsg
		r.portIndex = t.portIndex
		return r
	case 7:
		r := new(PhyAddr)
		r.parentMsg = t.parentMsg
		r.portIndex = t.portIndex
		return r
	case 8:
		r := new(AdmStatus)
		r.parentMsg = t.parentMsg
		r.portIndex = t.portIndex
		return r
	case 9:
		r := new(OperStatus)
		r.parentMsg = t.parentMsg
		r.portIndex = t.portIndex
		return r
	case 10:
		r := new(LastChange)
		r.parentMsg = t.parentMsg
		r.portIndex = t.portIndex
		return r
	case 11:
		r := new(HighSpeed)
		r.parentMsg = t.parentMsg
		r.portIndex = t.portIndex
		return r
	case 12:
		r := new(LinkTrap)
		r.parentMsg = t.parentMsg
		r.portIndex = t.portIndex
		return r
	case 13:
		r := new(PromMode)
		r.parentMsg = t.parentMsg
		r.portIndex = t.portIndex
		return r
	case 14:
		r := new(ConPres)
		r.parentMsg = t.parentMsg
		r.portIndex = t.portIndex
		return r
	default:
		log.Printf("IfEnet TLV type: %d not supported", int(b))
		return nil
	}
}

// parseTLVs parses IfEnet TLVs.
func (t *IfEnet) parseTLVs(b []byte) ([]RCP, error) {
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

// A IPAddress is an IpAddress TLV.
type IPAddress struct {
	TLV
	// Fake port index to creat an array of interfaces
	portIndex int8
}

// Name returns the type name of an IpAddress TLV.
func (t *IPAddress) Name() string { return "IpAddress" }

// IsComplex returns whether a IpAddress TLV is Complex or not.
func (t *IPAddress) IsComplex() bool { return true }

func (t *IPAddress) newTLV(b byte) RCP {
	switch int(b) {
	// IMPORTANT
	// Port Index is not the first value sent/received in this case.
	// This makes indexing for JSON output very challenging. Need to
	// come up with a best approach here as the current one is not correct.
	// Potentially carry the entire IPAddress data struct in the TLV.
	case 1:
		r := new(AddrType)
		r.parentMsg = t.parentMsg
		r.portIndex = t.portIndex
		t.parentMsg.REX.Sequence.RpdInfo.IPAddress = append(t.parentMsg.REX.Sequence.RpdInfo.IPAddress,
			IPAdd{},
		)
		return r
	case 2:
		r := new(IPAddr)
		r.parentMsg = t.parentMsg
		r.portIndex = t.portIndex
		return r
	case 3:
		r := new(PortIdx)
		r.parentMsg = t.parentMsg
		r.portIndex = t.portIndex
		return r
	case 4:
		r := new(IntType)
		r.parentMsg = t.parentMsg
		r.portIndex = t.portIndex
		return r
	case 5:
		r := new(PrefixLen)
		r.parentMsg = t.parentMsg
		r.portIndex = t.portIndex
		return r
	case 6:
		r := new(Origin)
		r.parentMsg = t.parentMsg
		r.portIndex = t.portIndex
		return r
	case 7:
		r := new(IntStatus)
		r.parentMsg = t.parentMsg
		r.portIndex = t.portIndex
		return r
	case 8:
		r := new(Created)
		r.parentMsg = t.parentMsg
		r.portIndex = t.portIndex
		return r
	case 9:
		r := new(LastChanged)
		r.parentMsg = t.parentMsg
		r.portIndex = t.portIndex
		return r
	default:
		log.Printf("IPAddress TLV type: %d not supported", int(b))
		return nil
	}
}

// parseTLVs parses IPAddress TLVs.
func (t *IPAddress) parseTLVs(b []byte) ([]RCP, error) {
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

// A EnPortIdx is a EnetPortIndex TLV.
type EnPortIdx struct {
	TLV
	portIndex int8
}

// Name returns the type name of a EnetPortIndex TLV.
func (t *EnPortIdx) Name() string { return "EnetPortIndex" }

// Val returns the value a EnetPortIndex TLV carries.
func (t *EnPortIdx) Val() interface{} {
	if len(t.Value) != 1 {
		return fmt.Errorf("unexpected lenght: %v, want: 1", len(t.Value))
	}
	if t.parentMsg == nil {
		fmt.Printf("\n***\n***\nDEBUG: EMPTY EnetPortIndex!\n***\n***\n")
		t.parentMsg = new(GCP)
	}
	s := u8Val(t.Value)
	t.parentMsg.REX.Sequence.RpdInfo.IfEnet[t.portIndex].EnetPortIndex = s
	return s
}

// A IfName is a Name TLV.
type IfName struct {
	TLV
	portIndex int8
}

// Name returns the type name of a Name TLV.
func (t *IfName) Name() string { return "Name" }

// Val returns the value a Name TLV carries.
func (t *IfName) Val() interface{} {
	if t.parentMsg == nil {
		fmt.Printf("\n***\n***\nDEBUG: EMPTY IfName!\n***\n***\n")
		t.parentMsg = new(GCP)
	}
	s := stringVal(t.Value)
	t.parentMsg.REX.Sequence.RpdInfo.IfEnet[t.portIndex].Name = s
	return s
}

// A Descr is a Descr TLV.
type Descr struct {
	TLV
	portIndex int8
}

// Name returns the type name of a Descr TLV.
func (t *Descr) Name() string { return "Description" }

// Val returns the value a Descr TLV carries.
func (t *Descr) Val() interface{} {
	if t.parentMsg == nil {
		fmt.Printf("\n***\n***\nDEBUG: EMPTY Descr!\n***\n***\n")
		t.parentMsg = new(GCP)
	}
	s := stringVal(t.Value)
	t.parentMsg.REX.Sequence.RpdInfo.IfEnet[t.portIndex].Descr = s
	return s
}

// A IifType is a Type TLV (IANAifType).
type IifType struct {
	TLV
	portIndex int8
}

// Name returns the type name of a Type TLV.
func (t *IifType) Name() string { return "Type" }

// Val returns the value a Type TLV carries.
func (t *IifType) Val() interface{} {
	if t.parentMsg == nil {
		fmt.Printf("\n***\n***\nDEBUG: EMPTY Type!\n***\n***\n")
		t.parentMsg = new(GCP)
	}
	s := u16Val(t.Value)
	switch s {
	case "1":
		s = "other"
	case "6":
		s = "ethernetCsmacd"
	default:
	}
	t.parentMsg.REX.Sequence.RpdInfo.IfEnet[t.portIndex].Type = s
	return s
}

// A Alias is a Alias TLV.
type Alias struct {
	TLV
	portIndex int8
}

// Name returns the type name of a Alias TLV.
func (t *Alias) Name() string { return "Alias" }

// Val returns the value a Alias TLV carries.
func (t *Alias) Val() interface{} {
	if t.parentMsg == nil {
		fmt.Printf("\n***\n***\nDEBUG: EMPTY Alias!\n***\n***\n")
		t.parentMsg = new(GCP)
	}
	s := stringVal(t.Value)
	t.parentMsg.REX.Sequence.RpdInfo.IfEnet[t.portIndex].Alias = s
	return s
}

// A Mtu is a Mtu TLV.
type Mtu struct {
	TLV
	portIndex int8
}

// Name returns the type name of a Mtu TLV.
func (t *Mtu) Name() string { return "Mtu" }

// Val returns the value a Mtu TLV carries.
func (t *Mtu) Val() interface{} {
	if t.parentMsg == nil {
		fmt.Printf("\n***\n***\nDEBUG: EMPTY Mtu!\n***\n***\n")
		t.parentMsg = new(GCP)
	}
	s := u32Val(t.Value)
	t.parentMsg.REX.Sequence.RpdInfo.IfEnet[t.portIndex].MTU = s
	return s
}

// A PhyAddr is a PhysAddress TLV.
type PhyAddr struct {
	TLV
	portIndex int8
}

// Name returns the type name of a PhysAddress TLV.
func (t *PhyAddr) Name() string { return "PhysAddress" }

// Val returns the value a PhysAddress TLV carries.
func (t *PhyAddr) Val() interface{} {
	if t.parentMsg == nil {
		fmt.Printf("\n***\n***\nDEBUG: EMPTY PhysAddress!\n***\n***\n")
		t.parentMsg = new(GCP)
	}
	s := macVal(t.Value)
	t.parentMsg.REX.Sequence.RpdInfo.IfEnet[t.portIndex].PhysAddress = s
	return s
}

// A AdmStatus is a AdminStatus TLV.
type AdmStatus struct {
	TLV
	portIndex int8
}

// Name returns the type name of a AdminStatus TLV.
func (t *AdmStatus) Name() string { return "AdminStatus" }

// Val returns the value a AdminStatus TLV carries.
func (t *AdmStatus) Val() interface{} {
	if len(t.Value) != 1 {
		return fmt.Sprintf("unexpected lenght: %v, want: 1", len(t.Value))
	}
	var s string
	switch int(t.Value[0]) {
	case 1:
		s = "up"
	case 2:
		s = "down"
	case 3:
		s = "testing"
	default:
		s = "Unknown AdminStatus"
	}
	if t.parentMsg == nil {
		fmt.Printf("\n***\n***\nDEBUG: EMPTY AdminStatus!\n***\n***\n")
		t.parentMsg = new(GCP)
	}
	t.parentMsg.REX.Sequence.RpdInfo.IfEnet[t.portIndex].AdminStatus = s
	return s
}

// A OperStatus is a OperStatus TLV.
type OperStatus struct {
	TLV
	portIndex int8
}

// Name returns the type name of a OperStatus TLV.
func (t *OperStatus) Name() string { return "OperStatus" }

// Val returns the value a OperStatus TLV carries.
func (t *OperStatus) Val() interface{} {
	if len(t.Value) != 1 {
		return fmt.Errorf("unexpected lenght: %v, want: 1", len(t.Value))
	}
	var s string
	switch int(t.Value[0]) {
	case 1:
		s = "up"
	case 2:
		s = "down"
	case 3:
		s = "testing"
	case 4:
		s = "unknown"
	case 5:
		s = "dormant"
	case 6:
		s = "notPresent"
	case 7:
		s = "lowerLayerDown"
	default:
		s = "Unknown OperStatus"
	}
	if t.parentMsg == nil {
		fmt.Printf("\n***\n***\nDEBUG: EMPTY OperStatus!\n***\n***\n")
		t.parentMsg = new(GCP)
	}
	t.parentMsg.REX.Sequence.RpdInfo.IfEnet[t.portIndex].OperStatus = s
	return s
}

// A LastChange is a LastChange TLV.
type LastChange struct {
	TLV
	portIndex int8
}

// Name returns the type name of a LastChange TLV.
func (t *LastChange) Name() string { return "LastChange" }

// Val returns the value a LastChange TLV carries.
func (t *LastChange) Val() interface{} {
	if t.parentMsg == nil {
		fmt.Printf("\n***\n***\nDEBUG: EMPTY LastChange!\n***\n***\n")
		t.parentMsg = new(GCP)
	}
	s := timeVal(t.Value)
	t.parentMsg.REX.Sequence.RpdInfo.IfEnet[t.portIndex].LastChange = s
	return s
}

// A HighSpeed is a HighSpeed TLV.
type HighSpeed struct {
	TLV
	portIndex int8
}

// Name returns the type name of a HighSpeed TLV.
func (t *HighSpeed) Name() string { return "HighSpeed" }

// Val returns the value a HighSpeed TLV carries.
func (t *HighSpeed) Val() interface{} {
	if t.parentMsg == nil {
		fmt.Printf("\n***\n***\nDEBUG: EMPTY HighSpeed!\n***\n***\n")
		t.parentMsg = new(GCP)
	}
	// Speed in units of 1,000,000 bits per second.
	s := u32Val(t.Value) + " Mbps"
	t.parentMsg.REX.Sequence.RpdInfo.IfEnet[t.portIndex].HighSpeed = s
	return s
}

// A LinkTrap is a Type LinkUpDownTrapEnable.
type LinkTrap struct {
	TLV
	portIndex int8
}

// Name returns the type name of a LinkUpDownTrapEnable TLV.
func (t *LinkTrap) Name() string { return "LinkUpDownTrapEnable" }

// Val returns the value a LinkUpDownTrapEnable TLV carries.
func (t *LinkTrap) Val() interface{} {
	if t.parentMsg == nil {
		fmt.Printf("\n***\n***\nDEBUG: EMPTY LinkUpDownTrapEnable!\n***\n***\n")
		t.parentMsg = new(GCP)
	}
	s := u8Val(t.Value)
	switch s {
	case "1":
		s = "true"
	case "2":
		s = "false"
	default:
	}
	t.parentMsg.REX.Sequence.RpdInfo.IfEnet[t.portIndex].LinkUpDownTrapEnable = s
	return s
}

// A PromMode is a Type PromiscuousMode.
type PromMode struct {
	TLV
	portIndex int8
}

// Name returns the type name of a PromiscuousMode TLV.
func (t *PromMode) Name() string { return "PromiscuousMode" }

// Val returns the value a PromiscuousMode TLV carries.
func (t *PromMode) Val() interface{} {
	if t.parentMsg == nil {
		fmt.Printf("\n***\n***\nDEBUG: EMPTY PromiscuousMode!\n***\n***\n")
		t.parentMsg = new(GCP)
	}
	s := u8Val(t.Value)
	switch s {
	case "1":
		s = "true"
	case "2":
		s = "false"
	default:
	}
	t.parentMsg.REX.Sequence.RpdInfo.IfEnet[t.portIndex].PromiscuousMode = s
	return s
}

// A ConPres is a Type ConnectorPresent .
type ConPres struct {
	TLV
	portIndex int8
}

// Name returns the type name of a ConnectorPresent  TLV.
func (t *ConPres) Name() string { return "PromiscuousMode" }

// Val returns the value a ConnectorPresent  TLV carries.
func (t *ConPres) Val() interface{} {
	if t.parentMsg == nil {
		fmt.Printf("\n***\n***\nDEBUG: EMPTY ConnectorPresent !\n***\n***\n")
		t.parentMsg = new(GCP)
	}
	s := u8Val(t.Value)
	var b bool
	switch s {
	case "1":
		b = true
	case "2":
		b = false
	default:
	}
	t.parentMsg.REX.Sequence.RpdInfo.IfEnet[t.portIndex].ConnectorPresent = b
	return b
}

// A AddrType is an AddrType TLV.
type AddrType struct {
	TLV
	portIndex int8
}

// Name returns the type name of an AddrType TLV.
func (t *AddrType) Name() string { return "AddrType" }

// Val returns the value an AddrType carries.
func (t *AddrType) Val() interface{} {
	if len(t.Value) != 4 {
		return fmt.Errorf("unexpected lenght: %v, want: 4", len(t.Value))
	}
	var s string
	switch int(t.Value[3]) {
	case 1:
		s = "ipv4"
	case 2:
		s = "ipv6"
	default:
		s = "Unknown InetAddressType"
	}
	t.parentMsg.REX.Sequence.RpdInfo.IPAddress[t.portIndex].AddrType = s
	return s
}

// A IPAddr is an IpAddress TLV.
type IPAddr struct {
	TLV
	portIndex int8
}

// Name returns the type name of an IpAddress TLV.
func (t *IPAddr) Name() string { return "IpAddress" }

// Val returns the value an IpAddress TLV carries.
func (t *IPAddr) Val() interface{} {
	s := ipVal(t.Value)
	t.parentMsg.REX.Sequence.RpdInfo.IPAddress[t.portIndex].IPAddress = s
	return s
}

// A PortIdx is an EnetPortIndex TLV.
type PortIdx struct {
	TLV
	portIndex int8
}

// Name returns the type name of an EnetPortIndex TLV.
func (t *PortIdx) Name() string { return "EnetPortIndex" }

// Val returns the value an EnetPortIndex TLV carries.
func (t *PortIdx) Val() interface{} {
	s := u8Val(t.Value)
	t.parentMsg.REX.Sequence.RpdInfo.IPAddress[t.portIndex].EnetPortIndex = s
	return s
}

// A IntType is an Type TLV.
type IntType struct {
	TLV
	portIndex int8
}

// Name returns the type name of a Type TLV.
func (t *IntType) Name() string { return "Type" }

// Val returns the value a Type carries.
func (t *IntType) Val() interface{} {
	if len(t.Value) != 1 {
		return fmt.Errorf("unexpected lenght: %v, want: 1", len(t.Value))
	}
	var s string
	switch int(t.Value[0]) {
	case 1:
		s = "unicast"
	case 2:
		s = "anycast"
	case 3:
		s = "broadcast"
	default:
		s = "Unknown Type"
	}
	t.parentMsg.REX.Sequence.RpdInfo.IPAddress[t.portIndex].Type = s
	return s
}

// A PrefixLen is a PrefixLen TLV.
type PrefixLen struct {
	TLV
	portIndex int8
}

// Name returns the type name of a PrefixLen TLV.
func (t *PrefixLen) Name() string { return "PrefixLen" }

// Val returns the value a PrefixLen TLV carries.
func (t *PrefixLen) Val() interface{} {
	s := u16Val(t.Value)
	t.parentMsg.REX.Sequence.RpdInfo.IPAddress[t.portIndex].PrefixLen = s
	return s
}

// A Origin is an OriginTLV.
type Origin struct {
	TLV
	portIndex int8
}

// Name returns the type name of an Origin TLV.
func (t *Origin) Name() string { return "Type" }

// Val returns the value an Origin TLV carries.
func (t *Origin) Val() interface{} {
	if len(t.Value) != 1 {
		return fmt.Errorf("unexpected lenght: %v, want: 1", len(t.Value))
	}
	var s string
	switch int(t.Value[0]) {
	case 1:
		s = "other"
	case 2:
		s = "manual"
	case 3:
		s = "wellKnown"
	case 4:
		s = "dhcp"
	case 5:
		s = "routerAdv"
	default:
		s = "Unknown Origin"
	}
	t.parentMsg.REX.Sequence.RpdInfo.IPAddress[t.portIndex].Origin = s
	return s
}

// An IntStatus is a Status TLV.
type IntStatus struct {
	TLV
	portIndex int8
}

// Name returns the type name of a Status TLV.
func (t *IntStatus) Name() string { return "Status" }

// Val returns the value a Status carries.
func (t *IntStatus) Val() interface{} {
	if len(t.Value) != 1 {
		return fmt.Errorf("unexpected lenght: %v, want: 1", len(t.Value))
	}
	var s string
	switch int(t.Value[0]) {
	case 1:
		s = "preferred"
	case 2:
		s = "deprecated"
	case 3:
		s = "invalid"
	case 4:
		s = "inaccessible"
	case 5:
		s = "unknown"
	case 6:
		s = "tentative"
	case 7:
		s = "duplicate"
	case 8:
		s = "optimistic"
	default:
		s = "Unknown Status"
	}
	t.parentMsg.REX.Sequence.RpdInfo.IPAddress[t.portIndex].Status = s
	return s
}

// A Created is a Created TLV.
type Created struct {
	TLV
	portIndex int8
}

// Name returns the type name of a Created TLV.
func (t *Created) Name() string { return "Created" }

// Val returns the value a Created TLV carries.
func (t *Created) Val() interface{} {
	s := timeVal(t.Value)
	t.parentMsg.REX.Sequence.RpdInfo.IPAddress[t.portIndex].Created = s
	return s
}

// A LastChanged is a LastChanged TLV.
type LastChanged struct {
	TLV
	portIndex int8
}

// Name returns the type name of a LastChanged TLV.
func (t *LastChanged) Name() string { return "LastChanged" }

// Val returns the value a LastChanged TLV carries.
func (t *LastChanged) Val() interface{} {
	s := timeVal(t.Value)
	t.parentMsg.REX.Sequence.RpdInfo.IPAddress[t.portIndex].LastChanged = s
	return s
}
