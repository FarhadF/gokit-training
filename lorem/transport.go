package lorem

import (
	"context"
	"gokit-training/lorem/pb"
)

//Encode and Decode Lorem Request and response
func EncodeGRPCLoremRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(loremRequest)
	return &pb.LoremRequest{
		RequestType: req.RequestType,
		Max: int32(req.Max),
		Min: int32(req.Min),
	} , nil
}

func DecodeGRPCLoremRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.LoremRequest)
	return loremRequest{
		RequestType: req.RequestType,
		Max: int(req.Max),
		Min: int(req.Min),
	}, nil
}

// Encode and Decode Lorem Response
func EncodeGRPCLoremResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(loremResponse)
	return &pb.LoremResponse{
		Message: resp.Message,
		Err: resp.Err,
	}, nil
}

func DecodeGRPCLoremResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(*pb.LoremResponse)
	return loremResponse{
		Message: resp.Message,
		Err: resp.Err,
	}, nil
}
