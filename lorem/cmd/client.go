package main

import (
	"google.golang.org/grpc"
	flag "github.com/spf13/pflag"
	"context"
	"github.com/rs/zerolog"
	"os"
	"gokit-training/lorem/client"
)



func main() {
	var (
		grpcAddr = flag.StringP("addr", "a", ":8081", "gRPC address")
		//method string
		requestType string
		min         int
		max         int
	)
	flag.StringVarP(&requestType, "requestType", "r", "word", "Should be word, sentence or paragraph")
	flag.IntVarP(&min, "min", "m", 5, "minimum value")
	flag.IntVarP(&max, "Max", "M", 10, "Maximum value")

	flag.Parse()
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	ctx := context.Background()
	conn, err := grpc.Dial(*grpcAddr, grpc.WithInsecure())
	if err != nil {
		logger.Fatal().Err(err).Msg("")
	}
	defer conn.Close()
	loremService := client.NewGRPCClient(conn)
	m, err := client.Lorem(ctx, loremService, requestType, min, max)
	if err != nil {
		logger.Error().Err(err).Msg("")
	}
	logger.Info().Msg(m)

}


