package comment

import (
	"fmt"
	"log"

	"github.com/sundogrd/content-api/grpc_gen/comment"
	"google.golang.org/grpc"
)

const (
	address = "localhost:8999"
	// address     = "sundog.comment"
	defaultName = "sundog.comment"
)

// NewGrpcCommentClient
func NewGrpcCommentClient() (comment.CommentServiceClient, *grpc.ClientConn, error) {
	// r, err := grpcUtils.NewGrpcResolover()
	// if err != nil {
	// 	return nil, nil, err
	// }
	// b := grpc.RoundRobin(r)
	// // Set up a connection to the server.
	// // WithBlock https://gocn.vip/question/931  wtf
	// conn, err := grpc.Dial(address, grpc.WithBalancer(b), grpc.WithInsecure(), grpc.WithBlock())
	// if err != nil {
	// 	log.Fatalf("did not connect: %v", err)
	// 	return nil, nil, err
	// }

	fmt.Println("客户端开始运行.....")
	conn, err := grpc.Dial("127.0.0.1:8999", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// name := defaultName
	// if len(os.Args) > 1 {
	// 	name = os.Args[1]
	// }

	// defer conn.Close()

	c := comment.NewCommentServiceClient(conn)

	return c, conn, nil
}
