package listing

import (
	"crypto/sha256"
	"encoding/hex"
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

func (p Package) HexEncodedHash() string {
	hash := sha256Hash([]byte(p.URI))
	return hex.EncodeToString(hash)
}

func sha256Hash(data []byte) []byte {
	h := sha256.New()
	h.Write(data)
	return h.Sum(nil)
}