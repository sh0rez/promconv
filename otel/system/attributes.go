package system

// Deprecated, use `cpu.logical_number` instead
type AttrCpuLogicalNumber string // system.cpu.logical_number

func (AttrCpuLogicalNumber) Development()    {}
func (AttrCpuLogicalNumber) Recommended()    {}
func (AttrCpuLogicalNumber) Key() string     { return "system_cpu_logical_number" }
func (a AttrCpuLogicalNumber) Value() string { return string(a) }

// Deprecated, use `cpu.mode` instead
type AttrCpuState string // system.cpu.state

func (AttrCpuState) Development()    {}
func (AttrCpuState) Recommended()    {}
func (AttrCpuState) Key() string     { return "system_cpu_state" }
func (a AttrCpuState) Value() string { return string(a) }

const CpuStateUser AttrCpuState = "user"
const CpuStateSystem AttrCpuState = "system"
const CpuStateNice AttrCpuState = "nice"
const CpuStateIdle AttrCpuState = "idle"
const CpuStateIowait AttrCpuState = "iowait"
const CpuStateInterrupt AttrCpuState = "interrupt"
const CpuStateSteal AttrCpuState = "steal"

// The device identifier
type AttrDevice string // system.device

func (AttrDevice) Development()    {}
func (AttrDevice) Recommended()    {}
func (AttrDevice) Key() string     { return "system_device" }
func (a AttrDevice) Value() string { return string(a) }

// The filesystem mode
type AttrFilesystemMode string // system.filesystem.mode

func (AttrFilesystemMode) Development()    {}
func (AttrFilesystemMode) Recommended()    {}
func (AttrFilesystemMode) Key() string     { return "system_filesystem_mode" }
func (a AttrFilesystemMode) Value() string { return string(a) }

// The filesystem mount path
type AttrFilesystemMountpoint string // system.filesystem.mountpoint

func (AttrFilesystemMountpoint) Development()    {}
func (AttrFilesystemMountpoint) Recommended()    {}
func (AttrFilesystemMountpoint) Key() string     { return "system_filesystem_mountpoint" }
func (a AttrFilesystemMountpoint) Value() string { return string(a) }

// The filesystem state
type AttrFilesystemState string // system.filesystem.state

func (AttrFilesystemState) Development()    {}
func (AttrFilesystemState) Recommended()    {}
func (AttrFilesystemState) Key() string     { return "system_filesystem_state" }
func (a AttrFilesystemState) Value() string { return string(a) }

const FilesystemStateUsed AttrFilesystemState = "used"
const FilesystemStateFree AttrFilesystemState = "free"
const FilesystemStateReserved AttrFilesystemState = "reserved"

// The filesystem type
type AttrFilesystemType string // system.filesystem.type

func (AttrFilesystemType) Development()    {}
func (AttrFilesystemType) Recommended()    {}
func (AttrFilesystemType) Key() string     { return "system_filesystem_type" }
func (a AttrFilesystemType) Value() string { return string(a) }

const FilesystemTypeFat32 AttrFilesystemType = "fat32"
const FilesystemTypeExfat AttrFilesystemType = "exfat"
const FilesystemTypeNtfs AttrFilesystemType = "ntfs"
const FilesystemTypeRefs AttrFilesystemType = "refs"
const FilesystemTypeHfsplus AttrFilesystemType = "hfsplus"
const FilesystemTypeExt4 AttrFilesystemType = "ext4"

// The memory state
type AttrMemoryState string // system.memory.state

func (AttrMemoryState) Development()    {}
func (AttrMemoryState) Recommended()    {}
func (AttrMemoryState) Key() string     { return "system_memory_state" }
func (a AttrMemoryState) Value() string { return string(a) }

const MemoryStateUsed AttrMemoryState = "used"
const MemoryStateFree AttrMemoryState = "free"
const MemoryStateShared AttrMemoryState = "shared"
const MemoryStateBuffers AttrMemoryState = "buffers"
const MemoryStateCached AttrMemoryState = "cached"

// Deprecated, use `network.connection.state` instead
type AttrNetworkState string // system.network.state

func (AttrNetworkState) Development()    {}
func (AttrNetworkState) Recommended()    {}
func (AttrNetworkState) Key() string     { return "system_network_state" }
func (a AttrNetworkState) Value() string { return string(a) }

const NetworkStateClose AttrNetworkState = "close"
const NetworkStateCloseWait AttrNetworkState = "close_wait"
const NetworkStateClosing AttrNetworkState = "closing"
const NetworkStateDelete AttrNetworkState = "delete"
const NetworkStateEstablished AttrNetworkState = "established"
const NetworkStateFinWait1 AttrNetworkState = "fin_wait_1"
const NetworkStateFinWait2 AttrNetworkState = "fin_wait_2"
const NetworkStateLastAck AttrNetworkState = "last_ack"
const NetworkStateListen AttrNetworkState = "listen"
const NetworkStateSynRecv AttrNetworkState = "syn_recv"
const NetworkStateSynSent AttrNetworkState = "syn_sent"
const NetworkStateTimeWait AttrNetworkState = "time_wait"

