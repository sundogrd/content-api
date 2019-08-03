package comment

import (
	"github.com/sirupsen/logrus"
	"github.com/sundogrd/content-api/grpc_gen/comment"
	grpcUtils "github.com/sundogrd/gopkg/grpc"
	"google.golang.org/grpc"
)

const (
	//address     = "localhost:50052"
	address     = "sundog.comment"
	defaultName = "sundog.comment"
)

// NewGrpcCommentClient
func NewGrpcCommentClient() (comment.CommentServiceClient, *grpc.ClientConn, error) {
	r, err := grpcUtils.NewGrpcResolover()
	if err != nil {
		return nil, nil, err
	}
	b := grpc.RoundRobin(r)
	// Set up a connection to the server.
	// WithBlock https://gocn.vip/question/931  wtf
	conn, err := grpc.Dial(address, grpc.WithBalancer(b), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		logrus.Fatalf("did not connect: %v", err)
		return nil, nil, err
	}
	c := comment.NewCommentServiceClient(conn)
	return c, conn, nil
}
