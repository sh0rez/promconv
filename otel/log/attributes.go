package log

// The basename of the file
type AttrFileName string // log.file.name

func (AttrFileName) Development()    {}
func (AttrFileName) Recommended()    {}
func (AttrFileName) Key() string     { return "log_file_name" }
func (a AttrFileName) Value() string { return string(a) }

// The basename of the file, with symlinks resolved
type AttrFileNameResolved string // log.file.name_resolved

func (AttrFileNameResolved) Development()    {}
func (AttrFileNameResolved) Recommended()    {}
func (AttrFileNameResolved) Key() string     { return "log_file_name_resolved" }
func (a AttrFileNameResolved) Value() string { return string(a) }

// The full path to the file
type AttrFilePath string // log.file.path

func (AttrFilePath) Development()    {}
func (AttrFilePath) Recommended()    {}
func (AttrFilePath) Key() string     { return "log_file_path" }
func (a AttrFilePath) Value() string { return string(a) }

// The full path to the file, with symlinks resolved
type AttrFilePathResolved string // log.file.path_resolved

func (AttrFilePathResolved) Development()    {}
func (AttrFilePathResolved) Recommended()    {}
func (AttrFilePathResolved) Key() string     { return "log_file_path_resolved" }
func (a AttrFilePathResolved) Value() string { return string(a) }

// The stream associated with the log. See below for a list of well-known values
type AttrIostream string // log.iostream

func (AttrIostream) Development()    {}
func (AttrIostream) Recommended()    {}
func (AttrIostream) Key() string     { return "log_iostream" }
func (a AttrIostream) Value() string { return string(a) }

const IostreamStdout AttrIostream = "stdout"
const IostreamStderr AttrIostream = "stderr"

// The complete original Log Record.
//
// This value MAY be added when processing a Log Record which was originally transmitted as a string or equivalent data type AND the Body field of the Log Record does not contain the same value. (e.g. a syslog or a log record read from a file.)
type AttrRecordOriginal string // log.record.original

func (AttrRecordOriginal) Development()    {}
func (AttrRecordOriginal) Recommended()    {}
func (AttrRecordOriginal) Key() string     { return "log_record_original" }
func (a AttrRecordOriginal) Value() string { return string(a) }

// A unique identifier for the Log Record.
//
// If an id is provided, other log records with the same id will be considered duplicates and can be removed safely. This means, that two distinguishable log records MUST have different values.
// The id MAY be an [Universally Unique Lexicographically Sortable Identifier (ULID)], but other identifiers (e.g. UUID) may be used as needed
//
// [Universally Unique Lexicographically Sortable Identifier (ULID)]: https://github.com/ulid/spec
type AttrRecordUid string // log.record.uid

func (AttrRecordUid) Development()    {}
func (AttrRecordUid) Recommended()    {}
func (AttrRecordUid) Key() string     { return "log_record_uid" }
func (a AttrRecordUid) Value() string { return string(a) }

/* State {
    name: "attr.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "ctx": {
            "attributes": [
                {
                    "brief": "The basename of the file.\n",
                    "examples": [
                        "audit.log",
                    ],
                    "name": "log.file.name",
                    "requirement_level": "recommended",
                    "root_namespace": "log",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The basename of the file, with symlinks resolved.\n",
                    "examples": [
                        "uuid.log",
                    ],
                    "name": "log.file.name_resolved",
                    "requirement_level": "recommended",
                    "root_namespace": "log",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The full path to the file.\n",
                    "examples": [
                        "/var/log/mysql/audit.log",
                    ],
                    "name": "log.file.path",
                    "requirement_level": "recommended",
                    "root_namespace": "log",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The full path to the file, with symlinks resolved.\n",
                    "examples": [
                        "/var/lib/docker/uuid.log",
                    ],
                    "name": "log.file.path_resolved",
                    "requirement_level": "recommended",
                    "root_namespace": "log",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The stream associated with the log. See below for a list of well-known values.\n",
                    "name": "log.iostream",
                    "requirement_level": "recommended",
                    "root_namespace": "log",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": "Logs from stdout stream",
                                "deprecated": none,
                                "id": "stdout",
                                "note": none,
                                "stability": "development",
                                "value": "stdout",
                            },
                            {
                                "brief": "Events from stderr stream",
                                "deprecated": none,
                                "id": "stderr",
                                "note": none,
                                "stability": "development",
                                "value": "stderr",
                            },
                        ],
                    },
                },
                {
                    "brief": "The complete original Log Record.\n",
                    "examples": [
                        "77 <86>1 2015-08-06T21:58:59.694Z 192.168.2.133 inactive - - - Something happened",
                        "[INFO] 8/3/24 12:34:56 Something happened",
                    ],
                    "name": "log.record.original",
                    "note": "This value MAY be added when processing a Log Record which was originally transmitted as a string or equivalent data type AND the Body field of the Log Record does not contain the same value. (e.g. a syslog or a log record read from a file.)\n",
                    "requirement_level": "recommended",
                    "root_namespace": "log",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "A unique identifier for the Log Record.\n",
                    "examples": [
                        "01ARZ3NDEKTSV4RRFFQ69G5FAV",
                    ],
                    "name": "log.record.uid",
                    "note": "If an id is provided, other log records with the same id will be considered duplicates and can be removed safely. This means, that two distinguishable log records MUST have different values.\nThe id MAY be an [Universally Unique Lexicographically Sortable Identifier (ULID)](https://github.com/ulid/spec), but other identifiers (e.g. UUID) may be used as needed.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "log",
                    "stability": "development",
                    "type": "string",
                },
            ],
            "root_namespace": "log",
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
