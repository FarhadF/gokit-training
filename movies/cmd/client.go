package main

import (
	"google.golang.org/grpc"
	"gokit-training/movies"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"gokit-training/movies/pb"
	flag "github.com/spf13/pflag"
	"context"
	"github.com/rs/zerolog"
	"os"
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
	return movies.Endpoints{
		getMoviesEndpoint,
		getMovieByIdEndpoint,
	}
}


func main(){
	var (
		grpcAddr string
		movieId string
	)
	flag.StringVarP(&grpcAddr,"addr", "a",":8081","gRPC address")
	flag.StringVarP(&movieId, "id", "i","","movieId")
	//flag.StringVarP(&requestType, "requestType", "r", "word", "Should be word, sentence or paragraph")
	//flag.IntVarP(&min,"min", "m", 5, "minimum value")
	//flag.IntVarP(&max,"Max", "M", 10, "Maximum value")


	flag.Parse()
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	ctx := context.Background()
	conn, err := grpc.Dial(grpcAddr, grpc.WithInsecure())
	if err != nil {
		logger.Fatal().Err(err).Msg("")
	}
	defer conn.Close()
	getMoviesService := NewGRPCClient(conn)


	callGetMovies(ctx, getMoviesService, logger)
	if movieId != "" {
		callGetMovieById(ctx, movieId, getMoviesService, logger)

	}


}

//callService helper
func callGetMovies(ctx context.Context, service movies.Service, logger zerolog.Logger) {
	mesg, err := service.GetMovies(ctx)
	if err != nil {
		logger.Fatal().Err(err).Msg("")
	}
	//j, err := json.Marshal(mesg)
	logger.Info().Interface("movie", mesg).Msg("")
}

func callGetMovieById(ctx context.Context, id string, service movies.Service, logger zerolog.Logger) {
	mesg, err := service.GetMovieById(ctx, id)
	if err != nil {
		logger.Fatal().Err(err).Msg("")
	}
	logger.Info().Interface("movie", mesg).Msg("")
}