// package main

// import (
// 	"context"
// 	"log"
// 	"net"

// 	"google.golang.org/grpc"
// 	pb "github.com/ARunni/gRPC_sample/gen"
// )

// type server struct {}
// func (s *server) SayHello(ctx context.Context,req *pb.) (*HelloResponce,error) {
// 	return &HelloResponce{Message: "Hello,"+ req.Name},nil
// }

// func main()  {
// 	lis,err := net.Listen("tcp",":50051")
// 	if err != nil {
// 		log.Fatalf("failed to listen : %v",err)
// 	}
// 	s := grpc.NewServer()
// 	RegisterHelloServiceServer(s,&server{})
// 	if err := s.Serve(lis); err != nil {
// 		log.Fatalf("failed to serve : %v",err)
// 	}
// }

// server.go

package main

import (
	"context"
	"log"
	"net"

	pb "github.com/ARunni/gRPC_sample/gen"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello, " + req.Name}, nil
}

// Ensure that server struct implements mustEmbedUnimplementedHelloServiceServer
func (s *server) mustEmbedUnimplementedHelloServiceServer() {}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &server{})
	log.Println("Server started on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
