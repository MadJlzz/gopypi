package web

import (
	"regexp"
	"strings"
)

var normalizeRegexp = regexp.MustCompile("[-_.]+")

type Package struct {
	Name          string
	URI           string
	normalizedUri string
}

func New(name, uri string) *Package {
	return &Package{
		Name:          name,
		URI:           uri,
		normalizedUri: normalize(uri),
	}
}

func (p *Package) NormalizeURI() string {
	return p.normalizedUri
}

func normalize(URI string) string {
	return strings.ToLower(normalizeRegexp.ReplaceAllString(URI, "-"))
}
