//+build e2e

package e2e

import (
	"cdx/internal/check"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestReleaseOpensRepository(t *testing.T) {
	createTempGitDir()
	command := exec.Command("cdx", "tag", "release", "-n", "app")
	err := command.Run()
	check.Ok(t, err)
}

func TestReleaseTagsRepository(t *testing.T) {
	createTempGitDir()
	command := exec.Command("cdx", "tag", "release", "-n", "app")
	err := command.Run()
	check.Ok(t, err)
	output, err := exec.Command("git", "tag", "--points-at", "HEAD").CombinedOutput()
	check.Ok(t, err)
	check.Equals(t, "app-0.1.0", strings.TrimSpace(string(output)))
}

func TestReleaseTagsRepository_patch(t *testing.T) {
	createTempGitDir()
	command := exec.Command("cdx", "tag", "release", "-n", "app", "-i", "patch")
	err := command.Run()
	check.Ok(t, err)
	output, err := exec.Command("git", "tag", "--points-at", "HEAD").CombinedOutput()
	check.Ok(t, err)
	check.Equals(t, "app-0.0.1", strings.TrimSpace(string(output)))
}

func TestReleaseTagsRepository_minor(t *testing.T) {
	createTempGitDir()
	command := exec.Command("cdx", "tag", "release", "-n", "app", "-i", "minor")
	err := command.Run()
	check.Ok(t, err)
	output, err := exec.Command("git", "tag", "--points-at", "HEAD").CombinedOutput()
	check.Ok(t, err)
	check.Equals(t, "app-0.1.0", strings.TrimSpace(string(output)))
}

func TestReleaseTagsRepository_major(t *testing.T) {
	createTempGitDir()
	command := exec.Command("cdx", "tag", "release", "-n", "app", "-i", "major")
	err := command.Run()
	check.Ok(t, err)
	output, err := exec.Command("git", "tag", "--points-at", "HEAD").CombinedOutput()
	check.Ok(t, err)
	check.Equals(t, "app-1.0.0", strings.TrimSpace(string(output)))
}

func TestReleaseTagsRepositoryAndPushesTags(t *testing.T) {
	fn := createTempGitDir()
	rd := createTempGitRemote(fn)

	command := exec.Command("cdx", "tag", "release", "-n", "app", "--push")
	err := command.Run()
	check.Ok(t, err)
	output, err := exec.Command("git", "tag", "--points-at", "HEAD").CombinedOutput()
	check.Equals(t, "app-0.1.0", strings.TrimSpace(string(output)))
	check.Ok(t, err)
	_ = os.Chdir(rd)
	output, err = exec.Command("git", "tag", "--points-at", "HEAD").CombinedOutput()
	check.Ok(t, err)
	check.Equals(t, "app-0.1.0", strings.TrimSpace(string(output)))
}

func TestReleaseTagsRepositoryAndPushesAllTags(t *testing.T) {
	fn := createTempGitDir()
	createTag(fn, "app-0.1.0")
	createCommit(fn, "commit 2")
	rd := createTempGitRemote(fn)

	command := exec.Command("cdx", "tag", "release", "-n", "app", "--push")
	err := command.Run()
	check.Ok(t, err)
	output, err := exec.Command("git", "tag", "--points-at", "HEAD").CombinedOutput()
	check.Equals(t, "app-0.2.0", strings.TrimSpace(string(output)))
	check.Ok(t, err)
	_ = os.Chdir(rd)
	output, err = exec.Command("git", "tag", "--points-at", "HEAD~1").CombinedOutput()
	check.Ok(t, err)
	check.Equals(t, "app-0.1.0", strings.TrimSpace(string(output)))
	output, err = exec.Command("git", "tag", "--points-at", "HEAD").CombinedOutput()
	check.Ok(t, err)
	check.Equals(t, "app-0.2.0", strings.TrimSpace(string(output)))
}

func TestReleaseTagsRepositoryAndPushesTags_detachedHead(t *testing.T) {
	fn := createTempGitDir()
	rd := createTempGitRemote(fn)
	_ = exec.Command("git", "checkout", "master", "--detach").Run()

	command := exec.Command("cdx", "tag", "release", "-n", "app", "--push")
	err := command.Run()
	check.Ok(t, err)
	output, err := exec.Command("git", "tag", "--points-at", "HEAD").CombinedOutput()
	check.Equals(t, "app-0.1.0", strings.TrimSpace(string(output)))
	check.Ok(t, err)
	_ = os.Chdir(rd)
	output, err = exec.Command("git", "tag", "--points-at", "HEAD").CombinedOutput()
	check.Ok(t, err)
	check.Equals(t, "app-0.1.0", strings.TrimSpace(string(output)))
}

func TestReleaseTagsRepositoryAndPushesTags_detachedHeadAndNonHeadTags(t *testing.T) {
	fn := createTempGitDir()
	rd := createTempGitRemote(fn)
	createTag(fn, "app-0.1.0")
	createCommit(fn, "commit 2")
	createCommit(fn, "commit 3")
	_ = exec.Command("git", "push", "origin", "master").Run()
	_ = exec.Command("git", "checkout", "HEAD~1", "--detach").Run()

	command := exec.Command("cdx", "tag", "release", "-n", "app", "--push")
	err := command.Run()
	check.Ok(t, err)
	_ = exec.Command("git", "checkout", "master").Run()
	output, err := exec.Command("git", "tag", "--points-at", "HEAD~1").CombinedOutput()
	check.Equals(t, "app-0.2.0", strings.TrimSpace(string(output)))
	check.Ok(t, err)
	_ = os.Chdir(rd)
	output, err = exec.Command("git", "tag", "--points-at", "HEAD~1").CombinedOutput()
	check.Ok(t, err)
	check.Equals(t, "app-0.2.0", strings.TrimSpace(string(output)))
}
