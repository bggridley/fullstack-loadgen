package config

import (
	"fmt"
	// "os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"backend/helper"
	// "strconv"
)

func DatabaseConnection() *gorm.DB {
	// var connString     = os.Getenv("CONN_STRING")

	pConn := postgres.Open("host=fullstackloadgen.postgres.database.azure.com port=5432 user=test123 password=test123 dbname=postgres sslmode=require")

	fmt.Println("got here", pConn)
	db, err := gorm.Open(pConn, &gorm.Config{})
	helper.ErrorPanic(err)

	fmt.Println("Did something")

	return db
}