//+build e2e

package tag

import (
	"bytes"
	"github.com/azunymous/cdx/test/check"
	"github.com/azunymous/cdx/test/e2e"
	"os/exec"
	"testing"
)

func TestLatestOpensRepositoryWithNoTagsFails(t *testing.T) {
	e2e.CreateTempGitDir()
	command := exec.Command(e2e.CDX, "tag", "latest", "-n", "app")
	err := command.Run()
	check.Assert(t, err != nil, "expecting error to not be nil, got %v", err)
}

func TestLatestGetsTagsFromRepository(t *testing.T) {
	dir := e2e.CreateTempGitDir()
	e2e.CreateTag(dir, "app-0.1.0")

	command := exec.Command(e2e.CDX, "tag", "latest", "-n", "app")
	var stdOut bytes.Buffer
	var stdErr bytes.Buffer
	command.Stdout = &stdOut
	command.Stderr = &stdErr
	err := command.Run()
	check.Ok(t, err)
	check.Equals(t, "0.1.0\n", stdOut.String())
	check.Equals(t, "", stdErr.String())
}

func TestLatestGetsAnnotatedTagsFromRepository(t *testing.T) {
	_ = e2e.CreateTempGitDir()
	_, _ = exec.Command("git", "tag", "app-0.1.0", "-a", "-m", "").CombinedOutput()

	command := exec.Command(e2e.CDX, "tag", "latest", "-n", "app")
	var stdOut bytes.Buffer
	var stdErr bytes.Buffer
	command.Stdout = &stdOut
	command.Stderr = &stdErr
	err := command.Run()
	check.Equals(t, "", stdErr.String())
	check.Ok(t, err)
	check.Equals(t, "0.1.0\n", stdOut.String())
}

func TestLatestGetsHashFromRepositoryWhenNoTagsAndHashFlagProvided(t *testing.T) {
	_ = e2e.CreateTempGitDir()

	command := exec.Command(e2e.CDX, "tag", "latest", "-n", "app", "--head", "--fallback")
	var stdOut bytes.Buffer
	var stdErr bytes.Buffer
	command.Stdout = &stdOut
	command.Stderr = &stdErr
	err := command.Run()
	check.Ok(t, err)
	expectedOutput, _ := exec.Command("git", "rev-parse", "HEAD").Output()
	check.Equals(t, "", stdErr.String())
	check.Equals(t, string(expectedOutput), stdOut.String())
}

func TestLatestGetsTagsFromRepositoryForMultipleCommits(t *testing.T) {
	dir := e2e.CreateTempGitDir()
	e2e.CreateTag(dir, "app-0.1.0")
	e2e.CreateCommit(dir, "commit 2")
	e2e.CreateTag(dir, "app-0.2.0")
	e2e.CreateCommit(dir, "commit 3")

	command := exec.Command(e2e.CDX, "tag", "latest", "-n", "app")
	var stdOut bytes.Buffer
	var stdErr bytes.Buffer
	command.Stdout = &stdOut
	command.Stderr = &stdErr
	err := command.Run()
	check.Ok(t, err)
	check.Equals(t, "0.2.0\n", stdOut.String())
	check.Equals(t, "", stdErr.String())
}

func TestLatestGetsTagsFromRepositoryForMultipleCommits_UnnaturalSorting(t *testing.T) {
	dir := e2e.CreateTempGitDir()
	e2e.CreateTag(dir, "app-0.99.0")
	e2e.CreateCommit(dir, "commit 2")
	e2e.CreateTag(dir, "app-0.100.0")
	e2e.CreateCommit(dir, "commit 3")

	command := exec.Command(e2e.CDX, "tag", "latest", "-n", "app")
	var stdOut bytes.Buffer
	var stdErr bytes.Buffer
	command.Stdout = &stdOut
	command.Stderr = &stdErr
	err := command.Run()
	check.Ok(t, err)
	check.Equals(t, "0.100.0\n", stdOut.String())
	check.Equals(t, "", stdErr.String())
}

