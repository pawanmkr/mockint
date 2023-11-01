package services

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
)

func SendMail(receiver string, msg string) {
	if msg == "" {
		log.Fatal("join URL is empty")
	} else {
		from := os.Getenv("FROM")
		password := os.Getenv("PASSWORD")
		to := []string{
			receiver,
		}
		body := fmt.Sprintf("Subject: Interview Meeting Join URL\n\n%s", msg)

		smtpHost := "smtp.gmail.com"
		smtpPort := "587"

		auth := smtp.PlainAuth("", from, password, smtpHost)

		err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, []byte(body))
		if err != nil {
			log.Fatalf("Error while sending email: %v", err)
			return
		}
		fmt.Println("Email Sent Successfully!")
	}
}
