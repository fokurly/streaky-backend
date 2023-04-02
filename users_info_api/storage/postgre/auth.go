package postgre

import (
	"fmt"

	"github.com/fokurly/streaky-backend/users_info_api/models"
)

func (d *Db) InsertNewUser(user models.User) error {
	const (
		insertNewUser = `INSERT INTO user_register_info(email, login, full_name, password) VALUES($1, $2, $3, $4)`
	)
	_, err := d.db.Exec(insertNewUser, user.Email, user.Login, user.FullName, user.Password)
	if err != nil {
		return fmt.Errorf("[InsertNewUser] - could not exec query. error: %v", err)
	}

	return nil
}

func (d *Db) GetUser(login, pass string) (*int64, error) {
	const (
		getUser = `SELECT id FROM user_register_info WHERE login=$1 and password=$2`
	)
	rows, err := d.db.Query(getUser, login, pass)
	if err != nil {
		return nil, fmt.Errorf("[GetBalance] - could not exec query. error: %v", err)
	}

	defer func() {
		_ = rows.Close()
	}()

	var ID int64
	if rows.Next() {
		err := rows.Scan(&ID)
		if err != nil {
			return nil, fmt.Errorf("[GetUser] - could not scan rows. error: %v", err)
		}
	} else {
		return nil, fmt.Errorf("[GetUser] - no rows with such id")
	}
	return &ID, nil
}
