package rpc

import (
	"context"

	"github.com/yasir2000/my-hexagonal-go/adaptors/framework/left/grpc/pb"
	"github.com/yasir2000/my-hexagonal-go/internal/adaptors/framework/left/grpc/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (grpca Adaptor) GetAddition(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, err) {
	var err error
	ans := &pb.Answer{}
	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}
	answer, err := grpca.api.GetAddition(req.A, req.B)
	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected error")
	}

	ans = &pb.Answer{
		Value: answer,
	}
	return ans, nil
}

func (grpca Adaptor) GetDivision(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, err) {
	var err error
	ans := &pb.Answer{}
	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}
	answer, err := grpca.api.GetDivision(req.A, req.B)
	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected error")
	}

	ans = &pb.Answer{
		Value: answer,
	}
	return ans, nil
}

func (grpca Adaptor) GetSubtraction(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, err) {
	var err error
	ans := &pb.Answer{}
	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}
	answer, err := grpca.api.GetSubtraction(req.A, req.B)
	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected error")
	}

	ans = &pb.Answer{
		Value: answer,
	}
	return ans, nil
}

func (grpca Adaptor) GetMultiplication(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, err) {
	var err error
	ans := &pb.Answer{}
	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}
	answer, err := grpca.api.GetMultiplication(req.A, req.B)
	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected error")
	}

	ans = &pb.Answer{
		Value: answer,
	}
	return ans, nil
}
