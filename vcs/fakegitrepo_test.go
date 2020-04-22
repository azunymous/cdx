package vcs

import "cdx/versioned"

type FakeGitRepo struct {
	isOnMaster         bool
	passedIncrementTag func() (string, versioned.Field)
	passedStringString func() (string, string)
	passedErr          error
	passedModuleErr    error
	passedHeadErr      error
	emptyModuleTags    bool
	emptyHeadTags      bool
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
	return f.passedErr
}

func (f *FakeGitRepo) Promote(app, stage string) error {
	f.passedStringString = func() (string, string) {
		return app, stage
	}
	return f.passedErr
}

func (f *FakeGitRepo) PushTags() error {
	f.pushed = true
	return f.pushTagsErr
}

func (f *FakeGitRepo) TagsForHead(app string, stage ...string) ([]string, error) {
	s := ""
	if len(stage) > 0 {
		s = stage[0]
	}
	f.passedStringString = func() (string, string) {
		return app, s
	}
	if f.emptyHeadTags {
		return nil, nil
	}
	return []string{"head-tag-0.0.0"}, f.passedHeadErr
}

func (f *FakeGitRepo) TagsForModule(app string, stage ...string) ([]string, error) {
	s := ""
	if len(stage) > 0 {
		s = stage[0]
	}
	f.passedStringString = func() (string, string) {
		return app, s
	}
	if f.emptyModuleTags {
		return nil, nil
	}
	return []string{"module-tag-1.1.1"}, f.passedModuleErr
}
