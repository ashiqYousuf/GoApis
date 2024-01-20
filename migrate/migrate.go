package main

import (
	"github.com/ashiqYousuf/GoApisStructure/initializers"
	"github.com/ashiqYousuf/GoApisStructure/models"
)

// ? Automatically migrate your schema, to keep your schema up to date
// ? db.AutoMigrate(&User{}, &Product{}, &Order{})

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
