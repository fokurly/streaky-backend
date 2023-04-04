package main

import (
	"github.com/fokurly/streaky-backend/users_info_api/service"
	"github.com/fokurly/streaky-backend/users_info_api/storage/postgre"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// TODO: подумать о том, как обрабатывать ошибки от бд
// TODO: возможно добавлять в описание ключевую фразу [dbException] и по ней возвращать internal status error. В остальных случаях BadRequest

// TODO: проверить валидацию ctx.Bind и возможно заменить её на что то другое
// TODO: Удалять записи, которые добавили в бд, но не применились в других таблицах
// TODO: апи для получения всей инфы по пользователю
func main() {
	storage := postgre.NewDatabase()

	userApi := service.NewUsersInfoApi(storage)
	router := gin.Default()

	router.POST("/sign_up", userApi.SignUp)
	router.POST("/sign_in", userApi.SignIn)

	usersApi := router.Group("/api", userApi.CheckAuth)
	{
		// TODO: logout
		usersApi.GET("/refresh_token", userApi.RefreshToken)
		usersApi.POST("/check_user_for_exists_by_id", userApi.CheckUserForExistsByID)
		usersApi.POST("/check_user_for_exists_by_login", userApi.CheckUserForExistsByLogin)
		usersApi.POST("/change_password", userApi.UpdateUserPassword)
		usersApi.POST("/send_friend_request", userApi.AddNewFriendRequest)
		usersApi.POST("/accept_friend_request", userApi.AcceptNewFriend)
		usersApi.POST("/get_friend_list", userApi.GetFriendListByUserID)
		usersApi.POST("/get_unconfirmed_friend_list", userApi.GetUnconfirmedFriendsIDs)
		usersApi.POST("/cancel_new_friend", userApi.CancelNewFriendRequest)
	}

	taskApi := router.Group("/api", userApi.CheckAuth)
	{
		taskApi.POST("/create_new_task", userApi.CreateNewTask)
		taskApi.POST("/get_tasks", userApi.GetUserTasks)
		taskApi.POST("/get_observed_tasks", userApi.GetObservedTasks)

		// TODO:
		// taskApi.POST("/update_task_status", userApi.UpdateTaskStatus)
	}

	if err := router.Run(":8080"); err != nil {
		logrus.Panicf("could not run server. error: %v", err)
	}
}

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		PrettyPrint:     true,
	})

	logrus.SetLevel(logrus.DebugLevel)
}
