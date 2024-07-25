package main

import (
  "encoding/json"
  "io/ioutil"
  "fmt"
  "os"
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

// this is strictly because AWS did not want to create secret object in kubernetes...
func loadEnvHack() {
  secretsFilePath := "/mnt/secrets-store/realsecrets" // can use a github action secret for this bit later

	secretsData, err := ioutil.ReadFile(secretsFilePath)
	if err != nil {
		fmt.Println("Secrets file not found. Exiting.")
		return
	}
	var secretsMap map[string]string
	if err := json.Unmarshal(secretsData, &secretsMap); err != nil {
		fmt.Println("Error decoding JSON content. Exiting.")
		return
	}

	for key, value := range secretsMap {
		os.Setenv(key, value)
	}

	fmt.Println("Environment variables loaded from secrets successfully.")
}