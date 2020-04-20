package e2e

import (
	"cdx/internal/check"
	"fmt"
	"os/exec"
	"sort"
	"strings"
	"testing"
)

func TestPromoteOpensRepository(t *testing.T) {
	dir := createTempGitDir()
	createTag(dir, "app-0.1.0")
	command := exec.Command("cdx", "tag", "promote", "-n", "app", "promoted")
	b, err := command.CombinedOutput()
	fmt.Println(string(b))
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
