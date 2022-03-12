package config_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/goark/gocli/config"
)

func TestConfigPath(t *testing.T) {
	dir, err := os.UserConfigDir()
	if err != nil {
		t.Errorf("os.UserConfigDir() is \"%v\", want nil error.", err)
	}
	testCases := []struct {
		appName  string
		fileName string
		path     string
	}{
		{appName: "foo", fileName: "bar", path: filepath.Join(dir, "foo", "bar")},
		{appName: "foo", fileName: "", path: ""},
		{appName: "", fileName: "bar", path: filepath.Join(dir, "bar")},
		{appName: "", fileName: "", path: ""},
		{appName: "../foo", fileName: "bar", path: ""},
		{appName: "foo", fileName: "bar/bar", path: ""},
	}
	for _, tc := range testCases {
		path := config.Path(tc.appName, tc.fileName)
		if path != tc.path {
			t.Errorf("config.Path(\"%v\", \"%v\")  is \"%v\", want \"%v\".", tc.appName, tc.fileName, path, tc.path)
		}
	}
}
