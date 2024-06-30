package main 

import (
	"log"
	"net"

	"dodi/db"
	pb "dodi/genproto/pro1tos" // Protokol paketini Go tiliga ko'chirish
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)


func main() {
	listener, _ := net.Listen("tcp", ":50051")
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &db.Server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
