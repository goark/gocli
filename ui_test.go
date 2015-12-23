package gocli

import (
	"bytes"
	"os"
	"runtime"
	"strings"
	"testing"
)

var inputMsgs []string
var inputMsg string
var lineEnding string

func TestMain(m *testing.M) {
	//initialization
	inputMsgs = []string{
		"Take the Go-lang!",
		"Go言語で行こう！",
	}
	inputMsg = strings.Join(inputMsgs, "\n")
	if runtime.GOOS == "windows" {
		lineEnding = "\r\n"
	} else {
		lineEnding = "\n"
	}

	//start test
	code := m.Run()

	//termination
	os.Exit(code)
}

func TestUiOutput(t *testing.T) {
	outBuf := new(bytes.Buffer)
	ui := &UI{Writer: outBuf}

	ui.Output(inputMsg)
	result := outBuf.String()
	if result != inputMsg {
		t.Errorf("CliUi.Output = \"%s\", want \"%s\".", result, inputMsg)
	}
}

func TestUiOutputln(t *testing.T) {
	outBuf := new(bytes.Buffer)
	ui := &UI{Writer: outBuf}

	ui.Outputln(inputMsg)
	result := outBuf.String()
	if result != inputMsg+"\n" {
		t.Errorf("CliUi.Output = \"%s\", want \"%s\".", result, inputMsg+"\n")
	}
}

func TestUiOutputBytes(t *testing.T) {
	outBuf := new(bytes.Buffer)
	ui := &UI{Writer: outBuf}

	ui.OutputBytes([]byte(inputMsg))
	result := outBuf.String()
	if result != inputMsg {
		t.Errorf("CliUi.Output = \"%s\", want \"%s\".", result, inputMsg)
	}
}

func TestUiOutputErr(t *testing.T) {
	outBuf := new(bytes.Buffer)
	ui := &UI{ErrorWriter: outBuf}

	ui.OutputErr(inputMsg)
	result := outBuf.String()
	if result != inputMsg {
		t.Errorf("CliUi.OutputErr = \"%s\", want \"%s\".", result, inputMsg)
	}
}

func TestUiOutputErrln(t *testing.T) {
	outBuf := new(bytes.Buffer)
	ui := &UI{ErrorWriter: outBuf}

	ui.OutputErrln(inputMsg)
	result := outBuf.String()
	if result != inputMsg+"\n" {
		t.Errorf("CliUi.OutputErr = \"%s\", want \"%s\".", result, inputMsg+"\n")
	}
}
