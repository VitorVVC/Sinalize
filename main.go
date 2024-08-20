package main

import (
	"SinalizeB/configs"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error to load .env: %v", err)
	}

	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("error to load configs: %v", err)
	}
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	port := cfg.API.Port
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))

}
