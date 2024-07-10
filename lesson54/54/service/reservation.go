package service

import (
	"context"
	pbu "reservation_service/genproto/auth"
	pb "reservation_service/genproto/reservations"
	"reservation_service/storage/postgres"

	"github.com/google/uuid"
)

type ReservationService struct {
	AuthClient  pbu.AuthClient
	Reservation *postgres.ReservationRepo
	pb.UnimplementedReservationServiceServer
}

func NewReservationService(reservation *postgres.ReservationRepo, auth pbu.AuthClient) *ReservationService {
	return &ReservationService{Reservation: reservation, AuthClient: auth}
}

func (r *ReservationService) CreateReservation(cntx context.Context, req *pb.Reservation) (*pb.Void, error) {

	exist, err := r.AuthClient.ValidateUserId(cntx, &pbu.Id{Id: req.UserId})
	if !exist.Exists || err != nil {
		return &pb.Void{}, err
	}

	err = r.Reservation.CreateReservation(req)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

func (r *ReservationService) UpdateReservation(cntx context.Context, req *pb.ReservationUpdate) (*pb.Void, error) {

	exist, err := r.AuthClient.ValidateUserId(cntx, &pbu.Id{Id: req.UserId})
	if !exist.Exists || err != nil {
		return &pb.Void{}, err
	}
	err = r.Reservation.UpdateReservation(req)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}

func (r *ReservationService) DeleteReservation(cnts context.Context, req *pb.Id) (*pb.Void, error) {
	_, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}

	err = r.Reservation.DeleteReservation(req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}

func (r *ReservationService) GetReservationById(cntx context.Context, req *pb.Id) (*pb.ReservationInfo, error) {
	_, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	resp, err := r.Reservation.GetReservationById(req.Id)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *ReservationService) GetAllReservations(cntx context.Context, req *pb.ReservationFilter) (*pb.Reservations, error) {
	exist, err := r.AuthClient.ValidateUserId(cntx, &pbu.Id{Id: req.UserId})
	if !exist.Exists || err != nil {
		return nil, err
	}

	resp, err := r.Reservation.GetAllReservations(req)
	if err != nil {
		return nil, err
	}
	return resp, nil

}