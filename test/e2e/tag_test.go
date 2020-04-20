//+build e2e

package e2e

import (
	"cdx/internal/check"
	"os/exec"
	"testing"
)

func TestTagCommandExists(t *testing.T) {
	command := exec.Command("cdx", "tag", "-n", "app")
	err := command.Run()
	check.Ok(t, err)
}
