package lorem

import (
	"github.com/go-kit/kit/endpoint"
	"context"
	"errors"
)

//model request and response
type loremRequest struct {
	RequestType string `json:"requesttype"`
	Min         int    `json:"min"`
	Max         int    `json:"max"`
}

//model request and response
type loremResponse struct {
	Message string `json:"message,omitempty"`
	Err     string `json:"err,omitempty"`
}

//Endpoints Wrapper
type Endpoints struct {
	LoremEndpoint endpoint.Endpoint
}

//Make actual endpoint per Method
func MakeLoremEndpoint(svc Service) (endpoint.Endpoint) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(loremRequest)
		str, err := svc.Lorem(ctx, req.RequestType, int(req.Min), int(req.Max))
		if err != nil {
			return loremResponse{"", err.Error()}, nil
		}
		return loremResponse{str, ""}, nil
	}
}

// Wrapping Endpoints as a Service implementation.
// Will be used in gRPC client
func (e Endpoints) Lorem(ctx context.Context, requestType string, min int, max int) (string, error) {
	req := loremRequest{
		requestType,
		min,
		max,
	}
	resp, err := e.LoremEndpoint(ctx, req)
	if err != nil {
		return "", err
	}
	loremResp := resp.(loremResponse)
	if loremResp.Err != "" {
		return "", errors.New(loremResp.Err)
	}
	return loremResp.Message, nil
}
