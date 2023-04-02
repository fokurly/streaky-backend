package service

import (
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

func (u *usersInfoApi) GetFriendListByUserID(ctx *gin.Context) {

}

func (u *usersInfoApi) GetAllUserTasksByID(ctx *gin.Context) {

}

func (u *usersInfoApi) GetUserTasks(ctx *gin.Context) {

}

func (u *usersInfoApi) AddNewTaskToUser(ctx *gin.Context) {
	var task struct {
		//ID        int64  `json:"id"`
		UserID    int64  `json:"user_id"`
		Name      string `json:"name"`
		Punish    string `json:"punish"`
		StartDate string `json:"start_date"`
		EndDate   string `json:"end_date"`
		State     string `json:"state"`
	}

	ctx.JSON(http.StatusOK, task)
}

func (u *usersInfoApi) AddNewTaskToWatch(ctx *gin.Context) {

}
