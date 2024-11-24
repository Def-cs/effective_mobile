package main

import (
	"effective_mobile/api"
	_ "effective_mobile/docs"
	"effective_mobile/internal/songs"
	logger "effective_mobile/logs"
	"effective_mobile/pkg/external"
	"effective_mobile/pkg/storage/db/postgres"
	"fmt"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"strconv"
)

// @title Songs API
// @version 1.0
// @description API для управления песнями.
// @host localhost:8080
// @BasePath /

// @contact.name Frolov Vladislav
// @contact.url https://hh.ru/resume/7b5e19efff0c43b3390039ed1f4e5a635a4558
// @contact.email vfrolov2004@gmail.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8000
// @BasePath /

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("No .env file found:" + err.Error())
	}
	loggers, err := logger.InitLoggers(os.Getenv("INFO_PATH"), os.Getenv("ERROR_PATH"), os.Getenv("FATAL_PATH"))
	if err != nil {
		panic(err.Error())
	}
	defer loggers.Close()

	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		panic("Cant transform str to int: " + err.Error())
	}
	postgres.InitConn(dbPort, os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), loggers)

	defer postgres.Connection.Close()

	apiClient := external.NewApiClient(os.Getenv("EXT_API_PATH"))
	songs.ApiClient = apiClient

	router := api.RegisterMux(loggers)
	fmt.Println("ЗАПУСКАЕМ СЕРВЕР")
	err = http.ListenAndServe(os.Getenv("SERVER_ADR"), router)
	if err != nil {
		loggers.Fatal(err.Error())
	}

}
