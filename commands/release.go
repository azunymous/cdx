package commands

import (
	"cdx/commands/options"
	"cdx/vcs"
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

func release(app *options.App, incr *options.Increment, git *options.Git) error {
	logrus.Printf("Releasing %v", app.Name)
	repo, err := vcs.NewRepo()
	if err != nil {
		return err
	}
	err = repo.IncrementTag(app.Name, incr.GetField())
	if err != nil {
		return err
	}
	if git.Push {
		return repo.PushTags()
	}
	return nil
}
