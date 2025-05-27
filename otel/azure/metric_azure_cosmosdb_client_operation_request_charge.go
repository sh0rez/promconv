package azure

import (
	"github.com/prometheus/client_golang/prometheus"
)

import (
	"shorez.de/promconv/otel/db"
	"shorez.de/promconv/otel/error"
	"shorez.de/promconv/otel/server"
)

// [Request units](https://learn.microsoft.com/azure/cosmos-db/request-units) consumed by the operation
type CosmosdbClientOperationRequestCharge struct {
	*prometheus.HistogramVec
	extra CosmosdbClientOperationRequestChargeExtra
}

func NewCosmosdbClientOperationRequestCharge() CosmosdbClientOperationRequestCharge {
	labels := []string{"db_operation_name", "azure_cosmosdb_consistency_level", "azure_cosmosdb_response_sub_status_code", "db_collection_name", "db_namespace", "db_response_status_code", "error_type", "server_port", "azure_cosmosdb_operation_contacted_regions", "server_address"}
	return CosmosdbClientOperationRequestCharge{HistogramVec: prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "azure",
		Name:      "cosmosdb_client_operation_request_charge",
		Help:      "[Request units](https://learn.microsoft.com/azure/cosmos-db/request-units) consumed by the operation",
	}, labels)}
}

func (m CosmosdbClientOperationRequestCharge) With(operationName db.AttrOperationName, extra interface {
	AttrAzureCosmosdbConsistencyLevel() AttrCosmosdbConsistencyLevel
	AttrAzureCosmosdbResponseSubStatusCode() AttrCosmosdbResponseSubStatusCode
	AttrDbCollectionName() db.AttrCollectionName
	AttrDbNamespace() db.AttrNamespace
	AttrDbResponseStatusCode() db.AttrResponseStatusCode
	AttrErrorType() error.AttrType
	AttrServerPort() server.AttrPort
	AttrAzureCosmosdbOperationContactedRegions() AttrCosmosdbOperationContactedRegions
	AttrServerAddress() server.AttrAddress
}) prometheus.Observer {
	if extra == nil {
		extra = m.extra
	}
	return m.WithLabelValues(
		string(operationName),
		string(extra.AttrAzureCosmosdbConsistencyLevel()),
		string(extra.AttrAzureCosmosdbResponseSubStatusCode()),
		string(extra.AttrDbCollectionName()),
		string(extra.AttrDbNamespace()),
		string(extra.AttrDbResponseStatusCode()),
		string(extra.AttrErrorType()),
		string(extra.AttrServerPort()),
		string(extra.AttrAzureCosmosdbOperationContactedRegions()),
		string(extra.AttrServerAddress()),
	)
}

func (a CosmosdbClientOperationRequestCharge) WithAzureCosmosdbConsistencyLevel(attr interface {
	AttrAzureCosmosdbConsistencyLevel() AttrCosmosdbConsistencyLevel
}) CosmosdbClientOperationRequestCharge {
	a.extra.AzureCosmosdbConsistencyLevel = attr.AttrAzureCosmosdbConsistencyLevel()
	return a
}
func (a CosmosdbClientOperationRequestCharge) WithAzureCosmosdbResponseSubStatusCode(attr interface {
	AttrAzureCosmosdbResponseSubStatusCode() AttrCosmosdbResponseSubStatusCode
}) CosmosdbClientOperationRequestCharge {
	a.extra.AzureCosmosdbResponseSubStatusCode = attr.AttrAzureCosmosdbResponseSubStatusCode()
	return a
}
func (a CosmosdbClientOperationRequestCharge) WithDbCollectionName(attr interface{ AttrDbCollectionName() db.AttrCollectionName }) CosmosdbClientOperationRequestCharge {
	a.extra.DbCollectionName = attr.AttrDbCollectionName()
	return a
}
func (a CosmosdbClientOperationRequestCharge) WithDbNamespace(attr interface{ AttrDbNamespace() db.AttrNamespace }) CosmosdbClientOperationRequestCharge {
	a.extra.DbNamespace = attr.AttrDbNamespace()
	return a
}
func (a CosmosdbClientOperationRequestCharge) WithDbResponseStatusCode(attr interface {
	AttrDbResponseStatusCode() db.AttrResponseStatusCode
}) CosmosdbClientOperationRequestCharge {
	a.extra.DbResponseStatusCode = attr.AttrDbResponseStatusCode()
	return a
}
func (a CosmosdbClientOperationRequestCharge) WithErrorType(attr interface{ AttrErrorType() error.AttrType }) CosmosdbClientOperationRequestCharge {
	a.extra.ErrorType = attr.AttrErrorType()
	return a
}
func (a CosmosdbClientOperationRequestCharge) WithServerPort(attr interface{ AttrServerPort() server.AttrPort }) CosmosdbClientOperationRequestCharge {
	a.extra.ServerPort = attr.AttrServerPort()
	return a
}
func (a CosmosdbClientOperationRequestCharge) WithAzureCosmosdbOperationContactedRegions(attr interface {
	AttrAzureCosmosdbOperationContactedRegions() AttrCosmosdbOperationContactedRegions
}) CosmosdbClientOperationRequestCharge {
	a.extra.AzureCosmosdbOperationContactedRegions = attr.AttrAzureCosmosdbOperationContactedRegions()
	return a
}
func (a CosmosdbClientOperationRequestCharge) WithServerAddress(attr interface{ AttrServerAddress() server.AttrAddress }) CosmosdbClientOperationRequestCharge {
	a.extra.ServerAddress = attr.AttrServerAddress()
	return a
}

