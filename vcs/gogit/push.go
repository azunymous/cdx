package gogit

import (
	"context"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"time"
)

func (r *Repo) PushTags() error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelFunc()

	rs := config.RefSpec("refs/tags/*:refs/tags/*")
	return r.gitRepo.PushContext(ctx, &git.PushOptions{RefSpecs: []config.RefSpec{rs}})
}
