package utils

import (
	"bytes"
	"email-verification/models"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
	"html/template"
	"log"
	"os"
)

var Msg string

func CheckError(err error, msg string) {
	if err != nil {
		log.Fatal(msg, " Reason : ", err)
	}
}
func HashPassword(password string) string {
	var passwordBytes = []byte(password)
	hashedPasswordBytes, hashErr := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.MinCost)
	CheckError(hashErr, "Error generating hash from password")
	return string(hashedPasswordBytes)
}
func MatchPassword(hashPassword, currPassword string) bool {
	matchError := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(currPassword))
	return matchError == nil
}
func SendMail(to string, u *models.User) {
	var (
		sendMail models.SendMail
		t        *template.Template
		err      error
	)
	//t = template.New("html-tmpl")
	t, err = t.ParseFiles("/home/heema/GolandProjects/uuid/email-verification/views/default/email-verification.html")
	CheckError(err, "Error parsing html file")
	buff := new(bytes.Buffer)
	err = t.Execute(buff, &u)
	if err != nil {
		fmt.Println(err)
		return
	}
	Msg = buff.String()
	fmt.Println("Getting credentials")
	sendMail.From, sendMail.Username, sendMail.Password, sendMail.Port = GetSmtpMailCredentials()
	sendMail.To = to
	mail := gomail.NewMessage()
	mail.SetHeader("From", sendMail.From)
	mail.SetHeader("To", sendMail.To)
	mail.SetHeader("Subject", "Please verify your email")
	mail.SetBody("text/html", Msg)
	dialer := gomail.NewDialer("smtp.gmail.com", sendMail.Port, sendMail.Username, sendMail.Password)
	if err := dialer.DialAndSend(mail); err != nil {
		log.Fatal("Error sending mail", err)
	}
}
func GetSmtpMailCredentials() (string, string, string, int) {
	from := os.Getenv("EMAIL_FROM")
	username := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")
	port := 587
	return from, username, password, port
}
