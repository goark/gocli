// Package config : Configuration file and directory
//
// These codes are licensed under CC0.
// http://creativecommons.org/publicdomain/zero/1.0/
package config

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func Path(appName, fileName string) string {
	if len(fileName) == 0 || includeSlash(fileName) {
		return ""
	}
	dir := Dir(appName)
	if len(dir) == 0 {
		return ""
	}
	return filepath.Join(dir, fileName)
}

//Dir returns Configuration directory
func Dir(appName string) string {
	if includeSlash(appName) {
		return ""
	}
	dir := XDGConfigHome()
	if len(dir) == 0 {
		dir = WindowsAppDataDir()
	}
	if len(dir) == 0 {
		var err error
		dir, err = os.UserHomeDir()
		if err != nil {
			dir = ""
		}
	}
	if len(dir) == 0 {
		return ""
	}
	return filepath.Join(dir, appName)
}

//XDGConfigHome returns $XDG_CONFIG_HOME directory
func XDGConfigHome() string {
	return os.Getenv("XDG_CONFIG_HOME")
}

//WindowsAppDataDir returns %APPDATA% directory if Windows
func WindowsAppDataDir() string {
	if runtime.GOOS == "windows" {
		appDir := os.Getenv("HOME")
		if len(appDir) == 0 {
			return os.Getenv("APPDATA")
		}
		return appDir
	}
	return ""
}

func includeSlash(path string) bool {
	if len(path) == 0 {
		return false
	}
	return strings.Contains(filepath.ToSlash(path), "/")
}
