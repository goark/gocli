package file

import (
	"reflect"
	"testing"
)

func TestUiSearch(t *testing.T) {
	testCases := []struct {
		path   string
		result []string
	}{
		{path: "XXX/**/", result: []string{}},
		{path: "**/", result: []string{"./", "testdata/"}},
		{path: "**/*", result: []string{"search.go", "search_test.go", "testdata/", "testdata/source.c", "testdata/source.h"}},
		{path: "./**/*", result: []string{"search.go", "search_test.go", "testdata/", "testdata/source.c", "testdata/source.h"}},
		{path: "testdata/**/*.[ch]", result: []string{"testdata/source.c", "testdata/source.h"}},
		{path: "**/*.[ch]", result: []string{"testdata/source.c", "testdata/source.h"}},
		{path: "./**/*.[ch]", result: []string{"testdata/source.c", "testdata/source.h"}},
		{path: "**/../**/*.[ch]", result: []string{"../file/testdata/source.c", "../file/testdata/source.h", "testdata/source.c", "testdata/source.h"}},
	}
	for _, tc := range testCases {
		str := Search(tc.path, NewSearchOption(
			WithToSlash(true),
		))
		if !reflect.DeepEqual(str, tc.result) {
			t.Errorf("Search(\"%s\")  = %v, want %v.", tc.path, str, tc.result)
		}
	}
}

func TestUiSearchFileOnly(t *testing.T) {
	testCases := []struct {
		path   string
		result []string
	}{
		{path: "XXX/**/", result: []string{}},
		{path: "**/", result: []string{}},
		{path: "./**/*", result: []string{"search.go", "search_test.go", "testdata/source.c", "testdata/source.h"}},
		{path: "testdata/**/*.[ch]", result: []string{"testdata/source.c", "testdata/source.h"}},
		{path: "**/*.[ch]", result: []string{"testdata/source.c", "testdata/source.h"}},
		{path: "./**/*.[ch]", result: []string{"testdata/source.c", "testdata/source.h"}},
		{path: "**/../**/*.[ch]", result: []string{"../file/testdata/source.c", "../file/testdata/source.h", "testdata/source.c", "testdata/source.h"}},
	}
	for _, tc := range testCases {
		str := Search(tc.path, NewSearchOption(
			WithToSlash(true),
			WithEnableFile(true),
			WithEnableDir(false),
		))
		if !reflect.DeepEqual(str, tc.result) {
			t.Errorf("Search(\"%s\")  = %v, want %v.", tc.path, str, tc.result)
		}
	}
}

func TestUiSearchDirOnly(t *testing.T) {
	testCases := []struct {
		path   string
		result []string
	}{
		{path: "XXX/**/", result: []string{}},
		{path: "**/", result: []string{"./", "testdata/"}},
		{path: "**/*", result: []string{"testdata/"}},
		{path: "./**/*", result: []string{"testdata/"}},
		{path: "testdata", result: []string{"testdata/"}},
		{path: "testdata/**/", result: []string{"testdata/"}},
		{path: "test*", result: []string{"testdata/"}},
		{path: "test*/**/", result: []string{"testdata/"}},
		{path: "test*/**/XXX", result: []string{}},
	}
	for _, tc := range testCases {
		str := Search(tc.path, NewSearchOption(
			WithToSlash(true),
			WithEnableFile(false),
			WithEnableDir(true),
		))
		if !reflect.DeepEqual(str, tc.result) {
			t.Errorf("Search(\"%s\")  = %v, want %v.", tc.path, str, tc.result)
		}
	}
}

func TestUiSearchAbs(t *testing.T) {
	testCases := []struct {
		path   string
		result []string
	}{
		{path: "XXX/**/", result: []string{}},
		{path: "**/", result: []string{"/path/to/file/", "/path/to/file/testdata/"}},
		{path: "**/*", result: []string{"/path/to/file/search.go", "/path/to/file/search_test.go", "/path/to/file/testdata/", "/path/to/file/testdata/source.c", "testdata/source.h"}},
		{path: "./**/*", result: []string{"/path/to/file/search.go", "/path/to/file/search_test.go", "/path/to/file/testdata/", "/path/to/file/testdata/source.c", "/path/to/file/testdata/source.h"}},
		{path: "testdata/**/*.[ch]", result: []string{"/path/to/file/testdata/source.c", "/path/to/file/testdata/source.h"}},
		{path: "**/*.[ch]", result: []string{"/path/to/file/testdata/source.c", "/path/to/file/testdata/source.h"}},
		{path: "./**/*.[ch]", result: []string{"/path/to/file/testdata/source.c", "/path/to/file/testdata/source.h"}},
		{path: "**/../**/*.[ch]", result: []string{"/path/to/file/testdata/source.c", "/path/to/file/testdata/source.h"}},
	}
	for _, tc := range testCases {
		str := Search(tc.path, NewSearchOption(
			WithToSlash(true),
			WithAbsPath(true),
		))
		if len(str) != len(tc.result) {
			t.Errorf("Search(\"%s\")  = %v, want %v.", tc.path, str, tc.result)
		}
	}
}
