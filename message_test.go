package gcp_test

import (
	"encoding/base64"
	"fmt"
	"testing"

	gcp "github.com/nleiva/gcp-rphy"
)

var (
	host = "::1"
	port = "8190"
	// Pre-generated GCP Notify message for testing
	ntf = "AgFqAAHAAQAAAAEDAV8JAVwKAAIAAQsAAQIyAUkTASUBAAVDaXNjbwIAAgAJAwAIUlBIWS1SUEQEAAag+ElvQxwFAAR2Ni40BgBvUHJpbWFyeTogVS1Cb290IDIwMTYuMDEgKEp1bCAzMSAyMDE3IC0gMDk6NTQ6NTEgKzA4MDApICo7R29sZGVuOiBVLUJvb3QgMjAxNi4wMSAoQXByIDEyIDIwMTcgLSAwOToxMzoyOCArMDgwMCk7BwADUlBECAADUlBECQALQ0FUMjEzM0UwQTUKAAIRPQsACEJDTTMxNjEwDAADVjExDQAIMDAwMDAwMDAOAAMxLjAPAAYxLjAuMTAQAAMxLjARAAASAAATAAgH4wQCEjIqBRQAEFJQRC1WNi00Lml0Yi5TU0EVABAgAQV4EAAREQAAAAAAAAJFFgABABgAHgEAAk5BAgAJKzAwMDAwMC4wAwAKKzAwMDAwMDAuMFYABAEAAQE="
	// Pre-generated GCP RCP Object Exchange message for testing
	rex = "BwMgCEQAAAAAAAAAEYsBAgMRCQMOCgACAAoLAAEEEwABAGQC/ggAbgEAAQICAAR2YmgxAwAmVmlydHVhbCBCYWNraGF1bCBUZW4gR2lnYWJpdCBJbnRlcmZhY2UEAAIABgUAAAYABAAABdwHAAag+ElvQx0IAAEBCQABBwoABAAANV0LAAQAACcQDAABAg0AAQIOAAECCABuAQABAQIABHZiaDADACZWaXJ0dWFsIEJhY2toYXVsIFRlbiBHaWdhYml0IEludGVyZmFjZQQAAgAGBQAABgAEAAAF3AcABqD4SW9DHAgAAQEJAAEBCgAEAAA1uwsABAAAJxAMAAECDQABAg4AAQEPADEBAAQAAAABAgAECgAB/gMAAQQEAAEBBQACABgGAAEEBwABAQgABAAAAAAJAAQAAAAADwAxAQAEAAAAAQIABH8AAAEDAAEHBAABAQUAAgAIBgABBAcAAQEIAAQAAAAACQAEAAAAAA8AMQEABAAAAAECAATAqAEBAwABAwQAAQEFAAIAGAYAAQQHAAEBCAAEAAAAAAkABAAAAAAPAD0BAAQAAAACAgAQAAAAAAAAAAAAAAAAAAAAAQMAAQcEAAEBBQACAIAGAAEBBwABAQgABAAAAAAJAAQAAAAADwA9AQAEAAAAAgIAECABBXgQAAESAAAAAAAAAwEDAAEBBAABAQUAAgBABgABBAcAAQEIAAQAAFyfCQAEAABcnw8APQEABAAAAAICABD+gAAAAAAAAKL4Sf/+b0McAwABAQQAAQEFAAIAQAYAAQEHAAEBCAAEAABcnwkABAAAXJ8PAD0BAAQAAAACAgAQ/oAAAAAAAACi+En//m9DHQMAAQIEAAEBBQACAEAGAAEBBwABAQgABAAAAAAJAAQAAAAADwA9AQAEAAAAAgIAEP6AAAAAAAAAovhJ//5vQx4DAAEDBAABAQUAAgBABgABAQcAAQEIAAQAAAAACQAEAAAAAA8APQEABAAAAAICABD+gAAAAAAAAKgzEf/+ZgAAAwABBAQAAQEFAAIAQAYAAQEHAAEBCAAEAAAAAAkABAAAAAA="
	// Pre-generated GCP Identification and Resource Advertising message for testing
	ira = "BwB3Ni0AAAAAAAAAEYsBAQBoCQBlCgACAAELAAEFEwABABkAEwEAECABBXgQAHWoAAAAAAAAAAEZABMBABAgAQV4EAB1oAAAAAAAAAABGQATAQAQIAEFeBAAdaoAAAAAAAAAARkAEwEAECABBXgQAHWiAAAAAAAAAAE="
)

