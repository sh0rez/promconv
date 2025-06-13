package os

// Unique identifier for a particular build or compilation of the operating system
type AttrBuildId string // os.build_id

func (AttrBuildId) Development()    {}
func (AttrBuildId) Recommended()    {}
func (AttrBuildId) Key() string     { return "os_build_id" }
func (a AttrBuildId) Value() string { return string(a) }

// Human readable (not intended to be parsed) OS version information, like e.g. reported by `ver` or `lsb_release -a` commands
type AttrDescription string // os.description

func (AttrDescription) Development()    {}
func (AttrDescription) Recommended()    {}
func (AttrDescription) Key() string     { return "os_description" }
func (a AttrDescription) Value() string { return string(a) }

// Human readable operating system name
type AttrName string // os.name

func (AttrName) Development()    {}
func (AttrName) Recommended()    {}
func (AttrName) Key() string     { return "os_name" }
func (a AttrName) Value() string { return string(a) }

// The operating system type
type AttrType string // os.type

func (AttrType) Development()    {}
func (AttrType) Recommended()    {}
func (AttrType) Key() string     { return "os_type" }
func (a AttrType) Value() string { return string(a) }

const TypeWindows AttrType = "windows"
const TypeLinux AttrType = "linux"
const TypeDarwin AttrType = "darwin"
const TypeFreebsd AttrType = "freebsd"
const TypeNetbsd AttrType = "netbsd"
const TypeOpenbsd AttrType = "openbsd"
const TypeDragonflybsd AttrType = "dragonflybsd"
const TypeHpux AttrType = "hpux"
const TypeAix AttrType = "aix"
const TypeSolaris AttrType = "solaris"
const TypeZOs AttrType = "z_os"
const TypeZos AttrType = "zos"

// The version string of the operating system as defined in [Version Attributes]
//
// [Version Attributes]: /docs/resource/README.md#version-attributes
type AttrVersion string // os.version

func (AttrVersion) Development()    {}
func (AttrVersion) Recommended()    {}
func (AttrVersion) Key() string     { return "os_version" }
func (a AttrVersion) Value() string { return string(a) }

/* State {
    name: "attr.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "ctx": {
            "attributes": [
                {
                    "brief": "Unique identifier for a particular build or compilation of the operating system.",
                    "examples": [
                        "TQ3C.230805.001.B2",
                        "20E247",
                        "22621",
                    ],
                    "name": "os.build_id",
                    "requirement_level": "recommended",
                    "root_namespace": "os",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Human readable (not intended to be parsed) OS version information, like e.g. reported by `ver` or `lsb_release -a` commands.\n",
                    "examples": [
                        "Microsoft Windows [Version 10.0.18363.778]",
                        "Ubuntu 18.04.1 LTS",
                    ],
                    "name": "os.description",
                    "requirement_level": "recommended",
                    "root_namespace": "os",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Human readable operating system name.",
                    "examples": [
                        "iOS",
                        "Android",
                        "Ubuntu",
                    ],
                    "name": "os.name",
                    "requirement_level": "recommended",
                    "root_namespace": "os",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The operating system type.\n",
                    "name": "os.type",
                    "requirement_level": "recommended",
                    "root_namespace": "os",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": "Microsoft Windows",
                                "deprecated": none,
                                "id": "windows",
                                "note": none,
                                "stability": "development",
                                "value": "windows",
                            },
                            {
                                "brief": "Linux",
                                "deprecated": none,
                                "id": "linux",
                                "note": none,
                                "stability": "development",
                                "value": "linux",
                            },
                            {
                                "brief": "Apple Darwin",
                                "deprecated": none,
                                "id": "darwin",
                                "note": none,
                                "stability": "development",
                                "value": "darwin",
                            },
                            {
                                "brief": "FreeBSD",
                                "deprecated": none,
                                "id": "freebsd",
                                "note": none,
                                "stability": "development",
                                "value": "freebsd",
                            },
                            {
                                "brief": "NetBSD",
                                "deprecated": none,
                                "id": "netbsd",
                                "note": none,
                                "stability": "development",
                                "value": "netbsd",
                            },
                            {
                                "brief": "OpenBSD",
                                "deprecated": none,
                                "id": "openbsd",
                                "note": none,
                                "stability": "development",
                                "value": "openbsd",
                            },
                            {
                                "brief": "DragonFly BSD",
                                "deprecated": none,
                                "id": "dragonflybsd",
                                "note": none,
                                "stability": "development",
                                "value": "dragonflybsd",
                            },
                            {
                                "brief": "HP-UX (Hewlett Packard Unix)",
                                "deprecated": none,
                                "id": "hpux",
                                "note": none,
                                "stability": "development",
                                "value": "hpux",
                            },
                            {
                                "brief": "AIX (Advanced Interactive eXecutive)",
                                "deprecated": none,
                                "id": "aix",
                                "note": none,
                                "stability": "development",
                                "value": "aix",
                            },
                            {
                                "brief": "SunOS, Oracle Solaris",
                                "deprecated": none,
                                "id": "solaris",
                                "note": none,
                                "stability": "development",
                                "value": "solaris",
                            },
                            {
                                "brief": "Deprecated. Use `zos` instead.",
                                "deprecated": none,
                                "id": "z_os",
                                "note": none,
                                "stability": "development",
                                "value": "z_os",
                            },
                            {
                                "brief": "IBM z/OS",
                                "deprecated": none,
                                "id": "zos",
                                "note": none,
                                "stability": "development",
                                "value": "zos",
                            },
                        ],
                    },
                },
                {
                    "brief": "The version string of the operating system as defined in [Version Attributes](/docs/resource/README.md#version-attributes).\n",
                    "examples": [
                        "14.2.1",
                        "18.04.1",
                    ],
                    "name": "os.version",
                    "requirement_level": "recommended",
                    "root_namespace": "os",
                    "stability": "development",
                    "type": "string",
                },
            ],
            "root_namespace": "os",
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
