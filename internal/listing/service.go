package listing

import "context"

type Repository interface {
	//// GetAllPackages returns all packages saved in storage.
	//GetAllPackages(ctx context.Context) []PackageReference

	// GetAllProjects return all unique projects stored in the repository.
	GetAllProjects(ctx context.Context) []Project
}
