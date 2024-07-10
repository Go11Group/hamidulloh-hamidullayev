package service

import (
	"context"
	"database/sql"
	pb "reservation_service/genproto/restaurant"
	"reservation_service/storage/postgres"
)

type RestaurantService struct {
	pb.UnimplementedRestaurantServer
	Repo *postgres.RestaurantRepo
}

func NewRestaurantService(db *sql.DB) *RestaurantService {
	return &RestaurantService{
		Repo: postgres.NewRestaurantRepo(db),
	}
}

func (u *RestaurantService) CreateRestaurant(ctx context.Context, req *pb.RestaurantCreate) error {
	err := u.Repo.CreateRestaurant(req)
	if err != nil {
		return err
	}
	return nil
}

func (u *RestaurantService) UpdateRestaurant(ctx context.Context, rep *pb.RestaurantUpdate) error {
	err := u.Repo.UpdateRestaurant(rep)
	if err != nil {
		return err
	}
	return nil
} 

func (u *RestaurantService) DeleteRestaurant(ctx context.Context, rep *pb.Id) error {
	err := u.Repo.DeleteRestaurant(rep)
	if err != nil {
		return err
	}
	return nil
}

func (u *RestaurantService) GetRestaurants(ctx context.Context, req *pb.RestaurantFilter) (*pb.Restaurants, error) {
	resp, err := u.Repo.GetRestaurants(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (u *RestaurantService) GetRestaurantById(rep *pb.Id) (*pb.RestaurantInfo, error) {
	resp, err := u.Repo.GetRestaurantById(rep)
	if err != nil {
		return nil, err
	}
	return resp, nil
}