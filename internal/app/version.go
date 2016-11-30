package app

import (
	_ "embed"
	"fmt"
	"strings"
)

var (
	//go:embed version
	version string
)

func Version() string {
	return fmt.Sprintf("%s v%s", Name, strings.TrimSpace(version))
}
