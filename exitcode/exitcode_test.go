package exitcode

import (
	"os"
	"os/exec"
	"strconv"
	"testing"
)

func TestExitCode(t *testing.T) {
	testCases := []struct { // Test case for ExitCode
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

func TestExit(t *testing.T) {
	testCases := []struct {
		ec   ExitCode
		want int
	}{
		{Normal, 0},
		{Abnormal, 1},
		{ExitCode(2), 2},
	}

	for _, tc := range testCases {
		if os.Getenv("TEST_EXIT_CODE") == strconv.Itoa(int(tc.ec)) {
			tc.ec.Exit()
		}
	}

	for _, tc := range testCases {
		cmd := exec.Command(os.Args[0], "-test.run=TestExit") //nolint:gosec
		cmd.Env = append(os.Environ(), "TEST_EXIT_CODE="+strconv.Itoa(int(tc.ec)))
		err := cmd.Run()
		if tc.ec == Normal {
			if err != nil {
				t.Errorf("Exit(%d): expected success, got %v", tc.ec, err)
			}
		} else {
			exitErr, ok := err.(*exec.ExitError)
			if !ok {
				t.Errorf("Exit(%d): expected ExitError, got %v", tc.ec, err)
				continue
			}
			if exitErr.ExitCode() != tc.want {
				t.Errorf("Exit(%d): exit code = %d, want %d", tc.ec, exitErr.ExitCode(), tc.want)
			}
		}
	}
}

func TestExitIfNotNormal(t *testing.T) {
	testCases := []struct {
		ec      ExitCode
		wantExt bool
		want    int
	}{
		{Normal, false, 0},
		{Abnormal, true, 1},
		{ExitCode(2), true, 2},
	}

	for _, tc := range testCases {
		if os.Getenv("TEST_EXIT_IF_CODE") == strconv.Itoa(int(tc.ec)) {
			tc.ec.ExitIfNotNormal()
			return // Normal case: should reach here
		}
	}

	for _, tc := range testCases {
		cmd := exec.Command(os.Args[0], "-test.run=TestExitIfNotNormal") //nolint:gosec
		cmd.Env = append(os.Environ(), "TEST_EXIT_IF_CODE="+strconv.Itoa(int(tc.ec)))
		err := cmd.Run()
		if !tc.wantExt {
			if err != nil {
				t.Errorf("ExitIfNotNormal(%d): expected no exit, got %v", tc.ec, err)
			}
		} else {
			exitErr, ok := err.(*exec.ExitError)
			if !ok {
				t.Errorf("ExitIfNotNormal(%d): expected ExitError, got %v", tc.ec, err)
				continue
			}
			if exitErr.ExitCode() != tc.want {
				t.Errorf("ExitIfNotNormal(%d): exit code = %d, want %d", tc.ec, exitErr.ExitCode(), tc.want)
			}
		}
	}
}
