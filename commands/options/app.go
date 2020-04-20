package options

import (
	"github.com/spf13/cobra"
)

// App struct contains options regarding the application
type App struct {
	Name string
}

func AddNameArg(cmd *cobra.Command, r *App) {
	cmd.PersistentFlags().StringVarP(&r.Name, "name", "n", "", "Application or module name")
	must(cmd.MarkPersistentFlagRequired("name"))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
