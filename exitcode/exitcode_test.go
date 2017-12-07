package exitcode

import (
	"testing"
)

type ecodeTestCase struct { //Test case for ExitCode
	ec  ExitCode
	str string
}

func TestExitCode(t *testing.T) {
	testCases := []struct { //Test case for ExitCode
		ec  ExitCode
		str string
	}{
		{Normal, "normal end"},
		{Abnormal, "abnormal end"},
		{ExitCode(2), "unknown"},
	}

	for _, testCase := range testCases {
		if testCase.ec.String() != testCase.str {
			t.Errorf("ExitCode.String()  = %v, want %v.", testCase.ec.String(), testCase.str)
		}
	}
}
