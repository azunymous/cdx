package vcs

import (
	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/cache"
	filesystem2 "github.com/go-git/go-git/v5/storage/filesystem"
)

type Repo struct {
	gitRepo *git.Repository
}

func NewRepo(fs billy.Filesystem) (*Repo, error) {
	gr, err := git.Open(filesystem2.NewStorage(fs, cache.NewObjectLRUDefault()), fs)
	if err != nil {
		return nil, err
	}
	return &Repo{gitRepo: gr}, nil
}
