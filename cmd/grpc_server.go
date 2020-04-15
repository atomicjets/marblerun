//go:generate protoc --proto_path=.. --go_out=plugins=grpc:./ --go_opt=paths=source_relative api/coordinator.proto

package main

import (
	"fmt"

	"edgeless.systems/mesh/coordinator/api"
)

// func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
// 	log.Printf("Received: %v", in.GetName())
// 	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
// }

// func main() {
// 	lis, err := net.Listen("tcp", port)
// 	if err != nil {
// 		log.Fatalf("failed to listen: %v", err)
// 	}
// 	s := grpc.NewServer()
// 	pb.RegisterGreeterServer(s, &server{})
// 	if err := s.Serve(lis); err != nil {
// 		log.Fatalf("failed to serve: %v", err)
// 	}
// }

func main() {
	r := api.HelloRequest{}
	fmt.Println(r)
}
