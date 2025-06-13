package db

// Deprecated, use `cassandra.consistency.level` instead
type AttrCassandraConsistencyLevel string // db.cassandra.consistency_level

func (AttrCassandraConsistencyLevel) Development()    {}
func (AttrCassandraConsistencyLevel) Recommended()    {}
func (AttrCassandraConsistencyLevel) Key() string     { return "db_cassandra_consistency_level" }
func (a AttrCassandraConsistencyLevel) Value() string { return string(a) }

const CassandraConsistencyLevelAll AttrCassandraConsistencyLevel = "all"
const CassandraConsistencyLevelEachQuorum AttrCassandraConsistencyLevel = "each_quorum"
const CassandraConsistencyLevelQuorum AttrCassandraConsistencyLevel = "quorum"
const CassandraConsistencyLevelLocalQuorum AttrCassandraConsistencyLevel = "local_quorum"
const CassandraConsistencyLevelOne AttrCassandraConsistencyLevel = "one"
const CassandraConsistencyLevelTwo AttrCassandraConsistencyLevel = "two"
const CassandraConsistencyLevelThree AttrCassandraConsistencyLevel = "three"
const CassandraConsistencyLevelLocalOne AttrCassandraConsistencyLevel = "local_one"
const CassandraConsistencyLevelAny AttrCassandraConsistencyLevel = "any"
const CassandraConsistencyLevelSerial AttrCassandraConsistencyLevel = "serial"
const CassandraConsistencyLevelLocalSerial AttrCassandraConsistencyLevel = "local_serial"

// Deprecated, use `cassandra.coordinator.dc` instead
type AttrCassandraCoordinatorDc string // db.cassandra.coordinator.dc

func (AttrCassandraCoordinatorDc) Development()    {}
func (AttrCassandraCoordinatorDc) Recommended()    {}
func (AttrCassandraCoordinatorDc) Key() string     { return "db_cassandra_coordinator_dc" }
func (a AttrCassandraCoordinatorDc) Value() string { return string(a) }

// Deprecated, use `cassandra.coordinator.id` instead
type AttrCassandraCoordinatorId string // db.cassandra.coordinator.id

func (AttrCassandraCoordinatorId) Development()    {}
func (AttrCassandraCoordinatorId) Recommended()    {}
func (AttrCassandraCoordinatorId) Key() string     { return "db_cassandra_coordinator_id" }
func (a AttrCassandraCoordinatorId) Value() string { return string(a) }

// Deprecated, use `cassandra.query.idempotent` instead
type AttrCassandraIdempotence string // db.cassandra.idempotence

func (AttrCassandraIdempotence) Development()    {}
func (AttrCassandraIdempotence) Recommended()    {}
func (AttrCassandraIdempotence) Key() string     { return "db_cassandra_idempotence" }
func (a AttrCassandraIdempotence) Value() string { return string(a) }

// Deprecated, use `cassandra.page.size` instead
type AttrCassandraPageSize string // db.cassandra.page_size

func (AttrCassandraPageSize) Development()    {}
func (AttrCassandraPageSize) Recommended()    {}
func (AttrCassandraPageSize) Key() string     { return "db_cassandra_page_size" }
func (a AttrCassandraPageSize) Value() string { return string(a) }

// Deprecated, use `cassandra.speculative_execution.count` instead
type AttrCassandraSpeculativeExecutionCount string // db.cassandra.speculative_execution_count

func (AttrCassandraSpeculativeExecutionCount) Development() {}
func (AttrCassandraSpeculativeExecutionCount) Recommended() {}
func (AttrCassandraSpeculativeExecutionCount) Key() string {
	return "db_cassandra_speculative_execution_count"
}
func (a AttrCassandraSpeculativeExecutionCount) Value() string { return string(a) }

// Deprecated, use `db.collection.name` instead
type AttrCassandraTable string // db.cassandra.table

func (AttrCassandraTable) Development()    {}
func (AttrCassandraTable) Recommended()    {}
func (AttrCassandraTable) Key() string     { return "db_cassandra_table" }
func (a AttrCassandraTable) Value() string { return string(a) }

// The name of the connection pool; unique within the instrumented application. In case the connection pool implementation doesn't provide a name, instrumentation SHOULD use a combination of parameters that would make the name unique, for example, combining attributes `server.address`, `server.port`, and `db.namespace`, formatted as `server.address:server.port/db.namespace`. Instrumentations that generate connection pool name following different patterns SHOULD document it
type AttrClientConnectionPoolName string // db.client.connection.pool.name

func (AttrClientConnectionPoolName) Development()    {}
func (AttrClientConnectionPoolName) Recommended()    {}
func (AttrClientConnectionPoolName) Key() string     { return "db_client_connection_pool_name" }
func (a AttrClientConnectionPoolName) Value() string { return string(a) }

// The state of a connection in the pool
type AttrClientConnectionState string // db.client.connection.state

func (AttrClientConnectionState) Development()    {}
func (AttrClientConnectionState) Recommended()    {}
func (AttrClientConnectionState) Key() string     { return "db_client_connection_state" }
func (a AttrClientConnectionState) Value() string { return string(a) }

const ClientConnectionStateIdle AttrClientConnectionState = "idle"
const ClientConnectionStateUsed AttrClientConnectionState = "used"

// Deprecated, use `db.client.connection.pool.name` instead
type AttrClientConnectionsPoolName string // db.client.connections.pool.name

func (AttrClientConnectionsPoolName) Development()    {}
func (AttrClientConnectionsPoolName) Recommended()    {}
func (AttrClientConnectionsPoolName) Key() string     { return "db_client_connections_pool_name" }
func (a AttrClientConnectionsPoolName) Value() string { return string(a) }

// Deprecated, use `db.client.connection.state` instead
type AttrClientConnectionsState string // db.client.connections.state

func (AttrClientConnectionsState) Development()    {}
func (AttrClientConnectionsState) Recommended()    {}
func (AttrClientConnectionsState) Key() string     { return "db_client_connections_state" }
func (a AttrClientConnectionsState) Value() string { return string(a) }

const ClientConnectionsStateIdle AttrClientConnectionsState = "idle"
const ClientConnectionsStateUsed AttrClientConnectionsState = "used"

// The name of a collection (table, container) within the database.
// It is RECOMMENDED to capture the value as provided by the application
// without attempting to do any case normalization.
//
// The collection name SHOULD NOT be extracted from `db.query.text`,
// when the database system supports query text with multiple collections
// in non-batch operations.
//
// For batch operations, if the individual operations are known to have the same
// collection name then that collection name SHOULD be used
type AttrCollectionName string // db.collection.name

func (AttrCollectionName) Stable()         {}
func (AttrCollectionName) Recommended()    {}
func (AttrCollectionName) Key() string     { return "db_collection_name" }
func (a AttrCollectionName) Value() string { return string(a) }

// Deprecated, use `server.address`, `server.port` attributes instead
type AttrConnectionString string // db.connection_string

func (AttrConnectionString) Development()    {}
func (AttrConnectionString) Recommended()    {}
func (AttrConnectionString) Key() string     { return "db_connection_string" }
func (a AttrConnectionString) Value() string { return string(a) }

// Deprecated, use `azure.client.id` instead
type AttrCosmosdbClientId string // db.cosmosdb.client_id

func (AttrCosmosdbClientId) Development()    {}
func (AttrCosmosdbClientId) Recommended()    {}
func (AttrCosmosdbClientId) Key() string     { return "db_cosmosdb_client_id" }
func (a AttrCosmosdbClientId) Value() string { return string(a) }

// Deprecated, use `azure.cosmosdb.connection.mode` instead
type AttrCosmosdbConnectionMode string // db.cosmosdb.connection_mode

