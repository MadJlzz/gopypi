package registry

type Registry interface {
	// GetAllProjects return all unique projects stored in the repository.
	GetAllProjects() []Project

	// GetAllProjectPackages returns all packages saved in storage.
	GetAllProjectPackages(project string) []Package
}
