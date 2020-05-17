package watch

import (
	"context"
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

func RequestDiffs(ctx context.Context, c diff.DiffClient, r *diff.DiffRequest, password string) (*diff.DiffCommits, error) {
	resp, err := c.SendDiff(ctx, r)
	if err != nil {
		return nil, err
	}
	if password != "" {
		resp.Commits, err = decrypt(resp.Commits, password, resp.Salt)
		if err != nil {
			return nil, err
		}
	}
	return resp, nil
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
