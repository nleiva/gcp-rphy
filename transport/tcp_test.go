package transport_test

import (
	"encoding/base64"
	"testing"
	"time"

	"github.com/nleiva/gcp-rphy/transport"
)

var (
	host = "::1"
	port = "8190"
	ntf  = "AgFqAAHAAQAAAAEDAV8JAVwKAAIAAQsAAQIyAUkTASUBAAVDaXNjbwIAAgAJAwAIUlBIWS1SUEQEAAag+ElvQxwFAAR2Ni40BgBvUHJpbWFyeTogVS1Cb290IDIwMTYuMDEgKEp1bCAzMSAyMDE3IC0gMDk6NTQ6NTEgKzA4MDApICo7R29sZGVuOiBVLUJvb3QgMjAxNi4wMSAoQXByIDEyIDIwMTcgLSAwOToxMzoyOCArMDgwMCk7BwADUlBECAADUlBECQALQ0FUMjEzM0UwQTUKAAIRPQsACEJDTTMxNjEwDAADVjExDQAIMDAwMDAwMDAOAAMxLjAPAAYxLjAuMTAQAAMxLjARAAASAAATAAgH4wQCEjIqBRQAEFJQRC1WNi00Lml0Yi5TU0EVABAgAQV4EAAREQAAAAAAAAJFFgABABgAHgEAAk5BAgAJKzAwMDAwMC4wAwAKKzAwMDAwMDAuMFYABAEAAQE="
)

func TestSendReceive(t *testing.T) {
	// Server
	go func() {
		server := transport.TCPEnd{
			Port: port,
		}
		err := server.Receive()
		if err != nil {
			t.Fatalf("could not setup a server: %v", err)
		}
	}()
	// Temp: Give enough time to the server to start listening for requests
	// time.Sleep(200 * time.Millisecond)

	// Client
	client := transport.TCPEnd{
		Host: host,
		Port: port,
	}

	data, err := base64.StdEncoding.DecodeString(ntf)
	if err != nil {
		t.Fatalf("could not decode base64 notify message: %v", err)
	}
	err = client.Send(data)
	if err != nil {
		t.Fatalf("could not send notify message: %v", err)
	}
	// Temp: Give enough time to print messages?
	time.Sleep(1000 * time.Millisecond)

}
