package main

import (
	"google.golang.org/grpc"
	flag "github.com/spf13/pflag"
	"context"
	"github.com/rs/zerolog"
	"os"
	"strings"
	"gokit-training/movies/client"
)

func main() {
	var (
		grpcAddr string
		movieId  string
		newMovie bool
		title    string
		director string
		year     string
		userId   string
		deleteMovie bool
		updateMovie bool
	)
	flag.StringVarP(&grpcAddr, "addr", "a", ":8081", "gRPC address")
	flag.StringVarP(&movieId, "id", "i", "", "movieId")
	flag.StringVarP(&title, "title", "t", "", "title")
	flag.StringVarP(&director, "director", "d", "", "director(s) comma seperated")
	flag.StringVarP(&year, "year", "y", "", "year")
	flag.StringVarP(&userId, "userid", "u", "", "userId")
	flag.BoolVarP(&newMovie, "newmovie", "n", false, "newMovie")
	flag.BoolVarP(&deleteMovie, "deletemovie", "D", false, "deleteMovie")
	flag.BoolVarP(&updateMovie, "updatemovie", "U", false, "updateMovie")
	//flag.StringVarP(&requestType, "requestType", "r", "word", "Should be word, sentence or paragraph")
	//flag.IntVarP(&min,"min", "m", 5, "minimum value")
	//flag.IntVarP(&max,"Max", "M", 10, "Maximum value")

	flag.Parse()
	logger := zerolog.New(os.Stderr).With().Timestamp().Caller().Logger()
	ctx := context.Background()
	conn, err := grpc.Dial(grpcAddr, grpc.WithInsecure())
	if err != nil {
		logger.Fatal().Err(err).Msg("")
	}
	defer conn.Close()
	moviesService := client.NewGRPCClient(conn)
	if movieId == "" && newMovie == false{
		movies, err := client.GetMovies(ctx, moviesService)
		if err != nil {
			logger.Fatal().Err(err).Msg("")
		}
		logger.Info().Interface("movies", movies).Msg("")
	}
	if movieId != "" && deleteMovie == false {
		movie, err := client.GetMovieById(ctx, movieId, moviesService)
		if err != nil {
			logger.Fatal().Err(err).Msg("")
		}
		logger.Info().Interface("movie", movie).Msg("")

	}
	if newMovie != false && title != "" && director != "" && year != "" && userId != "" {
		dir := strings.Split(director, ",")
		var dirSlice []string
		for _, d := range dir {
			dirSlice = append(dirSlice, d)
		}
		id, err := client.NewMovie(ctx, title, dirSlice, year, userId, moviesService)
		if err != nil {
			logger.Fatal().Err(err).Msg("")
		}
		logger.Info().Str("id", id).Msg("")
	}
	if deleteMovie != false && movieId != "" {
		err := client.DeleteMovie(ctx, movieId, moviesService)
		if err != nil {
			logger.Fatal().Err(err).Msg("")
		} else {
			logger.Info().Msg("Delete Successful for id: " + movieId)
		}
	}
	if updateMovie != false && movieId != "" && title != "" && director != "" && year != "" && userId != "" {
		dir := strings.Split(director, ",")
		var dirSlice []string
		for _, d := range dir {
			dirSlice = append(dirSlice, d)
		}
		err := client.UpdateMovie(ctx, movieId, title, dirSlice, year, userId, moviesService)
		if err != nil {
			logger.Fatal().Err(err).Msg("")
		}
		logger.Info().Msg("Successfully updated movie with id: " + movieId)
	}
}

