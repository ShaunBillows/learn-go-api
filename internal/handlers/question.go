package handlers

import (
	"github.com/ShaunBillows/learn-go-api-v2/internal/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

const (
	questionDeletedMessage = "Question deleted successfully"
	questionUpdatedMessage = "Question updated successfully"

	invalidJsonFormatMessage = "Invalid JSON format"
	databaseErrorMessage     = "Error occurred while accessing database"
)

type responseMessage struct {
	Message string `json:"message"`
}

type QuestionHandler struct {
	DB *gorm.DB
}

func NewQuestionHandler(db *gorm.DB) *QuestionHandler {
	return &QuestionHandler{
		DB: db,
	}
}

func (h *QuestionHandler) Create(c echo.Context) error {
	// Create empty struct from the questions model
	question := new(models.Question)
	// Bind the request to the struct
	if err := c.Bind(question); err != nil { // maps data to empty question struct
		return c.JSON(http.StatusBadRequest, responseMessage{Message: databaseErrorMessage})
	}
	// Insert the new question into the database
	result := h.DB.Create(question)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: databaseErrorMessage})
	}
	return c.JSON(http.StatusOK, question)
}

func (h *QuestionHandler) Read(c echo.Context) error {
	id := c.Param("id")
	question := new(models.Question)
	result := h.DB.First(&question, id)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: databaseErrorMessage})
	}
	return c.JSON(http.StatusOK, question)
}

func (h *QuestionHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	// Delete the question with the given ID from the database
	result := h.DB.Delete(&models.Question{}, id)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: databaseErrorMessage})
	}
	return c.JSON(http.StatusOK, responseMessage{Message: questionDeletedMessage})

}

func (h *QuestionHandler) Update(c echo.Context) error {
	question := new(models.Question)
	if err := c.Bind(question); err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: databaseErrorMessage})
	}
	result := h.DB.Save(question)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: databaseErrorMessage})
	}
	return c.JSON(http.StatusOK, responseMessage{Message: questionUpdatedMessage})
}

func (h *QuestionHandler) ReadRandom(c echo.Context) error {
	question := new(models.Question)
	result := h.DB.Order("RAND()").First(question)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: invalidJsonFormatMessage})
	}
	return c.JSON(http.StatusOK, question)
}

func (h *QuestionHandler) ReadAll(c echo.Context) error {
	questions := new([]models.Question)
	result := h.DB.Find(&questions)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: databaseErrorMessage})
	}
	return c.JSON(http.StatusOK, questions)
}

func (h *QuestionHandler) CreateMany(c echo.Context) error {
	questions := new([]models.Question)
	if err := c.Bind(questions); err != nil {
		return err
	}
	result := h.DB.Create(&questions)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: databaseErrorMessage})
	}
	return c.JSON(http.StatusOK, questions)
}
