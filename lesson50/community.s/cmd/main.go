package main

import (
	"EcoTrack/community/pkg/db"
	storage "EcoTrack/community/storage/postgres"
	s "EcoTrack/community/service"
	p "EcoTrack/community/genproto/protos"

	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":7070")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	server := grpc.NewServer()

	db, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	storage := storage.NewCommunityRepo(db)
	service := s.CommunityService{Storage: *storage}

	p.RegisterCommunityServiceServer(server, &service)

	log.Println("server is running on :7070 ...")
	if err := server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
