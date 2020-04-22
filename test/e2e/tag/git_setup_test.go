//+build e2e

package tag

import (
	"io/ioutil"
	"os"
	"os/exec"
)

func createTempGitDir() string {
	fn, _ := ioutil.TempDir(os.TempDir(), "cdx-test-")
	_ = os.Chdir(fn)
	_, _ = exec.Command("git", "init").CombinedOutput()
	_, _ = exec.Command("bash", "-c", "echo 'hello world' > file.txt").CombinedOutput()
	_, _ = exec.Command("git", "add", "file.txt").CombinedOutput()
	_, _ = exec.Command("git", "commit", "-m", `"Commit 1"`).CombinedOutput()
	return fn
}

func createTag(dir, tag string) {
	_ = os.Chdir(dir)
	_, _ = exec.Command("git", "tag", tag).CombinedOutput()
}

func createCommit(dir, msg string) {
	_ = os.Chdir(dir)
	_, _ = exec.Command("bash", "-c", "echo '"+msg+"' > file.txt").CombinedOutput()
	_, _ = exec.Command("git", "add", "file.txt").CombinedOutput()
	_, _ = exec.Command("git", "commit", "-m", `"`+msg+`"`).CombinedOutput()
}

func createTempGitRemote(gitDir string) string {
	remoteDir, _ := ioutil.TempDir(os.TempDir(), "cdx-remote-*.git")
	_ = os.Chdir(remoteDir)
	_, _ = exec.Command("git", "init", "--bare").CombinedOutput()

	_ = os.Chdir(gitDir)
	_, _ = exec.Command("git", "remote", "add", "origin", remoteDir).CombinedOutput()
	_, _ = exec.Command("git", "push", "origin", "master").CombinedOutput()

	return remoteDir
}
