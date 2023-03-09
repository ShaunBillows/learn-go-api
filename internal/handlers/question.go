package handlers

import (
	"github.com/ShaunBillows/learn-go-api-v2/internal/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

type Question struct {
	DB *gorm.DB
}

type ID struct {
	ID uint `json:"id"`
}

func NewQuestion(db *gorm.DB) *Question {
	return &Question{
		DB: db,
	}
}

func (h *Question) Create(c echo.Context) error {
	// Create empty struct from the questions model
	question := new(models.Question)
	// Bind the request to the struct
	if err := c.Bind(question); err != nil { // maps data to empty question struct
		return err
	}
	// Insert the new question into the database
	result := h.DB.Create(question)
	if result.Error != nil {
		return result.Error
	}
	return c.JSON(http.StatusOK, question)
}

func (h *Question) Delete(c echo.Context) error {

	id := new(ID)
	if err := c.Bind(id); err != nil {
		return err
	}

	// Delete the question with the given ID from the database
	result := h.DB.Delete(&models.Question{}, id)
	if result.Error != nil {
		return result.Error
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Question deleted successfully"})
}

func (h *Question) Update(c echo.Context) error {
	question := new(models.Question)
	if err := c.Bind(question); err != nil {
		return err
	}
	result := h.DB.Save(question)
	if result.Error != nil {
		return result.Error
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Question updated successfully"})
}

func (h *Question) Read(c echo.Context) error {
	question := new(models.Question)
	result := h.DB.Order("RAND()").First(question)
	if result.Error != nil {
		return result.Error
	}
	return c.JSON(http.StatusOK, question)
}

func (h *Question) ReadAll(c echo.Context) error {
	questions := new([]models.Question)
	result := h.DB.Find(&questions)
	if result.Error != nil {
		return result.Error
	}
	return c.JSON(http.StatusOK, questions)
}

func (h *Question) CreateMany(c echo.Context) error {
	questions := new([]models.Question)
	if err := c.Bind(questions); err != nil {
		return err
	}
	result := h.DB.Create(&questions)
	if result.Error != nil {
		return result.Error
	}
	return c.JSON(http.StatusOK, questions)
}
