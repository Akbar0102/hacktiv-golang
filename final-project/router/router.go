package router

import (
	"final-project/controllers"
	"final-project/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}

	photosRouter := r.Group("/photos")
	{
		photosRouter.Use(middlewares.Authentication())

		photosRouter.POST("/", controllers.CreatePhoto)
		photosRouter.GET("/", controllers.GetAllPhotos)
		photosRouter.GET("/:photoId", controllers.GetOnePhoto)
		photosRouter.PUT("/:photoId", middlewares.PhotoAuthorization, controllers.UpdatePhoto)
		photosRouter.DELETE("/:photoId", middlewares.PhotoAuthorization, controllers.DeletePhoto)
	}

	commentsRouter := r.Group("/comments")
	{
		commentsRouter.Use(middlewares.Authentication())

		commentsRouter.POST("/", controllers.CreateComment)
		commentsRouter.GET("/", controllers.GetAllComments)
		commentsRouter.GET("/:commentId", controllers.GetOneComment)
		commentsRouter.PUT("/:commentId", middlewares.CommentAuthorization, controllers.UpdateComment)
		commentsRouter.DELETE("/:commentId", middlewares.CommentAuthorization, controllers.DeleteComment)
	}

	socialMediasRouter := r.Group("/socialmedias")
	{
		socialMediasRouter.Use(middlewares.Authentication())

		socialMediasRouter.POST("/", controllers.CreateSocialMedia)
		socialMediasRouter.GET("/", controllers.GetAllSocialMedia)
		socialMediasRouter.GET("/:socialMediaId", controllers.GetOneSocialMedia)
		socialMediasRouter.PUT("/:socialMediaId", middlewares.SocialMediaAuthorization, controllers.UpdateSocialMedia)
		socialMediasRouter.DELETE("/:socialMediaId", middlewares.SocialMediaAuthorization, controllers.DeleteSocialMedia)
	}

	return r
}
