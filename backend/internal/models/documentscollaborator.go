package models

import "time"

type DocumentCollaborator struct {
	ID         uint `gorm:"primary key"`
	DocumentID uint `gorm:"not null"`
	Document   Document
	UserID     uint `gorm:"not null"`
	User       User
	Permission string `gorm:"size:20;default:'edit'"`
	CreatedAt  time.Time
}
