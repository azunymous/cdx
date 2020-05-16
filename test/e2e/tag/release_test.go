//+build e2e

package tag

import (
	"github.com/azunymous/cdx/test/check"
	"github.com/azunymous/cdx/test/e2e"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestReleaseOpensRepository(t *testing.T) {
	e2e.CreateTempGitDir()
	command := exec.Command(e2e.CDX, "tag", "release", "-n", "app")
	err := command.Run()
	check.Ok(t, err)
}

func TestReleaseTagsRepository(t *testing.T) {
	dir := e2e.CreateTempGitDir()
	e2e.CreateTag(dir, "app-0.1.0")
	e2e.CreateCommit(dir, "Commit 2")
	command := exec.Command(e2e.CDX, "tag", "release", "-n", "app")
	err := command.Run()
	check.Ok(t, err)
	output, err := exec.Command("git", "tag", "--points-at", "HEAD").CombinedOutput()
	check.Ok(t, err)
	check.Equals(t, "app-0.2.0", strings.TrimSpace(string(output)))
}

func TestReleaseTagsRepositoryAlreadyReleased(t *testing.T) {
	dir := e2e.CreateTempGitDir()
	e2e.CreateTag(dir, "app-0.1.0")
	command := exec.Command(e2e.CDX, "tag", "release", "-n", "app")
	err := command.Run()
	check.Ok(t, err)
	output, err := exec.Command("git", "tag", "--points-at", "HEAD").CombinedOutput()
	check.Ok(t, err)
	check.Equals(t, "app-0.1.0", strings.TrimSpace(string(output)))
}

func TestReleaseTagsRepositoryFirstRelease(t *testing.T) {
	e2e.CreateTempGitDir()
	command := exec.Command(e2e.CDX, "tag", "release", "-n", "app")
	err := command.Run()
	check.Ok(t, err)
	output, err := exec.Command("git", "tag", "--points-at", "HEAD").CombinedOutput()
	check.Ok(t, err)
	check.Equals(t, "app-0.1.0", strings.TrimSpace(string(output)))
}

func TestReleaseTagsRepository_patch(t *testing.T) {
	e2e.CreateTempGitDir()
	command := exec.Command(e2e.CDX, "tag", "release", "-n", "app", "-i", "patch")
	err := command.Run()
	check.Ok(t, err)
	output, err := exec.Command("git", "tag", "--points-at", "HEAD").CombinedOutput()
	check.Ok(t, err)
	check.Equals(t, "app-0.0.1", strings.TrimSpace(string(output)))
}

func TestReleaseTagsRepository_minor(t *testing.T) {
	e2e.CreateTempGitDir()
	command := exec.Command(e2e.CDX, "tag", "release", "-n", "app", "-i", "minor")
	err := command.Run()
	check.Ok(t, err)
	output, err := exec.Command("git", "tag", "--points-at", "HEAD").CombinedOutput()
	check.Ok(t, err)
	check.Equals(t, "app-0.1.0", strings.TrimSpace(string(output)))
}

func TestReleaseTagsRepository_major(t *testing.T) {
	e2e.CreateTempGitDir()
	command := exec.Command(e2e.CDX, "tag", "release", "-n", "app", "-i", "major")
	err := command.Run()
	check.Ok(t, err)
	output, err := exec.Command("git", "tag", "--points-at", "HEAD").CombinedOutput()
	check.Ok(t, err)
	check.Equals(t, "app-1.0.0", strings.TrimSpace(string(output)))
}

func TestReleaseTagsRepositoryAndPushesTags(t *testing.T) {
	fn := e2e.CreateTempGitDir()
	rd := e2e.CreateTempGitRemote(fn)

	command := exec.Command(e2e.CDX, "tag", "release", "-n", "app", "--push")
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

func TestReleaseDoesNotTagNonOriginMasterWithPushFlag(t *testing.T) {
	dir := e2e.CreateTempGitDir()
	rd := e2e.CreateTempGitRemote(dir)
	e2e.CreateTag(dir, "app-0.1.0")
	_ = exec.Command("git", "checkout", "-b", "test-branch").Run()

	e2e.CreateCommit(dir, "Commit 2")
	command := exec.Command(e2e.CDX, "tag", "release", "-n", "app", "--push")
	err := command.Run()
	check.Ok(t, err)
	output, err := exec.Command("git", "tag", "--points-at", "HEAD").CombinedOutput()
	check.Ok(t, err)
	check.Equals(t, "", strings.TrimSpace(string(output)))

	_ = os.Chdir(rd)
	output, err = exec.Command("git", "tag", "--points-at", "HEAD").CombinedOutput()
	check.Ok(t, err)
	check.Equals(t, "", strings.TrimSpace(string(output)))
}

func TestReleaseTagsDoesNotFailIfAlreadyTaggedLocallyWithPushFlag(t *testing.T) {
	dir := e2e.CreateTempGitDir()
	rd := e2e.CreateTempGitRemote(dir)
	e2e.CreateTag(dir, "app-0.1.0")
	e2e.CreateCommit(dir, "Commit 2")
	_ = exec.Command("git", "checkout", "HEAD~1", "--detach").Run()

	command := exec.Command(e2e.CDX, "tag", "release", "-n", "app", "--push")
	err := command.Run()
	check.Ok(t, err)
	output, err := exec.Command("git", "tag", "--points-at", "HEAD").CombinedOutput()
	check.Ok(t, err)
	check.Equals(t, "app-0.1.0", strings.TrimSpace(string(output)))

	_ = os.Chdir(rd)
	output, err = exec.Command("git", "tag", "--points-at", "HEAD").CombinedOutput()
	check.Ok(t, err)
	check.Equals(t, "app-0.1.0", strings.TrimSpace(string(output)))
}

func TestReleaseTagsRepositoryDoesNotFailifAlreadyTaggedRemotelyWithPushFlag(t *testing.T) {
	dir := e2e.CreateTempGitDir()
	e2e.CreateTag(dir, "app-0.1.0")
	rd := e2e.CreateTempGitRemote(dir)
	_, _ = exec.Command("git", "push", "--tags").CombinedOutput()
	_ = exec.Command("git", "checkout", "HEAD", "--detach").Run()

	command := exec.Command(e2e.CDX, "tag", "release", "-n", "app", "--push")
	err := command.Run()
	check.Ok(t, err)
	output, err := exec.Command("git", "tag", "--points-at", "HEAD").CombinedOutput()
	check.Ok(t, err)
	check.Equals(t, "app-0.1.0", strings.TrimSpace(string(output)))

	_ = os.Chdir(rd)
	output, err = exec.Command("git", "tag", "--points-at", "HEAD").CombinedOutput()
	check.Ok(t, err)
	check.Equals(t, "app-0.1.0", strings.TrimSpace(string(output)))
}

func TestReleaseWithPushFlagFailsWhenNoRemote(t *testing.T) {
	dir := e2e.CreateTempGitDir()
	e2e.CreateTag(dir, "app-0.1.0")
	e2e.CreateCommit(dir, "Commit 2")

	command := exec.Command(e2e.CDX, "tag", "release", "-n", "app", "--push")
	err := command.Run()
	check.Assert(t, err != nil, "expecting error to not be nil, got %v", err)
	output, err := exec.Command("git", "tag", "--points-at", "HEAD").CombinedOutput()
	check.Ok(t, err)
	check.Equals(t, "", strings.TrimSpace(string(output)))
}

func TestReleaseTagsRepositoryAndPushesAllTags(t *testing.T) {
	fn := e2e.CreateTempGitDir()
	e2e.CreateTag(fn, "app-0.1.0")
	e2e.CreateCommit(fn, "commit 2")
	rd := e2e.CreateTempGitRemote(fn)

	command := exec.Command(e2e.CDX, "tag", "release", "-n", "app", "--push")
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
	fn := e2e.CreateTempGitDir()
	rd := e2e.CreateTempGitRemote(fn)
	_ = exec.Command("git", "checkout", "master", "--detach").Run()

	command := exec.Command(e2e.CDX, "tag", "release", "-n", "app", "--push")
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
	fn := e2e.CreateTempGitDir()
	rd := e2e.CreateTempGitRemote(fn)
	e2e.CreateTag(fn, "app-0.1.0")
	e2e.CreateCommit(fn, "commit 2")
	e2e.CreateCommit(fn, "commit 3")
	_ = exec.Command("git", "push", "origin", "master").Run()
	_ = exec.Command("git", "checkout", "HEAD~1", "--detach").Run()

	command := exec.Command(e2e.CDX, "tag", "release", "-n", "app", "--push")
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

func TestReleaseTagsRepositoryAndPushesTags_detachedHead_FailsWhenNotOnOriginMaster(t *testing.T) {
	fn := e2e.CreateTempGitDir()
	rd := e2e.CreateTempGitRemote(fn)
	e2e.CreateTag(fn, "app-0.1.0")
	e2e.CreateCommit(fn, "commit 2")
	_ = exec.Command("git", "push", "origin", "master").Run()

	e2e.CreateCommit(fn, "commit 3")
	_ = exec.Command("git", "checkout", "HEAD", "--detach").Run()

	command := exec.Command(e2e.CDX, "tag", "release", "-n", "app", "--push")
	err := command.Run()
	check.Ok(t, err)
	_ = exec.Command("git", "checkout", "master").Run()
	output, err := exec.Command("git", "tag", "--points-at", "HEAD").CombinedOutput()
	check.Equals(t, "", strings.TrimSpace(string(output)))
	check.Ok(t, err)
	_ = os.Chdir(rd)
	output, err = exec.Command("git", "tag", "--points-at", "HEAD").CombinedOutput()
	check.Ok(t, err)
	check.Equals(t, "", strings.TrimSpace(string(output)))
}
