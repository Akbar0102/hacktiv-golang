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

func CreateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	SocialMedia := models.SocialMedia{}

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))
	SocialMedia.UserId = userId

	err := db.Debug().Create(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":               SocialMedia.ID,
		"name":             SocialMedia.Name,
		"social_media_url": SocialMedia.SocialMediaUrl,
		"user_id":          SocialMedia.UserId,
		"created_at":       SocialMedia.CreatedAt,
	})
}

func GetAllSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	SocialMedias := []models.SocialMedia{}
	err := db.Where("user_id = ?", userId).Preload("User").Find(&SocialMedias).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": SocialMedias,
	})
}

func GetOneSocialMedia(c *gin.Context) {
	db := database.GetDB()
	paramId, _ := strconv.Atoi(c.Param("socialMediaId"))
	SocialMedia := models.SocialMedia{}

	err := db.First(&SocialMedia, paramId).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": SocialMedia,
	})

}

func UpdateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	paramId, _ := strconv.Atoi(c.Param("socialMediaId"))
	contentType := helpers.GetContentType(c)
	SocialMedia := models.SocialMedia{}

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	err := db.Model(&SocialMedia).Where("id = ?", paramId).Updates(models.SocialMedia{Name: SocialMedia.Name, SocialMediaUrl: SocialMedia.SocialMediaUrl}).Take(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":               SocialMedia.ID,
		"name":             SocialMedia.Name,
		"social_media_urk": SocialMedia.SocialMediaUrl,
		"user_id":          SocialMedia.UserId,
		"updated_at":       SocialMedia.UpdatedAt,
	})

}

func DeleteSocialMedia(c *gin.Context) {
	db := database.GetDB()
	paramId, _ := strconv.Atoi(c.Param("socialMediaId"))
	SocialMedia := models.SocialMedia{}

	err := db.Where("id = ?", paramId).Delete(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your social media has been succesfully deleted",
	})
}
