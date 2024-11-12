package models

type User struct {
	ID int `json: "id" gorm:"primaryKey;autoIncrement"`
	Name string `json: "name"`
	Username string `json: "username"`
	Password string `json: "password"`
}