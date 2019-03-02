// Package rwi : Reader/Writer Interface for command-line
//
// These codes are licensed under CC0.
// http://creativecommons.org/publicdomain/zero/1.0/
package prompt

import "fmt"

//Errno is error number for CVSS
type Errno int

const (
	ErrTerminate Errno = iota + 1
	ErrNotTerminal
)

var errMessage = map[Errno]string{
	ErrTerminate:   "terminate prompt",
	ErrNotTerminal: "not terminal (or pipe?)",
}

func (n Errno) Error() string {
	if s, ok := errMessage[n]; ok {
		return s
	}
	return fmt.Sprintf("unknown error (%d)", int(n))
}
