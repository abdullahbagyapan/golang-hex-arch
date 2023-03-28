package grpc

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hexarch/internal/adapters/framework/left/grpc/pb"
	"hexarch/internal/ports"
)

func NewGRPCAdapter(api ports.APIPort) *Adapter {
	return &Adapter{api: api}
}

func (grpcA Adapter) GetAddition(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	answer, err := grpcA.api.GetAddition(req.A, req.B)

	if err != nil {
		return nil, status.Error(codes.Internal, "unexcepted error")
	}

	return &pb.Answer{Value: answer}, nil

}

func (grpcA Adapter) GetMultiplication(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	answer, err := grpcA.api.GetMultiplication(req.A, req.B)

	if err != nil {
		return nil, status.Error(codes.Internal, "unexcepted error")
	}

	return &pb.Answer{Value: answer}, nil
}

func (grpcA Adapter) GetDivision(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	answer, err := grpcA.api.GetDivision(req.A, req.B)

	if err != nil {
		return nil, status.Error(codes.Internal, "unexcepted error")
	}

	return &pb.Answer{Value: answer}, nil
}

func (grpcA Adapter) GetSubtraction(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	answer, err := grpcA.api.GetSubtraction(req.A, req.B)

	if err != nil {
		return nil, status.Error(codes.Internal, "unexcepted error")
	}

	return &pb.Answer{Value: answer}, nil
}
