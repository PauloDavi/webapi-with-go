package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID					uint 						`json:"id" gorm:"primaryKey"`
	Name				string 					`json:"name"`
	Description string 					`json:"description"`
	MediunPrice float32 				`json:"mediun_price"`
	Author 			string 					`json:"author"`
	ImageURL 		string 					`json:"image_url"`
	CreatedAt 	time.Time 			`json:"created"`
	UpdatedAt 	time.Time 			`json:"updated"`
	DeletedAt 	gorm.DeletedAt 	`gorm:"index" json:"deleted"`
}