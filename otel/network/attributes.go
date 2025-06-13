package network

// The ISO 3166-1 alpha-2 2-character country code associated with the mobile carrier network
type AttrCarrierIcc string // network.carrier.icc

func (AttrCarrierIcc) Development()    {}
func (AttrCarrierIcc) Recommended()    {}
func (AttrCarrierIcc) Key() string     { return "network_carrier_icc" }
func (a AttrCarrierIcc) Value() string { return string(a) }

// The mobile carrier country code
type AttrCarrierMcc string // network.carrier.mcc

func (AttrCarrierMcc) Development()    {}
func (AttrCarrierMcc) Recommended()    {}
func (AttrCarrierMcc) Key() string     { return "network_carrier_mcc" }
func (a AttrCarrierMcc) Value() string { return string(a) }

// The mobile carrier network code
type AttrCarrierMnc string // network.carrier.mnc

func (AttrCarrierMnc) Development()    {}
func (AttrCarrierMnc) Recommended()    {}
func (AttrCarrierMnc) Key() string     { return "network_carrier_mnc" }
func (a AttrCarrierMnc) Value() string { return string(a) }

// The name of the mobile carrier
type AttrCarrierName string // network.carrier.name

func (AttrCarrierName) Development()    {}
func (AttrCarrierName) Recommended()    {}
func (AttrCarrierName) Key() string     { return "network_carrier_name" }
func (a AttrCarrierName) Value() string { return string(a) }

// The state of network connection
// Connection states are defined as part of the [rfc9293]
//
// [rfc9293]: https://datatracker.ietf.org/doc/html/rfc9293#section-3.3.2
type AttrConnectionState string // network.connection.state

func (AttrConnectionState) Development()    {}
func (AttrConnectionState) Recommended()    {}
func (AttrConnectionState) Key() string     { return "network_connection_state" }
func (a AttrConnectionState) Value() string { return string(a) }

const ConnectionStateClosed AttrConnectionState = "closed"
const ConnectionStateCloseWait AttrConnectionState = "close_wait"
const ConnectionStateClosing AttrConnectionState = "closing"
const ConnectionStateEstablished AttrConnectionState = "established"
const ConnectionStateFinWait1 AttrConnectionState = "fin_wait_1"
const ConnectionStateFinWait2 AttrConnectionState = "fin_wait_2"
const ConnectionStateLastAck AttrConnectionState = "last_ack"
const ConnectionStateListen AttrConnectionState = "listen"
const ConnectionStateSynReceived AttrConnectionState = "syn_received"
const ConnectionStateSynSent AttrConnectionState = "syn_sent"
const ConnectionStateTimeWait AttrConnectionState = "time_wait"

// This describes more details regarding the connection.type. It may be the type of cell technology connection, but it could be used for describing details about a wifi connection
type AttrConnectionSubtype string // network.connection.subtype

func (AttrConnectionSubtype) Development()    {}
func (AttrConnectionSubtype) Recommended()    {}
func (AttrConnectionSubtype) Key() string     { return "network_connection_subtype" }
func (a AttrConnectionSubtype) Value() string { return string(a) }

const ConnectionSubtypeGprs AttrConnectionSubtype = "gprs"
const ConnectionSubtypeEdge AttrConnectionSubtype = "edge"
const ConnectionSubtypeUmts AttrConnectionSubtype = "umts"
const ConnectionSubtypeCdma AttrConnectionSubtype = "cdma"
const ConnectionSubtypeEvdo0 AttrConnectionSubtype = "evdo_0"
const ConnectionSubtypeEvdoA AttrConnectionSubtype = "evdo_a"
const ConnectionSubtypeCdma20001xrtt AttrConnectionSubtype = "cdma2000_1xrtt"
const ConnectionSubtypeHsdpa AttrConnectionSubtype = "hsdpa"
const ConnectionSubtypeHsupa AttrConnectionSubtype = "hsupa"
const ConnectionSubtypeHspa AttrConnectionSubtype = "hspa"
const ConnectionSubtypeIden AttrConnectionSubtype = "iden"
const ConnectionSubtypeEvdoB AttrConnectionSubtype = "evdo_b"
const ConnectionSubtypeLte AttrConnectionSubtype = "lte"
const ConnectionSubtypeEhrpd AttrConnectionSubtype = "ehrpd"
const ConnectionSubtypeHspap AttrConnectionSubtype = "hspap"
const ConnectionSubtypeGsm AttrConnectionSubtype = "gsm"
const ConnectionSubtypeTdScdma AttrConnectionSubtype = "td_scdma"
const ConnectionSubtypeIwlan AttrConnectionSubtype = "iwlan"
const ConnectionSubtypeNr AttrConnectionSubtype = "nr"
const ConnectionSubtypeNrnsa AttrConnectionSubtype = "nrnsa"
const ConnectionSubtypeLteCa AttrConnectionSubtype = "lte_ca"

