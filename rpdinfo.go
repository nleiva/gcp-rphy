package gcp

import (
	"fmt"
	"net"
)

// A RpdInfo is a RpdInfo TLV (Complex TLV).
type RpdInfo struct {
	TLV
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
			// DEBUG
			return nil, err
			// fmt.Printf("ERROR IfEnet -> i: %d, type: %v", i, b[i])
			// return nil, fmt.Errorf("ERROR -> i: %d, type: %v", i, b[i])
		}
		// fmt.Printf("DEBUG IfEnet -> i: %d, type: %v, length: %d\n", i, b[i], l)
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
			// DEBUG
			return nil, err
			// fmt.Printf("ERROR IPAddress -> i: %d, type: %v", i, b[i])
			// return nil, fmt.Errorf("ERROR -> i: %d, type: %v", i, b[i])
		}
		// fmt.Printf("DEBUG IPAddress -> i: %d, type: %v, length: %d\n", i, b[i], l)
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

// A EnPortIdx is a EnetPortIndex TLV.
type EnPortIdx struct {
	TLV
}

// Name returns the type name of a EnetPortIndex TLV.
func (t *EnPortIdx) Name() string { return "EnetPortIndex" }

// IsComplex returns whether a EnetPortIndex TLV is Complex or not.
func (t *EnPortIdx) IsComplex() bool { return false }

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

// IsComplex returns whether a Name TLV is Complex or not.
func (t *IfName) IsComplex() bool { return false }

// A Descr is a Descr TLV.
type Descr struct {
	TLV
}

// Name returns the type name of a Descr TLV.
func (t *Descr) Name() string { return "Description" }

// Val returns the value a Descr TLV carries.
func (t *Descr) Val() interface{} { return stringVal(t.Value) }

// IsComplex returns whether a Descr TLV is Complex or not.
func (t *Descr) IsComplex() bool { return false }

// A Alias is a Alias TLV.
type Alias struct {
	TLV
}

// Name returns the type name of a Alias TLV.
func (t *Alias) Name() string { return "Alias" }

// Val returns the value a Alias TLV carries.
func (t *Alias) Val() interface{} { return stringVal(t.Value) }

// IsComplex returns whether a Alias TLV is Complex or not.
func (t *Alias) IsComplex() bool { return false }

// A Mtu is a Mtu TLV.
type Mtu struct {
	TLV
}

// Name returns the type name of a Mtu TLV.
func (t *Mtu) Name() string { return "Mtu" }

// Val returns the value a Mtu TLV carries.
func (t *Mtu) Val() interface{} { return u32Val(t.Value) }

// IsComplex returns whether a Mtu TLV is Complex or not.
func (t *Mtu) IsComplex() bool { return false }

// A PhyAddr is a PhysAddress TLV.
type PhyAddr struct {
	TLV
}

// Name returns the type name of a PhysAddress TLV.
func (t *PhyAddr) Name() string { return "PhysAddress" }

// Val returns the value a PhysAddress TLV carries.
func (t *PhyAddr) Val() interface{} {
	if len(t.Value) != 6 {
		return fmt.Sprintf("unexpected lenght: %v, want: 6", len(t.Value))
	}
	return net.HardwareAddr(t.Value).String()
}

// IsComplex returns whether a PhysAddress TLV is Complex or not.
func (t *PhyAddr) IsComplex() bool { return false }

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

// IsComplex returns whether a AdminStatus TLV is Complex or not.
func (t *AdmStatus) IsComplex() bool { return false }

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

// IsComplex returns whether a OperStatus TLV is Complex or not.
func (t *OperStatus) IsComplex() bool { return false }

// A LastChange is a LastChange TLV.
type LastChange struct {
	TLV
}

// Name returns the type name of a LastChange TLV.
func (t *LastChange) Name() string { return "LastChange" }

// Val returns the value a LastChange TLV carries.
func (t *LastChange) Val() interface{} { return u32Val(t.Value) }

// IsComplex returns whether a LastChange TLV is Complex or not.
func (t *LastChange) IsComplex() bool { return false }

// A HighSpeed is a HighSpeed TLV.
type HighSpeed struct {
	TLV
}

// Name returns the type name of a LastChange TLV.
func (t *HighSpeed) Name() string { return "HighSpeed" }

// Val returns the value a HighSpeed TLV carries.
func (t *HighSpeed) Val() interface{} { return u32Val(t.Value) }

// IsComplex returns whether a HighSpeed TLV is Complex or not.
func (t *HighSpeed) IsComplex() bool { return false }
