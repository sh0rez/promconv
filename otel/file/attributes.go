package file

// Time when the file was last accessed, in ISO 8601 format.
//
// This attribute might not be supported by some file systems — NFS, FAT32, in embedded OS, etc
type AttrAccessed string // file.accessed

func (AttrAccessed) Development()    {}
func (AttrAccessed) Recommended()    {}
func (AttrAccessed) Key() string     { return "file_accessed" }
func (a AttrAccessed) Value() string { return string(a) }

// Array of file attributes.
//
// Attributes names depend on the OS or file system. Here’s a non-exhaustive list of values expected for this attribute: `archive`, `compressed`, `directory`, `encrypted`, `execute`, `hidden`, `immutable`, `journaled`, `read`, `readonly`, `symbolic link`, `system`, `temporary`, `write`
type AttrAttributes string // file.attributes

func (AttrAttributes) Development()    {}
func (AttrAttributes) Recommended()    {}
func (AttrAttributes) Key() string     { return "file_attributes" }
func (a AttrAttributes) Value() string { return string(a) }

// Time when the file attributes or metadata was last changed, in ISO 8601 format.
//
// `file.changed` captures the time when any of the file's properties or attributes (including the content) are changed, while `file.modified` captures the timestamp when the file content is modified
type AttrChanged string // file.changed

func (AttrChanged) Development()    {}
func (AttrChanged) Recommended()    {}
func (AttrChanged) Key() string     { return "file_changed" }
func (a AttrChanged) Value() string { return string(a) }

// Time when the file was created, in ISO 8601 format.
//
// This attribute might not be supported by some file systems — NFS, FAT32, in embedded OS, etc
type AttrCreated string // file.created

func (AttrCreated) Development()    {}
func (AttrCreated) Recommended()    {}
func (AttrCreated) Key() string     { return "file_created" }
func (a AttrCreated) Value() string { return string(a) }

// Directory where the file is located. It should include the drive letter, when appropriate
type AttrDirectory string // file.directory

func (AttrDirectory) Development()    {}
func (AttrDirectory) Recommended()    {}
func (AttrDirectory) Key() string     { return "file_directory" }
func (a AttrDirectory) Value() string { return string(a) }

// File extension, excluding the leading dot.
//
// When the file name has multiple extensions (example.tar.gz), only the last one should be captured ("gz", not "tar.gz")
type AttrExtension string // file.extension

func (AttrExtension) Development()    {}
func (AttrExtension) Recommended()    {}
func (AttrExtension) Key() string     { return "file_extension" }
func (a AttrExtension) Value() string { return string(a) }

// Name of the fork. A fork is additional data associated with a filesystem object.
//
// On Linux, a resource fork is used to store additional data with a filesystem object. A file always has at least one fork for the data portion, and additional forks may exist.
// On NTFS, this is analogous to an Alternate Data Stream (ADS), and the default data stream for a file is just called $DATA. Zone.Identifier is commonly used by Windows to track contents downloaded from the Internet. An ADS is typically of the form: C:\path\to\filename.extension:some_fork_name, and some_fork_name is the value that should populate `fork_name`. `filename.extension` should populate `file.name`, and `extension` should populate `file.extension`. The full path, `file.path`, will include the fork name
type AttrForkName string // file.fork_name

func (AttrForkName) Development()    {}
func (AttrForkName) Recommended()    {}
func (AttrForkName) Key() string     { return "file_fork_name" }
func (a AttrForkName) Value() string { return string(a) }

// Primary Group ID (GID) of the file
type AttrGroupId string // file.group.id

func (AttrGroupId) Development()    {}
func (AttrGroupId) Recommended()    {}
func (AttrGroupId) Key() string     { return "file_group_id" }
func (a AttrGroupId) Value() string { return string(a) }

// Primary group name of the file
type AttrGroupName string // file.group.name

func (AttrGroupName) Development()    {}
func (AttrGroupName) Recommended()    {}
func (AttrGroupName) Key() string     { return "file_group_name" }
func (a AttrGroupName) Value() string { return string(a) }

