package vcs

import (
	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
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

func (r *Repo) TagsForHead() (map[string]struct{}, error) {
	current, err := r.gitRepo.ResolveRevision("HEAD")
	if err != nil {
		return nil, err
	}
	tags, err := r.gitRepo.Tags()
	if err != nil {
		return nil, err
	}

	t := map[string]struct{}{}
	_ = tags.ForEach(func(reference *plumbing.Reference) error {
		if reference.Hash() == *current {
			t[reference.Name().Short()] = struct{}{}
		}
		return nil
	})
	return t, nil
}
