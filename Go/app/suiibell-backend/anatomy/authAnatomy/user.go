package authAnatomy

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID                   string         `gorm:"primaryKey"`
	Email                string         `gorm:"type:varchar(100);unique_index"`
	FailedLoginAtttempts int            `gorm:"type:int"`
	FailedStatus         bool           `gorm:"type:bool"`
	CreatedAt            time.Time      `gorm:"default:current_timestamp"`
	UpdatedAt            time.Time      `gorm:"default:current_timestamp"`
	DeletedAt            gorm.DeletedAt `gorm:"index"`
}

type UserCredential struct {
	Email    string `json:"email",gqlgen:"email",required:"true"`
	Password string `json:"password",gqlgen:"password",required:"true"`
}
