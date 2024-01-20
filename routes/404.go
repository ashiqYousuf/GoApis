package routes

import "github.com/gin-gonic/gin"

func FourOFour(router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		c.JSON(400, gin.H{
			"success": false,
			"message": "This is not a valid endpoint, please check our documentation",
		})
	})
}
