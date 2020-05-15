package commands

import (
	"bytes"
	"context"
	"github.com/azunymous/cdx/commands/options"
	"github.com/azunymous/cdx/watch"
	"github.com/azunymous/cdx/watch/diff"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os/exec"
)

// addRelease adds the increment command to a top level command.
func addApply(topLevel *cobra.Command) {
	patchOpts := &options.Patch{}
	applyCmd := &cobra.Command{
		Use:   "apply",
		Short: "Apply a shared workspace's changed to your local changes",
		Long: `The apply command connects to a server to update your workspace
`,
		Run: func(cmd *cobra.Command, args []string) {
			err := apply(patchOpts)
			if err != nil {
				logrus.Fatal(err)
			}
		},
	}
	options.AddResetArg(applyCmd, patchOpts)
	topLevel.AddCommand(applyCmd)
}

func apply(patchOpts *options.Patch) error {
	logrus.Printf("Joining ")
	c, closeFunc, err := watch.NewClient()
	if err != nil {
		return err
	}
	defer closeFunc()
	in := &diff.DiffRequest{Name: "from diff client"}
	ctx := context.Background()
	resp, err := c.SendDiff(ctx, in)
	if err != nil {
		return err
	}
	logrus.Info(resp.Committed)
	if patchOpts.Reset {
		output, err := exec.CommandContext(ctx, "git", "reset", "origin/master", "--hard").CombinedOutput()
		if err != nil {
			return err
		}
		logrus.Infof("Resetting to origin/master: \n%s", string(output))
	}
	cmd := exec.CommandContext(ctx, "git", "am", "-")
	cmd.Stdin = bytes.NewBufferString(resp.Committed)
	output, err := cmd.CombinedOutput()
	logrus.Info(string(output))
	return err

}
