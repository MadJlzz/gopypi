package disk

import (
	"crypto/sha256"
	"fmt"
	"github.com/MadJlzz/gopypi/pkg/listing"
	"go.uber.org/zap"
	"io/fs"
	"os"
	"path/filepath"
)

type Storage struct {
	logger *zap.SugaredLogger
	path   string
}

type StorageOption func(s *Storage)

func WithPath(path string) StorageOption {
	return func(s *Storage) {
		s.path = path
	}
}

func NewStorage(logger *zap.SugaredLogger, opts ...StorageOption) *Storage {
	const defaultPath = "/tmp"
	s := &Storage{
		logger: logger,
		path:   defaultPath,
	}
	// Loop through each option.
	for _, opt := range opts {
		opt(s)
	}
	if _, err := os.Stat(s.path); os.IsNotExist(err) {
		s.logger.Fatalf("path [%q] does not exist", s.path)
	}

	return s
}

func (s Storage) GetAllPackages() []listing.Package {
	var res []listing.Package
	var m = make(map[string]listing.Package)

	err := filepath.Walk(s.path, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			s.logger.Warnf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() {
			return nil
		}
		version := filepath.Dir(path)
		lib := filepath.Dir(version)
		hash := fmt.Sprintf("%x", sha256.Sum256([]byte(lib+version)))

		if validExt(info.Name()) {
			pkg, ok := m[hash]
			fmt.Println(ok)
			pkg.Distributions = append(pkg.Distributions, )
			m[hash].Distributions = append(m[hash].Distributions, listing.Distribution(path))
		} else {
			m[hash] = listing.Package{
				Name:          filepath.Base(lib),
				Version:       filepath.Base(version),
				Distributions: make([]listing.Distribution, 0),
			}

		}

		fmt.Printf("%x\n", hash)

		res = append(res, listing.Package{
			Name:     filepath.Base(lib),
			Version:  filepath.Base(version),
			Location: path,
		})

		fmt.Printf("visited file or dir: %q\n", lib)
		fmt.Printf("visited file or dir: %q\n", path)
		return nil
	})
	if err != nil {
		s.logger.Errorf("error walking the path %q: %v\n", s.path, err)
	}

	fmt.Println(res)
	return res
}

func validExt(path string) bool {
	switch filepath.Ext(path) {
	case ".tar", ".gz", ".whl":
		return true
	}
	return false
}

func (s Storage) String() string {
	return fmt.Sprintf("Storage[directory=%q]", s.path)
}
