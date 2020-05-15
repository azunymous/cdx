package watch

import (
	"context"
	"fmt"
	"github.com/azunymous/cdx/watch/diff"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"log"
	"net"
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

func (d DiffServer) SendDiff(ctx context.Context, request *diff.DiffRequest) (*diff.DiffCommits, error) {
	logrus.Infof("Received message %s", request.Name)
	commits, err := getDiff(request.GetName())
	if err != nil {
		return nil, err
	}
	return &diff.DiffCommits{Commits: commits}, nil
}

func (d DiffServer) UploadDiff(ctx context.Context, reply *diff.DiffCommits) (*diff.DiffConfirm, error) {
	logrus.Infof("Received message %s", reply.Name)
	err := saveDiff(reply.GetName(), reply.GetCommits())
	if err != nil {
		return nil, err
	}
	return &diff.DiffConfirm{}, nil
}

func getDiff(name string) (string, error) {
	return "", nil
}

func saveDiff(name string, commits string) error {
	return nil
}
