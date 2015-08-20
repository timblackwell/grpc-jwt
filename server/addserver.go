package main
import (
	"github.com/timblackwell/grpc-jwt/server/protobuf"
	"golang.org/x/net/context"
	"fmt"
	"google.golang.org/grpc/metadata"
)

type AddServer struct {}

func (AddServer) Add (ctx context.Context, addreq *protobuf.AddRequest) (*protobuf.AddReply, error) {
	result := addreq.A + addreq.B

	md, _ := metadata.FromContext(ctx)
	token :=  md["token"]

	fmt.Printf("token: %v", token[0])
	//todo check token
	return &protobuf.AddReply{ V: result}, nil
}