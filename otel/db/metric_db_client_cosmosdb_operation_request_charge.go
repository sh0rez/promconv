package db

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Deprecated, use `azure.cosmosdb.client.operation.request_charge` instead.
type ClientCosmosdbOperationRequestCharge struct {
	*prometheus.HistogramVec
	extra ClientCosmosdbOperationRequestChargeExtra
}

func NewClientCosmosdbOperationRequestCharge() ClientCosmosdbOperationRequestCharge {
	labels := []string{AttrCollectionName("").Key(), AttrCosmosdbConsistencyLevel("").Key(), AttrCosmosdbSubStatusCode("").Key(), AttrNamespace("").Key(), AttrOperationName("").Key(), AttrCosmosdbRegionsContacted("").Key()}
	return ClientCosmosdbOperationRequestCharge{HistogramVec: prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "db_client_cosmosdb_operation_request_charge",
		Help: "Deprecated, use `azure.cosmosdb.client.operation.request_charge` instead.",
	}, labels)}
}

func (m ClientCosmosdbOperationRequestCharge) With(extras ...interface {
	DbCollectionName() AttrCollectionName
	DbCosmosdbConsistencyLevel() AttrCosmosdbConsistencyLevel
	DbCosmosdbSubStatusCode() AttrCosmosdbSubStatusCode
	DbNamespace() AttrNamespace
	DbOperationName() AttrOperationName
	DbCosmosdbRegionsContacted() AttrCosmosdbRegionsContacted
}) prometheus.Observer {
	if extras == nil {
		extras = append(extras, m.extra)
	}
	extra := extras[0]

	return m.HistogramVec.WithLabelValues(extra.DbCollectionName().Value(), extra.DbCosmosdbConsistencyLevel().Value(), extra.DbCosmosdbSubStatusCode().Value(), extra.DbNamespace().Value(), extra.DbOperationName().Value(), extra.DbCosmosdbRegionsContacted().Value())
}

// Deprecated: Use [ClientCosmosdbOperationRequestCharge.With] instead
func (m ClientCosmosdbOperationRequestCharge) WithLabelValues(lvs ...string) prometheus.Observer {
	return m.HistogramVec.WithLabelValues(lvs...)
}

func (a ClientCosmosdbOperationRequestCharge) WithCollectionName(attr interface{ DbCollectionName() AttrCollectionName }) ClientCosmosdbOperationRequestCharge {
	a.extra.AttrCollectionName = attr.DbCollectionName()
	return a
}
func (a ClientCosmosdbOperationRequestCharge) WithCosmosdbConsistencyLevel(attr interface {
	DbCosmosdbConsistencyLevel() AttrCosmosdbConsistencyLevel
}) ClientCosmosdbOperationRequestCharge {
	a.extra.AttrCosmosdbConsistencyLevel = attr.DbCosmosdbConsistencyLevel()
	return a
}
func (a ClientCosmosdbOperationRequestCharge) WithCosmosdbSubStatusCode(attr interface {
	DbCosmosdbSubStatusCode() AttrCosmosdbSubStatusCode
}) ClientCosmosdbOperationRequestCharge {
	a.extra.AttrCosmosdbSubStatusCode = attr.DbCosmosdbSubStatusCode()
	return a
}
func (a ClientCosmosdbOperationRequestCharge) WithNamespace(attr interface{ DbNamespace() AttrNamespace }) ClientCosmosdbOperationRequestCharge {
	a.extra.AttrNamespace = attr.DbNamespace()
	return a
}
func (a ClientCosmosdbOperationRequestCharge) WithOperationName(attr interface{ DbOperationName() AttrOperationName }) ClientCosmosdbOperationRequestCharge {
	a.extra.AttrOperationName = attr.DbOperationName()
	return a
}
func (a ClientCosmosdbOperationRequestCharge) WithCosmosdbRegionsContacted(attr interface {
	DbCosmosdbRegionsContacted() AttrCosmosdbRegionsContacted
}) ClientCosmosdbOperationRequestCharge {
	a.extra.AttrCosmosdbRegionsContacted = attr.DbCosmosdbRegionsContacted()
	return a
}

type ClientCosmosdbOperationRequestChargeExtra struct {
	// Cosmos DB container name
	AttrCollectionName           AttrCollectionName           `otel:"db.collection.name"`            // Deprecated, use `cosmosdb.consistency.level` instead
	AttrCosmosdbConsistencyLevel AttrCosmosdbConsistencyLevel `otel:"db.cosmosdb.consistency_level"` // Deprecated, use `azure.cosmosdb.response.sub_status_code` instead
	AttrCosmosdbSubStatusCode    AttrCosmosdbSubStatusCode    `otel:"db.cosmosdb.sub_status_code"`   // The name of the database, fully qualified within the server address and port
	AttrNamespace                AttrNamespace                `otel:"db.namespace"`                  // The name of the operation or command being executed
	AttrOperationName            AttrOperationName            `otel:"db.operation.name"`             // Deprecated, use `azure.cosmosdb.operation.contacted_regions` instead
	AttrCosmosdbRegionsContacted AttrCosmosdbRegionsContacted `otel:"db.cosmosdb.regions_contacted"`
}

func (a ClientCosmosdbOperationRequestChargeExtra) DbCollectionName() AttrCollectionName {
	return a.AttrCollectionName
}
func (a ClientCosmosdbOperationRequestChargeExtra) DbCosmosdbConsistencyLevel() AttrCosmosdbConsistencyLevel {
	return a.AttrCosmosdbConsistencyLevel
}
func (a ClientCosmosdbOperationRequestChargeExtra) DbCosmosdbSubStatusCode() AttrCosmosdbSubStatusCode {
	return a.AttrCosmosdbSubStatusCode
}
func (a ClientCosmosdbOperationRequestChargeExtra) DbNamespace() AttrNamespace {
	return a.AttrNamespace
}
func (a ClientCosmosdbOperationRequestChargeExtra) DbOperationName() AttrOperationName {
	return a.AttrOperationName
}
func (a ClientCosmosdbOperationRequestChargeExtra) DbCosmosdbRegionsContacted() AttrCosmosdbRegionsContacted {
	return a.AttrCosmosdbRegionsContacted
}

/*
State {
    name: "vec.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "ClientCosmosdbOperationRequestChargeExtra",
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
