## cbef (Couchbase Eventing Functions)
[![Release](https://img.shields.io/github/v/release/Trendyol/cbef?sort=semver)](https://github.com/Trendyol/cbef/releases)
[![GoDoc](https://img.shields.io/static/v1?label=godoc&message=reference&color=blue)](https://pkg.go.dev/github.com/Trendyol/cbef)
[![Test Coverage](https://codecov.io/gh/Trendyol/cbef/branch/main/graph/badge.svg)](https://codecov.io/gh/Trendyol/cbef)
[![Tag Push](https://github.com/Trendyol/cbef/actions/workflows/tag-push.yml/badge.svg)](https://github.com/Trendyol/cbef/actions/workflows/tag-push.yml)

`cbef` is a simple, gitops and testing capability provider for [Couchbase eventing functions](https://www.couchbase.com/products/eventing/).

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
