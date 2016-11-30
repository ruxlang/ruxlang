package project

import (
	"path/filepath"
	"strings"
)

// PathName represents a package file path
type PathName string

// ProjectPath gets the root project path for this file path
func (p PathName) ProjectPath() string {
	return FindDir(string(p))
}

// ToPackageName generates the package name that this path would represent
func (p PathName) ToPackageName() string {
	path := filepath.Dir(string(p))
	projectPath := p.ProjectPath()
	if path == projectPath {
		path = filepath.Base(path)
	} else {
		path, _ = filepath.Rel(projectPath, path)
	}

	paths := make([]string, 0)
	for !isRoot(path) {
		var file string
		path, file = filepath.Split(path)
		path = filepath.Clean(path)
		paths = append([]string{file}, paths...)
	}

	return strings.Join(paths, ".")
}

func isRoot(path string) bool {
	return path == "" || path == "." || path == "/" || path == "\\" || strings.HasSuffix(path, ":/") || strings.HasSuffix(path, ":\\")
}
