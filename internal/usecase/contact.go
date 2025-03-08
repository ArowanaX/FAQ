package usecase

import (
	"FAQ/internal/entity"
	"FAQ/internal/repository"
)

type ContactUseCase interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	IsEmailUnique(email string) (bool, error)
	PostContactUs(us entity.ContactUs) (entity.ContactUs, error)
}
type contactUseCase struct {
	contactRepo repository.ContactRepository
}

// constructors
func NewContactUseCase(Repo repository.ContactRepository) ContactUseCase {
	return &contactUseCase{contactRepo: Repo}
}

// interface method
func (u *contactUseCase) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	return u.contactRepo.IsPhoneNumberUnique(phoneNumber)
}

func (u *contactUseCase) IsEmailUnique(email string) (bool, error) {
	return u.contactRepo.IsEmailUnique(email)
}

func (u *contactUseCase) PostContactUs(cu entity.ContactUs) (entity.ContactUs, error) {
	return u.contactRepo.PostContactUs(cu)
}
