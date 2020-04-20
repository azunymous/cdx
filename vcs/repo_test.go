package vcs

import (
	"cdx/internal/check"
	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/cache"
	"github.com/go-git/go-git/v5/plumbing/object"
	filesystem2 "github.com/go-git/go-git/v5/storage/filesystem"
	"testing"
	"time"
)

func TestNewRepo(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)

	repo, err := NewRepo(fs)
	check.Ok(t, err)
	log, err := repo.gitRepo.Log(&git.LogOptions{})
	check.Ok(t, err)
	_, err = log.Next()
	check.Ok(t, err)
}

func TestTagsForHead(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	createVersionTag(fs, "app-0.0.1")

	repo, err := NewRepo(fs)
	check.Ok(t, err)
	tagsForHead, err := repo.TagsForHead()
	check.Ok(t, err)
	check.Equals(t, 1, len(tagsForHead))

	_, ok := tagsForHead["app-0.0.1"]
	check.Assert(t, ok, "Expected 'app-0.0.1' in map")
}

func TestTagsForHeadForMultipleTags(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	createVersionTag(fs, "app-0.0.1-alpha")
	createVersionTag(fs, "app-0.0.1")

	repo, err := NewRepo(fs)
	check.Ok(t, err)
	tagsForHead, err := repo.TagsForHead()
	check.Ok(t, err)
	check.Equals(t, 2, len(tagsForHead))
	check.Equals(t, map[string]struct{}{"app-0.0.1-alpha": {}, "app-0.0.1": {}}, tagsForHead)
}

func TestTagsForHeadForMultipleCommits(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	createVersionTag(fs, "app-0.0.1")
	createCommit(fs, "New Version", "Hello world 2")
	createVersionTag(fs, "app-0.0.2")

	repo, err := NewRepo(fs)
	check.Ok(t, err)
	tagsForHead, err := repo.TagsForHead()
	check.Ok(t, err)
	check.Equals(t, 1, len(tagsForHead))

	_, ok := tagsForHead["app-0.0.2"]
	check.Assert(t, ok, "Expected 'app-0.0.2' in map")
}

func TestTagsForHeadWhenEmpty(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)

	repo, err := NewRepo(fs)
	check.Ok(t, err)
	tagsForHead, err := repo.TagsForHead()
	check.Ok(t, err)
	check.Equals(t, 0, len(tagsForHead))
}

func createGitRepo(fs billy.Filesystem) {
	_, _ = git.Init(filesystem2.NewStorage(fs, cache.NewObjectLRUDefault()), fs)
	createCommit(fs, "example-git-file", "hello world!")
}

func createCommit(fs billy.Filesystem, filename, msg string) {
	r, _ := git.Open(filesystem2.NewStorage(fs, cache.NewObjectLRUDefault()), fs)
	w, _ := r.Worktree()
	// create a new file inside of the worktree of the project using the go standard library.
	file, _ := fs.Create(filename)
	_, _ = file.Write([]byte(msg))

	// Adds the new file to the staging area.
	_, _ = w.Add(filename)

	// Commits the current staging area to the repository, with the new file
	// just created.
	_, _ = w.Commit("example go-git commit", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "John Doe",
			Email: "john@doe.org",
			When:  time.Now(),
		},
	})
}

func createVersionTag(fs billy.Filesystem, tag string) {
	r, _ := git.Open(filesystem2.NewStorage(fs, cache.NewObjectLRUDefault()), fs)
	revision, _ := r.ResolveRevision("HEAD")
	_, _ = r.CreateTag(tag, *revision, nil)
}
