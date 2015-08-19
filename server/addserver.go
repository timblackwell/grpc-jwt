package main
import (
	"github.com/timblackwell/grpc-jwt/server/protobuf"
	"golang.org/x/net/context"
)

type AddServer struct {}

func (AddServer) Add (ctx context.Context, addreq *protobuf.AddRequest) (*protobuf.AddReply, error) {
	result := addreq.A + addreq.B
	return &protobuf.AddReply{ V: result}, nil
}