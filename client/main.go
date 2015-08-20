package main
import (
	"fmt"
	"google.golang.org/grpc"
	"github.com/timblackwell/grpc-jwt/client/protobuf"
	"golang.org/x/net/context"
	"flag"
	"os"
	"google.golang.org/grpc/metadata"
)


func main() {
	fs := flag.NewFlagSet("", flag.ExitOnError)
	var (
		grpcAddr = fs.String("grpc.addr", ":8002", "Address for gRPC server")
		accessToken = fs.String("grpc.token", "test", "JWT used to gRPC calls")
	)
	flag.Usage = fs.Usage // only show our flags
	fs.Parse(os.Args[1:])

	conn, err := grpc.Dial(*grpcAddr)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	// create client and call
	client := protobuf.NewAddClient(conn)

	// create context with JWT
	md := metadata.Pairs("token", *accessToken)
	ctx := context.Background()
	ctx = metadata.NewContext(ctx, md)

	var header, trailer metadata.MD
	result, err := client.Add(ctx, &protobuf.AddRequest{A: 1, B: 2}, grpc.Header(&header), grpc.Trailer(&trailer))

	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	fmt.Printf("Result: %v\n", result.V)
	conn.Close()
}