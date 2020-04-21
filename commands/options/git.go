package options

import (
	"github.com/spf13/cobra"
)

// Git struct contains options regarding git settings
type Git struct {
	Push     bool
	HeadOnly bool
}

func AddPushArg(cmd *cobra.Command, g *Git) {
	cmd.Flags().BoolVarP(&g.Push, "push", "p", false, "Push tags")
}

func AddHeadOnlyArg(cmd *cobra.Command, g *Git) {
	cmd.Flags().BoolVar(&g.HeadOnly, "head", false, "Look at tags at HEAD only")
}
