package listing

type Repository interface {
	// GetAllPackages returns all packages saved in storage.
	GetAllPackages() []Package
}
