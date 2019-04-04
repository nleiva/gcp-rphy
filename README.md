# gcp-rphy

Go implementation of the Cable Lab's generic control plane protocol or GCP.

The latest version of the spec can be downloaded [here](https://specification-search.cablelabs.com/CM-SP-GCP).

## GCP CLI Examples

Sending a packet:

Sends the word `test` on a GCP Notify packet over TCP.

```bash
./gcp -m client -t ::1 -w test
```

Receiving a packet:

Listens on TCP port 8190 and prints out GCP packet received.

```bash
$ ./gcp -m server
Serving [::1]:65370
Incoming Message (Lenght: 23) ->
  Message Identifier: 2
  Length: 5
  Body:
	Transaction ID: 56320
	Mode: 0
	Status: 0
	Event Code: 116
	Event Data: [101 115 116 54]
2019/03/29 19:18:50 end of the transmition: EOF
```

## Reading list

Cable related:

- [Generic Control Plane Specification](https://specification-search.cablelabs.com/CM-SP-GCP)
- [MAC and Upper Layer Protocols Interface Specification](https://specification-search.cablelabs.com/CM-SP-MULPIv3.1)
- [DOCSIS 3.1 Physical Layer Specification](https://specification-search.cablelabs.com/CM-SP-PHYv3.1)

Go Network connections:

- [TCP/IP Networking](https://appliedgo.net/networking/). Also [tcpip](https://github.com/billglover/tcpip).
- [Building messaging in Go network clients](https://www.oreilly.com/ideas/building-messaging-in-go-network-clients).
- [Network Protocol Breakdown: NDP and Go](https://medium.com/@mdlayher/network-protocol-breakdown-ndp-and-go-3dc2900b1c20).