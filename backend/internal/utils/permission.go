package utils

import (
	"errors"
	"time"

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

func GetShareLink(token string) (*models.ShareLink, error) {
    var link models.ShareLink

    if err := database.DB.Where("token = ?", token).First(&link).Error; err != nil {
        return nil, err
    }

    if !link.IsEnabled {
        return nil, errors.New("share link is disabled")
    }

    if link.ExpirationTime != nil && time.Now().After(*link.ExpirationTime) {
        return nil, errors.New("share link expired")
    }

    return &link, nil
}

func CanRead(userID uint, docID uint, token string) (bool, error) {

    isOwner, err := IsOwner(userID, docID)
    if err != nil {
        return false, err
    }
    if isOwner {
        return true, nil
    }

    isCollab, perm, _ := IsCollaborator(userID, docID)
    if isCollab {
        return true, nil
    }

    if token != "" {
        link, err := GetShareLink(token)
        if err != nil {
            return false, nil
        }

        if link.IsPublic && link.AllowRead {
            return true, nil
        }

        if !link.IsPublic && link.AllowRead && userID != 0 {
            return true, nil
        }
    }

    return false, nil
}

func CanEdit(userID uint, docID uint, token string) (bool, error) {

    isOwner, err := IsOwner(userID, docID)
    if err != nil {
        return false, err
    }
    if isOwner {
        return true, nil
    }

    isCollab, perm, _ := IsCollaborator(userID, docID)
    if isCollab && perm == "edit" {
        return true, nil
    }

    if token != "" {
        link, err := GetShareLink(token)
        if err != nil {
            return false, nil
        }

        if link.IsPublic && link.AllowEdit {
            return true, nil
        }

        if !link.IsPublic && link.AllowEdit && userID != 0 {
            return true, nil
        }
    }

    return false, nil
}



