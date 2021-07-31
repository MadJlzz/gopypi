package utils

import (
	"os"
	"testing"
)

func TestBasePath(t *testing.T) {
	bp := BasePath()
	_ = os.Chdir("../..")
	pp, _ := os.Getwd()
	if bp != pp {
		t.Errorf("base path is not set to the root of the project.\ngot: [%s], want: [%s]", pp, bp)
	}
}