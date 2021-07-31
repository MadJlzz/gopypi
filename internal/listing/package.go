package listing

import (
	"regexp"
	"strings"
)

// normalizeRegexp represents the normalized name rules coming
// from PEP503.
//
// see [PEP503](https://www.python.org/dev/peps/pep-0503/#normalized-names)
var normalizeRegexp = regexp.MustCompile("[-_.]+")

type PackageReference string

// normalize apply normalizeRegexp for being PEP503 compliant.
func normalize(name string) string {
	return strings.ToLower(normalizeRegexp.ReplaceAllString(name, "-"))
}

// Normalize simple getter of Package.
func (p PackageReference) Normalize() string {
	return normalize(string(p))
}