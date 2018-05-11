package main

import (
"google.golang.org/grpc"
"gokit-training/lorem"
grpctransport "github.com/go-kit/kit/transport/grpc"
"gokit-training/lorem/pb"
flag "github.com/spf13/pflag"
"context"
"github.com/rs/zerolog"
"os"
"fmt"
)
//create new client returns Lorem Service
func NewGRPCClient(conn *grpc.ClientConn) lorem.Service {
	var loremEndpoint = grpctransport.NewClient(
		conn, "pb.Lorem", "Lorem",
		lorem.EncodeGRPCLoremRequest,
		lorem.DecodeGRPCLoremResponse,
		pb.LoremResponse{},
	).Endpoint()
	return lorem.Endpoints{
		loremEndpoint,
	}
}


func main(){
	var (
		grpcAddr = flag.StringP("addr", "a",":8081","gRPC address")
		//method string
		requestType string
		min int
		max int
	)
	//service = flag.StringP("service", "s","Lorem")
	//flag.StringVarP(&method,"method", "m", "lorem", "The only method available right now")
	flag.StringVarP(&requestType, "requestType", "r", "word", "Should be word, sentence or paragraph")
	flag.IntVarP(&min,"min", "m", 5, "minimum value")
	flag.IntVarP(&max,"Max", "M", 10, "Maximum value")


	flag.Parse()
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	ctx := context.Background()
	conn, err := grpc.Dial(*grpcAddr, grpc.WithInsecure())
	if err != nil {
		logger.Fatal().Err(err).Msg("")
	}
	defer conn.Close()
	loremService := NewGRPCClient(conn)
	//switch method {
	//case "lorem":

	callLorem(ctx, loremService, requestType, min, max, logger)

	//}
}

//callService helper
func callLorem(ctx context.Context, service lorem.Service, requestType string, min int, max int, logger zerolog.Logger) {
	mesg, err := service.Lorem(ctx, requestType, min, max)
	if err != nil {
		logger.Fatal().Err(err).Msg("")
	}
	fmt.Println(mesg)
}