package rest

import (
	"net/http"

	"FAQ/internal/entity"
	"FAQ/internal/usecase"
	"github.com/labstack/echo/v4"
)

type ContactHandler struct {
	contactUseCase usecase.ContactUseCase
}

func NewContactHandler(contactUseCase usecase.ContactUseCase) *ContactHandler {
	return &ContactHandler{contactUseCase: contactUseCase}
}

type ContactRequest struct {
	FirstName string `json:"first_name" validate:"required,min=2,max=50"`
	LastName  string `json:"last_name" validate:"required,min=2,max=50"`
	PhoneNum  string `json:"phone_number" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Message   string `json:"message" validate:"required"`
}

// هندلر برای ثبت کانتکت جدید
func (h *ContactHandler) CreateContact(c echo.Context) error {
	var req ContactRequest

	// پارس کردن ورودی
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	// اعتبارسنجی
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// تبدیل درخواست به entity
	contact := entity.ContactUs{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		PhoneNum:  req.PhoneNum,
		Email:     req.Email,
		Message:   req.Message,
	}

	// ذخیره در دیتابیس
	createdContact, err := h.contactUseCase.CreateContact(contact)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, createdContact)
}
