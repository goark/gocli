// Package gocli : Command line interface
//
// These codes are licensed under CC0.
// http://creativecommons.org/publicdomain/zero/1.0/
package gocli

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
)

// UI is Command line user interface
type UI struct {
	reader      io.Reader
	writer      io.Writer
	errorWriter io.Writer
}

// NewUI returns a new UI instance
func NewUI(r io.Reader, w, e io.Writer) *UI {
	if r == nil {
		r = ioutil.NopCloser(bytes.NewReader(nil))
	}
	if w == nil {
		w = ioutil.Discard
	}
	if e == nil {
		e = ioutil.Discard
	}
	return &UI{reader: r, writer: w, errorWriter: e}
}

//Reader returns io.Reader stream
func (c *UI) Reader() io.Reader {
	return c.reader
}

//Output to writer stream.
func (c *UI) Output(val ...interface{}) error {
	return doOutput(c.writer, val)
}

//Outputln to writer stream (add line-ending).
func (c *UI) Outputln(val ...interface{}) error {
	return doOutputln(c.writer, val)
}

//OutputBytes to writer stream ([]byte data).
func (c *UI) OutputBytes(data []byte) error {
	writer := bufio.NewWriter(c.writer)
	if _, err := writer.Write(data); err != nil {
		return err
	}
	return writer.Flush()
}

//OutputErr to errorWriter stream.
func (c *UI) OutputErr(val ...interface{}) error {
	return doOutput(c.errorWriter, val)
}

//OutputErrln to errorWriter stream (add line-ending).
func (c *UI) OutputErrln(val ...interface{}) error {
	return doOutputln(c.errorWriter, val)
}

//Output to io.Writer stream.
func doOutput(writer io.Writer, val []interface{}) error {
	_, err := fmt.Fprint(writer, val...)
	return err
}

//Output to io.Writer stream (add line-ending).
func doOutputln(writer io.Writer, val []interface{}) error {
	_, err := fmt.Fprintln(writer, val...)
	return err
}
