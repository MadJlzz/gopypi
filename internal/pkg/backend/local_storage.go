package backend

import (
	"github.com/MadJlzz/gopypi/internal/pkg/model"
	"github.com/sirupsen/logrus"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

const fileProtocol = "file:///"

// We can have a local storage with specify
// The location of the files to be stored
// Should we need something else ? Dunno
type LocalStorage struct {
	Location string
}

func NewLocalStorage(location string) *LocalStorage {
	return &LocalStorage{
		Location: location,
	}
}

func (ls *LocalStorage) Load() (map[string]*model.Package, error) {
	var pkgs = make(map[string]*model.Package)
	err := filepath.Walk(ls.Location, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		pkgName := filepath.Base(filepath.Dir(path))
		pkgURI := craftURI(path)
		pkg, found := pkgs[pkgName]
		if found {
			pkg.URI = append(pkgs[pkgName].URI, pkgURI)
		} else {
			pkgs[pkgName] = model.New(pkgName, pkgURI)
		}
		return nil
	})
	return pkgs, err
}

func craftURI(path string) *url.URL {
	cleanPath := fileProtocol + strings.ReplaceAll(path, "\\", "/")
	uri, err := url.ParseRequestURI(cleanPath)
	if err != nil {
		logrus.Error(err)
	}
	return uri
}
