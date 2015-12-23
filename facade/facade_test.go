// Package facade for CLI Application
//
// These codes are licensed under CC0.
// http://creativecommons.org/publicdomain/zero/1.0/
package facade

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

// Context defines context of add command
type FacadeTest struct {
	//Embedded facade.Context
	*Context
}

// Command returns a new Context instance
func Command(cxt *Context) *FacadeTest {
	return &FacadeTest{Context: cxt}
}

// Synopsis of add-command
func (c FacadeTest) Synopsis() string {
	return "Add new task"
}

// Help of add-command
func (c FacadeTest) Help() string {
	helpText := `
Usage: sample add [options] TASK
`
	return fmt.Sprintln(strings.TrimSpace(helpText))
}

// Run add-command
func (c FacadeTest) Run(args []string) int {
	// Write your code here
	c.Output("normal end")
	return 0
}

func GetContext(inpmsg string) *Context {
	inBuf := strings.NewReader(inpmsg)
	outBuf := new(bytes.Buffer)
	errBuf := new(bytes.Buffer)
	return NewContext(inBuf, outBuf, errBuf)
}

func TestAddCommand(t *testing.T) {
	inpmsg := ""
	cxt := GetContext(inpmsg)
	fcd := NewFacade(cxt)
	fcd.AddCommand("sub", Command(cxt))
	if _, ok := fcd.commands["sub"]; !ok {
		t.Errorf("facade.AddCommand() = %v, want true.", ok)
	}
}

func TestRunNormal(t *testing.T) {
	inpmsg := ""
	cxt := GetContext(inpmsg)
	fcd := NewFacade(cxt)
	fcd.AddCommand("sub", Command(cxt))
	rtn, _ := fcd.Run("cmd", "0.0.0", []string{"sub", "arg"})
	if rtn != 0 {
		t.Errorf("facade.Run() = %v, want 0.", rtn)
	}
}

func TestRunError(t *testing.T) {
	inpmsg := ""
	cxt := GetContext(inpmsg)
	fcd := NewFacade(cxt)
	fcd.AddCommand("sub", Command(cxt))
	rtn, _ := fcd.Run("cmd", "0.0.0", []string{"err", "arg"})
	if rtn == 0 {
		t.Errorf("facade.Run() = %v, want not 0.", rtn)
	}
}
