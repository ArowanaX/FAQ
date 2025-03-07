package usecase

import (
	"FAQ/internal/entity"
)

type ContactUseCase interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	IsEmailUnique(email string) (bool, error)
	PostContactUs(us entity.ContactUs) (entity.ContactUs, error)
}
type contactUseCase struct {
	contactRepo ContactRepository
}

// constructors
func NewContactUseCase(Repo ContactRepository) ContactUsecase {
	return &contactUsecase{contactRepo: repo}
}

// interface method
func (u *contactUsecase) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	return u.contactRepo.IsPhoneNumberUnique(phoneNumber)
}

func (u *contactUsecase) IsEmailUnique(email string) (bool, error) {
	return u.contactRepo.IsEmailUnique(email)
}

func (u *contactUsecase) PostContactUs(cu entity.ContactUs) (entity.ContactUs, error) {
	return u.contactRepo.PostContactUs(cu)
}
