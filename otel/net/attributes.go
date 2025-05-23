package net

// Deprecated, use `network.local.address`
type AttrHostIp string // net.host.ip

func (AttrHostIp) Development() {}
func (AttrHostIp) Recommended() {}

// Deprecated, use `server.address`
type AttrHostName string // net.host.name

func (AttrHostName) Development() {}
func (AttrHostName) Recommended() {}

// Deprecated, use `server.port`
type AttrHostPort string // net.host.port

func (AttrHostPort) Development() {}
func (AttrHostPort) Recommended() {}

// Deprecated, use `network.peer.address`
type AttrPeerIp string // net.peer.ip

func (AttrPeerIp) Development() {}
func (AttrPeerIp) Recommended() {}

// Deprecated, use `server.address` on client spans and `client.address` on server spans
type AttrPeerName string // net.peer.name

func (AttrPeerName) Development() {}
func (AttrPeerName) Recommended() {}

// Deprecated, use `server.port` on client spans and `client.port` on server spans
type AttrPeerPort string // net.peer.port

func (AttrPeerPort) Development() {}
func (AttrPeerPort) Recommended() {}

// Deprecated, use `network.protocol.name`
type AttrProtocolName string // net.protocol.name

func (AttrProtocolName) Development() {}
func (AttrProtocolName) Recommended() {}

// Deprecated, use `network.protocol.version`
type AttrProtocolVersion string // net.protocol.version

func (AttrProtocolVersion) Development() {}
func (AttrProtocolVersion) Recommended() {}

// Deprecated, use `network.transport` and `network.type`
type AttrSockFamily string // net.sock.family

func (AttrSockFamily) Development() {}
func (AttrSockFamily) Recommended() {}

// Deprecated, use `network.local.address`
type AttrSockHostAddr string // net.sock.host.addr

func (AttrSockHostAddr) Development() {}
func (AttrSockHostAddr) Recommended() {}

// Deprecated, use `network.local.port`
type AttrSockHostPort string // net.sock.host.port

func (AttrSockHostPort) Development() {}
func (AttrSockHostPort) Recommended() {}

// Deprecated, use `network.peer.address`
type AttrSockPeerAddr string // net.sock.peer.addr

func (AttrSockPeerAddr) Development() {}
func (AttrSockPeerAddr) Recommended() {}

// Deprecated, no replacement at this time
type AttrSockPeerName string // net.sock.peer.name

func (AttrSockPeerName) Development() {}
func (AttrSockPeerName) Recommended() {}

// Deprecated, use `network.peer.port`
type AttrSockPeerPort string // net.sock.peer.port

func (AttrSockPeerPort) Development() {}
func (AttrSockPeerPort) Recommended() {}

// Deprecated, use `network.transport`
type AttrTransport string // net.transport

func (AttrTransport) Development() {}
func (AttrTransport) Recommended() {}

