package service

import (
	"fmt"
	"net/http"

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

func (u *usersInfoApi) GetAllUserTasksByID(ctx *gin.Context) {

}

func (u *usersInfoApi) GetUserTasksById(ctx *gin.Context) {
	type data struct {
		Id int64 `json:"id"`
	}
	var user data
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Errorf("could not validate fields in body. error: %v", err))
		return
	}

	u.db.GetAllUserTasksByID(user.Id)
}

func (u *usersInfoApi) AddNewTaskToUser(ctx *gin.Context) {
	var task struct {
		UserID         int64  `json:"user_id"`
		FirstObserver  int64  `json:"first_observer_id"`
		SecondObserver int64  `json:"second_observer_id"`
		Name           string `json:"name"`
		Punish         string `json:"punish"`
		StartDate      string `json:"start_date"`
		EndDate        string `json:"end_date"`
		State          string `json:"state"`
	}

	ctx.JSON(http.StatusOK, task)
}

func (u *usersInfoApi) AddNewTaskToWatch(ctx *gin.Context) {

}