// The paging access direction
type AttrPagingDirection string // system.paging.direction

func (AttrPagingDirection) Development()    {}
func (AttrPagingDirection) Recommended()    {}
func (AttrPagingDirection) Key() string     { return "system_paging_direction" }
func (a AttrPagingDirection) Value() string { return string(a) }

const PagingDirectionIn AttrPagingDirection = "in"
const PagingDirectionOut AttrPagingDirection = "out"

// The memory paging state
type AttrPagingState string // system.paging.state

func (AttrPagingState) Development()    {}
func (AttrPagingState) Recommended()    {}
func (AttrPagingState) Key() string     { return "system_paging_state" }
func (a AttrPagingState) Value() string { return string(a) }

const PagingStateUsed AttrPagingState = "used"
const PagingStateFree AttrPagingState = "free"

// The memory paging type
type AttrPagingType string // system.paging.type

func (AttrPagingType) Development()    {}
func (AttrPagingType) Recommended()    {}
func (AttrPagingType) Key() string     { return "system_paging_type" }
func (a AttrPagingType) Value() string { return string(a) }

const PagingTypeMajor AttrPagingType = "major"
const PagingTypeMinor AttrPagingType = "minor"

// The process state, e.g., [Linux Process State Codes]
//
// [Linux Process State Codes]: https://man7.org/linux/man-pages/man1/ps.1.html#PROCESS_STATE_CODES
type AttrProcessStatus string // system.process.status

func (AttrProcessStatus) Development()    {}
func (AttrProcessStatus) Recommended()    {}
func (AttrProcessStatus) Key() string     { return "system_process_status" }
func (a AttrProcessStatus) Value() string { return string(a) }

const ProcessStatusRunning AttrProcessStatus = "running"
const ProcessStatusSleeping AttrProcessStatus = "sleeping"
const ProcessStatusStopped AttrProcessStatus = "stopped"
const ProcessStatusDefunct AttrProcessStatus = "defunct"

// Deprecated, use `system.process.status` instead
type AttrProcessesStatus string // system.processes.status

func (AttrProcessesStatus) Development()    {}
func (AttrProcessesStatus) Recommended()    {}
func (AttrProcessesStatus) Key() string     { return "system_processes_status" }
func (a AttrProcessesStatus) Value() string { return string(a) }

const ProcessesStatusRunning AttrProcessesStatus = "running"
const ProcessesStatusSleeping AttrProcessesStatus = "sleeping"
const ProcessesStatusStopped AttrProcessesStatus = "stopped"
const ProcessesStatusDefunct AttrProcessesStatus = "defunct"

