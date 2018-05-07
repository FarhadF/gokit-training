package main

import (
	flag "github.com/spf13/pflag"
	"gokit-training/vault"
	"context"
	"os"
	"os/signal"
	"syscall"
	"fmt"
	"net/http"
	"github.com/rs/zerolog"
	"net"
	"gokit-training/vault/pb"
	"google.golang.org/grpc"
)

func main() {
	var (
		httpAddr = flag.StringP("http", "H",":8080",
			"http listen address")
		gRPCAddr = flag.StringP("grpc", "G",":8081",
			"gRPC listen address")
	)
	flag.Parse()
	ctx := context.Background()
	//zerolog
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	svc := vault.NewService()
	errChan := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()
	hashEndpoint := vault.MakeHashEndpoint(svc)
	validateEndpoint := vault.MakeValidateEndpoint(svc)
	endpoints := vault.Endpoints{
		HashEndpoint:
		hashEndpoint,
		ValidateEndpoint: validateEndpoint,
	}
	// HTTP transport
	go func() {
		logger.Info().Str("http:", *httpAddr).Msg("")
		handler := vault.NewHTTPServer(ctx, endpoints)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()
	// GRPC transport
	go func() {
		listener, err := net.Listen("tcp", *gRPCAddr)
		if err != nil {
			errChan <- err
			return
		}
		logger.Info().Str("grpc:", *gRPCAddr).Msg("")
		handler := vault.NewGRPCServer(ctx, endpoints)
		gRPCServer := grpc.NewServer()
		pb.RegisterVaultServer(gRPCServer, handler)
		errChan <- gRPCServer.Serve(listener)
	}()

	logger.Info().Msg(*gRPCAddr)
	logger.Fatal().Err(<-errChan).Msg("")
}
