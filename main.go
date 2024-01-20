package main

import (
	"github.com/ashiqYousuf/GoApisStructure/initializers"
	"github.com/ashiqYousuf/GoApisStructure/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	router := gin.Default()

	routes.PostRouter(router)

	routes.FourOFour(router)

	router.Run("localhost:8080")
}
