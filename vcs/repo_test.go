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

func createGitRepo(fs billy.Filesystem) {
	r, _ := git.Init(filesystem2.NewStorage(fs, cache.NewObjectLRUDefault()), fs)
	w, _ := r.Worktree()

	// create a new file inside of the worktree of the project using the go standard library.
	file, _ := fs.Create("example-git-file")
	_, _ = file.Write([]byte("hello world!"))

	// Adds the new file to the staging area.
	_, _ = w.Add("example-git-file")

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
