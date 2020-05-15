package commands

import (
	"bytes"
	"context"
	"github.com/azunymous/cdx/watch"
	"github.com/azunymous/cdx/watch/diff"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os/exec"
	"strings"
)

// addRelease adds the increment command to a top level command.
func addJoin(topLevel *cobra.Command) {
	joinCmd := &cobra.Command{
		Use:   "join",
		Short: "Join a shared workspace",
		Long: `The join command connects server to update your workspace
`,
		Run: func(cmd *cobra.Command, args []string) {
			err := join()
			if err != nil {
				logrus.Fatal(err)
			}
		},
	}
	topLevel.AddCommand(joinCmd)
}

func join() error {
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
	logrus.Info(strings.SplitN(resp.Committed, "\n", 2)[0])
	cmd := exec.CommandContext(ctx, "git", "am", "-")
	cmd.Stdin = bytes.NewBufferString(resp.Committed)
	output, err := cmd.CombinedOutput()
	logrus.Info(string(output))
	return err

}
