package migration

import (
	"fmt"

	"blog.com/config"
	"blog.com/models"
)

func Migrate() {
	db := config.GoConnect()
	db.AutoMigrate(&models.Post{})
	fmt.Println("Migration successfull")
}
