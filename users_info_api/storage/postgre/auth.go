package postgre

import (
	"fmt"

	"github.com/fokurly/streaky-backend/users_info_api/models"
)

func (d *Db) InsertNewUser(user models.User) error {
	const (
		insertNewUser    = `INSERT INTO user_register_info(email, login, full_name, password) VALUES($1, $2, $3, $4)`
		insertUserFriend = `INSERT INTO user_friend_list(userid) VALUES ($1)`
		insertUserTasks  = `INSERT INTO user_tasks(user_id) VALUES ($1)`
	)
	_, err := d.db.Exec(insertNewUser, user.Email, user.Login, user.FullName, user.Password)
	if err != nil {
		return fmt.Errorf("[InsertNewUser] - could not exec query. error: %v", err)
	}

	{
		userAuth := models.UserAuth{
			Login:    user.Login,
			Password: user.Password,
		}

		id, err := d.GetUserID(userAuth)
		if err != nil {
			return fmt.Errorf("could not get user id. smth is wrong. error: %v", err)
		}

		{
			_, err = d.db.Exec(insertUserFriend, id)
			if err != nil {
				return fmt.Errorf("[InsertNewUser] - could not create friend list. error: %v", err)
			}
		}

		{
			_, err = d.db.Exec(insertUserTasks, id)
			if err != nil {
				return fmt.Errorf("[InsertNewUser] - could not create task list for user. error: %v", err)
			}
		}
	}

	return nil
}

func (d *Db) GetUserID(user models.UserAuth) (*int64, error) {
	const (
		getUser = `SELECT id FROM user_register_info WHERE login=$1 and password=$2`
	)
	rows, err := d.db.Query(getUser, user.Login, user.Password)
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
			return nil, fmt.Errorf("[GetUserID] - could not scan rows. error: %v", err)
		}
	}

	return &ID, nil
}
