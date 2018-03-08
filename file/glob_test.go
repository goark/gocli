package file

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGlob(t *testing.T) {
	testCases := []struct {
		path   string
		result []string
	}{
		{path: "XXX/**/", result: []string{}},
		{path: "**/", result: []string{"./", "testdata/", "testdata/include/"}},
		{path: "*/", result: []string{"testdata/"}},
		{path: "**/*", result: []string{"glob.go", "glob_test.go", "testdata/", "testdata/include/", "testdata/include/source.h", "testdata/source.c"}},
		{path: "**/**/*", result: []string{"glob.go", "glob_test.go", "testdata/", "testdata/include/", "testdata/include/source.h", "testdata/source.c"}},
		{path: "./**/", result: []string{"./", "testdata/", "testdata/include/"}},
		{path: "./*/*", result: []string{"testdata/include/", "testdata/source.c"}},
		{path: "./**/*", result: []string{"glob.go", "glob_test.go", "testdata/", "testdata/include/", "testdata/include/source.h", "testdata/source.c"}},
		{path: "./**/**/*", result: []string{"glob.go", "glob_test.go", "testdata/", "testdata/include/", "testdata/include/source.h", "testdata/source.c"}},
		{path: "test*", result: []string{"testdata/"}},
		{path: "test*/", result: []string{"testdata/"}},
		{path: "test*/**/*.[ch]", result: []string{"testdata/include/source.h", "testdata/source.c"}},
		{path: "**/*.[ch]", result: []string{"testdata/include/source.h", "testdata/source.c"}},
		{path: "./**/*.[ch]", result: []string{"testdata/include/source.h", "testdata/source.c"}},
		{path: "**/../**/*.[ch]", result: []string{"testdata/include/source.h", "testdata/source.c"}},
	}
	for _, tc := range testCases {
		str := Glob(tc.path, NewGlobOption(
			WithFlags(GlobStdFlags|GlobSeparatorSlash),
		))
		if !reflect.DeepEqual(str, tc.result) {
			t.Errorf("Glob(\"%s\")  = %v, want %v.", tc.path, str, tc.result)
		}
	}
}

func TestGlobNil(t *testing.T) {
	testCases := []struct {
		path   string
		result []string
	}{
		{path: "", result: []string{}},
		{path: "*.go", result: []string{"glob.go", "glob_test.go"}},
	}
	for _, tc := range testCases {
		str := Glob(tc.path, nil)
		if !reflect.DeepEqual(str, tc.result) {
			t.Errorf("Glob(\"%s\")  = %v, want %v.", tc.path, str, tc.result)
		}
	}
}

func TestGlobFileOnly(t *testing.T) {
	testCases := []struct {
		path   string
		result []string
	}{
		{path: "XXX/**/", result: []string{}},
		{path: "**/", result: []string{}},
		{path: "./*/*", result: []string{"testdata/source.c"}},
		{path: "./**/*", result: []string{"glob.go", "glob_test.go", "testdata/include/source.h", "testdata/source.c"}},
		{path: "test*", result: []string{}},
		{path: "test*/", result: []string{}},
		{path: "test*/**/*.[ch]", result: []string{"testdata/include/source.h", "testdata/source.c"}},
		{path: "test*/**/XXX", result: []string{}},
		{path: "**/*.[ch]", result: []string{"testdata/include/source.h", "testdata/source.c"}},
		{path: "./**/*.[ch]", result: []string{"testdata/include/source.h", "testdata/source.c"}},
		{path: "**/../**/*.[ch]", result: []string{"testdata/include/source.h", "testdata/source.c"}},
	}
	for _, tc := range testCases {
		str := Glob(tc.path, NewGlobOption(
			WithFlags(GlobContainsFile|GlobSeparatorSlash),
		))
		if !reflect.DeepEqual(str, tc.result) {
			t.Errorf("Glob(\"%s\")  = %v, want %v.", tc.path, str, tc.result)
		}
	}
}

func TestGlobDirOnly(t *testing.T) {
	testCases := []struct {
		path   string
		result []string
	}{
		{path: "XXX/**/", result: []string{}},
		{path: "**/", result: []string{"./", "testdata/", "testdata/include/"}},
		{path: "*", result: []string{"testdata/"}},
		{path: "**/*", result: []string{"testdata/", "testdata/include/"}},
		{path: "./*/*", result: []string{"testdata/include/"}},
		{path: "./**/*", result: []string{"testdata/", "testdata/include/"}},
		{path: "*/", result: []string{"testdata/"}},
		{path: "testdata/**/", result: []string{"testdata/", "testdata/include/"}},
		{path: "test*", result: []string{"testdata/"}},
		{path: "test*/", result: []string{"testdata/"}},
		{path: "test*/**/", result: []string{"testdata/", "testdata/include/"}},
		{path: "test*/**/XXX", result: []string{}},
	}
	for _, tc := range testCases {
		str := Glob(tc.path, NewGlobOption(
			WithFlags(GlobContainsDir|GlobSeparatorSlash),
		))
		if !reflect.DeepEqual(str, tc.result) {
			t.Errorf("Glob(\"%s\")  = %v, want %v.", tc.path, str, tc.result)
		}
	}
}

func TestGlobAbs(t *testing.T) {
	testCases := []struct {
		path   string
		result []string
	}{
		{path: "**/", result: []string{"/path/to/file/", "/path/to/file/testdata/", "/path/to/file/testdata/include"}},
		{path: "test*/", result: []string{"/path/to/file/testdata/"}},
		{path: "**/*.[ch]", result: []string{"/path/to/file/testdata/include/source.h", "/path/to/file/testdata/source.c"}},
	}
	for _, tc := range testCases {
		str := Glob(tc.path, NewGlobOption(
			WithFlags(GlobStdFlags|GlobSeparatorSlash|GlobAbsolutePath),
		))
		if len(str) != len(tc.result) {
			t.Errorf("Glob(\"%s\")  = %v, want %v.", tc.path, str, tc.result)
		}
	}
}

func ExampleGlob() {
	result := Glob("**/*.[ch]", NewGlobOption(WithFlags(GlobStdFlags|GlobSeparatorSlash)))
	fmt.Println(result)
	// Output:
	// [testdata/include/source.h testdata/source.c]
}
