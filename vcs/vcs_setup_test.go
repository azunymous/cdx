package vcs

import (
	"errors"
	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/cache"
	"github.com/go-git/go-git/v5/plumbing/object"
	filesystem2 "github.com/go-git/go-git/v5/storage/filesystem"
	"github.com/sirupsen/logrus"
	"time"
)

// newTestRepo is a test version of NewRepo which creates an in memory repository
func newTestRepo(fs billy.Filesystem) *Repo {
	gr, err := git.Open(filesystem2.NewStorage(fs, cache.NewObjectLRUDefault()), fs)
	if err != nil {
		panic(err)
	}
	return &Repo{gitRepo: gr, log: logrus.New()}

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

func tagExistsAtHead(fs billy.Filesystem, tag string) error {
	r, _ := git.Open(filesystem2.NewStorage(fs, cache.NewObjectLRUDefault()), fs)
	revision, _ := r.ResolveRevision("HEAD")
	t, err := r.Tag(tag)
	if err != nil {
		return err
	}
	if t.Hash() != *revision {
		return errors.New("tag not at head")
	}
	return nil
}
