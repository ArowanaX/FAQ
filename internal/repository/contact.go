package repository

import (
	"FAQ/internal/entity"
	"gorm.io/gorm"
)

type ContactRepository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	IsEmailUnique(email string) (bool, error)
	PostContactUs(us entity.ContactUs) (entity.ContactUs, error)
}
type contactRepo struct {
	db *gorm.DB
}

func NewContactUseCase(db *gorm.DB) ContactRepository {
	return &contactRepo{db: db}
}

func (cr *contactRepo) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	var contact entity.ContactUs
	result := cr.db.First(&contact, phoneNumber)

	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (cr *contactRepo) IsEmailUnique(email string) (bool, error) {
	var contact entity.ContactUs
	result := cr.db.First(&contact, email)

	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (cr *contactRepo) PostContactUs(us entity.ContactUs) (entity.ContactUs, error) {
	result := cr.db.Create(&us)

	if result.Error != nil {
		return entity.ContactUs{}, result.Error
	}

	return us, nil
}
