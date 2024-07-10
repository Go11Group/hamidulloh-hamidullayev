package storage

import (
	"database/sql"
	genproto "dodi/genproto/user/protos"
	"log"
)

type UserStorage struct {
	db *sql.DB
}

func NewUserStorage(db *sql.DB) *UserStorage {
	return &UserStorage{
		db: db,
	}
}

func (u *UserStorage) CreateUser(req *genproto.UserReq) (*genproto.User, error) {

	res := genproto.User{}

	row := u.db.QueryRow("insert into users(name, age) values ($1, $2) returning id, name, age", req.Name, req.Age)
	err := row.Scan(&res.Id, &res.Name, &res.Age)
	if err != nil {
		log.Println("error while scannig : ", err)
		return nil, err
	}

	return &res, nil
}