// The internet connection type
type AttrConnectionType string // network.connection.type

func (AttrConnectionType) Development()    {}
func (AttrConnectionType) Recommended()    {}
func (AttrConnectionType) Key() string     { return "network_connection_type" }
func (a AttrConnectionType) Value() string { return string(a) }

const ConnectionTypeWifi AttrConnectionType = "wifi"
const ConnectionTypeWired AttrConnectionType = "wired"
const ConnectionTypeCell AttrConnectionType = "cell"
const ConnectionTypeUnavailable AttrConnectionType = "unavailable"
const ConnectionTypeUnknown AttrConnectionType = "unknown"

// The network interface name
type AttrInterfaceName string // network.interface.name

func (AttrInterfaceName) Development()    {}
func (AttrInterfaceName) Recommended()    {}
func (AttrInterfaceName) Key() string     { return "network_interface_name" }
func (a AttrInterfaceName) Value() string { return string(a) }

// The network IO operation direction
type AttrIoDirection string // network.io.direction

func (AttrIoDirection) Development()    {}
func (AttrIoDirection) Recommended()    {}
func (AttrIoDirection) Key() string     { return "network_io_direction" }
func (a AttrIoDirection) Value() string { return string(a) }

const IoDirectionTransmit AttrIoDirection = "transmit"
const IoDirectionReceive AttrIoDirection = "receive"

// Local address of the network connection - IP address or Unix domain socket name
type AttrLocalAddress string // network.local.address

func (AttrLocalAddress) Stable()         {}
func (AttrLocalAddress) Recommended()    {}
func (AttrLocalAddress) Key() string     { return "network_local_address" }
func (a AttrLocalAddress) Value() string { return string(a) }

// Local port number of the network connection
type AttrLocalPort string // network.local.port

func (AttrLocalPort) Stable()         {}
func (AttrLocalPort) Recommended()    {}
func (AttrLocalPort) Key() string     { return "network_local_port" }
func (a AttrLocalPort) Value() string { return string(a) }

// Peer address of the network connection - IP address or Unix domain socket name
type AttrPeerAddress string // network.peer.address

func (AttrPeerAddress) Stable()         {}
func (AttrPeerAddress) Recommended()    {}
func (AttrPeerAddress) Key() string     { return "network_peer_address" }
func (a AttrPeerAddress) Value() string { return string(a) }

// Peer port number of the network connection
type AttrPeerPort string // network.peer.port

func (AttrPeerPort) Stable()         {}
func (AttrPeerPort) Recommended()    {}
func (AttrPeerPort) Key() string     { return "network_peer_port" }
func (a AttrPeerPort) Value() string { return string(a) }

// [OSI application layer] or non-OSI equivalent.
// The value SHOULD be normalized to lowercase
//
// [OSI application layer]: https://wikipedia.org/wiki/Application_layer
type AttrProtocolName string // network.protocol.name

func (AttrProtocolName) Stable()         {}
func (AttrProtocolName) Recommended()    {}
func (AttrProtocolName) Key() string     { return "network_protocol_name" }
func (a AttrProtocolName) Value() string { return string(a) }

// The actual version of the protocol used for network communication.
// If protocol version is subject to negotiation (for example using [ALPN]), this attribute SHOULD be set to the negotiated version. If the actual protocol version is not known, this attribute SHOULD NOT be set
//
// [ALPN]: https://www.rfc-editor.org/rfc/rfc7301.html
type AttrProtocolVersion string // network.protocol.version

func (AttrProtocolVersion) Stable()         {}
func (AttrProtocolVersion) Recommended()    {}
func (AttrProtocolVersion) Key() string     { return "network_protocol_version" }
func (a AttrProtocolVersion) Value() string { return string(a) }

