package host

// The CPU architecture the host system is running on
type AttrArch string // host.arch

func (AttrArch) Development() {}
func (AttrArch) Recommended() {}

// The amount of level 2 memory cache available to the processor (in Bytes)
type AttrCpuCacheL2Size string // host.cpu.cache.l2.size

func (AttrCpuCacheL2Size) Development() {}
func (AttrCpuCacheL2Size) Recommended() {}

// Family or generation of the CPU
type AttrCpuFamily string // host.cpu.family

func (AttrCpuFamily) Development() {}
func (AttrCpuFamily) Recommended() {}

// Model identifier. It provides more granular information about the CPU, distinguishing it from other CPUs within the same family
type AttrCpuModelId string // host.cpu.model.id

func (AttrCpuModelId) Development() {}
func (AttrCpuModelId) Recommended() {}

// Model designation of the processor
type AttrCpuModelName string // host.cpu.model.name

func (AttrCpuModelName) Development() {}
func (AttrCpuModelName) Recommended() {}

// Stepping or core revisions
type AttrCpuStepping string // host.cpu.stepping

func (AttrCpuStepping) Development() {}
func (AttrCpuStepping) Recommended() {}

// Processor manufacturer identifier. A maximum 12-character string.
//
// [CPUID] command returns the vendor ID string in EBX, EDX and ECX registers. Writing these to memory in this order results in a 12-character string
//
// [CPUID]: https://wiki.osdev.org/CPUID
type AttrCpuVendorId string // host.cpu.vendor.id

func (AttrCpuVendorId) Development() {}
func (AttrCpuVendorId) Recommended() {}

// Unique host ID. For Cloud, this must be the instance_id assigned by the cloud provider. For non-containerized systems, this should be the `machine-id`. See the table below for the sources to use to determine the `machine-id` based on operating system
type AttrId string // host.id

func (AttrId) Development() {}
func (AttrId) Recommended() {}

// VM image ID or host OS image ID. For Cloud, this value is from the provider
type AttrImageId string // host.image.id

func (AttrImageId) Development() {}
func (AttrImageId) Recommended() {}

// Name of the VM image or OS install the host was instantiated from
type AttrImageName string // host.image.name

func (AttrImageName) Development() {}
func (AttrImageName) Recommended() {}

// The version string of the VM image or host OS as defined in [Version Attributes]
//
// [Version Attributes]: /docs/resource/README.md#version-attributes
type AttrImageVersion string // host.image.version

func (AttrImageVersion) Development() {}
func (AttrImageVersion) Recommended() {}

// Available IP addresses of the host, excluding loopback interfaces.
//
// IPv4 Addresses MUST be specified in dotted-quad notation. IPv6 addresses MUST be specified in the [RFC 5952] format
//
// [RFC 5952]: https://www.rfc-editor.org/rfc/rfc5952.html
type AttrIp string // host.ip

func (AttrIp) Development() {}
func (AttrIp) Recommended() {}

// Available MAC addresses of the host, excluding loopback interfaces.
//
// MAC Addresses MUST be represented in [IEEE RA hexadecimal form]: as hyphen-separated octets in uppercase hexadecimal form from most to least significant
//
// [IEEE RA hexadecimal form]: https://standards.ieee.org/wp-content/uploads/import/documents/tutorials/eui.pdf
type AttrMac string // host.mac

func (AttrMac) Development() {}
func (AttrMac) Recommended() {}

// Name of the host. On Unix systems, it may contain what the hostname command returns, or the fully qualified hostname, or another name specified by the user
type AttrName string // host.name

func (AttrName) Development() {}
func (AttrName) Recommended() {}

// Type of host. For Cloud, this must be the machine type
type AttrType string // host.type

func (AttrType) Development() {}
func (AttrType) Recommended() {}

