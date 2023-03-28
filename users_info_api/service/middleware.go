package service

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (u *usersInfoApi) authMiddleware(ctx *gin.Context) {
	const (
		login    = "client"
		password = "clientPassword"
	)
	log, pass, _ := ctx.Request.BasicAuth()
	if log != login || pass != password {
		err := fmt.Errorf("authorization error")
		logrus.Warn(err)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
		return
	}
	ctx.Next()

}
