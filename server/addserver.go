package main
import (
	"github.com/timblackwell/grpc-jwt/server/protobuf"
	"golang.org/x/net/context"
	"fmt"
)

type AddServer struct {}

func (AddServer) Add (ctx context.Context, addreq *protobuf.AddRequest) (*protobuf.AddReply, error) {
	result := addreq.A + addreq.B
	fmt.Printf("authorization: %v", ctx.Value("authorization"))
	//todo check token
	return &protobuf.AddReply{ V: result}, nil
}