/* State {
    name: "attr.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "ctx": {
            "attributes": [
                {
                    "brief": "The CPU architecture the host system is running on.\n",
                    "name": "host.arch",
                    "requirement_level": "recommended",
                    "root_namespace": "host",
                    "stability": "development",
                    "type": {
                        "allow_custom_values": none,
                        "members": [
                            {
                                "brief": "AMD64",
                                "deprecated": none,
                                "id": "amd64",
                                "note": none,
                                "stability": "development",
                                "value": "amd64",
                            },
                            {
                                "brief": "ARM32",
                                "deprecated": none,
                                "id": "arm32",
                                "note": none,
                                "stability": "development",
                                "value": "arm32",
                            },
                            {
                                "brief": "ARM64",
                                "deprecated": none,
                                "id": "arm64",
                                "note": none,
                                "stability": "development",
                                "value": "arm64",
                            },
                            {
                                "brief": "Itanium",
                                "deprecated": none,
                                "id": "ia64",
                                "note": none,
                                "stability": "development",
                                "value": "ia64",
                            },
                            {
                                "brief": "32-bit PowerPC",
                                "deprecated": none,
                                "id": "ppc32",
                                "note": none,
                                "stability": "development",
                                "value": "ppc32",
                            },
                            {
                                "brief": "64-bit PowerPC",
                                "deprecated": none,
                                "id": "ppc64",
                                "note": none,
                                "stability": "development",
                                "value": "ppc64",
                            },
                            {
                                "brief": "IBM z/Architecture",
                                "deprecated": none,
                                "id": "s390x",
                                "note": none,
                                "stability": "development",
                                "value": "s390x",
                            },
                            {
                                "brief": "32-bit x86",
                                "deprecated": none,
                                "id": "x86",
                                "note": none,
                                "stability": "development",
                                "value": "x86",
                            },
                        ],
                    },
                },
                {
                    "brief": "The amount of level 2 memory cache available to the processor (in Bytes).\n",
                    "examples": [
                        12288000,
                    ],
                    "name": "host.cpu.cache.l2.size",
                    "requirement_level": "recommended",
                    "root_namespace": "host",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "Family or generation of the CPU.\n",
                    "examples": [
                        "6",
                        "PA-RISC 1.1e",
                    ],
                    "name": "host.cpu.family",
                    "requirement_level": "recommended",
                    "root_namespace": "host",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Model identifier. It provides more granular information about the CPU, distinguishing it from other CPUs within the same family.\n",
                    "examples": [
                        "6",
                        "9000/778/B180L",
                    ],
                    "name": "host.cpu.model.id",
                    "requirement_level": "recommended",
                    "root_namespace": "host",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Model designation of the processor.\n",
                    "examples": [
                        "11th Gen Intel(R) Core(TM) i7-1185G7 @ 3.00GHz",
                    ],
                    "name": "host.cpu.model.name",
                    "requirement_level": "recommended",
                    "root_namespace": "host",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Stepping or core revisions.\n",
                    "examples": [
                        "1",
                        "r1p1",
                    ],
                    "name": "host.cpu.stepping",
                    "requirement_level": "recommended",
                    "root_namespace": "host",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Processor manufacturer identifier. A maximum 12-character string.\n",
                    "examples": [
                        "GenuineIntel",
                    ],
                    "name": "host.cpu.vendor.id",
                    "note": "[CPUID](https://wiki.osdev.org/CPUID) command returns the vendor ID string in EBX, EDX and ECX registers. Writing these to memory in this order results in a 12-character string.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "host",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Unique host ID. For Cloud, this must be the instance_id assigned by the cloud provider. For non-containerized systems, this should be the `machine-id`. See the table below for the sources to use to determine the `machine-id` based on operating system.\n",
                    "examples": [
                        "fdbf79e8af94cb7f9e8df36789187052",
                    ],
                    "name": "host.id",
                    "requirement_level": "recommended",
                    "root_namespace": "host",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "VM image ID or host OS image ID. For Cloud, this value is from the provider.\n",
                    "examples": [
                        "ami-07b06b442921831e5",
                    ],
                    "name": "host.image.id",
                    "requirement_level": "recommended",
                    "root_namespace": "host",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Name of the VM image or OS install the host was instantiated from.\n",
                    "examples": [
                        "infra-ami-eks-worker-node-7d4ec78312",
                        "CentOS-8-x86_64-1905",
                    ],
                    "name": "host.image.name",
                    "requirement_level": "recommended",
                    "root_namespace": "host",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The version string of the VM image or host OS as defined in [Version Attributes](/docs/resource/README.md#version-attributes).\n",
                    "examples": [
                        "0.1",
                    ],
                    "name": "host.image.version",
                    "requirement_level": "recommended",
                    "root_namespace": "host",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Available IP addresses of the host, excluding loopback interfaces.\n",
                    "examples": [
                        [
                            "192.168.1.140",
                            "fe80::abc2:4a28:737a:609e",
                        ],
                    ],
                    "name": "host.ip",
                    "note": "IPv4 Addresses MUST be specified in dotted-quad notation. IPv6 addresses MUST be specified in the [RFC 5952](https://www.rfc-editor.org/rfc/rfc5952.html) format.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "host",
                    "stability": "development",
                    "type": "string[]",
                },
                {
                    "brief": "Available MAC addresses of the host, excluding loopback interfaces.\n",
                    "examples": [
                        [
                            "AC-DE-48-23-45-67",
                            "AC-DE-48-23-45-67-01-9F",
                        ],
                    ],
                    "name": "host.mac",
                    "note": "MAC Addresses MUST be represented in [IEEE RA hexadecimal form](https://standards.ieee.org/wp-content/uploads/import/documents/tutorials/eui.pdf): as hyphen-separated octets in uppercase hexadecimal form from most to least significant.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "host",
                    "stability": "development",
                    "type": "string[]",
                },
                {
                    "brief": "Name of the host. On Unix systems, it may contain what the hostname command returns, or the fully qualified hostname, or another name specified by the user.\n",
                    "examples": [
                        "opentelemetry-test",
                    ],
                    "name": "host.name",
                    "requirement_level": "recommended",
                    "root_namespace": "host",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Type of host. For Cloud, this must be the machine type.\n",
                    "examples": [
                        "n1-standard-1",
                    ],
                    "name": "host.type",
                    "requirement_level": "recommended",
                    "root_namespace": "host",
                    "stability": "development",
                    "type": "string",
                },
            ],
            "root_namespace": "host",
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
