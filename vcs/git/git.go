package git

import (
	"cdx/versioned"
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
	PushTags() error
}

type repoF func() (Repository, error)

func New(app string, field versioned.Field, push bool, r repoF) (*Git, error) {
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
	panic("implement me")
}

func (g *Git) Version(stage string, headOnly bool) (string, error) {
	panic("implement me")
}

func (g *Git) Distribute() error {
	if g.push {
		return g.r.PushTags()
	}
	return nil
}
