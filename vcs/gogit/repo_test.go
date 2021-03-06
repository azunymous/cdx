package gogit

import (
	"github.com/azunymous/cdx/test/check"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"strconv"
	"testing"
)

// This test assumes that it is running from a git repository and asserts that it can be parsed
func TestNewRepo(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)

	repo, err := NewRepo()
	check.Ok(t, err)
	log, err := repo.gitRepo.Log(&git.LogOptions{})
	check.Ok(t, err)
	_, err = log.Next()
	check.Ok(t, err)
}

func TestTagsForHead(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	createVersionTag(fs, "app-0.0.1")

	repo := newTestRepo(fs)
	tagsForHead, err := repo.TagsForHead("app")
	check.Ok(t, err)
	check.Equals(t, 1, len(tagsForHead))
	check.Equals(t, []string{"app-0.0.1"}, tagsForHead)
}

func TestTagsForHeadAnnotated(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	createAnnotatedVersionTag(fs, "app-0.0.1")

	repo := newTestRepo(fs)
	tagsForHead, err := repo.TagsForHead("app")
	check.Ok(t, err)
	check.Equals(t, 1, len(tagsForHead))
	check.Equals(t, []string{"app-0.0.1"}, tagsForHead)
}

func TestTagsForHeadForAppWithDashes(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	createVersionTag(fs, "my-app-0.0.1")

	repo := newTestRepo(fs)
	tagsForHead, err := repo.TagsForHead("my-app")
	check.Ok(t, err)
	check.Equals(t, 1, len(tagsForHead))
	check.Equals(t, []string{"my-app-0.0.1"}, tagsForHead)
}

func TestTagsForHeadForAppWithOverlappingNamePrefix(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	createVersionTag(fs, "my-app-0.0.1")
	createVersionTag(fs, "app-0.0.2")

	repo := newTestRepo(fs)
	tagsForHead, err := repo.TagsForHead("my-app")
	check.Ok(t, err)
	check.Equals(t, 1, len(tagsForHead))
	check.Equals(t, []string{"my-app-0.0.1"}, tagsForHead)
}

func TestTagsForHeadForAppWithOverlappingNameSuffix(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	createVersionTag(fs, "app-two-0.0.1")
	createVersionTag(fs, "app-0.0.2")

	repo := newTestRepo(fs)
	tagsForHead, err := repo.TagsForHead("app-two")
	check.Ok(t, err)
	check.Equals(t, 1, len(tagsForHead))
	check.Equals(t, []string{"app-two-0.0.1"}, tagsForHead)
}

func TestTagsForHeadForAppWithOverlappingNameNumberSuffix(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	createVersionTag(fs, "app-2-0.0.1")
	createVersionTag(fs, "app-0.0.2")

	repo := newTestRepo(fs)
	tagsForHead, err := repo.TagsForHead("app-2")
	check.Ok(t, err)
	check.Equals(t, 1, len(tagsForHead))
	check.Equals(t, []string{"app-2-0.0.1"}, tagsForHead)
}

func TestTagsForHeadForAppWithOverlappingName(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	createVersionTag(fs, "my-app-0.0.2")
	createVersionTag(fs, "app-0.0.1")

	repo := newTestRepo(fs)
	tagsForHead, err := repo.TagsForHead("app")
	check.Ok(t, err)
	check.Equals(t, 1, len(tagsForHead))
	check.Equals(t, []string{"app-0.0.1"}, tagsForHead)
}

func TestTagsForHead_emptyStage(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	createVersionTag(fs, "app-0.0.1")

	repo := newTestRepo(fs)
	tagsForHead, err := repo.TagsForHead("app", "")
	check.Ok(t, err)
	check.Equals(t, 1, len(tagsForHead))
	check.Equals(t, []string{"app-0.0.1"}, tagsForHead)
}

func TestTagsForHeadForMultipleTags(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	createVersionTag(fs, "app-0.0.2")
	createVersionTag(fs, "app-0.0.1")

	repo := newTestRepo(fs)
	tagsForHead, err := repo.TagsForHead("app")
	check.Ok(t, err)
	check.Equals(t, 2, len(tagsForHead))
	check.Equals(t, []string{"app-0.0.1", "app-0.0.2"}, tagsForHead)
}

func TestTagsForHeadAreSorted(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	var expectedTags []string
	for i := 0; i < 10; i++ {
		is := strconv.Itoa(i)
		tag := "app-0.0." + is
		createVersionTag(fs, tag)
		expectedTags = append(expectedTags, tag)
	}

	repo := newTestRepo(fs)
	tagsForHead, err := repo.TagsForHead("app")
	check.Ok(t, err)
	check.Equals(t, 10, len(tagsForHead))
	check.Equals(t, expectedTags, tagsForHead)
}

