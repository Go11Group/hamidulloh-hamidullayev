package main

import (
	"dodi/api"
	"dodi/genproto/protos"
	"fmt"

	"github.com/grpc/grpc-go/credentials/insecure"
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.NewClient(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(err)
	}

	weather := protos.NewWeatherServiceClient(conn)

	fmt.Println(weather)
	r := api.NewRouter(conn)

	r.Run()
}