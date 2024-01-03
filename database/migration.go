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

type File struct {
	gorm.Model
	FileName  string
	Path      string
	Extension string
}

func AutoMigrateModels() {
	DB.AutoMigrate(&User{}, &File{}) // Update with your model(s)
}
