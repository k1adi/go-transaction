package model

type Bank struct {
	Id   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}
