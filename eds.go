package gcp

import (
	"encoding/binary"
	"fmt"
)

// An EDSReq represents a GCP Exchange Data Structures Request message body.
type EDSReq struct {
	TransactionID uint16 // Transaction ID: 2 bytes
	Mode          uint8  // Mode: 1 byte
	Port          uint16 // Port: 2 bytes
	Channel       uint16 // Channel: 2 bytes
	VendorID      uint32 // Vendor ID: 4 bytes
	VendorIdx     uint8  // Vendor Index: 1 byte
	DataStr       []byte // Data Structures: N bytes
}

// Vendor ID Refers to IETF RFC 3232 "Assigned Number" by the IETF, Jan 2002.
// This spec refers to the IANA web page which is
// http://www.iana.org/assignments/enterprise-numbers.
const (
	Cisco     uint32 = 9    // Cisco Private Enterprise Code
	CableLabs uint32 = 4491 // Cable Labs Private Enterprise Code
)

// Len implements the Len method of MessageBody interface for a GCP Exchange Data Structures
// Request message.
func (p *EDSReq) Len() int {
	if p == nil {
		return 0
	}
	return 12 + len(p.DataStr)
}

// Print generates an output for a Notify Request message.
func (p *EDSReq) Print() string {
	if p == nil {
		return ""
	}
	data := "\n"
	tlvs, err := parseTLVs(p.DataStr)
	if err != nil {
		return err.Error()
	}
	for _, t := range tlvs {
		switch t.IsComplex() {
		case true:
			data = data + fmt.Sprintf("        Type: %s, \tLength: %v ->\n", t.Name(), t.Len())
		default:
			data = data + fmt.Sprintf("        Type: %s, \tLength: %v, \tValue: %v\n", t.Name(), t.Len(), t.Val())
		}
	}
	return fmt.Sprintf(`
    Transaction ID: %d
    Mode: %v
    Port: %d
    Channel: %d
    Vendor ID: %v
    Vendor Index: %v
    Data Structures: %s`,
		p.TransactionID, p.Mode, p.Port, p.Channel, p.VendorID,
		p.VendorIdx, data)
}

// Marshal implements the Marshal method of MessageBody interface.
func (p *EDSReq) Marshal() ([]byte, error) {
	b := make([]byte, 12+len(p.DataStr))
	binary.BigEndian.PutUint16(b[:2], p.TransactionID)
	b[2] = byte(p.Mode)
	binary.BigEndian.PutUint16(b[3:5], uint16(p.Port))
	binary.BigEndian.PutUint16(b[5:7], uint16(p.Channel))
	binary.BigEndian.PutUint32(b[7:11], uint32(p.VendorID))
	b[11] = byte(p.VendorIdx)
	copy(b[12:], p.DataStr)
	return b, nil
}

// EDSReqReq parses b as a GCP Exchange Data Structures Request message body.
func parseEDSReq(_ MessageID, b []byte) (MessageBody, error) {
	bodyLen := len(b)
	if bodyLen < 12 {
		return nil, ErrMessageTooShort
	}
	p := &EDSReq{
		TransactionID: binary.BigEndian.Uint16(b[:2]),
		Mode:          uint8(b[2]),
		Port:          binary.BigEndian.Uint16(b[3:5]),
		Channel:       binary.BigEndian.Uint16(b[5:7]),
		VendorID:      binary.BigEndian.Uint32(b[7:11]),
		VendorIdx:     uint8(b[11]),
	}
	if bodyLen > 8 {
		p.DataStr = make([]byte, bodyLen-12)
		copy(p.DataStr, b[12:])
	}
	return p, nil
}

// An EDSRes represents a GCP Exchange Data Structures Normal Response message body.
type EDSRes struct {
	TransactionID uint16 // Transaction ID: 2 bytes
	Mode          uint8  // Mode: 1 byte
	Port          uint16 // Port: 2 bytes
	Channel       uint16 // Channel: 2 bytes
	VendorID      uint32 // Vendor ID: 4 bytes
	VendorIdx     uint8  // Vendor Index: 1 byte
	DataStr       []byte // Data Structures: N bytes
}

// Len implements the Len method of MessageBody interface for a GCP Exchange Data Structures
// Normal Response message.
func (p *EDSRes) Len() int {
	if p == nil {
		return 0
	}
	return 12 + len(p.DataStr)
}

// Print generates an output for a GCP Exchange Data Structures Normal Response message.
func (p *EDSRes) Print() string {
	if p == nil {
		return ""
	}
	data := "\n"
	tlvs, err := parseTLVs(p.DataStr)
	if err != nil {
		return err.Error()
	}
	for _, t := range tlvs {
		switch t.IsComplex() {
		case true:
			data = data + fmt.Sprintf("        Type: %s, \tLength: %v ->\n", t.Name(), t.Len())
		default:
			data = data + fmt.Sprintf("        Type: %s, \tLength: %v, \tValue: %v\n", t.Name(), t.Len(), t.Val())
		}
	}
	return fmt.Sprintf(`
    Transaction ID: %d
    Mode: %v
    Port: %d
    Channel: %d
    Vendor ID: %v
    Vendor Index: %v
    Data Structures: %s`,
		p.TransactionID, p.Mode, p.Port, p.Channel, p.VendorID,
		p.VendorIdx, data)
}

// Marshal implements the Marshal method of MessageBody interface.
func (p *EDSRes) Marshal() ([]byte, error) {
	b := make([]byte, 12+len(p.DataStr))
	binary.BigEndian.PutUint16(b[:2], p.TransactionID)
	b[2] = byte(p.Mode)
	binary.BigEndian.PutUint16(b[3:5], uint16(p.Port))
	binary.BigEndian.PutUint16(b[5:7], uint16(p.Channel))
	binary.BigEndian.PutUint32(b[7:11], uint32(p.VendorID))
	b[11] = byte(p.VendorIdx)
	copy(b[12:], p.DataStr)
	return b, nil
}

// EDSReqReq parses b as a GCP Exchange Data Structures Request message body.
func parseEDSRes(_ MessageID, b []byte) (MessageBody, error) {
	bodyLen := len(b)
	if bodyLen < 12 {
		return nil, ErrMessageTooShort
	}
	p := &EDSRes{
		TransactionID: binary.BigEndian.Uint16(b[:2]),
		Mode:          uint8(b[2]),
		Port:          binary.BigEndian.Uint16(b[3:5]),
		Channel:       binary.BigEndian.Uint16(b[5:7]),
		VendorID:      binary.BigEndian.Uint32(b[7:11]),
		VendorIdx:     uint8(b[11]),
	}
	if bodyLen > 8 {
		p.DataStr = make([]byte, bodyLen-12)
		copy(p.DataStr, b[12:])
	}
	return p, nil
}
