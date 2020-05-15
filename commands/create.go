package commands

import (
	"github.com/azunymous/cdx/watch"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// addRelease adds the increment command to a top level command.
func addCreate(topLevel *cobra.Command) {
	createCmd := &cobra.Command{
		Use:   "create",
		Short: "Start sharing a workspace",
		Long: `The create command runs a server to share your workspace
`,
		Run: func(cmd *cobra.Command, args []string) {
			err := create()
			if err != nil {
				logrus.Fatal(err)
			}
		},
	}
	topLevel.AddCommand(createCmd)
}

func create() error {
	logrus.Printf("Sharing ")
	return watch.NewServer()
}
