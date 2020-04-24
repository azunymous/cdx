package parse

import "regexp"

var version = regexp.MustCompile(`[0-9]+\.[0-9]+\.[0-9]+`)

// Version returns the semantic version (X.Y.Z) from a tag, returning an empty string if not found.
func Version(tag string) string {
	return version.FindString(tag)
}
