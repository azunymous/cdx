package commands

import (
	"context"
	"github.com/azunymous/cdx/watch"
	"github.com/azunymous/cdx/watch/diff"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
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
	resp, err := c.SendDiff(context.Background(), in)
	if err != nil {
		return err
	}
	logrus.Info(resp.String())
	return nil
}
