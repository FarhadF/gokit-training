package main

import (
	flag "github.com/spf13/pflag"
	"context"
	"gokit-training/lorem"
	"net"
	"google.golang.org/grpc"
	"gokit-training/lorem/pb"
	"os"
	"os/signal"
	"syscall"
	"fmt"
	"github.com/rs/zerolog"
)

func main(){
	//zerolog
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	gRPCAddr := flag.StringP("grpc", "g", ":8081", "GRPC Address")
	flag.Parse()
	logger.Info().Msg("starting grpc server at"+ string(*gRPCAddr))
	ctx := context.Background()
	// init lorem service
	var svc lorem.Service
	loremStruct := lorem.NewService()
	svc = loremStruct
	errChan := make(chan error)
	// creating Endpoints struct
	endpoints := lorem.Endpoints{
		lorem.MakeLoremEndpoint(svc),
	}
	//execute grpc server
	go func() {
		listener, err := net.Listen("tcp", *gRPCAddr)
		if err != nil {
			errChan <- err
			return
		}
		handler := lorem.NewGRPCServer(ctx, endpoints)
		grpcServer := grpc.NewServer()
		pb.RegisterLoremServer(grpcServer, handler)
		errChan <- grpcServer.Serve(listener)
	}()
	//Handle os signals
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()
	logger.Error().Err(<- errChan).Msg("")
}
