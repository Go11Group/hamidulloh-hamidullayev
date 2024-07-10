package postgres

import (
	"database/sql"
	"fmt"
	"log"
	pb "reservation_service/genproto/order"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type OrderRepo struct {
	Db *sql.DB
}

func NewOrderRepo(db *sql.DB) *OrderRepo {
	return &OrderRepo{Db: db}
}

func (o *OrderRepo) CreateOrder(req *pb.Order) error {
	_, err := o.Db.Exec("insert into orders(reservation_id, menu_item_id, quantity, created_at) values($1, $2, $3, $4)", req.ReservationId, req.MenuItemId, req.Quantity, time.Now())
	if err != nil {
		return err
	}
	return nil
}

func (o *OrderRepo) UpdateOrder(req *pb.Updateorder) error {
	_, err := o.Db.Exec("update orders set reservation_id=$1, menu_item_id=$2, quantity=$3, updated_at=$4 where id=$5", req.ReservationId, req.MenuItemId, req.Quantity, time.Now(), )
	if err != nil {
		return err
	}
	return nil
}

func (o *OrderRepo) DeleteOrder(rep *pb.Id) error {
	_, err := uuid.Parse(rep.Id)
	if err != nil {
		log.Printf("Error parsing UUID: %v", err)
		return err
	}
	_, err = o.Db.Exec("update order set delete_at=$1 where id=$2", time.Now(), rep.Id)
	if err != nil {
		return err
	}
	return nil
}

func (o *OrderRepo) GetOrderById(id *pb.Id) (*pb.OrderInfo, error) {
	_, err := uuid.Parse(id.Id)
	if err != nil {
		log.Printf("Error parsing UUID: %v", err)
		return nil, err
	}
	resp := pb.OrderInfo{}
	err = o.Db.QueryRow("select * from orders where id=$1", id).Scan(&resp.ReservationId, &resp.MenuItemId, &resp.Quantity, &resp.CreatedAt, &resp.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (o *OrderRepo) GetAllOrder(req *pb.OrderFilter) (*pb.Orders, error) {
	var params []interface{}

	query := "select id,reservation_id,menu_item_id ,quantity, created_at, updated_at from orders where  deleted_at is null "

	if len(req.MenuItemId) > 0 {
		params = append(params, req.MenuItemId)
		query += fmt.Sprintf(" and menu_item_id = $%d", len(params))

	}
	if len(req.ReservationId) > 0 {
		params = append(params, req.ReservationId)
		query += fmt.Sprintf(" and reservation_id = $%d", len(params))
	}

	if req.Quantity > 0 {
		params = append(params, req.Quantity)
		query += fmt.Sprintf(" and quantity = $%d", req.Quantity)

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
	var orders pb.Orders
	for rows.Next() {
		var order pb.OrderInfo
		err := rows.Scan(&order.Id, &order.ReservationId, &order.MenuItemId, &order.Quantity, &order.CreatedAt, &order.UpdatedAt)
		if err != nil {
			return nil, err
		}
		orders.Orders = append(orders.Orders, &order)
	}
	return &orders, nil
}
