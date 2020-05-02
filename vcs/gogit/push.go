package gogit

import (
	"context"
	"github.com/sirupsen/logrus"
	"os/exec"
	"time"
)

func (r *Repo) PushTags() error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelFunc()

	logrus.Info("Pushing using git command")
	cmd := exec.CommandContext(ctx, "git", "push", "--tags", "-q")

	out, err := cmd.CombinedOutput()
	logrus.Info(string(out))
	if err != nil {
		return err
	}
	return nil
}
