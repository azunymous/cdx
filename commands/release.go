package commands

import (
	"cd/commands/options"
	"github.com/spf13/cobra"
	"log"
)

// addRelease adds the increment command to a top level command.
func addRelease(topLevel *cobra.Command) {
	appOpts := &options.App{}

	releaseCmd := &cobra.Command{
		Use:   "release",
		Short: "Release a new version",
		Long: `The release command increments the version via a git tag
`,
		Run: func(cmd *cobra.Command, args []string) {
			err := release(cmd, args, appOpts)
			if err != nil {
				log.Fatal(err)
			}
		},
	}

	options.AddNameArg(releaseCmd, appOpts)
	topLevel.AddCommand(releaseCmd)
}

func release(cmd *cobra.Command, args []string, appOpts *options.App) error {
	return cmd.Help()
}
