package code

// Deprecated, use `code.column.number`
type AttrColumn string // code.column

func (AttrColumn) Development() {}
func (AttrColumn) Recommended() {}

// The column number in `code.file.path` best representing the operation. It SHOULD point within the code unit named in `code.function.name`. This attribute MUST NOT be used on the Profile signal since the data is already captured in 'message Line'. This constraint is imposed to prevent redundancy and maintain data integrity
type AttrColumnNumber string // code.column.number

func (AttrColumnNumber) Stable()      {}
func (AttrColumnNumber) Recommended() {}

// The source code file name that identifies the code unit as uniquely as possible (preferably an absolute file path). This attribute MUST NOT be used on the Profile signal since the data is already captured in 'message Function'. This constraint is imposed to prevent redundancy and maintain data integrity
type AttrFilePath string // code.file.path

func (AttrFilePath) Stable()      {}
func (AttrFilePath) Recommended() {}

// Deprecated, use `code.file.path` instead
type AttrFilepath string // code.filepath

func (AttrFilepath) Development() {}
func (AttrFilepath) Recommended() {}

// Deprecated, use `code.function.name` instead
type AttrFunction string // code.function

func (AttrFunction) Development() {}
func (AttrFunction) Recommended() {}

// The method or function fully-qualified name without arguments. The value should fit the natural representation of the language runtime, which is also likely the same used within `code.stacktrace` attribute value. This attribute MUST NOT be used on the Profile signal since the data is already captured in 'message Function'. This constraint is imposed to prevent redundancy and maintain data integrity.
//
// Values and format depends on each language runtime, thus it is impossible to provide an exhaustive list of examples.
// The values are usually the same (or prefixes of) the ones found in native stack trace representation stored in
// `code.stacktrace` without information on arguments.
//
// Examples:
//
//   - Java method: `com.example.MyHttpService.serveRequest`
//   - Java anonymous class method: `com.mycompany.Main$1.myMethod`
//   - Java lambda method: `com.mycompany.Main$$Lambda/0x0000748ae4149c00.myMethod`
//   - PHP function: `GuzzleHttp\Client::transfer`
//   - Go function: `github.com/my/repo/pkg.foo.func5`
//   - Elixir: `OpenTelemetry.Ctx.new`
//   - Erlang: `opentelemetry_ctx:new`
//   - Rust: `playground::my_module::my_cool_func`
//   - C function: `fopen`
type AttrFunctionName string // code.function.name

func (AttrFunctionName) Stable()      {}
func (AttrFunctionName) Recommended() {}

// The line number in `code.file.path` best representing the operation. It SHOULD point within the code unit named in `code.function.name`. This attribute MUST NOT be used on the Profile signal since the data is already captured in 'message Line'. This constraint is imposed to prevent redundancy and maintain data integrity
type AttrLineNumber string // code.line.number

func (AttrLineNumber) Stable()      {}
func (AttrLineNumber) Recommended() {}

// Deprecated, use `code.line.number` instead
type AttrLineno string // code.lineno

func (AttrLineno) Development() {}
func (AttrLineno) Recommended() {}

// Deprecated, namespace is now included into `code.function.name`
type AttrNamespace string // code.namespace

func (AttrNamespace) Development() {}
func (AttrNamespace) Recommended() {}

// A stacktrace as a string in the natural representation for the language runtime. The representation is identical to [`exception.stacktrace`]. This attribute MUST NOT be used on the Profile signal since the data is already captured in 'message Location'. This constraint is imposed to prevent redundancy and maintain data integrity
//
// [`exception.stacktrace`]: /docs/exceptions/exceptions-spans.md#stacktrace-representation
type AttrStacktrace string // code.stacktrace

func (AttrStacktrace) Stable()      {}
func (AttrStacktrace) Recommended() {}

