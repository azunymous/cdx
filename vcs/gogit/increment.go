package gogit

import (
	"cdx/versioned"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var version = regexp.MustCompile(`[0-9]+\.[0-9]+\.[0-9]+`)

// Increment increases the version number tag for a module and creates the new tag.
func (r *Repo) IncrementTag(module string, field versioned.Field) error {
	tagsForHead, err := r.TagsForHead(module)
	if err != nil {
		return err
	}

	if len(tagsForHead) > 0 {
		r.log.Printf("HEAD already tagged with %s, continuing", tagsForHead[0])
		return nil
	}

	tagsForModule, err := r.TagsForModule(module)
	if err != nil {
		return err
	}
	if len(tagsForModule) == 0 {
		tagsForModule = []string{module + "-0.0.0"}
	}

	latest := tagsForModule[len(tagsForModule)-1]
	n, err := increase(latest, field)
	if err != nil {
		return err
	}
	tag := module + "-" + n
	r.log.Printf("Incrementing latest version %s -> %s", latest, tag)

	revision, err := r.gitRepo.ResolveRevision("HEAD")
	if err != nil {
		return err
	}

	_, err = r.gitRepo.CreateTag(tag, *revision, nil)
	return err
}

// increase takes a semver tag (see version regex) and bumps the given field returning the incremented X.Y.Z
// Note: this can take a semver tag string with a module but only returns the semantic version.
func increase(latest string, field versioned.Field) (string, error) {
	v := VersionFrom(latest)
	if v == "" {
		return "", fmt.Errorf("could not find version in tag: %s", latest)
	}

	split := strings.Split(v, ".")
	num, err := strconv.Atoi(split[field])
	if err != nil {
		return "", err
	}
	split[field] = strconv.Itoa(num + 1)
	if field < versioned.Patch {
		split[versioned.Patch] = "0"
	}
	if field < versioned.Minor {
		split[versioned.Minor] = "0"
	}
	return strings.Join(split, "."), nil
}

// VersionFrom returns the semantic version (X.Y.Z) from a tag, returning an empty string if not found.
func VersionFrom(tag string) string {
	return version.FindString(tag)
}
