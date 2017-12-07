# [gocli] - Minimal Packages for Command-Line Interface

[![Build Status](https://travis-ci.org/spiegel-im-spiegel/gocli.svg?branch=master)](https://travis-ci.org/spiegel-im-spiegel/gocli)
[![GitHub license](https://img.shields.io/badge/license-CC0-blue.svg)](https://raw.githubusercontent.com/spiegel-im-spiegel/gocli/master/LICENSE)

## Install

```
$ go get -u github.com/spiegel-im-spiegel/gocli
```

Installing by [dep].

```
$ dep ensure -add github.com/spiegel-im-spiegel/gocli
```

## Example

```go
package main

import (
    "fmt"

    "github.com/spiegel-im-spiegel/gocli/exitcode"
    "github.com/spiegel-im-spiegel/gocli/rwi"
)

func run(ui *rwi.RWI) exitcode.ExitCode {
    ui.Outputln("Hello world")
    return exitcode.Normal
}

func main() {
    run(rwi.New(
        rwi.Reader(os.Stdin),
        rwi.Writer(os.Stdout),
        rwi.ErrorWriter(os.Stderr),
    )).Exit()
}
```

[gocli]: https://github.com/spiegel-im-spiegel/gocli "spiegel-im-spiegel/gocli: Make Link with Markdown Format"
[dep]: https://github.com/golang/dep "golang/dep: Go dependency management tool"
