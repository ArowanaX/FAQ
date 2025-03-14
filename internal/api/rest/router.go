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

func SetupFaqRoutes(e *echo.Echo, faqUseCase usecase.FaqUseCase) {
	faqHandler := NewFaqHandler(faqUseCase)

	v2 := e.Group("/support/api/v2")
	v2.GET("/faq", faqHandler.GetAllFaq)
	v2.POST("/faq", faqHandler.CreateFaq)
	v2.GET("/faq/category/:title", faqHandler.GetCategorizedFaq)
	v2.POST("/faq/category", faqHandler.CreateCategory)
}
