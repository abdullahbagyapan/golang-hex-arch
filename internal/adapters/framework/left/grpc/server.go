package grpc

import (
	"google.golang.org/grpc"
	"hexarch/internal/adapters/framework/left/grpc/pb"
	"hexarch/internal/ports"
	"log"
	"net"
)

type Adapter struct {
	api ports.APIPort
	pb.UnimplementedArithmeticServiceServer
}

func (grpcA Adapter) Run() {

	listen, err := net.Listen("tcp", ":9000")

	if err != nil {
		log.Fatalf("failed to listening %v", err)
	}

	arithmeticServiceServer := grpcA

	grpcServer := grpc.NewServer()

	pb.RegisterArithmeticServiceServer(grpcServer, arithmeticServiceServer)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serving grpc server %v", err)
	}

}
