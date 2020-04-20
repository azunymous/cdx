//+build e2e

package e2e

import (
	"os/exec"
	"testing"
)

func TestReleaseCommandExists(t *testing.T) {
	command := exec.Command("cdx", "tag", "release", "-n", "app")
	err := command.Run()

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}
