package postgres

import (
	"database/sql"
	"fmt"
	pb "reservation_service/genproto/reservations"
	"time"

	"github.com/google/uuid"
)

type ReservationRepo struct {
	Db *sql.DB
}

func NewReservationRepo(db *sql.DB) *ReservationRepo {
	return &ReservationRepo{Db: db}
}

func (r *ReservationRepo) CreateReservation(reservation *pb.Reservation) error {
	query := `
	insert into reservations(
		id,
		restaurant_id,
		user_id,
		arriving_time,
		number_of_seats)
	values($1, $2, $3, $4)
	`

	_, err := r.Db.Exec(query, uuid.NewString(), reservation.RestaurantId, reservation.UserId,
		reservation.ArrivingTime, reservation.NumberOfSeats)
	if err != nil {
		return err
	}
	return nil
}

func (r *ReservationRepo) GetReservationById(id string) (*pb.ReservationInfo, error) {
	query := `
	select 
		id,
		restaurant_id,
		user_id,
		arriving_time,
		number_of_seats,
		created_at,
		updated_at
	from
		reservations
	where
		id = $1 and deleted_at is null
	`
	reservation := pb.ReservationInfo{}
	err := r.Db.QueryRow(query, id).Scan(&reservation.Id, &reservation.RestaurantId, &reservation.UserId,
		&reservation.ArrivingTime, &reservation.NumberOfSeats, &reservation.CreatedAt, &reservation.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &reservation, nil
}

func (r *ReservationRepo) GetAllReservations(filter *pb.ReservationFilter) (*pb.Reservations, error) {
	query := `
	select 
		id,
		restaurant_id,
		user_id,
		arriving_time,
		number_of_seats,
		created_at,
		updated_at
	from
		reservations
	where
		deleted_at is null 
	`
	params := []interface{}{}

	if len(filter.CreatedAt) > 0 {
		params = append(params, filter.CreatedAt)
		query += fmt.Sprintf(` and created_at = $%d`, len(params))
	}
	if filter.NumberOfSeats > 0 {
		params = append(params, filter.NumberOfSeats)
		query += fmt.Sprintf(` and number_of_seats = $%d`, len(params))
	}
	if len(filter.RestaurantId) > 0 {
		params = append(params, filter.NumberOfSeats)
		query += fmt.Sprintf(` and restaurant_id = $%d`, len(params))
	}
	if len(filter.UserId) > 0 {
		params = append(params, filter.NumberOfSeats)
		query += fmt.Sprintf(` and user_id = $%d`, len(params))
	}
	if filter.Limit > 0 {
		params = append(params, filter.Limit)
		query += fmt.Sprintf(` limit $%d`, len(params))
	}
	if filter.Offset > 0 {
		params = append(params, filter.Offset)
		query += fmt.Sprintf(` offset $%d`, len(params))
	}

	rows, err := r.Db.Query(query, params...)
	if err != nil {
		return nil, err
	}
	reservations := pb.Reservations{}

	for rows.Next() {
		r := pb.ReservationInfo{}

		err := rows.Scan(&r.Id, r.RestaurantId, r.UserId, r.ArrivingTime,
			r.NumberOfSeats, r.CreatedAt, r.UpdatedAt)
		if err != nil {
			return nil, err
		}
		reservations.Reservations = append(reservations.Reservations, &r)
	}

	return &reservations, rows.Err()
}

func (r *ReservationRepo) UpdateReservation(reservation *pb.ReservationUpdate) error {
	query := `
	update
		reservations
	set
		restaurant_id = $1,
		user_id = $2,
		arriving_time = $3,
		number_of_seats = $4,
		updated_at = $5
	where
		id = $6
	`

	_, err := r.Db.Exec(query, time.Now(), reservation.Id)

	return err
}

func (r *ReservationRepo) DeleteReservation(id string) error {
	query := `
	update
		reservations
	set
		deleted_at = $1
	where
		id = $2
	`

	_, err := r.Db.Exec(query, time.Now(), id)

	return err
}

func (r *ReservationRepo) ValidateReservationId(id string) (*pb.Exists, error){
	query := `
		select
			case 
				when id = $1 then true
			else
				false
			end
		from
			reservations
		where
		    id = $1 and deleted_at is null
	`

	res := pb.Exists{}
	err := r.Db.QueryRow(query, id).Scan(&res.Exists)
	return &res, err
}
