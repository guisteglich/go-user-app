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
	r.POST("/newUser", controllrers.CreateUserHandler) //Ok
	r.GET("/users", controllrers.ListUsersHandler) // Ok
	r.GET("/user", controllrers.ListUserHandler) // Ok
	r.DELETE("/deleteUser/:id", controllrers.DeleteUserHandler) // Ok
	
	r.GET("/buckets", controllrers.ListBucketsHandler)
	r.POST("/createBucket", controllrers.CreateBucketHandler)

	r.Run() // listen and serve on 0.0.0.0:8080
}