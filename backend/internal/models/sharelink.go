package models

import "time"

type ShareLink struct {
	ID             uint   `gorm:"primary key"`
	DocumentID     uint   `gorm:"not null;index"`
	Token          string `gorm:"size:64;uniqueIndex;not null"`
	IsPublic       bool   `gorm:"default:false"`
	AllowRead      bool   `gorm:"default:true"`
	AllowEdit      bool   `gorm:"default:false"`
	IsEnabled      bool   `gorm:"deafult:true"`
	ExpirationTime *time.Time
	CreatedAt      time.Time
}
