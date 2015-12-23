// Package facade for CLI Application
//
// These codes are licensed under CC0.
// http://creativecommons.org/publicdomain/zero/1.0/
package facade

import (
	"io"

	"github.com/mitchellh/cli"
)

// Exit Code
const (
	ExitCodeOK    int = 0
	ExitCodeError int = iota
)

//Context inheritance cli.BasicUi
type Context struct {
	//Embedded BasicUi
	*cli.BasicUi
}

// NewContext returns a new Context instance
func NewContext(r io.Reader, w, e io.Writer) *Context {
	return &Context{BasicUi: &cli.BasicUi{Reader: r, Writer: w, ErrorWriter: e}}
}

// Facade is context of facade
type Facade struct {
	//UI defines user interface of the Cli
	Cxt *Context
	// commands is a mapping of subcommand names to a factory function
	commands map[string]cli.CommandFactory
}

// NewFacade returns a new Facade instance
func NewFacade(cxt *Context) *Facade {
	return &Facade{Cxt: cxt, commands: make(map[string]cli.CommandFactory)}
}

// AddCommand add command
func (f *Facade) AddCommand(name string, command cli.Command) {
	f.commands[name] = func() (cli.Command, error) {
		return command, nil
	}
}

// Run facade
func (f *Facade) Run(appName, version string, args []string) (int, error) {
	c := cli.NewCLI(appName, version)
	c.Args = args
	c.Commands = f.commands
	c.HelpWriter = f.Cxt.Writer
	return c.Run()
}
