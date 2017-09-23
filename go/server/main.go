//go:generate protoc -I ../proto --go_out=plugins=grpc:../proto ../proto/hello.proto

package main

import (
  "log"
  "net"

  "golang.org/x/net/context"
  "google.golang.org/grpc"
  pb "github.com/furikuri/hello-grpc/go/proto"
  "google.golang.org/grpc/reflection"
)

const (
  port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.Request) (*pb.Reply, error) {
  return &pb.Reply{Message: "Hello " + in.Name + " from Go Server"}, nil
}

func main() {
  lis, err := net.Listen("tcp", port)
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }
  s := grpc.NewServer()
  pb.RegisterGreeterServer(s, &server{})
  // Register reflection service on gRPC server.
  reflection.Register(s)
  if err := s.Serve(lis); err != nil {
    log.Fatalf("failed to serve: %v", err)
  }
}
