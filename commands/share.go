package commands

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// addTag adds the primary tag command to a top level command.
func addShare(topLevel *cobra.Command) {
	shareCmd := &cobra.Command{
		Use:   "share",
		Short: "Share repositories",
		Long: `The share command allows you to share
your git repository
`,
		Run: func(cmd *cobra.Command, args []string) {
			err := share(cmd)
			if err != nil {
				logrus.Fatal(err)
			}
		},
		Args: cobra.NoArgs,
	}

	addStart(shareCmd)
	addApply(shareCmd)
	addUpload(shareCmd)
	topLevel.AddCommand(shareCmd)
}

func share(cmd *cobra.Command) error {
	return cmd.Help()
}