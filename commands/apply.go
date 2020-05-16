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
		Use:   "apply [patch name]",
		Short: "Apply a shared workspace's changed to your local changes",
		Long: `The apply command connects to a server to update your workspace with the required patch
`,
		Run: func(cmd *cobra.Command, args []string) {
			err := apply(args[0], patchOpts)
			if err != nil {
				logrus.Fatal(err)
			}
		},
		Args:    cobra.ExactArgs(1),
		Aliases: []string{"download"},
	}
	options.AddResetArg(applyCmd, patchOpts)
	options.AddInsecureArg(applyCmd, patchOpts)
	options.AddTargetArg(applyCmd, patchOpts)
	topLevel.AddCommand(applyCmd)
}

func apply(name string, patchOpts *options.Patch) error {
	logrus.Printf("Applying ")
	c, closeFunc, err := watch.NewClient(patchOpts.Target, patchOpts.Insecure)
	if err != nil {
		return err
	}
	defer closeFunc()
	in := &diff.DiffRequest{Name: name}
	ctx := context.Background()
	resp, err := c.SendDiff(ctx, in)
	if err != nil {
		return err
	}

	// This is true by default
	if patchOpts.Reset {
		output, err := exec.CommandContext(ctx, "git", "reset", "origin/master", "--hard").CombinedOutput()
		if err != nil {
			return err
		}
		logrus.Infof("Resetting to origin/master: \n%s", string(output))
	}
	cmd := exec.CommandContext(ctx, "git", "am", "-")
	cmd.Stdin = bytes.NewBufferString(resp.GetCommits())
	output, err := cmd.CombinedOutput()
	logrus.Info(string(output))
	return err

}
