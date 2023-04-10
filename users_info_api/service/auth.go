package service

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	_const "github.com/fokurly/streaky-backend/users_info_api/const"
	"github.com/fokurly/streaky-backend/users_info_api/models"
	"github.com/fokurly/streaky-backend/users_info_api/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

const (
	authorizationHeader = "Authorization"
)

func (u *serviceStreakyApi) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, fmt.Sprintf("ping"))
}

func (u *serviceStreakyApi) SignUp(ctx *gin.Context) {
	var user models.User

	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Errorf("could not validate fields in body. error: %v", err))
		return
	}

	user.Password = utils.HashPassword(user.Password)
	err := u.db.InsertNewUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Sprintf("Error: %v", err))
		return
	}

	// Добавить ошибку о том, что такой юзер уже есть

	ctx.JSON(http.StatusOK, fmt.Sprintf("New user %v has been created!", user.FullName))
}

func (u *serviceStreakyApi) SignIn(ctx *gin.Context) {
	var user models.UserAuth

	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("could not validate fields in body. error: %v", err))
		return
	}
	user.Password = utils.HashPassword(user.Password)
	id, err := u.db.GetUserID(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, fmt.Sprintf("could not get user from db. error: %s", err))
		return
	}

	if id == nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("could not find user with such login and password. try again."))
		return
	}

	expirationTime := time.Now().Add(60 * time.Minute)
	claims := &models.Claims{
		Username: user.Login,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(_const.JwtKey))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, fmt.Sprintf("could not generate bearer token"))
		return
	}

	// заменить hello message на id
	resp := &struct {
		Token string `json:"token"`
		User  string `json:"user"`
		ID    int64  `json:"id"`
	}{
		Token: tokenString,
		User:  user.Login,
		ID:    *id,
	}

	ctx.Header(authorizationHeader, tokenString)
	ctx.JSON(http.StatusOK, resp)
}

func (u *serviceStreakyApi) RefreshToken(ctx *gin.Context) {
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

	// TODO: Написать кэш для хранения валидных токенов, которые после логаута удаляются

	expirationTime := time.Now().Add(60 * time.Minute)
	claims = &models.Claims{
		Username: claims.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(_const.JwtKey))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, fmt.Errorf("could not generate bearer token"))
		return
	}

	resp := &struct {
		Token string `json:"token"`
	}{
		Token: tokenString,
	}

	ctx.Header(authorizationHeader, tokenString)
	ctx.JSON(http.StatusOK, resp)
}
