//go:generate protoc -I ../proto --go_out=plugins=grpc:../proto ../proto/hello.proto
package main

import (
  "log"
  "os"

  "golang.org/x/net/context"
  "google.golang.org/grpc"
  pb "github.com/furikuri/hello-grpc/go/proto"
)

const (
  address     = "localhost:50051"
  defaultName = "world"
)

func main() {
  conn, err := grpc.Dial(address, grpc.WithInsecure())
  if err != nil {
    log.Fatalf("did not connect: %v", err)
  }
  defer conn.Close()
  c := pb.NewGreeterClient(conn)

  name := defaultName
  if len(os.Args) > 1 {
    name = os.Args[1]
  }
  r, err := c.SayHello(context.Background(), &pb.Request{Name: name})
  if err != nil {
    log.Fatalf("could not greet: %v", err)
  }
  log.Printf("Greeting: %s", r.Message)
}
