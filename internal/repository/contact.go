package repository

import (
	"FAQ/internal/entity"
	"FAQ/internal/models"
	"gorm.io/gorm"
)

type ContactRepository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	IsEmailUnique(email string) (bool, error)
	Create(contact entity.ContactUs) error
}
type contactRepository struct {
	db *gorm.DB
}

func NewContactRepository(db *gorm.DB) ContactRepository {
	return &contactRepository{db: db}
}

func (r *contactRepository) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	var count int64
	err := r.db.Model(&models.ContactUsModel{}).Where("phone_num = ?", phoneNumber).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

func (r *contactRepository) IsEmailUnique(email string) (bool, error) {
	var count int64
	err := r.db.Model(&models.ContactUsModel{}).Where("email = ?", email).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

func (r *contactRepository) Create(contact entity.ContactUs) error {
	contactModel := models.ContactUsModel{
		FirstName: contact.FirstName,
		LastName:  contact.LastName,
		PhoneNum:  contact.PhoneNum,
		Email:     contact.Email,
		Message:   contact.Message,
		SentMail:  contact.SentMail,
		IpAddress: contact.IpAddress,
	}

	return r.db.Create(&contactModel).Error
}
