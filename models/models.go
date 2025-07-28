package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email" gorm:"unique"`
	Role     string `json:"role"`
}

type Song struct {
	ID        int    `json:"id"`
	TITLE     string `json:"title"`
	ALBUM_ID  int    `json:"album_id"`
	FILE_NAME string `json:"file_name"`
}

type Album struct {
	ID     int    `json:"id"`
	TITLE  string `json:"title"`
	ARTIST string `json:"artist"`
	COVER  string `json:"cover"`
}
