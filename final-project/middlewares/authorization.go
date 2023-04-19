package middlewares

import (
	"final-project/database"
	"final-project/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func PhotoAuthorization(c *gin.Context) {
	db := database.GetDB()

	paramId, err := strconv.Atoi(c.Param("photoId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Invalid parameter",
		})
		return
	}

	Photo := models.Photo{}
	err = db.Select("user_id").First(&Photo, uint(paramId)).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":   "Data Not Found",
			"message": "Data doesn't exist",
		})
		return
	}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	if Photo.UserId != userId {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "You are not allowed to access this data",
		})
		return
	}

	c.Next()
}

func CommentAuthorization(c *gin.Context) {
	db := database.GetDB()

	paramId, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Invalid parameter",
		})
		return
	}

	Comment := models.Comment{}
	err = db.Select("user_id").First(&Comment, uint(paramId)).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":   "Data Not Found",
			"message": "Data doesn't exist",
		})
		return
	}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	if Comment.UserId != userId {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "You are not allowed to access this data",
		})
		return
	}

	c.Next()
}

func SocialMediaAuthorization(c *gin.Context) {
	db := database.GetDB()

	paramId, err := strconv.Atoi(c.Param("socialMediaId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Invalid parameter",
		})
		return
	}

	SocialMedia := models.SocialMedia{}
	err = db.Select("user_id").First(&SocialMedia, uint(paramId)).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":   "Data Not Found",
			"message": "Data doesn't exist",
		})
		return
	}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	if SocialMedia.UserId != userId {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "You are not allowed to access this data",
		})
		return
	}

	c.Next()
}
