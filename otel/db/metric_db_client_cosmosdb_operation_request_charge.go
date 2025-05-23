package db

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Deprecated, use `azure.cosmosdb.client.operation.request_charge` instead.
type ClientCosmosdbOperationRequestCharge struct {
	*prometheus.HistogramVec
}

func NewClientCosmosdbOperationRequestCharge() ClientCosmosdbOperationRequestCharge {
	labels := []string{"db_collection_name", "db_cosmosdb_consistency_level", "db_cosmosdb_sub_status_code", "db_namespace", "db_operation_name", "db_cosmosdb_regions_contacted"}
	return ClientCosmosdbOperationRequestCharge{HistogramVec: prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "db",
		Name:      "client_cosmosdb_operation_request_charge",
		Help:      "Deprecated, use `azure.cosmosdb.client.operation.request_charge` instead.",
	}, labels)}
}

func (m ClientCosmosdbOperationRequestCharge) With(extra ClientCosmosdbOperationRequestChargeOptional) prometheus.Observer {
	return m.WithLabelValues(
		string(extra.DbCollectionName),
		string(extra.DbCosmosdbConsistencyLevel),
		string(extra.DbCosmosdbSubStatusCode),
		string(extra.DbNamespace),
		string(extra.DbOperationName),
		string(extra.DbCosmosdbRegionsContacted),
	)
}

type ClientCosmosdbOperationRequestChargeOptional struct {
	// Cosmos DB container name.
	DbCollectionName AttrCollectionName `otel:"db.collection.name"`
	// Deprecated, use `cosmosdb.consistency.level` instead.
	DbCosmosdbConsistencyLevel AttrCosmosdbConsistencyLevel `otel:"db.cosmosdb.consistency_level"`
	// Deprecated, use `azure.cosmosdb.response.sub_status_code` instead.
	DbCosmosdbSubStatusCode AttrCosmosdbSubStatusCode `otel:"db.cosmosdb.sub_status_code"`
	// The name of the database, fully qualified within the server address and port.
	DbNamespace AttrNamespace `otel:"db.namespace"`
	// The name of the operation or command being executed.
	DbOperationName AttrOperationName `otel:"db.operation.name"`
	// Deprecated, use `azure.cosmosdb.operation.contacted_regions` instead.
	DbCosmosdbRegionsContacted AttrCosmosdbRegionsContacted `otel:"db.cosmosdb.regions_contacted"`
}

