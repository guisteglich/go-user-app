package main

import (
	"files/controllrers"
	"files/initializers"

	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnvVariable()
	initializers.ConnToDB()
}

func main() {
	r := gin.Default()
	r.POST("/user", controllrers.UserHandler)
	r.GET("/users", controllrers.ListUsersHandler)
	r.GET("/user", controllrers.ListUserHandler)

	r.Run() // listen and serve on 0.0.0.0:8080
}