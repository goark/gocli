# [gocli] - Minimal Packages for Command-Line Interface

[![check vulns](https://github.com/goark/gocli/workflows/vulns/badge.svg)](https://github.com/goark/gocli/actions)
[![lint status](https://github.com/goark/gocli/workflows/lint/badge.svg)](https://github.com/goark/gocli/actions)
[![GitHub license](https://img.shields.io/badge/license-CC0-blue.svg)](https://raw.githubusercontent.com/goark/gocli/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/goark/gocli.svg)](https://github.com/goark/gocli/releases/latest)

This package is required Go 1.16 or later.

**Migrated repository to [github.com/goark/gocli][gocli]**

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

Support `$XDG_CONFIG_HOME` environment value (XDG Base Directory)

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

[gocli]: https://github.com/goark/gocli "goark/gocli: Make Link with Markdown Format"
[dep]: https://github.com/golang/dep "golang/dep: Go dependency management tool"
[Context]: https://golang.org/pkg/context/ "context - The Go Programming Language"
