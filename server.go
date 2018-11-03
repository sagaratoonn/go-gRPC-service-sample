
package main
import (
    "fmt"
    "context"
    "log"
    "net"
	"google.golang.org/grpc"
	pb "./pb"
)

func main() {
    listenPort, err := net.Listen("tcp", ":19003")

    if err != nil {
        log.Fatalln(err)
    }
    server := grpc.NewServer()
    service := &MyService{}
    pb.RegisterCutterServer(server, service)
    server.Serve(listenPort)
}

type MyService struct {
}

func (s *MyService) CutImage(ctx context.Context, message *pb.CutRequest) (*pb.CutReply, error) {
    fmt.Println("receviced message:")
    fmt.Println(message)
        return &pb.CutReply{ Number: 3 }, nil
}
