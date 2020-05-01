package commands

import (
	"github.com/azunymous/cdx/commands/options"
	"github.com/azunymous/cdx/vcs"
	"github.com/azunymous/cdx/vcs/gogit"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// addRelease adds the increment command to a top level command.
func addRelease(topLevel *cobra.Command, app *options.App) {
	incrOpts := &options.Increment{}
	gitOpts := &options.Git{}
	releaseCmd := &cobra.Command{
		Use:   "release",
		Short: "Release a new version",
		Long: `The release command increments the version via a git tag
`,
		Run: func(cmd *cobra.Command, args []string) {
			err := release(app, incrOpts, gitOpts)
			if err != nil {
				logrus.Fatal(err)
			}
		},
	}
	options.AddIncrementArg(releaseCmd, incrOpts)
	options.AddPushArg(releaseCmd, gitOpts)
	topLevel.AddCommand(releaseCmd)
}

func release(app *options.App, incr *options.Increment, gitOpts *options.Git) error {
	logrus.Printf("Releasing %v", app.Name)
	v, err := vcs.NewGit(app.Name, incr.GetField(), gitOpts.Push, func() (vcs.Repository, error) { return gogit.NewRepo() })
	if err != nil {
		return err
	}
	if !v.Ready() {
		return nil
	}
	err = v.Release()
	if err != nil {
		return err
	}
	return v.Distribute()
}