func (AttrCosmosdbConnectionMode) Development()    {}
func (AttrCosmosdbConnectionMode) Recommended()    {}
func (AttrCosmosdbConnectionMode) Key() string     { return "db_cosmosdb_connection_mode" }
func (a AttrCosmosdbConnectionMode) Value() string { return string(a) }

const CosmosdbConnectionModeGateway AttrCosmosdbConnectionMode = "gateway"
const CosmosdbConnectionModeDirect AttrCosmosdbConnectionMode = "direct"

// Deprecated, use `cosmosdb.consistency.level` instead
type AttrCosmosdbConsistencyLevel string // db.cosmosdb.consistency_level

func (AttrCosmosdbConsistencyLevel) Development()    {}
func (AttrCosmosdbConsistencyLevel) Recommended()    {}
func (AttrCosmosdbConsistencyLevel) Key() string     { return "db_cosmosdb_consistency_level" }
func (a AttrCosmosdbConsistencyLevel) Value() string { return string(a) }

const CosmosdbConsistencyLevelStrong AttrCosmosdbConsistencyLevel = "Strong"
const CosmosdbConsistencyLevelBoundedStaleness AttrCosmosdbConsistencyLevel = "BoundedStaleness"
const CosmosdbConsistencyLevelSession AttrCosmosdbConsistencyLevel = "Session"
const CosmosdbConsistencyLevelEventual AttrCosmosdbConsistencyLevel = "Eventual"
const CosmosdbConsistencyLevelConsistentPrefix AttrCosmosdbConsistencyLevel = "ConsistentPrefix"

// Deprecated, use `db.collection.name` instead
type AttrCosmosdbContainer string // db.cosmosdb.container

func (AttrCosmosdbContainer) Development()    {}
func (AttrCosmosdbContainer) Recommended()    {}
func (AttrCosmosdbContainer) Key() string     { return "db_cosmosdb_container" }
func (a AttrCosmosdbContainer) Value() string { return string(a) }

// Deprecated, no replacement at this time
type AttrCosmosdbOperationType string // db.cosmosdb.operation_type

func (AttrCosmosdbOperationType) Development()    {}
func (AttrCosmosdbOperationType) Recommended()    {}
func (AttrCosmosdbOperationType) Key() string     { return "db_cosmosdb_operation_type" }
func (a AttrCosmosdbOperationType) Value() string { return string(a) }

const CosmosdbOperationTypeBatch AttrCosmosdbOperationType = "batch"
const CosmosdbOperationTypeCreate AttrCosmosdbOperationType = "create"
const CosmosdbOperationTypeDelete AttrCosmosdbOperationType = "delete"
const CosmosdbOperationTypeExecute AttrCosmosdbOperationType = "execute"
const CosmosdbOperationTypeExecuteJavascript AttrCosmosdbOperationType = "execute_javascript"
const CosmosdbOperationTypeInvalid AttrCosmosdbOperationType = "invalid"
const CosmosdbOperationTypeHead AttrCosmosdbOperationType = "head"
const CosmosdbOperationTypeHeadFeed AttrCosmosdbOperationType = "head_feed"
const CosmosdbOperationTypePatch AttrCosmosdbOperationType = "patch"
const CosmosdbOperationTypeQuery AttrCosmosdbOperationType = "query"
const CosmosdbOperationTypeQueryPlan AttrCosmosdbOperationType = "query_plan"
const CosmosdbOperationTypeRead AttrCosmosdbOperationType = "read"
const CosmosdbOperationTypeReadFeed AttrCosmosdbOperationType = "read_feed"
const CosmosdbOperationTypeReplace AttrCosmosdbOperationType = "replace"
const CosmosdbOperationTypeUpsert AttrCosmosdbOperationType = "upsert"

// Deprecated, use `azure.cosmosdb.operation.contacted_regions` instead
type AttrCosmosdbRegionsContacted string // db.cosmosdb.regions_contacted

func (AttrCosmosdbRegionsContacted) Development()    {}
func (AttrCosmosdbRegionsContacted) Recommended()    {}
func (AttrCosmosdbRegionsContacted) Key() string     { return "db_cosmosdb_regions_contacted" }
func (a AttrCosmosdbRegionsContacted) Value() string { return string(a) }

// Deprecated, use `azure.cosmosdb.operation.request_charge` instead
type AttrCosmosdbRequestCharge string // db.cosmosdb.request_charge

func (AttrCosmosdbRequestCharge) Development()    {}
func (AttrCosmosdbRequestCharge) Recommended()    {}
func (AttrCosmosdbRequestCharge) Key() string     { return "db_cosmosdb_request_charge" }
func (a AttrCosmosdbRequestCharge) Value() string { return string(a) }

// Deprecated, use `azure.cosmosdb.request.body.size` instead
type AttrCosmosdbRequestContentLength string // db.cosmosdb.request_content_length

func (AttrCosmosdbRequestContentLength) Development()    {}
func (AttrCosmosdbRequestContentLength) Recommended()    {}
func (AttrCosmosdbRequestContentLength) Key() string     { return "db_cosmosdb_request_content_length" }
func (a AttrCosmosdbRequestContentLength) Value() string { return string(a) }

// Deprecated, use `db.response.status_code` instead
type AttrCosmosdbStatusCode string // db.cosmosdb.status_code

func (AttrCosmosdbStatusCode) Development()    {}
func (AttrCosmosdbStatusCode) Recommended()    {}
func (AttrCosmosdbStatusCode) Key() string     { return "db_cosmosdb_status_code" }
func (a AttrCosmosdbStatusCode) Value() string { return string(a) }

// Deprecated, use `azure.cosmosdb.response.sub_status_code` instead
type AttrCosmosdbSubStatusCode string // db.cosmosdb.sub_status_code

func (AttrCosmosdbSubStatusCode) Development()    {}
func (AttrCosmosdbSubStatusCode) Recommended()    {}
func (AttrCosmosdbSubStatusCode) Key() string     { return "db_cosmosdb_sub_status_code" }
func (a AttrCosmosdbSubStatusCode) Value() string { return string(a) }

// Deprecated, use `db.namespace` instead
type AttrElasticsearchClusterName string // db.elasticsearch.cluster.name

func (AttrElasticsearchClusterName) Development()    {}
func (AttrElasticsearchClusterName) Recommended()    {}
func (AttrElasticsearchClusterName) Key() string     { return "db_elasticsearch_cluster_name" }
func (a AttrElasticsearchClusterName) Value() string { return string(a) }

// Deprecated, use `elasticsearch.node.name` instead
type AttrElasticsearchNodeName string // db.elasticsearch.node.name

func (AttrElasticsearchNodeName) Development()    {}
func (AttrElasticsearchNodeName) Recommended()    {}
func (AttrElasticsearchNodeName) Key() string     { return "db_elasticsearch_node_name" }
func (a AttrElasticsearchNodeName) Value() string { return string(a) }

// Deprecated, use `db.operation.parameter` instead
type AttrElasticsearchPathParts string // db.elasticsearch.path_parts

func (AttrElasticsearchPathParts) Development()    {}
func (AttrElasticsearchPathParts) Recommended()    {}
func (AttrElasticsearchPathParts) Key() string     { return "db_elasticsearch_path_parts" }
func (a AttrElasticsearchPathParts) Value() string { return string(a) }

// Deprecated, no general replacement at this time. For Elasticsearch, use `db.elasticsearch.node.name` instead
type AttrInstanceId string // db.instance.id

func (AttrInstanceId) Development()    {}
func (AttrInstanceId) Recommended()    {}
func (AttrInstanceId) Key() string     { return "db_instance_id" }
func (a AttrInstanceId) Value() string { return string(a) }

// Removed, no replacement at this time
type AttrJdbcDriverClassname string // db.jdbc.driver_classname

