package system

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Reports a filesystem's space usage across different states.
type FilesystemUsage struct {
	*prometheus.GaugeVec
	extra FilesystemUsageExtra
}

func NewFilesystemUsage() FilesystemUsage {
	labels := []string{AttrDevice("").Key(), AttrFilesystemMode("").Key(), AttrFilesystemMountpoint("").Key(), AttrFilesystemState("").Key(), AttrFilesystemType("").Key()}
	return FilesystemUsage{GaugeVec: prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "system_filesystem_usage",
		Help: "Reports a filesystem's space usage across different states.",
	}, labels)}
}

func (m FilesystemUsage) With(extras ...interface {
	SystemDevice() AttrDevice
	SystemFilesystemMode() AttrFilesystemMode
	SystemFilesystemMountpoint() AttrFilesystemMountpoint
	SystemFilesystemState() AttrFilesystemState
	SystemFilesystemType() AttrFilesystemType
}) prometheus.Gauge {
	if extras == nil {
		extras = append(extras, m.extra)
	}
	extra := extras[0]

	return m.GaugeVec.WithLabelValues(extra.SystemDevice().Value(), extra.SystemFilesystemMode().Value(), extra.SystemFilesystemMountpoint().Value(), extra.SystemFilesystemState().Value(), extra.SystemFilesystemType().Value())
}

// Deprecated: Use [FilesystemUsage.With] instead
func (m FilesystemUsage) WithLabelValues(lvs ...string) prometheus.Gauge {
	return m.GaugeVec.WithLabelValues(lvs...)
}

func (a FilesystemUsage) WithDevice(attr interface{ SystemDevice() AttrDevice }) FilesystemUsage {
	a.extra.AttrDevice = attr.SystemDevice()
	return a
}
func (a FilesystemUsage) WithFilesystemMode(attr interface{ SystemFilesystemMode() AttrFilesystemMode }) FilesystemUsage {
	a.extra.AttrFilesystemMode = attr.SystemFilesystemMode()
	return a
}
func (a FilesystemUsage) WithFilesystemMountpoint(attr interface {
	SystemFilesystemMountpoint() AttrFilesystemMountpoint
}) FilesystemUsage {
	a.extra.AttrFilesystemMountpoint = attr.SystemFilesystemMountpoint()
	return a
}
func (a FilesystemUsage) WithFilesystemState(attr interface{ SystemFilesystemState() AttrFilesystemState }) FilesystemUsage {
	a.extra.AttrFilesystemState = attr.SystemFilesystemState()
	return a
}
func (a FilesystemUsage) WithFilesystemType(attr interface{ SystemFilesystemType() AttrFilesystemType }) FilesystemUsage {
	a.extra.AttrFilesystemType = attr.SystemFilesystemType()
	return a
}

type FilesystemUsageExtra struct {
	// Identifier for the device where the filesystem resides
	AttrDevice               AttrDevice               `otel:"system.device"`                // The filesystem mode
	AttrFilesystemMode       AttrFilesystemMode       `otel:"system.filesystem.mode"`       // The filesystem mount path
	AttrFilesystemMountpoint AttrFilesystemMountpoint `otel:"system.filesystem.mountpoint"` // The filesystem state
	AttrFilesystemState      AttrFilesystemState      `otel:"system.filesystem.state"`      // The filesystem type
	AttrFilesystemType       AttrFilesystemType       `otel:"system.filesystem.type"`
}

func (a FilesystemUsageExtra) SystemDevice() AttrDevice                 { return a.AttrDevice }
func (a FilesystemUsageExtra) SystemFilesystemMode() AttrFilesystemMode { return a.AttrFilesystemMode }
func (a FilesystemUsageExtra) SystemFilesystemMountpoint() AttrFilesystemMountpoint {
	return a.AttrFilesystemMountpoint
}
func (a FilesystemUsageExtra) SystemFilesystemState() AttrFilesystemState {
	return a.AttrFilesystemState
}
func (a FilesystemUsageExtra) SystemFilesystemType() AttrFilesystemType { return a.AttrFilesystemType }

/*
State {
    name: "vec.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "FilesystemUsageExtra",
        "Instr": "Gauge",
        "InstrMap": {
            "counter": "Counter",
            "gauge": "Gauge",
            "histogram": "Histogram",
            "updowncounter": "Gauge",
        },
        "Name": "filesystem.usage",
        "Type": "FilesystemUsage",
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
            "brief": "Reports a filesystem's space usage across different states.",
            "entity_associations": [
                "host",
            ],
            "events": [],
            "id": "metric.system.filesystem.usage",
            "instrument": "updowncounter",
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
            "metric_name": "system.filesystem.usage",
            "name": none,
            "note": "The sum of all `system.filesystem.usage` values over the different `system.filesystem.state` attributes\nSHOULD equal the total storage capacity of the filesystem, that is `system.filesystem.limit`.\n",
            "root_namespace": "system",
            "span_kind": none,
            "stability": "development",
            "type": "metric",
            "unit": "By",
        },
        "for_each_attr": <macro for_each_attr>,
        "module": "shorez.de/promconv/otel",
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
            "vec.go.j2",
        ],
    },
}
*/
