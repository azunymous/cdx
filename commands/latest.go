package commands

import (
	"fmt"
	"github.com/azunymous/cdx/commands/options"
	"github.com/azunymous/cdx/vcs"
	"github.com/azunymous/cdx/vcs/gogit"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// addLatest adds the latest command to a top level command.
func addLatest(topLevel *cobra.Command, app *options.App) {
	gitOpts := &options.Git{}
	latestCmd := &cobra.Command{
		Use:   "latest [promotion stage]",
		Short: "Get the latest version of an application",
		Long: `The latest command fetches the latest version of an application from git tags. 
If a stage is specified, the latest version promoted to that stage is returned.
`,
		Run: func(cmd *cobra.Command, args []string) {
			err := latest(args, app, gitOpts)
			if err != nil {
				logrus.Fatal(err)
			}
		},
		Args: cobra.MaximumNArgs(1),
	}

	options.AddHeadOnlyArg(latestCmd, gitOpts)
	options.AddFallbackHashArg(latestCmd, gitOpts)
	topLevel.AddCommand(latestCmd)
}

func latest(args []string, app *options.App, gitOpts *options.Git) error {
	stage := ""
	if len(args) > 0 {
		stage = args[0]
	}
	v, err := vcs.NewGit(app.Name, -1, false, func() (vcs.Repository, error) { return gogit.NewRepo() })
	if err != nil {
		return err
	}
	version, err := v.Version(stage, gitOpts.HeadOnly, gitOpts.FallbackHash)
	if err != nil {
		return err
	}
	fmt.Println(version)
	return nil
}
