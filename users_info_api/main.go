package main

import (
	"github.com/fokurly/streaky-backend/users_info_api/service"
	"github.com/fokurly/streaky-backend/users_info_api/storage/postgre"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	storage := postgre.NewDatabase()

	serviceApi := service.NewStreakyApi(storage)
	router := gin.Default()

	router.POST("/sign_up", serviceApi.SignUp)
	router.POST("/sign_in", serviceApi.SignIn)

	router.GET("/ping", serviceApi.Ping)
	usersApi := router.Group("/api", serviceApi.CheckAuth)
	{
		// TODO: logout
		usersApi.GET("/refresh_token", serviceApi.RefreshToken)
		usersApi.POST("/check_user_for_exists_by_id", serviceApi.CheckUserForExistsByID)
		usersApi.POST("/check_user_for_exists_by_login", serviceApi.CheckUserForExistsByLogin)
		usersApi.POST("/change_password", serviceApi.UpdateUserPassword)
		usersApi.POST("/send_friend_request", serviceApi.AddNewFriendRequest)
		usersApi.POST("/accept_friend_request", serviceApi.AcceptNewFriend)
		usersApi.POST("/get_friend_list", serviceApi.GetFriendListByUserID)
		usersApi.POST("/get_unconfirmed_friend_list", serviceApi.GetUnconfirmedFriendsIDs)
		usersApi.POST("/cancel_new_friend", serviceApi.CancelNewFriendRequest)
		usersApi.POST("/get_random_user", serviceApi.GetRandomUser)

		usersApi.POST("/get_notify", serviceApi.GetUserNotifications)
		usersApi.POST("/send_notify", serviceApi.SendNotification)

		usersApi.POST("/get_user_info", serviceApi.GetUserInfo)

		usersApi.POST("/update_observer_day", serviceApi.UpdateDayForObserver)
		usersApi.POST("/update_user_day", serviceApi.UpdateDayForUser)
		usersApi.POST("/get_days", serviceApi.GetDays)
		usersApi.POST("/get_current_day", serviceApi.GetCurrentDay)
	}

	taskApi := router.Group("/api", serviceApi.CheckAuth)
	{
		taskApi.POST("/create_new_task", serviceApi.CreateNewTask)
		taskApi.POST("/get_tasks", serviceApi.GetUserTasks)
		taskApi.POST("/get_observed_tasks", serviceApi.GetObservedTasks)

		// TODO:
		taskApi.POST("/update_task_status", serviceApi.UpdateTaskStatus)
		taskApi.POST("/get_task_by_id", serviceApi.GetTaskInfoByID)
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
