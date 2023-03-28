package postgre

import (
	"fmt"

	"github.com/fokurly/streaky-backend/users_info_api/models"
)

func (d *Db) InsertNewUser(user models.User) error {
	const (
		insertNewUser = `INSERT INTO user_register_info(email, login, password) VALUES($1, $2, $3)`
	)
	_, err := d.db.Exec(insertNewUser, user.Email, user.FullName, user.Password)
	if err != nil {
		return fmt.Errorf("[InsertNewUser] - could not exec query. error: %v", err)
	}

	return nil
}