// [OSI transport layer] or [inter-process communication method].
//
// The value SHOULD be normalized to lowercase.
//
// Consider always setting the transport when setting a port number, since
// a port number is ambiguous without knowing the transport. For example
// different processes could be listening on TCP port 12345 and UDP port 12345
//
// [OSI transport layer]: https://wikipedia.org/wiki/Transport_layer
// [inter-process communication method]: https://wikipedia.org/wiki/Inter-process_communication
type AttrTransport string // network.transport

func (AttrTransport) Stable()         {}
func (AttrTransport) Recommended()    {}
func (AttrTransport) Key() string     { return "network_transport" }
func (a AttrTransport) Value() string { return string(a) }

const TransportTcp AttrTransport = "tcp"
const TransportUdp AttrTransport = "udp"
const TransportPipe AttrTransport = "pipe"
const TransportUnix AttrTransport = "unix"
const TransportQuic AttrTransport = "quic"

// [OSI network layer] or non-OSI equivalent.
// The value SHOULD be normalized to lowercase
//
// [OSI network layer]: https://wikipedia.org/wiki/Network_layer
type AttrType string // network.type

func (AttrType) Stable()         {}
func (AttrType) Recommended()    {}
func (AttrType) Key() string     { return "network_type" }
func (a AttrType) Value() string { return string(a) }

const TypeIpv4 AttrType = "ipv4"
const TypeIpv6 AttrType = "ipv6"

