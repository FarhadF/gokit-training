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

// create new grpc server
func NewGRPCServer(ctx context.Context, endpoint Endpoints) pb.MoviesServer {
	return &grpcServer{
		grpctransport.NewServer(
			endpoint.GetMoviesEndpoint,
			DecodeGRPCGetMoviesRequest,
			EncodeGRPCGetMoviesResponse,
		),
		grpctransport.NewServer(
			endpoint.GetMovieByIdEndpoint,
			DecodeGRPCGetMovieByIdRequest,
			EncodeGRPCGetMovieByIdResponse,
		),
		grpctransport.NewServer(
			endpoint.NewMovieEndpoint,
			DecodeGRPCNewMovieRequest,
			EncodeGRPCNewMovieResponse,
		),
	}
}
