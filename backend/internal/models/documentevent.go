package models

import "time"

type DocumentEvent struct {
	ID         uint   `gorm:"primary key"`
	UserID     uint   `gorm:"not null"`
	DocumentID uint   `gorm:"not null"`
	Payload    string `gorm:"type:json"`
	CreatedAt  time.Time
}
