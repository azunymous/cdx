package options

import (
	"github.com/spf13/cobra"
)

// Git struct contains options regarding git settings
type Git struct {
	Push bool
}

func AddPushArg(cmd *cobra.Command, g *Git) {
	cmd.Flags().BoolVarP(&g.Push, "push", "p", false, "Push tags")
}
