package handlers

import (
	"github.com/ShaunBillows/learn-go-api-v2/internal/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"strconv"
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

func (h *QuestionHandler) CreateOne(c echo.Context) error {
	// Create empty struct from the questions model
	question := new(models.Question)
	// Bind the request to the struct
	if err := c.Bind(question); err != nil { // maps data to empty question struct
		return c.JSON(http.StatusBadRequest, responseMessage{Message: invalidJsonFormatMessage})
	}
	// Insert the new question into the database
	result := h.DB.Create(question)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: databaseErrorMessage})
	}
	return c.JSON(http.StatusOK, question)
}

func (h *QuestionHandler) ReadOne(c echo.Context) error {
	id := c.Param("id")
	question := new(models.Question)
	result := h.DB.First(&question, id)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: databaseErrorMessage})
	}
	return c.JSON(http.StatusOK, question)
}

func (h *QuestionHandler) DeleteOne(c echo.Context) error {
	id := c.Param("id")
	result := h.DB.Delete(&models.Question{}, id)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: databaseErrorMessage})
	}
	return c.JSON(http.StatusOK, responseMessage{Message: questionDeletedMessage})

}

func (h *QuestionHandler) UpdateOne(c echo.Context) error {
	question := new(models.Question)
	if err := c.Bind(question); err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: invalidJsonFormatMessage})
	}
	id, err := strconv.ParseUint(c.Param("id"), 10, 64) // can we convert directly to uint here?
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: invalidJsonFormatMessage})
	}
	question.ID = uint(id)
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
		return c.JSON(http.StatusBadRequest, responseMessage{Message: databaseErrorMessage})
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
		return c.JSON(http.StatusBadRequest, responseMessage{Message: invalidJsonFormatMessage})
	}
	result := h.DB.Create(&questions)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: databaseErrorMessage})
	}
	return c.JSON(http.StatusOK, questions)
}
