package usecase

import (
	"FAQ/internal/entity"
	"FAQ/internal/repository"
	"FAQ/internal/validator"
	"errors"
)

type ContactUseCase interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	IsEmailUnique(email string) (bool, error)
	CreateContact(us entity.ContactUs) (entity.ContactUs, error)
}

type contactUseCase struct {
	repo repository.ContactRepository
}

func NewContactUseCase(repo repository.ContactRepository) ContactUseCase {
	return &contactUseCase{repo: repo}
}

// interface method
func (u *contactUseCase) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	if !validator.IsValidPhone(phoneNumber) {
		return false, errors.New("invalid phone number format")
	}
	return u.repo.IsPhoneNumberUnique(phoneNumber)
}

func (u *contactUseCase) IsEmailUnique(email string) (bool, error) {
	if !validator.IsValidEmail(email) {
		return false, errors.New("invalid email format")
	}
	return u.repo.IsEmailUnique(email)
}

func (u *contactUseCase) CreateContact(us entity.ContactUs) (entity.ContactUs, error) {

	if !validator.IsValidName(us.FirstName) || !validator.IsValidName(us.LastName) {
		return entity.ContactUs{}, errors.New("invalid name format")
	}
	if !validator.IsValidPhone(us.PhoneNum) {
		return entity.ContactUs{}, errors.New("invalid phone number format")
	}
	if !validator.IsValidEmail(us.Email) {
		return entity.ContactUs{}, errors.New("invalid email format")
	}

	phoneUnique, _ := u.repo.IsPhoneNumberUnique(us.PhoneNum)
	if !phoneUnique {
		return entity.ContactUs{}, errors.New("phone number already exists")
	}

	emailUnique, _ := u.repo.IsEmailUnique(us.Email)
	if !emailUnique {
		return entity.ContactUs{}, errors.New("email already exists")
	}

	err := u.repo.Create(us)
	if err != nil {
		return entity.ContactUs{}, err
	}

	return us, nil
}
