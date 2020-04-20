package vcs

import (
	"cdx/internal/check"
	"github.com/go-git/go-billy/v5/memfs"
	"testing"
)

func TestIncrementTag(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	createVersionTag(fs, "app-0.1.0")
	createCommit(fs, "New Version", "Hello world 2")

	repo, err := NewRepo(fs)
	check.Ok(t, err)
	err = repo.IncrementTag("app", Minor)
	check.Ok(t, err)
	err = tagExistsAtHead(fs, "app-0.2.0")
	check.Ok(t, err)
}

func TestIncrementTagDifferentField(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	createVersionTag(fs, "app-1.1.1")
	createCommit(fs, "New Version", "Hello world 2")

	repo, err := NewRepo(fs)
	check.Ok(t, err)
	err = repo.IncrementTag("app", Major)
	check.Ok(t, err)
	err = tagExistsAtHead(fs, "app-2.0.0")
	check.Ok(t, err)
}

func TestIncrementTagCreatesNewTagWhenNoTagsExist(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)

	repo, err := NewRepo(fs)
	check.Ok(t, err)
	err = repo.IncrementTag("app", Minor)
	check.Ok(t, err)
	err = tagExistsAtHead(fs, "app-0.1.0")
	check.Ok(t, err)
}
