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
	r := strings.NewReader(inputMsg)
	outBuf := new(bytes.Buffer)
	ui := NewUI(Reader(r), Writer(outBuf))

	inBuf := make([]byte, 1024)
	len, _ := ui.Reader().Read(inBuf)

	ui.Output(string(inBuf[:len]))
	result := outBuf.String()
	if result != inputMsg {
		t.Errorf("UI.Output = \"%s\", want \"%s\".", result, inputMsg)
	}
}

func TestUiOutputln(t *testing.T) {
	outBuf := new(bytes.Buffer)
	ui := NewUI(Writer(outBuf))

	ui.Outputln(inputMsg)
	result := outBuf.String()
	if result != inputMsg+"\n" {
		t.Errorf("UI.Output = \"%s\", want \"%s\".", result, inputMsg+"\n")
	}
}

func TestUiOutputBytes(t *testing.T) {
	outBuf := new(bytes.Buffer)
	ui := NewUI(Writer(outBuf))

	ui.OutputBytes([]byte(inputMsg))
	result := outBuf.String()
	if result != inputMsg {
		t.Errorf("UI.Output = \"%s\", want \"%s\".", result, inputMsg)
	}
}

func TestUiOutputErr(t *testing.T) {
	outBuf := new(bytes.Buffer)
	ui := NewUI(ErrorWriter(outBuf))

	ui.OutputErr(inputMsg)
	result := outBuf.String()
	if result != inputMsg {
		t.Errorf("UI.OutputErr = \"%s\", want \"%s\".", result, inputMsg)
	}
}

func TestUiOutputErrln(t *testing.T) {
	outBuf := new(bytes.Buffer)
	ui := NewUI(ErrorWriter(outBuf))

	ui.OutputErrln(inputMsg)
	result := outBuf.String()
	if result != inputMsg+"\n" {
		t.Errorf("UI.OutputErr = \"%s\", want \"%s\".", result, inputMsg+"\n")
	}
}
