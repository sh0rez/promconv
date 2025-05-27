package db

import (
	"github.com/prometheus/client_golang/prometheus"
)

import (
	"shorez.de/promconv/otel/error"
	"shorez.de/promconv/otel/network"
	"shorez.de/promconv/otel/server"
)

// Duration of database client operations.
type ClientOperationDuration struct {
	*prometheus.HistogramVec
}

func NewClientOperationDuration() ClientOperationDuration {
	labels := []string{"db_system_name", "db_collection_name", "db_namespace", "db_operation_name", "db_response_status_code", "error_type", "server_port", "db_query_summary", "db_stored_procedure_name", "network_peer_address", "network_peer_port", "server_address", "db_query_text"}
	return ClientOperationDuration{HistogramVec: prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "db",
		Name:      "client_operation_duration",
		Help:      "Duration of database client operations.",
	}, labels)}
}

func (m ClientOperationDuration) With(systemName AttrSystemName, extra interface {
	AttrDbCollectionName() AttrCollectionName
	AttrDbNamespace() AttrNamespace
	AttrDbOperationName() AttrOperationName
	AttrDbResponseStatusCode() AttrResponseStatusCode
	AttrErrorType() error.AttrType
	AttrServerPort() server.AttrPort
	AttrDbQuerySummary() AttrQuerySummary
	AttrDbStoredProcedureName() AttrStoredProcedureName
	AttrNetworkPeerAddress() network.AttrPeerAddress
	AttrNetworkPeerPort() network.AttrPeerPort
	AttrServerAddress() server.AttrAddress
	AttrDbQueryText() AttrQueryText
}) prometheus.Observer {
	if extra == nil {
		extra = ClientOperationDurationExtra{}
	}
	return m.WithLabelValues(
		string(systemName),
		string(extra.AttrDbCollectionName()),
		string(extra.AttrDbNamespace()),
		string(extra.AttrDbOperationName()),
		string(extra.AttrDbResponseStatusCode()),
		string(extra.AttrErrorType()),
		string(extra.AttrServerPort()),
		string(extra.AttrDbQuerySummary()),
		string(extra.AttrDbStoredProcedureName()),
		string(extra.AttrNetworkPeerAddress()),
		string(extra.AttrNetworkPeerPort()),
		string(extra.AttrServerAddress()),
		string(extra.AttrDbQueryText()),
	)
}

type ClientOperationDurationExtra struct {
	// The name of a collection (table, container) within the database.
	DbCollectionName AttrCollectionName `otel:"db.collection.name"`
	// The name of the database, fully qualified within the server address and port.
	DbNamespace AttrNamespace `otel:"db.namespace"`
	// The name of the operation or command being executed.
	DbOperationName AttrOperationName `otel:"db.operation.name"`
	// Database response status code.
	DbResponseStatusCode AttrResponseStatusCode `otel:"db.response.status_code"`
	// Describes a class of error the operation ended with.
	ErrorType error.AttrType `otel:"error.type"`
	// Server port number.
	ServerPort server.AttrPort `otel:"server.port"`
	// Low cardinality summary of a database query.
	DbQuerySummary AttrQuerySummary `otel:"db.query.summary"`
	// The name of a stored procedure within the database.
	DbStoredProcedureName AttrStoredProcedureName `otel:"db.stored_procedure.name"`
	// Peer address of the database node where the operation was performed.
	NetworkPeerAddress network.AttrPeerAddress `otel:"network.peer.address"`
	// Peer port number of the network connection.
	NetworkPeerPort network.AttrPeerPort `otel:"network.peer.port"`
	// Name of the database host.
	ServerAddress server.AttrAddress `otel:"server.address"`
	// The database query being executed.
	DbQueryText AttrQueryText `otel:"db.query.text"`
}

