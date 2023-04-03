package service

import (
	"fmt"
	"net/http"

	"github.com/fokurly/streaky-backend/users_info_api/models"
	"github.com/fokurly/streaky-backend/users_info_api/utils"
	"github.com/gin-gonic/gin"
)

func (u *usersInfoApi) CheckUserForExistsByID(ctx *gin.Context) {
	type data struct {
		Id int64 `json:"id"`
	}
	var user data
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("could not validate fields in body. error: %v", err))
		return
	}

	userInfo, err := u.db.GetUserByID(user.Id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, fmt.Sprintf("could not get data from db. error: %v", err))
		return
	}

	if userInfo == nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("there is no user with such id!"))
		return
	}

	ctx.JSON(http.StatusOK, userInfo)
}

func (u *usersInfoApi) CheckUserForExistsByLogin(ctx *gin.Context) {
	type data struct {
		Login string `json:"login"`
	}

	var user data
	if err := ctx.BindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("could not validate fields in body. error: %v", err))
		return
	}

	userInfo, err := u.db.GetUserByLogin(user.Login)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, fmt.Sprintf("could not get data from db. error: %v", err))
		return
	}

	if userInfo == nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("there is no user with such login!"))
		return
	}

	ctx.JSON(http.StatusOK, userInfo)
}

// Добавить в случае если забыл пароль метод обновления??
func (u *usersInfoApi) UpdateUserPassword(ctx *gin.Context) {
	type data struct {
		Auth        models.UserAuth `json:"user_auth"`
		NewPassword string          `json:"new_password"`
	}

	var user data
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("could not validate fields in body. error: %v", err))
		return
	}

	user.Auth.Password = utils.HashPassword(user.Auth.Password)
	_, err := u.db.GetUserID(user.Auth)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Sprintf("could not get such user from db. error: %v", err))
		return
	}

	user.Auth.Password = utils.HashPassword(user.NewPassword)
	err = u.db.UpdateUserPassword(user.Auth)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, fmt.Sprintf("could not update password. error: %v", err))
		return
	}

	ctx.JSON(http.StatusOK, fmt.Sprintf("your password has been changed."))
}

// Принимает id
func (u *usersInfoApi) AddNewFriendRequest(ctx *gin.Context) {
	type params struct {
		UserID      int64 `json:"user_id"`
		NewFriendID int64 `json:"new_friend_id"`
	}
	var data params
	if err := ctx.BindJSON(&data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("could not validate fields in body. error: %v", err))
		return
	}

	err := u.db.AddNewFriendToUnconfirmed(data.UserID, data.NewFriendID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, fmt.Sprintf("could not send friend request. error: %v", err))
		return
	}

	ctx.JSON(http.StatusOK, fmt.Sprintf("You send request to new friend!"))
}

func (u *usersInfoApi) AcceptNewFriend(ctx *gin.Context) {
	type params struct {
		UserID      int64 `json:"user_id"`
		NewFriendID int64 `json:"new_friend_id"`
	}
	var data params
	if err := ctx.BindJSON(&data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("could not validate fields in body. error: %v", err))
		return
	}

	err := u.db.AcceptFriend(data.UserID, data.NewFriendID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, fmt.Sprintf("could not accept friend request. error: %v", err))
		return
	}

	ctx.JSON(http.StatusOK, fmt.Sprintf("You accepted friend request!"))
}

func (u *usersInfoApi) GetFriendListByUserID(ctx *gin.Context) {
	type params struct {
		UserID int64 `json:"user_id"`
	}
	var data params
	if err := ctx.BindJSON(&data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("could not validate fields in body. error: %v", err))
		return
	}

	friendList, err := u.db.GetFriendListByUserID(data.UserID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, fmt.Sprintf("could not get friend list. error"))
		return
	}

	if friendList == nil {
		ctx.AbortWithStatusJSON(http.StatusNoContent, fmt.Sprintf("there are no friends in your friend list!"))
		return
	}

	ctx.JSON(http.StatusOK, friendList)
}

// Метод в который будут тыкать и он будет отдавать запросы в друзьям конкретному id пользователя
func (u *usersInfoApi) GetUnconfirmedFriendsIDs(ctx *gin.Context) {
	type params struct {
		UserID int64 `json:"user_id"`
	}
	var data params
	if err := ctx.BindJSON(&data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("could not validate fields in body. error: %v", err))
		return
	}

	unconfirmedFriendList, err := u.db.GetUnconfirmedFriendListByUserID(data.UserID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, fmt.Sprintf("could not get friend list. error"))
		return
	}

	if unconfirmedFriendList == nil {
		ctx.AbortWithStatusJSON(http.StatusOK, fmt.Sprintf("there is no unconfirmed friends in your list!"))
		return
	}

	ctx.JSON(http.StatusOK, unconfirmedFriendList)
}

func (u *usersInfoApi) CancelNewFriendRequest(ctx *gin.Context) {
	type params struct {
		UserID         int64 `json:"user_id"`
		CancelFriendID int64 `json:"cancel_friend_id"`
	}
	var data params
	if err := ctx.BindJSON(&data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("could not validate fields in body. error: %v", err))
		return
	}

	err := u.db.DeleteFromUnconfirmedFriendList(data.UserID, data.CancelFriendID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, fmt.Sprintf("could not cancel friend request. error: %v", err))
		return
	}

	ctx.JSON(http.StatusOK, fmt.Sprintf("You have cancelled friend request!"))
}

func (u *usersInfoApi) DeleteFriend(ctx *gin.Context) {

}