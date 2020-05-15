package watch

import (
	"github.com/azunymous/cdx/watch/diff"
	"google.golang.org/grpc"
)

func NewClient() (diff.DiffClient, func(), error) {
	target := ":8080"
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}

	return diff.NewDiffClient(conn),
		func() {
			defer conn.Close()
		}, nil
}
