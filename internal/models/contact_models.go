package models

import "gorm.io/gorm"

type ContactUsModel struct {
	gorm.Model
	FirstName string `gorm:"size:50;not null"`
	LastName  string `gorm:"size:50;not null"`
	PhoneNum  string `gorm:"size:20;not null"`
	Email     string `gorm:"size:100;not null"`
	Message   string `gorm:"type:text;not null"`
	SentMail  bool
	IpAddress string `gorm:"size:45"`
}
