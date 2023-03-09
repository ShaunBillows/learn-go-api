package main

import (
	"github.com/ShaunBillows/learn-go-api-v2/internal/database"
	"github.com/ShaunBillows/learn-go-api-v2/internal/handlers"
	"github.com/ShaunBillows/learn-go-api-v2/internal/models"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {

	// Connect to MySQL db
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	// Migrate the schema
	// Makes sure db schema and models are in sync
	// Accepts a list of models
	if err := db.AutoMigrate(&models.Question{}); err != nil {
		log.Fatal(err)
	}

	questionHandler := handlers.NewQuestion(db)

	e := echo.New()

	e.GET("/", handlers.SayWelcome)

	// Single record endpoints
	e.POST("/question/create", questionHandler.Create)
	e.DELETE("/question/delete", questionHandler.Delete)
	e.PATCH("/question/update", questionHandler.Update)
	e.GET("/question/read", questionHandler.Read)

	// Many records endpoints
	e.GET("/question/read/all", questionHandler.ReadAll)
	e.POST("/question/create/many", questionHandler.CreateMany)

	e.Logger.Fatal(e.Start(":1323"))
}
