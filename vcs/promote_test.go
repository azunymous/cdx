package vcs

import (
	"cdx/internal/check"
	"github.com/go-git/go-billy/v5/memfs"
	"testing"
)

func TestPromoteTag(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	createVersionTag(fs, "app-0.1.0")

	repo := newTestRepo(fs)
	err := repo.Promote("app", "promoted")
	check.Ok(t, err)
	err = tagExistsAtHead(fs, "app-0.1.0-promoted")
	check.Ok(t, err)
}

func TestPromoteTagSucceedsIfAlreadyPromoted(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	createVersionTag(fs, "app-0.1.0")
	createVersionTag(fs, "app-0.1.0-promoted")

	repo := newTestRepo(fs)
	err := repo.Promote("app", "promoted")
	check.Ok(t, err)
	err = tagExistsAtHead(fs, "app-0.1.0-promoted")
	check.Ok(t, err)
}

func TestPromoteFailsWhenNoBaseTag(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	createVersionTag(fs, "app-0.1.0")
	createCommit(fs, "New Version", "Hello world 2")

	repo := newTestRepo(fs)
	err := repo.Promote("app", "promoted")
	check.Assert(t, err != nil, "expecting error to not be nil, got %v", err)
}