/*
State {
    name: "metric.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "ClientCosmosdbOperationRequestChargeOptional",
        "Instr": "Histogram",
        "InstrMap": {
            "counter": "Counter",
            "gauge": "Gauge",
            "histogram": "Histogram",
            "updowncounter": "Gauge",
        },
        "Name": "client.cosmosdb.operation.request_charge",
        "Type": "ClientCosmosdbOperationRequestCharge",
        "attributes": [
            {
                "brief": "Cosmos DB container name.",
                "examples": [
                    "public.users",
                    "customers",
                ],
                "name": "db.collection.name",
                "note": "It is RECOMMENDED to capture the value as provided by the application\nwithout attempting to do any case normalization.\n\nThe collection name SHOULD NOT be extracted from `db.query.text`,\nwhen the database system supports query text with multiple collections\nin non-batch operations.\n\nFor batch operations, if the individual operations are known to have the same\ncollection name then that collection name SHOULD be used.\n",
                "requirement_level": {
                    "conditionally_required": "If available.",
                },
                "stability": "stable",
                "type": "string",
            },
            {
                "brief": "Deprecated, use `cosmosdb.consistency.level` instead.",
                "deprecated": {
                    "note": "Replaced by `azure.cosmosdb.consistency.level`.",
                    "reason": "renamed",
                    "renamed_to": "azure.cosmosdb.consistency.level",
                },
                "examples": [
                    "Eventual",
                    "ConsistentPrefix",
                    "BoundedStaleness",
                    "Strong",
                    "Session",
                ],
                "name": "db.cosmosdb.consistency_level",
                "requirement_level": {
                    "conditionally_required": "If available.",
                },
                "stability": "development",
                "type": {
                    "allow_custom_values": none,
                    "members": [
                        {
                            "brief": none,
                            "deprecated": none,
                            "id": "strong",
                            "note": none,
                            "stability": "development",
                            "value": "Strong",
                        },
                        {
                            "brief": none,
                            "deprecated": none,
                            "id": "bounded_staleness",
                            "note": none,
                            "stability": "development",
                            "value": "BoundedStaleness",
                        },
                        {
                            "brief": none,
                            "deprecated": none,
                            "id": "session",
                            "note": none,
                            "stability": "development",
                            "value": "Session",
                        },
                        {
                            "brief": none,
                            "deprecated": none,
                            "id": "eventual",
                            "note": none,
                            "stability": "development",
                            "value": "Eventual",
                        },
                        {
                            "brief": none,
                            "deprecated": none,
                            "id": "consistent_prefix",
                            "note": none,
                            "stability": "development",
                            "value": "ConsistentPrefix",
                        },
                    ],
                },
            },
            {
                "brief": "Deprecated, use `azure.cosmosdb.response.sub_status_code` instead.",
                "deprecated": {
                    "note": "Replaced by `azure.cosmosdb.response.sub_status_code`.",
                    "reason": "renamed",
                    "renamed_to": "azure.cosmosdb.response.sub_status_code",
                },
                "examples": [
                    1000,
                    1002,
                ],
                "name": "db.cosmosdb.sub_status_code",
                "requirement_level": {
                    "conditionally_required": "when response was received and contained sub-code.",
                },
                "stability": "development",
                "type": "int",
            },
            {
                "brief": "The name of the database, fully qualified within the server address and port.\n",
                "examples": [
                    "customers",
                    "test.users",
                ],
                "name": "db.namespace",
                "requirement_level": {
                    "conditionally_required": "If available.",
                },
                "stability": "stable",
                "type": "string",
            },
            {
                "brief": "The name of the operation or command being executed.\n",
                "examples": [
                    "findAndModify",
                    "HMSET",
                    "SELECT",
                ],
                "name": "db.operation.name",
                "note": "It is RECOMMENDED to capture the value as provided by the application\nwithout attempting to do any case normalization.\n\nThe operation name SHOULD NOT be extracted from `db.query.text`,\nwhen the database system supports query text with multiple operations\nin non-batch operations.\n\nIf spaces can occur in the operation name, multiple consecutive spaces\nSHOULD be normalized to a single space.\n\nFor batch operations, if the individual operations are known to have the same operation name\nthen that operation name SHOULD be used prepended by `BATCH `,\notherwise `db.operation.name` SHOULD be `BATCH` or some other database\nsystem specific term if more applicable.\n",
                "requirement_level": {
                    "conditionally_required": "If readily available and if there is a single operation name that describes the database call. The operation name MAY be parsed from the query text, in which case it SHOULD be the single operation name found in the query.\n",
                },
                "stability": "stable",
                "type": "string",
            },
            {
                "brief": "Deprecated, use `azure.cosmosdb.operation.contacted_regions` instead.",
                "deprecated": {
                    "note": "Replaced by `azure.cosmosdb.operation.contacted_regions`.",
                    "reason": "renamed",
                    "renamed_to": "azure.cosmosdb.operation.contacted_regions",
                },
                "examples": [
                    [
                        "North Central US",
                        "Australia East",
                        "Australia Southeast",
                    ],
                ],
                "name": "db.cosmosdb.regions_contacted",
                "requirement_level": {
                    "recommended": "If available",
                },
                "stability": "development",
                "type": "string[]",
            },
        ],
        "ctx": {
            "attributes": [
                {
                    "brief": "Cosmos DB container name.",
                    "examples": [
                        "public.users",
                        "customers",
                    ],
                    "name": "db.collection.name",
                    "note": "It is RECOMMENDED to capture the value as provided by the application\nwithout attempting to do any case normalization.\n\nThe collection name SHOULD NOT be extracted from `db.query.text`,\nwhen the database system supports query text with multiple collections\nin non-batch operations.\n\nFor batch operations, if the individual operations are known to have the same\ncollection name then that collection name SHOULD be used.\n",
                    "requirement_level": {
                        "conditionally_required": "If available.",
                    },
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "The name of the database, fully qualified within the server address and port.\n",
                    "examples": [
                        "customers",
                        "test.users",
                    ],
                    "name": "db.namespace",
                    "requirement_level": {
                        "conditionally_required": "If available.",
                    },
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "The name of the operation or command being executed.\n",
                    "examples": [
                        "findAndModify",
                        "HMSET",
                        "SELECT",
                    ],
                    "name": "db.operation.name",
                    "note": "It is RECOMMENDED to capture the value as provided by the application\nwithout attempting to do any case normalization.\n\nThe operation name SHOULD NOT be extracted from `db.query.text`,\nwhen the database system supports query text with multiple operations\nin non-batch operations.\n\nIf spaces can occur in the operation name, multiple consecutive spaces\nSHOULD be normalized to a single space.\n\nFor batch operations, if the individual operations are known to have the same operation name\nthen that operation name SHOULD be used prepended by `BATCH `,\notherwise `db.operation.name` SHOULD be `BATCH` or some other database\nsystem specific term if more applicable.\n",
                    "requirement_level": {
                        "conditionally_required": "If readily available and if there is a single operation name that describes the database call. The operation name MAY be parsed from the query text, in which case it SHOULD be the single operation name found in the query.\n",
                    },
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "Deprecated, use `azure.cosmosdb.response.sub_status_code` instead.",
                    "deprecated": {
                        "note": "Replaced by `azure.cosmosdb.response.sub_status_code`.",
                        "reason": "renamed",
                        "renamed_to": "azure.cosmosdb.response.sub_status_code",
                    },
                    "examples": [
                        1000,
                        1002,
                    ],
                    "name": "db.cosmosdb.sub_status_code",
                    "requirement_level": {
                        "conditionally_required": "when response was received and contained sub-code.",
                    },
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "Deprecated, use `cosmosdb.consistency.level` instead.",
                    "deprecated": {
                        "note": "Replaced by `azure.cosmosdb.consistency.level`.",
                        "reason": "renamed",
                        "renamed_to": "azure.cosmosdb.consistency.level",
                    },
                    "examples": [
                        "Eventual",
                        "ConsistentPrefix",
                        "BoundedStaleness",
                        "Strong",
                        "Session",
                    ],
                    "name": "db.cosmosdb.consistency_level",
                    "requirement_level": {
                        "conditionally_required": "If available.",
                    },
                    "stability": "development",
                    "type": {
                        "allow_custom_values": none,
                        "members": [
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "strong",
                                "note": none,
                                "stability": "development",
                                "value": "Strong",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "bounded_staleness",
                                "note": none,
                                "stability": "development",
                                "value": "BoundedStaleness",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "session",
                                "note": none,
                                "stability": "development",
                                "value": "Session",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "eventual",
                                "note": none,
                                "stability": "development",
                                "value": "Eventual",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "consistent_prefix",
                                "note": none,
                                "stability": "development",
                                "value": "ConsistentPrefix",
                            },
                        ],
                    },
                },
                {
                    "brief": "Deprecated, use `azure.cosmosdb.operation.contacted_regions` instead.",
                    "deprecated": {
                        "note": "Replaced by `azure.cosmosdb.operation.contacted_regions`.",
                        "reason": "renamed",
                        "renamed_to": "azure.cosmosdb.operation.contacted_regions",
                    },
                    "examples": [
                        [
                            "North Central US",
                            "Australia East",
                            "Australia Southeast",
                        ],
                    ],
                    "name": "db.cosmosdb.regions_contacted",
                    "requirement_level": {
                        "recommended": "If available",
                    },
                    "stability": "development",
                    "type": "string[]",
                },
            ],
            "brief": "Deprecated, use `azure.cosmosdb.client.operation.request_charge` instead.",
            "deprecated": {
                "note": "Replaced by `azure.cosmosdb.client.operation.request_charge`.",
                "reason": "renamed",
                "renamed_to": "azure.cosmosdb.client.operation.request_charge",
            },
            "events": [],
            "id": "metric.db.client.cosmosdb.operation.request_charge",
            "instrument": "histogram",
            "lineage": {
                "attributes": {
                    "db.collection.name": {
                        "inherited_fields": [
                            "examples",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "brief",
                            "requirement_level",
                        ],
                        "source_group": "registry.db",
                    },
                    "db.cosmosdb.consistency_level": {
                        "inherited_fields": [
                            "brief",
                            "deprecated",
                            "examples",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "requirement_level",
                        ],
                        "source_group": "registry.db.deprecated",
                    },
                    "db.cosmosdb.regions_contacted": {
                        "inherited_fields": [
                            "brief",
                            "deprecated",
                            "examples",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "requirement_level",
                        ],
                        "source_group": "registry.db.deprecated",
                    },
                    "db.cosmosdb.sub_status_code": {
                        "inherited_fields": [
                            "brief",
                            "deprecated",
                            "examples",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "requirement_level",
                        ],
                        "source_group": "registry.db.deprecated",
                    },
                    "db.namespace": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "note",
                            "requirement_level",
                        ],
                        "source_group": "registry.db",
                    },
                    "db.operation.name": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "requirement_level",
                        ],
                        "source_group": "registry.db",
                    },
                },
                "provenance": {
                    "path": "https://github.com/open-telemetry/semantic-conventions.git[model]/database/deprecated/metrics-deprecated.yaml",
                    "registry_id": "main",
                },
            },
            "metric_name": "db.client.cosmosdb.operation.request_charge",
            "name": none,
            "root_namespace": "db",
            "span_kind": none,
            "stability": "development",
            "type": "metric",
            "unit": "{request_unit}",
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
