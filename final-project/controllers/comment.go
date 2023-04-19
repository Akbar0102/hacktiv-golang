package controllers

import (
	"final-project/database"
	"final-project/helpers"
	"final-project/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	Comment := models.Comment{}

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))
	Comment.UserId = userId

	err := db.Debug().Create(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         Comment.ID,
		"message":    Comment.Message,
		"photo_id":   Comment.PhotoId,
		"user_id":    Comment.UserId,
		"created_at": Comment.CreatedAt,
	})
}

func GetAllComments(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))
	
	Comments := []models.Comment{}
	err := db.Where("user_id = ?", userId).Preload("User").Preload("Photo").Find(&Comments).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": Comments,
	})
}

func GetOneComment(c *gin.Context) {
	db := database.GetDB()
	paramId, _ := strconv.Atoi(c.Param("commentId"))
	Comment := models.Comment{}

	err := db.First(&Comment, paramId).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": Comment,
	})
}

func UpdateComment(c *gin.Context) {
	db := database.GetDB()
	paramId, _ := strconv.Atoi(c.Param("commentId"))
	contentType := helpers.GetContentType(c)
	Comment := models.Comment{}

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	err := db.Model(&Comment).Where("id = ?", paramId).Updates(models.Comment{Message: Comment.Message}).Take(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         Comment.ID,
		"message":    Comment.Message,
		"photo_id":   Comment.PhotoId,
		"user_id":    Comment.UserId,
		"updated_at": Comment.UpdatedAt,
	})
}

func DeleteComment(c *gin.Context) {
	db := database.GetDB()
	paramId, _ := strconv.Atoi(c.Param("commentId"))
	Comment := models.Comment{}

	err := db.Where("id = ?", paramId).Delete(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your comment has been succesfully deleted",
	})
}
