package gogit

import (
	"context"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"strings"
	"time"
)

const basicAuthEnv = "CDX_BASIC_AUTH"
const noGitEnv = "CDX_NO_GIT"

func (r *Repo) PushTags() error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelFunc()

	if os.Getenv(basicAuthEnv) != "" || os.Getenv(noGitEnv) != "" {
		rs := config.RefSpec("refs/tags/*:refs/tags/*")
		return r.gitRepo.PushContext(
			ctx,
			&git.PushOptions{RefSpecs: []config.RefSpec{rs}, Auth: getAuth()})
	}
	logrus.Info("Pushing using git command")
	cmd := exec.CommandContext(ctx, "git", "push", "--tags", "-q")

	out, err := cmd.CombinedOutput()
	logrus.Info(string(out))
	if err != nil {
		return err
	}
	return nil
}

func getAuth() transport.AuthMethod {
	if os.Getenv(basicAuthEnv) != "" {
		auth := strings.SplitN(os.Getenv(basicAuthEnv), ":", 2)
		if len(auth) < 2 || auth[0] == "" || auth[1] == "" {
			logrus.Fatal("Invalid basic auth provided")
		}
		logrus.Info("Using basic auth:", auth)
		return &http.BasicAuth{
			Username: auth[0],
			Password: auth[1],
		}
	}
	return nil
}