func (AttrJdbcDriverClassname) Development()    {}
func (AttrJdbcDriverClassname) Recommended()    {}
func (AttrJdbcDriverClassname) Key() string     { return "db_jdbc_driver_classname" }
func (a AttrJdbcDriverClassname) Value() string { return string(a) }

// Deprecated, use `db.collection.name` instead
type AttrMongodbCollection string // db.mongodb.collection

func (AttrMongodbCollection) Development()    {}
func (AttrMongodbCollection) Recommended()    {}
func (AttrMongodbCollection) Key() string     { return "db_mongodb_collection" }
func (a AttrMongodbCollection) Value() string { return string(a) }

// Deprecated, SQL Server instance is now populated as a part of `db.namespace` attribute
type AttrMssqlInstanceName string // db.mssql.instance_name

func (AttrMssqlInstanceName) Development()    {}
func (AttrMssqlInstanceName) Recommended()    {}
func (AttrMssqlInstanceName) Key() string     { return "db_mssql_instance_name" }
func (a AttrMssqlInstanceName) Value() string { return string(a) }

// Deprecated, use `db.namespace` instead
type AttrName string // db.name

func (AttrName) Development()    {}
func (AttrName) Recommended()    {}
func (AttrName) Key() string     { return "db_name" }
func (a AttrName) Value() string { return string(a) }

// The name of the database, fully qualified within the server address and port.
//
// If a database system has multiple namespace components, they SHOULD be concatenated from the most general to the most specific namespace component, using `|` as a separator between the components. Any missing components (and their associated separators) SHOULD be omitted.
// Semantic conventions for individual database systems SHOULD document what `db.namespace` means in the context of that system.
// It is RECOMMENDED to capture the value as provided by the application without attempting to do any case normalization
type AttrNamespace string // db.namespace

func (AttrNamespace) Stable()         {}
func (AttrNamespace) Recommended()    {}
func (AttrNamespace) Key() string     { return "db_namespace" }
func (a AttrNamespace) Value() string { return string(a) }

// Deprecated, use `db.operation.name` instead
type AttrOperation string // db.operation

func (AttrOperation) Development()    {}
func (AttrOperation) Recommended()    {}
func (AttrOperation) Key() string     { return "db_operation" }
func (a AttrOperation) Value() string { return string(a) }

// The number of queries included in a batch operation.
// Operations are only considered batches when they contain two or more operations, and so `db.operation.batch.size` SHOULD never be `1`
type AttrOperationBatchSize string // db.operation.batch.size

func (AttrOperationBatchSize) Stable()         {}
func (AttrOperationBatchSize) Recommended()    {}
func (AttrOperationBatchSize) Key() string     { return "db_operation_batch_size" }
func (a AttrOperationBatchSize) Value() string { return string(a) }

// The name of the operation or command being executed.
//
// It is RECOMMENDED to capture the value as provided by the application
// without attempting to do any case normalization.
//
// The operation name SHOULD NOT be extracted from `db.query.text`,
// when the database system supports query text with multiple operations
// in non-batch operations.
//
// If spaces can occur in the operation name, multiple consecutive spaces
// SHOULD be normalized to a single space.
//
// For batch operations, if the individual operations are known to have the same operation name
// then that operation name SHOULD be used prepended by `BATCH `,
// otherwise `db.operation.name` SHOULD be `BATCH` or some other database
// system specific term if more applicable
type AttrOperationName string // db.operation.name

func (AttrOperationName) Stable()         {}
func (AttrOperationName) Recommended()    {}
func (AttrOperationName) Key() string     { return "db_operation_name" }
func (a AttrOperationName) Value() string { return string(a) }

// A database operation parameter, with `<key>` being the parameter name, and the attribute value being a string representation of the parameter value.
//
// For example, a client-side maximum number of rows to read from the database
// MAY be recorded as the `db.operation.parameter.max_rows` attribute.
//
// `db.query.text` parameters SHOULD be captured using `db.query.parameter.<key>`
// instead of `db.operation.parameter.<key>`
type AttrOperationParameter string // db.operation.parameter

func (AttrOperationParameter) Development()    {}
func (AttrOperationParameter) Recommended()    {}
func (AttrOperationParameter) Key() string     { return "db_operation_parameter" }
func (a AttrOperationParameter) Value() string { return string(a) }

// A database query parameter, with `<key>` being the parameter name, and the attribute value being a string representation of the parameter value.
//
// If a query parameter has no name and instead is referenced only by index,
// then `<key>` SHOULD be the 0-based index.
//
// `db.query.parameter.<key>` SHOULD match
// up with the parameterized placeholders present in `db.query.text`.
//
// `db.query.parameter.<key>` SHOULD NOT be captured on batch operations.
//
// Examples:
//
//   - For a query `SELECT * FROM users where username =  %s` with the parameter `"jdoe"`,
//     the attribute `db.query.parameter.0` SHOULD be set to `"jdoe"`.
//   - For a query `"SELECT * FROM users WHERE username = %(username)s;` with parameter
//     `username = "jdoe"`, the attribute `db.query.parameter.username` SHOULD be set to `"jdoe"`
type AttrQueryParameter string // db.query.parameter

func (AttrQueryParameter) Development()    {}
func (AttrQueryParameter) Recommended()    {}
func (AttrQueryParameter) Key() string     { return "db_query_parameter" }
func (a AttrQueryParameter) Value() string { return string(a) }

// Low cardinality summary of a database query.
//
// The query summary describes a class of database queries and is useful
// as a grouping key, especially when analyzing telemetry for database
// calls involving complex queries.
//
// Summary may be available to the instrumentation through
// instrumentation hooks or other means. If it is not available, instrumentations
// that support query parsing SHOULD generate a summary following
// [Generating query summary]
// section
//
// [Generating query summary]: /docs/database/database-spans.md#generating-a-summary-of-the-query
type AttrQuerySummary string // db.query.summary

func (AttrQuerySummary) Stable()         {}
func (AttrQuerySummary) Recommended()    {}
func (AttrQuerySummary) Key() string     { return "db_query_summary" }
func (a AttrQuerySummary) Value() string { return string(a) }

// The database query being executed.
//
// For sanitization see [Sanitization of `db.query.text`].
// For batch operations, if the individual operations are known to have the same query text then that query text SHOULD be used, otherwise all of the individual query texts SHOULD be concatenated with separator `; ` or some other database system specific separator if more applicable.
// Parameterized query text SHOULD NOT be sanitized. Even though parameterized query text can potentially have sensitive data, by using a parameterized query the user is giving a strong signal that any sensitive data will be passed as parameter values, and the benefit to observability of capturing the static part of the query text by default outweighs the risk
//
// [Sanitization of `db.query.text`]: /docs/database/database-spans.md#sanitization-of-dbquerytext
type AttrQueryText string // db.query.text

func (AttrQueryText) Stable()         {}
func (AttrQueryText) Recommended()    {}
func (AttrQueryText) Key() string     { return "db_query_text" }
func (a AttrQueryText) Value() string { return string(a) }

// Deprecated, use `db.namespace` instead
type AttrRedisDatabaseIndex string // db.redis.database_index

func (AttrRedisDatabaseIndex) Development()    {}
func (AttrRedisDatabaseIndex) Recommended()    {}
func (AttrRedisDatabaseIndex) Key() string     { return "db_redis_database_index" }
func (a AttrRedisDatabaseIndex) Value() string { return string(a) }

// Number of rows returned by the operation
type AttrResponseReturnedRows string // db.response.returned_rows

func (AttrResponseReturnedRows) Development()    {}
func (AttrResponseReturnedRows) Recommended()    {}
func (AttrResponseReturnedRows) Key() string     { return "db_response_returned_rows" }
func (a AttrResponseReturnedRows) Value() string { return string(a) }

// Database response status code.
// The status code returned by the database. Usually it represents an error code, but may also represent partial success, warning, or differentiate between various types of successful outcomes.
// Semantic conventions for individual database systems SHOULD document what `db.response.status_code` means in the context of that system
type AttrResponseStatusCode string // db.response.status_code

