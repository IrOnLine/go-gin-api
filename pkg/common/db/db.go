package db

import (
	"log"

	"go-gin-api/pkg/common/models" // import models

	"gorm.io/driver/mysql" // MySQL driver
	"gorm.io/gorm"         // GORM library
)

func Init(url string) *gorm.DB { // function to initialize database
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{}) // open connection to MySQL
	if err != nil {                                       // check for errors
		log.Fatalln(err) // log fatal error if can't connect
	}

	db.AutoMigrate(&models.Book{}) // automatically migrate Book model

	return db // return db connection
}
