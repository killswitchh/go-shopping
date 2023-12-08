package handlers

import (
	"crypto/tls"
	"fmt"

	"github.com/rs/zerolog/log"
	"gopkg.in/gomail.v2"

	"go-notification-service/models"
	"go-notification-service/utils"
)

func SendMail(receiver models.Order) {
	m := gomail.NewMessage()
	senderEmail := utils.GetEnvVar("SENDER_EMAIL")
	m.SetHeader("From", senderEmail)
	m.SetHeader("To", receiver.Email)
	m.SetHeader("Subject", "From Go-Shoppy")
	m.SetBody("text/plain", fmt.Sprintf("Greetings from Go-Shoppy! You have ordered %s which is %f dollars", receiver.Title, receiver.Price))

	d := gomail.NewDialer("smtp.gmail.com", 587, "guruprashanth78@gmail.com", "fphlpsgkvvleadip")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		fmt.Print("error")
		log.Err(err).Msg("error occured")
	}

	log.Info().Msgf("send email %s", receiver.Email)
}