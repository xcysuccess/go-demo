package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "hello-grpc/proto/grpc_hello_service"
	"log"
)

func main()  {
	conn ,err := grpc.Dial("localhost:1234",grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewHelloServiceClient(conn)
	reply,err := client.Hello(context.Background(), &pb.String{Value:"hello"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetValue())
}
