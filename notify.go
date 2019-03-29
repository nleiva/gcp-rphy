package gcp

import (
	"encoding/binary"
	"fmt"
)

// An NotifyReq represents a GCP Notify Request message body.
type NotifyReq struct {
	TransactionID int    // Transaction ID: 2 bytes
	Mode          byte   // Mode: 1 byte
	Status        int    // Status: 1 byte
	EvntCode      int    // Event Code: 4 bytes
	EvntData      []byte // Event Data: N bytes
}

// Len implements the Len method of MessageBody interface.
func (p *NotifyReq) Len() int {
	if p == nil {
		return 0
	}
	return 8 + len(p.EvntData)
}

// Print generates an output for a Notify Request message.
func (p *NotifyReq) Print() string {
	if p == nil {
		return ""
	}
	return fmt.Sprintf(`
	Transaction ID: %d
	Mode: %v
	Status: %d
	Event Code: %d
	Event Data: %v`,
		p.TransactionID, p.Mode, p.Status, p.EvntCode, p.EvntData)
}

// Marshal implements the Marshal method of MessageBody interface.
func (p *NotifyReq) Marshal() ([]byte, error) {
	b := make([]byte, 8+len(p.EvntData))
	binary.BigEndian.PutUint16(b[:2], uint16(p.TransactionID))
	b[2] = p.Mode
	b[3] = byte(p.Status)
	binary.BigEndian.PutUint32(b[4:8], uint32(p.EvntCode))
	copy(b[8:], p.EvntData)
	return b, nil
}

// parseNotifyReq parses b as a GCP notify request message body.
func parseNotifyReq(_ MessageID, b []byte) (MessageBody, error) {
	bodyLen := len(b)
	if bodyLen < 8 {
		return nil, ErrMessageTooShort
	}
	p := &NotifyReq{
		TransactionID: int(binary.BigEndian.Uint16(b[:2])),
		Mode:          (b[2]),
		Status:        int(b[3]),
		EvntCode:      int(binary.BigEndian.Uint32(b[4:8])),
	}
	if bodyLen > 8 {
		p.EvntData = make([]byte, bodyLen-8)
		copy(p.EvntData, b[8:])
	}
	return p, nil
}

// An NotifyRes represents a GCP Notify Response message body.
type NotifyRes struct {
	TransactionID int  // Transaction ID: 2 bytes
	Mode          byte // Mode: 1 byte
	EvntCode      int  // Event Code: 4 bytes
}

// Len implements the Len method of MessageBody interface.
func (p *NotifyRes) Len() int {
	if p == nil {
		return 0
	}
	return 7
}

// Print generates an output for a Notify Response message.
func (p *NotifyRes) Print() string {
	if p == nil {
		return ""
	}
	return fmt.Sprintf(`
	Transaction ID: %d
	Mode: %v
	Event Code: %d`,
		p.TransactionID, p.Mode, p.EvntCode)
}

// Marshal implements the Marshal method of MessageBody interface.
func (p *NotifyRes) Marshal() ([]byte, error) {
	b := make([]byte, 7)
	binary.BigEndian.PutUint16(b[:2], uint16(p.TransactionID))
	b[2] = p.Mode
	binary.BigEndian.PutUint32(b[3:7], uint32(p.EvntCode))
	return b, nil
}

// parseNotifyRes parses b as a GCP notify response message body.
func parseNotifyRes(_ MessageID, b []byte) (MessageBody, error) {
	bodyLen := len(b)
	if bodyLen < 7 {
		return nil, ErrMessageTooShort
	}
	p := &NotifyRes{
		TransactionID: int(binary.BigEndian.Uint16(b[:2])),
		Mode:          (b[2]),
		EvntCode:      int(binary.BigEndian.Uint32(b[3:7])),
	}
	return p, nil
}

// An NotifyErr represents a GCP Notify Response message body.
type NotifyErr struct {
	TransactionID int // Transaction ID: 2 bytes
	RtrnCode      int // Return Code: 1 byte
}

// Len implements the Len method of MessageBody interface.
func (p *NotifyErr) Len() int {
	if p == nil {
		return 0
	}
	return 3
}

// Print generates an output for a Notify Error message.
func (p *NotifyErr) Print() string {
	if p == nil {
		return ""
	}
	return fmt.Sprintf(`
	Transaction ID: %d
	Return Code: %d`,
		p.TransactionID, p.RtrnCode)
}

// Marshal implements the Marshal method of MessageBody interface.
func (p *NotifyErr) Marshal() ([]byte, error) {
	b := make([]byte, 3)
	binary.BigEndian.PutUint16(b[:2], uint16(p.TransactionID))
	b[2] = byte(p.RtrnCode)
	return b, nil
}

// parseNotifyErr parses b as a GCP notify error message body.
func parseNotifyErr(_ MessageID, b []byte) (MessageBody, error) {
	bodyLen := len(b)
	if bodyLen < 3 {
		return nil, ErrMessageTooShort
	}
	p := &NotifyErr{
		TransactionID: int(binary.BigEndian.Uint16(b[:2])),
		RtrnCode:      int(b[2]),
	}
	return p, nil
}
