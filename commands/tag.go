package commands

import (
	"cdx/commands/options"
	"github.com/spf13/cobra"
	"log"
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
				log.Fatal(err)
			}
		},
		Args: cobra.NoArgs,
	}

	options.AddNameArg(tagCmd, appOpts)
	topLevel.AddCommand(tagCmd)

	addRelease(tagCmd, appOpts)
}

func tag(cmd *cobra.Command, args []string, appOpts *options.App) error {
	return cmd.Help()
}
