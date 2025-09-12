package models

import "time"

type UsersModel struct {
	ID       int `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Rules    string `json:"rules"`
	CreateAt time.Time
}