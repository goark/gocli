package prompt

import (
	"bytes"
	"strings"
	"testing"

	"github.com/spiegel-im-spiegel/gocli/rwi"
)

var (
	testLogic = func(s string) (string, error) {
		if s == "q" || s == "quit" {
			return "quit prompt", ErrTerminate
		}
		runes := []rune(s)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		return string(runes), nil
	}
	inputMsg = "abcdef\nq\n"
)

func TestIsNotTeminal(t *testing.T) {
	ui := rwi.New(
		rwi.WithReader(strings.NewReader(inputMsg)),
		rwi.WithWriter(new(bytes.Buffer)),
	)
	p := New(ui, testLogic)
	if p.IsTerminal() {
		t.Errorf("Prompt.IsTerminal = %v, want false.", p.IsTerminal())
	}
}

func TestRun(t *testing.T) {
	outBuf := new(bytes.Buffer)
	outputMsg := "fedcba\nquit prompt\n"
	ui := rwi.New(
		rwi.WithReader(strings.NewReader(inputMsg)),
		rwi.WithWriter(outBuf),
	)
	p := New(ui, testLogic)
	if err := p.Run(); err != nil {
		t.Errorf("Prompt.Run = %v, want nil.", err)
	} else {
		result := outBuf.String()
		if result != outputMsg {
			t.Errorf("output of Prompt.Run = \"%s\", want \"%s\".", result, outputMsg)
		}
	}
}

func TestRunCustom(t *testing.T) {
	outBuf := new(bytes.Buffer)
	outputMsg := "Input 'q' or 'quit' to stop\nTest> fedcba\nTest> quit prompt\n"
	ui := rwi.New(
		rwi.WithReader(strings.NewReader(inputMsg)),
		rwi.WithWriter(outBuf),
	)
	p := New(ui,
		testLogic,
		WithPromptString("Test> "),
		WithHeaderMessage("Input 'q' or 'quit' to stop"),
	)
	if err := p.Run(); err != nil {
		t.Errorf("Prompt.Run = %v, want nil.", err)
	} else {
		result := outBuf.String()
		if result != outputMsg {
			t.Errorf("output of Prompt.Run = \"%s\", want \"%s\".", result, outputMsg)
		}
	}
}

func TestOnce(t *testing.T) {
	outBuf := new(bytes.Buffer)
	outputMsg := "Input string\nTest> fedcba\n"
	ui := rwi.New(
		rwi.WithReader(strings.NewReader(inputMsg)),
		rwi.WithWriter(outBuf),
	)
	p := New(ui,
		testLogic,
		WithPromptString("Test> "),
		WithHeaderMessage("Input string"),
	)
	if err := p.Once(); err != nil {
		t.Errorf("Prompt.Once = %v, want nil.", err)
	} else {
		result := outBuf.String()
		if result != outputMsg {
			t.Errorf("output of Prompt.Once = \"%s\", want \"%s\".", result, outputMsg)
		}
	}
}
