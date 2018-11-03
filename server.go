
package main
import (
    "fmt"
    "context"
    "log"
    "net"
    "bytes"
    "image"
    "image/png"
    _"os"
    "encoding/base64"
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

type MyService struct {}
func (s *MyService) CutImage(ctx context.Context, message *pb.CutRequest) (*pb.CutReply, error) {
    fmt.Println("receviced message:")

    data, _ := base64.StdEncoding.DecodeString(message.Data);
    result, _ := cut(data, int(message.Left), int(message.Top), int(message.Width), int(message.Height));

    return &pb.CutReply{ Result: result }, nil
}

func cut (imageData []byte, left int, top int, width int, height int) (string, error) {

  r := bytes.NewReader(imageData)
  decodedImage, err := png.Decode(r)
  if err != nil {
    fmt.Println("failed to decoee Image")
    return "", err;
  }

  fmt.Printf("left:%d, top:%d, width:%d, height:%d", left, top, width, height)
  outRect := image.Rect(left, top, left + width, top + height)

  outImage := decodedImage.(interface {
            SubImage(r image.Rectangle) image.Image
                  }).SubImage(outRect)
  /**
   * Output into file.
   */
  // outFile, err := os.Create("out.png")
  // if err != nil {
  //   fmt.Println("failed to create output file")
  // }
  // png.Encode(outFile,  outImage)
  // defer outFile.Close()

  var buff bytes.Buffer
  png.Encode(&buff, outImage)
  return base64.StdEncoding.EncodeToString(buff.Bytes()), nil
}