/* State {
    name: "attr.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "ctx": {
            "attributes": [
                {
                    "brief": "Deprecated, use `network.local.address`.",
                    "deprecated": {
                        "note": "Replaced by `network.local.address`.",
                        "reason": "renamed",
                        "renamed_to": "network.local.address",
                    },
                    "examples": "192.168.0.1",
                    "name": "net.host.ip",
                    "requirement_level": "recommended",
                    "root_namespace": "net",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Deprecated, use `server.address`.",
                    "deprecated": {
                        "note": "Replaced by `server.address`.",
                        "reason": "renamed",
                        "renamed_to": "server.address",
                    },
                    "examples": [
                        "example.com",
                    ],
                    "name": "net.host.name",
                    "requirement_level": "recommended",
                    "root_namespace": "net",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Deprecated, use `server.port`.",
                    "deprecated": {
                        "note": "Replaced by `server.port`.",
                        "reason": "renamed",
                        "renamed_to": "server.port",
                    },
                    "examples": [
                        8080,
                    ],
                    "name": "net.host.port",
                    "requirement_level": "recommended",
                    "root_namespace": "net",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "Deprecated, use `network.peer.address`.",
                    "deprecated": {
                        "note": "Replaced by `network.peer.address`.",
                        "reason": "renamed",
                        "renamed_to": "network.peer.address",
                    },
                    "examples": "127.0.0.1",
                    "name": "net.peer.ip",
                    "requirement_level": "recommended",
                    "root_namespace": "net",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Deprecated, use `server.address` on client spans and `client.address` on server spans.",
                    "deprecated": {
                        "note": "Replaced by `server.address` on client spans and `client.address` on server spans.",
                        "reason": "uncategorized",
                    },
                    "examples": [
                        "example.com",
                    ],
                    "name": "net.peer.name",
                    "requirement_level": "recommended",
                    "root_namespace": "net",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Deprecated, use `server.port` on client spans and `client.port` on server spans.",
                    "deprecated": {
                        "note": "Replaced by `server.port` on client spans and `client.port` on server spans.",
                        "reason": "uncategorized",
                    },
                    "examples": [
                        8080,
                    ],
                    "name": "net.peer.port",
                    "requirement_level": "recommended",
                    "root_namespace": "net",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "Deprecated, use `network.protocol.name`.",
                    "deprecated": {
                        "note": "Replaced by `network.protocol.name`.",
                        "reason": "renamed",
                        "renamed_to": "network.protocol.name",
                    },
                    "examples": [
                        "amqp",
                        "http",
                        "mqtt",
                    ],
                    "name": "net.protocol.name",
                    "requirement_level": "recommended",
                    "root_namespace": "net",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Deprecated, use `network.protocol.version`.",
                    "deprecated": {
                        "note": "Replaced by `network.protocol.version`.",
                        "reason": "renamed",
                        "renamed_to": "network.protocol.version",
                    },
                    "examples": "3.1.1",
                    "name": "net.protocol.version",
                    "requirement_level": "recommended",
                    "root_namespace": "net",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Deprecated, use `network.transport` and `network.type`.",
                    "deprecated": {
                        "note": "Split to `network.transport` and `network.type`.",
                        "reason": "uncategorized",
                    },
                    "name": "net.sock.family",
                    "requirement_level": "recommended",
                    "root_namespace": "net",
                    "stability": "development",
                    "type": {
                        "allow_custom_values": none,
                        "members": [
                            {
                                "brief": "IPv4 address",
                                "deprecated": none,
                                "id": "inet",
                                "note": none,
                                "stability": "development",
                                "value": "inet",
                            },
                            {
                                "brief": "IPv6 address",
                                "deprecated": none,
                                "id": "inet6",
                                "note": none,
                                "stability": "development",
                                "value": "inet6",
                            },
                            {
                                "brief": "Unix domain socket path",
                                "deprecated": none,
                                "id": "unix",
                                "note": none,
                                "stability": "development",
                                "value": "unix",
                            },
                        ],
                    },
                },
                {
                    "brief": "Deprecated, use `network.local.address`.",
                    "deprecated": {
                        "note": "Replaced by `network.local.address`.",
                        "reason": "renamed",
                        "renamed_to": "network.local.address",
                    },
                    "examples": [
                        "/var/my.sock",
                    ],
                    "name": "net.sock.host.addr",
                    "requirement_level": "recommended",
                    "root_namespace": "net",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Deprecated, use `network.local.port`.",
                    "deprecated": {
                        "note": "Replaced by `network.local.port`.",
                        "reason": "renamed",
                        "renamed_to": "network.local.port",
                    },
                    "examples": [
                        8080,
                    ],
                    "name": "net.sock.host.port",
                    "requirement_level": "recommended",
                    "root_namespace": "net",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "Deprecated, use `network.peer.address`.",
                    "deprecated": {
                        "note": "Replaced by `network.peer.address`.",
                        "reason": "renamed",
                        "renamed_to": "network.peer.address",
                    },
                    "examples": [
                        "192.168.0.1",
                    ],
                    "name": "net.sock.peer.addr",
                    "requirement_level": "recommended",
                    "root_namespace": "net",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Deprecated, no replacement at this time.",
                    "deprecated": {
                        "note": "Removed. No replacement at this time.",
                        "reason": "obsoleted",
                    },
                    "examples": [
                        "/var/my.sock",
                    ],
                    "name": "net.sock.peer.name",
                    "requirement_level": "recommended",
                    "root_namespace": "net",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Deprecated, use `network.peer.port`.",
                    "deprecated": {
                        "note": "Replaced by `network.peer.port`.",
                        "reason": "renamed",
                        "renamed_to": "network.peer.port",
                    },
                    "examples": [
                        65531,
                    ],
                    "name": "net.sock.peer.port",
                    "requirement_level": "recommended",
                    "root_namespace": "net",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "Deprecated, use `network.transport`.",
                    "deprecated": {
                        "note": "Replaced by `network.transport`.",
                        "reason": "renamed",
                        "renamed_to": "network.transport",
                    },
                    "name": "net.transport",
                    "requirement_level": "recommended",
                    "root_namespace": "net",
                    "stability": "development",
                    "type": {
                        "allow_custom_values": none,
                        "members": [
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "ip_tcp",
                                "note": none,
                                "stability": "development",
                                "value": "ip_tcp",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "ip_udp",
                                "note": none,
                                "stability": "development",
                                "value": "ip_udp",
                            },
                            {
                                "brief": "Named or anonymous pipe.",
                                "deprecated": none,
                                "id": "pipe",
                                "note": none,
                                "stability": "development",
                                "value": "pipe",
                            },
                            {
                                "brief": "In-process communication.",
                                "deprecated": none,
                                "id": "inproc",
                                "note": "Signals that there is only in-process communication not using a \"real\" network protocol in cases where network attributes would normally be expected. Usually all other network attributes can be left out in that case.\n",
                                "stability": "development",
                                "value": "inproc",
                            },
                            {
                                "brief": "Something else (non IP-based).",
                                "deprecated": none,
                                "id": "other",
                                "note": none,
                                "stability": "development",
                                "value": "other",
                            },
                        ],
                    },
                },
            ],
            "root_namespace": "net",
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
