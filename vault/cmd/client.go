package main

import (
	flag "github.com/spf13/pflag"
	"google.golang.org/grpc"
	"time"
	"context"
	"github.com/rs/zerolog"
	"os"
	"gokit-training/vault/client"
)

func main() {
	var (
		grpcAddr = flag.String("addr", ":8081", "gRPC address")
	)
	flag.Parse()
	ctx := context.Background()
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	conn, err := grpc.Dial(*grpcAddr, grpc.WithInsecure(),
		grpc.WithTimeout(1*time.Second))
	if err != nil {
		logger.Fatal().Err(err).Msg("grpc dial err")
	}
	defer conn.Close()
	vaultService := client.NewGRPCClient(conn)
	args := flag.Args()
	var cmd string
	cmd, args = pop(args)
	switch cmd {
	case "hash":
		var password string
		password, args = pop(args)
		hash, err := client.Hash(ctx, vaultService, password)
		if err != nil {
			logger.Error().Err(err).Msg("")
		} else {
			logger.Info().Msg(hash)
		}
	case "validate":
		var password, hash string
		password, args = pop(args)
		hash, args = pop(args)
		valid, err := client.Validate(ctx, vaultService, password, hash)
		if err != nil {
			logger.Error().Err(err).Msg("")
		} else {
			logger.Info().Bool("valid", valid).Msg("")
		}

	default:
		logger.Fatal().Str("unknown command", cmd).Msg("")
	}
}

func pop(s []string) (string, []string) {
	if len(s) == 0 {
		return "", s
	}
	return s[0], s[1:]
}

