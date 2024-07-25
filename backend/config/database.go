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
	var host     = os.Getenv("host")
	var p     = os.Getenv("port")
	var user     = os.Getenv("username")
	var password = os.Getenv("password")
	var dbName   = os.Getenv("dbname")

	port, err := strconv.Atoi(p)
	helper.ErrorPanic(err)



	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require", host, port, user, password, dbName)
	pConn := postgres.Open(sqlInfo)

	fmt.Println("got here", pConn)
	db, err := gorm.Open(pConn, &gorm.Config{})
	helper.ErrorPanic(err)

	fmt.Println("Did something")

	return db
}