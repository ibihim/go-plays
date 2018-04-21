package mailclient

import (
	"net/http"
)

const (
	mailJetSendURL = "https://api.mailjet.com/v3.1/send"
)

// Config that holds the credentials for the email provider's API access
type Config struct {
	UserName, Password      string
	SenderName, SenderEmail string
}

// New returns a new email client
func New(cfg *Config) *MailJet {
	return &MailJet{
		client:      newClient(cfg.UserName, cfg.Password),
		mailJetURL:  mailJetSendURL,
		senderName:  cfg.SenderEmail,
		senderEmail: cfg.SenderName,
	}
}

// MailJet returns a MailJetClient
type MailJet struct {
	client                  *http.Client
	mailJetURL              string
	senderName, senderEmail string
}

func newClient(userName, password string) *http.Client {
	return &http.Client{
		Transport: &authTransport{
			Transport: &http.Transport{},
			userName:  userName,
			password:  password,
		},
	}
}

type authTransport struct {
	*http.Transport
	userName, password string
}

func (t *authTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.SetBasicAuth(t.userName, t.password)

	return t.Transport.RoundTrip(req)
}
