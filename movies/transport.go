package movies

import (
	"context"
	"gokit-training/movies/pb"
	"github.com/golang/protobuf/ptypes"
)

//Encode and Decode GetMovies Request and response
func EncodeGRPCGetMoviesRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil , nil
}

func DecodeGRPCGetMoviesRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, nil
}

// Encode and Decode GetMovies Response
func EncodeGRPCGetMoviesResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(getMoviesResponse)
	var movies []*pb.Movie
	for _, movie := range resp.Movies {
		createdOn, err  := ptypes.TimestampProto(movie.CreatedOn)
		if err != nil {
			//todo bring logger
			return nil, err
		}
		updatedOn, err := ptypes.TimestampProto(movie.UpdatedOn)
		if err != nil {
			//todo bring logger
			return nil, err
		}
		m := &pb.Movie{
			Id: movie.Id,
			Title: movie.Title,
			Director: movie.Director,
			Year: movie.Year,
			Userid: movie.Userid,
			Createdon: createdOn ,
			Updatedon: updatedOn,
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
		createdOn, err := ptypes.Timestamp(movie.Createdon)
		if err != nil {
			//todo log error
			return nil, err
		}
		updatedOn, err := ptypes.Timestamp(movie.Updatedon)
		if err != nil {
			//todo log error
			return nil, err
		}
		m := Movie{
			Id: movie.Id,
			Title: movie.Title,
			Director: movie.Director,
			Year: movie.Year,
			Userid: movie.Userid,
			CreatedOn: createdOn,
			UpdatedOn: updatedOn,

		}
		movies = append(movies, m)
	}
	return getMoviesResponse{
		Movies: movies,
		Err: resp.Err,
	}, nil
}

