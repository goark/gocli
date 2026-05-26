# [gocli] - Minimal Packages for Command-Line Interface

[![CI Status](https://github.com/goark/gocli/workflows/ci/badge.svg)](https://github.com/goark/gocli/actions/workflows/ci.yml)
[![CodeQL](https://github.com/goark/gocli/workflows/CodeQL/badge.svg)](https://github.com/goark/gocli/actions/workflows/codeql.yml)
[![GitHub license](https://img.shields.io/badge/license-CC0-blue.svg)](https://raw.githubusercontent.com/goark/gocli/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/goark/gocli.svg)](https://github.com/goark/gocli/releases/latest)
[![Go Reference](https://pkg.go.dev/badge/github.com/goark/gocli.svg)](https://pkg.go.dev/github.com/goark/gocli)

This package requires Go 1.25 or later.

## Design goals

`gocli` provides a small collection of focused sub-packages for building command-line tools in Go.
Each package addresses a single concern:

- `exitcode` — OS exit code constants and helpers
- `rwi` — Reader/Writer interface wrapping stdin/stdout/stderr
- `signal` — Signal handling via `context.Context` (deprecated: use [`os/signal.NotifyContext`](https://pkg.go.dev/os/signal#NotifyContext) instead)
- `file` — File/directory globbing with wildcard support
- `config` — XDG-aware configuration file path helpers
- `cache` — XDG-aware cache file path helpers

## Development

Requires [Task] for local validation.

```text
task test
```

This runs `go mod verify`, `go test -shuffle on ./...`, and golangci-lint.

## Declare [gocli] module

See [go.mod](https://github.com/goark/gocli/blob/master/go.mod) file.

## Usage of [gocli] package

```go
package main

import (
    "os"

    "github.com/goark/gocli/exitcode"
    "github.com/goark/gocli/rwi"
)

func run(ui *rwi.RWI) exitcode.ExitCode {
    ui.Outputln("Hello world")
    return exitcode.Normal
}

func main() {
    run(rwi.New(
        rwi.WithReader(os.Stdin),
        rwi.WithWriter(os.Stdout),
        rwi.WithErrorWriter(os.Stderr),
    )).Exit()
}
```

### Handling SIGNAL with [Context] Package

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"

    "github.com/goark/gocli/signal"
)

func ticker(ctx context.Context) error {
    t := time.NewTicker(1 * time.Second) // 1 second cycle
    defer t.Stop()

    for {
        select {
        case now := <-t.C: // ticker event
            fmt.Println(now.Format(time.RFC3339))
        case <-ctx.Done(): // cancel event from context
            fmt.Println("Stop ticker")
            return ctx.Err()
        }
    }
}

func Run() error {
    errCh := make(chan error, 1)
    defer close(errCh)

    go func() {
        child, cancelChild := context.WithTimeout(
            signal.Context(context.Background(), os.Interrupt), // cancel event by SIGNAL
            10*time.Second, // timeout after 10 seconds
        )
        defer cancelChild()
        errCh <- ticker(child)
    }()

    err := <-errCh
    fmt.Println("Done")
    return err
}

func main() {
    if err := Run(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
}
```

### Search Files and Directories (w/ Wildcard)

```go
import (
    "fmt"

    "github.com/goark/gocli/file"
)

result := file.Glob("**/*.[ch]", file.NewGlobOption())
fmt.Println(result)
// Output:
// [testdata/include/source.h testdata/source.c]
```

### Configuration file and directory

Support `$XDG_CONFIG_HOME` environment value (XDG Base Directory)

```go
import (
    "fmt"

    "github.com/goark/gocli/config"
)

path := config.Path("app", "config.json")
fmt.Println(path)
// Output:
// /home/username/.config/app/config.json
```

### User cache file and directory

Support `$XDG_CACHE_HOME` environment value (XDG Base Directory)

```go
import (
    "fmt"

    "github.com/goark/gocli/cache"
)

path := cache.Path("app", "access.log")
fmt.Println(path)
// Output:
// /home/username/.cache/app/access.log
```

[gocli]: https://github.com/goark/gocli "goark/gocli: Minimal Packages for Command-Line Interface"
[Context]: https://pkg.go.dev/context "context - Go Packages"
[Task]: https://taskfile.dev "Task - A task runner / simpler Make alternative"
