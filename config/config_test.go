package config

import (
	"os"
	"testing"
)

func TestConfigPathXDG(t *testing.T) {
	testCases := []struct {
		appName  string
		fileName string
		path     string
	}{
		{appName: "foo", fileName: "bar", path: "xdg/foo/bar"},
		{appName: "foo", fileName: "", path: ""},
		{appName: "", fileName: "bar", path: "xdg/bar"},
		{appName: "", fileName: "", path: ""},
		{appName: "../foo", fileName: "bar", path: ""},
		{appName: "foo", fileName: "baa/bar", path: ""},
	}
	os.Setenv("HOME", "home")
	os.Setenv("XDG_CONFIG_HOME", "xdg")
	for _, tc := range testCases {
		path := Path(tc.appName, tc.fileName)
		if path != tc.path {
			t.Errorf("GetConfigPath(\"%v\", \"%v\")  is \"%v\", watnt \"%v\".", tc.appName, tc.fileName, path, tc.path)
		}
	}
}

func TestConfigPathHome(t *testing.T) {
	testCases := []struct {
		appName  string
		fileName string
		path     string
	}{
		{appName: "foo", fileName: "bar", path: "home/foo/bar"},
		{appName: "foo", fileName: "", path: ""},
		{appName: "", fileName: "bar", path: "home/bar"},
		{appName: "", fileName: "", path: ""},
		{appName: "../foo", fileName: "bar", path: ""},
		{appName: "foo", fileName: "bar/bar", path: ""},
	}
	os.Setenv("HOME", "home")
	os.Unsetenv("XDG_CONFIG_HOME")
	for _, tc := range testCases {
		path := Path(tc.appName, tc.fileName)
		if path != tc.path {
			t.Errorf("GetConfigPath(\"%v\", \"%v\")  is \"%v\", watnt \"%v\".", tc.appName, tc.fileName, path, tc.path)
		}
	}
}

func TestConfigNoHome(t *testing.T) {
	testCases := []struct {
		appName  string
		fileName string
		path     string
	}{
		{appName: "foo", fileName: "bar", path: ""},
		{appName: "foo", fileName: "", path: ""},
		{appName: "", fileName: "bar", path: ""},
		{appName: "", fileName: "", path: ""},
		{appName: "../foo", fileName: "bar", path: ""},
		{appName: "foo", fileName: "bar/bar", path: ""},
	}
	os.Unsetenv("HOME")
	os.Unsetenv("XDG_CONFIG_HOME")
	for _, tc := range testCases {
		path := Path(tc.appName, tc.fileName)
		if path != tc.path {
			t.Errorf("GetConfigPath(\"%v\", \"%v\")  is \"%v\", watnt \"%v\".", tc.appName, tc.fileName, path, tc.path)
		}
	}
}