func (a ClientOperationDurationExtra) AttrDbCollectionName() AttrCollectionName {
	return a.DbCollectionName
}
func (a ClientOperationDurationExtra) AttrDbNamespace() AttrNamespace { return a.DbNamespace }
func (a ClientOperationDurationExtra) AttrDbOperationName() AttrOperationName {
	return a.DbOperationName
}
func (a ClientOperationDurationExtra) AttrDbResponseStatusCode() AttrResponseStatusCode {
	return a.DbResponseStatusCode
}
func (a ClientOperationDurationExtra) AttrErrorType() error.AttrType        { return a.ErrorType }
func (a ClientOperationDurationExtra) AttrServerPort() server.AttrPort      { return a.ServerPort }
func (a ClientOperationDurationExtra) AttrDbQuerySummary() AttrQuerySummary { return a.DbQuerySummary }
func (a ClientOperationDurationExtra) AttrDbStoredProcedureName() AttrStoredProcedureName {
	return a.DbStoredProcedureName
}
func (a ClientOperationDurationExtra) AttrNetworkPeerAddress() network.AttrPeerAddress {
	return a.NetworkPeerAddress
}
func (a ClientOperationDurationExtra) AttrNetworkPeerPort() network.AttrPeerPort {
	return a.NetworkPeerPort
}
func (a ClientOperationDurationExtra) AttrServerAddress() server.AttrAddress { return a.ServerAddress }
func (a ClientOperationDurationExtra) AttrDbQueryText() AttrQueryText        { return a.DbQueryText }

