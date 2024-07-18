package main

import (
	"github.com/go-playground/validator/v10"
	"net/http"
	"backend/config"
	"backend/controller"
	"backend/helper"
	"backend/model"
	"backend/repository"
	"backend/router"
	"backend/service"
	"time"
	"github.com/joho/godotenv"
)

func main() {
  // Load env vars
  godotenv.Load() 

	//Database
	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("test").AutoMigrate(&model.Test{})

	//Init Repository
	testRepository := repository.NewTestRepositoryImpl(db)

	//Init Service
	testService := service.NewTestServiceImpl(testRepository, validate)

	//Init controller
	testController := controller.NewTestController(testService)

	//Router
	routes := router.NewRouter(testController)

	server := &http.Server{
		Addr:           ":8888",
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)

}