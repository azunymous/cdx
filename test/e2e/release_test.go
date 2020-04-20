//+build e2e

package e2e

import (
	"cdx/internal/check"
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
