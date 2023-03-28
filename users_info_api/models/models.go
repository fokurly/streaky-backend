package models

import "github.com/dgrijalva/jwt-go/v4"

type User struct {
	ID       int64  `json:"id"`
	Login    string `json:"login"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Task struct {
	ID        int64  `json:"id"`
	UserID    int64  `json:"user_id"`
	Name      string `json:"name"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	State     string `json:"state"`
}

type FriendList struct {
	Users []User `json:"users"`
}

type DatabaseConfig struct {
	User     string `json:"user" validate:"required"`
	Password string `json:"password" validate:"required"`
	Host     string `json:"host" validate:"required"`
	Dbname   string `json:"dbname" validate:"required"`
}

type Claims struct {
	jwt.StandardClaims
	UserName string `json:"user_name"`
}

type UserAuth struct {
	Username string `json:"username" bson:"_id"`
	Password string `json:"password" bson:"password"`
}
