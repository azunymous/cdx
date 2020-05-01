//+build e2e

package tag

import (
	"bytes"
	"github.com/azunymous/cdx/test/check"
	"os/exec"
	"testing"
)

func TestLatestOpensRepositoryWithNoTagsFails(t *testing.T) {
	createTempGitDir()
	command := exec.Command("cdx", "tag", "latest", "-n", "app")
	err := command.Run()
	check.Assert(t, err != nil, "expecting error to not be nil, got %v", err)
}

func TestLatestGetsTagsFromRepository(t *testing.T) {
	dir := createTempGitDir()
	createTag(dir, "app-0.1.0")

	command := exec.Command("cdx", "tag", "latest", "-n", "app")
	var stdOut bytes.Buffer
	var stdErr bytes.Buffer
	command.Stdout = &stdOut
	command.Stderr = &stdErr
	err := command.Run()
	check.Ok(t, err)
	check.Equals(t, "0.1.0\n", stdOut.String())
	check.Equals(t, "", stdErr.String())
}

func TestLatestGetsTagsFromRepositoryForMultipleCommits(t *testing.T) {
	dir := createTempGitDir()
	createTag(dir, "app-0.1.0")
	createCommit(dir, "commit 2")
	createTag(dir, "app-0.2.0")
	createCommit(dir, "commit 3")

	command := exec.Command("cdx", "tag", "latest", "-n", "app")
	var stdOut bytes.Buffer
	var stdErr bytes.Buffer
	command.Stdout = &stdOut
	command.Stderr = &stdErr
	err := command.Run()
	check.Ok(t, err)
	check.Equals(t, "0.2.0\n", stdOut.String())
	check.Equals(t, "", stdErr.String())
}

func TestLatestGetsTagsFromRepositoryOnlyOnHead(t *testing.T) {
	dir := createTempGitDir()
	createTag(dir, "app-0.1.0")
	createCommit(dir, "commit 2")
	createTag(dir, "app-0.2.0")
	_ = exec.Command("git", "checkout", "HEAD~1", "--detach").Run()

	command := exec.Command("cdx", "tag", "latest", "-n", "app", "--head")
	var stdOut bytes.Buffer
	var stdErr bytes.Buffer
	command.Stdout = &stdOut
	command.Stderr = &stdErr
	err := command.Run()
	check.Equals(t, "", stdErr.String())
	check.Ok(t, err)
	check.Equals(t, "0.1.0\n", stdOut.String())
}

func TestLatestOpensRepositoryForStageWithNoTagsFails(t *testing.T) {
	createTempGitDir()
	command := exec.Command("cdx", "tag", "latest", "-n", "app", "promoted")
	err := command.Run()
	check.Assert(t, err != nil, "expecting error to not be nil, got %v", err)
}

func TestLatestOpensRepositoryForStageWithNoStageTagsFails(t *testing.T) {
	dir := createTempGitDir()
	createTag(dir, "app-0.1.0")

	command := exec.Command("cdx", "tag", "latest", "-n", "app", "promoted")
	err := command.Run()
	check.Assert(t, err != nil, "expecting error to not be nil, got %v", err)
}

func TestLatestGetsTagsFromRepositoryForStage(t *testing.T) {
	dir := createTempGitDir()
	createTag(dir, "app-0.1.0")
	createTag(dir, "app-0.1.0-promoted")

	command := exec.Command("cdx", "tag", "latest", "-n", "app", "promoted")
	var stdOut bytes.Buffer
	var stdErr bytes.Buffer
	command.Stdout = &stdOut
	command.Stderr = &stdErr
	err := command.Run()
	check.Equals(t, "", stdErr.String())
	check.Ok(t, err)
	check.Equals(t, "0.1.0\n", stdOut.String())
}

func TestLatestGetsTagsFromRepositoryForStageWhenHeadIsNotTagged(t *testing.T) {
	dir := createTempGitDir()
	createTag(dir, "app-0.1.0")
	createTag(dir, "app-0.1.0-promoted")
	createCommit(dir, "commit 2")

	command := exec.Command("cdx", "tag", "latest", "-n", "app", "promoted")
	var stdOut bytes.Buffer
	var stdErr bytes.Buffer
	command.Stdout = &stdOut
	command.Stderr = &stdErr
	err := command.Run()
	check.Equals(t, "", stdErr.String())
	check.Ok(t, err)
	check.Equals(t, "0.1.0\n", stdOut.String())
}

func TestLatestGetsTagsFromRepositoryForStageForMultipleCommits(t *testing.T) {
	dir := createTempGitDir()
	createTag(dir, "app-0.1.0")
	createTag(dir, "app-0.1.0-promoted")
	createCommit(dir, "commit 2")
	createTag(dir, "app-0.2.0")
	createTag(dir, "app-0.2.0-promoted")
	createCommit(dir, "commit 3")

	command := exec.Command("cdx", "tag", "latest", "-n", "app", "promoted")
	var stdOut bytes.Buffer
	var stdErr bytes.Buffer
	command.Stdout = &stdOut
	command.Stderr = &stdErr
	err := command.Run()
	check.Equals(t, "", stdErr.String())
	check.Ok(t, err)
	check.Equals(t, "0.2.0\n", stdOut.String())
}

func TestLatestGetsTagsFromRepositoryForStageForMultipleCommits_onlyPromoted(t *testing.T) {
	dir := createTempGitDir()
	createTag(dir, "app-0.1.0")
	createTag(dir, "app-0.1.0-promoted")
	createCommit(dir, "commit 2")
	createTag(dir, "app-0.2.0")
	createTag(dir, "app-0.2.0")
	createCommit(dir, "commit 3")

	command := exec.Command("cdx", "tag", "latest", "-n", "app", "promoted")
	var stdOut bytes.Buffer
	var stdErr bytes.Buffer
	command.Stdout = &stdOut
	command.Stderr = &stdErr
	err := command.Run()
	check.Equals(t, "", stdErr.String())
	check.Ok(t, err)
	check.Equals(t, "0.1.0\n", stdOut.String())
}

func TestLatestGetsTagsFromRepositoryForStageOnlyOnHead(t *testing.T) {
	dir := createTempGitDir()
	createTag(dir, "app-0.1.0")
	createTag(dir, "app-0.1.0-promoted")
	createCommit(dir, "commit 2")
	createTag(dir, "app-0.2.0")
	createTag(dir, "app-0.2.0-promoted")
	_ = exec.Command("git", "checkout", "HEAD~1", "--detach").Run()

	command := exec.Command("cdx", "tag", "latest", "-n", "app", "promoted", "--head")
	var stdOut bytes.Buffer
	var stdErr bytes.Buffer
	command.Stdout = &stdOut
	command.Stderr = &stdErr
	err := command.Run()
	check.Equals(t, "", stdErr.String())
	check.Ok(t, err)
	check.Equals(t, "0.1.0\n", stdOut.String())
}
