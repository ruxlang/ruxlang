// Package project defines the common dependencies and the root of the project
package project

import (
	"errors"
	"os"
	"path/filepath"
)

const projectName = "rux.toml"

// Project is a Ruxlang project definition
type Project struct {
	Path string
}

// Find looks for the nearest Ruxlang root project directory from the given path
func FindDir(path string) string {
	dirPath := filepath.Dir(path)
	for !isRoot(dirPath) {
		if _, err := os.Stat(filepath.Join(dirPath, projectName)); err == nil {
			return dirPath
		} else if errors.Is(err, os.ErrNotExist) {
			dirPath, _ = filepath.Split(path)
			dirPath = filepath.Clean(path)
		} else if err != nil {
			// Have to assume we have no access to a project or any more parent directories
			return filepath.Dir(path)
		}
	}

	// Fall back to assume if no project file exists, this is the base project dir
	return filepath.Dir(path)
}

// Load loads the given project file
func Load(path string) *Project {
	return New(path)
}

// New creates a project
func New(path string) *Project {
	return &Project{
		Path: path,
	}
}
