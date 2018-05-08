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
	"github.com/rs/zerolog/log"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main(){
	//zerolog
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	var (
		console bool
		httpAddr string
		gRPCAddr string
	)
	flag.StringVarP(&httpAddr, "http", "H",":8080","http listen address")
	flag.StringVarP(&gRPCAddr,"grpc", "g", ":8081", "GRPC Address")
	flag.BoolVarP(&console, "console", "c", false, "turns on pretty console logging" )
	flag.Parse()
	logger.Info().Msg("starting grpc server at"+ string(gRPCAddr))
	if console {
		logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
	ctx := context.Background()

	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "my_group",
		Subsystem: "lorem_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "lorem_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
	countResult := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "lorem_service",
		Name:      "count_result",
		Help:      "The result of each count method.",
	}, []string{}) // no fields here




	// init lorem service
	var svc lorem.Service
	loremStruct := lorem.NewService()
	svc = loremStruct
	svc = lorem.InstrumentingMiddleware{requestCount, requestLatency, countResult, svc}
	errChan := make(chan error)
	// creating Endpoints struct
	endpoints := lorem.Endpoints{
		lorem.MakeLoremEndpoint(svc),
	}
	//execute grpc server
	go func() {
		listener, err := net.Listen("tcp", gRPCAddr)
		if err != nil {
			errChan <- err
			return
		}
		handler := lorem.NewGRPCServer(ctx, endpoints)
		grpcServer := grpc.NewServer()
		pb.RegisterLoremServer(grpcServer, handler)
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