/* State {
    name: "attr.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "ctx": {
            "attributes": [
                {
                    "brief": "The ISO 3166-1 alpha-2 2-character country code associated with the mobile carrier network.",
                    "examples": "DE",
                    "name": "network.carrier.icc",
                    "requirement_level": "recommended",
                    "root_namespace": "network",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The mobile carrier country code.",
                    "examples": "310",
                    "name": "network.carrier.mcc",
                    "requirement_level": "recommended",
                    "root_namespace": "network",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The mobile carrier network code.",
                    "examples": "001",
                    "name": "network.carrier.mnc",
                    "requirement_level": "recommended",
                    "root_namespace": "network",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The name of the mobile carrier.",
                    "examples": "sprint",
                    "name": "network.carrier.name",
                    "requirement_level": "recommended",
                    "root_namespace": "network",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The state of network connection",
                    "examples": [
                        "close_wait",
                    ],
                    "name": "network.connection.state",
                    "note": "Connection states are defined as part of the [rfc9293](https://datatracker.ietf.org/doc/html/rfc9293#section-3.3.2)",
                    "requirement_level": "recommended",
                    "root_namespace": "network",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "closed",
                                "note": none,
                                "stability": "development",
                                "value": "closed",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "close_wait",
                                "note": none,
                                "stability": "development",
                                "value": "close_wait",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "closing",
                                "note": none,
                                "stability": "development",
                                "value": "closing",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "established",
                                "note": none,
                                "stability": "development",
                                "value": "established",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "fin_wait_1",
                                "note": none,
                                "stability": "development",
                                "value": "fin_wait_1",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "fin_wait_2",
                                "note": none,
                                "stability": "development",
                                "value": "fin_wait_2",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "last_ack",
                                "note": none,
                                "stability": "development",
                                "value": "last_ack",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "listen",
                                "note": none,
                                "stability": "development",
                                "value": "listen",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "syn_received",
                                "note": none,
                                "stability": "development",
                                "value": "syn_received",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "syn_sent",
                                "note": none,
                                "stability": "development",
                                "value": "syn_sent",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "time_wait",
                                "note": none,
                                "stability": "development",
                                "value": "time_wait",
                            },
                        ],
                    },
                },
                {
                    "brief": "This describes more details regarding the connection.type. It may be the type of cell technology connection, but it could be used for describing details about a wifi connection.",
                    "examples": "LTE",
                    "name": "network.connection.subtype",
                    "requirement_level": "recommended",
                    "root_namespace": "network",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": "GPRS",
                                "deprecated": none,
                                "id": "gprs",
                                "note": none,
                                "stability": "development",
                                "value": "gprs",
                            },
                            {
                                "brief": "EDGE",
                                "deprecated": none,
                                "id": "edge",
                                "note": none,
                                "stability": "development",
                                "value": "edge",
                            },
                            {
                                "brief": "UMTS",
                                "deprecated": none,
                                "id": "umts",
                                "note": none,
                                "stability": "development",
                                "value": "umts",
                            },
                            {
                                "brief": "CDMA",
                                "deprecated": none,
                                "id": "cdma",
                                "note": none,
                                "stability": "development",
                                "value": "cdma",
                            },
                            {
                                "brief": "EVDO Rel. 0",
                                "deprecated": none,
                                "id": "evdo_0",
                                "note": none,
                                "stability": "development",
                                "value": "evdo_0",
                            },
                            {
                                "brief": "EVDO Rev. A",
                                "deprecated": none,
                                "id": "evdo_a",
                                "note": none,
                                "stability": "development",
                                "value": "evdo_a",
                            },
                            {
                                "brief": "CDMA2000 1XRTT",
                                "deprecated": none,
                                "id": "cdma2000_1xrtt",
                                "note": none,
                                "stability": "development",
                                "value": "cdma2000_1xrtt",
                            },
                            {
                                "brief": "HSDPA",
                                "deprecated": none,
                                "id": "hsdpa",
                                "note": none,
                                "stability": "development",
                                "value": "hsdpa",
                            },
                            {
                                "brief": "HSUPA",
                                "deprecated": none,
                                "id": "hsupa",
                                "note": none,
                                "stability": "development",
                                "value": "hsupa",
                            },
                            {
                                "brief": "HSPA",
                                "deprecated": none,
                                "id": "hspa",
                                "note": none,
                                "stability": "development",
                                "value": "hspa",
                            },
                            {
                                "brief": "IDEN",
                                "deprecated": none,
                                "id": "iden",
                                "note": none,
                                "stability": "development",
                                "value": "iden",
                            },
                            {
                                "brief": "EVDO Rev. B",
                                "deprecated": none,
                                "id": "evdo_b",
                                "note": none,
                                "stability": "development",
                                "value": "evdo_b",
                            },
                            {
                                "brief": "LTE",
                                "deprecated": none,
                                "id": "lte",
                                "note": none,
                                "stability": "development",
                                "value": "lte",
                            },
                            {
                                "brief": "EHRPD",
                                "deprecated": none,
                                "id": "ehrpd",
                                "note": none,
                                "stability": "development",
                                "value": "ehrpd",
                            },
                            {
                                "brief": "HSPAP",
                                "deprecated": none,
                                "id": "hspap",
                                "note": none,
                                "stability": "development",
                                "value": "hspap",
                            },
                            {
                                "brief": "GSM",
                                "deprecated": none,
                                "id": "gsm",
                                "note": none,
                                "stability": "development",
                                "value": "gsm",
                            },
                            {
                                "brief": "TD-SCDMA",
                                "deprecated": none,
                                "id": "td_scdma",
                                "note": none,
                                "stability": "development",
                                "value": "td_scdma",
                            },
                            {
                                "brief": "IWLAN",
                                "deprecated": none,
                                "id": "iwlan",
                                "note": none,
                                "stability": "development",
                                "value": "iwlan",
                            },
                            {
                                "brief": "5G NR (New Radio)",
                                "deprecated": none,
                                "id": "nr",
                                "note": none,
                                "stability": "development",
                                "value": "nr",
                            },
                            {
                                "brief": "5G NRNSA (New Radio Non-Standalone)",
                                "deprecated": none,
                                "id": "nrnsa",
                                "note": none,
                                "stability": "development",
                                "value": "nrnsa",
                            },
                            {
                                "brief": "LTE CA",
                                "deprecated": none,
                                "id": "lte_ca",
                                "note": none,
                                "stability": "development",
                                "value": "lte_ca",
                            },
                        ],
                    },
                },
                {
                    "brief": "The internet connection type.",
                    "examples": "wifi",
                    "name": "network.connection.type",
                    "requirement_level": "recommended",
                    "root_namespace": "network",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "wifi",
                                "note": none,
                                "stability": "development",
                                "value": "wifi",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "wired",
                                "note": none,
                                "stability": "development",
                                "value": "wired",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "cell",
                                "note": none,
                                "stability": "development",
                                "value": "cell",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "unavailable",
                                "note": none,
                                "stability": "development",
                                "value": "unavailable",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "unknown",
                                "note": none,
                                "stability": "development",
                                "value": "unknown",
                            },
                        ],
                    },
                },
                {
                    "brief": "The network interface name.",
                    "examples": [
                        "lo",
                        "eth0",
                    ],
                    "name": "network.interface.name",
                    "requirement_level": "recommended",
                    "root_namespace": "network",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The network IO operation direction.",
                    "examples": [
                        "transmit",
                    ],
                    "name": "network.io.direction",
                    "requirement_level": "recommended",
                    "root_namespace": "network",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "transmit",
                                "note": none,
                                "stability": "development",
                                "value": "transmit",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "receive",
                                "note": none,
                                "stability": "development",
                                "value": "receive",
                            },
                        ],
                    },
                },
                {
                    "brief": "Local address of the network connection - IP address or Unix domain socket name.",
                    "examples": [
                        "10.1.2.80",
                        "/tmp/my.sock",
                    ],
                    "name": "network.local.address",
                    "requirement_level": "recommended",
                    "root_namespace": "network",
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "Local port number of the network connection.",
                    "examples": [
                        65123,
                    ],
                    "name": "network.local.port",
                    "requirement_level": "recommended",
                    "root_namespace": "network",
                    "stability": "stable",
                    "type": "int",
                },
                {
                    "brief": "Peer address of the network connection - IP address or Unix domain socket name.",
                    "examples": [
                        "10.1.2.80",
                        "/tmp/my.sock",
                    ],
                    "name": "network.peer.address",
                    "requirement_level": "recommended",
                    "root_namespace": "network",
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "Peer port number of the network connection.",
                    "examples": [
                        65123,
                    ],
                    "name": "network.peer.port",
                    "requirement_level": "recommended",
                    "root_namespace": "network",
                    "stability": "stable",
                    "type": "int",
                },
                {
                    "brief": "[OSI application layer](https://wikipedia.org/wiki/Application_layer) or non-OSI equivalent.",
                    "examples": [
                        "amqp",
                        "http",
                        "mqtt",
                    ],
                    "name": "network.protocol.name",
                    "note": "The value SHOULD be normalized to lowercase.",
                    "requirement_level": "recommended",
                    "root_namespace": "network",
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "The actual version of the protocol used for network communication.",
                    "examples": [
                        "1.1",
                        "2",
                    ],
                    "name": "network.protocol.version",
                    "note": "If protocol version is subject to negotiation (for example using [ALPN](https://www.rfc-editor.org/rfc/rfc7301.html)), this attribute SHOULD be set to the negotiated version. If the actual protocol version is not known, this attribute SHOULD NOT be set.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "network",
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "[OSI transport layer](https://wikipedia.org/wiki/Transport_layer) or [inter-process communication method](https://wikipedia.org/wiki/Inter-process_communication).\n",
                    "examples": [
                        "tcp",
                        "udp",
                    ],
                    "name": "network.transport",
                    "note": "The value SHOULD be normalized to lowercase.\n\nConsider always setting the transport when setting a port number, since\na port number is ambiguous without knowing the transport. For example\ndifferent processes could be listening on TCP port 12345 and UDP port 12345.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "network",
                    "stability": "stable",
                    "type": {
                        "members": [
                            {
                                "brief": "TCP",
                                "deprecated": none,
                                "id": "tcp",
                                "note": none,
                                "stability": "stable",
                                "value": "tcp",
                            },
                            {
                                "brief": "UDP",
                                "deprecated": none,
                                "id": "udp",
                                "note": none,
                                "stability": "stable",
                                "value": "udp",
                            },
                            {
                                "brief": "Named or anonymous pipe.",
                                "deprecated": none,
                                "id": "pipe",
                                "note": none,
                                "stability": "stable",
                                "value": "pipe",
                            },
                            {
                                "brief": "Unix domain socket",
                                "deprecated": none,
                                "id": "unix",
                                "note": none,
                                "stability": "stable",
                                "value": "unix",
                            },
                            {
                                "brief": "QUIC",
                                "deprecated": none,
                                "id": "quic",
                                "note": none,
                                "stability": "stable",
                                "value": "quic",
                            },
                        ],
                    },
                },
                {
                    "brief": "[OSI network layer](https://wikipedia.org/wiki/Network_layer) or non-OSI equivalent.",
                    "examples": [
                        "ipv4",
                        "ipv6",
                    ],
                    "name": "network.type",
                    "note": "The value SHOULD be normalized to lowercase.",
                    "requirement_level": "recommended",
                    "root_namespace": "network",
                    "stability": "stable",
                    "type": {
                        "members": [
                            {
                                "brief": "IPv4",
                                "deprecated": none,
                                "id": "ipv4",
                                "note": none,
                                "stability": "stable",
                                "value": "ipv4",
                            },
                            {
                                "brief": "IPv6",
                                "deprecated": none,
                                "id": "ipv6",
                                "note": none,
                                "stability": "stable",
                                "value": "ipv6",
                            },
                        ],
                    },
                },
            ],
            "root_namespace": "network",
        },
    },
    env: Environment {
        globals: {
            "concat_if": weaver_forge::extensions::util::concat_if,
            "cycler": minijinja_contrib::globals::cycler,
            "debug": minijinja::functions::builtins::debug,
            "dict": minijinja::functions::builtins::dict,
            "joiner": minijinja_contrib::globals::joiner,
            "namespace": minijinja::functions::builtins::namespace,
            "params": {
                "params": {},
            },
            "range": minijinja::functions::builtins::range,
            "template": {},
        },
        tests: [
            "!=",
            "<",
            "<=",
            "==",
            ">",
            ">=",
            "array",
            "boolean",
            "defined",
            "deprecated",
            "divisibleby",
            "endingwith",
            "enum",
            "enum_type",
            "eq",
            "equalto",
            "escaped",
            "even",
            "experimental",
            "false",
            "filter",
            "float",
            "ge",
            "greaterthan",
            "gt",
            "in",
            "int",
            "integer",
            "iterable",
            "le",
            "lessthan",
            "lower",
            "lt",
            "mapping",
            "ne",
            "none",
            "number",
            "odd",
            "safe",
            "sameas",
            "sequence",
            "simple_type",
            "stable",
            "startingwith",
            "string",
            "template_type",
            "test",
            "true",
            "undefined",
            "upper",
        ],
        filters: [
            "abs",
            "acronym",
            "ansi_bg_black",
            "ansi_bg_blue",
            "ansi_bg_bright_black",
            "ansi_bg_bright_blue",
            "ansi_bg_bright_cyan",
            "ansi_bg_bright_green",
            "ansi_bg_bright_magenta",
            "ansi_bg_bright_red",
            "ansi_bg_bright_white",
            "ansi_bg_bright_yellow",
            "ansi_bg_cyan",
            "ansi_bg_green",
            "ansi_bg_magenta",
            "ansi_bg_red",
            "ansi_bg_white",
            "ansi_bg_yellow",
            "ansi_black",
            "ansi_blue",
            "ansi_bold",
            "ansi_bright_black",
            "ansi_bright_blue",
            "ansi_bright_cyan",
            "ansi_bright_green",
            "ansi_bright_magenta",
            "ansi_bright_red",
            "ansi_bright_white",
            "ansi_bright_yellow",
            "ansi_cyan",
            "ansi_green",
            "ansi_italic",
            "ansi_magenta",
            "ansi_red",
            "ansi_strikethrough",
            "ansi_underline",
            "ansi_white",
            "ansi_yellow",
            "attr",
            "attribute_id",
            "attribute_namespace",
            "attribute_registry_file",
            "attribute_registry_namespace",
            "attribute_registry_title",
            "attribute_sort",
            "batch",
            "body_fields",
            "bool",
            "camel_case",
            "camel_case_const",
            "capitalize",
            "capitalize_first",
            "comment",
            "comment_with_prefix",
            "count",
            "d",
            "default",
            "dictsort",
            "e",
            "enum_type",
            "escape",
            "filesizeformat",
            "first",
            "flatten",
            "float",
            "groupby",
            "indent",
            "instantiated_type",
            "int",
            "items",
            "join",
            "kebab_case",
            "kebab_case_const",
            "last",
            "length",
            "lines",
            "list",
            "lower",
            "lower_case",
            "map",
            "map_text",
            "markdown_to_html",
            "max",
            "metric_namespace",
            "min",
            "not_required",
            "pascal_case",
            "pascal_case_const",
            "pluralize",
            "pprint",
            "print_member_value",
            "regex_replace",
            "reject",
            "rejectattr",
            "replace",
            "required",
            "reverse",
            "round",
            "safe",
            "screaming_kebab_case",
            "screaming_snake_case",
            "screaming_snake_case_const",
            "select",
            "selectattr",
            "slice",
            "snake_case",
            "snake_case_const",
            "sort",
            "split",
            "split_id",
            "string",
            "striptags",
            "sum",
            "title",
            "title_case",
            "tojson",
            "toyaml",
            "trim",
            "truncate",
            "type_mapping",
            "unique",
            "upper",
            "upper_case",
            "urlencode",
        ],
        templates: [
            "attr.go.j2",
        ],
    },
} */
