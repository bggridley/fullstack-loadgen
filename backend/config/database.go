package config

import (
	"fmt"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"backend/helper"
	"strconv"
)

func DatabaseConnection() *gorm.DB {
	var host     = os.Getenv("DB_URL")
	var p     = os.Getenv("DB_PORT")
	var user     = os.Getenv("DB_USERNAME")
	var password = os.Getenv("DB_PASSWORD")
	var dbName   = "test"

	port, err := strconv.Atoi(p)
	helper.ErrorPanic(err)



	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	pConn := postgres.Open(sqlInfo)

	fmt.Println("got here", pConn)
	db, err := gorm.Open(pConn, &gorm.Config{})
	helper.ErrorPanic(err)

	fmt.Println("Did something")

	return db
}