func TestTagsForHeadForMultipleCommits(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	createVersionTag(fs, "app-0.0.1")
	createCommit(fs, "New Version", "Hello world 2")
	createVersionTag(fs, "app-0.0.2")

	repo := newTestRepo(fs)
	tagsForHead, err := repo.TagsForHead("app")
	check.Ok(t, err)
	check.Equals(t, 1, len(tagsForHead))

	check.Equals(t, []string{"app-0.0.2"}, tagsForHead)
}

func TestTagsForHeadWhenNone(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)

	repo := newTestRepo(fs)
	tagsForHead, err := repo.TagsForHead("app")
	check.Ok(t, err)
	check.Equals(t, 0, len(tagsForHead))
}

func TestTagsForHeadWhenPartialModuleNameMatch(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	createVersionTag(fs, "ap-0.0.1")

	repo := newTestRepo(fs)
	tagsForHead, err := repo.TagsForHead("app")
	check.Ok(t, err)
	check.Equals(t, 0, len(tagsForHead))
}

func TestTagsForHeadWhenIncorrectVersioning(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	createVersionTag(fs, "app-.0.1")

	repo := newTestRepo(fs)
	tagsForHead, err := repo.TagsForHead("app")
	check.Ok(t, err)
	check.Equals(t, 0, len(tagsForHead))
}

func TestTagsForHeadWhenDashSeparatedPartialModuleNameMatch(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	createVersionTag(fs, "another-app-0.0.1")

	repo := newTestRepo(fs)
	tagsForHead, err := repo.TagsForHead("app")
	check.Ok(t, err)
	check.Equals(t, 0, len(tagsForHead))
}

func TestTagsForHeadWhenOnlyPromoted(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	createVersionTag(fs, "app-0.0.1+promoted")

	repo := newTestRepo(fs)
	tagsForHead, err := repo.TagsForHead("app")
	check.Ok(t, err)
	check.Equals(t, 0, len(tagsForHead))
}

func TestTagsForHead_stage(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	createVersionTag(fs, "app-0.0.1")
	createVersionTag(fs, "app-0.0.1+promoted")

	repo := newTestRepo(fs)
	tagsForHead, err := repo.TagsForHead("app", "promoted")
	check.Ok(t, err)
	check.Equals(t, 1, len(tagsForHead))
	check.Equals(t, []string{"app-0.0.1+promoted"}, tagsForHead)
}

func TestTagsForHeadForMultipleCommits_stage(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	createVersionTag(fs, "app-0.0.1")
	createVersionTag(fs, "app-0.0.1+promoted")
	createCommit(fs, "New Version", "Hello world 2")
	createVersionTag(fs, "app-0.0.2")
	createVersionTag(fs, "app-0.0.2+promoted")

	repo := newTestRepo(fs)
	tagsForHead, err := repo.TagsForHead("app", "promoted")
	check.Ok(t, err)
	check.Equals(t, 1, len(tagsForHead))
	check.Equals(t, []string{"app-0.0.2+promoted"}, tagsForHead)
}

func TestTagsForHeadAreSorted_stage(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	var expectedTags []string
	for i := 0; i < 10; i++ {
		is := strconv.Itoa(i)
		tag := "app-0.0." + is + "+promoted"
		createVersionTag(fs, tag)
		expectedTags = append(expectedTags, tag)
	}

	repo := newTestRepo(fs)
	tagsForHead, err := repo.TagsForHead("app", "promoted")
	check.Ok(t, err)
	check.Equals(t, 10, len(tagsForHead))
	check.Equals(t, expectedTags, tagsForHead)
}

func TestTagsForHeadWhenNone_stage(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)

	repo := newTestRepo(fs)
	tagsForHead, err := repo.TagsForHead("app", "promoted")
	check.Ok(t, err)
	check.Equals(t, 0, len(tagsForHead))
}

func TestTagsForHeadWhenOnlyUnpromoted_stage(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	createVersionTag(fs, "app-0.0.1")

	repo := newTestRepo(fs)
	tagsForHead, err := repo.TagsForHead("app", "promoted")
	check.Ok(t, err)
	check.Equals(t, 0, len(tagsForHead))
}

func TestTagsForHeadWhenPartialMatch_stage(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	createVersionTag(fs, "ap-0.0.1")

	repo := newTestRepo(fs)
	tagsForHead, err := repo.TagsForHead("app", "promoted")
	check.Ok(t, err)
	check.Equals(t, 0, len(tagsForHead))
}

func TestTagsForHeadWhenIncorrectVersioning_stage(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	createVersionTag(fs, "app-.0.1")

	repo := newTestRepo(fs)
	tagsForHead, err := repo.TagsForHead("app", "promoted")
	check.Ok(t, err)
	check.Equals(t, 0, len(tagsForHead))
}

