package service

import (
	"fmt"
	"net/http"

	"github.com/fokurly/streaky-backend/users_info_api/models"
	"github.com/fokurly/streaky-backend/users_info_api/storage/postgre"
	"github.com/gin-gonic/gin"
)

type usersInfoApi struct {
	db *postgre.Db
}

func NewUsersInfoApi(db *postgre.Db) *usersInfoApi {
	return &usersInfoApi{db: db}
}

func (u *usersInfoApi) GetUserInfoByUserID(ctx *gin.Context) {

}

func (u *usersInfoApi) GetTaskStateByTaskID(ctx *gin.Context) {

}

func (u *usersInfoApi) GetFriendListByUserID(ctx *gin.Context) {

}

func (u *usersInfoApi) GetAllUserTasksByID(ctx *gin.Context) {

}

func (u *usersInfoApi) GetUserTasks(ctx *gin.Context) {

}

func (u *usersInfoApi) AddNewUser(ctx *gin.Context) {
	var user models.User

	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Errorf("could not validate fields in body. error: %v", err))
		return
	}

	err := u.db.InsertNewUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, fmt.Sprintf("New user %v has been created!", user.FullName))
}

func (u *usersInfoApi) AddNewTaskToUser(ctx *gin.Context) {

}

func (u *usersInfoApi) AddNewTaskToWatch(ctx *gin.Context) {

}
