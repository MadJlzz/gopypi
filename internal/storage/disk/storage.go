package disk

import (
	"fmt"
	"github.com/MadJlzz/gopypi/internal/listing"
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
	var pkgs []listing.Package

	err := filepath.Walk(s.path, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			s.logger.Warnf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() {
			return nil
		}
		pkgs = append(pkgs, listing.Package(path))
		return nil
	})
	if err != nil {
		s.logger.Errorf("error walking the path %q: %v\n", s.path, err)
	}
	return pkgs
}

func (s Storage) String() string {
	return fmt.Sprintf("DiskStorage[directory=%q]", s.path)
}
