package main
import (
    "context"
    "fmt"
    "log"
	"google.golang.org/grpc"
    pb "./pb"
)

func main() {

    conn, err := grpc.Dial("127.0.0.1:19003", grpc.WithInsecure())
    if err != nil {
        log.Fatal("client connection error:", err)
    }
    defer conn.Close()
    client := pb.NewCutterClient(conn)

    message := &pb.CutRequest{Data: "HELLO YOUTUBe"}
    res, err := client.CutImage(context.TODO(), message)
    fmt.Printf("result:%#v \n", res)
    fmt.Printf("error::%#v \n", err)
}
