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

	err := u.db.InsertNewTask(task)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, fmt.Sprintf("could not create new task. error: %v", err))
		return
	}

	ctx.JSON(http.StatusOK, fmt.Sprintf("task has been created!"))
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
	}
}