// Inode representing the file in the filesystem
type AttrInode string // file.inode

func (AttrInode) Development()    {}
func (AttrInode) Recommended()    {}
func (AttrInode) Key() string     { return "file_inode" }
func (a AttrInode) Value() string { return string(a) }

// Mode of the file in octal representation
type AttrMode string // file.mode

func (AttrMode) Development()    {}
func (AttrMode) Recommended()    {}
func (AttrMode) Key() string     { return "file_mode" }
func (a AttrMode) Value() string { return string(a) }

// Time when the file content was last modified, in ISO 8601 format
type AttrModified string // file.modified

func (AttrModified) Development()    {}
func (AttrModified) Recommended()    {}
func (AttrModified) Key() string     { return "file_modified" }
func (a AttrModified) Value() string { return string(a) }

// Name of the file including the extension, without the directory
type AttrName string // file.name

func (AttrName) Development()    {}
func (AttrName) Recommended()    {}
func (AttrName) Key() string     { return "file_name" }
func (a AttrName) Value() string { return string(a) }

// The user ID (UID) or security identifier (SID) of the file owner
type AttrOwnerId string // file.owner.id

func (AttrOwnerId) Development()    {}
func (AttrOwnerId) Recommended()    {}
func (AttrOwnerId) Key() string     { return "file_owner_id" }
func (a AttrOwnerId) Value() string { return string(a) }

// Username of the file owner
type AttrOwnerName string // file.owner.name

func (AttrOwnerName) Development()    {}
func (AttrOwnerName) Recommended()    {}
func (AttrOwnerName) Key() string     { return "file_owner_name" }
func (a AttrOwnerName) Value() string { return string(a) }

// Full path to the file, including the file name. It should include the drive letter, when appropriate
type AttrPath string // file.path

func (AttrPath) Development()    {}
func (AttrPath) Recommended()    {}
func (AttrPath) Key() string     { return "file_path" }
func (a AttrPath) Value() string { return string(a) }

// File size in bytes
type AttrSize string // file.size

func (AttrSize) Development()    {}
func (AttrSize) Recommended()    {}
func (AttrSize) Key() string     { return "file_size" }
func (a AttrSize) Value() string { return string(a) }

// Path to the target of a symbolic link.
//
// This attribute is only applicable to symbolic links
type AttrSymbolicLinkTargetPath string // file.symbolic_link.target_path

func (AttrSymbolicLinkTargetPath) Development()    {}
func (AttrSymbolicLinkTargetPath) Recommended()    {}
func (AttrSymbolicLinkTargetPath) Key() string     { return "file_symbolic_link_target_path" }
func (a AttrSymbolicLinkTargetPath) Value() string { return string(a) }

