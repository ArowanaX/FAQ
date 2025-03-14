package rest

import (
	"FAQ/internal/entity"
	"FAQ/internal/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type FaqHandler struct {
	faqUseCase usecase.FaqUseCase
}

func NewFaqHandler(faqUseCase usecase.FaqUseCase) *FaqHandler {
	return &FaqHandler{faqUseCase: faqUseCase}
}

type FaqRequest struct {
	CategoryTitle string `json:"category_title" validate:"required"`
	Question      string `json:"question" validate:"required"`
	Answer        string `json:"answer" validate:"required"`
}

// Create new FAQ
func (h *FaqHandler) CreateFaq(c echo.Context) error {
	var req FaqRequest

	// Parse input
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Convert request to entity
	faq := entity.Faq{
		Category: entity.Category{Title: req.CategoryTitle},
		Question: req.Question,
		Answer:   req.Answer,
	}

	// Commit on DB
	if err := h.faqUseCase.CreateFaq(faq); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "FAQ created successfully"})
}

func (h *FaqHandler) GetAllFaq(c echo.Context) error {
	faqs, err := h.faqUseCase.GetAllFaq()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, faqs)
}

// Get FAQs with a category
func (h *FaqHandler) GetCategorizedFaq(c echo.Context) error {
	categoryTitle := c.Param("category")
	if categoryTitle == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "category title is required"})
	}

	faqs, err := h.faqUseCase.GetCategorizeFaq(categoryTitle)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, faqs)
}

// Category handler
type CategoryRequest struct {
	Title string `json:"title" validate:"required,min=2,max=50"`
}

func (h *FaqHandler) CreateCategory(c echo.Context) error {
	var req CategoryRequest

	// پارس کردن ورودی
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// تبدیل درخواست به انتیتی
	category := entity.Category{
		Title: req.Title,
	}

	// ایجاد در دیتابیس
	err := h.faqUseCase.CreateCategory(category)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Category created successfully"})
}
