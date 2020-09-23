package template

import (
	"github.com/MadJlzz/gopypi/internal/pkg/utils"
	"os"
	"path/filepath"
	"testing"
)

func TestGenerate(t *testing.T) {
	path := filepath.Join(utils.BasePath(), "web", "index.gohtml")
	err := Generate(os.Stdout, path)
	if err != nil {
		t.Errorf("error occurred when generating template\ngot: [%v]", err)
	}
}
