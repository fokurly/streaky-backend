package models

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/lib/pq"
)

type User struct {
	ID       int64  `json:"id"`
	Login    string `json:"login"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserInfo struct {
	ID    int64  `json:"id"`
	Login string `json:"login"`
}

type UserAuth struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Task struct {
	ID             int64          `json:"id"`
	UserID         int64          `json:"user_id"`
	FirstObserver  int64          `json:"first_observer" binding:"required"`
	SecondObserver int64          `json:"second_observer"`
	Name           string         `json:"name"`
	Punish         string         `json:"punish"`
	StartDate      string         `json:"start_date"`
	EndDate        string         `json:"end_date"`
	State          string         `json:"state"`
	Description    string         `json:"description"`
	FrequenctPQ    pq.StringArray `json:"frequencyperiod"`
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
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type Notification struct {
	FromUserID int64  `json:"from_user_id"`
	ToUserID   int64  `json:"to_user_id"`
	Message    string `json:"message"`
}
