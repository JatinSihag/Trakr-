package utils

import (
	"fmt"
	"math/rand"
	"time"

	"gopkg.in/gomail.v2"
)

func GenerateOTP() string {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := rng.Intn(900000) + 100000
	return fmt.Sprintf("%d", code)
}
func SendEmail(toEmail string, code string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "jatinsihag18@gmail.com")
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", "Trackr Verification Code")
	m.SetBody("text/html", fmt.Sprintf("<h1>Welcome to Trackr!</h1><p>Your verification code is: <b>%s</b></p>", code))
	// d:=gomail.NewDialer("smtp.gmail.com",587,"jatinsihag18@gmail.com","password")
	// return d.DialAndSend(m)
	fmt.Println("=================================")
	fmt.Printf("📧 EMAIL TO %s: Your Code is %s\n", toEmail, code)
	fmt.Println("=================================")
	return nil
}
