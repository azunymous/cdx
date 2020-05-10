package vcs

import (
	"errors"
	"github.com/azunymous/cdx/parse"
	"github.com/azunymous/cdx/versioned"
	"github.com/sirupsen/logrus"
)

// Git is a VCS with an application that needs to be versioned via git
type Git struct {
	app   string
	field versioned.Field
	r     Repository
	push  bool
}

type Repository interface {
	OnMaster() bool
	IncrementTag(name string, field versioned.Field) error
	Promote(app, stage string) error
	TagsForHead(app string, stage ...string) ([]string, error)
	TagsForModule(app string, stage ...string) ([]string, error)
	PushTags() error
	HeadHash() (string, error)
}

type repoF func() (Repository, error)

func NewGit(app string, field versioned.Field, push bool, r repoF) (*Git, error) {
	repo, err := r()
	if err != nil {
		return nil, err
	}
	return &Git{
		app:   app,
		field: field,
		r:     repo,
		push:  push,
	}, nil
}

func (g *Git) Ready() bool {
	if g.push && !g.r.OnMaster() {
		logrus.Println("Not on origin/master, continuing")
		return false
	}
	return true
}

func (g *Git) Release() error {
	return g.r.IncrementTag(g.app, g.field)
}

func (g *Git) Promote(stage string) error {
	return g.r.Promote(g.app, stage)
}

// Version returns a semantic version from the git repository by parsing git tags
//
// Depending on the provided modifiers, version scans for different types of tags:
//
// - stage: if provided, scans only for that stage. If not provided, scans only for non-promoted tags.
// - headOnly: only scans HEAD for tags
// - useHash: falls back to the hash of HEAD if no tags found and headOnly is true
func (g *Git) Version(stage string, headOnly bool, useHash bool) (string, error) {
	if headOnly {
		tag, err := g.getHeadTag(stage)
		if errors.Is(err, errNoTagsFoundAtHead) && useHash {
			return g.r.HeadHash()
		}
		return tag, err
	}
	return g.getModuleTag(stage)
}

func (g *Git) Distribute() error {
	if g.push {
		return g.r.PushTags()
	}
	return nil
}

var errNoTagsFoundAtHead = errors.New("no tags found at HEAD")

func (g *Git) getHeadTag(stage string) (string, error) {
	tagsForHead, err := g.r.TagsForHead(g.app, stage)
	if err != nil {
		return "", err
	}
	if len(tagsForHead) == 0 {
		return "", errNoTagsFoundAtHead
	}
	return parse.Version(tagsForHead[len(tagsForHead)-1]), nil
}

func (g *Git) getModuleTag(stage string) (string, error) {
	tagsForModule, err := g.r.TagsForModule(g.app, stage)
	if err != nil {
		return "", err
	}
	if len(tagsForModule) == 0 {
		return "", errors.New("no tags found for module and stage")
	}
	return parse.Version(tagsForModule[len(tagsForModule)-1]), nil
}
