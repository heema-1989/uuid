package main

import (
	"gopkg.in/gomail.v2"
	"log"
)

func main() {
	mail := gomail.NewMessage()
	mail.SetHeader("From", "heema@simformsolutions.com")
	mail.SetHeader("To", "kishan.m@simformsolutions.com")
	mail.SetHeader("Subject", "Onboarding")
	mail.SetBody("text/plain", "Congratulations, Kishan Maheta! You are onboarded....")
	dialer := gomail.NewDialer("smtp.office365.com", 587, "heema@simformsolutions.com", "JL*$_un4wVNC4hMY")
	if err := dialer.DialAndSend(mail); err != nil {
		log.Fatal("Error sending email. ", err)
	}
}
