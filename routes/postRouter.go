package routes

import (
	"github.com/ashiqYousuf/GoApisStructure/controllers"
	"github.com/gin-gonic/gin"
)

func PostRouter(router *gin.Engine) {
	router.POST("/posts", controllers.CreatePost)
	router.GET("/posts", controllers.GetPosts)
	router.GET("/posts/:id", controllers.GetPost)
	router.PATCH("/posts/:id", controllers.UpdatePost)
	router.DELETE("/posts/:id", controllers.DeletePost)
}
