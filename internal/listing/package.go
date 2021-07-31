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

// normalize apply normalizeRegexp for being PEP503 compliant.
func normalize(name string) string {
	return strings.ToLower(normalizeRegexp.ReplaceAllString(name, "-"))
}

type Project string

// Normalize simple getter of Package.
func (p Project) Normalize() string {
	return normalize(string(p))
}

type Package struct {
	Filename string
	URI      string
}
