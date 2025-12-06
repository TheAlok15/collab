package models

import "time"

type Document struct {
	ID            uint `gorm:"primaryKey"`
	OwnerId       uint `gorm:"not null"`
	Owner         User
	Title         string `gorm:"size:255;not null"`
	Content       string `gorm:"type:text"`
	IsDeleted     bool   `gorm:"default:false"`
	Collaborators []DocumentCollaborator
	Events        []DocumentEvent
	CreatedAt     time.Time
	UpdatedAt 		time.Time
}