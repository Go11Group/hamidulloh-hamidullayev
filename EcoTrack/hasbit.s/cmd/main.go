package main

import (
	"Habits/Tracker/pkg/db"
	storage "Habits/Tracker/Storage/postgres"
	pb "Habits/Tracker/genproto/HabitTracker"
	service "Habits/Tracker/Service"
	"net"
	"log"
	"google.golang.org/grpc"
)

func main() {
	db, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	grpcServer := grpc.NewServer()
	ht := service.NewHabitsTracker(storage.NewHabitTracker(db))
	
	pb.RegisterHabitTrackerServiceServer(grpcServer,ht)

	grpcServer.Serve(listener)

	log.Printf("Server started on port 50051")
}
