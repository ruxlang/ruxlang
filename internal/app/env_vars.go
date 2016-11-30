package app

import (
	"os"
	"path"
)

const (
	DEFAULT_PATH = ".rux"
	PACKAGE_PATH = "RUX"
	TEMP_PATH    = "RUX_TMP"
)

func GetPackagePath() (string, error) {
	pkg := os.Getenv(PACKAGE_PATH)
	if pkg != "" {
		return pkg, nil
	}

	pkg, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return path.Join(pkg, DEFAULT_PATH), nil
}

func GetTempPath() string {
	tmp := os.Getenv(TEMP_PATH)
	if tmp != "" {
		return tmp
	}

	return os.TempDir()
}
