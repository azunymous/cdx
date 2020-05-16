package watch

import (
	"crypto/tls"
	"github.com/azunymous/cdx/watch/diff"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func NewClient(target string, insecure bool) (diff.DiffClient, func(), error) {
	conn, err := grpc.Dial(target, createDialOptions(insecure)...)
	if err != nil {
		return nil, nil, err
	}

	return diff.NewDiffClient(conn),
		func() {
			defer conn.Close()
		}, nil
}

func createDialOptions(insecure bool) []grpc.DialOption {
	var opts []grpc.DialOption
	if insecure {
		opts = append(opts, grpc.WithInsecure())
	} else {
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))
	}
	return opts
}
