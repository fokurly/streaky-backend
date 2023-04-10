package service

import (
	"net/http"
	"strings"

	_const "github.com/fokurly/streaky-backend/users_info_api/const"
	"github.com/fokurly/streaky-backend/users_info_api/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func (u *serviceStreakyApi) CheckAuth(ctx *gin.Context) {
	header := ctx.GetHeader(authorizationHeader)
	if header == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, "empty auth header")
		return
	}

	tokenParts := strings.Split(header, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, "empty auth header")
		return
	}

	if len(tokenParts[1]) == 0 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, "empty auth header")
		return
	}

	claims := &models.Claims{}
	tkn, err := jwt.ParseWithClaims(tokenParts[1], claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(_const.JwtKey), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.Next()
}
