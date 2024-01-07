package entities

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID          uint           `gorm:"primaryKey;autoIncrement"`
	Name        string         `gorm:"type:varchar(100)"`
	Description string         `gorm:"type:text"`
	Price       float64        `gorm:"type:decimal(10,2)"`
	SKU         string         `gorm:"type:varchar(30)"`
	Stock       int            `gorm:"type:int;not null;default:0"`
	Categories  datatypes.JSON `gorm:"type:json"`
	Images      datatypes.JSON `gorm:"type:json"`
	Attributes  datatypes.JSON `gorm:"type:json"`
	Ratings     datatypes.JSON `gorm:"type:json"`
	CreatedAt   time.Time      `gorm:"type:timestamp;not null;default:now()"`
	UpdatedAt   time.Time      `gorm:"type:timestamp;not null;default:now()"`
	DeletedAt   gorm.DeletedAt `gorm:"type:timestamp"`
}

type Category struct {
	Name string `json:"Name"`
}

type Image struct {
	URL    string `json:"URL"`
	Alt    string `json:"Alt"`
	IsMain bool   `json:"IsMain"`
}

type Attribute struct {
	Key   string `json:"Key"`
	Value string `json:"Value"`
}

type Rating struct {
	UserID    string    `json:"UserID"`
	ProductID uint      `json:"ProductID"`
	Rating    float32   `json:"Rating"`
	Comment   string    `json:"Comment"`
	Timestamp time.Time `json:"Timestamp"`
}
