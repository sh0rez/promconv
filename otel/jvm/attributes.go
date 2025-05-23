package jvm

// Name of the buffer pool.
// Pool names are generally obtained via [BufferPoolMXBean#getName()]
//
// [BufferPoolMXBean#getName()]: https://docs.oracle.com/en/java/javase/11/docs/api/java.management/java/lang/management/BufferPoolMXBean.html#getName()
type AttrBufferPoolName string // jvm.buffer.pool.name

func (AttrBufferPoolName) Development() {}
func (AttrBufferPoolName) Recommended() {}

// Name of the garbage collector action.
// Garbage collector action is generally obtained via [GarbageCollectionNotificationInfo#getGcAction()]
//
// [GarbageCollectionNotificationInfo#getGcAction()]: https://docs.oracle.com/en/java/javase/11/docs/api/jdk.management/com/sun/management/GarbageCollectionNotificationInfo.html#getGcAction()
type AttrGcAction string // jvm.gc.action

func (AttrGcAction) Stable()      {}
func (AttrGcAction) Recommended() {}

// Name of the garbage collector cause.
// Garbage collector cause is generally obtained via [GarbageCollectionNotificationInfo#getGcCause()]
//
// [GarbageCollectionNotificationInfo#getGcCause()]: https://docs.oracle.com/en/java/javase/11/docs/api/jdk.management/com/sun/management/GarbageCollectionNotificationInfo.html#getGcCause()
type AttrGcCause string // jvm.gc.cause

func (AttrGcCause) Development() {}
func (AttrGcCause) Recommended() {}

// Name of the garbage collector.
// Garbage collector name is generally obtained via [GarbageCollectionNotificationInfo#getGcName()]
//
// [GarbageCollectionNotificationInfo#getGcName()]: https://docs.oracle.com/en/java/javase/11/docs/api/jdk.management/com/sun/management/GarbageCollectionNotificationInfo.html#getGcName()
type AttrGcName string // jvm.gc.name

func (AttrGcName) Stable()      {}
func (AttrGcName) Recommended() {}

// Name of the memory pool.
// Pool names are generally obtained via [MemoryPoolMXBean#getName()]
//
// [MemoryPoolMXBean#getName()]: https://docs.oracle.com/en/java/javase/11/docs/api/java.management/java/lang/management/MemoryPoolMXBean.html#getName()
type AttrMemoryPoolName string // jvm.memory.pool.name

func (AttrMemoryPoolName) Stable()      {}
func (AttrMemoryPoolName) Recommended() {}

// The type of memory
type AttrMemoryType string // jvm.memory.type

func (AttrMemoryType) Stable()      {}
func (AttrMemoryType) Recommended() {}

// Whether the thread is daemon or not
type AttrThreadDaemon string // jvm.thread.daemon

func (AttrThreadDaemon) Stable()      {}
func (AttrThreadDaemon) Recommended() {}

// State of the thread
type AttrThreadState string // jvm.thread.state

func (AttrThreadState) Stable()      {}
func (AttrThreadState) Recommended() {}

