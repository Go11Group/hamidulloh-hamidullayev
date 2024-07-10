package postgres

import (
	"database/sql"
	"fmt"
	"log"
	pb "reservation_service/genproto/menu"
	"time"

	"github.com/google/uuid"
)


type MenuRepo struct{
	Db *sql.DB
}

func NewMenuRepo(db *sql.DB) *MenuRepo{
	return &MenuRepo{Db: db}
} 

func (o *MenuRepo) CreateMenu(rep *pb.MenuRequest) error {
	_, err := o.Db.Exec("insert into menu(item_type, name, price, descriptioon, created_at) values($1, $2, $3, $4, $5)", rep.ItemType, rep.Name, rep.Price, rep.Description, time.Now())
	if err != nil{
		return err
	}
	return nil
}

func (o *MenuRepo) UpdateMenu(rep *pb.MenuUpateRequest) error {
	_, err := o.Db.Exec("update menu set item_type=$1, name=$2, price=$3, descriptioon=$4, restaurant_id=$5, updated_at=$6 where id=$7", rep.ItemType, rep.Name, rep.Price, rep.Description, rep.RestaurantId, time.Now(), rep.Id)
	if err != nil{
		return err
	}
	return nil
}

func (o *MenuRepo) DeleteMenu(rep *pb.Id) error {
	_, err := uuid.Parse(rep.Id)
	if err != nil {
		log.Printf("Error parsing UUID: %v", err)
		return err
	}
	_, err = o.Db.Exec("update menu set delete_at=$1 where id=$2", time.Now(), rep.Id)
	if err != nil {
		return err
	}
	return nil
}

func (o *MenuRepo) GetByIdMenu(rep *pb.Id) (*pb.MenuResponse, error) {
	_, err := uuid.Parse(rep.Id)
	if err != nil {
		log.Printf("Error parsing UUID: %v", err)
		return nil, err
	}
	resp := pb.MenuResponse{}
	err = o.Db.QueryRow("select * from menu where id=$1", rep.Id).Scan(&resp.Id, &resp.ItemType, &resp.Name, &resp.Price, 
		&resp.Description, resp.RestaurantId, resp.CreatedAt, resp.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (o *MenuRepo) GetAllMenu(req *pb.MenuFilter) (*pb.Menus, error) {
	var params []interface{}

	query := "select id, item_type, name, price, descriptioon, RestaurantId, created_at, updated_at from orders where  deleted_at is null "

	if len(req.ItemType) > 0 {
		params = append(params, req.ItemType)
		query += fmt.Sprintf(" and item_type = $%d", len(params))

	}
	if len(req.Name) > 0 {
		params = append(params, req.Name)
		query += fmt.Sprintf(" and name = $%d", len(params))
	}

	if len(req.Description) > 0 {
		params = append(params, req.Description)
		query += fmt.Sprintf(" and descriptioon = $%d", len(params))

	}
	if len(req.RestaurantId) > 0 {
		params = append(params, req.RestaurantId)
		query += fmt.Sprintf(" and RestaurantId = $%d", len(params))
	}

	if req.Price > 0 {
		params = append(params, req.Price)
		query += fmt.Sprintf(" and price = $%d", len(params))

	}

	if req.Limit > 0 {
		params = append(params, req.Limit)
		query = fmt.Sprintf(" LIMIT = $%d", len(params))

	}
	if req.Offset > 0 {
		params = append(params, req.Offset)
		query = fmt.Sprintf(" OFFSET = $%d", len(params))

	}

	rows, err := o.Db.Query(query, params...)
	if err != nil {
		return nil, err
	}
	var menus pb.Menus
	for rows.Next() {
		var menu pb.MenuResponse
		err := rows.Scan(&menu.Id, &menu.ItemType, &menu.Name, &menu.Price, &menu.Description, &menu.RestaurantId, &menu.CreatedAt, &menu.UpdatedAt)
		if err != nil {
			return nil, err
		}
		menus.Menus = append(menus.Menus, &menu)
	}
	return &menus, nil
}