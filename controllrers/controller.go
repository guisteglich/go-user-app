package controllrers

import (
	"files/initializers"
	"files/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateBucketRequest struct {
	BucketName string `json:"bucketName"`
}

func UserHandler(c *gin.Context) {

	var newUser struct {
		Name string
	}

	c.Bind(&newUser)

	user := models.User{Name: newUser.Name}
	result := initializers.DB.Create(&user) // pass pointer of data to Create
	if result.Error != nil{
		c.Status(400)
		return
	}

	c.JSON(201, gin.H{
		"User": user,
	})
	}

func ListUsersHandler(c *gin.Context){
	var users []models.User
	// Get all records
	initializers.DB.Find(&users)
	c.JSON(200, gin.H{
		"Users": users,
	})

}

func ListUserHandler(c *gin.Context){
	name := c.Param("Name")

	var user []models.User
	// Get first matched record
	initializers.DB.First(&user, name)
	c.JSON(200, gin.H{
		"User": user,
	})

}

func ListBucketsHandler(c *gin.Context) {
	output, err := initializers.ListBuckets()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to list buckets",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"buckets": output,
	})
}


func CreateBucketHandler(c *gin.Context) {
	var req CreateBucketRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request payload",
		})
		return
	}

	err := initializers.CreateBucket(req.BucketName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create bucket",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Bucket created",
	})
}
