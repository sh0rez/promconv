package azure

import (
	"github.com/prometheus/client_golang/prometheus"
)

import (
	"shorez.de/promconv/otel/server"
)

// Number of active client instances
type CosmosdbClientActiveInstanceCount struct {
	*prometheus.GaugeVec
	extra CosmosdbClientActiveInstanceCountExtra
}

func NewCosmosdbClientActiveInstanceCount() CosmosdbClientActiveInstanceCount {
	labels := []string{server.AttrPort("").Key(), server.AttrAddress("").Key()}
	return CosmosdbClientActiveInstanceCount{GaugeVec: prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "azure_cosmosdb_client_active_instance_count",
		Help: "Number of active client instances",
	}, labels)}
}

func (m CosmosdbClientActiveInstanceCount) With(extras ...interface {
	ServerPort() server.AttrPort
	ServerAddress() server.AttrAddress
}) prometheus.Gauge {
	if extras == nil {
		extras = append(extras, m.extra)
	}
	extra := extras[0]

	return m.GaugeVec.WithLabelValues(extra.ServerPort().Value(), extra.ServerAddress().Value())
}

// Deprecated: Use [CosmosdbClientActiveInstanceCount.With] instead
func (m CosmosdbClientActiveInstanceCount) WithLabelValues(lvs ...string) prometheus.Gauge {
	return m.GaugeVec.WithLabelValues(lvs...)
}

func (a CosmosdbClientActiveInstanceCount) WithServerPort(attr interface{ ServerPort() server.AttrPort }) CosmosdbClientActiveInstanceCount {
	a.extra.AttrServerPort = attr.ServerPort()
	return a
}
func (a CosmosdbClientActiveInstanceCount) WithServerAddress(attr interface{ ServerAddress() server.AttrAddress }) CosmosdbClientActiveInstanceCount {
	a.extra.AttrServerAddress = attr.ServerAddress()
	return a
}

type CosmosdbClientActiveInstanceCountExtra struct {
	// Server port number
	AttrServerPort    server.AttrPort    `otel:"server.port"` // Name of the database host
	AttrServerAddress server.AttrAddress `otel:"server.address"`
}

func (a CosmosdbClientActiveInstanceCountExtra) ServerPort() server.AttrPort { return a.AttrServerPort }
func (a CosmosdbClientActiveInstanceCountExtra) ServerAddress() server.AttrAddress {
	return a.AttrServerAddress
}

/*
State {
    name: "vec.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "CosmosdbClientActiveInstanceCountExtra",
        "Instr": "Gauge",
        "InstrMap": {
            "counter": "Counter",
            "gauge": "Gauge",
            "histogram": "Histogram",
            "updowncounter": "Gauge",
        },
        "Name": "cosmosdb.client.active_instance.count",
        "Type": "CosmosdbClientActiveInstanceCount",
        "attributes": [
            {
                "brief": "Server port number.",
                "examples": [
                    80,
                    8080,
                    443,
                ],
                "name": "server.port",
                "note": "When observed from the client side, and when communicating through an intermediary, `server.port` SHOULD represent the server port behind any intermediaries, for example proxies, if it's available.\n",
                "requirement_level": {
                    "conditionally_required": "If using a port other than the default port for this DBMS and if `server.address` is set.",
                },
                "stability": "stable",
                "type": "int",
            },
            {
                "brief": "Name of the database host.\n",
                "examples": [
                    "example.com",
                    "10.1.2.80",
                    "/tmp/my.sock",
                ],
                "name": "server.address",
                "note": "When observed from the client side, and when communicating through an intermediary, `server.address` SHOULD represent the server address behind any intermediaries, for example proxies, if it's available.\n",
                "requirement_level": "recommended",
                "stability": "stable",
                "type": "string",
            },
        ],
        "ctx": {
            "attributes": [
                {
                    "brief": "Name of the database host.\n",
                    "examples": [
                        "example.com",
                        "10.1.2.80",
                        "/tmp/my.sock",
                    ],
                    "name": "server.address",
                    "note": "When observed from the client side, and when communicating through an intermediary, `server.address` SHOULD represent the server address behind any intermediaries, for example proxies, if it's available.\n",
                    "requirement_level": "recommended",
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "Server port number.",
                    "examples": [
                        80,
                        8080,
                        443,
                    ],
                    "name": "server.port",
                    "note": "When observed from the client side, and when communicating through an intermediary, `server.port` SHOULD represent the server port behind any intermediaries, for example proxies, if it's available.\n",
                    "requirement_level": {
                        "conditionally_required": "If using a port other than the default port for this DBMS and if `server.address` is set.",
                    },
                    "stability": "stable",
                    "type": "int",
                },
            ],
            "brief": "Number of active client instances",
            "events": [],
            "id": "metric.azure.cosmosdb.client.active_instance.count",
            "instrument": "updowncounter",
            "lineage": {
                "attributes": {
                    "server.address": {
                        "inherited_fields": [
                            "examples",
                            "note",
                            "requirement_level",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "brief",
                        ],
                        "source_group": "registry.server",
                    },
                    "server.port": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "requirement_level",
                        ],
                        "source_group": "registry.server",
                    },
                },
                "provenance": {
                    "path": "https://github.com/open-telemetry/semantic-conventions.git[model]/azure/cosmosdb-metrics.yaml",
                    "registry_id": "main",
                },
            },
            "metric_name": "azure.cosmosdb.client.active_instance.count",
            "name": none,
            "root_namespace": "azure",
            "span_kind": none,
            "stability": "development",
            "type": "metric",
            "unit": "{instance}",
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
