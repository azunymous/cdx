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

func Watch() error {

	return nil
}

type DiffServer struct {
}

func NewServer() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	diff.RegisterDiffServer(grpcServer, &DiffServer{})
	return grpcServer.Serve(lis)
}

func (d DiffServer) SendDiff(ctx context.Context, request *diff.DiffRequest) (*diff.DiffReply, error) {
	logrus.Infof("Received message %s", request.Name)
	return &diff.DiffReply{Message: "Hello world!"}, nil
}
