package handlers

import (
	"net/http"

	"github.com/TheAlok15/collab/internal/database"
	"github.com/TheAlok15/collab/internal/models"
	"github.com/gin-gonic/gin"
)

type DocumentInput struct{
	Title string
	Content string
}

func CreateDocument(c *gin.Context){

	uid, okk := c.Get("user_id")
	if !okk {
		c.JSON(http.StatusUnauthorized, gin.H{"message":"You are not authorize"})
		return
	}
	userID := uid.(uint)

	var input DocumentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message" : "Invalid document input"})
		return

	}

	doc := models.Document{
		Title: input.Title,
		Content: input.Content,
		OwnerID: userID,
	}

	if err := database.DB.Create(&doc).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not save document"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message" : "Doc created successfully",
		"document" : doc,
	})


}

func GetAllDocuments(c *gin.Context){

	uid, okk := c.Get("user_id")
	if !okk {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message" : "You are not authorized",
		})
		return
	}

	userId := uid.(uint)

	var alldocument []models.Document
	if err := database.DB.Where("owner_id = ? AND is_deleted = false", userId).Find(&alldocument).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : "Internal server error",
		})
		return
	}

	if len(alldocument) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message" : "No docs created yet",
			"document" : []models.Document{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"documents": alldocument})


}

func GetAllCollaborateDoc(c *gin.Context){

	uid, okk := c.Get("user_id")
	if !okk {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message" :"You are not authorized",
		})
		return
	}

	userId := uid.(uint)

	var collabdocs []models.DocumentCollaborator
	if err := database.DB.Where("user_id = ? ", userId).Find(&collabdocs).Error; err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : "Collaborators query failed, Internal server error",
		})
		return
	}

	var docIds []uint
	for _ , col := range collabdocs {
		docIds = append(docIds, col.DocumentID)
	}

	if len(docIds) == 0{
		c.JSON(http.StatusOK, gin.H{
			"message" :"You are not a collaborator on any document",
			"documents" : []models.Document{},
		})
		return
	}

	var docs []models.Document
	if err := database.DB.Where("id IN ? AND is_deleted = false", docIds).Find(&docs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to fetch documents",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message" : "Collaborator documents fetched successfully",
		"documents" : docs,
	})
}


