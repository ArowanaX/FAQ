package bootstrap

import (
	"FAQ/internal/api/rest"
	"FAQ/internal/repository"
	"FAQ/internal/usecase"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// Keep all AppContainer of project
type AppContainer struct {
	ContactUseCase usecase.ContactUseCase
	FaqUseCase     usecase.FaqUseCase
}

// InitializeApp Initialize all sections
func InitializeApp(db *gorm.DB) *AppContainer {

	contactRepo := repository.NewContactRepository(db)
	faqRepo := repository.NewFaqRepository(db)

	contactUseCase := usecase.NewContactUseCase(contactRepo)
	faqUseCase := usecase.NewFaqUseCase(faqRepo)

	return &AppContainer{
		ContactUseCase: contactUseCase,
		FaqUseCase:     faqUseCase,
	}
}

// SetupRoutes Handles APIs paths.
func SetupRoutes(e *echo.Echo, container *AppContainer) {
	rest.SetupContactRoutes(e, container.ContactUseCase)
	rest.SetupFaqRoutes(e, container.FaqUseCase)
}
