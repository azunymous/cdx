package options

import "github.com/spf13/cobra"

// App struct contains options regarding the application
type Patch struct {
	Reset    bool
	Insecure bool
}

func AddResetArg(cmd *cobra.Command, p *Patch) {
	cmd.Flags().BoolVarP(&p.Reset, "reset", "r", true, "Reset (hard) to origin/master before applying patch")
}

func AddInsecureArg(cmd *cobra.Command, p *Patch) {
	cmd.Flags().BoolVarP(&p.Insecure, "insecure", "i", false, "Insecure connection to GRPC server")
}