/* State {
    name: "attr.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "ctx": {
            "attributes": [
                {
                    "brief": "Deprecated, use `cpu.logical_number` instead.",
                    "examples": [
                        1,
                    ],
                    "name": "system.cpu.logical_number",
                    "requirement_level": "recommended",
                    "root_namespace": "system",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "Deprecated, use `cpu.mode` instead.",
                    "deprecated": {
                        "note": "Replaced by `cpu.mode`.",
                        "reason": "renamed",
                        "renamed_to": "cpu.mode",
                    },
                    "examples": [
                        "idle",
                        "interrupt",
                    ],
                    "name": "system.cpu.state",
                    "requirement_level": "recommended",
                    "root_namespace": "system",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "user",
                                "note": none,
                                "stability": "development",
                                "value": "user",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "system",
                                "note": none,
                                "stability": "development",
                                "value": "system",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "nice",
                                "note": none,
                                "stability": "development",
                                "value": "nice",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "idle",
                                "note": none,
                                "stability": "development",
                                "value": "idle",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "iowait",
                                "note": none,
                                "stability": "development",
                                "value": "iowait",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "interrupt",
                                "note": none,
                                "stability": "development",
                                "value": "interrupt",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "steal",
                                "note": none,
                                "stability": "development",
                                "value": "steal",
                            },
                        ],
                    },
                },
                {
                    "brief": "The device identifier",
                    "examples": [
                        "(identifier)",
                    ],
                    "name": "system.device",
                    "requirement_level": "recommended",
                    "root_namespace": "system",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The filesystem mode",
                    "examples": [
                        "rw, ro",
                    ],
                    "name": "system.filesystem.mode",
                    "requirement_level": "recommended",
                    "root_namespace": "system",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The filesystem mount path",
                    "examples": [
                        "/mnt/data",
                    ],
                    "name": "system.filesystem.mountpoint",
                    "requirement_level": "recommended",
                    "root_namespace": "system",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The filesystem state",
                    "examples": [
                        "used",
                    ],
                    "name": "system.filesystem.state",
                    "requirement_level": "recommended",
                    "root_namespace": "system",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "used",
                                "note": none,
                                "stability": "development",
                                "value": "used",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "free",
                                "note": none,
                                "stability": "development",
                                "value": "free",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "reserved",
                                "note": none,
                                "stability": "development",
                                "value": "reserved",
                            },
                        ],
                    },
                },
                {
                    "brief": "The filesystem type",
                    "examples": [
                        "ext4",
                    ],
                    "name": "system.filesystem.type",
                    "requirement_level": "recommended",
                    "root_namespace": "system",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "fat32",
                                "note": none,
                                "stability": "development",
                                "value": "fat32",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "exfat",
                                "note": none,
                                "stability": "development",
                                "value": "exfat",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "ntfs",
                                "note": none,
                                "stability": "development",
                                "value": "ntfs",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "refs",
                                "note": none,
                                "stability": "development",
                                "value": "refs",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "hfsplus",
                                "note": none,
                                "stability": "development",
                                "value": "hfsplus",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "ext4",
                                "note": none,
                                "stability": "development",
                                "value": "ext4",
                            },
                        ],
                    },
                },
                {
                    "brief": "The memory state",
                    "examples": [
                        "free",
                        "cached",
                    ],
                    "name": "system.memory.state",
                    "requirement_level": "recommended",
                    "root_namespace": "system",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "used",
                                "note": none,
                                "stability": "development",
                                "value": "used",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "free",
                                "note": none,
                                "stability": "development",
                                "value": "free",
                            },
                            {
                                "brief": none,
                                "deprecated": "Removed, report shared memory usage with `metric.system.memory.shared` metric",
                                "id": "shared",
                                "note": none,
                                "stability": "development",
                                "value": "shared",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "buffers",
                                "note": none,
                                "stability": "development",
                                "value": "buffers",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "cached",
                                "note": none,
                                "stability": "development",
                                "value": "cached",
                            },
                        ],
                    },
                },
                {
                    "brief": "Deprecated, use `network.connection.state` instead.",
                    "deprecated": {
                        "note": "Replaced by `network.connection.state`.",
                        "reason": "renamed",
                        "renamed_to": "network.connection.state",
                    },
                    "examples": [
                        "close_wait",
                    ],
                    "name": "system.network.state",
                    "requirement_level": "recommended",
                    "root_namespace": "system",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "close",
                                "note": none,
                                "stability": "development",
                                "value": "close",
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
                                "id": "delete",
                                "note": none,
                                "stability": "development",
                                "value": "delete",
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
                                "id": "syn_recv",
                                "note": none,
                                "stability": "development",
                                "value": "syn_recv",
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
                    "brief": "The paging access direction",
                    "examples": [
                        "in",
                    ],
                    "name": "system.paging.direction",
                    "requirement_level": "recommended",
                    "root_namespace": "system",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "in",
                                "note": none,
                                "stability": "development",
                                "value": "in",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "out",
                                "note": none,
                                "stability": "development",
                                "value": "out",
                            },
                        ],
                    },
                },
                {
                    "brief": "The memory paging state",
                    "examples": [
                        "free",
                    ],
                    "name": "system.paging.state",
                    "requirement_level": "recommended",
                    "root_namespace": "system",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "used",
                                "note": none,
                                "stability": "development",
                                "value": "used",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "free",
                                "note": none,
                                "stability": "development",
                                "value": "free",
                            },
                        ],
                    },
                },
                {
                    "brief": "The memory paging type",
                    "examples": [
                        "minor",
                    ],
                    "name": "system.paging.type",
                    "requirement_level": "recommended",
                    "root_namespace": "system",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "major",
                                "note": none,
                                "stability": "development",
                                "value": "major",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "minor",
                                "note": none,
                                "stability": "development",
                                "value": "minor",
                            },
                        ],
                    },
                },
                {
                    "brief": "The process state, e.g., [Linux Process State Codes](https://man7.org/linux/man-pages/man1/ps.1.html#PROCESS_STATE_CODES)\n",
                    "examples": [
                        "running",
                    ],
                    "name": "system.process.status",
                    "requirement_level": "recommended",
                    "root_namespace": "system",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "running",
                                "note": none,
                                "stability": "development",
                                "value": "running",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "sleeping",
                                "note": none,
                                "stability": "development",
                                "value": "sleeping",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "stopped",
                                "note": none,
                                "stability": "development",
                                "value": "stopped",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "defunct",
                                "note": none,
                                "stability": "development",
                                "value": "defunct",
                            },
                        ],
                    },
                },
                {
                    "brief": "Deprecated, use `system.process.status` instead.",
                    "deprecated": {
                        "note": "Replaced by `system.process.status`.",
                        "reason": "renamed",
                        "renamed_to": "system.process.status",
                    },
                    "examples": [
                        "running",
                    ],
                    "name": "system.processes.status",
                    "requirement_level": "recommended",
                    "root_namespace": "system",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "running",
                                "note": none,
                                "stability": "development",
                                "value": "running",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "sleeping",
                                "note": none,
                                "stability": "development",
                                "value": "sleeping",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "stopped",
                                "note": none,
                                "stability": "development",
                                "value": "stopped",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "defunct",
                                "note": none,
                                "stability": "development",
                                "value": "defunct",
                            },
                        ],
                    },
                },
            ],
            "root_namespace": "system",
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
