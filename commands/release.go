package commands

import (
	"cdx/commands/options"
	"github.com/spf13/cobra"
	"log"
)

// addRelease adds the increment command to a top level command.
func addRelease(topLevel *cobra.Command, app *options.App) {

	releaseCmd := &cobra.Command{
		Use:   "release",
		Short: "Release a new version",
		Long: `The release command increments the version via a git tag
`,
		Run: func(cmd *cobra.Command, args []string) {
			err := release(args, app)
			if err != nil {
				log.Fatal(err)
			}
		},
	}

	topLevel.AddCommand(releaseCmd)
}

func release(args []string, app *options.App) error {
	log.Printf("Releasing %v", app.Name)
	return nil
}
