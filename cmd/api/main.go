package main

import (
	"fmt"
	"github.com/ShaunBillows/learn-go-api-v2/internal/database"
	"log"
)

type Question struct {
	ID               uint   `gorm:"primaryKey" json:"id"`
	Topic            string `gorm:"not null" json:"topic"`
	Question         string `gorm:"not null;unique" json:"question"`
	Answer           string `gorm:"not null" json:"answer"`
	IncorrectAnswer1 string `gorm:"not null" json:"incorrect_answer_1"`
	IncorrectAnswer2 string `gorm:"not null" json:"incorrect_answer_2"`
	Difficulty       uint   `gorm:"not null" json:"difficulty"`
}

func main() {

	// Connect to MySQL db
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	// Migrate the schema
	db.AutoMigrate(&Question{}) // pass the tables you want to migrate as args

	// Create a pointer to a new Question object
	q := &Question{
		Topic:            "This is a test",
		Question:         "This is a test",
		IncorrectAnswer1: "This is a test",
		IncorrectAnswer2: "This is a test",
		Answer:           "This is a test",
		Difficulty:       3,
	}

	// Insert the new question into the database
	result := db.Table("questions").Create(q)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Printf("New question with ID %d created.\n", q.ID)

	// Retrieve a question by ID
	var questionByID Question
	result = db.Table("questions").First(&questionByID, q.ID)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Printf("Question with ID %d and Question: %s retrieved.\n", questionByID.ID, questionByID.Question)

	// Update the retrieved question
	questionByID.Topic = "Updated topic"
	result = db.Table("questions").Save(&questionByID)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Printf("Question with ID %d and Question: %s updated.\n", questionByID.ID, questionByID.Question)

	// Delete the retrieved question
	result = db.Table("questions").Delete(&questionByID)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Printf("Question with ID %d deleted.\n", questionByID.ID)

	// Retrieve all questions from the database
	var questions []Question
	result = db.Table("questions").Find(&questions)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	// Print all questions in the db
	fmt.Printf("Retrieved %d questions:\n", len(questions))
	fmt.Println(questions)
}
