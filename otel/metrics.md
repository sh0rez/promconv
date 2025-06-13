# Metrics


## aspnetcore

### aspnetcore.diagnostics.exceptions

```py
# TYPE aspnetcore.diagnostics.exceptions COUNTER
aspnetcore.diagnostics.exceptions{error.type, aspnetcore.diagnostics.handler.type, aspnetcore.diagnostics.exception.result}
```

Number of exceptions caught by exception handling middleware.


|Attribute|Type|Description|
|-|-|-|
| error.type | `_OTHER` | The full name of exception type.  |
| aspnetcore.diagnostics.handler.type | string | Full type name of the [`IExceptionHandler`](https://learn.microsoft.com/dotnet/api/microsoft.aspnetcore.diagnostics.iexceptionhandler) implementation that handled the exception.  |
| aspnetcore.diagnostics.exception.result | `handled` \| `unhandled` \| `skipped` \| `aborted` | ASP.NET Core exception middleware handling result  |




### aspnetcore.rate_limiting.active_request_leases

```py
# TYPE aspnetcore.rate_limiting.active_request_leases UPDOWNCOUNTER
aspnetcore.rate_limiting.active_request_leases{aspnetcore.rate_limiting.policy}
```

Number of requests that are currently active on the server that hold a rate limiting lease.


|Attribute|Type|Description|
|-|-|-|
| aspnetcore.rate_limiting.policy | string | Rate limiting policy name.  |




### aspnetcore.rate_limiting.queued_requests

```py
# TYPE aspnetcore.rate_limiting.queued_requests UPDOWNCOUNTER
aspnetcore.rate_limiting.queued_requests{aspnetcore.rate_limiting.policy}
```

Number of requests that are currently queued, waiting to acquire a rate limiting lease.


|Attribute|Type|Description|
|-|-|-|
| aspnetcore.rate_limiting.policy | string | Rate limiting policy name.  |




### aspnetcore.rate_limiting.request.time_in_queue

```py
# TYPE aspnetcore.rate_limiting.request.time_in_queue HISTOGRAM
aspnetcore.rate_limiting.request.time_in_queue{aspnetcore.rate_limiting.policy, aspnetcore.rate_limiting.result}
```

The time the request spent in a queue waiting to acquire a rate limiting lease.


|Attribute|Type|Description|
|-|-|-|
| aspnetcore.rate_limiting.policy | string | Rate limiting policy name.  |
| aspnetcore.rate_limiting.result | `acquired` \| `endpoint_limiter` \| `global_limiter` \| `request_canceled` | Rate-limiting result, shows whether the lease was acquired or contains a rejection reason  |




### aspnetcore.rate_limiting.request_lease.duration

```py
# TYPE aspnetcore.rate_limiting.request_lease.duration HISTOGRAM
aspnetcore.rate_limiting.request_lease.duration{aspnetcore.rate_limiting.policy}
```

The duration of rate limiting lease held by requests on the server.


|Attribute|Type|Description|
|-|-|-|
| aspnetcore.rate_limiting.policy | string | Rate limiting policy name.  |




### aspnetcore.rate_limiting.requests

```py
# TYPE aspnetcore.rate_limiting.requests COUNTER
aspnetcore.rate_limiting.requests{aspnetcore.rate_limiting.policy, aspnetcore.rate_limiting.result}
```

Number of requests that tried to acquire a rate limiting lease.


|Attribute|Type|Description|
|-|-|-|
| aspnetcore.rate_limiting.policy | string | Rate limiting policy name.  |
| aspnetcore.rate_limiting.result | `acquired` \| `endpoint_limiter` \| `global_limiter` \| `request_canceled` | Rate-limiting result, shows whether the lease was acquired or contains a rejection reason  |




### aspnetcore.routing.match_attempts

```py
# TYPE aspnetcore.routing.match_attempts COUNTER
aspnetcore.routing.match_attempts{http.route, aspnetcore.routing.is_fallback, aspnetcore.routing.match_status}
```

Number of requests that were attempted to be matched to an endpoint.


|Attribute|Type|Description|
|-|-|-|
| http.route | string | The matched route, that is, the path template in the format used by the respective server framework.
  |
| aspnetcore.routing.is_fallback | boolean | A value that indicates whether the matched route is a fallback route.  |
| aspnetcore.routing.match_status | `success` \| `failure` | Match result - success or failure  |





## azure

### azure.cosmosdb.client.active_instance.count

```py
# TYPE azure.cosmosdb.client.active_instance.count UPDOWNCOUNTER
azure.cosmosdb.client.active_instance.count{server.address, server.port}
```

Number of active client instances


|Attribute|Type|Description|
|-|-|-|
| server.address | string | Name of the database host.
  |
| server.port | int | Server port number.  |




### azure.cosmosdb.client.operation.request_charge

```py
# TYPE azure.cosmosdb.client.operation.request_charge HISTOGRAM
azure.cosmosdb.client.operation.request_charge{db.response.status_code, error.type, azure.cosmosdb.consistency.level, azure.cosmosdb.response.sub_status_code, azure.cosmosdb.operation.contacted_regions, db.collection.name, db.namespace, db.operation.name, server.address, server.port}
```

[Request units](https://learn.microsoft.com/azure/cosmos-db/request-units) consumed by the operation


|Attribute|Type|Description|
|-|-|-|
| db.response.status_code | string | Database response status code.  |
| error.type | `_OTHER` | Describes a class of error the operation ended with.
  |
| azure.cosmosdb.consistency.level | `Strong` \| `BoundedStaleness` \| `Session` \| `Eventual` \| `ConsistentPrefix` | Account or request [consistency level](https://learn.microsoft.com/azure/cosmos-db/consistency-levels).  |
| azure.cosmosdb.response.sub_status_code | int | Cosmos DB sub status code.  |
| azure.cosmosdb.operation.contacted_regions | string[] | List of regions contacted during operation in the order that they were contacted. If there is more than one region listed, it indicates that the operation was performed on multiple regions i.e. cross-regional call.
  |
| db.collection.name | string | Cosmos DB container name.
  |
| db.namespace | string | The name of the database, fully qualified within the server address and port.
  |
| db.operation.name | string | The name of the operation or command being executed.
  |
| server.address | string | Name of the database host.
  |
| server.port | int | Server port number.  |





## cicd

### cicd.pipeline.run.active

```py
# TYPE cicd.pipeline.run.active UPDOWNCOUNTER
cicd.pipeline.run.active{cicd.pipeline.name, cicd.pipeline.run.state}
```

The number of pipeline runs currently active in the system by state.


|Attribute|Type|Description|
|-|-|-|
| cicd.pipeline.name | string | The human readable name of the pipeline within a CI/CD system.
  |
| cicd.pipeline.run.state | `pending` \| `executing` \| `finalizing` | The pipeline run goes through these states during its lifecycle.
  |




### cicd.pipeline.run.duration

```py
# TYPE cicd.pipeline.run.duration HISTOGRAM
cicd.pipeline.run.duration{cicd.pipeline.name, cicd.pipeline.run.state, cicd.pipeline.result, error.type}
```

Duration of a pipeline run grouped by pipeline, state and result.


|Attribute|Type|Description|
|-|-|-|
| cicd.pipeline.name | string | The human readable name of the pipeline within a CI/CD system.
  |
| cicd.pipeline.run.state | `pending` \| `executing` \| `finalizing` | The pipeline run goes through these states during its lifecycle.
  |
| cicd.pipeline.result | `success` \| `failure` \| `error` \| `timeout` \| `cancellation` \| `skip` | The result of a pipeline run.
  |
| error.type | `_OTHER` | Describes a class of error the operation ended with.
  |




### cicd.pipeline.run.errors

```py
# TYPE cicd.pipeline.run.errors COUNTER
cicd.pipeline.run.errors{cicd.pipeline.name, error.type}
```

The number of errors encountered in pipeline runs (eg. compile, test failures).


|Attribute|Type|Description|
|-|-|-|
| cicd.pipeline.name | string | The human readable name of the pipeline within a CI/CD system.
  |
| error.type | `_OTHER` | Describes a class of error the operation ended with.
  |




### cicd.system.errors

```py
# TYPE cicd.system.errors COUNTER
cicd.system.errors{error.type, cicd.system.component}
```

The number of errors in a component of the CICD system (eg. controller, scheduler, agent).


|Attribute|Type|Description|
|-|-|-|
| error.type | `_OTHER` | Describes a class of error the operation ended with.
  |
| cicd.system.component | string | The name of a component of the CICD system.  |




### cicd.worker.count

```py
# TYPE cicd.worker.count UPDOWNCOUNTER
cicd.worker.count{cicd.worker.state}
```

The number of workers on the CICD system by state.


|Attribute|Type|Description|
|-|-|-|
| cicd.worker.state | `available` \| `busy` \| `offline` | The state of a CICD worker / agent.
  |





## container

### container.cpu.time

```py
# TYPE container.cpu.time COUNTER
container.cpu.time{cpu.mode}
```

Total CPU time consumed


|Attribute|Type|Description|
|-|-|-|
| cpu.mode | `user` \| `system` \| `nice` \| `idle` \| `iowait` \| `interrupt` \| `steal` \| `kernel` | The CPU mode for this data point. A container's CPU metric SHOULD be characterized _either_ by data points with no `mode` labels, _or only_ data points with `mode` labels.  |




### container.cpu.usage

```py
# TYPE container.cpu.usage GAUGE
container.cpu.usage{cpu.mode}
```

Container's CPU usage, measured in cpus. Range from 0 to the number of allocatable CPUs


|Attribute|Type|Description|
|-|-|-|
| cpu.mode | `user` \| `system` \| `nice` \| `idle` \| `iowait` \| `interrupt` \| `steal` \| `kernel` | The CPU mode for this data point. A container's CPU metric SHOULD be characterized _either_ by data points with no `mode` labels, _or only_ data points with `mode` labels.  |




### container.disk.io

```py
# TYPE container.disk.io COUNTER
container.disk.io{disk.io.direction, system.device}
```

Disk bytes for the container.


|Attribute|Type|Description|
|-|-|-|
| disk.io.direction | `read` \| `write` | The disk IO operation direction.  |
| system.device | string | The device identifier  |




### container.memory.usage

```py
# TYPE container.memory.usage COUNTER
container.memory.usage{}
```

Memory usage of the container.




### container.network.io

```py
# TYPE container.network.io COUNTER
container.network.io{network.io.direction, network.interface.name}
```

Network bytes for the container.


|Attribute|Type|Description|
|-|-|-|
| network.io.direction | `transmit` \| `receive` | The network IO operation direction.  |
| network.interface.name | string | The network interface name.  |




### container.uptime

```py
# TYPE container.uptime GAUGE
container.uptime{}
```

The time the container has been running





## cpu

### cpu.frequency

```py
# TYPE cpu.frequency GAUGE
cpu.frequency{}
```

Deprecated. Use `system.cpu.frequency` instead.




### cpu.time

```py
# TYPE cpu.time COUNTER
cpu.time{cpu.logical_number, cpu.mode}
```

Deprecated. Use `system.cpu.time` instead.


|Attribute|Type|Description|
|-|-|-|
| cpu.logical_number | int | The logical CPU number [0..n-1]  |
| cpu.mode | `user` \| `system` \| `nice` \| `idle` \| `iowait` \| `interrupt` \| `steal` \| `kernel` | The mode of the CPU  |




### cpu.utilization

```py
# TYPE cpu.utilization GAUGE
cpu.utilization{cpu.logical_number, cpu.mode}
```

Deprecated. Use `system.cpu.utilization` instead.


|Attribute|Type|Description|
|-|-|-|
| cpu.logical_number | int | The logical CPU number [0..n-1]  |
| cpu.mode | `user` \| `system` \| `nice` \| `idle` \| `iowait` \| `interrupt` \| `steal` \| `kernel` | The mode of the CPU  |





## cpython

### cpython.gc.collected_objects

```py
# TYPE cpython.gc.collected_objects COUNTER
cpython.gc.collected_objects{cpython.gc.generation}
```

The total number of objects collected inside a generation since interpreter start.


|Attribute|Type|Description|
|-|-|-|
| cpython.gc.generation | 0 \| 1 \| 2 | Value of the garbage collector collection generation.  |




### cpython.gc.collections

```py
# TYPE cpython.gc.collections COUNTER
cpython.gc.collections{cpython.gc.generation}
```

The number of times a generation was collected since interpreter start.


|Attribute|Type|Description|
|-|-|-|
| cpython.gc.generation | 0 \| 1 \| 2 | Value of the garbage collector collection generation.  |




### cpython.gc.uncollectable_objects

```py
# TYPE cpython.gc.uncollectable_objects COUNTER
cpython.gc.uncollectable_objects{cpython.gc.generation}
```

The total number of objects which were found to be uncollectable inside a generation since interpreter start.


|Attribute|Type|Description|
|-|-|-|
| cpython.gc.generation | 0 \| 1 \| 2 | Value of the garbage collector collection generation.  |





## db

### db.client.connection.count

```py
# TYPE db.client.connection.count UPDOWNCOUNTER
db.client.connection.count{db.client.connection.state, db.client.connection.pool.name}
```

The number of connections that are currently in state described by the `state` attribute


|Attribute|Type|Description|
|-|-|-|
| db.client.connection.state | `idle` \| `used` | The state of a connection in the pool  |
| db.client.connection.pool.name | string | The name of the connection pool; unique within the instrumented application. In case the connection pool implementation doesn't provide a name, instrumentation SHOULD use a combination of parameters that would make the name unique, for example, combining attributes `server.address`, `server.port`, and `db.namespace`, formatted as `server.address:server.port/db.namespace`. Instrumentations that generate connection pool name following different patterns SHOULD document it.
  |




### db.client.connection.create_time

```py
# TYPE db.client.connection.create_time HISTOGRAM
db.client.connection.create_time{db.client.connection.pool.name}
```

The time it took to create a new connection


|Attribute|Type|Description|
|-|-|-|
| db.client.connection.pool.name | string | The name of the connection pool; unique within the instrumented application. In case the connection pool implementation doesn't provide a name, instrumentation SHOULD use a combination of parameters that would make the name unique, for example, combining attributes `server.address`, `server.port`, and `db.namespace`, formatted as `server.address:server.port/db.namespace`. Instrumentations that generate connection pool name following different patterns SHOULD document it.
  |




### db.client.connection.idle.max

```py
# TYPE db.client.connection.idle.max UPDOWNCOUNTER
db.client.connection.idle.max{db.client.connection.pool.name}
```

The maximum number of idle open connections allowed


|Attribute|Type|Description|
|-|-|-|
| db.client.connection.pool.name | string | The name of the connection pool; unique within the instrumented application. In case the connection pool implementation doesn't provide a name, instrumentation SHOULD use a combination of parameters that would make the name unique, for example, combining attributes `server.address`, `server.port`, and `db.namespace`, formatted as `server.address:server.port/db.namespace`. Instrumentations that generate connection pool name following different patterns SHOULD document it.
  |




### db.client.connection.idle.min

```py
# TYPE db.client.connection.idle.min UPDOWNCOUNTER
db.client.connection.idle.min{db.client.connection.pool.name}
```

The minimum number of idle open connections allowed


|Attribute|Type|Description|
|-|-|-|
| db.client.connection.pool.name | string | The name of the connection pool; unique within the instrumented application. In case the connection pool implementation doesn't provide a name, instrumentation SHOULD use a combination of parameters that would make the name unique, for example, combining attributes `server.address`, `server.port`, and `db.namespace`, formatted as `server.address:server.port/db.namespace`. Instrumentations that generate connection pool name following different patterns SHOULD document it.
  |




### db.client.connection.max

```py
# TYPE db.client.connection.max UPDOWNCOUNTER
db.client.connection.max{db.client.connection.pool.name}
```

The maximum number of open connections allowed


|Attribute|Type|Description|
|-|-|-|
| db.client.connection.pool.name | string | The name of the connection pool; unique within the instrumented application. In case the connection pool implementation doesn't provide a name, instrumentation SHOULD use a combination of parameters that would make the name unique, for example, combining attributes `server.address`, `server.port`, and `db.namespace`, formatted as `server.address:server.port/db.namespace`. Instrumentations that generate connection pool name following different patterns SHOULD document it.
  |




### db.client.connection.pending_requests

```py
# TYPE db.client.connection.pending_requests UPDOWNCOUNTER
db.client.connection.pending_requests{db.client.connection.pool.name}
```

The number of current pending requests for an open connection


|Attribute|Type|Description|
|-|-|-|
| db.client.connection.pool.name | string | The name of the connection pool; unique within the instrumented application. In case the connection pool implementation doesn't provide a name, instrumentation SHOULD use a combination of parameters that would make the name unique, for example, combining attributes `server.address`, `server.port`, and `db.namespace`, formatted as `server.address:server.port/db.namespace`. Instrumentations that generate connection pool name following different patterns SHOULD document it.
  |




### db.client.connection.timeouts

```py
# TYPE db.client.connection.timeouts COUNTER
db.client.connection.timeouts{db.client.connection.pool.name}
```

The number of connection timeouts that have occurred trying to obtain a connection from the pool


|Attribute|Type|Description|
|-|-|-|
| db.client.connection.pool.name | string | The name of the connection pool; unique within the instrumented application. In case the connection pool implementation doesn't provide a name, instrumentation SHOULD use a combination of parameters that would make the name unique, for example, combining attributes `server.address`, `server.port`, and `db.namespace`, formatted as `server.address:server.port/db.namespace`. Instrumentations that generate connection pool name following different patterns SHOULD document it.
  |




### db.client.connection.use_time

```py
# TYPE db.client.connection.use_time HISTOGRAM
db.client.connection.use_time{db.client.connection.pool.name}
```

The time between borrowing a connection and returning it to the pool


|Attribute|Type|Description|
|-|-|-|
| db.client.connection.pool.name | string | The name of the connection pool; unique within the instrumented application. In case the connection pool implementation doesn't provide a name, instrumentation SHOULD use a combination of parameters that would make the name unique, for example, combining attributes `server.address`, `server.port`, and `db.namespace`, formatted as `server.address:server.port/db.namespace`. Instrumentations that generate connection pool name following different patterns SHOULD document it.
  |




### db.client.connection.wait_time

```py
# TYPE db.client.connection.wait_time HISTOGRAM
db.client.connection.wait_time{db.client.connection.pool.name}
```

The time it took to obtain an open connection from the pool


|Attribute|Type|Description|
|-|-|-|
| db.client.connection.pool.name | string | The name of the connection pool; unique within the instrumented application. In case the connection pool implementation doesn't provide a name, instrumentation SHOULD use a combination of parameters that would make the name unique, for example, combining attributes `server.address`, `server.port`, and `db.namespace`, formatted as `server.address:server.port/db.namespace`. Instrumentations that generate connection pool name following different patterns SHOULD document it.
  |




### db.client.connections.create_time

```py
# TYPE db.client.connections.create_time HISTOGRAM
db.client.connections.create_time{db.client.connections.pool.name}
```

Deprecated, use `db.client.connection.create_time` instead. Note: the unit also changed from `ms` to `s`.


|Attribute|Type|Description|
|-|-|-|
| db.client.connections.pool.name | string | Deprecated, use `db.client.connection.pool.name` instead.  |




### db.client.connections.idle.max

```py
# TYPE db.client.connections.idle.max UPDOWNCOUNTER
db.client.connections.idle.max{db.client.connections.pool.name}
```

Deprecated, use `db.client.connection.idle.max` instead.


|Attribute|Type|Description|
|-|-|-|
| db.client.connections.pool.name | string | Deprecated, use `db.client.connection.pool.name` instead.  |




### db.client.connections.idle.min

```py
# TYPE db.client.connections.idle.min UPDOWNCOUNTER
db.client.connections.idle.min{db.client.connections.pool.name}
```

Deprecated, use `db.client.connection.idle.min` instead.


|Attribute|Type|Description|
|-|-|-|
| db.client.connections.pool.name | string | Deprecated, use `db.client.connection.pool.name` instead.  |




### db.client.connections.max

```py
# TYPE db.client.connections.max UPDOWNCOUNTER
db.client.connections.max{db.client.connections.pool.name}
```

Deprecated, use `db.client.connection.max` instead.


|Attribute|Type|Description|
|-|-|-|
| db.client.connections.pool.name | string | Deprecated, use `db.client.connection.pool.name` instead.  |




### db.client.connections.pending_requests

```py
# TYPE db.client.connections.pending_requests UPDOWNCOUNTER
db.client.connections.pending_requests{db.client.connections.pool.name}
```

Deprecated, use `db.client.connection.pending_requests` instead.


|Attribute|Type|Description|
|-|-|-|
| db.client.connections.pool.name | string | Deprecated, use `db.client.connection.pool.name` instead.  |




### db.client.connections.timeouts

```py
# TYPE db.client.connections.timeouts COUNTER
db.client.connections.timeouts{db.client.connections.pool.name}
```

Deprecated, use `db.client.connection.timeouts` instead.


|Attribute|Type|Description|
|-|-|-|
| db.client.connections.pool.name | string | Deprecated, use `db.client.connection.pool.name` instead.  |




### db.client.connections.usage

```py
# TYPE db.client.connections.usage UPDOWNCOUNTER
db.client.connections.usage{db.client.connections.state, db.client.connections.pool.name}
```

Deprecated, use `db.client.connection.count` instead.


|Attribute|Type|Description|
|-|-|-|
| db.client.connections.state | `idle` \| `used` | Deprecated, use `db.client.connection.state` instead.  |
| db.client.connections.pool.name | string | Deprecated, use `db.client.connection.pool.name` instead.  |




### db.client.connections.use_time

```py
# TYPE db.client.connections.use_time HISTOGRAM
db.client.connections.use_time{db.client.connections.pool.name}
```

Deprecated, use `db.client.connection.use_time` instead. Note: the unit also changed from `ms` to `s`.


|Attribute|Type|Description|
|-|-|-|
| db.client.connections.pool.name | string | Deprecated, use `db.client.connection.pool.name` instead.  |




### db.client.connections.wait_time

```py
# TYPE db.client.connections.wait_time HISTOGRAM
db.client.connections.wait_time{db.client.connections.pool.name}
```

Deprecated, use `db.client.connection.wait_time` instead. Note: the unit also changed from `ms` to `s`.


|Attribute|Type|Description|
|-|-|-|
| db.client.connections.pool.name | string | Deprecated, use `db.client.connection.pool.name` instead.  |




### db.client.cosmosdb.active_instance.count

```py
# TYPE db.client.cosmosdb.active_instance.count UPDOWNCOUNTER
db.client.cosmosdb.active_instance.count{server.address, server.port}
```

Deprecated, use `azure.cosmosdb.client.active_instance.count` instead.


|Attribute|Type|Description|
|-|-|-|
| server.address | string | Name of the database host.
  |
| server.port | int | Server port number.  |




### db.client.cosmosdb.operation.request_charge

```py
# TYPE db.client.cosmosdb.operation.request_charge HISTOGRAM
db.client.cosmosdb.operation.request_charge{db.namespace, db.collection.name, db.operation.name, db.cosmosdb.sub_status_code, db.cosmosdb.consistency_level, db.cosmosdb.regions_contacted}
```

Deprecated, use `azure.cosmosdb.client.operation.request_charge` instead.


|Attribute|Type|Description|
|-|-|-|
| db.namespace | string | The name of the database, fully qualified within the server address and port.
  |
| db.collection.name | string | Cosmos DB container name.  |
| db.operation.name | string | The name of the operation or command being executed.
  |
| db.cosmosdb.sub_status_code | int | Deprecated, use `azure.cosmosdb.response.sub_status_code` instead.  |
| db.cosmosdb.consistency_level | `Strong` \| `BoundedStaleness` \| `Session` \| `Eventual` \| `ConsistentPrefix` | Deprecated, use `cosmosdb.consistency.level` instead.  |
| db.cosmosdb.regions_contacted | string[] | Deprecated, use `azure.cosmosdb.operation.contacted_regions` instead.  |




### db.client.operation.duration

```py
# TYPE db.client.operation.duration HISTOGRAM
db.client.operation.duration{db.response.status_code, error.type, db.stored_procedure.name, network.peer.address, db.collection.name, db.namespace, db.operation.name, db.query.summary, db.query.text, db.system.name, network.peer.port, server.address, server.port}
```

Duration of database client operations.


|Attribute|Type|Description|
|-|-|-|
| db.response.status_code | string | Database response status code.  |
| error.type | `_OTHER` | Describes a class of error the operation ended with.
  |
| db.stored_procedure.name | string | The name of a stored procedure within the database.  |
| network.peer.address | string | Peer address of the database node where the operation was performed.  |
| db.collection.name | string | The name of a collection (table, container) within the database.  |
| db.namespace | string | The name of the database, fully qualified within the server address and port.
  |
| db.operation.name | string | The name of the operation or command being executed.
  |
| db.query.summary | string | Low cardinality summary of a database query.
  |
| db.query.text | string | The database query being executed.
  |
| db.system.name | `other_sql` \| `softwareag.adabas` \| `actian.ingres` \| `aws.dynamodb` \| `aws.redshift` \| `azure.cosmosdb` \| `intersystems.cache` \| `cassandra` \| `clickhouse` \| `cockroachdb` \| `couchbase` \| `couchdb` \| `derby` \| `elasticsearch` \| `firebirdsql` \| `gcp.spanner` \| `geode` \| `h2database` \| `hbase` \| `hive` \| `hsqldb` \| `ibm.db2` \| `ibm.informix` \| `ibm.netezza` \| `influxdb` \| `instantdb` \| `mariadb` \| `memcached` \| `mongodb` \| `microsoft.sql_server` \| `mysql` \| `neo4j` \| `opensearch` \| `oracle.db` \| `postgresql` \| `redis` \| `sap.hana` \| `sap.maxdb` \| `sqlite` \| `teradata` \| `trino` | The database management system (DBMS) product as identified by the client instrumentation.  |
| network.peer.port | int | Peer port number of the network connection.  |
| server.address | string | Name of the database host.
  |
| server.port | int | Server port number.  |




### db.client.response.returned_rows

```py
# TYPE db.client.response.returned_rows HISTOGRAM
db.client.response.returned_rows{db.response.status_code, error.type, network.peer.address, db.collection.name, db.namespace, db.operation.name, db.query.summary, db.query.text, db.system.name, network.peer.port, server.address, server.port}
```

The actual number of records returned by the database operation.


|Attribute|Type|Description|
|-|-|-|
| db.response.status_code | string | Database response status code.  |
| error.type | `_OTHER` | Describes a class of error the operation ended with.
  |
| network.peer.address | string | Peer address of the database node where the operation was performed.  |
| db.collection.name | string | The name of a collection (table, container) within the database.  |
| db.namespace | string | The name of the database, fully qualified within the server address and port.
  |
| db.operation.name | string | The name of the operation or command being executed.
  |
| db.query.summary | string | Low cardinality summary of a database query.
  |
| db.query.text | string | The database query being executed.
  |
| db.system.name | `other_sql` \| `softwareag.adabas` \| `actian.ingres` \| `aws.dynamodb` \| `aws.redshift` \| `azure.cosmosdb` \| `intersystems.cache` \| `cassandra` \| `clickhouse` \| `cockroachdb` \| `couchbase` \| `couchdb` \| `derby` \| `elasticsearch` \| `firebirdsql` \| `gcp.spanner` \| `geode` \| `h2database` \| `hbase` \| `hive` \| `hsqldb` \| `ibm.db2` \| `ibm.informix` \| `ibm.netezza` \| `influxdb` \| `instantdb` \| `mariadb` \| `memcached` \| `mongodb` \| `microsoft.sql_server` \| `mysql` \| `neo4j` \| `opensearch` \| `oracle.db` \| `postgresql` \| `redis` \| `sap.hana` \| `sap.maxdb` \| `sqlite` \| `teradata` \| `trino` | The database management system (DBMS) product as identified by the client instrumentation.  |
| network.peer.port | int | Peer port number of the network connection.  |
| server.address | string | Name of the database host.
  |
| server.port | int | Server port number.  |





## dns

### dns.lookup.duration

```py
# TYPE dns.lookup.duration HISTOGRAM
dns.lookup.duration{dns.question.name, error.type}
```

Measures the time taken to perform a DNS lookup.


|Attribute|Type|Description|
|-|-|-|
| dns.question.name | string | The name being queried.  |
| error.type | `_OTHER` | Describes the error the DNS lookup failed with.  |





## dotnet

### dotnet.assembly.count

```py
# TYPE dotnet.assembly.count UPDOWNCOUNTER
dotnet.assembly.count{}
```

The number of .NET assemblies that are currently loaded.




### dotnet.exceptions

```py
# TYPE dotnet.exceptions COUNTER
dotnet.exceptions{error.type}
```

The number of exceptions that have been thrown in managed code.


|Attribute|Type|Description|
|-|-|-|
| error.type | `_OTHER` | Describes a class of error the operation ended with.
  |




### dotnet.gc.collections

```py
# TYPE dotnet.gc.collections COUNTER
dotnet.gc.collections{dotnet.gc.heap.generation}
```

The number of garbage collections that have occurred since the process has started.


|Attribute|Type|Description|
|-|-|-|
| dotnet.gc.heap.generation | `gen0` \| `gen1` \| `gen2` \| `loh` \| `poh` | Name of the garbage collector managed heap generation.  |




### dotnet.gc.heap.total_allocated

```py
# TYPE dotnet.gc.heap.total_allocated COUNTER
dotnet.gc.heap.total_allocated{}
```

The *approximate* number of bytes allocated on the managed GC heap since the process has started. The returned value does not include any native allocations.





### dotnet.gc.last_collection.heap.fragmentation.size

```py
# TYPE dotnet.gc.last_collection.heap.fragmentation.size UPDOWNCOUNTER
dotnet.gc.last_collection.heap.fragmentation.size{dotnet.gc.heap.generation}
```

The heap fragmentation, as observed during the latest garbage collection.



|Attribute|Type|Description|
|-|-|-|
| dotnet.gc.heap.generation | `gen0` \| `gen1` \| `gen2` \| `loh` \| `poh` | Name of the garbage collector managed heap generation.  |




### dotnet.gc.last_collection.heap.size

```py
# TYPE dotnet.gc.last_collection.heap.size UPDOWNCOUNTER
dotnet.gc.last_collection.heap.size{dotnet.gc.heap.generation}
```

The managed GC heap size (including fragmentation), as observed during the latest garbage collection.



|Attribute|Type|Description|
|-|-|-|
| dotnet.gc.heap.generation | `gen0` \| `gen1` \| `gen2` \| `loh` \| `poh` | Name of the garbage collector managed heap generation.  |




### dotnet.gc.last_collection.memory.committed_size

```py
# TYPE dotnet.gc.last_collection.memory.committed_size UPDOWNCOUNTER
dotnet.gc.last_collection.memory.committed_size{}
```

The amount of committed virtual memory in use by the .NET GC, as observed during the latest garbage collection.





### dotnet.gc.pause.time

```py
# TYPE dotnet.gc.pause.time COUNTER
dotnet.gc.pause.time{}
```

The total amount of time paused in GC since the process has started.




### dotnet.jit.compilation.time

```py
# TYPE dotnet.jit.compilation.time COUNTER
dotnet.jit.compilation.time{}
```

The amount of time the JIT compiler has spent compiling methods since the process has started.





### dotnet.jit.compiled_il.size

```py
# TYPE dotnet.jit.compiled_il.size COUNTER
dotnet.jit.compiled_il.size{}
```

Count of bytes of intermediate language that have been compiled since the process has started.




### dotnet.jit.compiled_methods

```py
# TYPE dotnet.jit.compiled_methods COUNTER
dotnet.jit.compiled_methods{}
```

The number of times the JIT compiler (re)compiled methods since the process has started.





### dotnet.monitor.lock_contentions

```py
# TYPE dotnet.monitor.lock_contentions COUNTER
dotnet.monitor.lock_contentions{}
```

The number of times there was contention when trying to acquire a monitor lock since the process has started.





### dotnet.process.cpu.count

```py
# TYPE dotnet.process.cpu.count UPDOWNCOUNTER
dotnet.process.cpu.count{}
```

The number of processors available to the process.




### dotnet.process.cpu.time

```py
# TYPE dotnet.process.cpu.time COUNTER
dotnet.process.cpu.time{cpu.mode}
```

CPU time used by the process.


|Attribute|Type|Description|
|-|-|-|
| cpu.mode | `user` \| `system` \| `nice` \| `idle` \| `iowait` \| `interrupt` \| `steal` \| `kernel` | The mode of the CPU  |




### dotnet.process.memory.working_set

```py
# TYPE dotnet.process.memory.working_set UPDOWNCOUNTER
dotnet.process.memory.working_set{}
```

The number of bytes of physical memory mapped to the process context.




### dotnet.thread_pool.queue.length

```py
# TYPE dotnet.thread_pool.queue.length UPDOWNCOUNTER
dotnet.thread_pool.queue.length{}
```

The number of work items that are currently queued to be processed by the thread pool.





### dotnet.thread_pool.thread.count

```py
# TYPE dotnet.thread_pool.thread.count UPDOWNCOUNTER
dotnet.thread_pool.thread.count{}
```

The number of thread pool threads that currently exist.




### dotnet.thread_pool.work_item.count

```py
# TYPE dotnet.thread_pool.work_item.count COUNTER
dotnet.thread_pool.work_item.count{}
```

The number of work items that the thread pool has completed since the process has started.





### dotnet.timer.count

```py
# TYPE dotnet.timer.count UPDOWNCOUNTER
dotnet.timer.count{}
```

The number of timer instances that are currently active.





## faas

### faas.coldstarts

```py
# TYPE faas.coldstarts COUNTER
faas.coldstarts{faas.trigger}
```

Number of invocation cold starts


|Attribute|Type|Description|
|-|-|-|
| faas.trigger | `datasource` \| `http` \| `pubsub` \| `timer` \| `other` | Type of the trigger which caused this function invocation.
  |




### faas.cpu_usage

```py
# TYPE faas.cpu_usage HISTOGRAM
faas.cpu_usage{faas.trigger}
```

Distribution of CPU usage per invocation


|Attribute|Type|Description|
|-|-|-|
| faas.trigger | `datasource` \| `http` \| `pubsub` \| `timer` \| `other` | Type of the trigger which caused this function invocation.
  |




### faas.errors

```py
# TYPE faas.errors COUNTER
faas.errors{faas.trigger}
```

Number of invocation errors


|Attribute|Type|Description|
|-|-|-|
| faas.trigger | `datasource` \| `http` \| `pubsub` \| `timer` \| `other` | Type of the trigger which caused this function invocation.
  |




### faas.init_duration

```py
# TYPE faas.init_duration HISTOGRAM
faas.init_duration{faas.trigger}
```

Measures the duration of the function's initialization, such as a cold start


|Attribute|Type|Description|
|-|-|-|
| faas.trigger | `datasource` \| `http` \| `pubsub` \| `timer` \| `other` | Type of the trigger which caused this function invocation.
  |




### faas.invocations

```py
# TYPE faas.invocations COUNTER
faas.invocations{faas.trigger}
```

Number of successful invocations


|Attribute|Type|Description|
|-|-|-|
| faas.trigger | `datasource` \| `http` \| `pubsub` \| `timer` \| `other` | Type of the trigger which caused this function invocation.
  |




### faas.invoke_duration

```py
# TYPE faas.invoke_duration HISTOGRAM
faas.invoke_duration{faas.trigger}
```

Measures the duration of the function's logic execution


|Attribute|Type|Description|
|-|-|-|
| faas.trigger | `datasource` \| `http` \| `pubsub` \| `timer` \| `other` | Type of the trigger which caused this function invocation.
  |




### faas.mem_usage

```py
# TYPE faas.mem_usage HISTOGRAM
faas.mem_usage{faas.trigger}
```

Distribution of max memory usage per invocation


|Attribute|Type|Description|
|-|-|-|
| faas.trigger | `datasource` \| `http` \| `pubsub` \| `timer` \| `other` | Type of the trigger which caused this function invocation.
  |




### faas.net_io

```py
# TYPE faas.net_io HISTOGRAM
faas.net_io{faas.trigger}
```

Distribution of net I/O usage per invocation


|Attribute|Type|Description|
|-|-|-|
| faas.trigger | `datasource` \| `http` \| `pubsub` \| `timer` \| `other` | Type of the trigger which caused this function invocation.
  |




### faas.timeouts

```py
# TYPE faas.timeouts COUNTER
faas.timeouts{faas.trigger}
```

Number of invocation timeouts


|Attribute|Type|Description|
|-|-|-|
| faas.trigger | `datasource` \| `http` \| `pubsub` \| `timer` \| `other` | Type of the trigger which caused this function invocation.
  |





## gen_ai

### gen_ai.client.operation.duration

```py
# TYPE gen_ai.client.operation.duration HISTOGRAM
gen_ai.client.operation.duration{gen_ai.response.model, gen_ai.operation.name, error.type, gen_ai.system, gen_ai.request.model, server.address, server.port}
```

GenAI operation duration


|Attribute|Type|Description|
|-|-|-|
| gen_ai.response.model | string | The name of the model that generated the response.  |
| gen_ai.operation.name | `chat` \| `generate_content` \| `text_completion` \| `embeddings` \| `create_agent` \| `invoke_agent` \| `execute_tool` | The name of the operation being performed.  |
| error.type | `_OTHER` | Describes a class of error the operation ended with.
  |
| gen_ai.system | `openai` \| `gcp.gen_ai` \| `gcp.vertex_ai` \| `gcp.gemini` \| `vertex_ai` \| `gemini` \| `anthropic` \| `cohere` \| `az.ai.inference` \| `az.ai.openai` \| `ibm.watsonx.ai` \| `aws.bedrock` \| `perplexity` \| `xai` \| `deepseek` \| `groq` \| `mistral_ai` | The Generative AI product as identified by the client or server instrumentation.  |
| gen_ai.request.model | string | The name of the GenAI model a request is being made to.  |
| server.address | string | GenAI server address.  |
| server.port | int | GenAI server port.  |




### gen_ai.client.token.usage

```py
# TYPE gen_ai.client.token.usage HISTOGRAM
gen_ai.client.token.usage{gen_ai.response.model, gen_ai.operation.name, gen_ai.system, gen_ai.request.model, gen_ai.token.type, server.address, server.port}
```

Measures number of input and output tokens used


|Attribute|Type|Description|
|-|-|-|
| gen_ai.response.model | string | The name of the model that generated the response.  |
| gen_ai.operation.name | `chat` \| `generate_content` \| `text_completion` \| `embeddings` \| `create_agent` \| `invoke_agent` \| `execute_tool` | The name of the operation being performed.  |
| gen_ai.system | `openai` \| `gcp.gen_ai` \| `gcp.vertex_ai` \| `gcp.gemini` \| `vertex_ai` \| `gemini` \| `anthropic` \| `cohere` \| `az.ai.inference` \| `az.ai.openai` \| `ibm.watsonx.ai` \| `aws.bedrock` \| `perplexity` \| `xai` \| `deepseek` \| `groq` \| `mistral_ai` | The Generative AI product as identified by the client or server instrumentation.  |
| gen_ai.request.model | string | The name of the GenAI model a request is being made to.  |
| gen_ai.token.type | `input` \| `output` \| `output` | The type of token being counted.  |
| server.address | string | GenAI server address.  |
| server.port | int | GenAI server port.  |




### gen_ai.server.request.duration

```py
# TYPE gen_ai.server.request.duration HISTOGRAM
gen_ai.server.request.duration{gen_ai.response.model, gen_ai.operation.name, gen_ai.system, gen_ai.request.model, error.type, server.address, server.port}
```

Generative AI server request duration such as time-to-last byte or last output token


|Attribute|Type|Description|
|-|-|-|
| gen_ai.response.model | string | The name of the model that generated the response.  |
| gen_ai.operation.name | `chat` \| `generate_content` \| `text_completion` \| `embeddings` \| `create_agent` \| `invoke_agent` \| `execute_tool` | The name of the operation being performed.  |
| gen_ai.system | `openai` \| `gcp.gen_ai` \| `gcp.vertex_ai` \| `gcp.gemini` \| `vertex_ai` \| `gemini` \| `anthropic` \| `cohere` \| `az.ai.inference` \| `az.ai.openai` \| `ibm.watsonx.ai` \| `aws.bedrock` \| `perplexity` \| `xai` \| `deepseek` \| `groq` \| `mistral_ai` | The Generative AI product as identified by the client or server instrumentation.  |
| gen_ai.request.model | string | The name of the GenAI model a request is being made to.  |
| error.type | `_OTHER` | Describes a class of error the operation ended with.
  |
| server.address | string | GenAI server address.  |
| server.port | int | GenAI server port.  |




### gen_ai.server.time_per_output_token

```py
# TYPE gen_ai.server.time_per_output_token HISTOGRAM
gen_ai.server.time_per_output_token{gen_ai.response.model, gen_ai.operation.name, gen_ai.system, gen_ai.request.model, server.address, server.port}
```

Time per output token generated after the first token for successful responses


|Attribute|Type|Description|
|-|-|-|
| gen_ai.response.model | string | The name of the model that generated the response.  |
| gen_ai.operation.name | `chat` \| `generate_content` \| `text_completion` \| `embeddings` \| `create_agent` \| `invoke_agent` \| `execute_tool` | The name of the operation being performed.  |
| gen_ai.system | `openai` \| `gcp.gen_ai` \| `gcp.vertex_ai` \| `gcp.gemini` \| `vertex_ai` \| `gemini` \| `anthropic` \| `cohere` \| `az.ai.inference` \| `az.ai.openai` \| `ibm.watsonx.ai` \| `aws.bedrock` \| `perplexity` \| `xai` \| `deepseek` \| `groq` \| `mistral_ai` | The Generative AI product as identified by the client or server instrumentation.  |
| gen_ai.request.model | string | The name of the GenAI model a request is being made to.  |
| server.address | string | GenAI server address.  |
| server.port | int | GenAI server port.  |




### gen_ai.server.time_to_first_token

```py
# TYPE gen_ai.server.time_to_first_token HISTOGRAM
gen_ai.server.time_to_first_token{gen_ai.response.model, gen_ai.operation.name, gen_ai.system, gen_ai.request.model, server.address, server.port}
```

Time to generate first token for successful responses


|Attribute|Type|Description|
|-|-|-|
| gen_ai.response.model | string | The name of the model that generated the response.  |
| gen_ai.operation.name | `chat` \| `generate_content` \| `text_completion` \| `embeddings` \| `create_agent` \| `invoke_agent` \| `execute_tool` | The name of the operation being performed.  |
| gen_ai.system | `openai` \| `gcp.gen_ai` \| `gcp.vertex_ai` \| `gcp.gemini` \| `vertex_ai` \| `gemini` \| `anthropic` \| `cohere` \| `az.ai.inference` \| `az.ai.openai` \| `ibm.watsonx.ai` \| `aws.bedrock` \| `perplexity` \| `xai` \| `deepseek` \| `groq` \| `mistral_ai` | The Generative AI product as identified by the client or server instrumentation.  |
| gen_ai.request.model | string | The name of the GenAI model a request is being made to.  |
| server.address | string | GenAI server address.  |
| server.port | int | GenAI server port.  |





## go

### go.config.gogc

```py
# TYPE go.config.gogc UPDOWNCOUNTER
go.config.gogc{}
```

Heap size target percentage configured by the user, otherwise 100.




### go.goroutine.count

```py
# TYPE go.goroutine.count UPDOWNCOUNTER
go.goroutine.count{}
```

Count of live goroutines.




### go.memory.allocated

```py
# TYPE go.memory.allocated COUNTER
go.memory.allocated{}
```

Memory allocated to the heap by the application.




### go.memory.allocations

```py
# TYPE go.memory.allocations COUNTER
go.memory.allocations{}
```

Count of allocations to the heap by the application.




### go.memory.gc.goal

```py
# TYPE go.memory.gc.goal UPDOWNCOUNTER
go.memory.gc.goal{}
```

Heap size target for the end of the GC cycle.




### go.memory.limit

```py
# TYPE go.memory.limit UPDOWNCOUNTER
go.memory.limit{}
```

Go runtime memory limit configured by the user, if a limit exists.




### go.memory.used

```py
# TYPE go.memory.used UPDOWNCOUNTER
go.memory.used{go.memory.type}
```

Memory used by the Go runtime.


|Attribute|Type|Description|
|-|-|-|
| go.memory.type | `stack` \| `other` | The type of memory.  |




### go.processor.limit

```py
# TYPE go.processor.limit UPDOWNCOUNTER
go.processor.limit{}
```

The number of OS threads that can execute user-level Go code simultaneously.




### go.schedule.duration

```py
# TYPE go.schedule.duration HISTOGRAM
go.schedule.duration{}
```

The time goroutines have spent in the scheduler in a runnable state before actually running.





## http

### http.client.active_requests

```py
# TYPE http.client.active_requests UPDOWNCOUNTER
http.client.active_requests{http.request.method, url.scheme, url.template, server.port, server.address}
```

Number of active HTTP requests.


|Attribute|Type|Description|
|-|-|-|
| http.request.method | `CONNECT` \| `DELETE` \| `GET` \| `HEAD` \| `OPTIONS` \| `PATCH` \| `POST` \| `PUT` \| `TRACE` \| `_OTHER` | HTTP request method.  |
| url.scheme | string | The [URI scheme](https://www.rfc-editor.org/rfc/rfc3986#section-3.1) component identifying the used protocol.
  |
| url.template | string | The low-cardinality template of an [absolute path reference](https://www.rfc-editor.org/rfc/rfc3986#section-4.2).
  |
| server.port | int | Port identifier of the ["URI origin"](https://www.rfc-editor.org/rfc/rfc9110.html#name-uri-origin) HTTP request is sent to.
  |
| server.address | string | Server domain name if available without reverse DNS lookup; otherwise, IP address or Unix domain socket name.  |




### http.client.connection.duration

```py
# TYPE http.client.connection.duration HISTOGRAM
http.client.connection.duration{network.peer.address, network.protocol.version, url.scheme, server.port, server.address}
```

The duration of the successfully established outbound HTTP connections.


|Attribute|Type|Description|
|-|-|-|
| network.peer.address | string | Peer address of the network connection - IP address or Unix domain socket name.  |
| network.protocol.version | string | The actual version of the protocol used for network communication.  |
| url.scheme | string | The [URI scheme](https://www.rfc-editor.org/rfc/rfc3986#section-3.1) component identifying the used protocol.
  |
| server.port | int | Port identifier of the ["URI origin"](https://www.rfc-editor.org/rfc/rfc9110.html#name-uri-origin) HTTP request is sent to.
  |
| server.address | string | Server domain name if available without reverse DNS lookup; otherwise, IP address or Unix domain socket name.  |




### http.client.open_connections

```py
# TYPE http.client.open_connections UPDOWNCOUNTER
http.client.open_connections{network.peer.address, network.protocol.version, url.scheme, http.connection.state, server.port, server.address}
```

Number of outbound HTTP connections that are currently active or idle on the client.


|Attribute|Type|Description|
|-|-|-|
| network.peer.address | string | Peer address of the network connection - IP address or Unix domain socket name.  |
| network.protocol.version | string | The actual version of the protocol used for network communication.  |
| url.scheme | string | The [URI scheme](https://www.rfc-editor.org/rfc/rfc3986#section-3.1) component identifying the used protocol.
  |
| http.connection.state | `active` \| `idle` | State of the HTTP connection in the HTTP connection pool.  |
| server.port | int | Port identifier of the ["URI origin"](https://www.rfc-editor.org/rfc/rfc9110.html#name-uri-origin) HTTP request is sent to.
  |
| server.address | string | Server domain name if available without reverse DNS lookup; otherwise, IP address or Unix domain socket name.  |




### http.client.request.body.size

```py
# TYPE http.client.request.body.size HISTOGRAM
http.client.request.body.size{error.type, http.response.status_code, network.protocol.name, network.protocol.version, url.scheme, http.request.method, url.template, server.address, server.port}
```

Size of HTTP client request bodies.


|Attribute|Type|Description|
|-|-|-|
| error.type | `_OTHER` | Describes a class of error the operation ended with.
  |
| http.response.status_code | int | [HTTP response status code](https://tools.ietf.org/html/rfc7231#section-6).  |
| network.protocol.name | string | [OSI application layer](https://wikipedia.org/wiki/Application_layer) or non-OSI equivalent.  |
| network.protocol.version | string | The actual version of the protocol used for network communication.  |
| url.scheme | string | The [URI scheme](https://www.rfc-editor.org/rfc/rfc3986#section-3.1) component identifying the used protocol.
  |
| http.request.method | `CONNECT` \| `DELETE` \| `GET` \| `HEAD` \| `OPTIONS` \| `PATCH` \| `POST` \| `PUT` \| `TRACE` \| `_OTHER` | HTTP request method.  |
| url.template | string | The low-cardinality template of an [absolute path reference](https://www.rfc-editor.org/rfc/rfc3986#section-4.2).
  |
| server.address | string | Host identifier of the ["URI origin"](https://www.rfc-editor.org/rfc/rfc9110.html#name-uri-origin) HTTP request is sent to.
  |
| server.port | int | Port identifier of the ["URI origin"](https://www.rfc-editor.org/rfc/rfc9110.html#name-uri-origin) HTTP request is sent to.
  |




### http.client.request.duration

```py
# TYPE http.client.request.duration HISTOGRAM
http.client.request.duration{error.type, http.response.status_code, network.protocol.name, network.protocol.version, url.scheme, url.template, http.request.method, server.address, server.port}
```

Duration of HTTP client requests.


|Attribute|Type|Description|
|-|-|-|
| error.type | `_OTHER` | Describes a class of error the operation ended with.
  |
| http.response.status_code | int | [HTTP response status code](https://tools.ietf.org/html/rfc7231#section-6).  |
| network.protocol.name | string | [OSI application layer](https://wikipedia.org/wiki/Application_layer) or non-OSI equivalent.  |
| network.protocol.version | string | The actual version of the protocol used for network communication.  |
| url.scheme | string | The [URI scheme](https://www.rfc-editor.org/rfc/rfc3986#section-3.1) component identifying the used protocol.
  |
| url.template | string | The low-cardinality template of an [absolute path reference](https://www.rfc-editor.org/rfc/rfc3986#section-4.2).
  |
| http.request.method | `CONNECT` \| `DELETE` \| `GET` \| `HEAD` \| `OPTIONS` \| `PATCH` \| `POST` \| `PUT` \| `TRACE` \| `_OTHER` | HTTP request method.  |
| server.address | string | Host identifier of the ["URI origin"](https://www.rfc-editor.org/rfc/rfc9110.html#name-uri-origin) HTTP request is sent to.
  |
| server.port | int | Port identifier of the ["URI origin"](https://www.rfc-editor.org/rfc/rfc9110.html#name-uri-origin) HTTP request is sent to.
  |




### http.client.response.body.size

```py
# TYPE http.client.response.body.size HISTOGRAM
http.client.response.body.size{error.type, http.response.status_code, network.protocol.name, network.protocol.version, url.scheme, http.request.method, url.template, server.address, server.port}
```

Size of HTTP client response bodies.


|Attribute|Type|Description|
|-|-|-|
| error.type | `_OTHER` | Describes a class of error the operation ended with.
  |
| http.response.status_code | int | [HTTP response status code](https://tools.ietf.org/html/rfc7231#section-6).  |
| network.protocol.name | string | [OSI application layer](https://wikipedia.org/wiki/Application_layer) or non-OSI equivalent.  |
| network.protocol.version | string | The actual version of the protocol used for network communication.  |
| url.scheme | string | The [URI scheme](https://www.rfc-editor.org/rfc/rfc3986#section-3.1) component identifying the used protocol.
  |
| http.request.method | `CONNECT` \| `DELETE` \| `GET` \| `HEAD` \| `OPTIONS` \| `PATCH` \| `POST` \| `PUT` \| `TRACE` \| `_OTHER` | HTTP request method.  |
| url.template | string | The low-cardinality template of an [absolute path reference](https://www.rfc-editor.org/rfc/rfc3986#section-4.2).
  |
| server.address | string | Host identifier of the ["URI origin"](https://www.rfc-editor.org/rfc/rfc9110.html#name-uri-origin) HTTP request is sent to.
  |
| server.port | int | Port identifier of the ["URI origin"](https://www.rfc-editor.org/rfc/rfc9110.html#name-uri-origin) HTTP request is sent to.
  |




### http.server.active_requests

```py
# TYPE http.server.active_requests UPDOWNCOUNTER
http.server.active_requests{http.request.method, url.scheme, server.address, server.port}
```

Number of active HTTP server requests.


|Attribute|Type|Description|
|-|-|-|
| http.request.method | `CONNECT` \| `DELETE` \| `GET` \| `HEAD` \| `OPTIONS` \| `PATCH` \| `POST` \| `PUT` \| `TRACE` \| `_OTHER` | HTTP request method.  |
| url.scheme | string | The [URI scheme](https://www.rfc-editor.org/rfc/rfc3986#section-3.1) component identifying the used protocol.
  |
| server.address | string | Name of the local HTTP server that received the request.
  |
| server.port | int | Port of the local HTTP server that received the request.
  |




### http.server.request.body.size

```py
# TYPE http.server.request.body.size HISTOGRAM
http.server.request.body.size{error.type, http.response.status_code, network.protocol.name, network.protocol.version, http.route, http.request.method, url.scheme, user_agent.synthetic.type, server.address, server.port}
```

Size of HTTP server request bodies.


|Attribute|Type|Description|
|-|-|-|
| error.type | `_OTHER` | Describes a class of error the operation ended with.
  |
| http.response.status_code | int | [HTTP response status code](https://tools.ietf.org/html/rfc7231#section-6).  |
| network.protocol.name | string | [OSI application layer](https://wikipedia.org/wiki/Application_layer) or non-OSI equivalent.  |
| network.protocol.version | string | The actual version of the protocol used for network communication.  |
| http.route | string | The matched route, that is, the path template in the format used by the respective server framework.
  |
| http.request.method | `CONNECT` \| `DELETE` \| `GET` \| `HEAD` \| `OPTIONS` \| `PATCH` \| `POST` \| `PUT` \| `TRACE` \| `_OTHER` | HTTP request method.  |
| url.scheme | string | The [URI scheme](https://www.rfc-editor.org/rfc/rfc3986#section-3.1) component identifying the used protocol.
  |
| user_agent.synthetic.type | `bot` \| `test` | Specifies the category of synthetic traffic, such as tests or bots.
  |
| server.address | string | Name of the local HTTP server that received the request.
  |
| server.port | int | Port of the local HTTP server that received the request.
  |




### http.server.request.duration

```py
# TYPE http.server.request.duration HISTOGRAM
http.server.request.duration{error.type, http.response.status_code, network.protocol.name, network.protocol.version, http.route, http.request.method, url.scheme, user_agent.synthetic.type, server.address, server.port}
```

Duration of HTTP server requests.


|Attribute|Type|Description|
|-|-|-|
| error.type | `_OTHER` | Describes a class of error the operation ended with.
  |
| http.response.status_code | int | [HTTP response status code](https://tools.ietf.org/html/rfc7231#section-6).  |
| network.protocol.name | string | [OSI application layer](https://wikipedia.org/wiki/Application_layer) or non-OSI equivalent.  |
| network.protocol.version | string | The actual version of the protocol used for network communication.  |
| http.route | string | The matched route, that is, the path template in the format used by the respective server framework.
  |
| http.request.method | `CONNECT` \| `DELETE` \| `GET` \| `HEAD` \| `OPTIONS` \| `PATCH` \| `POST` \| `PUT` \| `TRACE` \| `_OTHER` | HTTP request method.  |
| url.scheme | string | The [URI scheme](https://www.rfc-editor.org/rfc/rfc3986#section-3.1) component identifying the used protocol.
  |
| user_agent.synthetic.type | `bot` \| `test` | Specifies the category of synthetic traffic, such as tests or bots.
  |
| server.address | string | Name of the local HTTP server that received the request.
  |
| server.port | int | Port of the local HTTP server that received the request.
  |




### http.server.response.body.size

```py
# TYPE http.server.response.body.size HISTOGRAM
http.server.response.body.size{error.type, http.response.status_code, network.protocol.name, network.protocol.version, http.route, http.request.method, url.scheme, user_agent.synthetic.type, server.address, server.port}
```

Size of HTTP server response bodies.


|Attribute|Type|Description|
|-|-|-|
| error.type | `_OTHER` | Describes a class of error the operation ended with.
  |
| http.response.status_code | int | [HTTP response status code](https://tools.ietf.org/html/rfc7231#section-6).  |
| network.protocol.name | string | [OSI application layer](https://wikipedia.org/wiki/Application_layer) or non-OSI equivalent.  |
| network.protocol.version | string | The actual version of the protocol used for network communication.  |
| http.route | string | The matched route, that is, the path template in the format used by the respective server framework.
  |
| http.request.method | `CONNECT` \| `DELETE` \| `GET` \| `HEAD` \| `OPTIONS` \| `PATCH` \| `POST` \| `PUT` \| `TRACE` \| `_OTHER` | HTTP request method.  |
| url.scheme | string | The [URI scheme](https://www.rfc-editor.org/rfc/rfc3986#section-3.1) component identifying the used protocol.
  |
| user_agent.synthetic.type | `bot` \| `test` | Specifies the category of synthetic traffic, such as tests or bots.
  |
| server.address | string | Name of the local HTTP server that received the request.
  |
| server.port | int | Port of the local HTTP server that received the request.
  |





## hw

### hw.energy

```py
# TYPE hw.energy COUNTER
hw.energy{hw.name, hw.parent, hw.id, hw.type}
```

Energy consumed by the component


|Attribute|Type|Description|
|-|-|-|
| hw.name | string | An easily-recognizable name for the hardware component
  |
| hw.parent | string | Unique identifier of the parent component (typically the `hw.id` attribute of the enclosure, or disk controller)
  |
| hw.id | string | An identifier for the hardware component, unique within the monitored host
  |
| hw.type | `battery` \| `cpu` \| `disk_controller` \| `enclosure` \| `fan` \| `gpu` \| `logical_disk` \| `memory` \| `network` \| `physical_disk` \| `power_supply` \| `tape_drive` \| `temperature` \| `voltage` | Type of the component
  |




### hw.errors

```py
# TYPE hw.errors COUNTER
hw.errors{hw.name, hw.parent, hw.id, hw.type, error.type}
```

Number of errors encountered by the component


|Attribute|Type|Description|
|-|-|-|
| hw.name | string | An easily-recognizable name for the hardware component
  |
| hw.parent | string | Unique identifier of the parent component (typically the `hw.id` attribute of the enclosure, or disk controller)
  |
| hw.id | string | An identifier for the hardware component, unique within the monitored host
  |
| hw.type | `battery` \| `cpu` \| `disk_controller` \| `enclosure` \| `fan` \| `gpu` \| `logical_disk` \| `memory` \| `network` \| `physical_disk` \| `power_supply` \| `tape_drive` \| `temperature` \| `voltage` | Type of the component
  |
| error.type | `_OTHER` | The type of error encountered by the component  |




### hw.host.ambient_temperature

```py
# TYPE hw.host.ambient_temperature GAUGE
hw.host.ambient_temperature{hw.name, hw.parent, hw.id}
```

Ambient (external) temperature of the physical host


|Attribute|Type|Description|
|-|-|-|
| hw.name | string | An easily-recognizable name for the hardware component
  |
| hw.parent | string | Unique identifier of the parent component (typically the `hw.id` attribute of the enclosure, or disk controller)
  |
| hw.id | string | An identifier for the hardware component, unique within the monitored host
  |




### hw.host.energy

```py
# TYPE hw.host.energy COUNTER
hw.host.energy{hw.name, hw.parent, hw.id}
```

Total energy consumed by the entire physical host, in joules


|Attribute|Type|Description|
|-|-|-|
| hw.name | string | An easily-recognizable name for the hardware component
  |
| hw.parent | string | Unique identifier of the parent component (typically the `hw.id` attribute of the enclosure, or disk controller)
  |
| hw.id | string | An identifier for the hardware component, unique within the monitored host
  |




### hw.host.heating_margin

```py
# TYPE hw.host.heating_margin GAUGE
hw.host.heating_margin{hw.name, hw.parent, hw.id}
```

By how many degrees Celsius the temperature of the physical host can be increased, before reaching a warning threshold on one of the internal sensors



|Attribute|Type|Description|
|-|-|-|
| hw.name | string | An easily-recognizable name for the hardware component
  |
| hw.parent | string | Unique identifier of the parent component (typically the `hw.id` attribute of the enclosure, or disk controller)
  |
| hw.id | string | An identifier for the hardware component, unique within the monitored host
  |




### hw.host.power

```py
# TYPE hw.host.power GAUGE
hw.host.power{hw.name, hw.parent, hw.id}
```

Instantaneous power consumed by the entire physical host in Watts (`hw.host.energy` is preferred)



|Attribute|Type|Description|
|-|-|-|
| hw.name | string | An easily-recognizable name for the hardware component
  |
| hw.parent | string | Unique identifier of the parent component (typically the `hw.id` attribute of the enclosure, or disk controller)
  |
| hw.id | string | An identifier for the hardware component, unique within the monitored host
  |




### hw.power

```py
# TYPE hw.power GAUGE
hw.power{hw.name, hw.parent, hw.id, hw.type}
```

Instantaneous power consumed by the component


|Attribute|Type|Description|
|-|-|-|
| hw.name | string | An easily-recognizable name for the hardware component
  |
| hw.parent | string | Unique identifier of the parent component (typically the `hw.id` attribute of the enclosure, or disk controller)
  |
| hw.id | string | An identifier for the hardware component, unique within the monitored host
  |
| hw.type | `battery` \| `cpu` \| `disk_controller` \| `enclosure` \| `fan` \| `gpu` \| `logical_disk` \| `memory` \| `network` \| `physical_disk` \| `power_supply` \| `tape_drive` \| `temperature` \| `voltage` | Type of the component
  |




### hw.status

```py
# TYPE hw.status UPDOWNCOUNTER
hw.status{hw.name, hw.parent, hw.id, hw.type, hw.state}
```

Operational status: `1` (true) or `0` (false) for each of the possible states


|Attribute|Type|Description|
|-|-|-|
| hw.name | string | An easily-recognizable name for the hardware component
  |
| hw.parent | string | Unique identifier of the parent component (typically the `hw.id` attribute of the enclosure, or disk controller)
  |
| hw.id | string | An identifier for the hardware component, unique within the monitored host
  |
| hw.type | `battery` \| `cpu` \| `disk_controller` \| `enclosure` \| `fan` \| `gpu` \| `logical_disk` \| `memory` \| `network` \| `physical_disk` \| `power_supply` \| `tape_drive` \| `temperature` \| `voltage` | Type of the component
  |
| hw.state | `ok` \| `degraded` \| `failed` | The current state of the component
  |





## jvm

### jvm.buffer.count

```py
# TYPE jvm.buffer.count UPDOWNCOUNTER
jvm.buffer.count{jvm.buffer.pool.name}
```

Number of buffers in the pool.


|Attribute|Type|Description|
|-|-|-|
| jvm.buffer.pool.name | string | Name of the buffer pool.  |




### jvm.buffer.memory.limit

```py
# TYPE jvm.buffer.memory.limit UPDOWNCOUNTER
jvm.buffer.memory.limit{jvm.buffer.pool.name}
```

Measure of total memory capacity of buffers.


|Attribute|Type|Description|
|-|-|-|
| jvm.buffer.pool.name | string | Name of the buffer pool.  |




### jvm.buffer.memory.usage

```py
# TYPE jvm.buffer.memory.usage UPDOWNCOUNTER
jvm.buffer.memory.usage{jvm.buffer.pool.name}
```

Deprecated, use `jvm.buffer.memory.used` instead.


|Attribute|Type|Description|
|-|-|-|
| jvm.buffer.pool.name | string | Name of the buffer pool.  |




### jvm.buffer.memory.used

```py
# TYPE jvm.buffer.memory.used UPDOWNCOUNTER
jvm.buffer.memory.used{jvm.buffer.pool.name}
```

Measure of memory used by buffers.


|Attribute|Type|Description|
|-|-|-|
| jvm.buffer.pool.name | string | Name of the buffer pool.  |




### jvm.class.count

```py
# TYPE jvm.class.count UPDOWNCOUNTER
jvm.class.count{}
```

Number of classes currently loaded.




### jvm.class.loaded

```py
# TYPE jvm.class.loaded COUNTER
jvm.class.loaded{}
```

Number of classes loaded since JVM start.




### jvm.class.unloaded

```py
# TYPE jvm.class.unloaded COUNTER
jvm.class.unloaded{}
```

Number of classes unloaded since JVM start.




### jvm.cpu.count

```py
# TYPE jvm.cpu.count UPDOWNCOUNTER
jvm.cpu.count{}
```

Number of processors available to the Java virtual machine.




### jvm.cpu.recent_utilization

```py
# TYPE jvm.cpu.recent_utilization GAUGE
jvm.cpu.recent_utilization{}
```

Recent CPU utilization for the process as reported by the JVM.




### jvm.cpu.time

```py
# TYPE jvm.cpu.time COUNTER
jvm.cpu.time{}
```

CPU time used by the process as reported by the JVM.




### jvm.file_descriptor.count

```py
# TYPE jvm.file_descriptor.count UPDOWNCOUNTER
jvm.file_descriptor.count{}
```

Number of open file descriptors as reported by the JVM.




### jvm.gc.duration

```py
# TYPE jvm.gc.duration HISTOGRAM
jvm.gc.duration{jvm.gc.action, jvm.gc.name, jvm.gc.cause}
```

Duration of JVM garbage collection actions.


|Attribute|Type|Description|
|-|-|-|
| jvm.gc.action | string | Name of the garbage collector action.  |
| jvm.gc.name | string | Name of the garbage collector.  |
| jvm.gc.cause | string | Name of the garbage collector cause.  |




### jvm.memory.committed

```py
# TYPE jvm.memory.committed UPDOWNCOUNTER
jvm.memory.committed{jvm.memory.type, jvm.memory.pool.name}
```

Measure of memory committed.


|Attribute|Type|Description|
|-|-|-|
| jvm.memory.type | `heap` \| `non_heap` | The type of memory.  |
| jvm.memory.pool.name | string | Name of the memory pool.  |




### jvm.memory.init

```py
# TYPE jvm.memory.init UPDOWNCOUNTER
jvm.memory.init{jvm.memory.type, jvm.memory.pool.name}
```

Measure of initial memory requested.


|Attribute|Type|Description|
|-|-|-|
| jvm.memory.type | `heap` \| `non_heap` | The type of memory.  |
| jvm.memory.pool.name | string | Name of the memory pool.  |




### jvm.memory.limit

```py
# TYPE jvm.memory.limit UPDOWNCOUNTER
jvm.memory.limit{jvm.memory.type, jvm.memory.pool.name}
```

Measure of max obtainable memory.


|Attribute|Type|Description|
|-|-|-|
| jvm.memory.type | `heap` \| `non_heap` | The type of memory.  |
| jvm.memory.pool.name | string | Name of the memory pool.  |




### jvm.memory.used

```py
# TYPE jvm.memory.used UPDOWNCOUNTER
jvm.memory.used{jvm.memory.type, jvm.memory.pool.name}
```

Measure of memory used.


|Attribute|Type|Description|
|-|-|-|
| jvm.memory.type | `heap` \| `non_heap` | The type of memory.  |
| jvm.memory.pool.name | string | Name of the memory pool.  |




### jvm.memory.used_after_last_gc

```py
# TYPE jvm.memory.used_after_last_gc UPDOWNCOUNTER
jvm.memory.used_after_last_gc{jvm.memory.type, jvm.memory.pool.name}
```

Measure of memory used, as measured after the most recent garbage collection event on this pool.


|Attribute|Type|Description|
|-|-|-|
| jvm.memory.type | `heap` \| `non_heap` | The type of memory.  |
| jvm.memory.pool.name | string | Name of the memory pool.  |




### jvm.system.cpu.load_1m

```py
# TYPE jvm.system.cpu.load_1m GAUGE
jvm.system.cpu.load_1m{}
```

Average CPU load of the whole system for the last minute as reported by the JVM.




### jvm.system.cpu.utilization

```py
# TYPE jvm.system.cpu.utilization GAUGE
jvm.system.cpu.utilization{}
```

Recent CPU utilization for the whole system as reported by the JVM.




### jvm.thread.count

```py
# TYPE jvm.thread.count UPDOWNCOUNTER
jvm.thread.count{jvm.thread.daemon, jvm.thread.state}
```

Number of executing platform threads.


|Attribute|Type|Description|
|-|-|-|
| jvm.thread.daemon | boolean | Whether the thread is daemon or not.  |
| jvm.thread.state | `new` \| `runnable` \| `blocked` \| `waiting` \| `timed_waiting` \| `terminated` | State of the thread.  |





## k8s

### k8s.container.cpu.limit

```py
# TYPE k8s.container.cpu.limit GAUGE
k8s.container.cpu.limit{}
```

Maximum CPU resource limit set for the container




### k8s.container.cpu.request

```py
# TYPE k8s.container.cpu.request GAUGE
k8s.container.cpu.request{}
```

CPU resource requested for the container




### k8s.container.ephemeral_storage.limit

```py
# TYPE k8s.container.ephemeral_storage.limit GAUGE
k8s.container.ephemeral_storage.limit{}
```

Maximum ephemeral storage resource limit set for the container




### k8s.container.ephemeral_storage.request

```py
# TYPE k8s.container.ephemeral_storage.request GAUGE
k8s.container.ephemeral_storage.request{}
```

Ephemeral storage resource requested for the container




### k8s.container.memory.limit

```py
# TYPE k8s.container.memory.limit GAUGE
k8s.container.memory.limit{}
```

Maximum memory resource limit set for the container




### k8s.container.memory.request

```py
# TYPE k8s.container.memory.request GAUGE
k8s.container.memory.request{}
```

Memory resource requested for the container




### k8s.container.ready

```py
# TYPE k8s.container.ready UPDOWNCOUNTER
k8s.container.ready{}
```

Indicates whether the container is currently marked as ready to accept traffic, based on its readiness probe (1 = ready, 0 = not ready)





### k8s.container.restart.count

```py
# TYPE k8s.container.restart.count UPDOWNCOUNTER
k8s.container.restart.count{}
```

Describes how many times the container has restarted (since the last counter reset)




### k8s.container.status.reason

```py
# TYPE k8s.container.status.reason UPDOWNCOUNTER
k8s.container.status.reason{k8s.container.status.reason}
```

Describes the number of K8s containers that are currently in a state for a given reason


|Attribute|Type|Description|
|-|-|-|
| k8s.container.status.reason | `ContainerCreating` \| `CrashLoopBackOff` \| `CreateContainerConfigError` \| `ErrImagePull` \| `ImagePullBackOff` \| `OOMKilled` \| `Completed` \| `Error` \| `ContainerCannotRun` | The reason for the container state. Corresponds to the `reason` field of the: [K8s ContainerStateWaiting](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#containerstatewaiting-v1-core) or [K8s ContainerStateTerminated](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#containerstateterminated-v1-core)
  |




### k8s.container.status.state

```py
# TYPE k8s.container.status.state UPDOWNCOUNTER
k8s.container.status.state{k8s.container.status.state}
```

Describes the number of K8s containers that are currently in a given state


|Attribute|Type|Description|
|-|-|-|
| k8s.container.status.state | `terminated` \| `running` \| `waiting` | The state of the container. [K8s ContainerState](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#containerstate-v1-core)
  |




### k8s.container.storage.limit

```py
# TYPE k8s.container.storage.limit GAUGE
k8s.container.storage.limit{}
```

Maximum storage resource limit set for the container




### k8s.container.storage.request

```py
# TYPE k8s.container.storage.request GAUGE
k8s.container.storage.request{}
```

Storage resource requested for the container




### k8s.cronjob.active_jobs

```py
# TYPE k8s.cronjob.active_jobs UPDOWNCOUNTER
k8s.cronjob.active_jobs{}
```

The number of actively running jobs for a cronjob




### k8s.daemonset.current_scheduled_nodes

```py
# TYPE k8s.daemonset.current_scheduled_nodes UPDOWNCOUNTER
k8s.daemonset.current_scheduled_nodes{}
```

Number of nodes that are running at least 1 daemon pod and are supposed to run the daemon pod




### k8s.daemonset.desired_scheduled_nodes

```py
# TYPE k8s.daemonset.desired_scheduled_nodes UPDOWNCOUNTER
k8s.daemonset.desired_scheduled_nodes{}
```

Number of nodes that should be running the daemon pod (including nodes currently running the daemon pod)




### k8s.daemonset.misscheduled_nodes

```py
# TYPE k8s.daemonset.misscheduled_nodes UPDOWNCOUNTER
k8s.daemonset.misscheduled_nodes{}
```

Number of nodes that are running the daemon pod, but are not supposed to run the daemon pod




### k8s.daemonset.ready_nodes

```py
# TYPE k8s.daemonset.ready_nodes UPDOWNCOUNTER
k8s.daemonset.ready_nodes{}
```

Number of nodes that should be running the daemon pod and have one or more of the daemon pod running and ready




### k8s.deployment.available_pods

```py
# TYPE k8s.deployment.available_pods UPDOWNCOUNTER
k8s.deployment.available_pods{}
```

Total number of available replica pods (ready for at least minReadySeconds) targeted by this deployment




### k8s.deployment.desired_pods

```py
# TYPE k8s.deployment.desired_pods UPDOWNCOUNTER
k8s.deployment.desired_pods{}
```

Number of desired replica pods in this deployment




### k8s.hpa.current_pods

```py
# TYPE k8s.hpa.current_pods UPDOWNCOUNTER
k8s.hpa.current_pods{}
```

Current number of replica pods managed by this horizontal pod autoscaler, as last seen by the autoscaler




### k8s.hpa.desired_pods

```py
# TYPE k8s.hpa.desired_pods UPDOWNCOUNTER
k8s.hpa.desired_pods{}
```

Desired number of replica pods managed by this horizontal pod autoscaler, as last calculated by the autoscaler




### k8s.hpa.max_pods

```py
# TYPE k8s.hpa.max_pods UPDOWNCOUNTER
k8s.hpa.max_pods{}
```

The upper limit for the number of replica pods to which the autoscaler can scale up




### k8s.hpa.min_pods

```py
# TYPE k8s.hpa.min_pods UPDOWNCOUNTER
k8s.hpa.min_pods{}
```

The lower limit for the number of replica pods to which the autoscaler can scale down




### k8s.job.active_pods

```py
# TYPE k8s.job.active_pods UPDOWNCOUNTER
k8s.job.active_pods{}
```

The number of pending and actively running pods for a job




### k8s.job.desired_successful_pods

```py
# TYPE k8s.job.desired_successful_pods UPDOWNCOUNTER
k8s.job.desired_successful_pods{}
```

The desired number of successfully finished pods the job should be run with




### k8s.job.failed_pods

```py
# TYPE k8s.job.failed_pods UPDOWNCOUNTER
k8s.job.failed_pods{}
```

The number of pods which reached phase Failed for a job




### k8s.job.max_parallel_pods

```py
# TYPE k8s.job.max_parallel_pods UPDOWNCOUNTER
k8s.job.max_parallel_pods{}
```

The max desired number of pods the job should run at any given time




### k8s.job.successful_pods

```py
# TYPE k8s.job.successful_pods UPDOWNCOUNTER
k8s.job.successful_pods{}
```

The number of pods which reached phase Succeeded for a job




### k8s.namespace.phase

```py
# TYPE k8s.namespace.phase UPDOWNCOUNTER
k8s.namespace.phase{k8s.namespace.phase}
```

Describes number of K8s namespaces that are currently in a given phase.


|Attribute|Type|Description|
|-|-|-|
| k8s.namespace.phase | `active` \| `terminating` | The phase of the K8s namespace.
  |




### k8s.node.cpu.time

```py
# TYPE k8s.node.cpu.time COUNTER
k8s.node.cpu.time{}
```

Total CPU time consumed




### k8s.node.cpu.usage

```py
# TYPE k8s.node.cpu.usage GAUGE
k8s.node.cpu.usage{}
```

Node's CPU usage, measured in cpus. Range from 0 to the number of allocatable CPUs




### k8s.node.memory.usage

```py
# TYPE k8s.node.memory.usage GAUGE
k8s.node.memory.usage{}
```

Memory usage of the Node




### k8s.node.network.errors

```py
# TYPE k8s.node.network.errors COUNTER
k8s.node.network.errors{network.io.direction, network.interface.name}
```

Node network errors


|Attribute|Type|Description|
|-|-|-|
| network.io.direction | `transmit` \| `receive` | The network IO operation direction.  |
| network.interface.name | string | The network interface name.  |




### k8s.node.network.io

```py
# TYPE k8s.node.network.io COUNTER
k8s.node.network.io{network.io.direction, network.interface.name}
```

Network bytes for the Node


|Attribute|Type|Description|
|-|-|-|
| network.io.direction | `transmit` \| `receive` | The network IO operation direction.  |
| network.interface.name | string | The network interface name.  |




### k8s.node.uptime

```py
# TYPE k8s.node.uptime GAUGE
k8s.node.uptime{}
```

The time the Node has been running




### k8s.pod.cpu.time

```py
# TYPE k8s.pod.cpu.time COUNTER
k8s.pod.cpu.time{}
```

Total CPU time consumed




### k8s.pod.cpu.usage

```py
# TYPE k8s.pod.cpu.usage GAUGE
k8s.pod.cpu.usage{}
```

Pod's CPU usage, measured in cpus. Range from 0 to the number of allocatable CPUs




### k8s.pod.memory.usage

```py
# TYPE k8s.pod.memory.usage GAUGE
k8s.pod.memory.usage{}
```

Memory usage of the Pod




### k8s.pod.network.errors

```py
# TYPE k8s.pod.network.errors COUNTER
k8s.pod.network.errors{network.io.direction, network.interface.name}
```

Pod network errors


|Attribute|Type|Description|
|-|-|-|
| network.io.direction | `transmit` \| `receive` | The network IO operation direction.  |
| network.interface.name | string | The network interface name.  |




### k8s.pod.network.io

```py
# TYPE k8s.pod.network.io COUNTER
k8s.pod.network.io{network.io.direction, network.interface.name}
```

Network bytes for the Pod


|Attribute|Type|Description|
|-|-|-|
| network.io.direction | `transmit` \| `receive` | The network IO operation direction.  |
| network.interface.name | string | The network interface name.  |




### k8s.pod.uptime

```py
# TYPE k8s.pod.uptime GAUGE
k8s.pod.uptime{}
```

The time the Pod has been running




### k8s.replicaset.available_pods

```py
# TYPE k8s.replicaset.available_pods UPDOWNCOUNTER
k8s.replicaset.available_pods{}
```

Total number of available replica pods (ready for at least minReadySeconds) targeted by this replicaset




### k8s.replicaset.desired_pods

```py
# TYPE k8s.replicaset.desired_pods UPDOWNCOUNTER
k8s.replicaset.desired_pods{}
```

Number of desired replica pods in this replicaset




### k8s.replication_controller.available_pods

```py
# TYPE k8s.replication_controller.available_pods UPDOWNCOUNTER
k8s.replication_controller.available_pods{}
```

Deprecated, use `k8s.replicationcontroller.available_pods` instead.




### k8s.replication_controller.desired_pods

```py
# TYPE k8s.replication_controller.desired_pods UPDOWNCOUNTER
k8s.replication_controller.desired_pods{}
```

Deprecated, use `k8s.replicationcontroller.desired_pods` instead.




### k8s.replicationcontroller.available_pods

```py
# TYPE k8s.replicationcontroller.available_pods UPDOWNCOUNTER
k8s.replicationcontroller.available_pods{}
```

Total number of available replica pods (ready for at least minReadySeconds) targeted by this replication controller




### k8s.replicationcontroller.desired_pods

```py
# TYPE k8s.replicationcontroller.desired_pods UPDOWNCOUNTER
k8s.replicationcontroller.desired_pods{}
```

Number of desired replica pods in this replication controller




### k8s.statefulset.current_pods

```py
# TYPE k8s.statefulset.current_pods UPDOWNCOUNTER
k8s.statefulset.current_pods{}
```

The number of replica pods created by the statefulset controller from the statefulset version indicated by currentRevision




### k8s.statefulset.desired_pods

```py
# TYPE k8s.statefulset.desired_pods UPDOWNCOUNTER
k8s.statefulset.desired_pods{}
```

Number of desired replica pods in this statefulset




### k8s.statefulset.ready_pods

```py
# TYPE k8s.statefulset.ready_pods UPDOWNCOUNTER
k8s.statefulset.ready_pods{}
```

The number of replica pods created for this statefulset with a Ready Condition




### k8s.statefulset.updated_pods

```py
# TYPE k8s.statefulset.updated_pods UPDOWNCOUNTER
k8s.statefulset.updated_pods{}
```

Number of replica pods created by the statefulset controller from the statefulset version indicated by updateRevision





## kestrel

### kestrel.active_connections

```py
# TYPE kestrel.active_connections UPDOWNCOUNTER
kestrel.active_connections{server.address, server.port, network.type, network.transport}
```

Number of connections that are currently active on the server.


|Attribute|Type|Description|
|-|-|-|
| server.address | string | Server domain name if available without reverse DNS lookup; otherwise, IP address or Unix domain socket name.  |
| server.port | int | Server port number.  |
| network.type | `ipv4` \| `ipv6` | [OSI network layer](https://wikipedia.org/wiki/Network_layer) or non-OSI equivalent.  |
| network.transport | `tcp` \| `udp` \| `pipe` \| `unix` \| `quic` | [OSI transport layer](https://wikipedia.org/wiki/Transport_layer) or [inter-process communication method](https://wikipedia.org/wiki/Inter-process_communication).
  |




### kestrel.active_tls_handshakes

```py
# TYPE kestrel.active_tls_handshakes UPDOWNCOUNTER
kestrel.active_tls_handshakes{server.address, server.port, network.type, network.transport}
```

Number of TLS handshakes that are currently in progress on the server.


|Attribute|Type|Description|
|-|-|-|
| server.address | string | Server domain name if available without reverse DNS lookup; otherwise, IP address or Unix domain socket name.  |
| server.port | int | Server port number.  |
| network.type | `ipv4` \| `ipv6` | [OSI network layer](https://wikipedia.org/wiki/Network_layer) or non-OSI equivalent.  |
| network.transport | `tcp` \| `udp` \| `pipe` \| `unix` \| `quic` | [OSI transport layer](https://wikipedia.org/wiki/Transport_layer) or [inter-process communication method](https://wikipedia.org/wiki/Inter-process_communication).
  |




### kestrel.connection.duration

```py
# TYPE kestrel.connection.duration HISTOGRAM
kestrel.connection.duration{tls.protocol.version, network.protocol.version, server.address, server.port, network.type, network.transport, error.type, network.protocol.name}
```

The duration of connections on the server.


|Attribute|Type|Description|
|-|-|-|
| tls.protocol.version | string | Numeric part of the version parsed from the original string of the negotiated [SSL/TLS protocol version](https://docs.openssl.org/1.1.1/man3/SSL_get_version/#return-values)
  |
| network.protocol.version | string | The actual version of the protocol used for network communication.  |
| server.address | string | Server domain name if available without reverse DNS lookup; otherwise, IP address or Unix domain socket name.  |
| server.port | int | Server port number.  |
| network.type | `ipv4` \| `ipv6` | [OSI network layer](https://wikipedia.org/wiki/Network_layer) or non-OSI equivalent.  |
| network.transport | `tcp` \| `udp` \| `pipe` \| `unix` \| `quic` | [OSI transport layer](https://wikipedia.org/wiki/Transport_layer) or [inter-process communication method](https://wikipedia.org/wiki/Inter-process_communication).
  |
| error.type | `_OTHER` | The full name of exception type.  |
| network.protocol.name | string | [OSI application layer](https://wikipedia.org/wiki/Application_layer) or non-OSI equivalent.  |




### kestrel.queued_connections

```py
# TYPE kestrel.queued_connections UPDOWNCOUNTER
kestrel.queued_connections{server.address, server.port, network.type, network.transport}
```

Number of connections that are currently queued and are waiting to start.


|Attribute|Type|Description|
|-|-|-|
| server.address | string | Server domain name if available without reverse DNS lookup; otherwise, IP address or Unix domain socket name.  |
| server.port | int | Server port number.  |
| network.type | `ipv4` \| `ipv6` | [OSI network layer](https://wikipedia.org/wiki/Network_layer) or non-OSI equivalent.  |
| network.transport | `tcp` \| `udp` \| `pipe` \| `unix` \| `quic` | [OSI transport layer](https://wikipedia.org/wiki/Transport_layer) or [inter-process communication method](https://wikipedia.org/wiki/Inter-process_communication).
  |




### kestrel.queued_requests

```py
# TYPE kestrel.queued_requests UPDOWNCOUNTER
kestrel.queued_requests{network.protocol.version, server.address, server.port, network.type, network.transport, network.protocol.name}
```

Number of HTTP requests on multiplexed connections (HTTP/2 and HTTP/3) that are currently queued and are waiting to start.


|Attribute|Type|Description|
|-|-|-|
| network.protocol.version | string | The actual version of the protocol used for network communication.  |
| server.address | string | Server domain name if available without reverse DNS lookup; otherwise, IP address or Unix domain socket name.  |
| server.port | int | Server port number.  |
| network.type | `ipv4` \| `ipv6` | [OSI network layer](https://wikipedia.org/wiki/Network_layer) or non-OSI equivalent.  |
| network.transport | `tcp` \| `udp` \| `pipe` \| `unix` \| `quic` | [OSI transport layer](https://wikipedia.org/wiki/Transport_layer) or [inter-process communication method](https://wikipedia.org/wiki/Inter-process_communication).
  |
| network.protocol.name | string | [OSI application layer](https://wikipedia.org/wiki/Application_layer) or non-OSI equivalent.  |




### kestrel.rejected_connections

```py
# TYPE kestrel.rejected_connections COUNTER
kestrel.rejected_connections{server.address, server.port, network.type, network.transport}
```

Number of connections rejected by the server.


|Attribute|Type|Description|
|-|-|-|
| server.address | string | Server domain name if available without reverse DNS lookup; otherwise, IP address or Unix domain socket name.  |
| server.port | int | Server port number.  |
| network.type | `ipv4` \| `ipv6` | [OSI network layer](https://wikipedia.org/wiki/Network_layer) or non-OSI equivalent.  |
| network.transport | `tcp` \| `udp` \| `pipe` \| `unix` \| `quic` | [OSI transport layer](https://wikipedia.org/wiki/Transport_layer) or [inter-process communication method](https://wikipedia.org/wiki/Inter-process_communication).
  |




### kestrel.tls_handshake.duration

```py
# TYPE kestrel.tls_handshake.duration HISTOGRAM
kestrel.tls_handshake.duration{tls.protocol.version, server.address, server.port, network.type, network.transport, error.type}
```

The duration of TLS handshakes on the server.


|Attribute|Type|Description|
|-|-|-|
| tls.protocol.version | string | Numeric part of the version parsed from the original string of the negotiated [SSL/TLS protocol version](https://docs.openssl.org/1.1.1/man3/SSL_get_version/#return-values)
  |
| server.address | string | Server domain name if available without reverse DNS lookup; otherwise, IP address or Unix domain socket name.  |
| server.port | int | Server port number.  |
| network.type | `ipv4` \| `ipv6` | [OSI network layer](https://wikipedia.org/wiki/Network_layer) or non-OSI equivalent.  |
| network.transport | `tcp` \| `udp` \| `pipe` \| `unix` \| `quic` | [OSI transport layer](https://wikipedia.org/wiki/Transport_layer) or [inter-process communication method](https://wikipedia.org/wiki/Inter-process_communication).
  |
| error.type | `_OTHER` | The full name of exception type.  |




### kestrel.upgraded_connections

```py
# TYPE kestrel.upgraded_connections UPDOWNCOUNTER
kestrel.upgraded_connections{server.address, server.port, network.type, network.transport}
```

Number of connections that are currently upgraded (WebSockets). .


|Attribute|Type|Description|
|-|-|-|
| server.address | string | Server domain name if available without reverse DNS lookup; otherwise, IP address or Unix domain socket name.  |
| server.port | int | Server port number.  |
| network.type | `ipv4` \| `ipv6` | [OSI network layer](https://wikipedia.org/wiki/Network_layer) or non-OSI equivalent.  |
| network.transport | `tcp` \| `udp` \| `pipe` \| `unix` \| `quic` | [OSI transport layer](https://wikipedia.org/wiki/Transport_layer) or [inter-process communication method](https://wikipedia.org/wiki/Inter-process_communication).
  |





## messaging

### messaging.client.consumed.messages

```py
# TYPE messaging.client.consumed.messages COUNTER
messaging.client.consumed.messages{messaging.destination.partition.id, error.type, messaging.destination.name, messaging.destination.template, messaging.system, messaging.consumer.group.name, messaging.destination.subscription.name, messaging.operation.name, server.port, server.address}
```

Number of messages that were delivered to the application.


|Attribute|Type|Description|
|-|-|-|
| messaging.destination.partition.id | string | The identifier of the partition messages are sent to or received from, unique within the `messaging.destination.name`.
  |
| error.type | `_OTHER` | Describes a class of error the operation ended with.
  |
| messaging.destination.name | string | The message destination name  |
| messaging.destination.template | string | Low cardinality representation of the messaging destination name  |
| messaging.system | `activemq` \| `aws_sqs` \| `eventgrid` \| `eventhubs` \| `servicebus` \| `gcp_pubsub` \| `jms` \| `kafka` \| `rabbitmq` \| `rocketmq` \| `pulsar` | The messaging system as identified by the client instrumentation.  |
| messaging.consumer.group.name | string | The name of the consumer group with which a consumer is associated.
  |
| messaging.destination.subscription.name | string | The name of the destination subscription from which a message is consumed.  |
| messaging.operation.name | string | The system-specific name of the messaging operation.
  |
| server.port | int | Server port number.  |
| server.address | string | Server domain name if available without reverse DNS lookup; otherwise, IP address or Unix domain socket name.  |




### messaging.client.operation.duration

```py
# TYPE messaging.client.operation.duration HISTOGRAM
messaging.client.operation.duration{messaging.destination.partition.id, error.type, messaging.destination.name, messaging.destination.template, messaging.system, messaging.consumer.group.name, messaging.destination.subscription.name, messaging.operation.name, messaging.operation.type, server.port, server.address}
```

Duration of messaging operation initiated by a producer or consumer client.


|Attribute|Type|Description|
|-|-|-|
| messaging.destination.partition.id | string | The identifier of the partition messages are sent to or received from, unique within the `messaging.destination.name`.
  |
| error.type | `_OTHER` | Describes a class of error the operation ended with.
  |
| messaging.destination.name | string | The message destination name  |
| messaging.destination.template | string | Low cardinality representation of the messaging destination name  |
| messaging.system | `activemq` \| `aws_sqs` \| `eventgrid` \| `eventhubs` \| `servicebus` \| `gcp_pubsub` \| `jms` \| `kafka` \| `rabbitmq` \| `rocketmq` \| `pulsar` | The messaging system as identified by the client instrumentation.  |
| messaging.consumer.group.name | string | The name of the consumer group with which a consumer is associated.
  |
| messaging.destination.subscription.name | string | The name of the destination subscription from which a message is consumed.  |
| messaging.operation.name | string | The system-specific name of the messaging operation.
  |
| messaging.operation.type | `create` \| `send` \| `receive` \| `process` \| `settle` \| `deliver` \| `publish` | A string identifying the type of the messaging operation.
  |
| server.port | int | Server port number.  |
| server.address | string | Server domain name if available without reverse DNS lookup; otherwise, IP address or Unix domain socket name.  |




### messaging.client.published.messages

```py
# TYPE messaging.client.published.messages COUNTER
messaging.client.published.messages{messaging.destination.partition.id, error.type, messaging.operation.name, messaging.destination.name, messaging.destination.template, messaging.system, server.port, server.address}
```

Deprecated. Use `messaging.client.sent.messages` instead.


|Attribute|Type|Description|
|-|-|-|
| messaging.destination.partition.id | string | The identifier of the partition messages are sent to or received from, unique within the `messaging.destination.name`.
  |
| error.type | `_OTHER` | Describes a class of error the operation ended with.
  |
| messaging.operation.name | string | The system-specific name of the messaging operation.
  |
| messaging.destination.name | string | The message destination name  |
| messaging.destination.template | string | Low cardinality representation of the messaging destination name  |
| messaging.system | `activemq` \| `aws_sqs` \| `eventgrid` \| `eventhubs` \| `servicebus` \| `gcp_pubsub` \| `jms` \| `kafka` \| `rabbitmq` \| `rocketmq` \| `pulsar` | The messaging system as identified by the client instrumentation.  |
| server.port | int | Server port number.  |
| server.address | string | Server domain name if available without reverse DNS lookup; otherwise, IP address or Unix domain socket name.  |




### messaging.client.sent.messages

```py
# TYPE messaging.client.sent.messages COUNTER
messaging.client.sent.messages{messaging.destination.partition.id, error.type, messaging.destination.name, messaging.destination.template, messaging.system, messaging.operation.name, server.port, server.address}
```

Number of messages producer attempted to send to the broker.


|Attribute|Type|Description|
|-|-|-|
| messaging.destination.partition.id | string | The identifier of the partition messages are sent to or received from, unique within the `messaging.destination.name`.
  |
| error.type | `_OTHER` | Describes a class of error the operation ended with.
  |
| messaging.destination.name | string | The message destination name  |
| messaging.destination.template | string | Low cardinality representation of the messaging destination name  |
| messaging.system | `activemq` \| `aws_sqs` \| `eventgrid` \| `eventhubs` \| `servicebus` \| `gcp_pubsub` \| `jms` \| `kafka` \| `rabbitmq` \| `rocketmq` \| `pulsar` | The messaging system as identified by the client instrumentation.  |
| messaging.operation.name | string | The system-specific name of the messaging operation.
  |
| server.port | int | Server port number.  |
| server.address | string | Server domain name if available without reverse DNS lookup; otherwise, IP address or Unix domain socket name.  |




### messaging.process.duration

```py
# TYPE messaging.process.duration HISTOGRAM
messaging.process.duration{messaging.destination.partition.id, error.type, messaging.destination.name, messaging.destination.template, messaging.system, messaging.consumer.group.name, messaging.destination.subscription.name, messaging.operation.name, server.port, server.address}
```

Duration of processing operation.


|Attribute|Type|Description|
|-|-|-|
| messaging.destination.partition.id | string | The identifier of the partition messages are sent to or received from, unique within the `messaging.destination.name`.
  |
| error.type | `_OTHER` | Describes a class of error the operation ended with.
  |
| messaging.destination.name | string | The message destination name  |
| messaging.destination.template | string | Low cardinality representation of the messaging destination name  |
| messaging.system | `activemq` \| `aws_sqs` \| `eventgrid` \| `eventhubs` \| `servicebus` \| `gcp_pubsub` \| `jms` \| `kafka` \| `rabbitmq` \| `rocketmq` \| `pulsar` | The messaging system as identified by the client instrumentation.  |
| messaging.consumer.group.name | string | The name of the consumer group with which a consumer is associated.
  |
| messaging.destination.subscription.name | string | The name of the destination subscription from which a message is consumed.  |
| messaging.operation.name | string | The system-specific name of the messaging operation.
  |
| server.port | int | Server port number.  |
| server.address | string | Server domain name if available without reverse DNS lookup; otherwise, IP address or Unix domain socket name.  |




### messaging.process.messages

```py
# TYPE messaging.process.messages COUNTER
messaging.process.messages{error.type, messaging.operation.name, server.port, server.address}
```

Deprecated. Use `messaging.client.consumed.messages` instead.


|Attribute|Type|Description|
|-|-|-|
| error.type | `_OTHER` | Describes a class of error the operation ended with.
  |
| messaging.operation.name | string | The system-specific name of the messaging operation.
  |
| server.port | int | Server port number.  |
| server.address | string | Server domain name if available without reverse DNS lookup; otherwise, IP address or Unix domain socket name.  |




### messaging.publish.duration

```py
# TYPE messaging.publish.duration HISTOGRAM
messaging.publish.duration{error.type, messaging.operation.name, server.port, server.address}
```

Deprecated. Use `messaging.client.operation.duration` instead.


|Attribute|Type|Description|
|-|-|-|
| error.type | `_OTHER` | Describes a class of error the operation ended with.
  |
| messaging.operation.name | string | The system-specific name of the messaging operation.
  |
| server.port | int | Server port number.  |
| server.address | string | Server domain name if available without reverse DNS lookup; otherwise, IP address or Unix domain socket name.  |




### messaging.publish.messages

```py
# TYPE messaging.publish.messages COUNTER
messaging.publish.messages{error.type, messaging.operation.name, server.port, server.address}
```

Deprecated. Use `messaging.client.sent.messages` instead.


|Attribute|Type|Description|
|-|-|-|
| error.type | `_OTHER` | Describes a class of error the operation ended with.
  |
| messaging.operation.name | string | The system-specific name of the messaging operation.
  |
| server.port | int | Server port number.  |
| server.address | string | Server domain name if available without reverse DNS lookup; otherwise, IP address or Unix domain socket name.  |




### messaging.receive.duration

```py
# TYPE messaging.receive.duration HISTOGRAM
messaging.receive.duration{error.type, messaging.operation.name, server.port, server.address}
```

Deprecated. Use `messaging.client.operation.duration` instead.


|Attribute|Type|Description|
|-|-|-|
| error.type | `_OTHER` | Describes a class of error the operation ended with.
  |
| messaging.operation.name | string | The system-specific name of the messaging operation.
  |
| server.port | int | Server port number.  |
| server.address | string | Server domain name if available without reverse DNS lookup; otherwise, IP address or Unix domain socket name.  |




### messaging.receive.messages

```py
# TYPE messaging.receive.messages COUNTER
messaging.receive.messages{error.type, messaging.operation.name, server.port, server.address}
```

Deprecated. Use `messaging.client.consumed.messages` instead.


|Attribute|Type|Description|
|-|-|-|
| error.type | `_OTHER` | Describes a class of error the operation ended with.
  |
| messaging.operation.name | string | The system-specific name of the messaging operation.
  |
| server.port | int | Server port number.  |
| server.address | string | Server domain name if available without reverse DNS lookup; otherwise, IP address or Unix domain socket name.  |





## nodejs

### nodejs.eventloop.delay.max

```py
# TYPE nodejs.eventloop.delay.max GAUGE
nodejs.eventloop.delay.max{}
```

Event loop maximum delay.




### nodejs.eventloop.delay.mean

```py
# TYPE nodejs.eventloop.delay.mean GAUGE
nodejs.eventloop.delay.mean{}
```

Event loop mean delay.




### nodejs.eventloop.delay.min

```py
# TYPE nodejs.eventloop.delay.min GAUGE
nodejs.eventloop.delay.min{}
```

Event loop minimum delay.




### nodejs.eventloop.delay.p50

```py
# TYPE nodejs.eventloop.delay.p50 GAUGE
nodejs.eventloop.delay.p50{}
```

Event loop 50 percentile delay.




### nodejs.eventloop.delay.p90

```py
# TYPE nodejs.eventloop.delay.p90 GAUGE
nodejs.eventloop.delay.p90{}
```

Event loop 90 percentile delay.




### nodejs.eventloop.delay.p99

```py
# TYPE nodejs.eventloop.delay.p99 GAUGE
nodejs.eventloop.delay.p99{}
```

Event loop 99 percentile delay.




### nodejs.eventloop.delay.stddev

```py
# TYPE nodejs.eventloop.delay.stddev GAUGE
nodejs.eventloop.delay.stddev{}
```

Event loop standard deviation delay.




### nodejs.eventloop.time

```py
# TYPE nodejs.eventloop.time COUNTER
nodejs.eventloop.time{nodejs.eventloop.state}
```

Cumulative duration of time the event loop has been in each state.


|Attribute|Type|Description|
|-|-|-|
| nodejs.eventloop.state | `active` \| `idle` | The state of event loop time.  |




### nodejs.eventloop.utilization

```py
# TYPE nodejs.eventloop.utilization GAUGE
nodejs.eventloop.utilization{}
```

Event loop utilization.





## otel

### otel.sdk.exporter.log.exported

```py
# TYPE otel.sdk.exporter.log.exported COUNTER
otel.sdk.exporter.log.exported{otel.component.type, otel.component.name, error.type, server.address, server.port}
```

The number of log records for which the export has finished, either successful or failed


|Attribute|Type|Description|
|-|-|-|
| otel.component.type | `batching_span_processor` \| `simple_span_processor` \| `batching_log_processor` \| `simple_log_processor` \| `otlp_grpc_span_exporter` \| `otlp_http_span_exporter` \| `otlp_http_json_span_exporter` \| `otlp_grpc_log_exporter` \| `otlp_http_log_exporter` \| `otlp_http_json_log_exporter` \| `periodic_metric_reader` \| `otlp_grpc_metric_exporter` \| `otlp_http_metric_exporter` \| `otlp_http_json_metric_exporter` | A name identifying the type of the OpenTelemetry component.
  |
| otel.component.name | string | A name uniquely identifying the instance of the OpenTelemetry component within its containing SDK instance.
  |
| error.type | `_OTHER` | Describes a class of error the operation ended with.
  |
| server.address | string | Server domain name if available without reverse DNS lookup; otherwise, IP address or Unix domain socket name.  |
| server.port | int | Server port number.  |




### otel.sdk.exporter.log.inflight

```py
# TYPE otel.sdk.exporter.log.inflight UPDOWNCOUNTER
otel.sdk.exporter.log.inflight{otel.component.type, otel.component.name, server.address, server.port}
```

The number of log records which were passed to the exporter, but that have not been exported yet (neither successful, nor failed)


|Attribute|Type|Description|
|-|-|-|
| otel.component.type | `batching_span_processor` \| `simple_span_processor` \| `batching_log_processor` \| `simple_log_processor` \| `otlp_grpc_span_exporter` \| `otlp_http_span_exporter` \| `otlp_http_json_span_exporter` \| `otlp_grpc_log_exporter` \| `otlp_http_log_exporter` \| `otlp_http_json_log_exporter` \| `periodic_metric_reader` \| `otlp_grpc_metric_exporter` \| `otlp_http_metric_exporter` \| `otlp_http_json_metric_exporter` | A name identifying the type of the OpenTelemetry component.
  |
| otel.component.name | string | A name uniquely identifying the instance of the OpenTelemetry component within its containing SDK instance.
  |
| server.address | string | Server domain name if available without reverse DNS lookup; otherwise, IP address or Unix domain socket name.  |
| server.port | int | Server port number.  |




### otel.sdk.exporter.metric_data_point.exported

```py
# TYPE otel.sdk.exporter.metric_data_point.exported COUNTER
otel.sdk.exporter.metric_data_point.exported{otel.component.type, otel.component.name, error.type, server.address, server.port}
```

The number of metric data points for which the export has finished, either successful or failed


|Attribute|Type|Description|
|-|-|-|
| otel.component.type | `batching_span_processor` \| `simple_span_processor` \| `batching_log_processor` \| `simple_log_processor` \| `otlp_grpc_span_exporter` \| `otlp_http_span_exporter` \| `otlp_http_json_span_exporter` \| `otlp_grpc_log_exporter` \| `otlp_http_log_exporter` \| `otlp_http_json_log_exporter` \| `periodic_metric_reader` \| `otlp_grpc_metric_exporter` \| `otlp_http_metric_exporter` \| `otlp_http_json_metric_exporter` | A name identifying the type of the OpenTelemetry component.
  |
| otel.component.name | string | A name uniquely identifying the instance of the OpenTelemetry component within its containing SDK instance.
  |
| error.type | `_OTHER` | Describes a class of error the operation ended with.
  |
| server.address | string | Server domain name if available without reverse DNS lookup; otherwise, IP address or Unix domain socket name.  |
| server.port | int | Server port number.  |




### otel.sdk.exporter.metric_data_point.inflight

```py
# TYPE otel.sdk.exporter.metric_data_point.inflight UPDOWNCOUNTER
otel.sdk.exporter.metric_data_point.inflight{otel.component.type, otel.component.name, server.address, server.port}
```

The number of metric data points which were passed to the exporter, but that have not been exported yet (neither successful, nor failed)


|Attribute|Type|Description|
|-|-|-|
| otel.component.type | `batching_span_processor` \| `simple_span_processor` \| `batching_log_processor` \| `simple_log_processor` \| `otlp_grpc_span_exporter` \| `otlp_http_span_exporter` \| `otlp_http_json_span_exporter` \| `otlp_grpc_log_exporter` \| `otlp_http_log_exporter` \| `otlp_http_json_log_exporter` \| `periodic_metric_reader` \| `otlp_grpc_metric_exporter` \| `otlp_http_metric_exporter` \| `otlp_http_json_metric_exporter` | A name identifying the type of the OpenTelemetry component.
  |
| otel.component.name | string | A name uniquely identifying the instance of the OpenTelemetry component within its containing SDK instance.
  |
| server.address | string | Server domain name if available without reverse DNS lookup; otherwise, IP address or Unix domain socket name.  |
| server.port | int | Server port number.  |




### otel.sdk.exporter.operation.duration

```py
# TYPE otel.sdk.exporter.operation.duration HISTOGRAM
otel.sdk.exporter.operation.duration{otel.component.name, otel.component.type, error.type, http.response.status_code, server.address, server.port, rpc.grpc.status_code}
```

The duration of exporting a batch of telemetry records.


|Attribute|Type|Description|
|-|-|-|
| otel.component.name | string | A name uniquely identifying the instance of the OpenTelemetry component within its containing SDK instance.
  |
| otel.component.type | `batching_span_processor` \| `simple_span_processor` \| `batching_log_processor` \| `simple_log_processor` \| `otlp_grpc_span_exporter` \| `otlp_http_span_exporter` \| `otlp_http_json_span_exporter` \| `otlp_grpc_log_exporter` \| `otlp_http_log_exporter` \| `otlp_http_json_log_exporter` \| `periodic_metric_reader` \| `otlp_grpc_metric_exporter` \| `otlp_http_metric_exporter` \| `otlp_http_json_metric_exporter` | A name identifying the type of the OpenTelemetry component.
  |
| error.type | `_OTHER` | Describes a class of error the operation ended with.
  |
| http.response.status_code | int | The HTTP status code of the last HTTP request performed in scope of this export call.  |
| server.address | string | Server domain name if available without reverse DNS lookup; otherwise, IP address or Unix domain socket name.  |
| server.port | int | Server port number.  |
| rpc.grpc.status_code | 0 \| 1 \| 2 \| 3 \| 4 \| 5 \| 6 \| 7 \| 8 \| 9 \| 10 \| 11 \| 12 \| 13 \| 14 \| 15 \| 16 | The gRPC status code of the last gRPC requests performed in scope of this export call.  |




### otel.sdk.exporter.span.exported

```py
# TYPE otel.sdk.exporter.span.exported COUNTER
otel.sdk.exporter.span.exported{otel.component.type, otel.component.name, error.type, server.address, server.port}
```

The number of spans for which the export has finished, either successful or failed


|Attribute|Type|Description|
|-|-|-|
| otel.component.type | `batching_span_processor` \| `simple_span_processor` \| `batching_log_processor` \| `simple_log_processor` \| `otlp_grpc_span_exporter` \| `otlp_http_span_exporter` \| `otlp_http_json_span_exporter` \| `otlp_grpc_log_exporter` \| `otlp_http_log_exporter` \| `otlp_http_json_log_exporter` \| `periodic_metric_reader` \| `otlp_grpc_metric_exporter` \| `otlp_http_metric_exporter` \| `otlp_http_json_metric_exporter` | A name identifying the type of the OpenTelemetry component.
  |
| otel.component.name | string | A name uniquely identifying the instance of the OpenTelemetry component within its containing SDK instance.
  |
| error.type | `_OTHER` | Describes a class of error the operation ended with.
  |
| server.address | string | Server domain name if available without reverse DNS lookup; otherwise, IP address or Unix domain socket name.  |
| server.port | int | Server port number.  |




### otel.sdk.exporter.span.exported.count

```py
# TYPE otel.sdk.exporter.span.exported.count UPDOWNCOUNTER
otel.sdk.exporter.span.exported.count{}
```

Deprecated, use `otel.sdk.exporter.span.exported` instead.




### otel.sdk.exporter.span.inflight

```py
# TYPE otel.sdk.exporter.span.inflight UPDOWNCOUNTER
otel.sdk.exporter.span.inflight{otel.component.type, otel.component.name, server.address, server.port}
```

The number of spans which were passed to the exporter, but that have not been exported yet (neither successful, nor failed)


|Attribute|Type|Description|
|-|-|-|
| otel.component.type | `batching_span_processor` \| `simple_span_processor` \| `batching_log_processor` \| `simple_log_processor` \| `otlp_grpc_span_exporter` \| `otlp_http_span_exporter` \| `otlp_http_json_span_exporter` \| `otlp_grpc_log_exporter` \| `otlp_http_log_exporter` \| `otlp_http_json_log_exporter` \| `periodic_metric_reader` \| `otlp_grpc_metric_exporter` \| `otlp_http_metric_exporter` \| `otlp_http_json_metric_exporter` | A name identifying the type of the OpenTelemetry component.
  |
| otel.component.name | string | A name uniquely identifying the instance of the OpenTelemetry component within its containing SDK instance.
  |
| server.address | string | Server domain name if available without reverse DNS lookup; otherwise, IP address or Unix domain socket name.  |
| server.port | int | Server port number.  |




### otel.sdk.exporter.span.inflight.count

```py
# TYPE otel.sdk.exporter.span.inflight.count UPDOWNCOUNTER
otel.sdk.exporter.span.inflight.count{}
```

Deprecated, use `otel.sdk.exporter.span.inflight` instead.




### otel.sdk.log.created

```py
# TYPE otel.sdk.log.created COUNTER
otel.sdk.log.created{}
```

The number of logs submitted to enabled SDK Loggers




### otel.sdk.metric_reader.collection.duration

```py
# TYPE otel.sdk.metric_reader.collection.duration HISTOGRAM
otel.sdk.metric_reader.collection.duration{otel.component.type, otel.component.name, error.type}
```

The duration of the collect operation of the metric reader.


|Attribute|Type|Description|
|-|-|-|
| otel.component.type | `batching_span_processor` \| `simple_span_processor` \| `batching_log_processor` \| `simple_log_processor` \| `otlp_grpc_span_exporter` \| `otlp_http_span_exporter` \| `otlp_http_json_span_exporter` \| `otlp_grpc_log_exporter` \| `otlp_http_log_exporter` \| `otlp_http_json_log_exporter` \| `periodic_metric_reader` \| `otlp_grpc_metric_exporter` \| `otlp_http_metric_exporter` \| `otlp_http_json_metric_exporter` | A name identifying the type of the OpenTelemetry component.
  |
| otel.component.name | string | A name uniquely identifying the instance of the OpenTelemetry component within its containing SDK instance.
  |
| error.type | `_OTHER` | Describes a class of error the operation ended with.
  |




### otel.sdk.processor.log.processed

```py
# TYPE otel.sdk.processor.log.processed COUNTER
otel.sdk.processor.log.processed{otel.component.type, otel.component.name, error.type}
```

The number of log records for which the processing has finished, either successful or failed


|Attribute|Type|Description|
|-|-|-|
| otel.component.type | `batching_span_processor` \| `simple_span_processor` \| `batching_log_processor` \| `simple_log_processor` \| `otlp_grpc_span_exporter` \| `otlp_http_span_exporter` \| `otlp_http_json_span_exporter` \| `otlp_grpc_log_exporter` \| `otlp_http_log_exporter` \| `otlp_http_json_log_exporter` \| `periodic_metric_reader` \| `otlp_grpc_metric_exporter` \| `otlp_http_metric_exporter` \| `otlp_http_json_metric_exporter` | A name identifying the type of the OpenTelemetry component.
  |
| otel.component.name | string | A name uniquely identifying the instance of the OpenTelemetry component within its containing SDK instance.
  |
| error.type | `_OTHER` | A low-cardinality description of the failure reason. SDK Batching Log Record Processors MUST use `queue_full` for log records dropped due to a full queue.
  |




### otel.sdk.processor.log.queue.capacity

```py
# TYPE otel.sdk.processor.log.queue.capacity UPDOWNCOUNTER
otel.sdk.processor.log.queue.capacity{otel.component.type, otel.component.name}
```

The maximum number of log records the queue of a given instance of an SDK Log Record processor can hold


|Attribute|Type|Description|
|-|-|-|
| otel.component.type | `batching_span_processor` \| `simple_span_processor` \| `batching_log_processor` \| `simple_log_processor` \| `otlp_grpc_span_exporter` \| `otlp_http_span_exporter` \| `otlp_http_json_span_exporter` \| `otlp_grpc_log_exporter` \| `otlp_http_log_exporter` \| `otlp_http_json_log_exporter` \| `periodic_metric_reader` \| `otlp_grpc_metric_exporter` \| `otlp_http_metric_exporter` \| `otlp_http_json_metric_exporter` | A name identifying the type of the OpenTelemetry component.
  |
| otel.component.name | string | A name uniquely identifying the instance of the OpenTelemetry component within its containing SDK instance.
  |




### otel.sdk.processor.log.queue.size

```py
# TYPE otel.sdk.processor.log.queue.size UPDOWNCOUNTER
otel.sdk.processor.log.queue.size{otel.component.type, otel.component.name}
```

The number of log records in the queue of a given instance of an SDK log processor


|Attribute|Type|Description|
|-|-|-|
| otel.component.type | `batching_span_processor` \| `simple_span_processor` \| `batching_log_processor` \| `simple_log_processor` \| `otlp_grpc_span_exporter` \| `otlp_http_span_exporter` \| `otlp_http_json_span_exporter` \| `otlp_grpc_log_exporter` \| `otlp_http_log_exporter` \| `otlp_http_json_log_exporter` \| `periodic_metric_reader` \| `otlp_grpc_metric_exporter` \| `otlp_http_metric_exporter` \| `otlp_http_json_metric_exporter` | A name identifying the type of the OpenTelemetry component.
  |
| otel.component.name | string | A name uniquely identifying the instance of the OpenTelemetry component within its containing SDK instance.
  |




### otel.sdk.processor.span.processed

```py
# TYPE otel.sdk.processor.span.processed COUNTER
otel.sdk.processor.span.processed{otel.component.type, otel.component.name, error.type}
```

The number of spans for which the processing has finished, either successful or failed


|Attribute|Type|Description|
|-|-|-|
| otel.component.type | `batching_span_processor` \| `simple_span_processor` \| `batching_log_processor` \| `simple_log_processor` \| `otlp_grpc_span_exporter` \| `otlp_http_span_exporter` \| `otlp_http_json_span_exporter` \| `otlp_grpc_log_exporter` \| `otlp_http_log_exporter` \| `otlp_http_json_log_exporter` \| `periodic_metric_reader` \| `otlp_grpc_metric_exporter` \| `otlp_http_metric_exporter` \| `otlp_http_json_metric_exporter` | A name identifying the type of the OpenTelemetry component.
  |
| otel.component.name | string | A name uniquely identifying the instance of the OpenTelemetry component within its containing SDK instance.
  |
| error.type | `_OTHER` | A low-cardinality description of the failure reason. SDK Batching Span Processors MUST use `queue_full` for spans dropped due to a full queue.
  |




### otel.sdk.processor.span.processed.count

```py
# TYPE otel.sdk.processor.span.processed.count UPDOWNCOUNTER
otel.sdk.processor.span.processed.count{}
```

Deprecated, use `otel.sdk.processor.span.processed` instead.




### otel.sdk.processor.span.queue.capacity

```py
# TYPE otel.sdk.processor.span.queue.capacity UPDOWNCOUNTER
otel.sdk.processor.span.queue.capacity{otel.component.type, otel.component.name}
```

The maximum number of spans the queue of a given instance of an SDK span processor can hold


|Attribute|Type|Description|
|-|-|-|
| otel.component.type | `batching_span_processor` \| `simple_span_processor` \| `batching_log_processor` \| `simple_log_processor` \| `otlp_grpc_span_exporter` \| `otlp_http_span_exporter` \| `otlp_http_json_span_exporter` \| `otlp_grpc_log_exporter` \| `otlp_http_log_exporter` \| `otlp_http_json_log_exporter` \| `periodic_metric_reader` \| `otlp_grpc_metric_exporter` \| `otlp_http_metric_exporter` \| `otlp_http_json_metric_exporter` | A name identifying the type of the OpenTelemetry component.
  |
| otel.component.name | string | A name uniquely identifying the instance of the OpenTelemetry component within its containing SDK instance.
  |




### otel.sdk.processor.span.queue.size

```py
# TYPE otel.sdk.processor.span.queue.size UPDOWNCOUNTER
otel.sdk.processor.span.queue.size{otel.component.type, otel.component.name}
```

The number of spans in the queue of a given instance of an SDK span processor


|Attribute|Type|Description|
|-|-|-|
| otel.component.type | `batching_span_processor` \| `simple_span_processor` \| `batching_log_processor` \| `simple_log_processor` \| `otlp_grpc_span_exporter` \| `otlp_http_span_exporter` \| `otlp_http_json_span_exporter` \| `otlp_grpc_log_exporter` \| `otlp_http_log_exporter` \| `otlp_http_json_log_exporter` \| `periodic_metric_reader` \| `otlp_grpc_metric_exporter` \| `otlp_http_metric_exporter` \| `otlp_http_json_metric_exporter` | A name identifying the type of the OpenTelemetry component.
  |
| otel.component.name | string | A name uniquely identifying the instance of the OpenTelemetry component within its containing SDK instance.
  |




### otel.sdk.span.ended

```py
# TYPE otel.sdk.span.ended COUNTER
otel.sdk.span.ended{otel.span.sampling_result}
```

The number of created spans for which the end operation was called


|Attribute|Type|Description|
|-|-|-|
| otel.span.sampling_result | `DROP` \| `RECORD_ONLY` \| `RECORD_AND_SAMPLE` | The result value of the sampler for this span  |




### otel.sdk.span.ended.count

```py
# TYPE otel.sdk.span.ended.count COUNTER
otel.sdk.span.ended.count{}
```

Deprecated, use `otel.sdk.span.ended` instead.




### otel.sdk.span.live

```py
# TYPE otel.sdk.span.live UPDOWNCOUNTER
otel.sdk.span.live{otel.span.sampling_result}
```

The number of created spans for which the end operation has not been called yet


|Attribute|Type|Description|
|-|-|-|
| otel.span.sampling_result | `DROP` \| `RECORD_ONLY` \| `RECORD_AND_SAMPLE` | The result value of the sampler for this span  |




### otel.sdk.span.live.count

```py
# TYPE otel.sdk.span.live.count UPDOWNCOUNTER
otel.sdk.span.live.count{}
```

Deprecated, use `otel.sdk.span.live` instead.





## process

### process.context_switches

```py
# TYPE process.context_switches COUNTER
process.context_switches{process.context_switch_type}
```

Number of times the process has been context switched.


|Attribute|Type|Description|
|-|-|-|
| process.context_switch_type | `voluntary` \| `involuntary` | Specifies whether the context switches for this data point were voluntary or involuntary.  |




### process.cpu.time

```py
# TYPE process.cpu.time COUNTER
process.cpu.time{cpu.mode}
```

Total CPU seconds broken down by different states.


|Attribute|Type|Description|
|-|-|-|
| cpu.mode | `user` \| `system` \| `nice` \| `idle` \| `iowait` \| `interrupt` \| `steal` \| `kernel` | A process SHOULD be characterized _either_ by data points with no `mode` labels, _or only_ data points with `mode` labels.
  |




### process.cpu.utilization

```py
# TYPE process.cpu.utilization GAUGE
process.cpu.utilization{cpu.mode}
```

Difference in process.cpu.time since the last measurement, divided by the elapsed time and number of CPUs available to the process.


|Attribute|Type|Description|
|-|-|-|
| cpu.mode | `user` \| `system` \| `nice` \| `idle` \| `iowait` \| `interrupt` \| `steal` \| `kernel` | A process SHOULD be characterized _either_ by data points with no `mode` labels, _or only_ data points with `mode` labels.
  |




### process.disk.io

```py
# TYPE process.disk.io COUNTER
process.disk.io{disk.io.direction}
```

Disk bytes transferred.


|Attribute|Type|Description|
|-|-|-|
| disk.io.direction | `read` \| `write` | The disk IO operation direction.  |




### process.memory.usage

```py
# TYPE process.memory.usage UPDOWNCOUNTER
process.memory.usage{}
```

The amount of physical memory in use.




### process.memory.virtual

```py
# TYPE process.memory.virtual UPDOWNCOUNTER
process.memory.virtual{}
```

The amount of committed virtual memory.




### process.network.io

```py
# TYPE process.network.io COUNTER
process.network.io{network.io.direction}
```

Network bytes transferred.


|Attribute|Type|Description|
|-|-|-|
| network.io.direction | `transmit` \| `receive` | The network IO operation direction.  |




### process.open_file_descriptor.count

```py
# TYPE process.open_file_descriptor.count UPDOWNCOUNTER
process.open_file_descriptor.count{}
```

Number of file descriptors in use by the process.




### process.paging.faults

```py
# TYPE process.paging.faults COUNTER
process.paging.faults{process.paging.fault_type}
```

Number of page faults the process has made.


|Attribute|Type|Description|
|-|-|-|
| process.paging.fault_type | `major` \| `minor` | The type of page fault for this data point. Type `major` is for major/hard page faults, and `minor` is for minor/soft page faults.
  |




### process.thread.count

```py
# TYPE process.thread.count UPDOWNCOUNTER
process.thread.count{}
```

Process threads count.




### process.uptime

```py
# TYPE process.uptime GAUGE
process.uptime{}
```

The time the process has been running.





## rpc

### rpc.client.duration

```py
# TYPE rpc.client.duration HISTOGRAM
rpc.client.duration{}
```

Measures the duration of outbound RPC.




### rpc.client.request.size

```py
# TYPE rpc.client.request.size HISTOGRAM
rpc.client.request.size{}
```

Measures the size of RPC request messages (uncompressed).




### rpc.client.requests_per_rpc

```py
# TYPE rpc.client.requests_per_rpc HISTOGRAM
rpc.client.requests_per_rpc{}
```

Measures the number of messages received per RPC.




### rpc.client.response.size

```py
# TYPE rpc.client.response.size HISTOGRAM
rpc.client.response.size{}
```

Measures the size of RPC response messages (uncompressed).




### rpc.client.responses_per_rpc

```py
# TYPE rpc.client.responses_per_rpc HISTOGRAM
rpc.client.responses_per_rpc{}
```

Measures the number of messages sent per RPC.




### rpc.server.duration

```py
# TYPE rpc.server.duration HISTOGRAM
rpc.server.duration{}
```

Measures the duration of inbound RPC.




### rpc.server.request.size

```py
# TYPE rpc.server.request.size HISTOGRAM
rpc.server.request.size{}
```

Measures the size of RPC request messages (uncompressed).




### rpc.server.requests_per_rpc

```py
# TYPE rpc.server.requests_per_rpc HISTOGRAM
rpc.server.requests_per_rpc{}
```

Measures the number of messages received per RPC.




### rpc.server.response.size

```py
# TYPE rpc.server.response.size HISTOGRAM
rpc.server.response.size{}
```

Measures the size of RPC response messages (uncompressed).




### rpc.server.responses_per_rpc

```py
# TYPE rpc.server.responses_per_rpc HISTOGRAM
rpc.server.responses_per_rpc{}
```

Measures the number of messages sent per RPC.





## signalr

### signalr.server.active_connections

```py
# TYPE signalr.server.active_connections UPDOWNCOUNTER
signalr.server.active_connections{signalr.connection.status, signalr.transport}
```

Number of connections that are currently active on the server.


|Attribute|Type|Description|
|-|-|-|
| signalr.connection.status | `normal_closure` \| `timeout` \| `app_shutdown` | SignalR HTTP connection closure status.  |
| signalr.transport | `server_sent_events` \| `long_polling` \| `web_sockets` | [SignalR transport type](https://github.com/dotnet/aspnetcore/blob/main/src/SignalR/docs/specs/TransportProtocols.md)  |




### signalr.server.connection.duration

```py
# TYPE signalr.server.connection.duration HISTOGRAM
signalr.server.connection.duration{signalr.connection.status, signalr.transport}
```

The duration of connections on the server.


|Attribute|Type|Description|
|-|-|-|
| signalr.connection.status | `normal_closure` \| `timeout` \| `app_shutdown` | SignalR HTTP connection closure status.  |
| signalr.transport | `server_sent_events` \| `long_polling` \| `web_sockets` | [SignalR transport type](https://github.com/dotnet/aspnetcore/blob/main/src/SignalR/docs/specs/TransportProtocols.md)  |





## system

### system.cpu.frequency

```py
# TYPE system.cpu.frequency GAUGE
system.cpu.frequency{cpu.logical_number}
```

Operating frequency of the logical CPU in Hertz.


|Attribute|Type|Description|
|-|-|-|
| cpu.logical_number | int | The logical CPU number [0..n-1]  |




### system.cpu.logical.count

```py
# TYPE system.cpu.logical.count UPDOWNCOUNTER
system.cpu.logical.count{}
```

Reports the number of logical (virtual) processor cores created by the operating system to manage multitasking




### system.cpu.physical.count

```py
# TYPE system.cpu.physical.count UPDOWNCOUNTER
system.cpu.physical.count{}
```

Reports the number of actual physical processor cores on the hardware




### system.cpu.time

```py
# TYPE system.cpu.time COUNTER
system.cpu.time{cpu.logical_number, cpu.mode}
```

Seconds each logical CPU spent on each mode


|Attribute|Type|Description|
|-|-|-|
| cpu.logical_number | int | The logical CPU number [0..n-1]  |
| cpu.mode | `user` \| `system` \| `nice` \| `idle` \| `iowait` \| `interrupt` \| `steal` \| `kernel` | The mode of the CPU  |




### system.cpu.utilization

```py
# TYPE system.cpu.utilization GAUGE
system.cpu.utilization{cpu.logical_number, cpu.mode}
```

For each logical CPU, the utilization is calculated as the change in cumulative CPU time (cpu.time) over a measurement interval, divided by the elapsed time.


|Attribute|Type|Description|
|-|-|-|
| cpu.logical_number | int | The logical CPU number [0..n-1]  |
| cpu.mode | `user` \| `system` \| `nice` \| `idle` \| `iowait` \| `interrupt` \| `steal` \| `kernel` | The mode of the CPU  |




### system.disk.io

```py
# TYPE system.disk.io COUNTER
system.disk.io{disk.io.direction, system.device}
```




|Attribute|Type|Description|
|-|-|-|
| disk.io.direction | `read` \| `write` | The disk IO operation direction.  |
| system.device | string | The device identifier  |




### system.disk.io_time

```py
# TYPE system.disk.io_time COUNTER
system.disk.io_time{system.device}
```

Time disk spent activated


|Attribute|Type|Description|
|-|-|-|
| system.device | string | The device identifier  |




### system.disk.limit

```py
# TYPE system.disk.limit UPDOWNCOUNTER
system.disk.limit{system.device}
```

The total storage capacity of the disk


|Attribute|Type|Description|
|-|-|-|
| system.device | string | The device identifier  |




### system.disk.merged

```py
# TYPE system.disk.merged COUNTER
system.disk.merged{disk.io.direction, system.device}
```




|Attribute|Type|Description|
|-|-|-|
| disk.io.direction | `read` \| `write` | The disk IO operation direction.  |
| system.device | string | The device identifier  |




### system.disk.operation_time

```py
# TYPE system.disk.operation_time COUNTER
system.disk.operation_time{disk.io.direction, system.device}
```

Sum of the time each operation took to complete


|Attribute|Type|Description|
|-|-|-|
| disk.io.direction | `read` \| `write` | The disk IO operation direction.  |
| system.device | string | The device identifier  |




### system.disk.operations

```py
# TYPE system.disk.operations COUNTER
system.disk.operations{disk.io.direction, system.device}
```




|Attribute|Type|Description|
|-|-|-|
| disk.io.direction | `read` \| `write` | The disk IO operation direction.  |
| system.device | string | The device identifier  |




### system.filesystem.limit

```py
# TYPE system.filesystem.limit UPDOWNCOUNTER
system.filesystem.limit{system.filesystem.type, system.filesystem.mode, system.filesystem.mountpoint, system.device}
```

The total storage capacity of the filesystem


|Attribute|Type|Description|
|-|-|-|
| system.filesystem.type | `fat32` \| `exfat` \| `ntfs` \| `refs` \| `hfsplus` \| `ext4` | The filesystem type  |
| system.filesystem.mode | string | The filesystem mode  |
| system.filesystem.mountpoint | string | The filesystem mount path  |
| system.device | string | Identifier for the device where the filesystem resides.  |




### system.filesystem.usage

```py
# TYPE system.filesystem.usage UPDOWNCOUNTER
system.filesystem.usage{system.filesystem.state, system.filesystem.type, system.filesystem.mode, system.filesystem.mountpoint, system.device}
```

Reports a filesystem's space usage across different states.


|Attribute|Type|Description|
|-|-|-|
| system.filesystem.state | `used` \| `free` \| `reserved` | The filesystem state  |
| system.filesystem.type | `fat32` \| `exfat` \| `ntfs` \| `refs` \| `hfsplus` \| `ext4` | The filesystem type  |
| system.filesystem.mode | string | The filesystem mode  |
| system.filesystem.mountpoint | string | The filesystem mount path  |
| system.device | string | Identifier for the device where the filesystem resides.  |




### system.filesystem.utilization

```py
# TYPE system.filesystem.utilization GAUGE
system.filesystem.utilization{system.filesystem.state, system.filesystem.type, system.filesystem.mode, system.filesystem.mountpoint, system.device}
```




|Attribute|Type|Description|
|-|-|-|
| system.filesystem.state | `used` \| `free` \| `reserved` | The filesystem state  |
| system.filesystem.type | `fat32` \| `exfat` \| `ntfs` \| `refs` \| `hfsplus` \| `ext4` | The filesystem type  |
| system.filesystem.mode | string | The filesystem mode  |
| system.filesystem.mountpoint | string | The filesystem mount path  |
| system.device | string | Identifier for the device where the filesystem resides.  |




### system.linux.memory.available

```py
# TYPE system.linux.memory.available UPDOWNCOUNTER
system.linux.memory.available{}
```

An estimate of how much memory is available for starting new applications, without causing swapping




### system.linux.memory.slab.usage

```py
# TYPE system.linux.memory.slab.usage UPDOWNCOUNTER
system.linux.memory.slab.usage{linux.memory.slab.state}
```

Reports the memory used by the Linux kernel for managing caches of frequently used objects.


|Attribute|Type|Description|
|-|-|-|
| linux.memory.slab.state | `reclaimable` \| `unreclaimable` | The Linux Slab memory state  |




### system.memory.limit

```py
# TYPE system.memory.limit UPDOWNCOUNTER
system.memory.limit{}
```

Total memory available in the system.




### system.memory.shared

```py
# TYPE system.memory.shared UPDOWNCOUNTER
system.memory.shared{}
```

Shared memory used (mostly by tmpfs).




### system.memory.usage

```py
# TYPE system.memory.usage UPDOWNCOUNTER
system.memory.usage{system.memory.state}
```

Reports memory in use by state.


|Attribute|Type|Description|
|-|-|-|
| system.memory.state | `used` \| `free` \| `shared` \| `buffers` \| `cached` | The memory state  |




### system.memory.utilization

```py
# TYPE system.memory.utilization GAUGE
system.memory.utilization{system.memory.state}
```




|Attribute|Type|Description|
|-|-|-|
| system.memory.state | `used` \| `free` \| `shared` \| `buffers` \| `cached` | The memory state  |




### system.network.connections

```py
# TYPE system.network.connections UPDOWNCOUNTER
system.network.connections{network.transport, network.interface.name, network.connection.state}
```




|Attribute|Type|Description|
|-|-|-|
| network.transport | `tcp` \| `udp` \| `pipe` \| `unix` \| `quic` | [OSI transport layer](https://wikipedia.org/wiki/Transport_layer) or [inter-process communication method](https://wikipedia.org/wiki/Inter-process_communication).
  |
| network.interface.name | string | The network interface name.  |
| network.connection.state | `closed` \| `close_wait` \| `closing` \| `established` \| `fin_wait_1` \| `fin_wait_2` \| `last_ack` \| `listen` \| `syn_received` \| `syn_sent` \| `time_wait` | The state of network connection  |




### system.network.dropped

```py
# TYPE system.network.dropped COUNTER
system.network.dropped{network.io.direction, network.interface.name}
```

Count of packets that are dropped or discarded even though there was no error


|Attribute|Type|Description|
|-|-|-|
| network.io.direction | `transmit` \| `receive` | The network IO operation direction.  |
| network.interface.name | string | The network interface name.  |




### system.network.errors

```py
# TYPE system.network.errors COUNTER
system.network.errors{network.io.direction, network.interface.name}
```

Count of network errors detected


|Attribute|Type|Description|
|-|-|-|
| network.io.direction | `transmit` \| `receive` | The network IO operation direction.  |
| network.interface.name | string | The network interface name.  |




### system.network.io

```py
# TYPE system.network.io COUNTER
system.network.io{network.io.direction, network.interface.name}
```




|Attribute|Type|Description|
|-|-|-|
| network.io.direction | `transmit` \| `receive` | The network IO operation direction.  |
| network.interface.name | string | The network interface name.  |




### system.network.packets

```py
# TYPE system.network.packets COUNTER
system.network.packets{network.io.direction, system.device}
```




|Attribute|Type|Description|
|-|-|-|
| network.io.direction | `transmit` \| `receive` | The network IO operation direction.  |
| system.device | string | The device identifier  |




### system.paging.faults

```py
# TYPE system.paging.faults COUNTER
system.paging.faults{system.paging.type}
```




|Attribute|Type|Description|
|-|-|-|
| system.paging.type | `major` \| `minor` | The memory paging type  |




### system.paging.operations

```py
# TYPE system.paging.operations COUNTER
system.paging.operations{system.paging.type, system.paging.direction}
```




|Attribute|Type|Description|
|-|-|-|
| system.paging.type | `major` \| `minor` | The memory paging type  |
| system.paging.direction | `in` \| `out` | The paging access direction  |




### system.paging.usage

```py
# TYPE system.paging.usage UPDOWNCOUNTER
system.paging.usage{system.paging.state, system.device}
```

Unix swap or windows pagefile usage


|Attribute|Type|Description|
|-|-|-|
| system.paging.state | `used` \| `free` | The memory paging state  |
| system.device | string | Unique identifier for the device responsible for managing paging operations.  |




### system.paging.utilization

```py
# TYPE system.paging.utilization GAUGE
system.paging.utilization{system.paging.state, system.device}
```




|Attribute|Type|Description|
|-|-|-|
| system.paging.state | `used` \| `free` | The memory paging state  |
| system.device | string | Unique identifier for the device responsible for managing paging operations.  |




### system.process.count

```py
# TYPE system.process.count UPDOWNCOUNTER
system.process.count{system.process.status}
```

Total number of processes in each state


|Attribute|Type|Description|
|-|-|-|
| system.process.status | `running` \| `sleeping` \| `stopped` \| `defunct` | The process state, e.g., [Linux Process State Codes](https://man7.org/linux/man-pages/man1/ps.1.html#PROCESS_STATE_CODES)
  |




### system.process.created

```py
# TYPE system.process.created COUNTER
system.process.created{}
```

Total number of processes created over uptime of the host




### system.uptime

```py
# TYPE system.uptime GAUGE
system.uptime{}
```

The time the system has been running





## v8js

### v8js.gc.duration

```py
# TYPE v8js.gc.duration HISTOGRAM
v8js.gc.duration{v8js.gc.type}
```

Garbage collection duration.


|Attribute|Type|Description|
|-|-|-|
| v8js.gc.type | `major` \| `minor` \| `incremental` \| `weakcb` | The type of garbage collection.  |




### v8js.heap.space.available_size

```py
# TYPE v8js.heap.space.available_size UPDOWNCOUNTER
v8js.heap.space.available_size{v8js.heap.space.name}
```

Heap space available size.


|Attribute|Type|Description|
|-|-|-|
| v8js.heap.space.name | `new_space` \| `old_space` \| `code_space` \| `map_space` \| `large_object_space` | The name of the space type of heap memory.  |




### v8js.heap.space.physical_size

```py
# TYPE v8js.heap.space.physical_size UPDOWNCOUNTER
v8js.heap.space.physical_size{v8js.heap.space.name}
```

Committed size of a heap space.


|Attribute|Type|Description|
|-|-|-|
| v8js.heap.space.name | `new_space` \| `old_space` \| `code_space` \| `map_space` \| `large_object_space` | The name of the space type of heap memory.  |




### v8js.memory.heap.limit

```py
# TYPE v8js.memory.heap.limit UPDOWNCOUNTER
v8js.memory.heap.limit{v8js.heap.space.name}
```

Total heap memory size pre-allocated.


|Attribute|Type|Description|
|-|-|-|
| v8js.heap.space.name | `new_space` \| `old_space` \| `code_space` \| `map_space` \| `large_object_space` | The name of the space type of heap memory.  |




### v8js.memory.heap.used

```py
# TYPE v8js.memory.heap.used UPDOWNCOUNTER
v8js.memory.heap.used{v8js.heap.space.name}
```

Heap Memory size allocated.


|Attribute|Type|Description|
|-|-|-|
| v8js.heap.space.name | `new_space` \| `old_space` \| `code_space` \| `map_space` \| `large_object_space` | The name of the space type of heap memory.  |





## vcs

### vcs.change.count

```py
# TYPE vcs.change.count UPDOWNCOUNTER
vcs.change.count{vcs.repository.name, vcs.owner.name, vcs.change.state, vcs.repository.url.full, vcs.provider.name}
```

The number of changes (pull requests/merge requests/changelists) in a repository, categorized by their state (e.g. open or merged)


|Attribute|Type|Description|
|-|-|-|
| vcs.repository.name | string | The human readable name of the repository. It SHOULD NOT include any additional identifier like Group/SubGroup in GitLab or organization in GitHub.
  |
| vcs.owner.name | string | The group owner within the version control system.
  |
| vcs.change.state | `open` \| `wip` \| `closed` \| `merged` | The state of the change (pull request/merge request/changelist).
  |
| vcs.repository.url.full | string | The [canonical URL](https://support.google.com/webmasters/answer/10347851?hl=en#:~:text=A%20canonical%20URL%20is%20the,Google%20chooses%20one%20as%20canonical.) of the repository providing the complete HTTP(S) address in order to locate and identify the repository through a browser.
  |
| vcs.provider.name | `github` \| `gitlab` \| `gittea` \| `gitea` \| `bitbucket` | The name of the version control system provider.
  |




### vcs.change.duration

```py
# TYPE vcs.change.duration GAUGE
vcs.change.duration{vcs.repository.name, vcs.owner.name, vcs.change.state, vcs.repository.url.full, vcs.provider.name, vcs.ref.head.name}
```

The time duration a change (pull request/merge request/changelist) has been in a given state.


|Attribute|Type|Description|
|-|-|-|
| vcs.repository.name | string | The human readable name of the repository. It SHOULD NOT include any additional identifier like Group/SubGroup in GitLab or organization in GitHub.
  |
| vcs.owner.name | string | The group owner within the version control system.
  |
| vcs.change.state | `open` \| `wip` \| `closed` \| `merged` | The state of the change (pull request/merge request/changelist).
  |
| vcs.repository.url.full | string | The [canonical URL](https://support.google.com/webmasters/answer/10347851?hl=en#:~:text=A%20canonical%20URL%20is%20the,Google%20chooses%20one%20as%20canonical.) of the repository providing the complete HTTP(S) address in order to locate and identify the repository through a browser.
  |
| vcs.provider.name | `github` \| `gitlab` \| `gittea` \| `gitea` \| `bitbucket` | The name of the version control system provider.
  |
| vcs.ref.head.name | string | The name of the [reference](https://git-scm.com/docs/gitglossary#def_ref) such as **branch** or **tag** in the repository.
  |




### vcs.change.time_to_approval

```py
# TYPE vcs.change.time_to_approval GAUGE
vcs.change.time_to_approval{vcs.repository.name, vcs.ref.base.name, vcs.owner.name, vcs.repository.url.full, vcs.provider.name, vcs.ref.head.name, vcs.ref.head.revision, vcs.ref.base.revision}
```

The amount of time since its creation it took a change (pull request/merge request/changelist) to get the first approval.


|Attribute|Type|Description|
|-|-|-|
| vcs.repository.name | string | The human readable name of the repository. It SHOULD NOT include any additional identifier like Group/SubGroup in GitLab or organization in GitHub.
  |
| vcs.ref.base.name | string | The name of the [reference](https://git-scm.com/docs/gitglossary#def_ref) such as **branch** or **tag** in the repository.
  |
| vcs.owner.name | string | The group owner within the version control system.
  |
| vcs.repository.url.full | string | The [canonical URL](https://support.google.com/webmasters/answer/10347851?hl=en#:~:text=A%20canonical%20URL%20is%20the,Google%20chooses%20one%20as%20canonical.) of the repository providing the complete HTTP(S) address in order to locate and identify the repository through a browser.
  |
| vcs.provider.name | `github` \| `gitlab` \| `gittea` \| `gitea` \| `bitbucket` | The name of the version control system provider.
  |
| vcs.ref.head.name | string | The name of the [reference](https://git-scm.com/docs/gitglossary#def_ref) such as **branch** or **tag** in the repository.
  |
| vcs.ref.head.revision | string | The revision, literally [revised version](https://www.merriam-webster.com/dictionary/revision), The revision most often refers to a commit object in Git, or a revision number in SVN.
  |
| vcs.ref.base.revision | string | The revision, literally [revised version](https://www.merriam-webster.com/dictionary/revision), The revision most often refers to a commit object in Git, or a revision number in SVN.
  |




### vcs.change.time_to_merge

```py
# TYPE vcs.change.time_to_merge GAUGE
vcs.change.time_to_merge{vcs.repository.name, vcs.ref.base.name, vcs.owner.name, vcs.repository.url.full, vcs.provider.name, vcs.ref.head.name, vcs.ref.head.revision, vcs.ref.base.revision}
```

The amount of time since its creation it took a change (pull request/merge request/changelist) to get merged into the target(base) ref.


|Attribute|Type|Description|
|-|-|-|
| vcs.repository.name | string | The human readable name of the repository. It SHOULD NOT include any additional identifier like Group/SubGroup in GitLab or organization in GitHub.
  |
| vcs.ref.base.name | string | The name of the [reference](https://git-scm.com/docs/gitglossary#def_ref) such as **branch** or **tag** in the repository.
  |
| vcs.owner.name | string | The group owner within the version control system.
  |
| vcs.repository.url.full | string | The [canonical URL](https://support.google.com/webmasters/answer/10347851?hl=en#:~:text=A%20canonical%20URL%20is%20the,Google%20chooses%20one%20as%20canonical.) of the repository providing the complete HTTP(S) address in order to locate and identify the repository through a browser.
  |
| vcs.provider.name | `github` \| `gitlab` \| `gittea` \| `gitea` \| `bitbucket` | The name of the version control system provider.
  |
| vcs.ref.head.name | string | The name of the [reference](https://git-scm.com/docs/gitglossary#def_ref) such as **branch** or **tag** in the repository.
  |
| vcs.ref.head.revision | string | The revision, literally [revised version](https://www.merriam-webster.com/dictionary/revision), The revision most often refers to a commit object in Git, or a revision number in SVN.
  |
| vcs.ref.base.revision | string | The revision, literally [revised version](https://www.merriam-webster.com/dictionary/revision), The revision most often refers to a commit object in Git, or a revision number in SVN.
  |




### vcs.contributor.count

```py
# TYPE vcs.contributor.count GAUGE
vcs.contributor.count{vcs.repository.name, vcs.owner.name, vcs.repository.url.full, vcs.provider.name}
```

The number of unique contributors to a repository


|Attribute|Type|Description|
|-|-|-|
| vcs.repository.name | string | The human readable name of the repository. It SHOULD NOT include any additional identifier like Group/SubGroup in GitLab or organization in GitHub.
  |
| vcs.owner.name | string | The group owner within the version control system.
  |
| vcs.repository.url.full | string | The [canonical URL](https://support.google.com/webmasters/answer/10347851?hl=en#:~:text=A%20canonical%20URL%20is%20the,Google%20chooses%20one%20as%20canonical.) of the repository providing the complete HTTP(S) address in order to locate and identify the repository through a browser.
  |
| vcs.provider.name | `github` \| `gitlab` \| `gittea` \| `gitea` \| `bitbucket` | The name of the version control system provider.
  |




### vcs.ref.count

```py
# TYPE vcs.ref.count UPDOWNCOUNTER
vcs.ref.count{vcs.repository.name, vcs.owner.name, vcs.repository.url.full, vcs.provider.name, vcs.ref.type}
```

The number of refs of type branch or tag in a repository.


|Attribute|Type|Description|
|-|-|-|
| vcs.repository.name | string | The human readable name of the repository. It SHOULD NOT include any additional identifier like Group/SubGroup in GitLab or organization in GitHub.
  |
| vcs.owner.name | string | The group owner within the version control system.
  |
| vcs.repository.url.full | string | The [canonical URL](https://support.google.com/webmasters/answer/10347851?hl=en#:~:text=A%20canonical%20URL%20is%20the,Google%20chooses%20one%20as%20canonical.) of the repository providing the complete HTTP(S) address in order to locate and identify the repository through a browser.
  |
| vcs.provider.name | `github` \| `gitlab` \| `gittea` \| `gitea` \| `bitbucket` | The name of the version control system provider.
  |
| vcs.ref.type | `branch` \| `tag` | The type of the [reference](https://git-scm.com/docs/gitglossary#def_ref) in the repository.
  |




### vcs.ref.lines_delta

```py
# TYPE vcs.ref.lines_delta GAUGE
vcs.ref.lines_delta{vcs.repository.name, vcs.owner.name, vcs.repository.url.full, vcs.provider.name, vcs.ref.head.name, vcs.change.id, vcs.ref.head.type, vcs.ref.base.name, vcs.ref.base.type, vcs.line_change.type}
```

The number of lines added/removed in a ref (branch) relative to the ref from the `vcs.ref.base.name` attribute.


|Attribute|Type|Description|
|-|-|-|
| vcs.repository.name | string | The human readable name of the repository. It SHOULD NOT include any additional identifier like Group/SubGroup in GitLab or organization in GitHub.
  |
| vcs.owner.name | string | The group owner within the version control system.
  |
| vcs.repository.url.full | string | The [canonical URL](https://support.google.com/webmasters/answer/10347851?hl=en#:~:text=A%20canonical%20URL%20is%20the,Google%20chooses%20one%20as%20canonical.) of the repository providing the complete HTTP(S) address in order to locate and identify the repository through a browser.
  |
| vcs.provider.name | `github` \| `gitlab` \| `gittea` \| `gitea` \| `bitbucket` | The name of the version control system provider.
  |
| vcs.ref.head.name | string | The name of the [reference](https://git-scm.com/docs/gitglossary#def_ref) such as **branch** or **tag** in the repository.
  |
| vcs.change.id | string | The ID of the change (pull request/merge request/changelist) if applicable. This is usually a unique (within repository) identifier generated by the VCS system.
  |
| vcs.ref.head.type | `branch` \| `tag` | The type of the [reference](https://git-scm.com/docs/gitglossary#def_ref) in the repository.
  |
| vcs.ref.base.name | string | The name of the [reference](https://git-scm.com/docs/gitglossary#def_ref) such as **branch** or **tag** in the repository.
  |
| vcs.ref.base.type | `branch` \| `tag` | The type of the [reference](https://git-scm.com/docs/gitglossary#def_ref) in the repository.
  |
| vcs.line_change.type | `added` \| `removed` | The type of line change being measured on a branch or change.
  |




### vcs.ref.revisions_delta

```py
# TYPE vcs.ref.revisions_delta GAUGE
vcs.ref.revisions_delta{vcs.repository.name, vcs.owner.name, vcs.repository.url.full, vcs.provider.name, vcs.ref.head.name, vcs.change.id, vcs.ref.head.type, vcs.ref.base.name, vcs.ref.base.type, vcs.revision_delta.direction}
```

The number of revisions (commits) a ref (branch) is ahead/behind the branch from the `vcs.ref.base.name` attribute


|Attribute|Type|Description|
|-|-|-|
| vcs.repository.name | string | The human readable name of the repository. It SHOULD NOT include any additional identifier like Group/SubGroup in GitLab or organization in GitHub.
  |
| vcs.owner.name | string | The group owner within the version control system.
  |
| vcs.repository.url.full | string | The [canonical URL](https://support.google.com/webmasters/answer/10347851?hl=en#:~:text=A%20canonical%20URL%20is%20the,Google%20chooses%20one%20as%20canonical.) of the repository providing the complete HTTP(S) address in order to locate and identify the repository through a browser.
  |
| vcs.provider.name | `github` \| `gitlab` \| `gittea` \| `gitea` \| `bitbucket` | The name of the version control system provider.
  |
| vcs.ref.head.name | string | The name of the [reference](https://git-scm.com/docs/gitglossary#def_ref) such as **branch** or **tag** in the repository.
  |
| vcs.change.id | string | The ID of the change (pull request/merge request/changelist) if applicable. This is usually a unique (within repository) identifier generated by the VCS system.
  |
| vcs.ref.head.type | `branch` \| `tag` | The type of the [reference](https://git-scm.com/docs/gitglossary#def_ref) in the repository.
  |
| vcs.ref.base.name | string | The name of the [reference](https://git-scm.com/docs/gitglossary#def_ref) such as **branch** or **tag** in the repository.
  |
| vcs.ref.base.type | `branch` \| `tag` | The type of the [reference](https://git-scm.com/docs/gitglossary#def_ref) in the repository.
  |
| vcs.revision_delta.direction | `behind` \| `ahead` | The type of revision comparison.
  |




### vcs.ref.time

```py
# TYPE vcs.ref.time GAUGE
vcs.ref.time{vcs.repository.name, vcs.owner.name, vcs.repository.url.full, vcs.provider.name, vcs.ref.head.name, vcs.ref.head.type}
```

Time a ref (branch) created from the default branch (trunk) has existed. The `ref.type` attribute will always be `branch`


|Attribute|Type|Description|
|-|-|-|
| vcs.repository.name | string | The human readable name of the repository. It SHOULD NOT include any additional identifier like Group/SubGroup in GitLab or organization in GitHub.
  |
| vcs.owner.name | string | The group owner within the version control system.
  |
| vcs.repository.url.full | string | The [canonical URL](https://support.google.com/webmasters/answer/10347851?hl=en#:~:text=A%20canonical%20URL%20is%20the,Google%20chooses%20one%20as%20canonical.) of the repository providing the complete HTTP(S) address in order to locate and identify the repository through a browser.
  |
| vcs.provider.name | `github` \| `gitlab` \| `gittea` \| `gitea` \| `bitbucket` | The name of the version control system provider.
  |
| vcs.ref.head.name | string | The name of the [reference](https://git-scm.com/docs/gitglossary#def_ref) such as **branch** or **tag** in the repository.
  |
| vcs.ref.head.type | `branch` \| `tag` | The type of the [reference](https://git-scm.com/docs/gitglossary#def_ref) in the repository.
  |




### vcs.repository.count

```py
# TYPE vcs.repository.count UPDOWNCOUNTER
vcs.repository.count{vcs.owner.name, vcs.provider.name}
```

The number of repositories in an organization.


|Attribute|Type|Description|
|-|-|-|
| vcs.owner.name | string | The group owner within the version control system.
  |
| vcs.provider.name | `github` \| `gitlab` \| `gittea` \| `gitea` \| `bitbucket` | The name of the version control system provider.
  |




