package repository

import (
	"FAQ/internal/entity"
	"gorm.io/gorm"
)

type FaqRepository interface {
	GetAllFaq() ([]entity.Faq, error)
	GetCategorizeFaq(categoryTitle string) ([]entity.Faq, error)
	CreateFaq(faq entity.Faq) error
	CreateCategory(cat entity.Category) error
}
type faqRepository struct {
	db *gorm.DB
}

func NewFaqRepository(db *gorm.DB) FaqRepository {
	return &faqRepository{db: db}
}

func (fr *faqRepository) GetAllFaq() ([]entity.Faq, error) {
	var faqs []entity.Faq
	result := fr.db.Find(&faqs)
	if result.Error != nil {
		return nil, result.Error
	}
	return faqs, nil
}

func (fr *faqRepository) GetCategorizeFaq(categoryTitle string) ([]entity.Faq, error) {
	var faqs []entity.Faq

	result := fr.db.Where("category_title = ?", categoryTitle).Find(&faqs)

	if result.Error != nil {
		return nil, result.Error
	}

	return faqs, nil
}
func (fr *faqRepository) CreateFaq(faq entity.Faq) error {
	result := fr.db.Create(&faq)

	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (fr *faqRepository) CreateCategory(cat entity.Category) error {
	result := fr.db.Create(&cat)

	if result.Error != nil {
		return result.Error
	}
	return nil
}