func TestLatestGetsTagsFromRepositoryOnlyOnHead(t *testing.T) {
	dir := e2e.CreateTempGitDir()
	e2e.CreateTag(dir, "app-0.1.0")
	e2e.CreateCommit(dir, "commit 2")
	e2e.CreateTag(dir, "app-0.2.0")
	_ = exec.Command("git", "checkout", "HEAD~1", "--detach").Run()

	command := exec.Command(e2e.CDX, "tag", "latest", "-n", "app", "--head")
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
	e2e.CreateTempGitDir()
	command := exec.Command(e2e.CDX, "tag", "latest", "-n", "app", "promoted")
	err := command.Run()
	check.Assert(t, err != nil, "expecting error to not be nil, got %v", err)
}

func TestLatestOpensRepositoryForStageWithNoStageTagsFails(t *testing.T) {
	dir := e2e.CreateTempGitDir()
	e2e.CreateTag(dir, "app-0.1.0")

	command := exec.Command(e2e.CDX, "tag", "latest", "-n", "app", "promoted")
	err := command.Run()
	check.Assert(t, err != nil, "expecting error to not be nil, got %v", err)
}

func TestLatestGetsTagsFromRepositoryForStage(t *testing.T) {
	dir := e2e.CreateTempGitDir()
	e2e.CreateTag(dir, "app-0.1.0")
	e2e.CreateTag(dir, "app-0.1.0-promoted")

	command := exec.Command(e2e.CDX, "tag", "latest", "-n", "app", "promoted")
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
	dir := e2e.CreateTempGitDir()
	e2e.CreateTag(dir, "app-0.1.0")
	e2e.CreateTag(dir, "app-0.1.0-promoted")
	e2e.CreateCommit(dir, "commit 2")

	command := exec.Command(e2e.CDX, "tag", "latest", "-n", "app", "promoted")
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
	dir := e2e.CreateTempGitDir()
	e2e.CreateTag(dir, "app-0.1.0")
	e2e.CreateTag(dir, "app-0.1.0-promoted")
	e2e.CreateCommit(dir, "commit 2")
	e2e.CreateTag(dir, "app-0.2.0")
	e2e.CreateTag(dir, "app-0.2.0-promoted")
	e2e.CreateCommit(dir, "commit 3")

	command := exec.Command(e2e.CDX, "tag", "latest", "-n", "app", "promoted")
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
	dir := e2e.CreateTempGitDir()
	e2e.CreateTag(dir, "app-0.1.0")
	e2e.CreateTag(dir, "app-0.1.0-promoted")
	e2e.CreateCommit(dir, "commit 2")
	e2e.CreateTag(dir, "app-0.2.0")
	e2e.CreateTag(dir, "app-0.2.0")
	e2e.CreateCommit(dir, "commit 3")

	command := exec.Command(e2e.CDX, "tag", "latest", "-n", "app", "promoted")
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
	dir := e2e.CreateTempGitDir()
	e2e.CreateTag(dir, "app-0.1.0")
	e2e.CreateTag(dir, "app-0.1.0-promoted")
	e2e.CreateCommit(dir, "commit 2")
	e2e.CreateTag(dir, "app-0.2.0")
	e2e.CreateTag(dir, "app-0.2.0-promoted")
	_ = exec.Command("git", "checkout", "HEAD~1", "--detach").Run()

	command := exec.Command(e2e.CDX, "tag", "latest", "-n", "app", "promoted", "--head")
	var stdOut bytes.Buffer
	var stdErr bytes.Buffer
	command.Stdout = &stdOut
	command.Stderr = &stdErr
	err := command.Run()
	check.Equals(t, "", stdErr.String())
	check.Ok(t, err)
	check.Equals(t, "0.1.0\n", stdOut.String())
}
