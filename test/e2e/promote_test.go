package e2e

import (
	"cdx/internal/check"
	"os"
	"os/exec"
	"sort"
	"strings"
	"testing"
)

func TestPromoteOpensRepository(t *testing.T) {
	dir := createTempGitDir()
	createTag(dir, "app-0.1.0")
	command := exec.Command("cdx", "tag", "promote", "-n", "app", "promoted")
	err := command.Run()
	check.Ok(t, err)
}

func TestPromoteTagsRepository(t *testing.T) {
	dir := createTempGitDir()
	createTag(dir, "app-0.1.0")
	command := exec.Command("cdx", "tag", "promote", "-n", "app", "promoted")
	err := command.Run()
	check.Ok(t, err)
	output, err := exec.Command("git", "tag", "--points-at", "HEAD").CombinedOutput()
	check.Ok(t, err)
	tags := strings.Split(string(output), "\n")
	sort.Strings(tags)
	check.Equals(t, "app-0.1.0-promoted", tags[len(tags)-1])
}

func TestPromoteTagsRepositoryAlreadyPromoted(t *testing.T) {
	dir := createTempGitDir()
	createTag(dir, "app-0.1.0")
	createTag(dir, "app-0.1.0-promoted")
	command := exec.Command("cdx", "tag", "promote", "-n", "app", "promoted")
	err := command.Run()
	check.Ok(t, err)
	output, err := exec.Command("git", "tag", "--points-at", "HEAD").CombinedOutput()
	check.Ok(t, err)
	tags := strings.Split(string(output), "\n")
	sort.Strings(tags)
	check.Equals(t, "app-0.1.0-promoted", tags[len(tags)-1])
}

func TestPromoteTagsRepositoryAndPushesTags(t *testing.T) {
	dir := createTempGitDir()
	createTag(dir, "app-0.1.0")
	rd := createTempGitRemote(dir)

	command := exec.Command("cdx", "tag", "promote", "-n", "app", "promoted", "--push")
	err := command.Run()
	check.Ok(t, err)
	output, err := exec.Command("git", "tag", "--points-at", "HEAD").CombinedOutput()
	check.Ok(t, err)
	tags := strings.Split(string(output), "\n")
	sort.Strings(tags)
	check.Equals(t, "app-0.1.0-promoted", tags[len(tags)-1])

	_ = os.Chdir(rd)
	output, err = exec.Command("git", "tag", "--points-at", "HEAD").CombinedOutput()
	check.Ok(t, err)
	tags = strings.Split(string(output), "\n")
	sort.Strings(tags)
	check.Equals(t, "app-0.1.0-promoted", tags[len(tags)-1])
}

func TestPromoteTagsRepositoryAndPushesTags_detachedHead(t *testing.T) {
	dir := createTempGitDir()
	createTag(dir, "app-0.1.0")
	rd := createTempGitRemote(dir)
	_ = exec.Command("git", "checkout", "master", "--detach").Run()

	command := exec.Command("cdx", "tag", "promote", "-n", "app", "promoted", "--push")
	err := command.Run()
	check.Ok(t, err)
	output, err := exec.Command("git", "tag", "--points-at", "HEAD").CombinedOutput()
	check.Ok(t, err)
	tags := strings.Split(string(output), "\n")
	sort.Strings(tags)
	check.Equals(t, "app-0.1.0-promoted", tags[len(tags)-1])

	_ = os.Chdir(rd)
	output, err = exec.Command("git", "tag", "--points-at", "HEAD").CombinedOutput()
	check.Ok(t, err)
	tags = strings.Split(string(output), "\n")
	sort.Strings(tags)
	check.Equals(t, "app-0.1.0-promoted", tags[len(tags)-1])
}

func TestPromoteTagsRepositoryAndPushesTags_detachedHeadAndNonHeadTags(t *testing.T) {
	fn := createTempGitDir()
	rd := createTempGitRemote(fn)
	createTag(fn, "app-0.1.0")
	createCommit(fn, "commit 2")
	createTag(fn, "app-0.2.0")
	createCommit(fn, "commit 3")
	_ = exec.Command("git", "push", "origin", "master").Run()
	_ = exec.Command("git", "checkout", "HEAD~1", "--detach").Run()

	command := exec.Command("cdx", "tag", "promote", "-n", "app", "promoted", "--push")
	err := command.Run()
	check.Ok(t, err)
	_ = exec.Command("git", "checkout", "master").Run()

	output, err := exec.Command("git", "tag", "--points-at", "HEAD~1").CombinedOutput()
	check.Ok(t, err)
	tags := strings.Split(string(output), "\n")
	sort.Strings(tags)
	check.Equals(t, "app-0.2.0-promoted", tags[len(tags)-1])

	_ = os.Chdir(rd)
	output, err = exec.Command("git", "tag", "--points-at", "HEAD~1").CombinedOutput()
	check.Ok(t, err)
	tags = strings.Split(string(output), "\n")
	sort.Strings(tags)
	check.Equals(t, "app-0.2.0-promoted", tags[len(tags)-1])
}