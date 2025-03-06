package entity

import (
	"regexp"
)

type ContactUs struct {
	FirstName ValidNameField
	LastName  ValidNameField
	PhoneNum  string
	Email     ValidEmailField
	Message   string
	SentMail  bool
	IpAddress string
}

type ValidNameField string

func (v ValidNameField) IsValid() bool {
	if len(v) <= 1 && len(v) >= 50 {
		return true
	}
	return false
}

type ValidEmailField string

func (v ValidEmailField) IsValid() bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(string(v))
}

type ValidPhoneNumField string

func (v ValidPhoneNumField) IsValid() bool {
	re := regexp.MustCompile(`^(?:\+98\d{10}|0\d{10})$`)
	return re.MatchString(string(v))
}
