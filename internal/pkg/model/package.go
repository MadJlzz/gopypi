package model

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

var normalizeRegexp = regexp.MustCompile("[-_.]+")

type Package struct {
	Name           string
	URI            []*url.URL
	normalizedName string
}

func New(name string, URI ...*url.URL) *Package {
	return &Package{
		Name:           name,
		URI:            URI,
		normalizedName: normalize(name),
	}
}

func (p *Package) NormalizeName() string {
	return p.normalizedName
}

func normalize(URI string) string {
	return strings.ToLower(normalizeRegexp.ReplaceAllString(URI, "-"))
}

func (p *Package) String() string {
	return fmt.Sprintf("Package: [%#v]\n", p)
}
