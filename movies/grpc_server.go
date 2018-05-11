package movies

import (
	"context"
	"gokit-training/movies/pb"
	grpctransport "github.com/go-kit/kit/transport/grpc"
)

//grpcServer Wrapper
type grpcServer struct {
	getMovies grpctransport.Handler
}

// implement getMovies server Interface in lorem.pb.go
func (s *grpcServer) GetMovies(ctx context.Context, r *pb.Empty)(*pb.GetMoviesResponse, error) {
	_, resp, err := s.getMovies.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.GetMoviesResponse), nil
}

// create new grpc server
func NewGRPCServer(ctx context.Context, endpoint Endpoints) pb.MoviesServer {
	return &grpcServer{
		grpctransport.NewServer(
			endpoint.GetMoviesEndpoint,
			DecodeGRPCGetMoviesRequest,
			EncodeGRPCGetMoviesResponse,
		),
	}
}

