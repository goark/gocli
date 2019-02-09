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
	ui := New(WithWriter(outBuf))

	if err := ui.Output(inputMsg); err != nil {
		t.Errorf("RWI.Output = \"%v\", want nil.", err)
	} else {
		result := outBuf.String()
		if result != inputMsg {
			t.Errorf("RWI.Output = \"%s\", want \"%s\".", result, inputMsg)
		}
	}
}

func TestUiOutputln(t *testing.T) {
	outBuf := new(bytes.Buffer)
	ui := New(WithWriter(outBuf))

	if err := ui.Outputln(inputMsg); err != nil {
		t.Errorf("RWI.Outputln = \"%v\", want nil.", err)
	} else {
		result := outBuf.String()
		if result != inputMsg+"\n" {
			t.Errorf("RWI.Outputln = \"%s\", want \"%s\".", result, inputMsg+"\n")
		}
	}
}

func TestUiOutputBytes(t *testing.T) {
	outBuf := new(bytes.Buffer)
	ui := New(WithWriter(outBuf))

	if err := ui.OutputBytes([]byte(inputMsg)); err != nil {
		t.Errorf("RWI.OutputBytes = \"%v\", want nil.", err)

	} else {
		result := outBuf.String()
		if result != inputMsg {
			t.Errorf("RWI.OutputBytes = \"%s\", want \"%s\".", result, inputMsg)
		}
	}
}

func TestUiOutputErr(t *testing.T) {
	outBuf := new(bytes.Buffer)
	ui := New(WithErrorWriter(outBuf))

	if err := ui.OutputErr(inputMsg); err != nil {
		t.Errorf("RWI.OutputErr = \"%v\", want nil.", err)
	} else {
		result := outBuf.String()
		if result != inputMsg {
			t.Errorf("RWI.OutputErr = \"%s\", want \"%s\".", result, inputMsg)
		}
	}
}

func TestUiOutputErrln(t *testing.T) {
	outBuf := new(bytes.Buffer)
	ui := New(WithErrorWriter(outBuf))

	if err := ui.OutputErrln(inputMsg); err != nil {
		t.Errorf("RWI.OutputErrln = \"%v\", want nil.", err)
	} else {
		result := outBuf.String()
		if result != inputMsg+"\n" {
			t.Errorf("RWI.OutputErrln = \"%s\", want \"%s\".", result, inputMsg+"\n")
		}
	}
}

func TestUiOutputErrBytes(t *testing.T) {
	outBuf := new(bytes.Buffer)
	ui := New(WithErrorWriter(outBuf))

	if err := ui.OutputErrBytes([]byte(inputMsg)); err != nil {
		t.Errorf("RWI.OutputErrBytes = \"%v\", want nil.", err)
	} else {
		result := outBuf.String()
		if result != inputMsg {
			t.Errorf("RWI.OutputErrBytes = \"%s\", want \"%s\".", result, inputMsg)
		}
	}
}

func TestUiReader(t *testing.T) {
	r := strings.NewReader(inputMsg)
	ui := New(WithReader(r))

	inBuf := new(bytes.Buffer)
	if _, err := io.Copy(inBuf, ui.Reader()); err != nil {
		t.Errorf("RWI.Copy = \"%v\", want nil.", err)
	} else {
		result := inBuf.String()
		if result != inputMsg {
			t.Errorf("RWI.Reader = \"%s\", want \"%s\".", result, inputMsg)
		}
	}
}

func TestUiWriter(t *testing.T) {
	r := strings.NewReader(inputMsg)
	outBuf := new(bytes.Buffer)
	ui := New(WithReader(r), WithWriter(outBuf))

	if _, err := io.Copy(ui.Writer(), ui.Reader()); err != nil {
		t.Errorf("RWI.Copy = \"%v\", want nil.", err)
	} else {
		result := outBuf.String()
		if result != inputMsg {
			t.Errorf("RWI.Writer = \"%s\", want \"%s\".", result, inputMsg)
		}
	}
}

func TestUiErrorWriter(t *testing.T) {
	r := strings.NewReader(inputMsg)
	outBuf := new(bytes.Buffer)
	ui := New(WithReader(r), WithErrorWriter(outBuf))

	if _, err := io.Copy(ui.ErrorWriter(), ui.Reader()); err != nil {
		t.Errorf("RWI.Copy = \"%v\", want nil.", err)
	} else {
		result := outBuf.String()
		if result != inputMsg {
			t.Errorf("RWI.ErrorWriter = \"%s\", want \"%s\".", result, inputMsg)
		}
	}
}