func TestTagsForModule(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	createVersionTag(fs, "app-0.0.1")

	repo := newTestRepo(fs)
	tagsForModule, err := repo.TagsForModule("app")
	check.Ok(t, err)
	check.Equals(t, 1, len(tagsForModule))
	check.Equals(t, []string{"app-0.0.1"}, tagsForModule)
}

func TestAnnotatedTagsForModule(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	createAnnotatedVersionTag(fs, "app-0.0.1")

	repo := newTestRepo(fs)
	tagsForModule, err := repo.TagsForModule("app")
	check.Ok(t, err)
	check.Equals(t, 1, len(tagsForModule))
	check.Equals(t, []string{"app-0.0.1"}, tagsForModule)
}

func TestTagsForModuleWithOverlappingNamePrefix(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	createVersionTag(fs, "my-app-0.0.1")
	createVersionTag(fs, "app-0.0.2")

	repo := newTestRepo(fs)
	tagsForModule, err := repo.TagsForModule("my-app")
	check.Ok(t, err)
	check.Equals(t, 1, len(tagsForModule))
	check.Equals(t, []string{"my-app-0.0.1"}, tagsForModule)
}

func TestTagsForModuleWithOverlappingNameSuffix(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	createVersionTag(fs, "app-two-0.0.1")
	createVersionTag(fs, "app-0.0.2")

	repo := newTestRepo(fs)
	tagsForModule, err := repo.TagsForModule("app-two")
	check.Ok(t, err)
	check.Equals(t, 1, len(tagsForModule))
	check.Equals(t, []string{"app-two-0.0.1"}, tagsForModule)
}

func TestTagsForModuleWithOverlappingNameNumberSuffix(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	createVersionTag(fs, "app-2-0.0.1")
	createVersionTag(fs, "app-0.0.2")

	repo := newTestRepo(fs)
	tagsForModule, err := repo.TagsForModule("app-2")
	check.Ok(t, err)
	check.Equals(t, 1, len(tagsForModule))
	check.Equals(t, []string{"app-2-0.0.1"}, tagsForModule)
}

func TestTagsForModuleWithOverlappingName(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	createVersionTag(fs, "my-app-0.0.2")
	createVersionTag(fs, "app-0.0.1")

	repo := newTestRepo(fs)
	tagsForModule, err := repo.TagsForModule("app")
	check.Ok(t, err)
	check.Equals(t, 1, len(tagsForModule))
	check.Equals(t, []string{"app-0.0.1"}, tagsForModule)
}

func TestTagsForModule_emptyStage(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	createVersionTag(fs, "app-0.0.1")

	repo := newTestRepo(fs)
	tagsForModule, err := repo.TagsForModule("app", "")
	check.Ok(t, err)
	check.Equals(t, 1, len(tagsForModule))
	check.Equals(t, []string{"app-0.0.1"}, tagsForModule)
}

func TestTagsForModuleForMultipleCommitsAndTags(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	createVersionTag(fs, "app-0.0.1")
	createCommit(fs, "New Version", "Hello world 2")
	createVersionTag(fs, "app-0.0.2")

	repo := newTestRepo(fs)
	tagsForModule, err := repo.TagsForModule("app")
	check.Ok(t, err)
	check.Equals(t, 2, len(tagsForModule))
	check.Equals(t, []string{"app-0.0.1", "app-0.0.2"}, tagsForModule)
}

func TestTagsForModuleIsSorted_Patch(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	var expectedTags []string
	for i := 0; i < 101; i++ {
		is := strconv.Itoa(i)
		createCommit(fs, "New Version "+is, "Hello world "+is)
		tag := "app-0.0." + is
		createVersionTag(fs, tag)
		expectedTags = append(expectedTags, tag)
	}

	repo := newTestRepo(fs)
	tagsForModule, err := repo.TagsForModule("app")
	check.Ok(t, err)
	check.Equals(t, 101, len(tagsForModule))
	check.Equals(t, expectedTags, tagsForModule)
}

func TestTagsForModuleIsSorted_Minor(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	var expectedTags []string
	for i := 0; i < 101; i++ {
		is := strconv.Itoa(i)
		createCommit(fs, "New Version "+is, "Hello world "+is)
		tag := "app-0." + is + ".0"
		createVersionTag(fs, tag)
		expectedTags = append(expectedTags, tag)
	}

	repo := newTestRepo(fs)
	tagsForModule, err := repo.TagsForModule("app")
	check.Ok(t, err)
	check.Equals(t, 101, len(tagsForModule))
	check.Equals(t, expectedTags, tagsForModule)
}

