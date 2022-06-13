package service

import (
	"context"
	pb "hello-grpc/proto/grpc_hello_service"
)

type HelloServiceImpl struct {
	pb.UnimplementedHelloServiceServer
}

func (p *HelloServiceImpl) Hello(ctx context.Context,args *pb.String)(*pb.String, error)  {
	reply := &pb.String{Value: "hello"+args.GetValue()}
	return reply,nil
}



