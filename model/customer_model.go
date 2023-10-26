package model

import "time"

type Customer struct {
	Id          string    `json:"id" binding:"required"`
	FullName    string    `json:"fullname" binding:"required"`
	UserName    string    `json:"username" binding:"required"`
	Password    string    `json:"password,omitempty" binding:"required"`
	PhoneNumber string    `json:"phone_number" binding:"required"`
	CreatedAt   time.Time `json:"created_at"`
}
