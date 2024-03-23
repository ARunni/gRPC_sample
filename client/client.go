// package main

// import (
// 	"context"
// 	"log"

// 	"github.com/ARunni/gRPC_sample/gen"
// 	"google.golang.org/grpc"
// )

// func main() {
// 	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
// 	if err != nil {
// 		log.Fatalf("did not connect: %v", err)
// 	}
// 	defer conn.Close()
// 	c := NewHelloServiceClient(conn)

// 	name := "John"
// 	res, err := c.SayHello(context.Background(), &HelloRequest{Name: name})
// 	if err != nil {
// 		log.Fatalf("could not greet: %v", err)
// 	}
// 	log.Printf("Response: %s", res.Message)
// }


// client.go

package main

import (
    "context"
    "log"

    "google.golang.org/grpc"
    pb "github.com/ARunni/gRPC_sample/gen"
)

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := pb.NewHelloServiceClient(conn)

    name := "John"
    res, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
    if err != nil {
        log.Fatalf("could not greet: %v", err)
    }
    log.Printf("Response: %s", res.Message)
}
