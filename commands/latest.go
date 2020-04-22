package commands

import (
	"cdx/commands/options"
	"cdx/vcs"
	"errors"
	"fmt"
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
	topLevel.AddCommand(latestCmd)
}

func latest(args []string, app *options.App, git *options.Git) error {
	stage := ""
	if len(args) > 0 {
		stage = args[0]
	}
	repo, err := vcs.NewRepo()
	if err != nil {
		return err
	}
	if git.HeadOnly {
		return getHeadTags(repo, app, stage)
	}
	return getModuleTags(repo, app, stage)
}

func getHeadTags(repo *vcs.Repo, app *options.App, stage string) error {
	tagsForHead, err := repo.TagsForHead(app.Name, stage)
	if err != nil {
		return err
	}
	if len(tagsForHead) == 0 {
		return errors.New("no tags found at HEAD")
	}
	fmt.Println(vcs.VersionFrom(tagsForHead[len(tagsForHead)-1]))
	return nil
}

func getModuleTags(repo *vcs.Repo, app *options.App, stage string) error {
	tagsForModule, err := repo.TagsForModule(app.Name, stage)
	if err != nil {
		return err
	}
	if len(tagsForModule) == 0 {
		return errors.New("no tags found for module and stage")
	}
	fmt.Println(vcs.VersionFrom(tagsForModule[len(tagsForModule)-1]))
	return nil
}
