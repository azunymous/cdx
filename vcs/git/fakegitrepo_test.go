package git

import "cdx/versioned"

type FakeGitRepo struct {
	isOnMaster         bool
	passedIncrementTag func() (string, versioned.Field)
	incrementTagErr    error
	passedPromote      func() (string, string)
	promoteErr         error
	pushed             bool
	pushTagsErr        error
}

func (f *FakeGitRepo) OnMaster() bool {
	return f.isOnMaster
}

func (f *FakeGitRepo) IncrementTag(name string, field versioned.Field) error {
	f.passedIncrementTag = func() (string, versioned.Field) {
		return name, field
	}
	return f.incrementTagErr
}

func (f *FakeGitRepo) Promote(app, stage string) error {
	f.passedPromote = func() (string, string) {
		return app, stage
	}
	return f.promoteErr
}

func (f *FakeGitRepo) PushTags() error {
	f.pushed = true
	return f.pushTagsErr
}
