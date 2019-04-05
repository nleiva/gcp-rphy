// Command gcp is a utility for working with the Generic Control Plane Protocol.
package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"time"

	gcp "github.com/nleiva/gcp-rphy"
	"github.com/nleiva/gcp-rphy/transport"
)

func main() {
	var (
		modeFlag   = flag.String("m", "", "connection mode: server or client")
		targetFlag = flag.String("t", "", "target address for GCP connection")
		wordFlag   = flag.String("w", "", "word to send over the GCP connection")
		port       = "8190"
	)

	flag.Usage = func() {
		fmt.Println(usage)
		fmt.Println("Flags:")
		flag.PrintDefaults()
	}

	flag.Parse()
	if flag.NArg() > 1 {
		log.Fatalf("too many args on command line: %v", flag.Args()[1:])
	}

	if *modeFlag == "server" {
		e := transport.TCPEnd{
			Port: port,
		}
		err := e.Receive()
		if err != nil {
			log.Fatalf("Server failed: %v", err)
		}

	} else if *modeFlag == "client" {
		e := transport.TCPEnd{
			Host: *targetFlag,
			Port: port,
		}
		// Create a Notify Message, Message ID: 2.
		m, err := Encapsulate(2, 0, []byte(*wordFlag))
		if err != nil {
			log.Fatalf("couldn't create a Notify message: %v", err)
		}
		// Create the GCP packet.
		b, err := m.Marshal()
		if err != nil {
			log.Fatalf("couldn't marshall a Notify message: %v", err)
		}
		err = e.Send(b)
		if err != nil {
			log.Fatalf("Client failed: %v", err)
		}
	} else {
		log.Fatalf("unknown mode")
	}
}

// Encapsulate is a temp function to encapsulate a GCP message
func Encapsulate(id uint8, tid uint16, b []byte) (gcp.Message, error) {
	m := gcp.Message{
		MessageID: id,
		Lenght:    uint16(len(b)),
	}
	rand.Seed(time.Now().UTC().UnixNano())
	switch n := gcp.MessageID(id); n {
	case gcp.MessageIDNotifyReq:
		m.Body = &gcp.NotifyReq{
			TransactionID: uint16(rand.Int()),
			// TODO: Change this.
			Mode: 0,
			// TODO: Change this.
			Status: 0,
			// TODO: Change this.
			EvntCode: 0,
			EvntData: b,
		}
		return m, nil
	case gcp.MessageIDNotifyRes:
		m.Body = &gcp.NotifyRes{
			TransactionID: tid,
			// TODO: Change this.
			Mode: 0,
			// TODO: Change this.
			EvntCode: 0,
		}
		return m, nil
	case gcp.MessageIDNotifyErr:
		m.Body = &gcp.NotifyErr{
			TransactionID: tid,
			// TODO: Change this.
			RtrnCode: 0,
		}
		return m, nil
	default:
		return m, fmt.Errorf("unrecognized Message ID: %d", id)
	}
}

const usage = `GCP: utility for working with the Generic Control Plane Protocol.
Examples:
  Listen for incoming GCP messages.
    $ ./gcp -m server
  Send a test message.
    $ ./gcp -m client -t ::1 -w test1
  `
