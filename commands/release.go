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
	releaseCmd := &cobra.Command{
		Use:   "release",
		Short: "Release a new version",
		Long: `The release command increments the version via a git tag
`,
		Run: func(cmd *cobra.Command, args []string) {
			err := release(app, incrOpts)
			if err != nil {
				logrus.Fatal(err)
			}
		},
	}
	options.AddIncrementArg(releaseCmd, incrOpts)
	topLevel.AddCommand(releaseCmd)
}

func release(app *options.App, incr *options.Increment) error {
	logrus.Printf("Releasing %v", app.Name)
	repo, err := vcs.NewRepo()
	if err != nil {
		return err
	}
	return repo.IncrementTag(app.Name, incr.GetField())
}