/*
State {
    name: "metric.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "ClientOperationDurationExtra",
        "Instr": "Histogram",
        "InstrMap": {
            "counter": "Counter",
            "gauge": "Gauge",
            "histogram": "Histogram",
            "updowncounter": "Gauge",
        },
        "Name": "client.operation.duration",
        "Type": "ClientOperationDuration",
        "attributes": [
            {
                "brief": "The database management system (DBMS) product as identified by the client instrumentation.",
                "name": "db.system.name",
                "note": "The actual DBMS may differ from the one identified by the client. For example, when using PostgreSQL client libraries to connect to a CockroachDB, the `db.system.name` is set to `postgresql` based on the instrumentation's best knowledge.\n",
                "requirement_level": "required",
                "stability": "stable",
                "type": {
                    "allow_custom_values": none,
                    "members": [
                        {
                            "brief": "Some other SQL database. Fallback only.",
                            "deprecated": none,
                            "id": "other_sql",
                            "note": none,
                            "stability": "development",
                            "value": "other_sql",
                        },
                        {
                            "brief": "[Adabas (Adaptable Database System)](https://documentation.softwareag.com/?pf=adabas)",
                            "deprecated": none,
                            "id": "softwareag.adabas",
                            "note": none,
                            "stability": "development",
                            "value": "softwareag.adabas",
                        },
                        {
                            "brief": "[Actian Ingres](https://www.actian.com/databases/ingres/)",
                            "deprecated": none,
                            "id": "actian.ingres",
                            "note": none,
                            "stability": "development",
                            "value": "actian.ingres",
                        },
                        {
                            "brief": "[Amazon DynamoDB](https://aws.amazon.com/pm/dynamodb/)",
                            "deprecated": none,
                            "id": "aws.dynamodb",
                            "note": none,
                            "stability": "development",
                            "value": "aws.dynamodb",
                        },
                        {
                            "brief": "[Amazon Redshift](https://aws.amazon.com/redshift/)",
                            "deprecated": none,
                            "id": "aws.redshift",
                            "note": none,
                            "stability": "development",
                            "value": "aws.redshift",
                        },
                        {
                            "brief": "[Azure Cosmos DB](https://learn.microsoft.com/azure/cosmos-db)",
                            "deprecated": none,
                            "id": "azure.cosmosdb",
                            "note": none,
                            "stability": "development",
                            "value": "azure.cosmosdb",
                        },
                        {
                            "brief": "[InterSystems Caché](https://www.intersystems.com/products/cache/)",
                            "deprecated": none,
                            "id": "intersystems.cache",
                            "note": none,
                            "stability": "development",
                            "value": "intersystems.cache",
                        },
                        {
                            "brief": "[Apache Cassandra](https://cassandra.apache.org/)",
                            "deprecated": none,
                            "id": "cassandra",
                            "note": none,
                            "stability": "development",
                            "value": "cassandra",
                        },
                        {
                            "brief": "[ClickHouse](https://clickhouse.com/)",
                            "deprecated": none,
                            "id": "clickhouse",
                            "note": none,
                            "stability": "development",
                            "value": "clickhouse",
                        },
                        {
                            "brief": "[CockroachDB](https://www.cockroachlabs.com/)",
                            "deprecated": none,
                            "id": "cockroachdb",
                            "note": none,
                            "stability": "development",
                            "value": "cockroachdb",
                        },
                        {
                            "brief": "[Couchbase](https://www.couchbase.com/)",
                            "deprecated": none,
                            "id": "couchbase",
                            "note": none,
                            "stability": "development",
                            "value": "couchbase",
                        },
                        {
                            "brief": "[Apache CouchDB](https://couchdb.apache.org/)",
                            "deprecated": none,
                            "id": "couchdb",
                            "note": none,
                            "stability": "development",
                            "value": "couchdb",
                        },
                        {
                            "brief": "[Apache Derby](https://db.apache.org/derby/)",
                            "deprecated": none,
                            "id": "derby",
                            "note": none,
                            "stability": "development",
                            "value": "derby",
                        },
                        {
                            "brief": "[Elasticsearch](https://www.elastic.co/elasticsearch)",
                            "deprecated": none,
                            "id": "elasticsearch",
                            "note": none,
                            "stability": "development",
                            "value": "elasticsearch",
                        },
                        {
                            "brief": "[Firebird](https://www.firebirdsql.org/)",
                            "deprecated": none,
                            "id": "firebirdsql",
                            "note": none,
                            "stability": "development",
                            "value": "firebirdsql",
                        },
                        {
                            "brief": "[Google Cloud Spanner](https://cloud.google.com/spanner)",
                            "deprecated": none,
                            "id": "gcp.spanner",
                            "note": none,
                            "stability": "development",
                            "value": "gcp.spanner",
                        },
                        {
                            "brief": "[Apache Geode](https://geode.apache.org/)",
                            "deprecated": none,
                            "id": "geode",
                            "note": none,
                            "stability": "development",
                            "value": "geode",
                        },
                        {
                            "brief": "[H2 Database](https://h2database.com/)",
                            "deprecated": none,
                            "id": "h2database",
                            "note": none,
                            "stability": "development",
                            "value": "h2database",
                        },
                        {
                            "brief": "[Apache HBase](https://hbase.apache.org/)",
                            "deprecated": none,
                            "id": "hbase",
                            "note": none,
                            "stability": "development",
                            "value": "hbase",
                        },
                        {
                            "brief": "[Apache Hive](https://hive.apache.org/)",
                            "deprecated": none,
                            "id": "hive",
                            "note": none,
                            "stability": "development",
                            "value": "hive",
                        },
                        {
                            "brief": "[HyperSQL Database](https://hsqldb.org/)",
                            "deprecated": none,
                            "id": "hsqldb",
                            "note": none,
                            "stability": "development",
                            "value": "hsqldb",
                        },
                        {
                            "brief": "[IBM Db2](https://www.ibm.com/db2)",
                            "deprecated": none,
                            "id": "ibm.db2",
                            "note": none,
                            "stability": "development",
                            "value": "ibm.db2",
                        },
                        {
                            "brief": "[IBM Informix](https://www.ibm.com/products/informix)",
                            "deprecated": none,
                            "id": "ibm.informix",
                            "note": none,
                            "stability": "development",
                            "value": "ibm.informix",
                        },
                        {
                            "brief": "[IBM Netezza](https://www.ibm.com/products/netezza)",
                            "deprecated": none,
                            "id": "ibm.netezza",
                            "note": none,
                            "stability": "development",
                            "value": "ibm.netezza",
                        },
                        {
                            "brief": "[InfluxDB](https://www.influxdata.com/)",
                            "deprecated": none,
                            "id": "influxdb",
                            "note": none,
                            "stability": "development",
                            "value": "influxdb",
                        },
                        {
                            "brief": "[Instant](https://www.instantdb.com/)",
                            "deprecated": none,
                            "id": "instantdb",
                            "note": none,
                            "stability": "development",
                            "value": "instantdb",
                        },
                        {
                            "brief": "[MariaDB](https://mariadb.org/)",
                            "deprecated": none,
                            "id": "mariadb",
                            "note": none,
                            "stability": "stable",
                            "value": "mariadb",
                        },
                        {
                            "brief": "[Memcached](https://memcached.org/)",
                            "deprecated": none,
                            "id": "memcached",
                            "note": none,
                            "stability": "development",
                            "value": "memcached",
                        },
                        {
                            "brief": "[MongoDB](https://www.mongodb.com/)",
                            "deprecated": none,
                            "id": "mongodb",
                            "note": none,
                            "stability": "development",
                            "value": "mongodb",
                        },
                        {
                            "brief": "[Microsoft SQL Server](https://www.microsoft.com/sql-server)",
                            "deprecated": none,
                            "id": "microsoft.sql_server",
                            "note": none,
                            "stability": "stable",
                            "value": "microsoft.sql_server",
                        },
                        {
                            "brief": "[MySQL](https://www.mysql.com/)",
                            "deprecated": none,
                            "id": "mysql",
                            "note": none,
                            "stability": "stable",
                            "value": "mysql",
                        },
                        {
                            "brief": "[Neo4j](https://neo4j.com/)",
                            "deprecated": none,
                            "id": "neo4j",
                            "note": none,
                            "stability": "development",
                            "value": "neo4j",
                        },
                        {
                            "brief": "[OpenSearch](https://opensearch.org/)",
                            "deprecated": none,
                            "id": "opensearch",
                            "note": none,
                            "stability": "development",
                            "value": "opensearch",
                        },
                        {
                            "brief": "[Oracle Database](https://www.oracle.com/database/)",
                            "deprecated": none,
                            "id": "oracle.db",
                            "note": none,
                            "stability": "development",
                            "value": "oracle.db",
                        },
                        {
                            "brief": "[PostgreSQL](https://www.postgresql.org/)",
                            "deprecated": none,
                            "id": "postgresql",
                            "note": none,
                            "stability": "stable",
                            "value": "postgresql",
                        },
                        {
                            "brief": "[Redis](https://redis.io/)",
                            "deprecated": none,
                            "id": "redis",
                            "note": none,
                            "stability": "development",
                            "value": "redis",
                        },
                        {
                            "brief": "[SAP HANA](https://www.sap.com/products/technology-platform/hana/what-is-sap-hana.html)",
                            "deprecated": none,
                            "id": "sap.hana",
                            "note": none,
                            "stability": "development",
                            "value": "sap.hana",
                        },
                        {
                            "brief": "[SAP MaxDB](https://maxdb.sap.com/)",
                            "deprecated": none,
                            "id": "sap.maxdb",
                            "note": none,
                            "stability": "development",
                            "value": "sap.maxdb",
                        },
                        {
                            "brief": "[SQLite](https://www.sqlite.org/)",
                            "deprecated": none,
                            "id": "sqlite",
                            "note": none,
                            "stability": "development",
                            "value": "sqlite",
                        },
                        {
                            "brief": "[Teradata](https://www.teradata.com/)",
                            "deprecated": none,
                            "id": "teradata",
                            "note": none,
                            "stability": "development",
                            "value": "teradata",
                        },
                        {
                            "brief": "[Trino](https://trino.io/)",
                            "deprecated": none,
                            "id": "trino",
                            "note": none,
                            "stability": "development",
                            "value": "trino",
                        },
                    ],
                },
            },
            {
                "brief": "The name of a collection (table, container) within the database.",
                "examples": [
                    "public.users",
                    "customers",
                ],
                "name": "db.collection.name",
                "note": "It is RECOMMENDED to capture the value as provided by the application\nwithout attempting to do any case normalization.\n\nThe collection name SHOULD NOT be extracted from `db.query.text`,\nwhen the database system supports query text with multiple collections\nin non-batch operations.\n\nFor batch operations, if the individual operations are known to have the same\ncollection name then that collection name SHOULD be used.\n",
                "requirement_level": {
                    "conditionally_required": "If readily available and if a database call is performed on a single collection.\n",
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
                "note": "If a database system has multiple namespace components, they SHOULD be concatenated from the most general to the most specific namespace component, using `|` as a separator between the components. Any missing components (and their associated separators) SHOULD be omitted.\nSemantic conventions for individual database systems SHOULD document what `db.namespace` means in the context of that system.\nIt is RECOMMENDED to capture the value as provided by the application without attempting to do any case normalization.\n",
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
                    "conditionally_required": "If readily available and if there is a single operation name that describes the database call.\n",
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
                "brief": "Low cardinality summary of a database query.\n",
                "examples": [
                    "SELECT wuser_table",
                    "INSERT shipping_details SELECT orders",
                    "get user by id",
                ],
                "name": "db.query.summary",
                "note": "The query summary describes a class of database queries and is useful\nas a grouping key, especially when analyzing telemetry for database\ncalls involving complex queries.\n\nSummary may be available to the instrumentation through\ninstrumentation hooks or other means. If it is not available, instrumentations\nthat support query parsing SHOULD generate a summary following\n[Generating query summary](/docs/database/database-spans.md#generating-a-summary-of-the-query)\nsection.\n",
                "requirement_level": {
                    "recommended": "if available through instrumentation hooks or if the instrumentation supports generating a query summary.",
                },
                "stability": "stable",
                "type": "string",
            },
            {
                "brief": "The name of a stored procedure within the database.",
                "examples": [
                    "GetCustomer",
                ],
                "name": "db.stored_procedure.name",
                "note": "It is RECOMMENDED to capture the value as provided by the application\nwithout attempting to do any case normalization.\n\nFor batch operations, if the individual operations are known to have the same\nstored procedure name then that stored procedure name SHOULD be used.\n",
                "requirement_level": {
                    "recommended": "If operation applies to a specific stored procedure.",
                },
                "stability": "stable",
                "type": "string",
            },
            {
                "brief": "Peer address of the database node where the operation was performed.",
                "examples": [
                    "10.1.2.80",
                    "/tmp/my.sock",
                ],
                "name": "network.peer.address",
                "note": "Semantic conventions for individual database systems SHOULD document whether `network.peer.*` attributes are applicable. Network peer address and port are useful when the application interacts with individual database nodes directly.\nIf a database operation involved multiple network calls (for example retries), the address of the last contacted node SHOULD be used.\n",
                "requirement_level": {
                    "recommended": "If applicable for this database system.",
                },
                "stability": "stable",
                "type": "string",
            },
            {
                "brief": "Peer port number of the network connection.",
                "examples": [
                    65123,
                ],
                "name": "network.peer.port",
                "requirement_level": {
                    "recommended": "If and only if `network.peer.address` is set.",
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
            {
                "brief": "The database query being executed.\n",
                "examples": [
                    "SELECT * FROM wuser_table where username = ?",
                    "SET mykey ?",
                ],
                "name": "db.query.text",
                "note": "For sanitization see [Sanitization of `db.query.text`](/docs/database/database-spans.md#sanitization-of-dbquerytext).\nFor batch operations, if the individual operations are known to have the same query text then that query text SHOULD be used, otherwise all of the individual query texts SHOULD be concatenated with separator `; ` or some other database system specific separator if more applicable.\nParameterized query text SHOULD NOT be sanitized. Even though parameterized query text can potentially have sensitive data, by using a parameterized query the user is giving a strong signal that any sensitive data will be passed as parameter values, and the benefit to observability of capturing the static part of the query text by default outweighs the risk.\n",
                "requirement_level": "opt_in",
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
                    "brief": "The name of a stored procedure within the database.",
                    "examples": [
                        "GetCustomer",
                    ],
                    "name": "db.stored_procedure.name",
                    "note": "It is RECOMMENDED to capture the value as provided by the application\nwithout attempting to do any case normalization.\n\nFor batch operations, if the individual operations are known to have the same\nstored procedure name then that stored procedure name SHOULD be used.\n",
                    "requirement_level": {
                        "recommended": "If operation applies to a specific stored procedure.",
                    },
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "Peer address of the database node where the operation was performed.",
                    "examples": [
                        "10.1.2.80",
                        "/tmp/my.sock",
                    ],
                    "name": "network.peer.address",
                    "note": "Semantic conventions for individual database systems SHOULD document whether `network.peer.*` attributes are applicable. Network peer address and port are useful when the application interacts with individual database nodes directly.\nIf a database operation involved multiple network calls (for example retries), the address of the last contacted node SHOULD be used.\n",
                    "requirement_level": {
                        "recommended": "If applicable for this database system.",
                    },
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "Low cardinality summary of a database query.\n",
                    "examples": [
                        "SELECT wuser_table",
                        "INSERT shipping_details SELECT orders",
                        "get user by id",
                    ],
                    "name": "db.query.summary",
                    "note": "The query summary describes a class of database queries and is useful\nas a grouping key, especially when analyzing telemetry for database\ncalls involving complex queries.\n\nSummary may be available to the instrumentation through\ninstrumentation hooks or other means. If it is not available, instrumentations\nthat support query parsing SHOULD generate a summary following\n[Generating query summary](/docs/database/database-spans.md#generating-a-summary-of-the-query)\nsection.\n",
                    "requirement_level": {
                        "recommended": "if available through instrumentation hooks or if the instrumentation supports generating a query summary.",
                    },
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "The name of a collection (table, container) within the database.",
                    "examples": [
                        "public.users",
                        "customers",
                    ],
                    "name": "db.collection.name",
                    "note": "It is RECOMMENDED to capture the value as provided by the application\nwithout attempting to do any case normalization.\n\nThe collection name SHOULD NOT be extracted from `db.query.text`,\nwhen the database system supports query text with multiple collections\nin non-batch operations.\n\nFor batch operations, if the individual operations are known to have the same\ncollection name then that collection name SHOULD be used.\n",
                    "requirement_level": {
                        "conditionally_required": "If readily available and if a database call is performed on a single collection.\n",
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
                        "conditionally_required": "If readily available and if there is a single operation name that describes the database call.\n",
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
                    "note": "If a database system has multiple namespace components, they SHOULD be concatenated from the most general to the most specific namespace component, using `|` as a separator between the components. Any missing components (and their associated separators) SHOULD be omitted.\nSemantic conventions for individual database systems SHOULD document what `db.namespace` means in the context of that system.\nIt is RECOMMENDED to capture the value as provided by the application without attempting to do any case normalization.\n",
                    "requirement_level": {
                        "conditionally_required": "If available.",
                    },
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "The database query being executed.\n",
                    "examples": [
                        "SELECT * FROM wuser_table where username = ?",
                        "SET mykey ?",
                    ],
                    "name": "db.query.text",
                    "note": "For sanitization see [Sanitization of `db.query.text`](/docs/database/database-spans.md#sanitization-of-dbquerytext).\nFor batch operations, if the individual operations are known to have the same query text then that query text SHOULD be used, otherwise all of the individual query texts SHOULD be concatenated with separator `; ` or some other database system specific separator if more applicable.\nParameterized query text SHOULD NOT be sanitized. Even though parameterized query text can potentially have sensitive data, by using a parameterized query the user is giving a strong signal that any sensitive data will be passed as parameter values, and the benefit to observability of capturing the static part of the query text by default outweighs the risk.\n",
                    "requirement_level": "opt_in",
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "The database management system (DBMS) product as identified by the client instrumentation.",
                    "name": "db.system.name",
                    "note": "The actual DBMS may differ from the one identified by the client. For example, when using PostgreSQL client libraries to connect to a CockroachDB, the `db.system.name` is set to `postgresql` based on the instrumentation's best knowledge.\n",
                    "requirement_level": "required",
                    "stability": "stable",
                    "type": {
                        "allow_custom_values": none,
                        "members": [
                            {
                                "brief": "Some other SQL database. Fallback only.",
                                "deprecated": none,
                                "id": "other_sql",
                                "note": none,
                                "stability": "development",
                                "value": "other_sql",
                            },
                            {
                                "brief": "[Adabas (Adaptable Database System)](https://documentation.softwareag.com/?pf=adabas)",
                                "deprecated": none,
                                "id": "softwareag.adabas",
                                "note": none,
                                "stability": "development",
                                "value": "softwareag.adabas",
                            },
                            {
                                "brief": "[Actian Ingres](https://www.actian.com/databases/ingres/)",
                                "deprecated": none,
                                "id": "actian.ingres",
                                "note": none,
                                "stability": "development",
                                "value": "actian.ingres",
                            },
                            {
                                "brief": "[Amazon DynamoDB](https://aws.amazon.com/pm/dynamodb/)",
                                "deprecated": none,
                                "id": "aws.dynamodb",
                                "note": none,
                                "stability": "development",
                                "value": "aws.dynamodb",
                            },
                            {
                                "brief": "[Amazon Redshift](https://aws.amazon.com/redshift/)",
                                "deprecated": none,
                                "id": "aws.redshift",
                                "note": none,
                                "stability": "development",
                                "value": "aws.redshift",
                            },
                            {
                                "brief": "[Azure Cosmos DB](https://learn.microsoft.com/azure/cosmos-db)",
                                "deprecated": none,
                                "id": "azure.cosmosdb",
                                "note": none,
                                "stability": "development",
                                "value": "azure.cosmosdb",
                            },
                            {
                                "brief": "[InterSystems Caché](https://www.intersystems.com/products/cache/)",
                                "deprecated": none,
                                "id": "intersystems.cache",
                                "note": none,
                                "stability": "development",
                                "value": "intersystems.cache",
                            },
                            {
                                "brief": "[Apache Cassandra](https://cassandra.apache.org/)",
                                "deprecated": none,
                                "id": "cassandra",
                                "note": none,
                                "stability": "development",
                                "value": "cassandra",
                            },
                            {
                                "brief": "[ClickHouse](https://clickhouse.com/)",
                                "deprecated": none,
                                "id": "clickhouse",
                                "note": none,
                                "stability": "development",
                                "value": "clickhouse",
                            },
                            {
                                "brief": "[CockroachDB](https://www.cockroachlabs.com/)",
                                "deprecated": none,
                                "id": "cockroachdb",
                                "note": none,
                                "stability": "development",
                                "value": "cockroachdb",
                            },
                            {
                                "brief": "[Couchbase](https://www.couchbase.com/)",
                                "deprecated": none,
                                "id": "couchbase",
                                "note": none,
                                "stability": "development",
                                "value": "couchbase",
                            },
                            {
                                "brief": "[Apache CouchDB](https://couchdb.apache.org/)",
                                "deprecated": none,
                                "id": "couchdb",
                                "note": none,
                                "stability": "development",
                                "value": "couchdb",
                            },
                            {
                                "brief": "[Apache Derby](https://db.apache.org/derby/)",
                                "deprecated": none,
                                "id": "derby",
                                "note": none,
                                "stability": "development",
                                "value": "derby",
                            },
                            {
                                "brief": "[Elasticsearch](https://www.elastic.co/elasticsearch)",
                                "deprecated": none,
                                "id": "elasticsearch",
                                "note": none,
                                "stability": "development",
                                "value": "elasticsearch",
                            },
                            {
                                "brief": "[Firebird](https://www.firebirdsql.org/)",
                                "deprecated": none,
                                "id": "firebirdsql",
                                "note": none,
                                "stability": "development",
                                "value": "firebirdsql",
                            },
                            {
                                "brief": "[Google Cloud Spanner](https://cloud.google.com/spanner)",
                                "deprecated": none,
                                "id": "gcp.spanner",
                                "note": none,
                                "stability": "development",
                                "value": "gcp.spanner",
                            },
                            {
                                "brief": "[Apache Geode](https://geode.apache.org/)",
                                "deprecated": none,
                                "id": "geode",
                                "note": none,
                                "stability": "development",
                                "value": "geode",
                            },
                            {
                                "brief": "[H2 Database](https://h2database.com/)",
                                "deprecated": none,
                                "id": "h2database",
                                "note": none,
                                "stability": "development",
                                "value": "h2database",
                            },
                            {
                                "brief": "[Apache HBase](https://hbase.apache.org/)",
                                "deprecated": none,
                                "id": "hbase",
                                "note": none,
                                "stability": "development",
                                "value": "hbase",
                            },
                            {
                                "brief": "[Apache Hive](https://hive.apache.org/)",
                                "deprecated": none,
                                "id": "hive",
                                "note": none,
                                "stability": "development",
                                "value": "hive",
                            },
                            {
                                "brief": "[HyperSQL Database](https://hsqldb.org/)",
                                "deprecated": none,
                                "id": "hsqldb",
                                "note": none,
                                "stability": "development",
                                "value": "hsqldb",
                            },
                            {
                                "brief": "[IBM Db2](https://www.ibm.com/db2)",
                                "deprecated": none,
                                "id": "ibm.db2",
                                "note": none,
                                "stability": "development",
                                "value": "ibm.db2",
                            },
                            {
                                "brief": "[IBM Informix](https://www.ibm.com/products/informix)",
                                "deprecated": none,
                                "id": "ibm.informix",
                                "note": none,
                                "stability": "development",
                                "value": "ibm.informix",
                            },
                            {
                                "brief": "[IBM Netezza](https://www.ibm.com/products/netezza)",
                                "deprecated": none,
                                "id": "ibm.netezza",
                                "note": none,
                                "stability": "development",
                                "value": "ibm.netezza",
                            },
                            {
                                "brief": "[InfluxDB](https://www.influxdata.com/)",
                                "deprecated": none,
                                "id": "influxdb",
                                "note": none,
                                "stability": "development",
                                "value": "influxdb",
                            },
                            {
                                "brief": "[Instant](https://www.instantdb.com/)",
                                "deprecated": none,
                                "id": "instantdb",
                                "note": none,
                                "stability": "development",
                                "value": "instantdb",
                            },
                            {
                                "brief": "[MariaDB](https://mariadb.org/)",
                                "deprecated": none,
                                "id": "mariadb",
                                "note": none,
                                "stability": "stable",
                                "value": "mariadb",
                            },
                            {
                                "brief": "[Memcached](https://memcached.org/)",
                                "deprecated": none,
                                "id": "memcached",
                                "note": none,
                                "stability": "development",
                                "value": "memcached",
                            },
                            {
                                "brief": "[MongoDB](https://www.mongodb.com/)",
                                "deprecated": none,
                                "id": "mongodb",
                                "note": none,
                                "stability": "development",
                                "value": "mongodb",
                            },
                            {
                                "brief": "[Microsoft SQL Server](https://www.microsoft.com/sql-server)",
                                "deprecated": none,
                                "id": "microsoft.sql_server",
                                "note": none,
                                "stability": "stable",
                                "value": "microsoft.sql_server",
                            },
                            {
                                "brief": "[MySQL](https://www.mysql.com/)",
                                "deprecated": none,
                                "id": "mysql",
                                "note": none,
                                "stability": "stable",
                                "value": "mysql",
                            },
                            {
                                "brief": "[Neo4j](https://neo4j.com/)",
                                "deprecated": none,
                                "id": "neo4j",
                                "note": none,
                                "stability": "development",
                                "value": "neo4j",
                            },
                            {
                                "brief": "[OpenSearch](https://opensearch.org/)",
                                "deprecated": none,
                                "id": "opensearch",
                                "note": none,
                                "stability": "development",
                                "value": "opensearch",
                            },
                            {
                                "brief": "[Oracle Database](https://www.oracle.com/database/)",
                                "deprecated": none,
                                "id": "oracle.db",
                                "note": none,
                                "stability": "development",
                                "value": "oracle.db",
                            },
                            {
                                "brief": "[PostgreSQL](https://www.postgresql.org/)",
                                "deprecated": none,
                                "id": "postgresql",
                                "note": none,
                                "stability": "stable",
                                "value": "postgresql",
                            },
                            {
                                "brief": "[Redis](https://redis.io/)",
                                "deprecated": none,
                                "id": "redis",
                                "note": none,
                                "stability": "development",
                                "value": "redis",
                            },
                            {
                                "brief": "[SAP HANA](https://www.sap.com/products/technology-platform/hana/what-is-sap-hana.html)",
                                "deprecated": none,
                                "id": "sap.hana",
                                "note": none,
                                "stability": "development",
                                "value": "sap.hana",
                            },
                            {
                                "brief": "[SAP MaxDB](https://maxdb.sap.com/)",
                                "deprecated": none,
                                "id": "sap.maxdb",
                                "note": none,
                                "stability": "development",
                                "value": "sap.maxdb",
                            },
                            {
                                "brief": "[SQLite](https://www.sqlite.org/)",
                                "deprecated": none,
                                "id": "sqlite",
                                "note": none,
                                "stability": "development",
                                "value": "sqlite",
                            },
                            {
                                "brief": "[Teradata](https://www.teradata.com/)",
                                "deprecated": none,
                                "id": "teradata",
                                "note": none,
                                "stability": "development",
                                "value": "teradata",
                            },
                            {
                                "brief": "[Trino](https://trino.io/)",
                                "deprecated": none,
                                "id": "trino",
                                "note": none,
                                "stability": "development",
                                "value": "trino",
                            },
                        ],
                    },
                },
                {
                    "brief": "Peer port number of the network connection.",
                    "examples": [
                        65123,
                    ],
                    "name": "network.peer.port",
                    "requirement_level": {
                        "recommended": "If and only if `network.peer.address` is set.",
                    },
                    "stability": "stable",
                    "type": "int",
                },
            ],
            "brief": "Duration of database client operations.",
            "events": [],
            "id": "metric.db.client.operation.duration",
            "instrument": "histogram",
            "lineage": {
                "attributes": {
                    "db.collection.name": {
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
                    "db.namespace": {
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
                    "db.query.summary": {
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
                    "db.query.text": {
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
                    "db.stored_procedure.name": {
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
                    "db.system.name": {
                        "inherited_fields": [
                            "brief",
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
                    "network.peer.address": {
                        "inherited_fields": [
                            "examples",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "brief",
                            "note",
                            "requirement_level",
                        ],
                        "source_group": "registry.network",
                    },
                    "network.peer.port": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "requirement_level",
                        ],
                        "source_group": "registry.network",
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
                    "path": "https://github.com/open-telemetry/semantic-conventions.git[model]/database/metrics.yaml",
                    "registry_id": "main",
                },
            },
            "metric_name": "db.client.operation.duration",
            "name": none,
            "note": "Batch operations SHOULD be recorded as a single operation.\n",
            "root_namespace": "db",
            "span_kind": none,
            "stability": "stable",
            "type": "metric",
            "unit": "s",
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
