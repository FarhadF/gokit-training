package main

import (
	"os"
	flag "github.com/spf13/pflag"
	"net"
	"github.com/julienschmidt/httprouter"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"os/signal"
	"syscall"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"context"
	"gokit-training/movies"
	"google.golang.org/grpc"
	"gokit-training/movies/pb"
	"database/sql"
	_ "github.com/lib/pq"
)

func main(){
	//zerolog
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	var (
		console bool
		httpAddr string
		gRPCAddr string
	)
	flag.StringVarP(&httpAddr, "http", "H",":8082","http listen address")
	flag.StringVarP(&gRPCAddr,"grpc", "g", ":8081", "GRPC Address")
	flag.BoolVarP(&console, "console", "c", false, "turns on pretty console logging" )
	flag.Parse()
	logger.Info().Msg("starting grpc server at"+ string(gRPCAddr))
	if console {
		logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
	ctx := context.Background()



	// init movies service
	var svc movies.Service
	db, err := sql.Open("postgres", "postgresql://app_user@localhost:26257/app_database?sslmode=disable")
	if err != nil {
		logger.Fatal().Err(err).Msg("db connection failed")
	}
	svc = movies.NewService(db, logger)
	//svc = movies.LoggingMiddleware{logger, svc}
	//svc = movies.InstrumentingMiddleware{requestCount, requestLatency, countResult, svc}
	errChan := make(chan error)
	// creating Endpoints struct
	endpoints := movies.Endpoints{
		movies.MakeGetMoviesEndpoint(svc),
	}
	//execute grpc server
	go func() {
		listener, err := net.Listen("tcp", gRPCAddr)
		if err != nil {
			errChan <- err
			return
		}
		handler := movies.NewGRPCServer(ctx, endpoints)
		grpcServer := grpc.NewServer()
		pb.RegisterMoviesServer(grpcServer, handler)
		errChan <- grpcServer.Serve(listener)
	}()
	// HTTP transport
	go func() {
		//httprouter initialization
		router := httprouter.New()
		//handler will be used for net/http handle compatibility
		router.Handler("GET","/metrics", promhttp.Handler())
		errChan <- http.ListenAndServe(httpAddr, router)
	}()
	//Handle os signals
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()
	logger.Error().Err(<- errChan).Msg("")
}

