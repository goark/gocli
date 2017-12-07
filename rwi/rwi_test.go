package rwi

import (
	"bytes"
	"io"
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
	ui := New(Writer(outBuf))

	ui.Output(inputMsg)
	result := outBuf.String()
	if result != inputMsg {
		t.Errorf("UI.Output = \"%s\", want \"%s\".", result, inputMsg)
	}
}

func TestUiOutputln(t *testing.T) {
	outBuf := new(bytes.Buffer)
	ui := New(Writer(outBuf))

	ui.Outputln(inputMsg)
	result := outBuf.String()
	if result != inputMsg+"\n" {
		t.Errorf("UI.Outputln = \"%s\", want \"%s\".", result, inputMsg+"\n")
	}
}

func TestUiOutputBytes(t *testing.T) {
	outBuf := new(bytes.Buffer)
	ui := New(Writer(outBuf))

	ui.OutputBytes([]byte(inputMsg))
	result := outBuf.String()
	if result != inputMsg {
		t.Errorf("UI.OutputBytes = \"%s\", want \"%s\".", result, inputMsg)
	}
}

func TestUiOutputErr(t *testing.T) {
	outBuf := new(bytes.Buffer)
	ui := New(ErrorWriter(outBuf))

	ui.OutputErr(inputMsg)
	result := outBuf.String()
	if result != inputMsg {
		t.Errorf("UI.OutputErr = \"%s\", want \"%s\".", result, inputMsg)
	}
}

func TestUiOutputErrln(t *testing.T) {
	outBuf := new(bytes.Buffer)
	ui := New(ErrorWriter(outBuf))

	ui.OutputErrln(inputMsg)
	result := outBuf.String()
	if result != inputMsg+"\n" {
		t.Errorf("UI.OutputErrln = \"%s\", want \"%s\".", result, inputMsg+"\n")
	}
}

func TestUiOutputErrBytes(t *testing.T) {
	outBuf := new(bytes.Buffer)
	ui := New(ErrorWriter(outBuf))

	ui.OutputErrBytes([]byte(inputMsg))
	result := outBuf.String()
	if result != inputMsg {
		t.Errorf("UI.OutputErrBytes = \"%s\", want \"%s\".", result, inputMsg)
	}
}

func TestUiReader(t *testing.T) {
	r := strings.NewReader(inputMsg)
	ui := New(Reader(r))

	inBuf := new(bytes.Buffer)
	io.Copy(inBuf, ui.Reader())
	result := inBuf.String()
	if result != inputMsg {
		t.Errorf("UI.Reader = \"%s\", want \"%s\".", result, inputMsg)
	}
}

func TestUiWriter(t *testing.T) {
	r := strings.NewReader(inputMsg)
	outBuf := new(bytes.Buffer)
	ui := New(Reader(r), Writer(outBuf))

	io.Copy(ui.Writer(), ui.Reader())
	result := outBuf.String()
	if result != inputMsg {
		t.Errorf("UI.Writer = \"%s\", want \"%s\".", result, inputMsg)
	}
}

func TestUiErrorWriter(t *testing.T) {
	r := strings.NewReader(inputMsg)
	outBuf := new(bytes.Buffer)
	ui := New(Reader(r), ErrorWriter(outBuf))

	io.Copy(ui.ErrorWriter(), ui.Reader())
	result := outBuf.String()
	if result != inputMsg {
		t.Errorf("UI.ErrorWriter = \"%s\", want \"%s\".", result, inputMsg)
	}
}
