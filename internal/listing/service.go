package listing

import "context"

type Repository interface {
	// GetAllPackages returns all packages saved in storage.
	GetAllPackages(ctx context.Context) []PackageReference
}
