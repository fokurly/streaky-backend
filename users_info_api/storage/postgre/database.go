package postgre

import (
	"fmt"

	"github.com/fokurly/streaky-backend/users_info_api/models"
	"github.com/fokurly/streaky-backend/users_info_api/utils"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type Db struct {
	db     *sqlx.DB
	config models.DatabaseConfig
}

func NewDatabase() *Db {
	config := utils.ParseDatabaseConfigByKey("database_config", false)
	query := fmt.Sprintf("host=0.0.0.0 port=5432 user=%s password=%s dbname=%s sslmode=disable", config.User, config.Password, config.Dbname)
	logrus.Println(query)
	logrus.Println(config.Host)
	db, err := sqlx.Open("postgres", query)

	if err != nil {
		logrus.Panicf(fmt.Sprintf("could not open db"))
	}

	return &Db{db: db, config: config}
}
