## cbef (Couchbase Eventing Functions)
[![Release](https://img.shields.io/github/v/release/Trendyol/cbef?sort=semver)](https://github.com/Trendyol/cbef/releases)
[![GoDoc](https://img.shields.io/static/v1?label=godoc&message=reference&color=blue)](https://pkg.go.dev/github.com/Trendyol/cbef)
[![Test Coverage](https://codecov.io/gh/Trendyol/cbef/branch/main/graph/badge.svg)](https://codecov.io/gh/Trendyol/cbef)
[![Tag Push](https://github.com/Trendyol/cbef/actions/workflows/tag-push.yml/badge.svg)](https://github.com/Trendyol/cbef/actions/workflows/tag-push.yml)

`cbef` is a simple, gitops and testing capability provider for [Couchbase eventing functions](https://www.couchbase.com/products/eventing/).

### Why?
In production projects where code reliability is a key consideration, it is essential to control and review the implemented processes using Git.

`cbef` facilitates GitOps support for the operations performed within eventing functions.

Moreover, to ensure the seamless functionality of eventing functions before deploying to the live environment, testing support is provided. This testing ensures that the functions operate as expected, offering confidence in their reliability.

In contrast to operations carried out through the UI, transitioning to a new function version is designed to be zero-downtime, allowing for rapid rollback if needed. The inclusion of revision history support further enhances the ability to track and manage changes effectively.

### Quickstart

#### Local
> Precompiled cbef for all arch types is available in the [releases](https://github.com/trendyol/cbef/releases) page.

Clone the `cbef` repository from the given GitHub URL
```bash
git clone https://github.com/trendyol/cbef
```

Change the current working directory to the `cbef` directory
```bash
cd cbef
```

Update function settings
```bash
nano ./examples/functions/basic.json
```

Update function file
```bash
nano ./examples/functions/basic.js
```

Download dependencies
```bash
go mod tidy
```

Set the environment variable `CONFIG_FILE` to the path of the `basic.json` configuration file
```bash
export CONFIG_FILE=./examples/functions/basic.json
```

Set the environment variable `FUNCTION_FILE` to the path of the `basic.js` JavaScript file
```bash
export FUNCTION_FILE=./examples/functions/basic.js
```

Set the environment variable `CI_COMMIT_SHORT_SHA` to the value `foo`, representing a short SHA identifier
```bash
export CI_COMMIT_SHORT_SHA=foo
```

Set the environment variable `CI_COMMIT_AUTHOR` to the value `foo`, representing the author of the commit
```bash
export CI_COMMIT_AUTHOR=foo
```

Set the environment variable `EXECUTION_TIMEOUT` to `3m`, representing a timeout of 3 minutes
```bash
export EXECUTION_TIMEOUT=3m
```

Run the Go program located in the `cmd` directory
```bash
go run ./cmd
```

### Function Configurations
> Configurations that can be used in files with `.json` extension located in the functions folder

| Field                                    | Description                                                  |  Value Type  | Required | Options                            |
|------------------------------------------|--------------------------------------------------------------|:------------:|:--------:|------------------------------------|
| cluster                                  | Couchbase connection configurations                          |    object    |    ✅     |                                    |
| cluster.connection_string                | Couchbase cluster connection string                          |    string    |    ✅     |                                    |
| cluster.user                             | Username for cluster authentication                          |    string    |    ✅     |                                    |
| cluster.pass                             | Password for cluster authentication                          |    string    |    ✅     |                                    |
| name                                     | Function name                                                |    string    |    ✅     |                                    |
| metadata_keyspace                        | Metadata keyspace configurations                             |    object    |    ✅     |                                    |
| metadata_keyspace.bucket                 | Metadata bucket name                                         |    string    |    ✅     |                                    |
| metadata_keyspace.scope                  | Metadata scope name                                          |    string    |    ✅     |                                    |
| metadata_keyspace.collection             | Metadata collection name                                     |    string    |    ✅     |                                    |
| source_keyspace                          | Source keyspace configurations                               |    object    |    ✅     |                                    |
| source_keyspace.bucket                   | Source bucket name                                           |    string    |    ✅     |                                    |
| source_keyspace.scope                    | Source scope name                                            |    string    |    ✅     |                                    |
| source_keyspace.collection               | Source collection name                                       |    string    |    ✅     |                                    |
| bucket_bindings                          | Bucket bindings configurations                               | object array |    ❌     |                                    |
| bucket_bindings.alias                    | Alias for bucket binding                                     |    string    |    ❌     |                                    |
| bucket_bindings.bucket                   | Target bucket name for binding                               |    string    |    ❌     |                                    |
| bucket_bindings.scope                    | Target scope name for binding                                |    string    |    ❌     |                                    |
| bucket_bindings.collection               | Target collection name for binding                           |    string    |    ❌     |                                    |
| bucket_bindings.access                   | Access level for binding                                     |    string    |    ❌     | r (read-only), rw (read-write)     |
| url_bindings                             | URL bindings configurations                                  | object array |    ❌     |                                    |
| bucket_bindings.hostname                 | Hostname for URL binding                                     |    string    |    ❌     |                                    |
| bucket_bindings.alias                    | Alias for URL binding                                        |    string    |    ❌     |                                    |
| bucket_bindings.allow_cookies            | Allow cookies for URL binding                                |     bool     |    ❌     |                                    |
| bucket_bindings.validate_ssl_certificate | Validate SSL certificate for URL binding                     |     bool     |    ❌     |                                    |
| bucket_bindings.auth                     | Authentication configurations for URL binding                |    object    |    ❌     |                                    |
| bucket_bindings.auth.type                | Authentication type for URL binding                          |    string    |    ❌     | basic, digest, bearer              |
| bucket_bindings.auth.user                | Username for URL binding authentication                      |    string    |    ❌     |                                    |
| bucket_bindings.auth.pass                | Password for URL binding authentication                      |    string    |    ❌     |                                    |
| bucket_bindings.auth.token               | Token for URL binding authentication                         |    string    |    ❌     |                                    |
| constant_bindings                        | Constant bindings configurations                             | object array |    ❌     |                                    |
| constant_bindings.alias                  | Alias for constant binding                                   |    string    |    ❌     |                                    |
| constant_bindings.literal                | Literal value for constant binding                           |    string    |    ❌     |                                    |
| settings                                 | Function configurations                                      |    object    |    ❌     |                                    |
| settings.dcp_stream_boundary             | The preferred deployment time feed boundary for the function |    string    |    ❌     | everything, from_now               |
| settings.description                     | Description for function                                     |    string    |    ❌     |                                    |
| settings.log_level                       | Granularity of system events being captured in the log       |    string    |    ❌     | INFO, ERROR, WARNING, DEBUG, TRACE |
| settings.query_consistency               | Consistency level of N1QL statements in the function         |     uint     |    ❌     | 1 (NotBounded), 2 (RequestPlus)    |
| settings.worker_count                    | Number of workers per node to process the events             |     uint     |    ❌     |                                    |
| settings.language_compatibility          | Language compatibility of the function                       |    string    |    ❌     | 6.0.0, 6.5.0, 6.6.2                |
| settings.execution_timeout               | Time after which the function's execution will be timed out  |     uint     |    ❌     |                                    |
| settings.timer_context_size              | Maximum allowed value of the timer context size in bytes     |     uint     |    ❌     |                                    |
