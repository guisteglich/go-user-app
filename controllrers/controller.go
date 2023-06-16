package controllrers

import (
	"files/initializers"
	"files/models"

	"github.com/gin-gonic/gin"
)

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
	// fmt.Println(result.RowsAffected)
	// if result.Error != nil{
	// 	c.Status(404)
	// 	return
	// }
	c.JSON(200, gin.H{
		"Users": users,
	})

}

func ListUserHandler(c *gin.Context){
	var user struct {
		Name string
	}
	c.Bind(&user)

	var users []models.User
	// Get first matched record
	initializers.DB.Where("name = ?", user.Name).First(&users)
	c.JSON(200, gin.H{
		"User": users,
	})

}