package model

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	CreateAT  time.Time       `gorm:"column:created_at;defauLt:CURRENT_TIMESTAMP" json:"create_at"`
	UpdateAT  time.Time       `gorm:"column:created_at;defauLt:CURRENT_TIMESTAMP" json:"update_at"`
	DeletedAT *gorm.DeletedAt `json:"deleted_at,omitempty"`
}
