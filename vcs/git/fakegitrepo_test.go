package git

import "cdx/versioned"

type FakeGitRepo struct {
	isOnMaster         bool
	passedIncrementTag func() (string, versioned.Field)
	incrementTagErr    error
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

func (f *FakeGitRepo) PushTags() error {
	f.pushed = true
	return f.pushTagsErr
}
