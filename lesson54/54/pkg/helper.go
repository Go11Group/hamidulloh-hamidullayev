package pkg

import (
	"errors"
	"log"
	"reservation_service/config"
	 pbu "reservation_service/genproto/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func CreateReservationManagementClient(cfg *config.Config) pbu.AuthClient {
	conn, err := grpc.NewClient(cfg.Reservation_SERVICE_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(errors.New("failed to connect to the address: " + err.Error()))
		return nil
	}

	return pbu.NewAuthClient(conn)
}
