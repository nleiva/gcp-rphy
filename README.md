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

## Sanity check

```bash
go test ./...
```

## Reading list

Cable related:

- [Generic Control Plane Specification](https://specification-search.cablelabs.com/CM-SP-GCP)
- [MAC and Upper Layer Protocols Interface Specification](https://specification-search.cablelabs.com/CM-SP-MULPIv3.1)
- [DOCSIS 3.1 Physical Layer Specification](https://specification-search.cablelabs.com/CM-SP-PHYv3.1)
- [Remote PHY OSS Interface Specification](https://specification-search.cablelabs.com/CM-SP-R-OSSI)

Go Network connections:

- [TCP/IP Networking](https://appliedgo.net/networking/). Also [tcpip](https://github.com/billglover/tcpip).
- [Building messaging in Go network clients](https://www.oreilly.com/ideas/building-messaging-in-go-network-clients).
- [Network Protocol Breakdown: NDP and Go](https://medium.com/@mdlayher/network-protocol-breakdown-ndp-and-go-3dc2900b1c20).

### GCP Usage (Normative)

From [Remote PHY Specification](https://specification-search.cablelabs.com/CM-SP-R-PHY).

GCP (Generic Control Plane) is described in [GCP]. GCP is fundamentally a control plane tunnel that allows data structures from other protocols to be reused in a new context. This is useful if there is configuration information that is well defined in an external specification. GCP can repurpose the information from other specifications rather than redefining it. For example, MHAv2 uses GCP to reuse predefined DOCSIS TLVs for configuration and operation of the RPD. GCP has three basic features:

- Device management, such as power management;
- Structured access, such as TLV tunneling;
- Diagnostic access.

GCP defines the structured access using a combination of:

- 32 bit Vendor ID as defined in [Vendor ID];
- 16 bit Structure ID as uniquely defined by the vendor. For MHAv2, the default vendor ID is the CableLabs vendor ID of 4491 (decimal).

When GCP tunnels the data structures of another protocol, the syntax GCP(protocol name) can be used.

### R-PHY Control Protocol

From [Remote PHY Specification](https://specification-search.cablelabs.com/CM-SP-R-PHY).

The following section defines the rules for the application of GCP as a Remote PHY control plane protocol. This set of rules is referred to as R-PHY Control Protocol or RCP.

RCP operates as an abstraction layer over the foundation of GCP protocol as defined in [GCP]. RCP provides the set of CCAP Core with to ability to remotely manage a set of objects, such as channels, ports, performance variables, etc.
RCP relies on the following GCP messages: Notify, Device Management and Exchange Data Structures. The
encodings of the GCP messages are provided in tables below.

### RCP Messages Types

- **IRA, Identification and Resource Advertising**: An initial message exchanged after authentication in which the CCAP Core obtains all parameters identifying the RPD and its available resources. Sent by CCAP Core in GCP EDS message.
- **REX, RCP Object Exchange**: A message in which CCAP Core allocates or deallocates resources and configures the resources in the RPD or requests information from the RPD, i.e., statistics or other status data. Sent by the CCAP Core in GCP EDS message. Responded to by the RPD when operation is complete.
- **NTF, Notification**: A message sent by the RPD to inform the CC
about a specific event or a set of events. Sent by the RPD in GCP Notify Message. CC does not respond to NTF messages.

### RCP Objects and TLVs

The RCP protocol operates on set of managed objects/TLVs sometimes referred to as ROTs (RCP Objects/TLVs). The ROTs are organized in a hierarchical tree. The top hierarchy consists of top level TLVs, which typically have a complex structure and are referred as Container ROTs. Container ROTs typically represent a set of managed attributes. The bottom of the hierarchy is formed from Leaf ROTs, which are scalars or strings that represent a single
managed attribute.

An RCP message consists of one or more Sequence(9) TLVs. A valid Sequence(9) TLV of a RCP message consists of:

- Exactly one SequenceNumber(10) TLV;
- Exactly one Operation(11) TLV;
- Exactly one top-level container object TLVs called the “Object Set-TLV”, except for a Read Response message that may contain multiple top-level TLVs expanding the wildcarded indexes of a Read Request top-level TLV;
- Optionally one ReadCount(26) TLV with an Operation(11) value of REX Read(1) or in an IRA message;
- Exactly one ResponseCode(19) in each Sequence(9) of a response message;
- Optionally one or more ErrorMessage(20) TLVs in each Sequence(9) of a response message.

A valid RCP Sequence(9) may contain its constituent TLVs in any order.

The top-level container object TLV in the sequence specifies one top-level container in the object hierarchy, but that container may contain an arbitrary set of subsidiary object hierarchy points.