func TestParseMessage(t *testing.T) {
	tt := []struct {
		name    string
		message string
		Parser  func(g *gcp.GCP) error
	}{
		{name: "Notify", message: ntf, Parser: ParseNTF},
		{name: "RCP Object Exchange", message: rex, Parser: ParseREX},
		{name: "Identification and Resource Advertising", message: ira, Parser: ParseIRA},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			data, err := base64.StdEncoding.DecodeString(tc.message)
			if err != nil {
				t.Fatalf("could not decode base64 %s message: %v", tc.name, err)
			}
			msg, err := gcp.ParseMessage(data)
			if err != nil {
				t.Fatalf("could not parse %s message: %v", tc.name, err)
			}
			// output, g := msg.Body.Process()
			_, g := msg.Body.Process()

			err = tc.Parser(g)
			if err != nil {
				t.Fatalf("could not match value %s for %s message", err, tc.name)
			}
			// fmt.Printf("%s\n", output)
		})
	}

}

// ParseNTF validates parsing of a pre-generated NTF message.
/*
{
  "NTF": {
    "Sequence": {
      "Sequence Number": "1",
      "Operation": "Write",
      "RPD Capabilities": {
        "RpdIdentification": {
          "VendorName": "Cisco",
          "VendorId": "9",
          "ModelNumber": "RPHY-RPD",
          "DeviceMacAddress": "a0:f8:49:6f:43:1c",
          "CurrentSwVersion": "v6.4",
          "BootRomVersion": "Primary: U-Boot 2016.01 (Jul 31 2017 - 09:54:51 +0800) *;Golden: U-Boot 2016.01 (Apr 12 2017 - 09:13:28 +0800);",
          "DeviceDescription": "RPD",
          "DeviceAlias": "RPD",
          "SerialNumber": "CAT2133E0A5",
          "UsBurstReceiverVendorId": "4413",
          "UsBurstReceiverModelNumber": "BCM31610",
          "UsBurstReceiverDriverVersion": "V11",
          "UsBurstReceiverSerialNumber": "00000000",
          "RpdRcpProtocolVersion": "1.0",
          "RpdRcpSchemaVersion": "1.0.10",
          "HwRevision": "1.0",
          "CurrentSwImageLastUpdate": "2019-04-02 18:50:42.5 +0000 +0000",
          "CurrentSwImageName": "RPD-V6-4.itb.SSA",
          "CurrentSwImageServer": "2001:578:1000:1111::245",
          "CurrrentSwImageIndex": "0"
        },
        "Device Location": {
          "Device Location Description": "NA",
          "Geographic Location Latitude": "+000000.0",
          "Geographic Location Longitude": "+0000000.0"
        }
      },
      "General Notification": {
        "Type": "StartUpNotification"
      }
    }
  }
}
*/
func ParseNTF(g *gcp.GCP) error {
	if g.NTF.Sequence.SequenceNumber != "1" {
		return fmt.Errorf("SequenceNumber got: %v, want: %s", g.NTF.Sequence.SequenceNumber, "1")
	}
	RPDIdent := g.NTF.Sequence.RpdCapabilities.RpdIdentification
	if RPDIdent.VendorName != "Cisco" {
		return fmt.Errorf("VendorName got: %v, want: %s", RPDIdent.VendorName, "Cisco")
	}
	if RPDIdent.VendorID != "9" {
		return fmt.Errorf("VendorID got: %v, want: %s", RPDIdent.VendorID, "9")
	}
	if RPDIdent.ModelNumber != "RPHY-RPD" {
		return fmt.Errorf("ModelNumber got: %v, want: %s", RPDIdent.ModelNumber, "RPHY-RPD")
	}
	DevLocation := g.NTF.Sequence.RpdCapabilities.DeviceLocation
	if DevLocation.Description != "NA" {
		return fmt.Errorf("Location Description got: %v, want: %s", DevLocation.Description, "NA")
	}
	GnrlNtf := g.NTF.Sequence.GeneralNtf
	if GnrlNtf.NotificationType != "StartUpNotification" {
		return fmt.Errorf("Notification Type got: %v, want: %s", GnrlNtf.NotificationType, "StartUpNotification")
	}
	return nil
}

