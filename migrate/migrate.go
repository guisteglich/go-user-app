package main

import (
	"files/initializers"
	"files/models"
)


func init(){
	initializers.LoadEnvVariable()
	initializers.ConnToDB()
}
func main() {
	initializers.DB.AutoMigrate(&models.User{})
}
// Migrate the schema
  