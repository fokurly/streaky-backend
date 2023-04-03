package service

import (
	"github.com/fokurly/streaky-backend/users_info_api/storage/postgre"
)

type usersInfoApi struct {
	db *postgre.Db
}

func NewUsersInfoApi(db *postgre.Db) *usersInfoApi {
	return &usersInfoApi{db: db}
}