// ParseREX validates parsing of a pre-generated REX message.
/*
{
  "REX": {
    "Sequence": {
      "Sequence Number": "10",
      "Operation": "ReadResponse",
      "Response Code": "NoError",
      "RPD Info": {
        "IfEnet": [
          {
            "Port Index": "2",
            "Name": "vbh1",
            "Description": "Virtual Backhaul Ten Gigabit Interface",
            "Type": "ethernetCsmacd",
            "MTU": "1500",
            "Physical Address": "a0:f8:49:6f:43:1d",
            "Admin State": "up",
            "Operational State": "lowerLayerDown",
            "Last Change": "1970-01-16 14:28:20 -0500 EST",
            "Bandwidth": "10000 Mbps",
            "LinkUpDownTrapEnable": "false",
            "PromiscuousMode": "false"
          },
          {
            "Port Index": "1",
            "Name": "vbh0",
            "Description": "Virtual Backhaul Ten Gigabit Interface",
            "Type": "ethernetCsmacd",
            "MTU": "1500",
            "Physical Address": "a0:f8:49:6f:43:1c",
            "Admin State": "up",
            "Operational State": "up",
            "Last Change": "1970-01-16 17:05:00 -0500 EST",
            "Bandwidth": "10000 Mbps",
            "LinkUpDownTrapEnable": "false",
            "PromiscuousMode": "false",
            "Connector Present": true
          }
        ],
        "IpAddress": [
          {
            "Address Type": "ipv4",
            "IP Address": "10.0.1.254",
            "Port Index": "4",
            "Type": "unicast",
            "Prefix Length": "24",
            "Origin": "dhcp",
            "Status": "preferred",
            "Created": "0",
            "Last Changed": "0"
          },
          {
            "Address Type": "ipv4",
            "IP Address": "127.0.0.1",
            "Port Index": "7",
            "Type": "unicast",
            "Prefix Length": "8",
            "Origin": "dhcp",
            "Status": "preferred",
            "Created": "0",
            "Last Changed": "0"
          },
          {
            "Address Type": "ipv4",
            "IP Address": "192.168.1.1",
            "Port Index": "3",
            "Type": "unicast",
            "Prefix Length": "24",
            "Origin": "dhcp",
            "Status": "preferred",
            "Created": "0",
            "Last Changed": "0"
          },
          {
            "Address Type": "ipv6",
            "IP Address": "::1",
            "Port Index": "7",
            "Type": "unicast",
            "Prefix Length": "128",
            "Origin": "other",
            "Status": "preferred",
            "Created": "0",
            "Last Changed": "0"
          },
          {
            "Address Type": "ipv6",
            "IP Address": "2001:578:1000:112::301",
            "Port Index": "1",
            "Type": "unicast",
            "Prefix Length": "64",
            "Origin": "dhcp",
            "Status": "preferred",
            "Created": "1970-01-28 05:38:20 -0500 EST",
            "Last Changed": "1970-01-28 05:38:20 -0500 EST"
          },
          {
            "Address Type": "ipv6",
            "IP Address": "fe80::a2f8:49ff:fe6f:431c",
            "Port Index": "1",
            "Type": "unicast",
            "Prefix Length": "64",
            "Origin": "other",
            "Status": "preferred",
            "Created": "1970-01-28 05:38:20 -0500 EST",
            "Last Changed": "1970-01-28 05:38:20 -0500 EST"
          },
          {
            "Address Type": "ipv6",
            "IP Address": "fe80::a2f8:49ff:fe6f:431d",
            "Port Index": "2",
            "Type": "unicast",
            "Prefix Length": "64",
            "Origin": "other",
            "Status": "preferred",
            "Created": "0",
            "Last Changed": "0"
          },
          {
            "Address Type": "ipv6",
            "IP Address": "fe80::a2f8:49ff:fe6f:431e",
            "Port Index": "3",
            "Type": "unicast",
            "Prefix Length": "64",
            "Origin": "other",
            "Status": "preferred",
            "Created": "0",
            "Last Changed": "0"
          },
          {
            "Address Type": "ipv6",
            "IP Address": "fe80::a833:11ff:fe66:0",
            "Port Index": "4",
            "Type": "unicast",
            "Prefix Length": "64",
            "Origin": "other",
            "Status": "preferred",
            "Created": "0",
            "Last Changed": "0"
          }
        ]
      }
    }
  }
}
*/
func ParseREX(g *gcp.GCP) error {
	ifInet := g.REX.Sequence.RpdInfo.IfEnet
	if ifInet[0].Name != "vbh1" {
		return fmt.Errorf("IfEnet[0] Name got: %v, want: %s", ifInet[0].Name, "vbh1")
	}
	return nil
}

// ParseIRA validates parsing of a pre-generated IRA message.
/*
{
  "IRA": {
    "Sequence": {
      "Sequence Number": "1",
      "Operation": "WriteResponse",
      "Response Code": "NoError",
      "RPD Redirect": {
        "IP Address": [
          "2001:578:1000:75a8::1",
          "2001:578:1000:75a0::1",
          "2001:578:1000:75aa::1",
          "2001:578:1000:75a2::1"
        ]
      }
    }
  }
}
*/
func ParseIRA(g *gcp.GCP) error {
	if g.IRA.Sequence.Operation != "WriteResponse" {
		return fmt.Errorf("Operation got: %v, want: %s", g.IRA.Sequence.Operation, "WriteResponse")
	}
	IPAddr := g.IRA.Sequence.RpdRedirect.RpdRedirectIPAddress
	if IPAddr[0] != "2001:578:1000:75a8::1" {
		return fmt.Errorf("RPD Redirect IP Address got: %v, want: %s", IPAddr[0], "2001:578:1000:75a8::1")
	}
	return nil
}
