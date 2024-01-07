package entities

import "gorm.io/gorm"

type BaseModel struct {
	ID        int64          `json:"id" gorm:"column:id;type:bigint;primary_key;auto_increment"`
	CreatedAt string         `json:"created_at" gorm:"column:created_at;type:timestamp;not null;default:now()"`
	UpdatedAt string         `json:"updated_at" gorm:"column:updated_at;type:timestamp;not null;default:now()"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at;type:timestamp"`
	CreatedBy string         `json:"created_by" gorm:"column:created_by"`
	UpdatedBy string         `json:"updated_by" gorm:"column:updated_by"`
}
