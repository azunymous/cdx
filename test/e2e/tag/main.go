package tag

import (
	"github.com/sirupsen/logrus"
	"os"
)

var cdxCmd = "cdx"

func init() {
	cmd, set := os.LookupEnv("CDX_CMD")
	if set {
		logrus.Info("Using command " + cmd)
		cdxCmd = cmd
	}
}
