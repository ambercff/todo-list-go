package models

import "time"

type Task struct {
	ID          int    `json: "id" gorm:"primaryKey;autoIncrement"`
	Title       string `json:"title"`
	Description string `json: "description"`
	Status      string `json: "status"`
	CreatedAt    time.Time `json: "createdAt"`
}