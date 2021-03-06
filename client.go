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

    message := &pb.CutRequest{
        Data: "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",
        Left: 0,
        Top: 32,
        Width: 32,
        Height: 49,
    }
    res, err := client.CutImage(context.TODO(), message)
    fmt.Printf("result:%#v \n", res)
    fmt.Printf("error::%#v \n", err)
}
