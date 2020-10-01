package model

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

var normalizeRegexp = regexp.MustCompile("[-_.]+")

type PackageFile struct {
	Name string
	URL  *url.URL
}

type Package struct {
	Name           string
	packageFiles   []*PackageFile
	normalizedName string
}

func New(name string, packageFiles ...*PackageFile) *Package {
	return &Package{
		Name:           name,
		packageFiles:   packageFiles,
		normalizedName: normalize(name),
	}
}

func (p *Package) AppendPackageFile(pf *PackageFile) {
	p.packageFiles = append(p.packageFiles, pf)
}

func (p *Package) NormalizeName() string {
	return p.normalizedName
}

func (p *Package) PackageFiles() []*PackageFile {
	return p.packageFiles
}

func (p *Package) String() string {
	return fmt.Sprintf("Package: [%#v]\n", p)
}

func normalize(URI string) string {
	return strings.ToLower(normalizeRegexp.ReplaceAllString(URI, "-"))
}