func TestTagsForModuleIsSorted_Major(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	var expectedTags []string
	for i := 0; i < 101; i++ {
		is := strconv.Itoa(i)
		createCommit(fs, "New Version "+is, "Hello world "+is)
		tag := "app-" + is + ".0.0"
		createVersionTag(fs, tag)
		expectedTags = append(expectedTags, tag)
	}

	repo := newTestRepo(fs)
	tagsForModule, err := repo.TagsForModule("app")
	check.Ok(t, err)
	check.Equals(t, 101, len(tagsForModule))
	check.Equals(t, expectedTags, tagsForModule)
}

func TestTagsForModuleWhenNone(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)

	repo := newTestRepo(fs)
	tagsForModule, err := repo.TagsForModule("app")
	check.Ok(t, err)
	check.Equals(t, 0, len(tagsForModule))
}

func TestTagsForModuleWhenOnlyPromoted(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	createVersionTag(fs, "app-0.0.1+promoted")

	repo := newTestRepo(fs)
	tagsForModule, err := repo.TagsForModule("app")
	check.Ok(t, err)
	check.Equals(t, 0, len(tagsForModule))
}

func TestTagsForModuleWhenPartialModuleNameMatch(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	createVersionTag(fs, "ap-0.0.1")

	repo := newTestRepo(fs)
	tagsForModule, err := repo.TagsForModule("app")
	check.Ok(t, err)
	check.Equals(t, 0, len(tagsForModule))
}

func TestTagsForModuleWhenDashSeparatedPartialModuleNameMatch(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	createVersionTag(fs, "another-app-0.0.1")

	repo := newTestRepo(fs)
	tagsForModule, err := repo.TagsForModule("app")
	check.Ok(t, err)
	check.Equals(t, 0, len(tagsForModule))
}

func TestTagsForModuleWhenIncorrectVersioning(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	createVersionTag(fs, "app-.0.1")

	repo := newTestRepo(fs)
	tagsForModule, err := repo.TagsForModule("app")
	check.Ok(t, err)
	check.Equals(t, 0, len(tagsForModule))
}

func TestTagsForModule_stage(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	createVersionTag(fs, "app-0.0.1")
	createVersionTag(fs, "app-0.0.1+promoted")

	repo := newTestRepo(fs)
	tagsForModule, err := repo.TagsForModule("app", "promoted")
	check.Ok(t, err)
	check.Equals(t, 1, len(tagsForModule))
	check.Equals(t, []string{"app-0.0.1+promoted"}, tagsForModule)
}

func TestTagsForModuleForMultipleCommitsAndTags_stage(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	createVersionTag(fs, "app-0.0.1")
	createVersionTag(fs, "app-0.0.1+promoted")
	createCommit(fs, "New Version", "Hello world 2")
	createVersionTag(fs, "app-0.0.2")
	createVersionTag(fs, "app-0.0.2+promoted")

	repo := newTestRepo(fs)
	tagsForModule, err := repo.TagsForModule("app", "promoted")
	check.Ok(t, err)
	check.Equals(t, 2, len(tagsForModule))
	check.Equals(t, []string{"app-0.0.1+promoted", "app-0.0.2+promoted"}, tagsForModule)
}

func TestTagsForModuleIsSorted_stage(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	var expectedTags []string
	for i := 0; i < 10; i++ {
		is := strconv.Itoa(i)
		createCommit(fs, "New Version "+is, "Hello world "+is)
		tag := "app-0.0." + is + "+promoted"
		createVersionTag(fs, tag)
		expectedTags = append(expectedTags, tag)
	}

	repo := newTestRepo(fs)
	tagsForModule, err := repo.TagsForModule("app", "promoted")
	check.Ok(t, err)
	check.Equals(t, 10, len(tagsForModule))
	check.Equals(t, expectedTags, tagsForModule)
}

func TestTagsForModuleWhenNone_stage(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)

	repo := newTestRepo(fs)
	tagsForModule, err := repo.TagsForModule("app", "promoted")
	check.Ok(t, err)
	check.Equals(t, 0, len(tagsForModule))
}

func TestTagsForModuleOnlyUnpromoted_stage(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	createVersionTag(fs, "app-0.0.1")

	repo := newTestRepo(fs)
	tagsForModule, err := repo.TagsForModule("app", "promoted")
	check.Ok(t, err)
	check.Equals(t, 0, len(tagsForModule))
}

func TestTagsForModulePartialModuleNameMatch_stage(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	createVersionTag(fs, "ap-0.0.1+promoted")

	repo := newTestRepo(fs)
	tagsForModule, err := repo.TagsForModule("app", "promoted")
	check.Ok(t, err)
	check.Equals(t, 0, len(tagsForModule))
}

func TestTagsForModuleIncorrectVersioning_stage(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	createVersionTag(fs, "app-.0.1+promoted")

	repo := newTestRepo(fs)
	tagsForModule, err := repo.TagsForModule("app", "promoted")
	check.Ok(t, err)
	check.Equals(t, 0, len(tagsForModule))
}
