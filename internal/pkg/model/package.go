package model

import (
	"fmt"
	"regexp"
	"strings"
)

// normalizeRegexp represents the normalized name rules coming
// from PEP503.
//
// see [PEP503](https://www.python.org/dev/peps/pep-0503/#normalized-names)
var normalizeRegexp = regexp.MustCompile("[-_.]+")

// PackageFile represent a single file contained by a Package.
//
// Should be instantiated using NewPackageFile.
type PackageFile struct {
	name      string
	signedURL string
}

// NewPackageFile is the simplest way to get started with a PackageFile.
func NewPackageFile(name, signedURL string) *PackageFile {
	return &PackageFile{
		name:      name,
		signedURL: signedURL,
	}
}

// Name simple getter of PackageFile
func (pf *PackageFile) Name() string {
	return pf.name
}

// SignedURL simple getter of PackageFile
func (pf *PackageFile) SignedURL() string {
	return pf.signedURL
}

// Package is the core model of gopypi.
//
// It stores any reference to files and mimic the
// representation of a Python package.
//
// Should be instantiated using NewPackage.
//
// e.g.
//
// example-pkg/
//   example-pkg-0.0.1.tar.gz
//   example_pkg-0.0.1-py3-none-any.whl
type Package struct {
	name           string
	packageFiles   []*PackageFile
	normalizedName string
}

// NewPackage is the simplest way to get started with a Package.
func NewPackage(name string, packageFiles ...*PackageFile) *Package {
	return &Package{
		name:           name,
		packageFiles:   packageFiles,
		normalizedName: normalize(name),
	}
}

// AppendPackageFile adds a new PackageFile to a Package.
func (p *Package) AppendPackageFile(pf *PackageFile) {
	p.packageFiles = append(p.packageFiles, pf)
}

// normalize apply normalizeRegexp for being PEP503 compliant.
func normalize(URI string) string {
	return strings.ToLower(normalizeRegexp.ReplaceAllString(URI, "-"))
}

// Name simple getter of Package
func (p *Package) Name() string {
	return p.name
}

// NormalizeName simple getter of Package.
func (p *Package) NormalizeName() string {
	return p.normalizedName
}

// PackageFiles simple getter of Package.
func (p *Package) PackageFiles() []*PackageFile {
	return p.packageFiles
}

// String redefines how we print Package.
func (p *Package) String() string {
	return fmt.Sprintf("Package: [%#v]\n", p)
}
