package movies

import (
	"github.com/go-kit/kit/endpoint"
	"context"
	"errors"
)



//model request and response
type getMoviesResponse struct {
	Movies []Movie 	`json:"movies,omitempty"`
	Err 	string 	`json:"err,omitempty"`
}

//Endpoints Wrapper
type Endpoints struct {
	GetMoviesEndpoint endpoint.Endpoint
}

//Make actual endpoint per Method
func MakeGetMoviesEndpoint(svc Service)(endpoint.Endpoint) {
	return func(ctx context.Context, req interface{})(interface{}, error){
		movies, err := svc.GetMovies(ctx)
		if err != nil {
			return getMoviesResponse{nil, err.Error()}, nil
		}
		return getMoviesResponse{movies, ""}, nil
	}
}

// Wrapping Endpoints as a Service implementation.
// Will be used in gRPC client
func (e Endpoints) GetMovies (ctx context.Context)([]Movie, error){

	resp, err := e.GetMoviesEndpoint(ctx, nil)
	if err != nil {
		return nil, err
	}
	getMoviesResp:= resp.(getMoviesResponse)
	if getMoviesResp.Err != ""{
		return nil, errors.New(getMoviesResp.Err)
	}
	return getMoviesResp.Movies, nil
}



