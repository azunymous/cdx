package options

import "github.com/spf13/cobra"

// App struct contains options regarding the application
type Patch struct {
	Reset    bool
	Insecure bool
	Target   string
	Port     int
}

func AddResetArg(cmd *cobra.Command, p *Patch) {
	cmd.Flags().BoolVarP(&p.Reset, "reset", "r", true, "Reset (hard) to origin/master before applying patch")
}

func AddInsecureArg(cmd *cobra.Command, p *Patch) {
	cmd.Flags().BoolVarP(&p.Insecure, "insecure", "i", false, "Insecure connection to GRPC server")
}

const defaultServer = "cdx.vvv.run:443"

func AddTargetArg(cmd *cobra.Command, r *Patch) {
	cmd.Flags().StringVarP(&r.Target, "uri", "u", defaultServer, "URI to use for pushing and pulling patches")
}

func AddPortArg(cmd *cobra.Command, r *Patch) {
	cmd.Flags().IntVarP(&r.Port, "port", "p", 19443, "Server port to use")
}
