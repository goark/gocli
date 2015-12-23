// Package gocli : Command line interface
//
// These codes are licensed under CC0.
// http://creativecommons.org/publicdomain/zero/1.0/
package gocli

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// UI is Command line user interface
type UI struct {
	Reader      io.Reader
	Writer      io.Writer
	ErrorWriter io.Writer
}

// NewUI returns a new Context instance
func NewUI() *UI {
	return &UI{Reader: os.Stdin, Writer: os.Stdout, ErrorWriter: os.Stderr}
}

//Output to Writer stream.
func (c *UI) Output(val ...interface{}) error {
	return doOutput(c.Writer, val)
}

//Outputln to Writer stream (add line-ending).
func (c *UI) Outputln(val ...interface{}) error {
	return doOutputln(c.Writer, val)
}

//OutputBytes to Writer stream ([]byte data).
func (c *UI) OutputBytes(data []byte) error {
	writer := bufio.NewWriter(c.Writer)
	if _, err := writer.Write(data); err != nil {
		return err
	}
	return writer.Flush()
}

//OutputErr to ErrorWriter stream.
func (c *UI) OutputErr(val ...interface{}) error {
	return doOutput(c.ErrorWriter, val)
}

//OutputErrln to ErrorWriter stream (add line-ending).
func (c *UI) OutputErrln(val ...interface{}) error {
	return doOutputln(c.ErrorWriter, val)
}

//Output to ErrorWriter stream.
func doOutput(writer io.Writer, val []interface{}) error {
	_, err := fmt.Fprint(writer, val...)
	return err
}

//Output to ErrorWriter stream (add line-ending).
func doOutputln(writer io.Writer, val []interface{}) error {
	_, err := fmt.Fprintln(writer, val...)
	return err
}
