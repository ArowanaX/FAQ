package rest

import (
	"FAQ/internal/usecase"
	"github.com/labstack/echo/v4"
)

type ContactHandler struct {
	contactUC usecase.ContactUseCase
}

func NewContactHandler(contactUC usecase.ContactUseCase) *ContactHandler {
	return &ContactHandler{contactUC: contactUC}
}

func (h *ContactHandler) IsPhoneNumberUnique(c echo.Context) error {

}
