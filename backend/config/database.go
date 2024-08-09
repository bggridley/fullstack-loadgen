package config

import (
	"fmt"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"backend/helper"
	// "strconv"
)

func DatabaseConnection() *gorm.DB {
	var dbUsername     = os.Getenv("DB_USERNAME")
	var dbPassword     = os.Getenv("DB_PASSWORD")

	sqlInfo := fmt.Sprintf("host=fullstackloadgen.postgres.database.azure.com port=5432 user=%s password=%s dbname=postgres sslmode=require", dbUsername, dbPassword)
	pConn := postgres.Open(sqlInfo)

	fmt.Println("got here", pConn)
	db, err := gorm.Open(pConn, &gorm.Config{})
	helper.ErrorPanic(err)

	fmt.Println("Did something")

	return db
}