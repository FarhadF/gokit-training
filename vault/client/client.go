package client

import (
	"google.golang.org/grpc"
	"gokit-training/vault"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"gokit-training/vault/pb"
	"context"
)

func NewGRPCClient(conn *grpc.ClientConn) vault.Service {
	var hashEndpoint = grpctransport.NewClient(
		conn, "pb.Vault", "Hash",
		vault.EncodeGRPCHashRequest,
		vault.DecodeGRPCHashResponse,
		pb.HashResponse{},
	).Endpoint()
	var validateEndpoint = grpctransport.NewClient(
		conn, "pb.Vault", "Validate",
		vault.EncodeGRPCValidateRequest,
		vault.DecodeGRPCValidateResponse,
		pb.ValidateResponse{},
	).Endpoint()
	return vault.Endpoints{
		HashEndpoint:     hashEndpoint,
		ValidateEndpoint: validateEndpoint,
	}
}

func Hash(ctx context.Context, service vault.Service, password string)(string, error) {
	h, err := service.Hash(ctx, password)
	if err != nil {
		return "", err
	}
	return h, nil
}
func Validate(ctx context.Context, service vault.Service, password, hash string)(bool, error) {
	valid, err := service.Validate(ctx, password, hash)
	if err != nil {
		return false, err
	}
	if !valid {
		return false, nil
	}
	return true, nil
}
