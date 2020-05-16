package options

import "github.com/spf13/cobra"

// App struct contains options regarding the application
type Patch struct {
	Reset    bool
	Insecure bool
	Target   string
}

func AddResetArg(cmd *cobra.Command, p *Patch) {
	cmd.Flags().BoolVarP(&p.Reset, "reset", "r", true, "Reset (hard) to origin/master before applying patch")
}

// TODO This is a bad default! Set this way currently for convenience. This will be fixed when a real default server is made
func AddInsecureArg(cmd *cobra.Command, p *Patch) {
	cmd.Flags().BoolVarP(&p.Insecure, "insecure", "i", true, "Insecure connection to GRPC server")
}

const defaultServer = "35.209.179.20:30443"

func AddTargetArg(cmd *cobra.Command, r *Patch) {
	cmd.Flags().StringVarP(&r.Target, "uri", "u", defaultServer, "URI to use for pushing and pulling patches")
}
