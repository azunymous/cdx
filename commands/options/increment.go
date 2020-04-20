package options

import (
	"cdx/vcs"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"strings"
)

// App struct contains options regarding the application
type Increment struct {
	Field string
}

func AddIncrementArg(cmd *cobra.Command, r *Increment) {
	cmd.Flags().StringVarP(&r.Field, "increment", "i", "minor", "Semantic version field to increment")
}

func (i *Increment) GetField() vcs.Field {
	i.Field = strings.ToLower(i.Field)
	switch i.Field {
	case "patch":
		return vcs.Patch
	case "minor":
		return vcs.Minor
	case "major":
		return vcs.Major
	}
	logrus.Fatalln("Increment field was not one of [major, minor, patch]")
	return -1
}
