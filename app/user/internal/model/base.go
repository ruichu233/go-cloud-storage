package model

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        uint           `gorm:"primaryKey" json:"id,string"`
	CreatedAt time.Time      `json:"created_at,string"`
	UpdatedAt time.Time      `json:"updated_at,string"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,string"`
}
