package watch

import (
	"bytes"
	"context"
	"errors"
	"github.com/azunymous/cdx/watch/diff"
	"google.golang.org/grpc"
	"os/exec"
)

func NewShareClient() (diff.DiffClient, func(), error) {
	target := ":19443"
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}

	return diff.NewDiffClient(conn),
		func() {
			defer conn.Close()
		}, nil
}

func Upload(ctx context.Context, client diff.DiffClient, name string) error {
	commits, err := getDiffCommits(ctx)
	if err != nil {
		return err
	}
	reply := &diff.DiffCommits{
		Name:    name,
		Commits: commits,
	}
	_, err = client.UploadDiff(ctx, reply)
	return err
}

func getDiffCommits(ctx context.Context) (string, error) {
	cmd := exec.CommandContext(ctx, "git", "format-patch", "origin/master", "--stdout")
	var stdOut bytes.Buffer
	var stdErr bytes.Buffer
	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr

	err := cmd.Run()
	if err != nil {
		return "", errors.New(stdErr.String())
	}
	return stdOut.String(), nil
}
