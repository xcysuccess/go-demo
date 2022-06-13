package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	pb "hello-grpc/proto/grpc_hello_service"
	"hello-grpc/service"
)

func main()  {
	grpcServer := grpc.NewServer()
	pb.RegisterHelloServiceServer(grpcServer, &service.HelloServiceImpl{})

	lis, err := net.Listen("tcp","localhost:1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)
}
