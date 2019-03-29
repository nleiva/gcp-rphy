// Package transport describes the transport protocol to be used to transmit GCP messages.
// GCP is a messaging protocol and should be used with a reliable transport protocol.
// The primary choice of a transport protocol for GCP is TCP. TCP uses the IP protocol
// as a network protocol.
//
// Under special circumstances, where reliability is not required, it is possible to use
// UDP. GCP messaging can also be used with other protocols such as L2TPv3 to offload
// the L2TPv3 signaling functionality.
package transport

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"

	gcp "github.com/nleiva/gcp-rphy"
)

// Transport represent a transport protocol for GCP messages.
type Transport interface {
	Receive() error
	Send(b []byte) error
}

// TCPmessage represents the TCP Payload Encapsulation.
type TCPmessage struct {
	tid uint16 // Transaction Identifier 2 bytes
	pid uint16 // Protocol Identifier 2 bytes
	len uint16 // Length 2 bytes
	uid int    // Unit Identifier 1 byte
	msg []byte // Message Field N bytes
}

// Marshal encapsulates a GCP TCP message.
func (p TCPmessage) Marshal() ([]byte, error) {
	b := make([]byte, 7+len(p.msg))
	binary.BigEndian.PutUint16(b[:2], uint16(p.tid))
	binary.BigEndian.PutUint16(b[2:4], uint16(p.pid))
	binary.BigEndian.PutUint16(b[4:6], uint16(p.len))
	b[6] = byte(p.uid)
	copy(b[7:], p.msg)
	return b, nil
}

// UnMarshal de-encapsulates a GCP TCP message.
func UnMarshal(b []byte) (TCPmessage, error) {
	bodyLen := len(b)
	if bodyLen < 7 {
		return TCPmessage{}, gcp.ErrMessageTooShort
	}
	p := TCPmessage{
		tid: binary.BigEndian.Uint16(b[:2]),
		pid: binary.BigEndian.Uint16(b[2:4]),
		len: binary.BigEndian.Uint16(b[4:6]),
		uid: int(b[6]),
	}
	if bodyLen > 7 {
		p.msg = make([]byte, bodyLen-7)
		copy(p.msg, b[7:])
	}
	return p, nil
}

// TCPEnd is a TCP endpoint that satisfies the Transport interface.
type TCPEnd struct {
	Host string
	Port string
}

func (e TCPEnd) listen() (net.Listener, error) {
	// TODO: Validate port
	l, err := net.Listen("tcp", ":"+e.Port)
	if err != nil {
		return nil, fmt.Errorf("failed to start the server: %v", err)
	}
	return l, nil
}

// Receive listens for new GCP/TCP messages.
func (e TCPEnd) Receive() error {
	listener, err := e.listen()
	if err != nil {
		return fmt.Errorf("failed to start the server: %v", err)
	}
	defer listener.Close()
	for {
		c, err := listener.Accept()
		if err != nil {
			log.Printf("failed to accept the connection: %v", err)
			continue
		}
		go handleMessage(c)
	}
}

func (e TCPEnd) dial() (net.Conn, error) {
	// TODO: Validate address
	addr := net.JoinHostPort(e.Host, e.Port)
	c, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("failed to setup a connection: %v", err)
	}
	return c, nil
}

// Send transmits a GCP/TCP message.
func (e TCPEnd) Send(b []byte) error {
	m, err := e.CraftPkt(b)
	if err != nil {
		return fmt.Errorf("could not craft a TCP packet: %v", err)
	}
	conn, err := e.dial()
	if err != nil {
		return fmt.Errorf("could not establish a connection: %v", err)
	}
	defer conn.Close()
	t, err := m.Marshal()
	if err != nil {
		return fmt.Errorf("failed to marshall a message: %v", err)
	}
	_, err = conn.Write(t)
	if err != nil {
		return fmt.Errorf("failed to send a message: %v", err)
	}
	return nil
}

// CraftPkt encapsulates a GCP/TCP message.
func (e TCPEnd) CraftPkt(b []byte) (TCPmessage, error) {
	m := TCPmessage{
		tid: 0,                  // Unique transaction ID. A value of 0 means to ignore this field.
		pid: 1,                  // 1 = GCP Protocol Version 1.
		len: uint16(1 + len(b)), // Length of Unit Identifier Field plus Message Field
		uid: 0,                  // Unit addressing with a device. Default is 0.
		msg: b,                  // One or more GCP messages
	}
	return m, nil
}

func handleMessage(c net.Conn) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	// TODO, do NOT hardcode the MTU.
	MTU := 1500
	defer c.Close()
	for {
		buf := make([]byte, MTU)
		n, err := c.Read(buf)
		switch {
		case err == io.EOF:
			log.Printf("end of the transmition: %s\n", err.Error())
			return
		case err != nil:
			log.Printf("failed reading response: %s\n", err.Error())
			continue
		default:
			pkt, err := UnMarshal(buf[:n])
			if err != nil {
				log.Printf("failed unmarshaling TCP message: %s\n", err.Error())
				continue
			}
			m, err := gcp.ParseMessage(pkt.msg)
			if err != nil {
				log.Printf("could not parse GCP message: %s\n", err.Error())
				continue
			}
			fmt.Printf("Incoming Message (Lenght: %d) ->\n  Message Identifier: %v\n  Length: %v\n  Body: %s\n",
				n, m.MessageID, m.Lenght, m.Body.Print())
		}
	}
}
