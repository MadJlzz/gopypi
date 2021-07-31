package utils

import (
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Join(filepath.Dir(b), "../..")
)

// BasePath returns the root directory of the project.
func BasePath() string {
	return basepath
}
