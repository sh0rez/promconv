package system

import (
	"github.com/prometheus/client_golang/prometheus"
)

type FilesystemUtilization struct {
	*prometheus.GaugeVec
	extra FilesystemUtilizationExtra
}

func NewFilesystemUtilization() FilesystemUtilization {
	labels := []string{"system_device", "system_filesystem_mode", "system_filesystem_mountpoint", "system_filesystem_state", "system_filesystem_type"}
	return FilesystemUtilization{GaugeVec: prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "system",
		Name:      "filesystem_utilization",
		Help:      "",
	}, labels)}
}

func (m FilesystemUtilization) With(extra interface {
	AttrSystemDevice() AttrDevice
	AttrSystemFilesystemMode() AttrFilesystemMode
	AttrSystemFilesystemMountpoint() AttrFilesystemMountpoint
	AttrSystemFilesystemState() AttrFilesystemState
	AttrSystemFilesystemType() AttrFilesystemType
}) prometheus.Gauge {
	if extra == nil {
		extra = m.extra
	}
	return m.WithLabelValues(
		string(extra.AttrSystemDevice()),
		string(extra.AttrSystemFilesystemMode()),
		string(extra.AttrSystemFilesystemMountpoint()),
		string(extra.AttrSystemFilesystemState()),
		string(extra.AttrSystemFilesystemType()),
	)
}

func (a FilesystemUtilization) WithSystemDevice(attr interface{ AttrSystemDevice() AttrDevice }) FilesystemUtilization {
	a.extra.SystemDevice = attr.AttrSystemDevice()
	return a
}
func (a FilesystemUtilization) WithSystemFilesystemMode(attr interface{ AttrSystemFilesystemMode() AttrFilesystemMode }) FilesystemUtilization {
	a.extra.SystemFilesystemMode = attr.AttrSystemFilesystemMode()
	return a
}
func (a FilesystemUtilization) WithSystemFilesystemMountpoint(attr interface {
	AttrSystemFilesystemMountpoint() AttrFilesystemMountpoint
}) FilesystemUtilization {
	a.extra.SystemFilesystemMountpoint = attr.AttrSystemFilesystemMountpoint()
	return a
}
func (a FilesystemUtilization) WithSystemFilesystemState(attr interface{ AttrSystemFilesystemState() AttrFilesystemState }) FilesystemUtilization {
	a.extra.SystemFilesystemState = attr.AttrSystemFilesystemState()
	return a
}
func (a FilesystemUtilization) WithSystemFilesystemType(attr interface{ AttrSystemFilesystemType() AttrFilesystemType }) FilesystemUtilization {
	a.extra.SystemFilesystemType = attr.AttrSystemFilesystemType()
	return a
}

type FilesystemUtilizationExtra struct {
	// Identifier for the device where the filesystem resides.
	SystemDevice AttrDevice `otel:"system.device"`
	// The filesystem mode
	SystemFilesystemMode AttrFilesystemMode `otel:"system.filesystem.mode"`
	// The filesystem mount path
	SystemFilesystemMountpoint AttrFilesystemMountpoint `otel:"system.filesystem.mountpoint"`
	// The filesystem state
	SystemFilesystemState AttrFilesystemState `otel:"system.filesystem.state"`
	// The filesystem type
	SystemFilesystemType AttrFilesystemType `otel:"system.filesystem.type"`
}

func (a FilesystemUtilizationExtra) AttrSystemDevice() AttrDevice { return a.SystemDevice }
func (a FilesystemUtilizationExtra) AttrSystemFilesystemMode() AttrFilesystemMode {
	return a.SystemFilesystemMode
}
func (a FilesystemUtilizationExtra) AttrSystemFilesystemMountpoint() AttrFilesystemMountpoint {
	return a.SystemFilesystemMountpoint
}
func (a FilesystemUtilizationExtra) AttrSystemFilesystemState() AttrFilesystemState {
	return a.SystemFilesystemState
}
func (a FilesystemUtilizationExtra) AttrSystemFilesystemType() AttrFilesystemType {
	return a.SystemFilesystemType
}

/*
State {
    name: "metric.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "FilesystemUtilizationExtra",
        "Instr": "Gauge",
        "InstrMap": {
            "counter": "Counter",
            "gauge": "Gauge",
            "histogram": "Histogram",
            "updowncounter": "Gauge",
        },
        "Name": "filesystem.utilization",
        "Type": "FilesystemUtilization",
        "attributes": [
            {
                "brief": "Identifier for the device where the filesystem resides.",
                "examples": [
                    "/dev/sda",
                    "\\network-drive",
                ],
                "name": "system.device",
                "requirement_level": "recommended",
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
                "stability": "development",
                "type": {
                    "allow_custom_values": none,
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
                "stability": "development",
                "type": {
                    "allow_custom_values": none,
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
        ],
        "ctx": {
            "attributes": [
                {
                    "brief": "The filesystem state",
                    "examples": [
                        "used",
                    ],
                    "name": "system.filesystem.state",
                    "requirement_level": "recommended",
                    "stability": "development",
                    "type": {
                        "allow_custom_values": none,
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
                    "stability": "development",
                    "type": {
                        "allow_custom_values": none,
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
                    "brief": "The filesystem mode",
                    "examples": [
                        "rw, ro",
                    ],
                    "name": "system.filesystem.mode",
                    "requirement_level": "recommended",
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
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Identifier for the device where the filesystem resides.",
                    "examples": [
                        "/dev/sda",
                        "\\network-drive",
                    ],
                    "name": "system.device",
                    "requirement_level": "recommended",
                    "stability": "development",
                    "type": "string",
                },
            ],
            "entity_associations": [
                "host",
            ],
            "events": [],
            "id": "metric.system.filesystem.utilization",
            "instrument": "gauge",
            "lineage": {
                "attributes": {
                    "system.device": {
                        "inherited_fields": [
                            "note",
                            "requirement_level",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "brief",
                            "examples",
                        ],
                        "source_group": "registry.system",
                    },
                    "system.filesystem.mode": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "requirement_level",
                            "stability",
                        ],
                        "source_group": "registry.system.filesystem",
                    },
                    "system.filesystem.mountpoint": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "requirement_level",
                            "stability",
                        ],
                        "source_group": "registry.system.filesystem",
                    },
                    "system.filesystem.state": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "requirement_level",
                            "stability",
                        ],
                        "source_group": "registry.system.filesystem",
                    },
                    "system.filesystem.type": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "requirement_level",
                            "stability",
                        ],
                        "source_group": "registry.system.filesystem",
                    },
                },
                "provenance": {
                    "path": "https://github.com/open-telemetry/semantic-conventions.git[model]/system/metrics.yaml",
                    "registry_id": "main",
                },
            },
            "metric_name": "system.filesystem.utilization",
            "name": none,
            "root_namespace": "system",
            "span_kind": none,
            "stability": "development",
            "type": "metric",
            "unit": "1",
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
            "metric.go.j2",
        ],
    },
}
*/
