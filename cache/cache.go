// Package cache : User cache file and directory
//
// These codes are licensed under CC0.
// http://creativecommons.org/publicdomain/zero/1.0/
package cache

import (
	"os"
	"path/filepath"
	"strings"
)

//Path returns path of user cache file
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

//Dir returns user cache directory
func Dir(appName string) string {
	if includeSlash(appName) {
		return ""
	}
	dir, err := os.UserCacheDir()
	if err != nil || len(dir) == 0 {
		dir = ""
	}
	return filepath.Join(dir, appName)
}

func includeSlash(path string) bool {
	if len(path) == 0 {
		return false
	}
	return strings.Contains(filepath.ToSlash(path), "/")
}
