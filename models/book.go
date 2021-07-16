package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID          uint64  `json:"id" gorm:"primaryKey"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	MediumPrice float32 `json:"medium_price"`
	Author      string  `json:"author"`
	ImageURL    string  `json:"image_url"`

	CreatedAt time.Time      `json:"created"`
	UpdatedAt time.Time      `json:"updated"`
	DeletedAt gorm.DeletedAt `json:"deleted"`
}
