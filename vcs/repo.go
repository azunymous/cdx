package vcs

import (
	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/cache"
	filesystem2 "github.com/go-git/go-git/v5/storage/filesystem"
	"github.com/sirupsen/logrus"
	"regexp"
	"sort"
)

// Repo is a VCS repository that can be manipulated
type Repo struct {
	gitRepo *git.Repository
	log     logrus.StdLogger
}

// NewRepo returns a new repository from the given filesystem
func NewRepo(fs billy.Filesystem) (*Repo, error) {
	var gr *git.Repository
	var err error
	// This is currently done for testing. TODO remove non plain open
	if fs == nil {
		gr, err = git.PlainOpenWithOptions(".", &git.PlainOpenOptions{DetectDotGit: true})
	} else {
		gr, err = git.Open(filesystem2.NewStorage(fs, cache.NewObjectLRUDefault()), fs)
	}

	if err != nil {
		return nil, err
	}
	return &Repo{gitRepo: gr, log: logrus.New()}, nil
}

// TagsForHead returns sorted all tags at HEAD
func (r *Repo) TagsForHead() ([]string, error) {
	current, err := r.gitRepo.ResolveRevision("HEAD")
	if err != nil {
		return nil, err
	}
	tags, err := r.gitRepo.Tags()
	if err != nil {
		return nil, err
	}

	var t []string
	_ = tags.ForEach(func(reference *plumbing.Reference) error {
		if reference.Hash() == *current {
			t = append(t, reference.Name().Short())
		}
		return nil
	})
	sort.Strings(t)
	return t, nil
}

// TagsForModule returns sorted all semantic version tags for a module and a promotion stage.
// If no promotion stage is provided, only unpromoted tags are returned.
// Only the first provided promotion stage is used for filtering.
func (r *Repo) TagsForModule(module string, stage ...string) ([]string, error) {
	suffix := ""
	if len(stage) > 0 {
		suffix = "-" + stage[0]
	}

	tags, err := r.gitRepo.Tags()
	if err != nil {
		return nil, err
	}

	var t []string
	regex, err := regexp.Compile("^" + module + "-[0-9]+\\.[0-9]+\\.[0-9]+" + suffix + "$")
	if err != nil {
		return nil, err
	}
	_ = tags.ForEach(func(reference *plumbing.Reference) error {
		if regex.MatchString(reference.Name().Short()) {
			t = append(t, reference.Name().Short())
		}
		return nil
	})
	sort.Strings(t)
	return t, nil
}
