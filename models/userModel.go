package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
  gorm.Model
  Name        string    `json:"name"`
  Email       string    `json:"email"`
  Password    string    `json:"password"`
  DateOfBirth time.Time `json:"date_of_birth"`
  Phone       string    `json:"phone"`
}
