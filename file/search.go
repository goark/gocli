package file

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
)

//SearchOption is setting for Search() function
type SearchOption struct {
	enableFile bool
	enableDir  bool
	toSlash    bool
	absPath    bool
}

//SrchOptFunc is self-referential function for functional options pattern
type SrchOptFunc func(*SearchOption)

//NewSearchOption returns SearchOption instance
func NewSearchOption(opts ...SrchOptFunc) *SearchOption {
	o := &SearchOption{enableFile: true, enableDir: true, toSlash: false, absPath: false}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

//WithEnableFile returns function for setting SearchOption
func WithEnableFile(b bool) SrchOptFunc {
	return func(o *SearchOption) {
		o.enableFile = b
	}
}

//WithEnableDir returns function for setting SearchOption
func WithEnableDir(b bool) SrchOptFunc {
	return func(o *SearchOption) {
		o.enableDir = b
	}
}

//WithToSlash returns function for setting SearchOption
func WithToSlash(b bool) SrchOptFunc {
	return func(o *SearchOption) {
		o.toSlash = b
	}
}

//WithAbsPath returns function for setting SearchOption
func WithAbsPath(b bool) SrchOptFunc {
	return func(o *SearchOption) {
		o.absPath = b
	}
}

//Search returns a list of search results.
func Search(path string, opt *SearchOption) []string {
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
				if abs, err := filepath.Abs(p); err == nil {
					if info, err := os.Stat(abs); err == nil {
						mode := info.Mode()
						if (dirFlag && (mode&os.ModeDir) != 0) || !dirFlag {
							paths = append(paths, normalizePath(abs, mode))
						}
					}
				}
			}
		}
		return paths
	}

	if abs, err := filepath.Abs(root); err == nil {
		if info, err := os.Stat(abs); err == nil {
			return []string{normalizePath(abs, info.Mode())}
		}
	}
	return []string{}
}

func getRoots(rootDir string) []string {
	if len(rootDir) == 0 {
		return []string{""}
	}
	if abs, err := filepath.Abs(rootDir); err == nil {
		rootDir = abs
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

func removeDuplicate(paths []string, opt *SearchOption) []string {
	if len(paths) == 0 {
		return paths
	}
	pathMap := make(map[string]struct{})
	cwd, _ := os.Getwd()
	for _, path := range paths {
		if !opt.absPath {
			if re, err := filepath.Rel(cwd, path); err == nil {
				if info, err := os.Stat(re); err == nil {
					path = normalizePath(re, info.Mode())
				}
			}
		}
		if strings.HasSuffix(path, string(filepath.Separator)) {
			if opt.enableDir {
				pathMap[path] = struct{}{}
			}
		} else if opt.enableFile {
			pathMap[path] = struct{}{}
		}
	}
	unqPaths := make([]string, 0, len(pathMap))
	for path := range pathMap {
		if opt.toSlash {
			path = filepath.ToSlash(path)
		}
		unqPaths = append(unqPaths, path)
	}
	sort.Strings(unqPaths)
	return unqPaths
}
