// db/migration.go
package db

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Username string
	Email    string
	Password []byte // Change the type to []byte for storing hashed password
}

func AutoMigrateModels() {
	DB.AutoMigrate(&User{}) // Update with your model(s)
}
