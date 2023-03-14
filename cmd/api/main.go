package main

import (
	"github.com/ShaunBillows/learn-go-api-v2/internal/database"
	"github.com/ShaunBillows/learn-go-api-v2/internal/handlers"
	"github.com/ShaunBillows/learn-go-api-v2/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.GET("/", handlers.SayWelcome)

	e.GET("/questions", questionHandler.ReadAll)
	e.GET("/questions/:id", questionHandler.ReadOne)
	e.PUT("/questions/:id", questionHandler.UpdateOne)
	e.DELETE("/questions/:id", questionHandler.DeleteOne)
	e.POST("/questions", questionHandler.CreateOne)
	e.GET("/questions/random", questionHandler.ReadRandom)
	e.POST("/questions/many", questionHandler.CreateMany)

	e.Logger.Fatal(e.Start(":1323"))
}
