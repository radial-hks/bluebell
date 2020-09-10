package routes

import (
	"bluebell/controllers"
	"bluebell/logger"
	"bluebell/middlewares"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	//
	v1 := r.Group("/api/v1")
	v1.POST("/signup", controllers.SignUpHandler)
	v1.POST("/login", controllers.LoginHandler)
	v1.Use(middlewares.JWTAuthMiddleware())
	//r.GET("/main", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
	//	//c.Request.Header.Get("Authorizatuion")
	//	c.JSON(http.StatusOK, gin.H{
	//		"version": "Success At Here ok ",
	//	})
	//})
	{
		v1.GET("/community", controllers.CommunityHandler)
		v1.GET("/community/:id", controllers.CommunityDetailHandler)

		v1.POST("/post", controllers.CreatePostHandler)
		v1.GET("/post/:id", controllers.GetPostHandler)
		v1.GET("/post_", controllers.GetPostListHandler)
	}

	return r
}
