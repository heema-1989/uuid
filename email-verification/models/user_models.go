package models

import (
	uuid "github.com/satori/go.uuid"
)

type User struct {
	UserId   int       `gorm:"<-:false;type:integer primary key GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 );"`
	VerifyId uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name     string    `gorm:"type:varchar;not null"`
	UserName string    `gorm:"type:varchar(100); not null"`
	Email    string    `gorm:"type:varchar(100);not null;unique"`
	Password string    `gorm:"type:varchar(50);not null"`
}
type SendMail struct {
	From     string
	To       string
	Username string
	Password string
	Port     int
}
