package client

import (
	"google.golang.org/grpc"
	"gokit-training/movies"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"gokit-training/movies/pb"
	"context"
)

//create new client returns GetMovies Service
func NewGRPCClient(conn *grpc.ClientConn) movies.Service {
	var getMoviesEndpoint = grpctransport.NewClient(
		conn, "pb.Movies", "GetMovies",
		movies.EncodeGRPCGetMoviesRequest,
		movies.DecodeGRPCGetMoviesResponse,
		pb.GetMoviesResponse{},
	).Endpoint()
	var getMovieByIdEndpoint = grpctransport.NewClient(
		conn, "pb.Movies", "GetMovieById",
		movies.EncodeGRPCGetMovieByIdRequest,
		movies.DecodeGRPCGetMovieByIdResponse,
		pb.GetMovieByIdResponse{},
	).Endpoint()
	var newMovieEndpoint = grpctransport.NewClient(
		conn, "pb.Movies", "NewMovie",
		movies.EncodeGRPCNewMovieRequest,
		movies.DecodeGRPCNewMovieResponse,
		pb.NewMovieResponse{},
	).Endpoint()
	var deleteMovieEndpoint = grpctransport.NewClient(
		conn, "pb.Movies", "DeleteMovie",
		movies.EncodeGRPCDeleteMovieRequest,
		movies.DecodeGRPCDeleteMovieResponse,
		pb.DeleteMovieResponse{},
	).Endpoint()
	var updateMovieEndpoint = grpctransport.NewClient(
		conn, "pb.Movies", "UpdateMovie",
		movies.EncodeGRPCUpdateMovieRequest,
		movies.DecodeGRPCUpdateMovieResponse,
		pb.UpdateMovieResponse{},
	).Endpoint()
	return movies.Endpoints{
		GetMoviesEndpoint: getMoviesEndpoint,
		GetMovieByIdEndpoint: getMovieByIdEndpoint,
		NewMovieEndpoint: newMovieEndpoint,
		DeleteMovieEndpoint: deleteMovieEndpoint,
		UpdateMovieEndpoint: updateMovieEndpoint,
	}
}

//callService helpers
func GetMovies(ctx context.Context, service movies.Service) ([]movies.Movie, error) {
	return service.GetMovies(ctx)
}

func GetMovieById(ctx context.Context, id string, service movies.Service) (movies.Movie, error) {
	return service.GetMovieById(ctx, id)
}

func NewMovie(ctx context.Context, title string, director []string, year string, userId string, service movies.Service) (string, error) {
	return service.NewMovie(ctx, title, director, year, userId)
}

func DeleteMovie(ctx context.Context, id string, service movies.Service) error {
	return service.DeleteMovie(ctx, id)
}

func UpdateMovie(ctx context.Context, id string, title string, director []string, year string, userId string,
	service movies.Service) error {
	err := service.UpdateMovie(ctx, id, title, director, year, userId)
	if err != nil {
		return err
	}
	return nil
}