package server

import (
	"context"
	ssov1 "gRPC_get_message/protos/gen/go/sso"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Processing interface {
	Process(ctx context.Context, data string, dataType string) (string, error)
}

// чтоб управлять функциями
type ServerAPI struct {
	ssov1.UnimplementedNeuralProcessingServiceServer
	Processer Processing
}

func Register(gRPC *grpc.Server) {
	ssov1.RegisterNeuralProcessingServiceServer(gRPC, &ServerAPI{})
}

func (s *ServerAPI) SubmitJob(ctx context.Context, req *ssov1.SubmitJobRequest) (*ssov1.SubmitJobResponse, error) {

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request is nil")
	}

	process, err := s.Processer.Process(ctx, req.GetInputData(), req.GetDataType())

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to process data: %v", err)
	}

	return &ssov1.SubmitJobResponse{
		SessionId: "abc123",
		Status:    "success",
		Output:    process,
	}, nil
}
