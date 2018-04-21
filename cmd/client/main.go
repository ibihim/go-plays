package main

import (
	"github.com/ibihim/go-plays/pkg/mailclient"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	var cfg mailclient.Config

	kingpin.
		Flag("user-name", "the user name used for accessing mail provider's API").
		Envar("USER_NAME").
		StringVar(&cfg.UserName)

	kingpin.
		Flag("password", "the password used for accessing mail provider's API").
		Envar("PASSWORD").
		StringVar(&cfg.Password)

	kingpin.
		Flag("sender-name", "the sender's name in the email").
		Envar("SENDER_NAME").
		StringVar(&cfg.SenderName)

	kingpin.
		Flag("sender-email", "the sender's email address in the email").
		Envar("SENDER_EMAIL").
		StringVar(&cfg.SenderEmail)

	c := mailclient.New(&cfg)

	c.sendEmail("Subjet", "Body")
}
