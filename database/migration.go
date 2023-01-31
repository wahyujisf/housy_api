package database

import (
	"fmt"
	"housy/models"
	"housy/pkg/mysql"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.User{},
		&models.Property{},
		// &models.Product{},
		// &models.Category{},
		// &models.Transaction{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
