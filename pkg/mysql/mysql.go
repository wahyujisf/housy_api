package mysql

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {

	env := godotenv.Load()
	if env != nil {
		log.Fatal("Error loading .env file")
	}

	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/housy?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to Database")
}
