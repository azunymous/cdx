//+build e2e

package tag

import (
	"cdx/test/check"
	"os/exec"
	"testing"
)

func TestTagCommandExists(t *testing.T) {
	command := exec.Command("cdx", "tag", "-n", "app")
	err := command.Run()
	check.Ok(t, err)
}