func (AttrResponseStatusCode) Stable()         {}
func (AttrResponseStatusCode) Recommended()    {}
func (AttrResponseStatusCode) Key() string     { return "db_response_status_code" }
func (a AttrResponseStatusCode) Value() string { return string(a) }

// Deprecated, use `db.collection.name` instead
type AttrSqlTable string // db.sql.table

func (AttrSqlTable) Development()    {}
func (AttrSqlTable) Recommended()    {}
func (AttrSqlTable) Key() string     { return "db_sql_table" }
func (a AttrSqlTable) Value() string { return string(a) }

// The database statement being executed
type AttrStatement string // db.statement

func (AttrStatement) Development()    {}
func (AttrStatement) Recommended()    {}
func (AttrStatement) Key() string     { return "db_statement" }
func (a AttrStatement) Value() string { return string(a) }

// The name of a stored procedure within the database.
// It is RECOMMENDED to capture the value as provided by the application
// without attempting to do any case normalization.
//
// For batch operations, if the individual operations are known to have the same
// stored procedure name then that stored procedure name SHOULD be used
type AttrStoredProcedureName string // db.stored_procedure.name

func (AttrStoredProcedureName) Stable()         {}
func (AttrStoredProcedureName) Recommended()    {}
func (AttrStoredProcedureName) Key() string     { return "db_stored_procedure_name" }
func (a AttrStoredProcedureName) Value() string { return string(a) }

// Deprecated, use `db.system.name` instead
type AttrSystem string // db.system

func (AttrSystem) Development()    {}
func (AttrSystem) Recommended()    {}
func (AttrSystem) Key() string     { return "db_system" }
func (a AttrSystem) Value() string { return string(a) }

const SystemOtherSql AttrSystem = "other_sql"
const SystemAdabas AttrSystem = "adabas"
const SystemCache AttrSystem = "cache"
const SystemIntersystemsCache AttrSystem = "intersystems_cache"
const SystemCassandra AttrSystem = "cassandra"
const SystemClickhouse AttrSystem = "clickhouse"
const SystemCloudscape AttrSystem = "cloudscape"
const SystemCockroachdb AttrSystem = "cockroachdb"
const SystemColdfusion AttrSystem = "coldfusion"
const SystemCosmosdb AttrSystem = "cosmosdb"
const SystemCouchbase AttrSystem = "couchbase"
const SystemCouchdb AttrSystem = "couchdb"
const SystemDb2 AttrSystem = "db2"
const SystemDerby AttrSystem = "derby"
const SystemDynamodb AttrSystem = "dynamodb"
const SystemEdb AttrSystem = "edb"
const SystemElasticsearch AttrSystem = "elasticsearch"
const SystemFilemaker AttrSystem = "filemaker"
const SystemFirebird AttrSystem = "firebird"
const SystemFirstsql AttrSystem = "firstsql"
const SystemGeode AttrSystem = "geode"
const SystemH2 AttrSystem = "h2"
const SystemHanadb AttrSystem = "hanadb"
const SystemHbase AttrSystem = "hbase"
const SystemHive AttrSystem = "hive"
const SystemHsqldb AttrSystem = "hsqldb"
const SystemInfluxdb AttrSystem = "influxdb"
const SystemInformix AttrSystem = "informix"
const SystemIngres AttrSystem = "ingres"
const SystemInstantdb AttrSystem = "instantdb"
const SystemInterbase AttrSystem = "interbase"
const SystemMariadb AttrSystem = "mariadb"
const SystemMaxdb AttrSystem = "maxdb"
const SystemMemcached AttrSystem = "memcached"
const SystemMongodb AttrSystem = "mongodb"
const SystemMssql AttrSystem = "mssql"
const SystemMssqlcompact AttrSystem = "mssqlcompact"
const SystemMysql AttrSystem = "mysql"
const SystemNeo4j AttrSystem = "neo4j"
const SystemNetezza AttrSystem = "netezza"
const SystemOpensearch AttrSystem = "opensearch"
const SystemOracle AttrSystem = "oracle"
const SystemPervasive AttrSystem = "pervasive"
const SystemPointbase AttrSystem = "pointbase"
const SystemPostgresql AttrSystem = "postgresql"
const SystemProgress AttrSystem = "progress"
const SystemRedis AttrSystem = "redis"
const SystemRedshift AttrSystem = "redshift"
const SystemSpanner AttrSystem = "spanner"
const SystemSqlite AttrSystem = "sqlite"
const SystemSybase AttrSystem = "sybase"
const SystemTeradata AttrSystem = "teradata"
const SystemTrino AttrSystem = "trino"
const SystemVertica AttrSystem = "vertica"

// The database management system (DBMS) product as identified by the client instrumentation.
// The actual DBMS may differ from the one identified by the client. For example, when using PostgreSQL client libraries to connect to a CockroachDB, the `db.system.name` is set to `postgresql` based on the instrumentation's best knowledge
type AttrSystemName string // db.system.name

func (AttrSystemName) Stable()         {}
func (AttrSystemName) Recommended()    {}
func (AttrSystemName) Key() string     { return "db_system_name" }
func (a AttrSystemName) Value() string { return string(a) }

const SystemNameOtherSql AttrSystemName = "other_sql"
const SystemNameSoftwareagAdabas AttrSystemName = "softwareag.adabas"
const SystemNameActianIngres AttrSystemName = "actian.ingres"
const SystemNameAwsDynamodb AttrSystemName = "aws.dynamodb"
const SystemNameAwsRedshift AttrSystemName = "aws.redshift"
const SystemNameAzureCosmosdb AttrSystemName = "azure.cosmosdb"
const SystemNameIntersystemsCache AttrSystemName = "intersystems.cache"
const SystemNameCassandra AttrSystemName = "cassandra"
const SystemNameClickhouse AttrSystemName = "clickhouse"
const SystemNameCockroachdb AttrSystemName = "cockroachdb"
const SystemNameCouchbase AttrSystemName = "couchbase"
const SystemNameCouchdb AttrSystemName = "couchdb"
const SystemNameDerby AttrSystemName = "derby"
const SystemNameElasticsearch AttrSystemName = "elasticsearch"
const SystemNameFirebirdsql AttrSystemName = "firebirdsql"
const SystemNameGcpSpanner AttrSystemName = "gcp.spanner"
const SystemNameGeode AttrSystemName = "geode"
const SystemNameH2database AttrSystemName = "h2database"
const SystemNameHbase AttrSystemName = "hbase"
const SystemNameHive AttrSystemName = "hive"
const SystemNameHsqldb AttrSystemName = "hsqldb"
const SystemNameIbmDb2 AttrSystemName = "ibm.db2"
const SystemNameIbmInformix AttrSystemName = "ibm.informix"
const SystemNameIbmNetezza AttrSystemName = "ibm.netezza"
const SystemNameInfluxdb AttrSystemName = "influxdb"
const SystemNameInstantdb AttrSystemName = "instantdb"
const SystemNameMariadb AttrSystemName = "mariadb"
const SystemNameMemcached AttrSystemName = "memcached"
const SystemNameMongodb AttrSystemName = "mongodb"
const SystemNameMicrosoftSqlServer AttrSystemName = "microsoft.sql_server"
const SystemNameMysql AttrSystemName = "mysql"
const SystemNameNeo4j AttrSystemName = "neo4j"
const SystemNameOpensearch AttrSystemName = "opensearch"
const SystemNameOracleDb AttrSystemName = "oracle.db"
const SystemNamePostgresql AttrSystemName = "postgresql"
const SystemNameRedis AttrSystemName = "redis"
const SystemNameSapHana AttrSystemName = "sap.hana"
const SystemNameSapMaxdb AttrSystemName = "sap.maxdb"
const SystemNameSqlite AttrSystemName = "sqlite"
const SystemNameTeradata AttrSystemName = "teradata"
const SystemNameTrino AttrSystemName = "trino"

