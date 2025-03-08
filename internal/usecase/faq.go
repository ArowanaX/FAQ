package usecase

import (
	"FAQ/internal/entity"
	"FAQ/internal/repository"
)

type FaqUseCase interface {
	GetAllFaq() ([]entity.Faq, error)
	GetCategorizeFaq(category entity.Category) ([]entity.Faq, error)
	CreateFaq(faq entity.Faq) error
	CreateCategory(cat entity.Category) error
}
type faqUseCase struct {
	faqRepo repository.FaqRepository
}

// constructors
func NewFaqUseCase(Repo repository.FaqRepository) FaqUseCase {
	return &faqUseCase{faqRepo: Repo}
}

// interface method
func (fu *faqUseCase) GetAllFaq() ([]entity.Faq, error) {
	return fu.faqRepo.GetAllFaq()
}

func (fu *faqUseCase) GetCategorizeFaq(category entity.Category) ([]entity.Faq, error) {
	return fu.faqRepo.GetCategorizeFaq(category)
}

func (fu *faqUseCase) CreateFaq(faq entity.Faq) error {
	return fu.faqRepo.CreateFaq(faq)
}

func (fu *faqUseCase) CreateCategory(cat entity.Category) error {
	return fu.faqRepo.CreateCategory(cat)
}