/* State {
    name: "attr.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "ctx": {
            "attributes": [
                {
                    "brief": "Time when the file was last accessed, in ISO 8601 format.\n",
                    "examples": [
                        "2021-01-01T12:00:00Z",
                    ],
                    "name": "file.accessed",
                    "note": "This attribute might not be supported by some file systems — NFS, FAT32, in embedded OS, etc.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "file",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Array of file attributes.\n",
                    "examples": [
                        [
                            "readonly",
                            "hidden",
                        ],
                    ],
                    "name": "file.attributes",
                    "note": "Attributes names depend on the OS or file system. Here’s a non-exhaustive list of values expected for this attribute: `archive`, `compressed`, `directory`, `encrypted`, `execute`, `hidden`, `immutable`, `journaled`, `read`, `readonly`, `symbolic link`, `system`, `temporary`, `write`.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "file",
                    "stability": "development",
                    "type": "string[]",
                },
                {
                    "brief": "Time when the file attributes or metadata was last changed, in ISO 8601 format.\n",
                    "examples": [
                        "2021-01-01T12:00:00Z",
                    ],
                    "name": "file.changed",
                    "note": "`file.changed` captures the time when any of the file's properties or attributes (including the content) are changed, while `file.modified` captures the timestamp when the file content is modified.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "file",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Time when the file was created, in ISO 8601 format.\n",
                    "examples": [
                        "2021-01-01T12:00:00Z",
                    ],
                    "name": "file.created",
                    "note": "This attribute might not be supported by some file systems — NFS, FAT32, in embedded OS, etc.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "file",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Directory where the file is located. It should include the drive letter, when appropriate.\n",
                    "examples": [
                        "/home/user",
                        "C:\\Program Files\\MyApp",
                    ],
                    "name": "file.directory",
                    "requirement_level": "recommended",
                    "root_namespace": "file",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "File extension, excluding the leading dot.\n",
                    "examples": [
                        "png",
                        "gz",
                    ],
                    "name": "file.extension",
                    "note": "When the file name has multiple extensions (example.tar.gz), only the last one should be captured (\"gz\", not \"tar.gz\").\n",
                    "requirement_level": "recommended",
                    "root_namespace": "file",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Name of the fork. A fork is additional data associated with a filesystem object.\n",
                    "examples": [
                        "Zone.Identifer",
                    ],
                    "name": "file.fork_name",
                    "note": "On Linux, a resource fork is used to store additional data with a filesystem object. A file always has at least one fork for the data portion, and additional forks may exist.\nOn NTFS, this is analogous to an Alternate Data Stream (ADS), and the default data stream for a file is just called $DATA. Zone.Identifier is commonly used by Windows to track contents downloaded from the Internet. An ADS is typically of the form: C:\\path\\to\\filename.extension:some_fork_name, and some_fork_name is the value that should populate `fork_name`. `filename.extension` should populate `file.name`, and `extension` should populate `file.extension`. The full path, `file.path`, will include the fork name.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "file",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Primary Group ID (GID) of the file.\n",
                    "examples": [
                        "1000",
                    ],
                    "name": "file.group.id",
                    "requirement_level": "recommended",
                    "root_namespace": "file",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Primary group name of the file.\n",
                    "examples": [
                        "users",
                    ],
                    "name": "file.group.name",
                    "requirement_level": "recommended",
                    "root_namespace": "file",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Inode representing the file in the filesystem.\n",
                    "examples": [
                        "256383",
                    ],
                    "name": "file.inode",
                    "requirement_level": "recommended",
                    "root_namespace": "file",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Mode of the file in octal representation.\n",
                    "examples": [
                        "0640",
                    ],
                    "name": "file.mode",
                    "requirement_level": "recommended",
                    "root_namespace": "file",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Time when the file content was last modified, in ISO 8601 format.\n",
                    "examples": [
                        "2021-01-01T12:00:00Z",
                    ],
                    "name": "file.modified",
                    "requirement_level": "recommended",
                    "root_namespace": "file",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Name of the file including the extension, without the directory.\n",
                    "examples": [
                        "example.png",
                    ],
                    "name": "file.name",
                    "requirement_level": "recommended",
                    "root_namespace": "file",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The user ID (UID) or security identifier (SID) of the file owner.\n",
                    "examples": [
                        "1000",
                    ],
                    "name": "file.owner.id",
                    "requirement_level": "recommended",
                    "root_namespace": "file",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Username of the file owner.\n",
                    "examples": [
                        "root",
                    ],
                    "name": "file.owner.name",
                    "requirement_level": "recommended",
                    "root_namespace": "file",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Full path to the file, including the file name. It should include the drive letter, when appropriate.\n",
                    "examples": [
                        "/home/alice/example.png",
                        "C:\\Program Files\\MyApp\\myapp.exe",
                    ],
                    "name": "file.path",
                    "requirement_level": "recommended",
                    "root_namespace": "file",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "File size in bytes.\n",
                    "name": "file.size",
                    "requirement_level": "recommended",
                    "root_namespace": "file",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "Path to the target of a symbolic link.\n",
                    "examples": [
                        "/usr/bin/python3",
                    ],
                    "name": "file.symbolic_link.target_path",
                    "note": "This attribute is only applicable to symbolic links.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "file",
                    "stability": "development",
                    "type": "string",
                },
            ],
            "root_namespace": "file",
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
