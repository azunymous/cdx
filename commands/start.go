package commands

import (
	"github.com/azunymous/cdx/commands/options"
	"github.com/azunymous/cdx/watch"
	"github.com/azunymous/cdx/watch/gocache"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"time"
)

// addRelease adds the increment command to a top level command.
func addStart(topLevel *cobra.Command) {
	patchOpts := &options.Patch{}
	createCmd := &cobra.Command{
		Use:   "start",
		Short: "Start a sharing server",
		Long: `The start command runs a server to facilitate sharing workspaces
`,
		Run: func(cmd *cobra.Command, args []string) {
			err := start(patchOpts)
			if err != nil {
				logrus.Fatal(err)
			}
		},
	}
	options.AddInsecureArg(createCmd, patchOpts)
	topLevel.AddCommand(createCmd)
}

const defaultExpiration = 5 * time.Minute
const cleanupInterval = 10 * time.Minute

func start(patchOpts *options.Patch) error {
	logrus.Printf("Sharing ")
	return watch.NewServer(gocache.NewGoCache(defaultExpiration, cleanupInterval), patchOpts.Insecure)
}
