package main
import (
	"fmt"
	"google.golang.org/grpc"
	"github.com/timblackwell/grpc-jwt/client/protobuf"
	"golang.org/x/net/context"
	"flag"
	"os"
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

	// create context with JWT
	ctx := context.Background()
	ctx = context.WithValue(ctx, "authorization", accessToken)

	// create client and call
	client := protobuf.NewAddClient(conn)
	result, err := client.Add(ctx, &protobuf.AddRequest{A: 1, B: 2})
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	fmt.Printf("Result: %v\n", result.V)
	conn.Close()
}