package movies

import (
	"context"
	"gokit-training/movies/pb"
	grpctransport "github.com/go-kit/kit/transport/grpc"
)

//grpcServer Wrapper
type grpcServer struct {
	getMovies    grpctransport.Handler
	getMovieById grpctransport.Handler
	newMovie     grpctransport.Handler
	deleteMovie  grpctransport.Handler
}

// implement getMovies server Interface in movies.pb.go
func (s *grpcServer) GetMovies(ctx context.Context, r *pb.Empty) (*pb.GetMoviesResponse, error) {
	_, resp, err := s.getMovies.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.GetMoviesResponse), nil
}

// implement getMovieById server Interface in movies.pb.go
func (s *grpcServer) GetMovieById(ctx context.Context, r *pb.GetMovieByIdRequest) (*pb.GetMovieByIdResponse, error) {
	_, resp, err := s.getMovieById.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.GetMovieByIdResponse), nil
}

// implement NewMovie server Interface in movies.pb.go
func (s *grpcServer) NewMovie(ctx context.Context, r *pb.NewMovieRequest) (*pb.NewMovieResponse, error) {
	_, resp, err := s.newMovie.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.NewMovieResponse), nil
}

// implement DeleteMovie server Interface in movies.pb.go
func (s *grpcServer) DeleteMovie(ctx context.Context, r *pb.DeleteMovieRequest) (*pb.DeleteMovieResponse, error) {
	_, resp, err := s.deleteMovie.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.DeleteMovieResponse), nil
}

// create new grpc server
func NewGRPCServer(ctx context.Context, endpoint Endpoints) pb.MoviesServer {
	return &grpcServer{
		getMovies: grpctransport.NewServer(
			endpoint.GetMoviesEndpoint,
			DecodeGRPCGetMoviesRequest,
			EncodeGRPCGetMoviesResponse,
		),
		getMovieById: grpctransport.NewServer(
			endpoint.GetMovieByIdEndpoint,
			DecodeGRPCGetMovieByIdRequest,
			EncodeGRPCGetMovieByIdResponse,
		),
		newMovie: grpctransport.NewServer(
			endpoint.NewMovieEndpoint,
			DecodeGRPCNewMovieRequest,
			EncodeGRPCNewMovieResponse,
		),
		deleteMovie: grpctransport.NewServer(
			endpoint.DeleteMovieEndpoint,
			DecodeGRPCDeleteMovieRequest,
			EncodeGRPCDeleteMovieResponse,
		),
	}
}
