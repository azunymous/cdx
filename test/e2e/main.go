package e2e

import (
	"github.com/sirupsen/logrus"
	"os"
)

var CDX = "cdx"

func init() {
	cmd, set := os.LookupEnv("CDX_CMD")
	if set {
		logrus.Info("Using command " + cmd)
		CDX = cmd
	}
}
