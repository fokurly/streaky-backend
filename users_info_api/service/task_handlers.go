package service

import (
	"fmt"
	"net/http"

	"github.com/fokurly/streaky-backend/users_info_api/models"
	"github.com/gin-gonic/gin"
)

func (u *usersInfoApi) CreateNewTask(ctx *gin.Context) {
	var task models.Task

	if err := ctx.BindJSON(&task); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("validation body error: %v", err))
		return
	}

	id, err := u.db.InsertNewTask(task)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, fmt.Sprintf("could not create new task. error: %v", err))
		return
	}

	resp := &struct {
		ID int64 `json:"task_id"`
	}{
		ID: *id,
	}
	ctx.JSON(http.StatusOK, resp)
}

func (u *usersInfoApi) GetUserTasks(ctx *gin.Context) {
	var user struct {
		ID int64 `json:"user_id"`
	}

	if err := ctx.BindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("validation body error: %v", err))
		return
	}

	tasks, err := u.db.GetUserTasks(user.ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, fmt.Sprintf("could not get tasks. error: %v", err))
		return
	}

	if tasks == nil {
		ctx.AbortWithStatus(http.StatusNoContent)
		return
	}

	ctx.JSON(http.StatusOK, tasks)
}

func (u *usersInfoApi) GetObservedTasks(ctx *gin.Context) {
	var user struct {
		ID int64 `json:"user_id"`
	}

	if err := ctx.BindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("validation body error: %v", err))
		return
	}

	tasks, err := u.db.GetObservedTasks(user.ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, fmt.Sprintf("could not get tasks. error: %v", err))
		return
	}

	if tasks == nil {
		ctx.AbortWithStatus(http.StatusNoContent)
		return
	}

	ctx.JSON(http.StatusOK, tasks)
}

func (u *usersInfoApi) UpdateTaskStatus(ctx *gin.Context) {
	var params struct {
		Status string `json:"status"`
		TaskID int64  `json:"task_id"`
	}

	if err := ctx.BindJSON(&params); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("could not validate body. err: %v", err))
		return
	}

	err := u.db.UpdateTaskStatus(params.Status, params.TaskID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, fmt.Sprintf("could not update status. error: %v", err))
		return
	}

	ctx.Status(http.StatusOK)
}

func (u *usersInfoApi) GetTaskInfoByID(ctx *gin.Context) {
	var params struct {
		ID int64 `json:"task_id"`
	}

	if err := ctx.BindJSON(&params); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("could not validate body. err: %v", err))
		return
	}

	taskInfo, err := u.db.GetTaskInfo(params.ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, fmt.Sprintf("could not get task. error: %v", err))
		return
	}

	ctx.JSON(http.StatusOK, taskInfo)
}

func (u *usersInfoApi) GetDays(ctx *gin.Context) {
	type params struct {
		TaskID int64 `json:"task_id"`
	}

	var data params
	if err := ctx.BindJSON(&data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("could not validate body. err: %v", err))
		return
	}

	days, err := u.db.GetDays(data.TaskID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, fmt.Sprintf("could not get days. error: %v", err))
		return
	}

	ctx.JSON(http.StatusOK, days)
}

func (u *usersInfoApi) GetCurrentDay(ctx *gin.Context) {
	type params struct {
		TaskID int64  `json:"task_id"`
		Day    string `json:"day"`
	}

	var data params
	if err := ctx.BindJSON(&data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("could not validate body. err: %v", err))
		return
	}

	days, err := u.db.GetCurrentDay(data.TaskID, data.Day)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, fmt.Sprintf("could not get day. error: %v", err))
		return
	}

	ctx.JSON(http.StatusOK, days)
}
