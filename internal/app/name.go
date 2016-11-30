package app

import (
	"os"
	"path"
	"strings"
)

var Name = strings.Title(path.Base(os.Args[0]))
