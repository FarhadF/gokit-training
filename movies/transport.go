package movies

import (
	"context"
	"gokit-training/movies/pb"
)

//Encode and Decode Lorem Request and response
func EncodeGRPCGetMoviesRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil , nil
}

func DecodeGRPCGetMoviesRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, nil
}

// Encode and Decode Lorem Response
func EncodeGRPCGetMoviesResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(getMoviesResponse)
	var movies []*pb.Movie
	for _, movie := range resp.Movies {
		m := &pb.Movie{
			Id: movie.Id,
			Title: movie.Title,
			Director: movie.Director,
			Year: movie.Year,
			Userid: movie.Userid,
		}
		movies = append(movies, m)
	}
	return &pb.GetMoviesResponse{
		Movies: movies,
		Err: resp.Err,
	}, nil
}

func DecodeGRPCGetMoviesResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(*pb.GetMoviesResponse)
	var movies []Movie
	for _, movie := range resp.Movies {
		m := Movie{
			Id: movie.Id,
			Title: movie.Title,
			Director: movie.Director,
			Year: movie.Year,
			Userid: movie.Userid,
		}
		movies = append(movies, m)
	}
	return getMoviesResponse{
		Movies: movies,
		Err: resp.Err,
	}, nil
}

