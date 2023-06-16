package initializers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var DB *gorm.DB

func ConnToDB(){
	var err error
	dsn := os.Getenv("DB_CURL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		log.Fatal("Error db conn")
	}
	fmt.Println("DB connected")
}