type CosmosdbClientOperationRequestChargeExtra struct {
	// Account or request [consistency level](https://learn.microsoft.com/azure/cosmos-db/consistency-levels).
	AzureCosmosdbConsistencyLevel AttrCosmosdbConsistencyLevel `otel:"azure.cosmosdb.consistency.level"`
	// Cosmos DB sub status code.
	AzureCosmosdbResponseSubStatusCode AttrCosmosdbResponseSubStatusCode `otel:"azure.cosmosdb.response.sub_status_code"`
	// Cosmos DB container name.
	DbCollectionName db.AttrCollectionName `otel:"db.collection.name"`
	// The name of the database, fully qualified within the server address and port.
	DbNamespace db.AttrNamespace `otel:"db.namespace"`
	// Database response status code.
	DbResponseStatusCode db.AttrResponseStatusCode `otel:"db.response.status_code"`
	// Describes a class of error the operation ended with.
	ErrorType error.AttrType `otel:"error.type"`
	// Server port number.
	ServerPort server.AttrPort `otel:"server.port"`
	// List of regions contacted during operation in the order that they were contacted. If there is more than one region listed, it indicates that the operation was performed on multiple regions i.e. cross-regional call.
	AzureCosmosdbOperationContactedRegions AttrCosmosdbOperationContactedRegions `otel:"azure.cosmosdb.operation.contacted_regions"`
	// Name of the database host.
	ServerAddress server.AttrAddress `otel:"server.address"`
}

func (a CosmosdbClientOperationRequestChargeExtra) AttrAzureCosmosdbConsistencyLevel() AttrCosmosdbConsistencyLevel {
	return a.AzureCosmosdbConsistencyLevel
}
func (a CosmosdbClientOperationRequestChargeExtra) AttrAzureCosmosdbResponseSubStatusCode() AttrCosmosdbResponseSubStatusCode {
	return a.AzureCosmosdbResponseSubStatusCode
}
func (a CosmosdbClientOperationRequestChargeExtra) AttrDbCollectionName() db.AttrCollectionName {
	return a.DbCollectionName
}
func (a CosmosdbClientOperationRequestChargeExtra) AttrDbNamespace() db.AttrNamespace {
	return a.DbNamespace
}
func (a CosmosdbClientOperationRequestChargeExtra) AttrDbResponseStatusCode() db.AttrResponseStatusCode {
	return a.DbResponseStatusCode
}
func (a CosmosdbClientOperationRequestChargeExtra) AttrErrorType() error.AttrType { return a.ErrorType }
func (a CosmosdbClientOperationRequestChargeExtra) AttrServerPort() server.AttrPort {
	return a.ServerPort
}
func (a CosmosdbClientOperationRequestChargeExtra) AttrAzureCosmosdbOperationContactedRegions() AttrCosmosdbOperationContactedRegions {
	return a.AzureCosmosdbOperationContactedRegions
}
func (a CosmosdbClientOperationRequestChargeExtra) AttrServerAddress() server.AttrAddress {
	return a.ServerAddress
}

