package main
import (
	"golang.org/x/oauth2"
	"google.golang.org/grpc/credentials/oauth"
	"fmt"
	"google.golang.org/grpc"
	"github.com/timblackwell/grpc-jwt/client/protobuf"
	"golang.org/x/net/context"
	"flag"
	"os"
)


func main() {
	// over
	fs := flag.NewFlagSet("", flag.ExitOnError)
	var (
		grpcAddr = fs.String("grpc.addr", ":8002", "Address for gRPC server")
		accessToken = fs.String("grpc.token", "", "JWT used to gRPC calls")
	)
	flag.Usage = fs.Usage // only show our flags
	fs.Parse(os.Args[1:])


	// obtaining the scope is not in this projects scope
	token := oauth2.Token{
		AccessToken: *accessToken,
	}

	// create credentials with token, and use them for all calls to server
	cred := oauth.NewOauthAccess(&token)
	conn, err := grpc.Dial(*grpcAddr, grpc.WithPerRPCCredentials(cred))
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	// create client and call
	client := protobuf.NewAddClient(conn)
	result, err := client.Add(context.Background(), &protobuf.AddRequest{A: 1, B: 2})
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	fmt.Printf("Result: %v\n", result.V)
	conn.Close()
}