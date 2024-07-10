package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	pb "dodi/genproto/user/protos"
)

func main() {
	conn, err := grpc.Dial(":7070", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewUserServiceClient(conn)
	user := &pb.UserReq{
		Name: "John Doe",
		Age:  30,
	}

	createdUser, err := client.CreateUser(context.Background(), user)
	if err != nil {
		log.Fatalf("could not create user: %v", err)
	}

	fmt.Printf("Created user: %v\n", createdUser)

}
