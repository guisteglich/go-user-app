package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"files/controllers"
	"files/initializers"
	"files/models"
)

type createUserResponse struct {
	User struct {
		ID           int    `json:"ID"`
		DateOfBirth  string `json:"date_of_birth"`
		Email        string `json:"email"`
		Name         string `json:"name"`
		Password     string `json:"password"`
		Phone        string `json:"phone"`
	} `json:"user"`
}

func compareResponses(expected createUserResponse, actual createUserResponse) bool {
	return expected.User.ID == actual.User.ID &&
		expected.User.DateOfBirth == actual.User.DateOfBirth &&
		expected.User.Email == actual.User.Email &&
		expected.User.Name == actual.User.Name &&
		expected.User.Phone == actual.User.Phone
}

func TestCreateUserHandler(t *testing.T) {
	router := gin.Default()
	router.POST("/createUser", controllers.CreateUserHandler)

	user := struct {
		Name        string `json:"name"`
		Email       string `json:"email"`
		Password    string `json:"password"`
		DateOfBirth string `json:"date_of_birth"`
		Phone       string `json:"phone"`
	}{
		Name:        "John Doe",
		Email:       "johndoe@example.com",
		Password:    "secretpassword",
		DateOfBirth: "01/01/1990",
		Phone:       "1234567890",
	}

	jsonUser, _ := json.Marshal(user)

	request, _ := http.NewRequest("POST", "/createUser", bytes.NewBuffer(jsonUser))
	request.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusCreated, recorder.Result().StatusCode)

	responseBody := strings.TrimSpace(recorder.Body.String())

	var actualResponse createUserResponse
	err := json.Unmarshal([]byte(responseBody), &actualResponse)
	if err != nil {
		t.Fatal(err)
	}

	expectedResponse := createUserResponse{
		User: struct {
			ID           int    `json:"ID"`
			DateOfBirth  string `json:"date_of_birth"`
			Email        string `json:"email"`
			Name         string `json:"name"`
			Password     string `json:"password"`
			Phone        string `json:"phone"`
		}{
			ID:           actualResponse.User.ID,
			DateOfBirth:  "1990-01-01T00:00:00Z",
			Email:        "johndoe@example.com",
			Name:         "John Doe",
			Password:     "", // Aceita senha vazia na resposta esperada
			Phone:        "1234567890",
		},
	}

	// Comparar as respostas, ignorando o campo de senha
	if !compareResponses(expectedResponse, actualResponse) {
		t.Errorf("Resposta esperada e resposta atual não são iguais")
	}
}

func TestMain(m *testing.M) {
	initializers.ConnToDB()
	initializers.DB.AutoMigrate(&models.User{})
	defer initializers.DB.Migrator().DropTable(&models.User{})
	m.Run()
}
