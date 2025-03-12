package rest

import (
	"FAQ/internal/usecase"
	"github.com/labstack/echo/v4"
)

func SetupContactRoutes(e *echo.Echo, contactUseCase usecase.ContactUseCase) {
	contactHandler := NewContactHandler(contactUseCase)

	v2 := e.Group("/support/api/v2")
	v2.POST("/contact-us", contactHandler.CreateContact)
}