// Deprecated, no replacement at this time
type AttrUser string // db.user

func (AttrUser) Development()    {}
func (AttrUser) Recommended()    {}
func (AttrUser) Key() string     { return "db_user" }
func (a AttrUser) Value() string { return string(a) }

/* State {
    name: "attr.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "ctx": {
            "attributes": [
                {
                    "brief": "Deprecated, use `cassandra.consistency.level` instead.",
                    "deprecated": {
                        "note": "Replaced by `cassandra.consistency.level`.",
                        "reason": "renamed",
                        "renamed_to": "cassandra.consistency.level",
                    },
                    "name": "db.cassandra.consistency_level",
                    "requirement_level": "recommended",
                    "root_namespace": "db",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "all",
                                "note": none,
                                "stability": "development",
                                "value": "all",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "each_quorum",
                                "note": none,
                                "stability": "development",
                                "value": "each_quorum",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "quorum",
                                "note": none,
                                "stability": "development",
                                "value": "quorum",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "local_quorum",
                                "note": none,
                                "stability": "development",
                                "value": "local_quorum",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "one",
                                "note": none,
                                "stability": "development",
                                "value": "one",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "two",
                                "note": none,
                                "stability": "development",
                                "value": "two",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "three",
                                "note": none,
                                "stability": "development",
                                "value": "three",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "local_one",
                                "note": none,
                                "stability": "development",
                                "value": "local_one",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "any",
                                "note": none,
                                "stability": "development",
                                "value": "any",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "serial",
                                "note": none,
                                "stability": "development",
                                "value": "serial",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "local_serial",
                                "note": none,
                                "stability": "development",
                                "value": "local_serial",
                            },
                        ],
                    },
                },
                {
                    "brief": "Deprecated, use `cassandra.coordinator.dc` instead.",
                    "deprecated": {
                        "note": "Replaced by `cassandra.coordinator.dc`.",
                        "reason": "renamed",
                        "renamed_to": "cassandra.coordinator.dc",
                    },
                    "examples": "us-west-2",
                    "name": "db.cassandra.coordinator.dc",
                    "requirement_level": "recommended",
                    "root_namespace": "db",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Deprecated, use `cassandra.coordinator.id` instead.",
                    "deprecated": {
                        "note": "Replaced by `cassandra.coordinator.id`.",
                        "reason": "renamed",
                        "renamed_to": "cassandra.coordinator.id",
                    },
                    "examples": "be13faa2-8574-4d71-926d-27f16cf8a7af",
                    "name": "db.cassandra.coordinator.id",
                    "requirement_level": "recommended",
                    "root_namespace": "db",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Deprecated, use `cassandra.query.idempotent` instead.",
                    "deprecated": {
                        "note": "Replaced by `cassandra.query.idempotent`.",
                        "reason": "renamed",
                        "renamed_to": "cassandra.query.idempotent",
                    },
                    "name": "db.cassandra.idempotence",
                    "requirement_level": "recommended",
                    "root_namespace": "db",
                    "stability": "development",
                    "type": "boolean",
                },
                {
                    "brief": "Deprecated, use `cassandra.page.size` instead.",
                    "deprecated": {
                        "note": "Replaced by `cassandra.page.size`.",
                        "reason": "renamed",
                        "renamed_to": "cassandra.page.size",
                    },
                    "examples": [
                        5000,
                    ],
                    "name": "db.cassandra.page_size",
                    "requirement_level": "recommended",
                    "root_namespace": "db",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "Deprecated, use `cassandra.speculative_execution.count` instead.",
                    "deprecated": {
                        "note": "Replaced by `cassandra.speculative_execution.count`.",
                        "reason": "renamed",
                        "renamed_to": "cassandra.speculative_execution.count",
                    },
                    "examples": [
                        0,
                        2,
                    ],
                    "name": "db.cassandra.speculative_execution_count",
                    "requirement_level": "recommended",
                    "root_namespace": "db",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "Deprecated, use `db.collection.name` instead.",
                    "deprecated": {
                        "note": "Replaced by `db.collection.name`.",
                        "reason": "renamed",
                        "renamed_to": "db.collection.name",
                    },
                    "examples": "mytable",
                    "name": "db.cassandra.table",
                    "requirement_level": "recommended",
                    "root_namespace": "db",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The name of the connection pool; unique within the instrumented application. In case the connection pool implementation doesn't provide a name, instrumentation SHOULD use a combination of parameters that would make the name unique, for example, combining attributes `server.address`, `server.port`, and `db.namespace`, formatted as `server.address:server.port/db.namespace`. Instrumentations that generate connection pool name following different patterns SHOULD document it.\n",
                    "examples": [
                        "myDataSource",
                    ],
                    "name": "db.client.connection.pool.name",
                    "requirement_level": "recommended",
                    "root_namespace": "db",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The state of a connection in the pool",
                    "examples": [
                        "idle",
                    ],
                    "name": "db.client.connection.state",
                    "requirement_level": "recommended",
                    "root_namespace": "db",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "idle",
                                "note": none,
                                "stability": "development",
                                "value": "idle",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "used",
                                "note": none,
                                "stability": "development",
                                "value": "used",
                            },
                        ],
                    },
                },
                {
                    "brief": "Deprecated, use `db.client.connection.pool.name` instead.",
                    "deprecated": {
                        "note": "Replaced by `db.client.connection.pool.name`.",
                        "reason": "renamed",
                        "renamed_to": "db.client.connection.pool.name",
                    },
                    "examples": [
                        "myDataSource",
                    ],
                    "name": "db.client.connections.pool.name",
                    "requirement_level": "recommended",
                    "root_namespace": "db",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Deprecated, use `db.client.connection.state` instead.",
                    "deprecated": {
                        "note": "Replaced by `db.client.connection.state`.",
                        "reason": "renamed",
                        "renamed_to": "db.client.connection.state",
                    },
                    "examples": [
                        "idle",
                    ],
                    "name": "db.client.connections.state",
                    "requirement_level": "recommended",
                    "root_namespace": "db",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "idle",
                                "note": none,
                                "stability": "development",
                                "value": "idle",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "used",
                                "note": none,
                                "stability": "development",
                                "value": "used",
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
                    "requirement_level": "recommended",
                    "root_namespace": "db",
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "Deprecated, use `server.address`, `server.port` attributes instead.",
                    "deprecated": {
                        "note": "Replaced by `server.address` and `server.port`.\n",
                        "reason": "uncategorized",
                    },
                    "examples": "Server=(localdb)\\v11.0;Integrated Security=true;",
                    "name": "db.connection_string",
                    "requirement_level": "recommended",
                    "root_namespace": "db",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Deprecated, use `azure.client.id` instead.",
                    "deprecated": {
                        "note": "Replaced by `azure.client.id`.",
                        "reason": "renamed",
                        "renamed_to": "azure.client.id",
                    },
                    "examples": "3ba4827d-4422-483f-b59f-85b74211c11d",
                    "name": "db.cosmosdb.client_id",
                    "requirement_level": "recommended",
                    "root_namespace": "db",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Deprecated, use `azure.cosmosdb.connection.mode` instead.",
                    "deprecated": {
                        "note": "Replaced by `azure.cosmosdb.connection.mode`.",
                        "reason": "renamed",
                        "renamed_to": "azure.cosmosdb.connection.mode",
                    },
                    "name": "db.cosmosdb.connection_mode",
                    "requirement_level": "recommended",
                    "root_namespace": "db",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": "Gateway (HTTP) connection.",
                                "deprecated": none,
                                "id": "gateway",
                                "note": none,
                                "stability": "development",
                                "value": "gateway",
                            },
                            {
                                "brief": "Direct connection.",
                                "deprecated": none,
                                "id": "direct",
                                "note": none,
                                "stability": "development",
                                "value": "direct",
                            },
                        ],
                    },
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
                    "requirement_level": "recommended",
                    "root_namespace": "db",
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
                    "brief": "Deprecated, use `db.collection.name` instead.",
                    "deprecated": {
                        "note": "Replaced by `db.collection.name`.",
                        "reason": "renamed",
                        "renamed_to": "db.collection.name",
                    },
                    "examples": "mytable",
                    "name": "db.cosmosdb.container",
                    "requirement_level": "recommended",
                    "root_namespace": "db",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Deprecated, no replacement at this time.",
                    "deprecated": {
                        "note": "Removed, no replacement at this time.\n",
                        "reason": "obsoleted",
                    },
                    "name": "db.cosmosdb.operation_type",
                    "requirement_level": "recommended",
                    "root_namespace": "db",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "batch",
                                "note": none,
                                "stability": "development",
                                "value": "batch",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "create",
                                "note": none,
                                "stability": "development",
                                "value": "create",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "delete",
                                "note": none,
                                "stability": "development",
                                "value": "delete",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "execute",
                                "note": none,
                                "stability": "development",
                                "value": "execute",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "execute_javascript",
                                "note": none,
                                "stability": "development",
                                "value": "execute_javascript",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "invalid",
                                "note": none,
                                "stability": "development",
                                "value": "invalid",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "head",
                                "note": none,
                                "stability": "development",
                                "value": "head",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "head_feed",
                                "note": none,
                                "stability": "development",
                                "value": "head_feed",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "patch",
                                "note": none,
                                "stability": "development",
                                "value": "patch",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "query",
                                "note": none,
                                "stability": "development",
                                "value": "query",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "query_plan",
                                "note": none,
                                "stability": "development",
                                "value": "query_plan",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "read",
                                "note": none,
                                "stability": "development",
                                "value": "read",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "read_feed",
                                "note": none,
                                "stability": "development",
                                "value": "read_feed",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "replace",
                                "note": none,
                                "stability": "development",
                                "value": "replace",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "upsert",
                                "note": none,
                                "stability": "development",
                                "value": "upsert",
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
                    "requirement_level": "recommended",
                    "root_namespace": "db",
                    "stability": "development",
                    "type": "string[]",
                },
                {
                    "brief": "Deprecated, use `azure.cosmosdb.operation.request_charge` instead.",
                    "deprecated": {
                        "note": "Replaced by `azure.cosmosdb.operation.request_charge`.",
                        "reason": "renamed",
                        "renamed_to": "azure.cosmosdb.operation.request_charge",
                    },
                    "examples": [
                        46.18,
                        1.0,
                    ],
                    "name": "db.cosmosdb.request_charge",
                    "requirement_level": "recommended",
                    "root_namespace": "db",
                    "stability": "development",
                    "type": "double",
                },
                {
                    "brief": "Deprecated, use `azure.cosmosdb.request.body.size` instead.",
                    "deprecated": {
                        "note": "Replaced by `azure.cosmosdb.request.body.size`.",
                        "reason": "renamed",
                        "renamed_to": "azure.cosmosdb.request.body.size",
                    },
                    "name": "db.cosmosdb.request_content_length",
                    "requirement_level": "recommended",
                    "root_namespace": "db",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "Deprecated, use `db.response.status_code` instead.",
                    "deprecated": {
                        "note": "Replaced by `db.response.status_code`.",
                        "reason": "renamed",
                        "renamed_to": "db.response.status_code",
                    },
                    "examples": [
                        200,
                        201,
                    ],
                    "name": "db.cosmosdb.status_code",
                    "requirement_level": "recommended",
                    "root_namespace": "db",
                    "stability": "development",
                    "type": "int",
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
                    "requirement_level": "recommended",
                    "root_namespace": "db",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "Deprecated, use `db.namespace` instead.\n",
                    "deprecated": {
                        "note": "Replaced by `db.namespace`.",
                        "reason": "renamed",
                        "renamed_to": "db.namespace",
                    },
                    "examples": [
                        "e9106fc68e3044f0b1475b04bf4ffd5f",
                    ],
                    "name": "db.elasticsearch.cluster.name",
                    "requirement_level": "recommended",
                    "root_namespace": "db",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Deprecated, use `elasticsearch.node.name` instead.\n",
                    "deprecated": {
                        "note": "Replaced by `elasticsearch.node.name`.",
                        "reason": "renamed",
                        "renamed_to": "elasticsearch.node.name",
                    },
                    "examples": [
                        "instance-0000000001",
                    ],
                    "name": "db.elasticsearch.node.name",
                    "requirement_level": "recommended",
                    "root_namespace": "db",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Deprecated, use `db.operation.parameter` instead.\n",
                    "deprecated": {
                        "note": "Replaced by `db.operation.parameter`.",
                        "reason": "renamed",
                        "renamed_to": "db.operation.parameter",
                    },
                    "examples": [
                        "test-index",
                        "123",
                    ],
                    "name": "db.elasticsearch.path_parts",
                    "requirement_level": "recommended",
                    "root_namespace": "db",
                    "stability": "development",
                    "type": "template[string]",
                },
                {
                    "brief": "Deprecated, no general replacement at this time. For Elasticsearch, use `db.elasticsearch.node.name` instead.",
                    "deprecated": {
                        "note": "Removed, no general replacement at this time. For Elasticsearch, use `db.elasticsearch.node.name` instead.\n",
                        "reason": "obsoleted",
                    },
                    "examples": "mysql-e26b99z.example.com",
                    "name": "db.instance.id",
                    "requirement_level": "recommended",
                    "root_namespace": "db",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Removed, no replacement at this time.",
                    "deprecated": {
                        "note": "Removed, no replacement at this time.\n",
                        "reason": "obsoleted",
                    },
                    "examples": [
                        "org.postgresql.Driver",
                        "com.microsoft.sqlserver.jdbc.SQLServerDriver",
                    ],
                    "name": "db.jdbc.driver_classname",
                    "requirement_level": "recommended",
                    "root_namespace": "db",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Deprecated, use `db.collection.name` instead.",
                    "deprecated": {
                        "note": "Replaced by `db.collection.name`.",
                        "reason": "renamed",
                        "renamed_to": "db.collection.name",
                    },
                    "examples": "mytable",
                    "name": "db.mongodb.collection",
                    "requirement_level": "recommended",
                    "root_namespace": "db",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Deprecated, SQL Server instance is now populated as a part of `db.namespace` attribute.",
                    "deprecated": {
                        "note": "Removed, no replacement at this time.",
                        "reason": "obsoleted",
                    },
                    "examples": "MSSQLSERVER",
                    "name": "db.mssql.instance_name",
                    "requirement_level": "recommended",
                    "root_namespace": "db",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Deprecated, use `db.namespace` instead.",
                    "deprecated": {
                        "note": "Replaced by `db.namespace`.",
                        "reason": "renamed",
                        "renamed_to": "db.namespace",
                    },
                    "examples": [
                        "customers",
                        "main",
                    ],
                    "name": "db.name",
                    "requirement_level": "recommended",
                    "root_namespace": "db",
                    "stability": "development",
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
                    "requirement_level": "recommended",
                    "root_namespace": "db",
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "Deprecated, use `db.operation.name` instead.",
                    "deprecated": {
                        "note": "Replaced by `db.operation.name`.",
                        "reason": "renamed",
                        "renamed_to": "db.operation.name",
                    },
                    "examples": [
                        "findAndModify",
                        "HMSET",
                        "SELECT",
                    ],
                    "name": "db.operation",
                    "requirement_level": "recommended",
                    "root_namespace": "db",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The number of queries included in a batch operation.",
                    "examples": [
                        2,
                        3,
                        4,
                    ],
                    "name": "db.operation.batch.size",
                    "note": "Operations are only considered batches when they contain two or more operations, and so `db.operation.batch.size` SHOULD never be `1`.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "db",
                    "stability": "stable",
                    "type": "int",
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
                    "requirement_level": "recommended",
                    "root_namespace": "db",
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "A database operation parameter, with `<key>` being the parameter name, and the attribute value being a string representation of the parameter value.\n",
                    "examples": [
                        "someval",
                        "55",
                    ],
                    "name": "db.operation.parameter",
                    "note": "For example, a client-side maximum number of rows to read from the database\nMAY be recorded as the `db.operation.parameter.max_rows` attribute.\n\n`db.query.text` parameters SHOULD be captured using `db.query.parameter.<key>`\ninstead of `db.operation.parameter.<key>`.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "db",
                    "stability": "development",
                    "type": "template[string]",
                },
                {
                    "brief": "A database query parameter, with `<key>` being the parameter name, and the attribute value being a string representation of the parameter value.\n",
                    "examples": [
                        "someval",
                        "55",
                    ],
                    "name": "db.query.parameter",
                    "note": "If a query parameter has no name and instead is referenced only by index,\nthen `<key>` SHOULD be the 0-based index.\n\n`db.query.parameter.<key>` SHOULD match\nup with the parameterized placeholders present in `db.query.text`.\n\n`db.query.parameter.<key>` SHOULD NOT be captured on batch operations.\n\nExamples:\n\n- For a query `SELECT * FROM users where username =  %s` with the parameter `\"jdoe\"`,\n  the attribute `db.query.parameter.0` SHOULD be set to `\"jdoe\"`.\n\n- For a query `\"SELECT * FROM users WHERE username = %(username)s;` with parameter\n  `username = \"jdoe\"`, the attribute `db.query.parameter.username` SHOULD be set to `\"jdoe\"`.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "db",
                    "stability": "development",
                    "type": "template[string]",
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
                    "requirement_level": "recommended",
                    "root_namespace": "db",
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
                    "requirement_level": "recommended",
                    "root_namespace": "db",
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "Deprecated, use `db.namespace` instead.",
                    "deprecated": {
                        "note": "Replaced by `db.namespace`.",
                        "reason": "renamed",
                        "renamed_to": "db.namespace",
                    },
                    "examples": [
                        0,
                        1,
                        15,
                    ],
                    "name": "db.redis.database_index",
                    "requirement_level": "recommended",
                    "root_namespace": "db",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "Number of rows returned by the operation.",
                    "examples": [
                        10,
                        30,
                        1000,
                    ],
                    "name": "db.response.returned_rows",
                    "requirement_level": "recommended",
                    "root_namespace": "db",
                    "stability": "development",
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
                    "requirement_level": "recommended",
                    "root_namespace": "db",
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "Deprecated, use `db.collection.name` instead.",
                    "deprecated": {
                        "note": "Replaced by `db.collection.name`, but only if not extracting the value from `db.query.text`.",
                        "reason": "uncategorized",
                    },
                    "examples": "mytable",
                    "name": "db.sql.table",
                    "requirement_level": "recommended",
                    "root_namespace": "db",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The database statement being executed.",
                    "deprecated": {
                        "note": "Replaced by `db.query.text`.",
                        "reason": "renamed",
                        "renamed_to": "db.query.text",
                    },
                    "examples": [
                        "SELECT * FROM wuser_table",
                        "SET mykey \"WuValue\"",
                    ],
                    "name": "db.statement",
                    "requirement_level": "recommended",
                    "root_namespace": "db",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The name of a stored procedure within the database.",
                    "examples": [
                        "GetCustomer",
                    ],
                    "name": "db.stored_procedure.name",
                    "note": "It is RECOMMENDED to capture the value as provided by the application\nwithout attempting to do any case normalization.\n\nFor batch operations, if the individual operations are known to have the same\nstored procedure name then that stored procedure name SHOULD be used.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "db",
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "Deprecated, use `db.system.name` instead.",
                    "deprecated": {
                        "note": "Replaced by `db.system.name`.",
                        "reason": "renamed",
                        "renamed_to": "db.system.name",
                    },
                    "name": "db.system",
                    "requirement_level": "recommended",
                    "root_namespace": "db",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": "Some other SQL database. Fallback only. See notes.",
                                "deprecated": none,
                                "id": "other_sql",
                                "note": none,
                                "stability": "development",
                                "value": "other_sql",
                            },
                            {
                                "brief": "Adabas (Adaptable Database System)",
                                "deprecated": none,
                                "id": "adabas",
                                "note": none,
                                "stability": "development",
                                "value": "adabas",
                            },
                            {
                                "brief": "Deprecated, use `intersystems_cache` instead.",
                                "deprecated": "Replaced by `intersystems_cache`.",
                                "id": "cache",
                                "note": none,
                                "stability": "development",
                                "value": "cache",
                            },
                            {
                                "brief": "InterSystems Caché",
                                "deprecated": none,
                                "id": "intersystems_cache",
                                "note": none,
                                "stability": "development",
                                "value": "intersystems_cache",
                            },
                            {
                                "brief": "Apache Cassandra",
                                "deprecated": none,
                                "id": "cassandra",
                                "note": none,
                                "stability": "development",
                                "value": "cassandra",
                            },
                            {
                                "brief": "ClickHouse",
                                "deprecated": none,
                                "id": "clickhouse",
                                "note": none,
                                "stability": "development",
                                "value": "clickhouse",
                            },
                            {
                                "brief": "Deprecated, use `other_sql` instead.",
                                "deprecated": "Replaced by `other_sql`.",
                                "id": "cloudscape",
                                "note": none,
                                "stability": "development",
                                "value": "cloudscape",
                            },
                            {
                                "brief": "CockroachDB",
                                "deprecated": none,
                                "id": "cockroachdb",
                                "note": none,
                                "stability": "development",
                                "value": "cockroachdb",
                            },
                            {
                                "brief": "Deprecated, no replacement at this time.",
                                "deprecated": "Removed.",
                                "id": "coldfusion",
                                "note": none,
                                "stability": "development",
                                "value": "coldfusion",
                            },
                            {
                                "brief": "Microsoft Azure Cosmos DB",
                                "deprecated": none,
                                "id": "cosmosdb",
                                "note": none,
                                "stability": "development",
                                "value": "cosmosdb",
                            },
                            {
                                "brief": "Couchbase",
                                "deprecated": none,
                                "id": "couchbase",
                                "note": none,
                                "stability": "development",
                                "value": "couchbase",
                            },
                            {
                                "brief": "CouchDB",
                                "deprecated": none,
                                "id": "couchdb",
                                "note": none,
                                "stability": "development",
                                "value": "couchdb",
                            },
                            {
                                "brief": "IBM Db2",
                                "deprecated": none,
                                "id": "db2",
                                "note": none,
                                "stability": "development",
                                "value": "db2",
                            },
                            {
                                "brief": "Apache Derby",
                                "deprecated": none,
                                "id": "derby",
                                "note": none,
                                "stability": "development",
                                "value": "derby",
                            },
                            {
                                "brief": "Amazon DynamoDB",
                                "deprecated": none,
                                "id": "dynamodb",
                                "note": none,
                                "stability": "development",
                                "value": "dynamodb",
                            },
                            {
                                "brief": "EnterpriseDB",
                                "deprecated": none,
                                "id": "edb",
                                "note": none,
                                "stability": "development",
                                "value": "edb",
                            },
                            {
                                "brief": "Elasticsearch",
                                "deprecated": none,
                                "id": "elasticsearch",
                                "note": none,
                                "stability": "development",
                                "value": "elasticsearch",
                            },
                            {
                                "brief": "FileMaker",
                                "deprecated": none,
                                "id": "filemaker",
                                "note": none,
                                "stability": "development",
                                "value": "filemaker",
                            },
                            {
                                "brief": "Firebird",
                                "deprecated": none,
                                "id": "firebird",
                                "note": none,
                                "stability": "development",
                                "value": "firebird",
                            },
                            {
                                "brief": "Deprecated, use `other_sql` instead.",
                                "deprecated": "Replaced by `other_sql`.",
                                "id": "firstsql",
                                "note": none,
                                "stability": "development",
                                "value": "firstsql",
                            },
                            {
                                "brief": "Apache Geode",
                                "deprecated": none,
                                "id": "geode",
                                "note": none,
                                "stability": "development",
                                "value": "geode",
                            },
                            {
                                "brief": "H2",
                                "deprecated": none,
                                "id": "h2",
                                "note": none,
                                "stability": "development",
                                "value": "h2",
                            },
                            {
                                "brief": "SAP HANA",
                                "deprecated": none,
                                "id": "hanadb",
                                "note": none,
                                "stability": "development",
                                "value": "hanadb",
                            },
                            {
                                "brief": "Apache HBase",
                                "deprecated": none,
                                "id": "hbase",
                                "note": none,
                                "stability": "development",
                                "value": "hbase",
                            },
                            {
                                "brief": "Apache Hive",
                                "deprecated": none,
                                "id": "hive",
                                "note": none,
                                "stability": "development",
                                "value": "hive",
                            },
                            {
                                "brief": "HyperSQL DataBase",
                                "deprecated": none,
                                "id": "hsqldb",
                                "note": none,
                                "stability": "development",
                                "value": "hsqldb",
                            },
                            {
                                "brief": "InfluxDB",
                                "deprecated": none,
                                "id": "influxdb",
                                "note": none,
                                "stability": "development",
                                "value": "influxdb",
                            },
                            {
                                "brief": "Informix",
                                "deprecated": none,
                                "id": "informix",
                                "note": none,
                                "stability": "development",
                                "value": "informix",
                            },
                            {
                                "brief": "Ingres",
                                "deprecated": none,
                                "id": "ingres",
                                "note": none,
                                "stability": "development",
                                "value": "ingres",
                            },
                            {
                                "brief": "InstantDB",
                                "deprecated": none,
                                "id": "instantdb",
                                "note": none,
                                "stability": "development",
                                "value": "instantdb",
                            },
                            {
                                "brief": "InterBase",
                                "deprecated": none,
                                "id": "interbase",
                                "note": none,
                                "stability": "development",
                                "value": "interbase",
                            },
                            {
                                "brief": "MariaDB",
                                "deprecated": none,
                                "id": "mariadb",
                                "note": none,
                                "stability": "development",
                                "value": "mariadb",
                            },
                            {
                                "brief": "SAP MaxDB",
                                "deprecated": none,
                                "id": "maxdb",
                                "note": none,
                                "stability": "development",
                                "value": "maxdb",
                            },
                            {
                                "brief": "Memcached",
                                "deprecated": none,
                                "id": "memcached",
                                "note": none,
                                "stability": "development",
                                "value": "memcached",
                            },
                            {
                                "brief": "MongoDB",
                                "deprecated": none,
                                "id": "mongodb",
                                "note": none,
                                "stability": "development",
                                "value": "mongodb",
                            },
                            {
                                "brief": "Microsoft SQL Server",
                                "deprecated": none,
                                "id": "mssql",
                                "note": none,
                                "stability": "development",
                                "value": "mssql",
                            },
                            {
                                "brief": "Deprecated, Microsoft SQL Server Compact is discontinued.",
                                "deprecated": "Removed, use `other_sql` instead.",
                                "id": "mssqlcompact",
                                "note": none,
                                "stability": "development",
                                "value": "mssqlcompact",
                            },
                            {
                                "brief": "MySQL",
                                "deprecated": none,
                                "id": "mysql",
                                "note": none,
                                "stability": "development",
                                "value": "mysql",
                            },
                            {
                                "brief": "Neo4j",
                                "deprecated": none,
                                "id": "neo4j",
                                "note": none,
                                "stability": "development",
                                "value": "neo4j",
                            },
                            {
                                "brief": "Netezza",
                                "deprecated": none,
                                "id": "netezza",
                                "note": none,
                                "stability": "development",
                                "value": "netezza",
                            },
                            {
                                "brief": "OpenSearch",
                                "deprecated": none,
                                "id": "opensearch",
                                "note": none,
                                "stability": "development",
                                "value": "opensearch",
                            },
                            {
                                "brief": "Oracle Database",
                                "deprecated": none,
                                "id": "oracle",
                                "note": none,
                                "stability": "development",
                                "value": "oracle",
                            },
                            {
                                "brief": "Pervasive PSQL",
                                "deprecated": none,
                                "id": "pervasive",
                                "note": none,
                                "stability": "development",
                                "value": "pervasive",
                            },
                            {
                                "brief": "PointBase",
                                "deprecated": none,
                                "id": "pointbase",
                                "note": none,
                                "stability": "development",
                                "value": "pointbase",
                            },
                            {
                                "brief": "PostgreSQL",
                                "deprecated": none,
                                "id": "postgresql",
                                "note": none,
                                "stability": "development",
                                "value": "postgresql",
                            },
                            {
                                "brief": "Progress Database",
                                "deprecated": none,
                                "id": "progress",
                                "note": none,
                                "stability": "development",
                                "value": "progress",
                            },
                            {
                                "brief": "Redis",
                                "deprecated": none,
                                "id": "redis",
                                "note": none,
                                "stability": "development",
                                "value": "redis",
                            },
                            {
                                "brief": "Amazon Redshift",
                                "deprecated": none,
                                "id": "redshift",
                                "note": none,
                                "stability": "development",
                                "value": "redshift",
                            },
                            {
                                "brief": "Cloud Spanner",
                                "deprecated": none,
                                "id": "spanner",
                                "note": none,
                                "stability": "development",
                                "value": "spanner",
                            },
                            {
                                "brief": "SQLite",
                                "deprecated": none,
                                "id": "sqlite",
                                "note": none,
                                "stability": "development",
                                "value": "sqlite",
                            },
                            {
                                "brief": "Sybase",
                                "deprecated": none,
                                "id": "sybase",
                                "note": none,
                                "stability": "development",
                                "value": "sybase",
                            },
                            {
                                "brief": "Teradata",
                                "deprecated": none,
                                "id": "teradata",
                                "note": none,
                                "stability": "development",
                                "value": "teradata",
                            },
                            {
                                "brief": "Trino",
                                "deprecated": none,
                                "id": "trino",
                                "note": none,
                                "stability": "development",
                                "value": "trino",
                            },
                            {
                                "brief": "Vertica",
                                "deprecated": none,
                                "id": "vertica",
                                "note": none,
                                "stability": "development",
                                "value": "vertica",
                            },
                        ],
                    },
                },
                {
                    "brief": "The database management system (DBMS) product as identified by the client instrumentation.",
                    "name": "db.system.name",
                    "note": "The actual DBMS may differ from the one identified by the client. For example, when using PostgreSQL client libraries to connect to a CockroachDB, the `db.system.name` is set to `postgresql` based on the instrumentation's best knowledge.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "db",
                    "stability": "stable",
                    "type": {
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
                    "brief": "Deprecated, no replacement at this time.",
                    "deprecated": {
                        "note": "Removed, no replacement at this time.",
                        "reason": "obsoleted",
                    },
                    "examples": [
                        "readonly_user",
                        "reporting_user",
                    ],
                    "name": "db.user",
                    "requirement_level": "recommended",
                    "root_namespace": "db",
                    "stability": "development",
                    "type": "string",
                },
            ],
            "root_namespace": "db",
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
