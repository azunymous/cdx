//+build e2e

package tag

import (
	"github.com/azunymous/cdx/test/check"
	"os/exec"
	"testing"
)

func TestTagCommandExists(t *testing.T) {
	command := exec.Command(cdxCmd, "tag", "-n", "app")
	err := command.Run()
	check.Ok(t, err)
}
