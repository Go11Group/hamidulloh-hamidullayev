package postgres

import (
	"database/sql"
	"fmt"
	"log"
	pb "reservation_service/genproto/restaurant"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type RestaurantRepo struct {
	Db *sql.DB
}

func NewRestaurantRepo(db *sql.DB) *RestaurantRepo {
	return &RestaurantRepo{Db: db}
}

func (r *RestaurantRepo) CreateRestaurant(req *pb.RestaurantCreate) error {
	_, err := r.Db.Exec("insert into restaurants(name, address, total_avb_seats, phone_number, description, created_at) values($1, $2, $3, $4, $5, $6)",
		req.Name, req.Address, req.TotalAvbSeats, req.PhoneNumber, req.Description, time.Now())

	if err != nil {
		log.Fatal("Create query failed: ", err)
		return err
	}
	return nil
}

func (r *RestaurantRepo) UpdateRestaurant(restaurant *pb.RestaurantUpdate) error {
	_, err := r.Db.Exec("update restaurants set name=$1, address=$2, total_avb_seats=$3, phone_number=$4, description=$5, updated_at=$6 where id=$7",
		restaurant.Name, restaurant.Address, restaurant.TotalAvbSeats, restaurant.PhoneNumber, restaurant.Description, time.Now(), restaurant.Id)
	if err != nil {
		log.Println("Update query failed: ", err)
		return err
	}
	return nil
}

func (r *RestaurantRepo) DeleteRestaurant(id *pb.Id) error {
	_, err := uuid.Parse(id.Id)
	if err != nil {
		log.Printf("Error parsing UUID: %v", err)
		return err
	}
	_, err = r.Db.Exec("update restaurants set deleted_at=$1 where id=$2", time.Now(), id)

	if err != nil {
		log.Fatal("Delete query failed: ", err)
		return err
	}

	return nil
}

func (r *RestaurantRepo) GetRestaurantById(id *pb.Id) (*pb.RestaurantInfo, error) {
	_, err := uuid.Parse(id.Id)
	if err != nil {
		log.Printf("Error parsing UUID: %v", err)
		return nil, err
	}
	rest := pb.RestaurantInfo{}
	err = r.Db.QueryRow("select name, address, total_avb_seats, phone_number, description, updated_at, created_at "+
		" from restaurants where id=$1 and deleted_at is null", id).Scan(&rest.Id, &rest.Name, &rest.Address, &rest.TotalAvbSeats,
		&rest.PhoneNumber, &rest.Description, &rest.UpdatedAt, &rest.CreatedAt)

	if err != nil {
		log.Fatal("Get by id query failed: ", err)
		return nil, err
	}
	return &rest, nil
}

func (r *RestaurantRepo) GetRestaurants(req *pb.RestaurantFilter) (*pb.Restaurants, error) {
	var params []interface{}

	query := `select
					id, 
					name, 
					address, 
					total_avb_seats, 
					phone_number, 
					description, 
					updated_at, 
					created_at
		from restaurants where  deleted_at is null `

	if len(req.Name) > 0 {
		params = append(params, req.Name)
		query += fmt.Sprintf(" and name = $%d", len(params))
	}

	if len(req.Address) > 0 {
		params = append(params, req.Address)
		query += fmt.Sprintf(" and address = $%d", len(params))
	}

	if req.Limit > 0 {
		params = append(params, req.Limit)
		query += fmt.Sprintf(" and limit = $%d", len(params))
	}

	if req.Offset > 0 {
		params = append(params, req.Offset)
		query += fmt.Sprintf(" and offset = $%d", len(params))
	}

	rows, err := r.Db.Query(query, params...)
	if err != nil {
		return nil, err
	}
	var respons pb.Restaurants
	for rows.Next() {
		var resp pb.RestaurantInfo
		err := rows.Scan(&resp.Id, &resp.Name, &resp.Address, &resp.TotalAvbSeats, &resp.PhoneNumber, &resp.Description, &resp.CreatedAt, &resp.UpdatedAt)
		if err != nil {
			return nil, err
		}
		respons.Restaurants = append(respons.Restaurants, &resp)
	}
	return &respons, nil
}
