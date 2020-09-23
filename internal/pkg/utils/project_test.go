package utils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestBasePath(t *testing.T) {
	bp := BasePath()
	gopath := os.Getenv("GOPATH")
	pp := filepath.Join(gopath, "src", "github.com", "MadJlzz", "gopypi")
	if bp != pp {
		t.Errorf("base path is not set to the root of the project.\ngot: [%s], want: [%s]", pp, bp)
	}
}
