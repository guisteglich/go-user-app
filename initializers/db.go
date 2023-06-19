package initializers

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var DB *gorm.DB

func ConnToDB(){
	var err error
	// dsn := os.Getenv("DB_URL")
	dsn := "postgres://qptncvep:IRKc7pFuzOzdd75sPxfWVyuOlVckbbP5@silly.db.elephantsql.com:5432/qptncvep"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		log.Fatal("Error db conn:", err)
	}
	fmt.Println("DB connected")
}