package service

import (
	"github.com/fokurly/streaky-backend/users_info_api/storage/postgre"
)

type serviceStreakyApi struct {
	db *postgre.Db
}

func NewStreakyApi(db *postgre.Db) *serviceStreakyApi {
	return &serviceStreakyApi{db: db}
}
