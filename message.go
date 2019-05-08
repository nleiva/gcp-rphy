package gcp

import (
	"encoding/binary"
	"errors"
	"fmt"
)

// A Message represents a GCP message.
type Message struct {
	MessageID uint8  // Message Identifier: 1 byte
	Lenght    uint16 // Message Length: 2 bytes
	Body      MessageBody
}

// A MessageBody represents a GCP message body.
type MessageBody interface {
	Len() int
	Marshal() ([]byte, error)
	Process() (string, *GCP)
}

// A MessageID represents a Message ID of a GCP message.
type MessageID uint8

// Error messages
var (
	ErrMessageTooShort = errors.New("message too short")
	errMessageID       = errors.New("invalid message id")
)

// GCP Message ID Parameters, Updated: 2018-05-09
const (
	MessageIDNotifyReq MessageID = 2   // GCP Notify Request
	MessageIDNotifyRes MessageID = 3   // GCP Notify Normal Response
	MessageIDGDMReq    MessageID = 4   // GCP Device Management (GDM) Request
	MessageIDGDMRes    MessageID = 5   // GCP Device Management (GDM) Normal Response
	MessageIDEDSReq    MessageID = 6   // Exchange Data Structures (EDS) Request
	MessageIDEDSRes    MessageID = 7   // Exchange Data Structures (EDS) Normal Response
	MessageIDEDRReq    MessageID = 16  // Exchange Data Registers (EDR) Request
	MessageIDEDRRes    MessageID = 17  // Exchange Data Registers (EDR) Normal Response
	MessageIDMWRReq    MessageID = 18  // Mask Write Register (MWR) Request
	MessageIDMWRRes    MessageID = 19  // Mask Write Register (MWR) Normal Response
	MessageIDNotifyErr MessageID = 131 // GCP Notify Error Response
	MessageIDGDMErr    MessageID = 133 // GCP Device Management Error Response
	MessageIDEDSErr    MessageID = 135 // Exchange Data Structures Error Response
	MessageIDEDRErr    MessageID = 145 // Exchange Data Registers (EDR) Error Response
	MessageIDMWRErr    MessageID = 147 // Mask Write Register (MWR) Error Response
)

// A RtrnCode represents a Return Code of a GCP message.
type RtrnCode int

// GCP Return Codes, Updated: 2018-05-09
const (
	MsgSuccess         RtrnCode = 0   // MESSAGE SUCCESSFUL
	UnsupportedMsg     RtrnCode = 1   // UNSUPPORTED MESSAGE
	IllegalMsgLen      RtrnCode = 2   // ILLEGAL MESSAGE LENGTH
	IllegalTransID     RtrnCode = 3   // ILLEGAL TRANSACTION ID
	IllegalMode        RtrnCode = 4   // ILLEGAL MODE
	IllegalPort        RtrnCode = 5   // ILLEGAL PORT
	IllegalChannel     RtrnCode = 6   // ILLEGAL CHANNEL
	IllegalCmd         RtrnCode = 7   // ILLEGAL COMMAND
	IllegalVendorID    RtrnCode = 8   // ILLEGAL VENDOR ID
	IllegalVendorIndex RtrnCode = 9   // ILLEGAL VENDOR INDEX
	IllegalAddr        RtrnCode = 10  // ILLEGAL ADDRESS
	IllegalDataValue   RtrnCode = 11  // ILLEGAL DATA VALUE
	MsgFail            RtrnCode = 12  // MESSAGE FAILURE
	SlaveDevFail       RtrnCode = 255 // SLAVE DEVICE FAILURE
)

var parseFns = map[MessageID]func(MessageID, []byte) (MessageBody, error){
	MessageIDNotifyReq: parseNotifyReq,
	MessageIDGDMReq:    parseDMReq,
	MessageIDEDSReq:    parseEDSReq,
	MessageIDEDSRes:    parseEDSRes,
}

// A RawBody represents a raw message body.
//
// A raw message body is excluded from message processing and can be
// used to construct applications such as protocol conformance
// testing.
type RawBody struct {
	Data []byte // data
}

// Len implements the Len method of MessageBody interface.
func (p *RawBody) Len() int {
	if p == nil {
		return 0
	}
	return len(p.Data)
}

// Process generates an output for a Raw message.
func (p *RawBody) Process() (string, *GCP) {
	if p == nil {
		return "", nil
	}
	return fmt.Sprintf(`
	Data: %v`, p.Data), nil
}

// Marshal implements the Marshal method of MessageBody interface.
func (p *RawBody) Marshal() ([]byte, error) {
	return p.Data, nil
}

// parseRawBody parses b as a GCP message body.
func parseRawBody(b []byte) (MessageBody, error) {
	p := &RawBody{Data: make([]byte, len(b))}
	copy(p.Data, b)
	return p, nil
}

// ParseMessage parses b as a GCP message.
func ParseMessage(b []byte) (*Message, error) {
	if len(b) < 3 {
		return nil, ErrMessageTooShort
	}
	id := uint8(b[0])
	mlen := binary.BigEndian.Uint16(b[1:3])
	var err error
	m := &Message{MessageID: id, Lenght: mlen}
	if fn, ok := parseFns[MessageID(id)]; !ok {
		m.Body, err = parseRawBody(b[3:])
	} else {
		m.Body, err = fn(MessageID(id), b[3:])
	}
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Marshal converts a GCP message into a byte array.
func (p *Message) Marshal() ([]byte, error) {
	bd, err := p.Body.Marshal()
	if err != nil {
		return nil, err
	}
	b := make([]byte, 3+len(bd))
	b[0] = byte(p.MessageID)
	binary.BigEndian.PutUint16(b[1:3], p.Lenght)
	copy(b[3:], bd)
	return b, nil
}
