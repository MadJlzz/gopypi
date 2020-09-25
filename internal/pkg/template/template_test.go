package template

import (
	"github.com/MadJlzz/gopypi/internal/pkg/utils"
	"github.com/MadJlzz/gopypi/internal/pkg/web"
	"os"
	"path/filepath"
	"testing"
)

func TestGenerate(t *testing.T) {
	path := filepath.Join(utils.BasePath(), "web", "index.gohtml")
	err := Generate(os.Stdout, path, []*web.Package{web.New("Sauron", "https://sauron.lotr.dev/_--=--../")})
	if err != nil {
		t.Errorf("error occurred when generating template\ngot: [%v]", err)
	}
}
