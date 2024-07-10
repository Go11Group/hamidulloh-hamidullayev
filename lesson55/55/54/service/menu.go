package service

import (
	"context"
	"database/sql"
	pb "reservation_service/genproto/menu"
	"reservation_service/storage/postgres"
)

type MenuService struct {
	pb.UnimplementedMenuServiceServer
	Repo *postgres.MenuRepo
}

func NewMenuService(db *sql.DB) *MenuService{
	return &MenuService{
		Repo: postgres.NewMenuRepo(db),
	}
}

func (u *MenuService) CreateMenu(ctx context.Context, req *pb.MenuRequest) error {
	err := u.Repo.CreateMenu(req)
	if err != nil{
		return err
	}
	return nil
}

func (u *MenuService) UpdateMenu(ctx context.Context, rep *pb.MenuUpateRequest) error {
	err := u.Repo.UpdateMenu(rep)
	if err != nil{
		return err
	}
	return nil
}

func (u *MenuService) DeleteMenu(ctx context.Context, rep *pb.Id) error {
	err := u.Repo.DeleteMenu(rep)
	if err != nil{
		return err
	}
	return nil
}

func (u *MenuService) GetByIdMenu(ctx context.Context, rep *pb.Id) (*pb.MenuResponse, error) {
	resp, err := u.Repo.GetByIdMenu(rep)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (u *MenuService) GetAllMenu(ctx context.Context, rep *pb.MenuFilter) (*pb.Menus, error) {
	resp, err := u.Repo.GetAllMenu(rep) 
	if err != nil {
		return nil, err
	}
	return resp, err
}