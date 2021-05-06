package listing

import (
	"go.uber.org/zap"
)

type Repository interface {
	// GetAllPackages returns all packages saved in storage.
	GetAllPackages() []Package
}

type Service interface {
	GetAllPackages() []Package
}

type service struct {
	logger *zap.SugaredLogger
	r      Repository
}

func NewService(logger *zap.SugaredLogger, repository Repository) Service {
	return &service{
		logger: logger,
		r:      repository,
	}
}

func (s service) GetAllPackages() []Package {
	return s.r.GetAllPackages()
}
