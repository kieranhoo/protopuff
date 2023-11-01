package mailers

import (
	"bytes"
	"log"
	"text/template"

	"gopkg.in/gomail.v2"
)

var (
	_email       string
	_appPassword string
)

func Config(email, appPassword string) {
	_email = email
	_appPassword = appPassword
}

func SendHTML(tempPath, to string) error {
	t, err := template.ParseFiles(tempPath)
	if err != nil {
		log.Fatal(err)
	}
	var body bytes.Buffer
	err = t.Execute(&body, struct{ Name string }{Name: "Kawasaki"})
	if err != nil {
		return err
	}

	mail := gomail.NewMessage()
	mail.SetHeader("From", _email)
	mail.SetHeader("To", to)
	// mail.SetAddressHeader("Cc", "dan@example.com", "Dan")
	mail.SetHeader("Subject", "Thư xác nhận sử dụng phòng Lab")
	mail.SetBody("text/html", body.String())
	// mail.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer("smtp.gmail.com", 587, _email, _appPassword)
	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(mail); err != nil {
		return err
	}
	return nil
}
