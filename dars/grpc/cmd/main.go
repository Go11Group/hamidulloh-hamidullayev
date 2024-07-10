package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	p "dodi/genproto/user/protos"
	"dodi/storage"

	_ "github.com/lib/pq"
)

func main() {

	listener, err := net.Listen("tcp", ":7070")
	if err != nil {
		log.Fatal(err)
	}

	defer listener.Close()

	server := grpc.NewServer()

	var psqlUrl = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost",
		"5432",
		"postgres",
		"dodi",
		"lesson36",
	)

	psqlConn, err := sql.Open("postgres", psqlUrl)
	if err != nil {
		log.Fatalf("failed to connect database: %s", err)
	}

	storage := storage.NewUserStorage(psqlConn)

	p.RegisterUserServiceServer(server, &UserServer{
		Storage: *storage,
	})

	log.Println("server is running on :7070 ...")
	if err := server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}

type UserServer struct {
	p.UnimplementedUserServiceServer

	Storage storage.UserStorage
}

var userData []p.User

func (u *UserServer) CreateUser(ctx context.Context, req *p.UserReq) (*p.User, error) {

	res, err := u.Storage.CreateUser(req)
	if err != nil {
		log.Println("error while creating user : ", err)
		return nil, err
	}

	return res, nil

}

func (u *UserServer) GetById(ctx context.Context, req *p.Id) (*p.User, error) {

	for _, user := range userData {
		if user.Id == req.Id {
			return &user, nil
		}
	}

	return nil, fmt.Errorf("User not found")
}
