package watch

import (
	"context"
	"fmt"
	"github.com/azunymous/cdx/watch/diff"
	"github.com/golang/protobuf/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

type DiffServer struct {
	db DiffStore
}

type DiffStore interface {
	Get(key string) (string, error)
	Set(key string, value string) error
}

func NewServer(store DiffStore, port int, insecure bool) error {
	srv := &DiffServer{db: store}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var srvOpts []grpc.ServerOption
	if !insecure {
		c, err := credentials.NewServerTLSFromFile("server.pem", "server.key")
		if err != nil {
			return fmt.Errorf("failed to find certificate files server.pem and server.key: %v", err)
		}
		srvOpts = append(srvOpts, grpc.Creds(c))
	}
	grpcServer := grpc.NewServer(srvOpts...)
	diff.RegisterDiffServer(grpcServer, srv)
	return grpcServer.Serve(lis)
}

func (d *DiffServer) SendDiff(_ context.Context, request *diff.DiffRequest) (*diff.DiffCommits, error) {
	logrus.Infof("Received message %s", request.Name)
	message, err := d.db.Get(request.GetName())
	if err != nil {
		return nil, err
	}
	reply := &diff.DiffCommits{}
	err = proto.UnmarshalText(message, reply)
	return reply, err
}

func (d *DiffServer) UploadDiff(_ context.Context, reply *diff.DiffCommits) (*diff.DiffConfirm, error) {
	logrus.Infof("Received message %s", reply.Name)
	err := d.db.Set(reply.GetName(), proto.MarshalTextString(reply))
	if err != nil {
		return nil, err
	}
	return &diff.DiffConfirm{}, nil
}
