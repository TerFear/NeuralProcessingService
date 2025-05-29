package server

import (
	"context"
	ssov1 "gRPC_get_message/protos/gen/go/sso"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Processing interface {
	Process(ctx context.Context, userID string, data string, topic string) (string, error)
}

// чтоб управлять функциями
type ServerAPI struct {
	ssov1.UnimplementedNeuralProcessingServiceServer
	Processer Processing
}

func Register(gRPC *grpc.Server, processing Processing) {
	ssov1.RegisterNeuralProcessingServiceServer(gRPC, &ServerAPI{Processer: processing})
}

func (s *ServerAPI) SubmitJob(ctx context.Context, req *ssov1.SubmitJobRequest) (*ssov1.SubmitJobResponse, error) {

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request is nil")
	}

	userID := req.GetUserId()
	data := req.GetData()
	topic := req.GetTopic()

	res, err := s.Processer.Process(ctx, userID, data, topic)

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	return &ssov1.SubmitJobResponse{
		Data: res,
	}, status.Error(codes.OK, "")
}
