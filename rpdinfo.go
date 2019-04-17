package gcp

import (
	"fmt"
)

// A RpdInfo is a RpdInfo TLV (Complex TLV).
type RpdInfo struct {
	TLV
	index uint8
}

// Name returns the type name of a RpdInfo TLV.
func (t *RpdInfo) Name() string { return "RpdInfo" }

// IsComplex returns whether a RpdInfo TLV is Complex or not.
func (t *RpdInfo) IsComplex() bool { return true }

func (t *RpdInfo) newTLV(b byte) RCP {
	switch int(b) {
	case 8:
		return new(IfEnet)
	case 15:
		return new(IPAddress)
	default:
		return new(TLV)
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
}

// Name returns the type name of an IfEnet TLV.
func (t *IfEnet) Name() string { return "IfEnet" }

// IsComplex returns whether a IfEnet TLV is Complex or not.
func (t *IfEnet) IsComplex() bool { return true }

func (t *IfEnet) newTLV(b byte) RCP {
	switch int(b) {
	case 1:
		return new(EnPortIdx)
	case 2:
		return new(IfName)
	case 3:
		return new(Descr)
	case 5:
		return new(Alias)
	case 6:
		return new(Mtu)
	case 7:
		return new(PhyAddr)
	case 8:
		return new(AdmStatus)
	case 9:
		return new(OperStatus)
	case 10:
		return new(LastChange)
	case 11:
		return new(HighSpeed)
	default:
		return new(TLV)
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
}

// Name returns the type name of an IpAddress TLV.
func (t *IPAddress) Name() string { return "IpAddress" }

// IsComplex returns whether a IpAddress TLV is Complex or not.
func (t *IPAddress) IsComplex() bool { return true }

func (t *IPAddress) newTLV(b byte) RCP {
	switch int(b) {
	case 1:
		return new(AddrType)
	case 2:
		return new(IPAddr)
	case 3:
		return new(PortIdx)
	case 4:
		return new(IntType)
	case 5:
		return new(PrefixLen)
	case 6:
		return new(Origin)
	case 7:
		return new(IntStatus)
	case 8:
		return new(Created)
	case 9:
		return new(LastChanged)
	default:
		return new(TLV)
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
}

// Name returns the type name of a EnetPortIndex TLV.
func (t *EnPortIdx) Name() string { return "EnetPortIndex" }

// Val returns the value a EnetPortIndex TLV carries.
func (t *EnPortIdx) Val() interface{} {
	if len(t.Value) != 1 {
		return fmt.Errorf("unexpected lenght: %v, want: 1", len(t.Value))
	}
	return string(t.Value[0])
}

// A IfName is a Name TLV.
type IfName struct {
	TLV
}

// Name returns the type name of a Name TLV.
func (t *IfName) Name() string { return "Name" }

// Val returns the value a Name TLV carries.
func (t *IfName) Val() interface{} { return stringVal(t.Value) }

// A Descr is a Descr TLV.
type Descr struct {
	TLV
}

// Name returns the type name of a Descr TLV.
func (t *Descr) Name() string { return "Description" }

// Val returns the value a Descr TLV carries.
func (t *Descr) Val() interface{} { return stringVal(t.Value) }

// A Alias is a Alias TLV.
type Alias struct {
	TLV
}

// Name returns the type name of a Alias TLV.
func (t *Alias) Name() string { return "Alias" }

// Val returns the value a Alias TLV carries.
func (t *Alias) Val() interface{} { return stringVal(t.Value) }

// A Mtu is a Mtu TLV.
type Mtu struct {
	TLV
}

// Name returns the type name of a Mtu TLV.
func (t *Mtu) Name() string { return "Mtu" }

// Val returns the value a Mtu TLV carries.
func (t *Mtu) Val() interface{} { return u32Val(t.Value) }

// A PhyAddr is a PhysAddress TLV.
type PhyAddr struct {
	TLV
}

// Name returns the type name of a PhysAddress TLV.
func (t *PhyAddr) Name() string { return "PhysAddress" }

// Val returns the value a PhysAddress TLV carries.
func (t *PhyAddr) Val() interface{} { return macVal(t.Value) }

// A AdmStatus is a AdminStatus TLV.
type AdmStatus struct {
	TLV
}

// Name returns the type name of a AdminStatus TLV.
func (t *AdmStatus) Name() string { return "AdminStatus" }

// Val returns the value a AdminStatus TLV carries.
func (t *AdmStatus) Val() interface{} {
	if len(t.Value) != 1 {
		return fmt.Sprintf("unexpected lenght: %v, want: 1", len(t.Value))
	}
	switch int(t.Value[0]) {
	case 1:
		return "up"
	case 2:
		return "down"
	case 3:
		return "testing"
	default:
		return "Unknown AdminStatus"
	}
}

// A OperStatus is a OperStatus TLV.
type OperStatus struct {
	TLV
}

// Name returns the type name of a OperStatus TLV.
func (t *OperStatus) Name() string { return "OperStatus" }

// Val returns the value a OperStatusTLV carries.
func (t *OperStatus) Val() interface{} {
	if len(t.Value) != 1 {
		return fmt.Errorf("unexpected lenght: %v, want: 1", len(t.Value))
	}
	switch int(t.Value[0]) {
	case 1:
		return "up"
	case 2:
		return "down"
	case 3:
		return "testing"
	case 4:
		return "unknown"
	case 5:
		return "dormant"
	case 6:
		return "notPresent"
	case 7:
		return "lowerLayerDown"
	default:
		return "Unknown OperStatus"
	}
}

// A LastChange is a LastChange TLV.
type LastChange struct {
	TLV
}

// Name returns the type name of a LastChange TLV.
func (t *LastChange) Name() string { return "LastChange" }

// Val returns the value a LastChange TLV carries.
func (t *LastChange) Val() interface{} { return timeVal(t.Value) }

// A HighSpeed is a HighSpeed TLV.
type HighSpeed struct {
	TLV
}

// Name returns the type name of a LastChange TLV.
func (t *HighSpeed) Name() string { return "HighSpeed" }

// Val returns the value a HighSpeed TLV carries.
func (t *HighSpeed) Val() interface{} { return u32Val(t.Value) }

// A AddrType is an AddrType TLV.
type AddrType struct {
	TLV
}

// Name returns the type name of an AddrType TLV.
func (t *AddrType) Name() string { return "AddrType" }

// Val returns the value an AddrType carries.
func (t *AddrType) Val() interface{} {
	if len(t.Value) != 4 {
		return fmt.Errorf("unexpected lenght: %v, want: 4", len(t.Value))
	}
	switch int(t.Value[3]) {
	case 1:
		return "ipv4"
	case 2:
		return "ipv6"
	default:
		return "Unknown InetAddressType"
	}
}

// A IPAddr is an IpAddress TLV.
type IPAddr struct {
	TLV
}

// Name returns the type name of an IpAddress TLV.
func (t *IPAddr) Name() string { return "IpAddress" }

// Val returns the value an IpAddress TLV carries.
func (t *IPAddr) Val() interface{} { return ipVal(t.Value) }

// A PortIdx is an EnetPortIndex TLV.
type PortIdx struct {
	TLV
}

// Name returns the type name of an EnetPortIndex TLV.
func (t *PortIdx) Name() string { return "EnetPortIndex" }

// Val returns the value an EnetPortIndex TLV carries.
func (t *PortIdx) Val() interface{} { return u8Val(t.Value) }

// A IntType is an Type TLV.
type IntType struct {
	TLV
}

// Name returns the type name of a Type TLV.
func (t *IntType) Name() string { return "Type" }

// Val returns the value a Type carries.
func (t *IntType) Val() interface{} {
	if len(t.Value) != 1 {
		return fmt.Errorf("unexpected lenght: %v, want: 1", len(t.Value))
	}
	switch int(t.Value[0]) {
	case 1:
		return "unicast"
	case 2:
		return "anycast"
	case 3:
		return "broadcast"
	default:
		return "Unknown Type"
	}
}

// A PrefixLen is a PrefixLen TLV.
type PrefixLen struct {
	TLV
}

// Name returns the type name of a PrefixLen TLV.
func (t *PrefixLen) Name() string { return "PrefixLen" }

// Val returns the value a PrefixLen TLV carries.
func (t *PrefixLen) Val() interface{} { return u16Val(t.Value) }

// A Origin is an OriginTLV.
type Origin struct {
	TLV
}

// Name returns the type name of an Origin TLV.
func (t *Origin) Name() string { return "Type" }

// Val returns the value an Origin TLV carries.
func (t *Origin) Val() interface{} {
	if len(t.Value) != 1 {
		return fmt.Errorf("unexpected lenght: %v, want: 1", len(t.Value))
	}
	switch int(t.Value[0]) {
	case 1:
		return "other"
	case 2:
		return "manual"
	case 3:
		return "wellKnown"
	case 4:
		return "dhcp"
	case 5:
		return "routerAdv"
	default:
		return "Unknown Origin"
	}
}

// An IntStatus is a Status TLV.
type IntStatus struct {
	TLV
}

// Name returns the type name of a Status TLV.
func (t *IntStatus) Name() string { return "Status" }

// Val returns the value a Status carries.
func (t *IntStatus) Val() interface{} {
	if len(t.Value) != 1 {
		return fmt.Errorf("unexpected lenght: %v, want: 1", len(t.Value))
	}
	switch int(t.Value[0]) {
	case 1:
		return "preferred"
	case 2:
		return "deprecated"
	case 3:
		return "invalid"
	case 4:
		return "inaccessible"
	case 5:
		return "unknown"
	case 6:
		return "tentative"
	case 7:
		return "duplicate"
	case 8:
		return "optimistic"
	default:
		return "Unknown Status"
	}
}

// A Created is a Created TLV.
type Created struct {
	TLV
}

// Name returns the type name of a Created TLV.
func (t *Created) Name() string { return "Created" }

// Val returns the value a Created TLV carries.
func (t *Created) Val() interface{} { return timeVal(t.Value) }

// A LastChanged is a LastChanged TLV.
type LastChanged struct {
	TLV
}

// Name returns the type name of a LastChanged TLV.
func (t *LastChanged) Name() string { return "LastChanged" }

// Val returns the value a LastChanged TLV carries.
func (t *LastChanged) Val() interface{} { return timeVal(t.Value) }
