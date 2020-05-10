package options

import (
	"github.com/spf13/cobra"
)

// Git struct contains options regarding git settings
type Git struct {
	Push         bool
	HeadOnly     bool
	FallbackHash bool
}

func AddPushArg(cmd *cobra.Command, g *Git) {
	cmd.Flags().BoolVarP(&g.Push, "push", "p", false, "Push tags (must be on origin/master)")
}

func AddHeadOnlyArg(cmd *cobra.Command, g *Git) {
	cmd.Flags().BoolVar(&g.HeadOnly, "head", false, "Look at tags at HEAD only")
}

func AddFallbackHashArg(cmd *cobra.Command, g *Git) {
	cmd.Flags().BoolVar(&g.FallbackHash, "fallback", false, "Fallback to git hash of current commit if no tag found. Must be used with --head.")
}
