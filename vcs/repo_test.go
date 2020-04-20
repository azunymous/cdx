package vcs

import (
	"cdx/internal/check"
	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/cache"
	"github.com/go-git/go-git/v5/plumbing/object"
	filesystem2 "github.com/go-git/go-git/v5/storage/filesystem"
	"strconv"
	"testing"
	"time"
)

func TestNewRepo(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)

	repo, err := NewRepo(fs)
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

	repo, err := NewRepo(fs)
	check.Ok(t, err)
	tagsForHead, err := repo.TagsForHead()
	check.Ok(t, err)
	check.Equals(t, 1, len(tagsForHead))
	check.Equals(t, []string{"app-0.0.1"}, tagsForHead)
}

func TestTagsForHeadForMultipleTags(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	createVersionTag(fs, "app-0.0.1-alpha")
	createVersionTag(fs, "app-0.0.1")

	repo, err := NewRepo(fs)
	check.Ok(t, err)
	tagsForHead, err := repo.TagsForHead()
	check.Ok(t, err)
	check.Equals(t, 2, len(tagsForHead))
	check.Equals(t, []string{"app-0.0.1", "app-0.0.1-alpha"}, tagsForHead)
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

	repo, err := NewRepo(fs)
	check.Ok(t, err)
	tagsForHead, err := repo.TagsForHead()
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

	repo, err := NewRepo(fs)
	check.Ok(t, err)
	tagsForHead, err := repo.TagsForHead()
	check.Ok(t, err)
	check.Equals(t, 1, len(tagsForHead))

	check.Equals(t, []string{"app-0.0.2"}, tagsForHead)
}

func TestTagsForHeadWhenNone(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)

	repo, err := NewRepo(fs)
	check.Ok(t, err)
	tagsForHead, err := repo.TagsForHead()
	check.Ok(t, err)
	check.Equals(t, 0, len(tagsForHead))
}

func TestTagsForModule(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	createVersionTag(fs, "app-0.0.1")

	repo, err := NewRepo(fs)
	check.Ok(t, err)
	tagsForModule, err := repo.TagsForModule("app")
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

	repo, err := NewRepo(fs)
	check.Ok(t, err)
	tagsForModule, err := repo.TagsForModule("app")
	check.Ok(t, err)
	check.Equals(t, 2, len(tagsForModule))
	check.Equals(t, []string{"app-0.0.1", "app-0.0.2"}, tagsForModule)
}

func TestTagsForModuleIsSorted(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	var expectedTags []string
	for i := 0; i < 10; i++ {
		is := strconv.Itoa(i)
		createCommit(fs, "New Version "+is, "Hello world "+is)
		tag := "app-0.0." + is
		createVersionTag(fs, tag)
		expectedTags = append(expectedTags, tag)
	}

	repo, err := NewRepo(fs)
	check.Ok(t, err)
	tagsForModule, err := repo.TagsForModule("app")
	check.Ok(t, err)
	check.Equals(t, 10, len(tagsForModule))
	check.Equals(t, expectedTags, tagsForModule)
}

func TestTagsForModuleWhenNone(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)

	repo, err := NewRepo(fs)
	check.Ok(t, err)
	tagsForModule, err := repo.TagsForModule("app")
	check.Ok(t, err)
	check.Equals(t, 0, len(tagsForModule))
}

func TestTagsForModuleWhenOnlyPromoted(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	createVersionTag(fs, "app-0.0.1-promoted")

	repo, err := NewRepo(fs)
	check.Ok(t, err)
	tagsForModule, err := repo.TagsForModule("app")
	check.Ok(t, err)
	check.Equals(t, 0, len(tagsForModule))
}

func TestTagsForModuleWhenPartialModuleNameMatch(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	createVersionTag(fs, "ap-0.0.1")

	repo, err := NewRepo(fs)
	check.Ok(t, err)
	tagsForModule, err := repo.TagsForModule("app")
	check.Ok(t, err)
	check.Equals(t, 0, len(tagsForModule))
}

func TestTagsForModuleWhenIncorrectVersioning(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	createVersionTag(fs, "app-.0.1")

	repo, err := NewRepo(fs)
	check.Ok(t, err)
	tagsForModule, err := repo.TagsForModule("app")
	check.Ok(t, err)
	check.Equals(t, 0, len(tagsForModule))
}

