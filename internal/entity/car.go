package entity

import (
	"time"

	"gorm.io/gorm"
)

type Car struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Brand     string         `json:"brand"`
	Color     string         `json:"color"`
	Seats     uint8          `json:"seats"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
