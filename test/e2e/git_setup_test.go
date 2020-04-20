package e2e

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
