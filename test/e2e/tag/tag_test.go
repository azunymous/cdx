//+build e2e

package tag

import (
	"github.com/azunymous/cdx/test/check"
	"github.com/azunymous/cdx/test/e2e"
	"os/exec"
	"testing"
)

func TestTagCommandExists(t *testing.T) {
	command := exec.Command(e2e.CDX, "tag", "-n", "app")
	err := command.Run()
	check.Ok(t, err)
}
