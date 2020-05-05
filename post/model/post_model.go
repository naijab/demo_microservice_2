package model

type Post struct {
	Id     int    `gorm:"primary_key" json:"id"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}