/*
State {
    name: "metric.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "CosmosdbClientOperationRequestChargeExtra",
        "Instr": "Histogram",
        "InstrMap": {
            "counter": "Counter",
            "gauge": "Gauge",
            "histogram": "Histogram",
            "updowncounter": "Gauge",
        },
        "Name": "cosmosdb.client.operation.request_charge",
        "Type": "CosmosdbClientOperationRequestCharge",
        "attributes": [
            {
                "brief": "The name of the operation or command being executed.\n",
                "examples": [
                    "findAndModify",
                    "HMSET",
                    "SELECT",
                ],
                "name": "db.operation.name",
                "note": "It is RECOMMENDED to capture the value as provided by the application\nwithout attempting to do any case normalization.\n\nThe operation name SHOULD NOT be extracted from `db.query.text`,\nwhen the database system supports query text with multiple operations\nin non-batch operations.\n\nIf spaces can occur in the operation name, multiple consecutive spaces\nSHOULD be normalized to a single space.\n\nFor batch operations, if the individual operations are known to have the same operation name\nthen that operation name SHOULD be used prepended by `BATCH `,\notherwise `db.operation.name` SHOULD be `BATCH` or some other database\nsystem specific term if more applicable.\n",
                "requirement_level": "required",
                "stability": "stable",
                "type": "string",
            },
            {
                "brief": "Account or request [consistency level](https://learn.microsoft.com/azure/cosmos-db/consistency-levels).",
                "examples": [
                    "Eventual",
                    "ConsistentPrefix",
                    "BoundedStaleness",
                    "Strong",
                    "Session",
                ],
                "name": "azure.cosmosdb.consistency.level",
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
                "brief": "Cosmos DB sub status code.",
                "examples": [
                    1000,
                    1002,
                ],
                "name": "azure.cosmosdb.response.sub_status_code",
                "requirement_level": {
                    "conditionally_required": "when response was received and contained sub-code.",
                },
                "stability": "development",
                "type": "int",
            },
            {
                "brief": "Cosmos DB container name.\n",
                "examples": [
                    "public.users",
                    "customers",
                ],
                "name": "db.collection.name",
                "note": "It is RECOMMENDED to capture the value as provided by the application without attempting to do any case normalization.\n",
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
                "brief": "Database response status code.",
                "examples": [
                    "102",
                    "ORA-17002",
                    "08P01",
                    "404",
                ],
                "name": "db.response.status_code",
                "note": "The status code returned by the database. Usually it represents an error code, but may also represent partial success, warning, or differentiate between various types of successful outcomes.\nSemantic conventions for individual database systems SHOULD document what `db.response.status_code` means in the context of that system.\n",
                "requirement_level": {
                    "conditionally_required": "If the operation failed and status code is available.",
                },
                "stability": "stable",
                "type": "string",
            },
            {
                "brief": "Describes a class of error the operation ended with.\n",
                "examples": [
                    "timeout",
                    "java.net.UnknownHostException",
                    "server_certificate_invalid",
                    "500",
                ],
                "name": "error.type",
                "note": "The `error.type` SHOULD match the `db.response.status_code` returned by the database or the client library, or the canonical name of exception that occurred.\nWhen using canonical exception type name, instrumentation SHOULD do the best effort to report the most relevant type. For example, if the original exception is wrapped into a generic one, the original exception SHOULD be preferred.\nInstrumentations SHOULD document how `error.type` is populated.\n",
                "requirement_level": {
                    "conditionally_required": "If and only if the operation failed.",
                },
                "stability": "stable",
                "type": {
                    "allow_custom_values": none,
                    "members": [
                        {
                            "brief": "A fallback error value to be used when the instrumentation doesn't define a custom value.\n",
                            "deprecated": none,
                            "id": "other",
                            "note": none,
                            "stability": "stable",
                            "value": "_OTHER",
                        },
                    ],
                },
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
            {
                "brief": "List of regions contacted during operation in the order that they were contacted. If there is more than one region listed, it indicates that the operation was performed on multiple regions i.e. cross-regional call.\n",
                "examples": [
                    [
                        "North Central US",
                        "Australia East",
                        "Australia Southeast",
                    ],
                ],
                "name": "azure.cosmosdb.operation.contacted_regions",
                "note": "Region name matches the format of `displayName` in [Azure Location API](https://learn.microsoft.com/rest/api/subscription/subscriptions/list-locations?view=rest-subscription-2021-10-01&tabs=HTTP#location)\n",
                "requirement_level": {
                    "recommended": "If available",
                },
                "stability": "development",
                "type": "string[]",
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
                    "brief": "Account or request [consistency level](https://learn.microsoft.com/azure/cosmos-db/consistency-levels).",
                    "examples": [
                        "Eventual",
                        "ConsistentPrefix",
                        "BoundedStaleness",
                        "Strong",
                        "Session",
                    ],
                    "name": "azure.cosmosdb.consistency.level",
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
                    "brief": "Cosmos DB sub status code.",
                    "examples": [
                        1000,
                        1002,
                    ],
                    "name": "azure.cosmosdb.response.sub_status_code",
                    "requirement_level": {
                        "conditionally_required": "when response was received and contained sub-code.",
                    },
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "List of regions contacted during operation in the order that they were contacted. If there is more than one region listed, it indicates that the operation was performed on multiple regions i.e. cross-regional call.\n",
                    "examples": [
                        [
                            "North Central US",
                            "Australia East",
                            "Australia Southeast",
                        ],
                    ],
                    "name": "azure.cosmosdb.operation.contacted_regions",
                    "note": "Region name matches the format of `displayName` in [Azure Location API](https://learn.microsoft.com/rest/api/subscription/subscriptions/list-locations?view=rest-subscription-2021-10-01&tabs=HTTP#location)\n",
                    "requirement_level": {
                        "recommended": "If available",
                    },
                    "stability": "development",
                    "type": "string[]",
                },
                {
                    "brief": "Cosmos DB container name.\n",
                    "examples": [
                        "public.users",
                        "customers",
                    ],
                    "name": "db.collection.name",
                    "note": "It is RECOMMENDED to capture the value as provided by the application without attempting to do any case normalization.\n",
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
                    "requirement_level": "required",
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "Database response status code.",
                    "examples": [
                        "102",
                        "ORA-17002",
                        "08P01",
                        "404",
                    ],
                    "name": "db.response.status_code",
                    "note": "The status code returned by the database. Usually it represents an error code, but may also represent partial success, warning, or differentiate between various types of successful outcomes.\nSemantic conventions for individual database systems SHOULD document what `db.response.status_code` means in the context of that system.\n",
                    "requirement_level": {
                        "conditionally_required": "If the operation failed and status code is available.",
                    },
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "Describes a class of error the operation ended with.\n",
                    "examples": [
                        "timeout",
                        "java.net.UnknownHostException",
                        "server_certificate_invalid",
                        "500",
                    ],
                    "name": "error.type",
                    "note": "The `error.type` SHOULD match the `db.response.status_code` returned by the database or the client library, or the canonical name of exception that occurred.\nWhen using canonical exception type name, instrumentation SHOULD do the best effort to report the most relevant type. For example, if the original exception is wrapped into a generic one, the original exception SHOULD be preferred.\nInstrumentations SHOULD document how `error.type` is populated.\n",
                    "requirement_level": {
                        "conditionally_required": "If and only if the operation failed.",
                    },
                    "stability": "stable",
                    "type": {
                        "allow_custom_values": none,
                        "members": [
                            {
                                "brief": "A fallback error value to be used when the instrumentation doesn't define a custom value.\n",
                                "deprecated": none,
                                "id": "other",
                                "note": none,
                                "stability": "stable",
                                "value": "_OTHER",
                            },
                        ],
                    },
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
            "brief": "[Request units](https://learn.microsoft.com/azure/cosmos-db/request-units) consumed by the operation",
            "events": [],
            "id": "metric.azure.cosmosdb.client.operation.request_charge",
            "instrument": "histogram",
            "lineage": {
                "attributes": {
                    "azure.cosmosdb.consistency.level": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "requirement_level",
                        ],
                        "source_group": "registry.azure.cosmosdb",
                    },
                    "azure.cosmosdb.operation.contacted_regions": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "requirement_level",
                        ],
                        "source_group": "registry.azure.cosmosdb",
                    },
                    "azure.cosmosdb.response.sub_status_code": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "requirement_level",
                        ],
                        "source_group": "registry.azure.cosmosdb",
                    },
                    "db.collection.name": {
                        "inherited_fields": [
                            "examples",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "brief",
                            "note",
                            "requirement_level",
                        ],
                        "source_group": "registry.db",
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
                    "db.response.status_code": {
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
                    "error.type": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "note",
                            "requirement_level",
                        ],
                        "source_group": "registry.error",
                    },
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
            "metric_name": "azure.cosmosdb.client.operation.request_charge",
            "name": none,
            "root_namespace": "azure",
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
