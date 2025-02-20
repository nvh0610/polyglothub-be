package send_otp

import (
	"gopkg.in/gomail.v2"
	"learn/pkg/config"
)

type SendOtpEmail struct {
	dialer *gomail.Dialer
	Email  string
	ApiKey string
}

func NewSendOtpEmail() *SendOtpEmail {
	email := config.StringEnv("SENDER_EMAIL")
	apiKey := config.StringEnv("API_KEY_EMAIL")
	return &SendOtpEmail{
		dialer: gomail.NewDialer("smtp.gmail.com", 587, email, apiKey),
		Email:  email,
		ApiKey: apiKey,
	}
}

func (s *SendOtpEmail) SendOtp(toEmail, otp string) error {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", s.Email)
	mailer.SetHeader("To", toEmail)
	mailer.SetHeader("Subject", "Your OTP Code")
	mailer.SetBody("text/plain", "Your OTP is: "+otp)

	if err := s.dialer.DialAndSend(mailer); err != nil {
		return err
	}
	return nil
}
