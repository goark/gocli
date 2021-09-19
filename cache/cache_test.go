package cache_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/spiegel-im-spiegel/gocli/cache"
)

func TestCachePath(t *testing.T) {
	dir, err := os.UserCacheDir()
	if err != nil {
		t.Errorf("os.UserConfigDir() error is \"%v\", want nil error.", err)
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
		path := cache.Path(tc.appName, tc.fileName)
		if path != tc.path {
			t.Errorf("cache.Path(\"%v\", \"%v\")  is \"%v\", want \"%v\".", tc.appName, tc.fileName, path, tc.path)
		}
	}
}
