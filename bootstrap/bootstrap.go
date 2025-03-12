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
	//FAQUseCase     usecase.FAQUseCase
}

// InitializeApp Initialize all sections
func InitializeApp(db *gorm.DB) *AppContainer {

	contactRepo := repository.NewContactRepository(db)
	//faqRepo := repository.NewFAQRepository(db)

	contactUseCase := usecase.NewContactUseCase(contactRepo)
	//faqUseCase := usecase.NewFAQUseCase(faqRepo)

	return &AppContainer{
		ContactUseCase: contactUseCase,
		//FAQUseCase:     faqUseCase,
	}
}

// SetupRoutes Handles APIs paths.
func SetupRoutes(e *echo.Echo, container *AppContainer) {
	rest.SetupContactRoutes(e, container.ContactUseCase)
	//rest.SetupFAQRoutes(e, container.FAQUseCase)
}
