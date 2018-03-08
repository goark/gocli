// Package file : Operating files and directories
//
// These codes are licensed under CC0.
// http://creativecommons.org/publicdomain/zero/1.0/
package file

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
)

//GlobOption is setting for Glob() function
type GlobOption struct {
	enableFile bool
	enableDir  bool
	toSlash    bool
	absPath    bool
}

//GlogOptFunc is self-referential function for functional options pattern
type GlogOptFunc func(*GlobOption)

//NewGlobOption returns GlobOption instance
func NewGlobOption(opts ...GlogOptFunc) *GlobOption {
	o := &GlobOption{enableFile: true, enableDir: true, toSlash: false, absPath: false}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

//WithEnableFile returns function for setting GlobOption
func WithEnableFile(b bool) GlogOptFunc {
	return func(o *GlobOption) {
		o.enableFile = b
	}
}

//WithEnableDir returns function for setting GlobOption
func WithEnableDir(b bool) GlogOptFunc {
	return func(o *GlobOption) {
		o.enableDir = b
	}
}

//WithToSlash returns function for setting GlobOption
func WithToSlash(b bool) GlogOptFunc {
	return func(o *GlobOption) {
		o.toSlash = b
	}
}

//WithAbsPath returns function for setting GlobOption
func WithAbsPath(b bool) GlogOptFunc {
	return func(o *GlobOption) {
		o.absPath = b
	}
}

//Glob returns an array containing the matching directory/file names..
func Glob(path string, opt *GlobOption) []string {
	if path == "" {
		return []string{}
	}
	if opt == nil {
		opt = NewGlobOption()
	}
	return removeDuplicate(getPaths("", path), opt)
}

func getPaths(rootDir, path string) []string {
	roots := getRoots(rootDir)
	if len(rootDir) > 0 && len(roots) == 0 {
		return []string{}
	}
	paths := []string{}
	for _, root := range roots {
		paths = append(paths, getPathsNext(root, path)...)
	}
	return paths
}

func getPathsNext(root, path string) []string {
	path = filepath.ToSlash(path)
	if strings.HasPrefix(path, "**/") {
		var roots []string
		if len(root) == 0 {
			roots = walkDir("./")
		} else {
			roots = walkDir(root)
		}
		subPath := path[3:]
		if len(roots) > 0 {
			paths := []string{}
			for _, rt := range roots {
				ps := getPaths(rt, subPath)
				if len(ps) > 0 {
					paths = append(paths, ps...)
				}
			}
			return paths
		}
		path = subPath
	} else if i := strings.Index(path, "/**/"); i >= 0 {
		return getPaths(root+path[:i+1], path[i+1:])
	}

	if len(path) > 0 {
		dirFlag := false
		if strings.HasSuffix(path, "/") {
			path = path[:len(path)-1]
			dirFlag = true
		}
		paths := []string{}
		if ps, err := filepath.Glob(root + path); err == nil {
			for _, p := range ps {
				if info, err := os.Stat(p); err == nil {
					mode := info.Mode()
					if (dirFlag && (mode&os.ModeDir) != 0) || !dirFlag {
						paths = append(paths, normalizePath(p, mode))
					}
				}
			}
		}
		return paths
	}

	if info, err := os.Stat(root); err == nil {
		return []string{normalizePath(root, info.Mode())}
	}
	return []string{}
}

func getRoots(rootDir string) []string {
	if len(rootDir) == 0 {
		return []string{""}
	}
	if strings.HasSuffix(rootDir, "/") {
		rootDir = rootDir[:len(rootDir)-1]
	}
	roots := []string{}
	if paths, err := filepath.Glob(rootDir); err == nil {
		for _, path := range paths {
			if info, err := os.Stat(path); err == nil && info.IsDir() {
				roots = append(roots, normalizePath(path, info.Mode()))
			}
		}
	}
	return roots
}

func walkDir(root string) []string {
	paths := []string{}
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err == nil && info.IsDir() {
			normalizePath(path, info.Mode())
			paths = append(paths, filepath.ToSlash(normalizePath(path, info.Mode())))
		}
		return nil
	})
	return paths
}

func normalizePath(path string, mode os.FileMode) string {
	tail := ""
	if (mode & os.ModeDir) != 0 { //directory
		tail = string(filepath.Separator)
	}
	return filepath.Clean(path) + tail
}

func removeDuplicate(paths []string, opt *GlobOption) []string {
	if len(paths) == 0 {
		return paths
	}
	pathMap := make(map[string]string)
	for _, path := range paths {
		absPath, err := filepath.Abs(path)
		if err != nil {
			absPath = ""
		} else if info, err := os.Stat(absPath); err == nil {
			absPath = normalizePath(absPath, info.Mode())
			if opt.absPath {
				path = absPath
			}
		}
		if strings.HasSuffix(absPath, string(filepath.Separator)) {
			if opt.enableDir {
				pathMap[absPath] = path
			}
		} else if opt.enableFile {
			pathMap[absPath] = path
		}
	}
	unqPaths := make([]string, 0, len(pathMap))
	for _, path := range pathMap {
		if opt.toSlash {
			path = filepath.ToSlash(path)
		}
		unqPaths = append(unqPaths, path)
	}
	sort.Strings(unqPaths)
	return unqPaths
}
