package postgre

import (
	"fmt"

	"github.com/fokurly/streaky-backend/users_info_api/models"
	"github.com/lib/pq"
)

// TODO: удаление пользователей из друзей

func (d *Db) GetUserByID(ID int64) (*models.UserInfo, error) {
	const (
		getUser = `SELECT id, login FROM user_register_info WHERE id=$1`
	)
	rows, err := d.db.Query(getUser, ID)
	if err != nil {
		return nil, fmt.Errorf("[GetUserByID] - could not exec query. error: %v", err)
	}

	defer func() {
		_ = rows.Close()
	}()

	var user models.UserInfo
	if rows.Next() {
		err := rows.Scan(&user.ID, &user.Login)
		if err != nil {
			return nil, fmt.Errorf("[GetUserByID] - could not scan rows. error: %v", err)
		}
	}
	// Костыль, надо как то убрать
	if len(user.Login) == 0 {
		return nil, nil
	}
	return &user, nil
}

func (d *Db) GetUserByLogin(ID string) (*models.UserInfo, error) {
	const (
		getUser = `SELECT id, login FROM user_register_info WHERE login=$1`
	)
	rows, err := d.db.Query(getUser, ID)
	if err != nil {
		return nil, fmt.Errorf("[GetUserByLogin] - could not exec query. error: %v", err)
	}

	defer func() {
		_ = rows.Close()
	}()

	var user models.UserInfo
	if rows.Next() {
		err := rows.Scan(&user.ID, &user.Login)
		if err != nil {
			return nil, fmt.Errorf("[GetUserByLogin] - could not scan rows. error: %v", err)
		}
	}

	// Костыль, надо как то убрать
	if len(user.Login) == 0 {
		return nil, nil
	}

	return &user, nil
}

func (d *Db) UpdateUserPassword(user models.UserAuth) error {
	const (
		updatePassword = `UPDATE user_register_info SET password=$1 WHERE login=$2`
	)

	_, err := d.db.Exec(updatePassword, user.Password, user.Login)
	if err != nil {
		return fmt.Errorf("[UpdateUserPassword] - could not update password. error: %v", err)
	}

	return nil
}

func (d *Db) GetFriendListByUserID(ID int64) ([]int64, error) {
	const (
		selectFriendsList = `SELECT friends_ids FROM user_friend_list where userid=$1`
	)

	rows, err := d.db.Query(selectFriendsList, ID)
	if err != nil {
		return nil, fmt.Errorf("[GetFriendListByUserID] - could not exec query. error: %v", err)
	}

	var friendList []int64
	var tmp pq.Int64Array
	if rows.Next() {
		err := rows.Scan(&tmp)
		if err != nil {
			return nil, fmt.Errorf("[GetFriendListByUserID] - could not scan rows. error: %v", err)
		}
	}

	if tmp != nil {
		friendList = tmp
	}

	return friendList, nil
}

func (d *Db) GetUnconfirmedFriendListByUserID(ID int64) ([]int64, error) {
	const (
		selectFriendsList = `SELECT unconfirmed_friends_ids FROM user_friend_list where userid=$1`
	)

	rows, err := d.db.Query(selectFriendsList, ID)
	if err != nil {
		return nil, fmt.Errorf("[GetUnconfirmedFriendListByUserID] - could not exec query. error: %v", err)
	}

	var friendList []int64
	var tmp pq.Int64Array
	if rows.Next() {
		err := rows.Scan(&tmp)
		if err != nil {
			return nil, fmt.Errorf("[GetUnconfirmedFriendListByUserID] - could not scan rows. error: %v", err)
		}
	}
	if tmp != nil {
		friendList = tmp
	}

	return friendList, nil
}

// Переделать именно под update, а не insert новый
func (d *Db) AddNewFriendToUnconfirmed(userID, newFriendID int64) error {
	const (
		updateUnconfirmedFriendList = `UPDATE user_friend_list SET unconfirmed_friends_ids=$1 WHERE userid=$2`
	)

	// Добавить проверку на уникальность
	friendList, err := d.GetFriendListByUserID(newFriendID)
	for i := range friendList {
		if friendList[i] == userID {
			return fmt.Errorf("user is already in your friend list")
		}
	}
	// Получаем друзей пользователя, которого хотят добавить
	unconfirmedFriendList, err := d.GetUnconfirmedFriendListByUserID(newFriendID)
	if err != nil {
		return fmt.Errorf("could not get unconfirmed user friends. error: %v", err)
	}

	// Добавляем к пользователю, которого хотят добавить айди пользователя, который хочет его добавить
	unconfirmedFriendList = append(unconfirmedFriendList, userID)
	var tmp pq.Int64Array
	tmp = unconfirmedFriendList
	_, err = d.db.Exec(updateUnconfirmedFriendList, tmp, newFriendID)
	if err != nil {
		return fmt.Errorf("[AddNewFriendToUnconfirmed] - could not exec query. error: %v", err)
	}

	return nil
}

func (d *Db) AddNewFriendToConfirmed(userID, newFriendID int64) error {
	const (
		updateConfirmedFriendList = `UPDATE user_friend_list SET friends_ids=$1 WHERE userid=$2`
	)

	// Добавить проверку на уникальность
	// Получаем друзей пользователя, которого хотят добавить
	confirmedFriendList, err := d.GetFriendListByUserID(userID)
	if err != nil {
		return fmt.Errorf("could not get confirmed user friends. error: %v", err)
	}
	// Добавить проверку на дупликацию

	for i := range confirmedFriendList {
		if confirmedFriendList[i] == newFriendID {
			return fmt.Errorf("user is already in your friend list")
		}
	}

	// Добавляем к пользователю, которого хотят добавить айди пользователя, который хочет его добавить
	confirmedFriendList = append(confirmedFriendList, newFriendID)
	var tmp pq.Int64Array
	tmp = confirmedFriendList
	_, err = d.db.Exec(updateConfirmedFriendList, tmp, userID)
	if err != nil {
		return fmt.Errorf("[AddNewFriendToUnconfirmed] - could not exec query. error: %v", err)
	}

	return nil
}

func (d *Db) AcceptFriend(userID, newFriendID int64) error {
	err := d.AddNewFriendToConfirmed(userID, newFriendID)
	if err != nil {
		return fmt.Errorf("could not confirm friend request. %v", err)
	}
	err = d.AddNewFriendToConfirmed(newFriendID, userID)
	if err != nil {
		return fmt.Errorf("could not confirm friend request. %v", err)
	}
	err = d.DeleteFromUnconfirmedFriendList(userID, newFriendID)
	if err != nil {
		return fmt.Errorf("could not delete friend from unconfirmed list")
	}

	return nil
}

func (d *Db) DeleteFromUnconfirmedFriendList(userID, deleteUserID int64) error {
	const (
		deleteIDFromUnconfirmedFriendList = `UPDATE user_friend_list SET unconfirmed_friends_ids=$1 where userid=$2`
	)
	list, err := d.GetUnconfirmedFriendListByUserID(userID)
	if err != nil {
		return fmt.Errorf("could not get unconfirmed friend list. error: %v", err)
	}

	for i := range list {
		if list[i] == deleteUserID {
			list[i] = list[len(list)-1] // Copy last element to index i.
			list[len(list)-1] = -1      // Erase last element (write zero value).
			list = list[:len(list)-1]   // Truncate slice.
			break
		}
	}

	var tmp pq.Int64Array
	tmp = list
	_, err = d.db.Exec(deleteIDFromUnconfirmedFriendList, tmp, userID)
	if err != nil {
		return fmt.Errorf("could not exec query. error: %v", err)
	}

	return nil
}
