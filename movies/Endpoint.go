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
	GetMovieByIdEndpoint endpoint.Endpoint
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

//model request and response
type getMovieByIdRequest struct {
	Id string `json:"id"`
}

type getMovieByIdResponse struct {
	Movie Movie `json:="movie"`
	Err string `json:="err"`
}

//Make actual endpoint per Method
func MakeGetMovieByIdEndpoint(svc Service)(endpoint.Endpoint) {
	return func(ctx context.Context, req interface{})(interface{}, error){
		r := req.(getMovieByIdRequest)
		movie, err := svc.GetMovieById(ctx, r.Id)
		if err != nil {
			return getMovieByIdResponse{movie, err.Error()}, nil
		}
		return getMovieByIdResponse{movie, ""}, nil
	}
}

// Wrapping Endpoints as a Service implementation.
// Will be used in gRPC client
func (e Endpoints) GetMovieById (ctx context.Context, id string)(Movie, error){
	req := getMovieByIdRequest{
		Id: id,
	}
	var movie Movie
	resp, err := e.GetMovieByIdEndpoint(ctx, req)
	if err != nil {
		return movie, err
	}
	getMovieByIdResp := resp.(getMovieByIdResponse)
	if getMovieByIdResp.Err != ""{
		return movie, errors.New(getMovieByIdResp.Err)
	}
	return getMovieByIdResp.Movie, nil
}