/* State {
    name: "attr.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "ctx": {
            "attributes": [
                {
                    "brief": "Deprecated, use `code.column.number`\n",
                    "deprecated": {
                        "note": "Replaced by `code.column.number`.",
                        "reason": "renamed",
                        "renamed_to": "code.column.number",
                    },
                    "examples": 16,
                    "name": "code.column",
                    "requirement_level": "recommended",
                    "root_namespace": "code",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "The column number in `code.file.path` best representing the operation. It SHOULD point within the code unit named in `code.function.name`. This attribute MUST NOT be used on the Profile signal since the data is already captured in 'message Line'. This constraint is imposed to prevent redundancy and maintain data integrity.\n",
                    "examples": 16,
                    "name": "code.column.number",
                    "requirement_level": "recommended",
                    "root_namespace": "code",
                    "stability": "stable",
                    "type": "int",
                },
                {
                    "brief": "The source code file name that identifies the code unit as uniquely as possible (preferably an absolute file path). This attribute MUST NOT be used on the Profile signal since the data is already captured in 'message Function'. This constraint is imposed to prevent redundancy and maintain data integrity.\n",
                    "examples": "/usr/local/MyApplication/content_root/app/index.php",
                    "name": "code.file.path",
                    "requirement_level": "recommended",
                    "root_namespace": "code",
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "Deprecated, use `code.file.path` instead\n",
                    "deprecated": {
                        "note": "Replaced by `code.file.path`.",
                        "reason": "renamed",
                        "renamed_to": "code.file.path",
                    },
                    "examples": "/usr/local/MyApplication/content_root/app/index.php",
                    "name": "code.filepath",
                    "requirement_level": "recommended",
                    "root_namespace": "code",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Deprecated, use `code.function.name` instead\n",
                    "deprecated": {
                        "note": "Value should be included in `code.function.name` which is expected to be a fully-qualified name.\n",
                        "reason": "uncategorized",
                    },
                    "examples": "serveRequest",
                    "name": "code.function",
                    "requirement_level": "recommended",
                    "root_namespace": "code",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The method or function fully-qualified name without arguments. The value should fit the natural representation of the language runtime, which is also likely the same used within `code.stacktrace` attribute value. This attribute MUST NOT be used on the Profile signal since the data is already captured in 'message Function'. This constraint is imposed to prevent redundancy and maintain data integrity.\n",
                    "examples": [
                        "com.example.MyHttpService.serveRequest",
                        "GuzzleHttp\\Client::transfer",
                        "fopen",
                    ],
                    "name": "code.function.name",
                    "note": "Values and format depends on each language runtime, thus it is impossible to provide an exhaustive list of examples.\nThe values are usually the same (or prefixes of) the ones found in native stack trace representation stored in\n`code.stacktrace` without information on arguments.\n\nExamples:\n\n* Java method: `com.example.MyHttpService.serveRequest`\n* Java anonymous class method: `com.mycompany.Main$1.myMethod`\n* Java lambda method: `com.mycompany.Main$$Lambda/0x0000748ae4149c00.myMethod`\n* PHP function: `GuzzleHttp\\Client::transfer`\n* Go function: `github.com/my/repo/pkg.foo.func5`\n* Elixir: `OpenTelemetry.Ctx.new`\n* Erlang: `opentelemetry_ctx:new`\n* Rust: `playground::my_module::my_cool_func`\n* C function: `fopen`\n",
                    "requirement_level": "recommended",
                    "root_namespace": "code",
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "The line number in `code.file.path` best representing the operation. It SHOULD point within the code unit named in `code.function.name`. This attribute MUST NOT be used on the Profile signal since the data is already captured in 'message Line'. This constraint is imposed to prevent redundancy and maintain data integrity.\n",
                    "examples": 42,
                    "name": "code.line.number",
                    "requirement_level": "recommended",
                    "root_namespace": "code",
                    "stability": "stable",
                    "type": "int",
                },
                {
                    "brief": "Deprecated, use `code.line.number` instead\n",
                    "deprecated": {
                        "note": "Replaced by `code.line.number`.",
                        "reason": "renamed",
                        "renamed_to": "code.line.number",
                    },
                    "examples": 42,
                    "name": "code.lineno",
                    "requirement_level": "recommended",
                    "root_namespace": "code",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "Deprecated, namespace is now included into `code.function.name`\n",
                    "deprecated": {
                        "note": "Value should be included in `code.function.name` which is expected to be a fully-qualified name.\n",
                        "reason": "uncategorized",
                    },
                    "examples": "com.example.MyHttpService",
                    "name": "code.namespace",
                    "requirement_level": "recommended",
                    "root_namespace": "code",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "A stacktrace as a string in the natural representation for the language runtime. The representation is identical to [`exception.stacktrace`](/docs/exceptions/exceptions-spans.md#stacktrace-representation). This attribute MUST NOT be used on the Profile signal since the data is already captured in 'message Location'. This constraint is imposed to prevent redundancy and maintain data integrity.\n",
                    "examples": "at com.example.GenerateTrace.methodB(GenerateTrace.java:13)\\n at com.example.GenerateTrace.methodA(GenerateTrace.java:9)\\n at com.example.GenerateTrace.main(GenerateTrace.java:5)\n",
                    "name": "code.stacktrace",
                    "requirement_level": "recommended",
                    "root_namespace": "code",
                    "stability": "stable",
                    "type": "string",
                },
            ],
            "root_namespace": "code",
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
