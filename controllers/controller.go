package controllers

import (
	"files/initializers"
	"files/models"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// ...


func CreateUserHandler(c *gin.Context) {
	var newUser struct {
		Name        string    `json:"name"`
		Email       string    `json:"email"`
		Password    string    `json:"password"`
		DateOfBirth string    `json:"date_of_birth"`
		Phone       string    `json:"phone"`
	}

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Criptografar a senha usando bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criptografar a senha"})
		return
	}

	// Converter a string da data de nascimento para o tipo time.Time
	dateOfBirth, err := time.Parse("02/01/2006", newUser.DateOfBirth)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de data de nascimento inválido"})
		return
	}

	user := models.User{
		Name:        newUser.Name,
		Email:       newUser.Email,
		Password:    string(hashedPassword),
		DateOfBirth: dateOfBirth,
		Phone:       newUser.Phone,
	}

	initializers.ConnToDB()

	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user": user,
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

func DeleteUserHandler(c *gin.Context) {
	userID := c.Param("id")

	result := initializers.DB.Delete(&models.User{}, userID)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao excluir usuário"})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuário excluído com sucesso"})
}



func ListUserHandler(c *gin.Context) {
	userID := c.Param("id")
	name := c.Param("Name")

	var user models.User

	if name != "" {
		result := initializers.DB.First(&user, name)
		if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}
	}
	result := initializers.DB.First(&user, userID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func UploadUserFile(c *gin.Context) {
	// recebe o id do usuário
	userID := c.Param("id")
	// recebe o arquivo para upload
	fileForm, err := c.FormFile("file")
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	openedFile, err := fileForm.Open()
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	defer openedFile.Close()

	fileBytes, err := ioutil.ReadAll(openedFile)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to read file",
		})
		return
	}

	err = initializers.UploadFile("img-"+userID, fileBytes)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to upload file to bucket",
		})
		return
	}

	// Se o upload foi bem-sucedido
	c.JSON(http.StatusOK, gin.H{
		"message": "File uploaded successfully",
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
	type CreateBucketRequest struct {
	BucketName string `json:"bucketName"`
	}

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
