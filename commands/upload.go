package commands

import (
	"context"
	"github.com/azunymous/cdx/commands/options"
	"github.com/azunymous/cdx/watch"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// addUpload adds the upload command to a top level command
func addUpload(topLevel *cobra.Command) {
	patchOpts := &options.Patch{}
	uploadCmd := &cobra.Command{
		Use:   "upload [patch name]",
		Short: "Upload uploads your workspace's git commits changed from origin/master",
		Long: `The upload command connects to a server to upload your workspace under a specified name
`,
		Run: func(cmd *cobra.Command, args []string) {
			err := upload(args[0], patchOpts)
			if err != nil {
				logrus.Fatal(err)
			}
		},
		Args: cobra.ExactArgs(1),
	}
	options.AddInsecureArg(uploadCmd, patchOpts)
	options.AddTargetArg(uploadCmd, patchOpts)
	topLevel.AddCommand(uploadCmd)
}

func upload(name string, opts *options.Patch) error {
	logrus.Printf("Uploading ")
	c, closeFunc, err := watch.NewShareClient(opts.Target, opts.Insecure)
	ctx := context.Background()
	if err != nil {
		return err
	}
	defer closeFunc()

	return watch.Upload(ctx, c, name)
}
