package repository

import (
	"FAQ/internal/entity"
	"gorm.io/gorm"
)

type FaqRepository interface {
	GetAllFaq() ([]entity.Faq, error)
	GetCategorizeFaq(category entity.Category) ([]entity.Faq, error)
	CreateFaq(faq entity.Faq) error
	CreateCategory(cat entity.Category) error
}
type faqRepo struct {
	db *gorm.DB
}

func NewFaqUseCase(db *gorm.DB) FaqRepository {
	return &faqRepo{db: db}
}

func (fr *faqRepo) GetAllFaq() ([]entity.Faq, error) {
	var faqs []entity.Faq
	result := fr.db.Find(&faqs)
	if result.Error != nil {
		return nil, result.Error
	}
	return faqs, nil
}

func (fr *faqRepo) GetCategorizeFaq(category entity.Category) ([]entity.Faq, error) {
	var faqs []entity.Faq

	result := fr.db.Where("category_title = ?", category.Title).Find(&faqs)

	if result.Error != nil {
		return nil, result.Error
	}

	return faqs, nil
}
func (fr *faqRepo) CreateFaq(faq entity.Faq) error {
	result := fr.db.Create(&faq)

	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (fr *faqRepo) CreateCategory(cat entity.Category) error {
	result := fr.db.Create(&cat)

	if result.Error != nil {
		return result.Error
	}
	return nil
}
