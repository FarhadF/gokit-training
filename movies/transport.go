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

//encode GetMoviesByIdRequest
func EncodeGRPCGetMovieByIdRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(getMovieByIdRequest)
	return &pb.GetMovieByIdRequest{
		Id: req.Id,
	} , nil
}

//decode GetMovieByIdRequest
func DecodeGRPCGetMovieByIdRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.GetMovieByIdRequest)
	return getMovieByIdRequest{
		Id: req.Id,
	}, nil
}

// encode GetMovieById Response
func EncodeGRPCGetMovieByIdResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(getMovieByIdResponse)
		createdOn, err  := ptypes.TimestampProto(resp.Movie.CreatedOn)
		if err != nil {
			//todo bring logger
			return nil, err
		}
		updatedOn, err := ptypes.TimestampProto(resp.Movie.UpdatedOn)
		if err != nil {
			//todo bring logger
			return nil, err
		}
		m := &pb.Movie{
			Id: resp.Movie.Id,
			Title: resp.Movie.Title,
			Director: resp.Movie.Director,
			Year: resp.Movie.Year,
			Userid: resp.Movie.Userid,
			Createdon: createdOn ,
			Updatedon: updatedOn,
		}

	return &pb.GetMovieByIdResponse{
		Movie: m,
		Err: resp.Err,
	}, nil
}

// decode GetMovieById Response
func DecodeGRPCGetMovieByIdResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(*pb.GetMovieByIdResponse)
		createdOn, err := ptypes.Timestamp(resp.Movie.Createdon)
		if err != nil {
			//todo log error
			return nil, err
		}
		updatedOn, err := ptypes.Timestamp(resp.Movie.Updatedon)
		if err != nil {
			//todo log error
			return nil, err
		}
		m := Movie{
			Id: resp.Movie.Id,
			Title: resp.Movie.Title,
			Director: resp.Movie.Director,
			Year: resp.Movie.Year,
			Userid: resp.Movie.Userid,
			CreatedOn: createdOn,
			UpdatedOn: updatedOn,
		}

	return getMovieByIdResponse{
		Movie: m,
		Err: resp.Err,
	}, nil
}

//encode NewMovieRequest
func EncodeGRPCNewMovieRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(newMovieRequest)
	return &pb.NewMovieRequest{
		Title: req.Title,
		Director: req.Director,
		Year: req.Year,
		Userid: req.Userid,
	} , nil
}

//decode NewMovieRequest
func DecodeGRPCNewMovieRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.NewMovieRequest)
	return newMovieRequest{
		Title: req.Title,
		Director: req.Director,
		Year: req.Year,
		Userid: req.Userid,
	}, nil
}

// Encode and Decode NewMovieResponse
func EncodeGRPCNewMovieResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(newMovieResponse)
	return &pb.NewMovieResponse{
		Id: resp.Id,
		Err: resp.Err,
	}, nil
}

func DecodeGRPCNewMovieResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(*pb.NewMovieResponse)
	return newMovieResponse{
		Id: resp.Id,
		Err: resp.Err,
	}, nil
}