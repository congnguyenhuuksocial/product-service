package entities

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID          uint           `gorm:"primaryKey;autoIncrement"`
	Name        string         `gorm:"type:varchar(100)"`
	Description string         `gorm:"type:text"`
	Price       float64        `gorm:"type:decimal(10,2)"`
	SKU         string         `gorm:"type:varchar(30)"`
	Stock       int64          `gorm:"type:int;not null;default:0"`
	CreatedAt   time.Time      `gorm:"type:timestamp;not null;default:now()"`
	UpdatedAt   time.Time      `gorm:"type:timestamp;not null;default:now()"`
	DeletedAt   gorm.DeletedAt `gorm:"type:timestamp"`
}
