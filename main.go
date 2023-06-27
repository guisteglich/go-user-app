package main

import (
	"files/controllers"
	"files/initializers"

	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnvVariable()
	initializers.ConnToDB()
}

func main() {
	r := gin.Default()
	r.POST("/newUser", controllers.CreateUserHandler) //Ok
	r.GET("/users", controllers.ListUsersHandler) // Ok
	r.GET("/user", controllers.ListUserHandler) // Ok
	r.DELETE("/deleteUser/:id", controllers.DeleteUserHandler) // Ok
	r.POST("/authentication", controllers.Auth)

	r.POST("/upload/:id", controllers.UploadUserFile)
	r.GET("/buckets", controllers.ListBucketsHandler)
	r.POST("/createBucket", controllers.CreateBucketHandler)

	r.Run() // listen and serve on 0.0.0.0:8080
}