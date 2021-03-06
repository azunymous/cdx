package commands

import (
	"github.com/azunymous/cdx/commands/options"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// addTag adds the primary tag command to a top level command.
func addTag(topLevel *cobra.Command) {
	appOpts := &options.App{}

	tagCmd := &cobra.Command{
		Use:   "tag",
		Short: "Tag git repositories",
		Long: `The tag command allows you to semantically version and promote
your git repository
`,
		Run: func(cmd *cobra.Command, args []string) {
			err := tag(cmd, args, appOpts)
			if err != nil {
				logrus.Fatal(err)
			}
		},
		Args: cobra.NoArgs,
	}

	options.AddNameArg(tagCmd, appOpts)
	topLevel.AddCommand(tagCmd)

	addRelease(tagCmd, appOpts)
	addPromote(tagCmd, appOpts)
	addLatest(tagCmd, appOpts)
}

func tag(cmd *cobra.Command, args []string, appOpts *options.App) error {
	return cmd.Help()
}
