// Package rwi : Reader/Writer Interface for command-line
//
// These codes are licensed under CC0.
// http://creativecommons.org/publicdomain/zero/1.0/
package prompt

import (
	"fmt"
	"testing"
)

func TestErrno(t *testing.T) {
	testCases := []struct {
		err error
		str string
	}{
		{err: Errno(0), str: "unknown error (0)"},
		{err: ErrTerminate, str: "terminate prompt"},
		{err: ErrNotTerminal, str: "not terminal (or pipe?)"},
		{err: Errno(3), str: "unknown error (3)"},
	}

	for _, tc := range testCases {
		errStr := tc.err.Error()
		if errStr != tc.str {
			t.Errorf("\"%v\" != \"%v\"", errStr, tc.str)
		}
		fmt.Printf("Info(TestErrno): %+v\n", tc.err)
	}
}
