package watch

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/azunymous/cdx/watch/diff"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"log"
	"net"
	"os/exec"
)

type DiffServer struct {
}

func NewServer() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 19443))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	diff.RegisterDiffServer(grpcServer, &DiffServer{})
	return grpcServer.Serve(lis)
}

func (d DiffServer) SendDiff(ctx context.Context, request *diff.DiffRequest) (*diff.DiffReply, error) {
	logrus.Infof("Received message %s", request.Name)
	//diffs := getDiff(ctx)
	commits, err := getDiffCommits(ctx)
	if err != nil {
		return nil, err
	}
	return &diff.DiffReply{Committed: commits}, nil
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
