package vcs

import (
	"cdx/parse"
	"cdx/versioned"
	"errors"
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

func (g *Git) Version(stage string, headOnly bool) (string, error) {
	if headOnly {
		return g.getHeadTag(stage)
	}
	return g.getModuleTag(stage)
}

func (g *Git) Distribute() error {
	if g.push {
		return g.r.PushTags()
	}
	return nil
}

func (g *Git) getHeadTag(stage string) (string, error) {
	tagsForHead, err := g.r.TagsForHead(g.app, stage)
	if err != nil {
		return "", err
	}
	if len(tagsForHead) == 0 {
		return "", errors.New("no tags found at HEAD")
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
