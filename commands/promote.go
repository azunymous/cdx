package commands

import (
	"cdx/commands/options"
	"cdx/vcs"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// addRelease adds the increment command to a top level command.
func addPromote(topLevel *cobra.Command, app *options.App) {
	releaseCmd := &cobra.Command{
		Use:   "promote <stage>",
		Short: "Promote this commit",
		Long: `The promote command promotes the current commit via a git tag
`,
		Run: func(cmd *cobra.Command, args []string) {
			err := promote(app, args[0])
			if err != nil {
				logrus.Fatal(err)
			}
		},
		Args: cobra.ExactArgs(1),
	}
	topLevel.AddCommand(releaseCmd)
}

func promote(app *options.App, stage string) error {
	logrus.Printf("Releasing %v", app.Name)
	repo, err := vcs.NewRepo()
	if err != nil {
		return err
	}
	return repo.Promote(app.Name, stage)
}
