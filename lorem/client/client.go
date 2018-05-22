package client

import (
	"google.golang.org/grpc"
	"gokit-training/lorem"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"gokit-training/lorem/pb"
	"context"
)

//create new client returns Lorem Service
func NewGRPCClient(conn *grpc.ClientConn) lorem.Service {
	var loremEndpoint = grpctransport.NewClient(
		conn, "pb.Lorem", "Lorem",
		lorem.EncodeGRPCLoremRequest,
		lorem.DecodeGRPCLoremResponse,
		pb.LoremResponse{},
	).Endpoint()
	return lorem.Endpoints{
		loremEndpoint,
	}
}

//callService helper
func Lorem(ctx context.Context, service lorem.Service, requestType string, min int, max int) (string, error){
	mesg, err := service.Lorem(ctx, requestType, min, max)
	if err != nil {
		return "", err
	}
	return mesg, nil
}
