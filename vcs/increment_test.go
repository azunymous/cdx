package vcs

import (
	"cdx/test/check"
	"github.com/go-git/go-billy/v5/memfs"
	"testing"
)

func TestIncrementTag(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	createVersionTag(fs, "app-0.1.0")
	createCommit(fs, "New Version", "Hello world 2")

	repo := newTestRepo(fs)
	err := repo.IncrementTag("app", Minor)
	check.Ok(t, err)
	err = tagExistsAtHead(fs, "app-0.2.0")
	check.Ok(t, err)
}

func TestIncrementTagAlreadyExists(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	createVersionTag(fs, "app-0.1.0")

	repo := newTestRepo(fs)
	err := repo.IncrementTag("app", Minor)
	check.Ok(t, err)
	err = tagExistsAtHead(fs, "app-0.1.0")
	check.Ok(t, err)
	err = tagDoesNotExist(fs, "app-0.2.0")
	check.Ok(t, err)
}

func TestIncrementTagDifferentField(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	createVersionTag(fs, "app-1.1.1")
	createCommit(fs, "New Version", "Hello world 2")

	repo := newTestRepo(fs)
	err := repo.IncrementTag("app", Major)
	check.Ok(t, err)
	err = tagExistsAtHead(fs, "app-2.0.0")
	check.Ok(t, err)
}

func TestIncrementTagCreatesNewTagWhenNoTagsExist(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)

	repo := newTestRepo(fs)
	err := repo.IncrementTag("app", Minor)
	check.Ok(t, err)
	err = tagExistsAtHead(fs, "app-0.1.0")
	check.Ok(t, err)
}
