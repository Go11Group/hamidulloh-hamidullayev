package service

import (
	"context"
	"database/sql"
	pb "reservation_service/genproto/order"
	"reservation_service/storage/postgres"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	Repo *postgres.OrderRepo
}

func NewOrderService(db *sql.DB) *OrderService {
	return &OrderService{
		Repo: postgres.NewOrderRepo(db),
	}
}

func (u *OrderService) CreateOrder(ctx context.Context, req *pb.Order) error {
	err := u.Repo.CreateOrder(req)
	if err != nil{
		return err
	}
	return nil
}

func (u *OrderService) UpdateOrder(ctx context.Context, rep *pb.Updateorder) error {
	err := u.Repo.UpdateOrder(rep)
	if err != nil{
		return err
	}
	return nil
}

func (u *OrderService) DeleteOrder(ctx context.Context, rep *pb.Id) error {
	err := u.Repo.DeleteOrder(rep)
	if err != nil{
		return err
	}
	return nil
}

func (u *OrderService) GetOrderById(ctx context.Context, rep *pb.Id) (*pb.OrderInfo, error) {
	resp, err := u.Repo.GetOrderById(rep)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (u *OrderService) GetAllOrder(ctx context.Context, rep *pb.OrderFilter) (*pb.Orders, error) {
	resp, err := u.Repo.GetAllOrder(rep) 
	if err != nil {
		return nil, err
	}
	return resp, err
}