//+build e2e

package share

import (
	"github.com/azunymous/cdx/test/check"
	"github.com/azunymous/cdx/test/e2e"
	"os/exec"
	"strings"
	"testing"
)

func TestUploadUploadsPatch(t *testing.T) {
	server, c := e2e.StartCdxShareServer()
	t.Cleanup(c)

	fn := e2e.CreateTempGitDir()
	_ = e2e.CreateTempGitRemote(fn)

	command := exec.Command(e2e.CDX, "share", "upload", "patchName", "--insecure", "--uri", server)
	err := command.Run()
	check.Ok(t, err)
}

func TestUploadUploadsPatchCanBeDownloaded(t *testing.T) {
	server, c := e2e.StartCdxShareServer()
	t.Cleanup(c)

	dir := e2e.CreateTempGitDir()
	_ = e2e.CreateTempGitRemote(dir)
	e2e.CreateCommit(dir, "Commit 2")

	command := exec.Command(e2e.CDX, "share", "upload", "patchName", "--insecure", "--uri", server)
	err := command.Run()
	check.Ok(t, err)

	dir2 := e2e.CreateTempGitDir()
	_ = e2e.CreateTempGitRemote(dir2)
	command = exec.Command(e2e.CDX, "share", "apply", "patchName", "--insecure", "--uri", server)
	err = command.Run()
	check.Ok(t, err)
	gitLog, err := exec.Command("git", "log", "-1", "--pretty=%B").Output()
	check.Ok(t, err)
	check.Equals(t, `"Commit 2"`, strings.TrimSpace(string(gitLog)))
}

func TestUploadCannotUploadExistingPatch(t *testing.T) {
	server, c := e2e.StartCdxShareServer()
	t.Cleanup(c)

	fn := e2e.CreateTempGitDir()
	_ = e2e.CreateTempGitRemote(fn)

	command := exec.Command(e2e.CDX, "share", "upload", "patchName", "--insecure", "--uri", server)
	err := command.Run()
	check.Ok(t, err)

	command = exec.Command(e2e.CDX, "share", "upload", "patchName", "--insecure", "--uri", server)
	err = command.Run()
	check.Assert(t, err != nil, "Expected error not to be nil, got: %v", err)
}

func TestUploadUploadsPatchWithPasswordProtection(t *testing.T) {
	server, c := e2e.StartCdxShareServer()
	t.Cleanup(c)

	dir := e2e.CreateTempGitDir()
	_ = e2e.CreateTempGitRemote(dir)
	e2e.CreateCommit(dir, "Commit 2")

	command := exec.Command(e2e.CDX, "share", "upload", "patchName", "--insecure", "--uri", server, "--password", "testPassword")
	err := command.Run()
	check.Ok(t, err)

	dir2 := e2e.CreateTempGitDir()
	_ = e2e.CreateTempGitRemote(dir2)
	command = exec.Command(e2e.CDX, "share", "apply", "patchName", "--insecure", "--uri", server, "--password", "testPassword")
	err = command.Run()
	check.Ok(t, err)
	gitLog, err := exec.Command("git", "log", "-1", "--pretty=%B").Output()
	check.Ok(t, err)
	check.Equals(t, `"Commit 2"`, strings.TrimSpace(string(gitLog)))
}

func TestUploadUploadsPatchWithPasswordRequiresPassword(t *testing.T) {
	server, c := e2e.StartCdxShareServer()
	t.Cleanup(c)

	dir := e2e.CreateTempGitDir()
	_ = e2e.CreateTempGitRemote(dir)
	e2e.CreateCommit(dir, "Commit 2")

	command := exec.Command(e2e.CDX, "share", "upload", "patchName", "--insecure", "--uri", server, "--password", "testPassword")
	err := command.Run()
	check.Ok(t, err)

	dir2 := e2e.CreateTempGitDir()
	_ = e2e.CreateTempGitRemote(dir2)
	command = exec.Command(e2e.CDX, "share", "apply", "patchName", "--insecure", "--uri", server, "--password", "wrongPassword")
	err = command.Run()
	check.Assert(t, err != nil, "Expected error not to be nil, got nil")
	gitLog, err := exec.Command("git", "log", "-1", "--pretty=%B").Output()
	check.Ok(t, err)
	check.Assert(t, `"Commit 2"` != strings.TrimSpace(string(gitLog)), "Expected commit not to be found, instead found %s", string(gitLog))
}