/* State {
    name: "attr.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "ctx": {
            "attributes": [
                {
                    "brief": "Name of the buffer pool.",
                    "examples": [
                        "mapped",
                        "direct",
                    ],
                    "name": "jvm.buffer.pool.name",
                    "note": "Pool names are generally obtained via [BufferPoolMXBean#getName()](https://docs.oracle.com/en/java/javase/11/docs/api/java.management/java/lang/management/BufferPoolMXBean.html#getName()).\n",
                    "requirement_level": "recommended",
                    "root_namespace": "jvm",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Name of the garbage collector action.",
                    "examples": [
                        "end of minor GC",
                        "end of major GC",
                    ],
                    "name": "jvm.gc.action",
                    "note": "Garbage collector action is generally obtained via [GarbageCollectionNotificationInfo#getGcAction()](https://docs.oracle.com/en/java/javase/11/docs/api/jdk.management/com/sun/management/GarbageCollectionNotificationInfo.html#getGcAction()).\n",
                    "requirement_level": "recommended",
                    "root_namespace": "jvm",
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "Name of the garbage collector cause.",
                    "examples": [
                        "System.gc()",
                        "Allocation Failure",
                    ],
                    "name": "jvm.gc.cause",
                    "note": "Garbage collector cause is generally obtained via [GarbageCollectionNotificationInfo#getGcCause()](https://docs.oracle.com/en/java/javase/11/docs/api/jdk.management/com/sun/management/GarbageCollectionNotificationInfo.html#getGcCause()).\n",
                    "requirement_level": "recommended",
                    "root_namespace": "jvm",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Name of the garbage collector.",
                    "examples": [
                        "G1 Young Generation",
                        "G1 Old Generation",
                    ],
                    "name": "jvm.gc.name",
                    "note": "Garbage collector name is generally obtained via [GarbageCollectionNotificationInfo#getGcName()](https://docs.oracle.com/en/java/javase/11/docs/api/jdk.management/com/sun/management/GarbageCollectionNotificationInfo.html#getGcName()).\n",
                    "requirement_level": "recommended",
                    "root_namespace": "jvm",
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "Name of the memory pool.",
                    "examples": [
                        "G1 Old Gen",
                        "G1 Eden space",
                        "G1 Survivor Space",
                    ],
                    "name": "jvm.memory.pool.name",
                    "note": "Pool names are generally obtained via [MemoryPoolMXBean#getName()](https://docs.oracle.com/en/java/javase/11/docs/api/java.management/java/lang/management/MemoryPoolMXBean.html#getName()).\n",
                    "requirement_level": "recommended",
                    "root_namespace": "jvm",
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "The type of memory.",
                    "examples": [
                        "heap",
                        "non_heap",
                    ],
                    "name": "jvm.memory.type",
                    "requirement_level": "recommended",
                    "root_namespace": "jvm",
                    "stability": "stable",
                    "type": {
                        "allow_custom_values": none,
                        "members": [
                            {
                                "brief": "Heap memory.",
                                "deprecated": none,
                                "id": "heap",
                                "note": none,
                                "stability": "stable",
                                "value": "heap",
                            },
                            {
                                "brief": "Non-heap memory",
                                "deprecated": none,
                                "id": "non_heap",
                                "note": none,
                                "stability": "stable",
                                "value": "non_heap",
                            },
                        ],
                    },
                },
                {
                    "brief": "Whether the thread is daemon or not.",
                    "name": "jvm.thread.daemon",
                    "requirement_level": "recommended",
                    "root_namespace": "jvm",
                    "stability": "stable",
                    "type": "boolean",
                },
                {
                    "brief": "State of the thread.",
                    "examples": [
                        "runnable",
                        "blocked",
                    ],
                    "name": "jvm.thread.state",
                    "requirement_level": "recommended",
                    "root_namespace": "jvm",
                    "stability": "stable",
                    "type": {
                        "allow_custom_values": none,
                        "members": [
                            {
                                "brief": "A thread that has not yet started is in this state.",
                                "deprecated": none,
                                "id": "new",
                                "note": none,
                                "stability": "stable",
                                "value": "new",
                            },
                            {
                                "brief": "A thread executing in the Java virtual machine is in this state.",
                                "deprecated": none,
                                "id": "runnable",
                                "note": none,
                                "stability": "stable",
                                "value": "runnable",
                            },
                            {
                                "brief": "A thread that is blocked waiting for a monitor lock is in this state.",
                                "deprecated": none,
                                "id": "blocked",
                                "note": none,
                                "stability": "stable",
                                "value": "blocked",
                            },
                            {
                                "brief": "A thread that is waiting indefinitely for another thread to perform a particular action is in this state.",
                                "deprecated": none,
                                "id": "waiting",
                                "note": none,
                                "stability": "stable",
                                "value": "waiting",
                            },
                            {
                                "brief": "A thread that is waiting for another thread to perform an action for up to a specified waiting time is in this state.",
                                "deprecated": none,
                                "id": "timed_waiting",
                                "note": none,
                                "stability": "stable",
                                "value": "timed_waiting",
                            },
                            {
                                "brief": "A thread that has exited is in this state.",
                                "deprecated": none,
                                "id": "terminated",
                                "note": none,
                                "stability": "stable",
                                "value": "terminated",
                            },
                        ],
                    },
                },
            ],
            "root_namespace": "jvm",
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
