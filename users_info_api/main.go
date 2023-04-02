package main

import (
	"github.com/fokurly/streaky-backend/users_info_api/service"
	"github.com/fokurly/streaky-backend/users_info_api/storage/postgre"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	storage := postgre.NewDatabase()

	//err := storage.InsertNewUser(models.User{Password: "fewf", Email: "ewfew", FullName: "fewf", Login: "fewfew"})
	//if err != nil {
	//	logrus.Panic(err)
	//}
	userApi := service.NewUsersInfoApi(storage)
	router := gin.Default()

	router.POST("/sign_up", userApi.SignUp)
	router.POST("/sign_in", userApi.SignIn)

	apis := router.Group("/api", userApi.CheckAuth)
	{
		apis.POST("/get_all_users_task_by_id", userApi.GetAllUserTasksByID)
		apis.GET("/refresh_token", userApi.RefreshToken)
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
