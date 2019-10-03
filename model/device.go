package model

import (
	"time"
)

// Device 设备表
type Device struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Sno       string `gorm:"primary_key"`
	Mac       string
}
