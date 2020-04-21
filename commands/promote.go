package commands

import (
	"cdx/commands/options"
	"cdx/vcs"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// addRelease adds the increment command to a top level command.
func addPromote(topLevel *cobra.Command, app *options.App) {
	gitOpts := &options.Git{}
	releaseCmd := &cobra.Command{
		Use:   "promote <stage>",
		Short: "Promote this commit",
		Long: `The promote command promotes the current commit via a git tag
`,
		Run: func(cmd *cobra.Command, args []string) {
			err := promote(app, gitOpts, args[0])
			if err != nil {
				logrus.Fatal(err)
			}
		},
		Args: cobra.ExactArgs(1),
	}
	options.AddPushArg(releaseCmd, gitOpts)
	topLevel.AddCommand(releaseCmd)
}

func promote(app *options.App, git *options.Git, stage string) error {
	logrus.Printf("Promoting %s to %s", app.Name, stage)
	repo, err := vcs.NewRepo()
	if err != nil {
		return err
	}
	err = repo.Promote(app.Name, stage)
	if err != nil {
		return err
	}
	if git.Push {
		return repo.PushTags()
	}
	return nil
}
