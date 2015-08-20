package main
import (
	"net"
	"google.golang.org/grpc"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"github.com/timblackwell/grpc-jwt/server/protobuf"
	"flag"
)


func main() {
	fs := flag.NewFlagSet("", flag.ExitOnError)
	var (
		grpcAddr = fs.String("grpc.addr", ":8002", "Address for gRPC server")
	)
	flag.Usage = fs.Usage // only show our flags
	fs.Parse(os.Args[1:])

	errc := make(chan error)
	go func() {
		errc <- interrupt()
	}()


	go func() {
		ln, err := net.Listen("tcp", *grpcAddr)
		if err != nil {
			errc <- err
			return
		}

		s := grpc.NewServer()
		protobuf.RegisterAddServer(s, AddServer{})
		errc <- s.Serve(ln)
	}()

	fmt.Errorf("%s", <-errc)
}

func interrupt() error {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	return fmt.Errorf("%s", <-c)
}