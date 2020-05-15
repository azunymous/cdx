package watch

import (
	"context"
	"fmt"
	"github.com/azunymous/cdx/watch/diff"
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

const defaultExpiration = 5 * time.Minute
const cleanupInterval = 10 * time.Minute

type DiffServer struct {
	db *cache.Cache
}

func NewServer() error {
	c := cache.New(defaultExpiration, cleanupInterval)
	srv := &DiffServer{db: c}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 19443))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	diff.RegisterDiffServer(grpcServer, srv)
	return grpcServer.Serve(lis)
}

func (d *DiffServer) SendDiff(ctx context.Context, request *diff.DiffRequest) (*diff.DiffCommits, error) {
	logrus.Infof("Received message %s", request.Name)
	commits, err := d.getDiff(request.GetName())
	if err != nil {
		return nil, err
	}
	return &diff.DiffCommits{Commits: commits}, nil
}

func (d *DiffServer) UploadDiff(ctx context.Context, reply *diff.DiffCommits) (*diff.DiffConfirm, error) {
	logrus.Infof("Received message %s", reply.Name)
	err := d.saveDiff(reply.GetName(), reply.GetCommits())
	if err != nil {
		return nil, err
	}
	return &diff.DiffConfirm{}, nil
}

func (d *DiffServer) getDiff(name string) (string, error) {
	get, ok := d.db.Get(name)
	if !ok {
		return "", fmt.Errorf("no patch with name %s found", name)
	}
	return get.(string), nil
}

func (d *DiffServer) saveDiff(name string, commits string) error {
	d.db.Set(name, commits, cache.DefaultExpiration)
	return nil
}
