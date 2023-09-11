# component-base

[![test](https://github.com/shipengqi/component-base/actions/workflows/test.yaml/badge.svg)](https://github.com/shipengqi/component-base/actions/workflows/test.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/shipengqi/component-base)](https://goreportcard.com/report/github.com/shipengqi/component-base)
[![release](https://img.shields.io/github/release/shipengqi/component-base.svg)](https://github.com/shipengqi/component-base/releases)
[![license](https://img.shields.io/github/license/shipengqi/component-base)](https://github.com/shipengqi/component-base/blob/main/LICENSE)

## Getting Started

### json

Use `json-iterator` instead of `encoding/json` by `-tags=jsoniter`
Use `go-json` instead of `encoding/json` by `-tags=gojson`
Use `sonic` instead of `encoding/json` by `-tags=sonic`, requirements: avx, linux/windows/darwin, amd64

Example:

```
$ go build -tag `-tags=jsoniter`
```

### version

Note: `VERSION_PKG` must be `github.com/shipengqi/component-base/version`.

Example:
```makefile
# The project's root import path
PKG := github.com/example/repo
# set version package
VERSION_PKG=github.com/shipengqi/component-base/version

ifeq ($(origin VERSION), undefined)
VERSION := $(shell git describe --tags --always --match='v*')
endif

# set git commit and tree state
GIT_COMMIT = $(shell git rev-parse HEAD)
ifneq ($(shell git status --porcelain 2> /dev/null),)
	GIT_TREE_STATE ?= dirty
else
	GIT_TREE_STATE ?= clean
endif

# set ldflags
GO_LDFLAGS += -X $(VERSION_PKG).Version=$(VERSION) \
	-X $(VERSION_PKG).GitCommit=$(GIT_COMMIT) \
	-X $(VERSION_PKG).GitTreeState=$(GIT_TREE_STATE) \
	-X $(VERSION_PKG).BuildDate=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
	
.PHONY: go.build
go.build:
	@echo "===========> Building: $(OUTPUT_DIR)/$(BIN)"
	@CGO_ENABLED=0 go build -ldflags "$(GO_LDFLAGS)" -o $(OUTPUT_DIR)/$(BIN) ${PKG}
```

Print the version message:

```go
package main

import (
	"fmt"
	
	"github.com/shipengqi/component-base/version"
)

func main() {
	fmt.Println(version.Get().String())
}
```

The output resembling as follows:

```
Version:      5fa0fea
Commit:       5fa0fea6e39fabdad0eda2dcfd55f70da2fa89ea
GitTreeState: dirty
BuildTime:    1970-01-01T00:00:00Z
GoVersion:    go1.21.0
Compiler:     gc
Platform:     linux/amd64
```

### cli

```go
package main

import (
	"github.com/spf13/cobra"
	cliflag "github.com/shipengqi/component-base/cli/flag"
	"github.com/shipengqi/component-base/cli/globalflag"
	"github.com/shipengqi/component-base/term"
)

func main() {
	cmd := &cobra.Command{
		Use:   "demo",
		Short: "demo description",
		RunE:  func (c *cobra.Command, args []string) error {
			return nil
		},
	}
	cliflag.InitFlags(cmd.Flags())

	var fss cliflag.NamedFlagSets
	// add one or more FlagSet
	fakes := fss.FlagSet("fake")
	fakes.StringVar(&o.Username, "username", o.Username, "fake username.")
	fakes.StringVar(&o.Password, "password", o.Password, "fake password.")

	// applies the FlagSets to this command 
	fs := cmd.Flags()
	for _, set := range fss.FlagSets {
		fs.AddFlagSet(set)
	}

	// applies global help flag to this command 
	globalflag.AddGlobalFlags(fss.FlagSet("global"), cmd.Name())
	
	// set both usage and help function.
	width, _, _ := term.TerminalSize(cmd.OutOrStdout())
	cliflag.SetUsageAndHelpFunc(cmd, fss, width)
}
```

### term

```go
package main

import (
	"log"
	"os"
	
	"github.com/shipengqi/component-base/term"
)

func main() {
	// get the current width and height of the user's terminal.
	// If it isn't a terminal, nil is returned. 
	width, height, err := term.TerminalSize(os.Stdout)
	if err != nil {
		log.Fatalln(err)
	}
}
```

## Documentation

You can find the docs at [go docs](https://pkg.go.dev/github.com/shipengqi/component-base).