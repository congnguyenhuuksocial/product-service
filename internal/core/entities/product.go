package entities

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	gorm.Model              // This adds ID, CreatedAt, UpdatedAt, DeletedAt fields
	Name        string      `gorm:"type:varchar(100);not null"`
	Description string      `gorm:"type:text;not null"`
	Price       float64     `gorm:"type:decimal(10,2);not null"`
	SKU         string      `gorm:"type:varchar(30);not null;unique"`
	Stock       int         `gorm:"type:int;not null;default:0"`
	Categories  []Category  `gorm:"many2many:product_categories;"` // Association to a Category struct
	Images      []Image     `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Attributes  []Attribute `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Ratings     []Rating    `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Category struct {
	ID   string `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(100);not null;unique"`
}

type Image struct {
	gorm.Model
	URL       string `gorm:"type:varchar(255);not null"`
	Alt       string `gorm:"type:varchar(255)"`
	IsMain    bool   `gorm:"type:boolean;not null;default:false"`
	ProductID uint   `gorm:"type:bigint;not null"`
}

type Attribute struct {
	gorm.Model
	Key       string `gorm:"type:varchar(50);not null"`
	Value     string `gorm:"type:varchar(100);not null"`
	ProductID uint   `gorm:"type:bigint;not null"`
}

type Rating struct {
	gorm.Model
	UserID    string    `gorm:"type:uuid;not null"`
	ProductID uint      `gorm:"type:bigint;not null"`
	Rating    float32   `gorm:"type:decimal(3,2);not null"`
	Comment   string    `gorm:"type:text"`
	Timestamp time.Time `gorm:"type:timestamp"`
}
