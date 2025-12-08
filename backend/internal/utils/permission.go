package utils

import (

	"github.com/TheAlok15/collab/internal/database"
	"github.com/TheAlok15/collab/internal/models"
)

func IsOwner(userID uint, docID uint) (bool, error){

	var doc models.Document

	if err := database.DB.First(&doc, docID).Error; err != nil {
		return false, err
	}
	
	if doc.OwnerID == userID{
		return true, nil
	}

	return false, nil
}

func IsCollaborator(userID uint, docID uint) (bool,string, error){

	var collab models.DocumentCollaborator
	if err := database.DB.Where("user_id = ? AND doc_id = ?", userID, docID).First(&collab).Error; err != nil {
		return false,"", err
	}

	return true, collab.Permission, nil
}