package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"time"

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

func UpdateDocument(c *gin.Context){

	uid, okk := c.Get("user_id")
	if !okk {
		c.JSON(http.StatusBadRequest, gin.H{
			"message" :"You are not authorized",
		})
		return
	}

	userid := uid.(uint)

	docId := c.Param("id")

	// permission block and pass the userid 

	var input DocumentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : "Invalid input",
		})
		return
	}

	var doc models.Document
	if err := database.DB.First(&doc, docId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message" : "Doc isnt available",
		})
		return
	}

	if doc.OwnerID != userid {
		c.JSON(http.StatusForbidden, gin.H{
			"message" :"You dont have permision to perform write operation",
		})
		return
	}

	doc.Title = input.Title
	doc.Content = input.Content
	
	if err := database.DB.Save(&doc).Error; err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to update document",
		})
     return
  }


	c.JSON(http.StatusOK, gin.H{
		"message" :"You document updates",
		"document" : doc,
	})
}

func generatetoken(n uint) (string, error){
	bytes := make([]byte,n)
	if _,err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil

}

type LinkCreate struct {
	AllowRead  bool `json:allow_read`
	AllowEdit bool `json:allow_edit`
	IsPublic bool `json:is_public`
	ExpirationTime *time.Time
}

func CreateLink(c *gin.Context){
	uid, ok := c.Get("user_id")
	if !ok {
		 c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorized",
		})
    return
	}

	userId := uid.(uint)

	docId := c.Param("id")

	var doc models.Document
	if err := database.DB.First(&doc, docId).Error; err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message" : "document not found",
		})
		return
	}

	if doc.Owner.ID != userId{
		c.JSON(http.StatusForbidden, gin.H{
			"message":"Only Owner can generate token",
		})
		return
	}

	var input LinkCreate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : "invalid share link input",
		})
		return
	}

	token, err := generatetoken(16)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message" : "Not able to generate token",
		})
		return
	}

	share := models.ShareLink{
		Token: token,
		DocumentID: doc.ID,
		AllowRead: input.AllowRead,
		AllowEdit: input.AllowEdit,
		IsPublic: input.IsPublic,
		ExpirationTime: input.ExpirationTime,
		IsEnabled: true,
}

	if err := database.DB.Create(&share).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message" : "Failure to save in database",
		})
		return
	}

	shareLink := "http://localhost:8080/share/" + token

	c.JSON(http.StatusOK, gin.H{
		"message" : "Unique link is created",
		"share_url" : shareLink,
		"settings":  share,
	})
}

