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
	reset1 := setTestEnv("HOME", "home")
	defer reset1()
	reset2 := setTestEnv("XDG_CONFIG_HOME", "xdg")
	defer reset2()
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
	reset1 := setTestEnv("HOME", "home")
	defer reset1()
	reset2 := unsetTestEnv("XDG_CONFIG_HOME")
	defer reset2()
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
	reset1 := unsetTestEnv("HOME")
	defer reset1()
	reset2 := unsetTestEnv("XDG_CONFIG_HOME")
	defer reset2()
	for _, tc := range testCases {
		path := Path(tc.appName, tc.fileName)
		if path != tc.path {
			t.Errorf("GetConfigPath(\"%v\", \"%v\")  is \"%v\", watnt \"%v\".", tc.appName, tc.fileName, path, tc.path)
		}
	}
}

func setTestEnv(key, val string) func() {
	preVal := os.Getenv(key)
	os.Setenv(key, val)
	return func() {
		os.Setenv(key, preVal)
	}
}

func unsetTestEnv(key string) func() {
	preVal := os.Getenv(key)
	os.Unsetenv(key)
	return func() {
		os.Setenv(key, preVal)
	}
}