func TestTagsForModule_stage(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	createVersionTag(fs, "app-0.0.1")
	createVersionTag(fs, "app-0.0.1-promoted")

	repo, err := NewRepo(fs)
	check.Ok(t, err)
	tagsForModule, err := repo.TagsForModule("app", "promoted")
	check.Ok(t, err)
	check.Equals(t, 1, len(tagsForModule))
	check.Equals(t, []string{"app-0.0.1-promoted"}, tagsForModule)
}

func TestTagsForModuleForMultipleCommitsAndTags_stage(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	createVersionTag(fs, "app-0.0.1")
	createVersionTag(fs, "app-0.0.1-promoted")
	createCommit(fs, "New Version", "Hello world 2")
	createVersionTag(fs, "app-0.0.2")
	createVersionTag(fs, "app-0.0.2-promoted")

	repo, err := NewRepo(fs)
	check.Ok(t, err)
	tagsForModule, err := repo.TagsForModule("app", "promoted")
	check.Ok(t, err)
	check.Equals(t, 2, len(tagsForModule))
	check.Equals(t, []string{"app-0.0.1-promoted", "app-0.0.2-promoted"}, tagsForModule)
}

func TestTagsForModuleIsSorted_stage(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	var expectedTags []string
	for i := 0; i < 10; i++ {
		is := strconv.Itoa(i)
		createCommit(fs, "New Version "+is, "Hello world "+is)
		tag := "app-0.0." + is + "-promoted"
		createVersionTag(fs, tag)
		expectedTags = append(expectedTags, tag)
	}

	repo, err := NewRepo(fs)
	check.Ok(t, err)
	tagsForModule, err := repo.TagsForModule("app", "promoted")
	check.Ok(t, err)
	check.Equals(t, 10, len(tagsForModule))
	check.Equals(t, expectedTags, tagsForModule)
}

func TestTagsForModuleWhenNone_stage(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)

	repo, err := NewRepo(fs)
	check.Ok(t, err)
	tagsForModule, err := repo.TagsForModule("app", "promoted")
	check.Ok(t, err)
	check.Equals(t, 0, len(tagsForModule))
}

func TestTagsForModuleOnlyUnpromoted_stage(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	createVersionTag(fs, "app-0.0.1")

	repo, err := NewRepo(fs)
	check.Ok(t, err)
	tagsForModule, err := repo.TagsForModule("app", "promoted")
	check.Ok(t, err)
	check.Equals(t, 0, len(tagsForModule))
}

func TestTagsForModulePartialModuleNameMatch_stage(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	createVersionTag(fs, "ap-0.0.1-promoted")

	repo, err := NewRepo(fs)
	check.Ok(t, err)
	tagsForModule, err := repo.TagsForModule("app", "promoted")
	check.Ok(t, err)
	check.Equals(t, 0, len(tagsForModule))
}

func TestTagsForModuleIncorrectVersioning_stage(t *testing.T) {
	fs := memfs.New()
	createGitRepo(fs)
	createVersionTag(fs, "app-.0.1-promoted")

	repo, err := NewRepo(fs)
	check.Ok(t, err)
	tagsForModule, err := repo.TagsForModule("app", "promoted")
	check.Ok(t, err)
	check.Equals(t, 0, len(tagsForModule))
}

func createGitRepo(fs billy.Filesystem) {
	_, _ = git.Init(filesystem2.NewStorage(fs, cache.NewObjectLRUDefault()), fs)
	createCommit(fs, "example-git-file", "hello world!")
}

func createCommit(fs billy.Filesystem, filename, msg string) {
	r, _ := git.Open(filesystem2.NewStorage(fs, cache.NewObjectLRUDefault()), fs)
	w, _ := r.Worktree()
	// create a new file inside of the worktree of the project using the go standard library.
	file, _ := fs.Create(filename)
	_, _ = file.Write([]byte(msg))

	// Adds the new file to the staging area.
	_, _ = w.Add(filename)

	// Commits the current staging area to the repository, with the new file
	// just created.
	_, _ = w.Commit("example go-git commit", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "John Doe",
			Email: "john@doe.org",
			When:  time.Now(),
		},
	})
}

func createVersionTag(fs billy.Filesystem, tag string) {
	r, _ := git.Open(filesystem2.NewStorage(fs, cache.NewObjectLRUDefault()), fs)
	revision, _ := r.ResolveRevision("HEAD")
	_, _ = r.CreateTag(tag, *revision, nil)
}
