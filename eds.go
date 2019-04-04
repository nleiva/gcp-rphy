package gcp

import (
	"encoding/binary"
	"fmt"
)

// An EDSReq represents a GCP Exchange Data Structures Request message body.
type EDSReq struct {
	TransactionID int    // Transaction ID: 2 bytes
	Mode          byte   // Mode: 1 byte
	Port          uint16 // Port: 2 bytes
	Channel       uint16 // Channel: 2 bytes
	VendorID      uint32 // Vendor ID: 4 bytes
	VendorIdx     uint8  // Vendor Index: 1 byte
	DataStr       []byte // Data Structures: N bytes
}

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
	return fmt.Sprintf(`
	Transaction ID: %d
	Mode: %v
	Port: %d
	Channel: %d
	Vendor ID: %v
	Vendor Index: %v
	Data Structures: %v`,
		p.TransactionID, p.Mode, p.Port, p.Channel, p.VendorID,
		p.VendorIdx, p.DataStr)
}

// Marshal implements the Marshal method of MessageBody interface.
func (p *EDSReq) Marshal() ([]byte, error) {
	b := make([]byte, 12+len(p.DataStr))
	binary.BigEndian.PutUint16(b[:2], uint16(p.TransactionID))
	b[2] = p.Mode
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
		TransactionID: int(binary.BigEndian.Uint16(b[:2])),
		Mode:          (b[2]),
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
	TransactionID int    // Transaction ID: 2 bytes
	Mode          byte   // Mode: 1 byte
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
	return fmt.Sprintf(`
	Transaction ID: %d
	Mode: %v
	Port: %d
	Channel: %d
	Vendor ID: %v
	Vendor Index: %v
	Data Structures: %v`,
		p.TransactionID, p.Mode, p.Port, p.Channel, p.VendorID,
		p.VendorIdx, p.DataStr)
}

// Marshal implements the Marshal method of MessageBody interface.
func (p *EDSRes) Marshal() ([]byte, error) {
	b := make([]byte, 12+len(p.DataStr))
	binary.BigEndian.PutUint16(b[:2], uint16(p.TransactionID))
	b[2] = p.Mode
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
		TransactionID: int(binary.BigEndian.Uint16(b[:2])),
		Mode:          (b[2]),
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
