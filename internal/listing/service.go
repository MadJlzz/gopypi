package listing

import "context"

type Repository interface {
	// GetAllProjects return all unique projects stored in the repository.
	GetAllProjects(ctx context.Context) []Project

	// GetAllProjectPackages returns all packages saved in storage.
	GetAllProjectPackages(ctx context.Context, project string) []Package
}
