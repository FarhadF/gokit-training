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
	return movies.Endpoints{
		getMoviesEndpoint,
	}
}


func main(){
	var (
		grpcAddr = flag.StringP("addr", "a",":8081","gRPC address")
	)

	//flag.StringVarP(&requestType, "requestType", "r", "word", "Should be word, sentence or paragraph")
	//flag.IntVarP(&min,"min", "m", 5, "minimum value")
	//flag.IntVarP(&max,"Max", "M", 10, "Maximum value")


	flag.Parse()
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	ctx := context.Background()
	conn, err := grpc.Dial(*grpcAddr, grpc.WithInsecure())
	if err != nil {
		logger.Fatal().Err(err).Msg("")
	}
	defer conn.Close()
	getMoviesService := NewGRPCClient(conn)


	callGetMovies(ctx, getMoviesService, logger)


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