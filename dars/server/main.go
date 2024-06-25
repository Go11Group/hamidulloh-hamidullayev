package main

import (
	"context"
	pb "dodi/genproto"
	"google.golang.org/grpc"
	"math/rand"
	"net"
	"time"
)

type server struct {
	pb.UnimplementedGeneratorServer
}

func (s *server) RandomPicker(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	numbers := []int{}
	
	for i := 1; i <= 24; i++ {
		if _, ok := in.Exception[int32(i)]; !ok {
			numbers = append(numbers, i)
		}
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(numbers), func(i, j int) { numbers[i], numbers[j] = numbers[j], numbers[i] })

	assignment := make(map[string]int32)
	for i, name := range in.Names {
		assignment[name] = int32(numbers[i])
	}

	return &pb.Response{Result: assignment}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	s := server{}
	grpc := grpc.NewServer()
	pb.RegisterGeneratorServer(grpc, &s)

	err = grpc.Serve(listener)
	if err != nil {
		panic(err)
	}
}