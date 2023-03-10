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

	questionHandler := handlers.NewQuestionHandler(db)

	e := echo.New()

	e.GET("/", handlers.SayWelcome)

	e.POST("/questions", questionHandler.Create)
	e.GET("/questions/:id", questionHandler.Read)
	e.PATCH("/questions", questionHandler.Update)
	e.DELETE("/questions/:id", questionHandler.Delete)
	e.GET("/questions/random", questionHandler.ReadRandom)
	e.GET("/questions/all", questionHandler.ReadAll)
	e.POST("/questions/many", questionHandler.CreateMany)

	e.Logger.Fatal(e.Start(":1323"))
}
