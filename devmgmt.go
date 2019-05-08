package gcp

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"strconv"
)

// An DMReq represents a GCP Device Management (GDM) Request message body.
type DMReq struct {
	TransactionID uint16 // Transaction ID: 2 bytes
	Mode          uint8  // Mode: 1 byte
	Port          uint16 // Port: 2 bytes
	Channel       uint16 // Channel: 2 bytes
	Command       uint8  // Command: 1 byte
}

// Len implements the Len method of MessageBody interface.
func (p *DMReq) Len() int {
	if p == nil {
		return 0
	}
	return 8
}

// Process generates an output for a GCP Device Management (GDM) Request message.
func (p *DMReq) Process() (string, *GCP) {
	if p == nil {
		return "", nil
	}

	// The RCP Top Level TLV for this message.
	var g GCP
	g.DM = new(cmnd)
	g.DM.Command = strconv.Itoa(int(p.Command))

	js, _ := json.MarshalIndent(g, "", "  ")
	data := "\n" + fmt.Sprintf("%s\n", js)
	return fmt.Sprintf(`
    Transaction ID: %d
    Mode: %v
    Port: %d
    Channel: %d
    Command: %s`,
		p.TransactionID, p.Mode, p.Port, p.Channel, data), &g
}

// Marshal implements the Marshal method of MessageBody interface.
func (p *DMReq) Marshal() ([]byte, error) {
	b := make([]byte, 8)
	binary.BigEndian.PutUint16(b[:2], p.TransactionID)
	b[2] = byte(p.Mode)
	binary.BigEndian.PutUint16(b[3:5], p.Port)
	binary.BigEndian.PutUint16(b[5:7], p.Channel)
	b[7] = byte(p.Command)
	return b, nil
}

// parseDMReq parses b as a GCP Device Management (GDM) Request message body.
func parseDMReq(_ MessageID, b []byte) (MessageBody, error) {
	bodyLen := len(b)
	if bodyLen != 8 {
		return nil, ErrMessageTooShort
	}
	p := &DMReq{
		TransactionID: binary.BigEndian.Uint16(b[:2]),
		Mode:          uint8(b[2]),
		Port:          binary.BigEndian.Uint16(b[3:5]),
		Channel:       binary.BigEndian.Uint16(b[5:7]),
		Command:       uint8(b[7]),
	}
	return p, nil
}
