package mailjet

import (
	"fmt"

	"github.com/kwandapchumba/bookmarkmonster/utils"
	"github.com/mailjet/mailjet-apiv3-go/v4"
)

type EmailVerificationMail struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

func NewEmailVerificationMail(email, code string) *EmailVerificationMail {
	return &EmailVerificationMail{
		Email: email,
		Code:  code,
	}
}

func (m EmailVerificationMail) SendEmailVificationMail() error {
	config, err := utils.LoadConfig(".")
	if err != nil {
		return fmt.Errorf("could not load config: %v", err)
	}

	client := mailjet.NewMailjetClient(config.MailJetApiKey, config.MailJetSecretKey)

	link := fmt.Sprintf("https://beta.bookmarkmonster.xyz/signup/verifyEmailAddress?token=%s", m.Code)

	body := fmt.Sprintf(`
	<p>Hey</p>

    <p>Verify your email address by clicking the following link:</p>

    <p>%s</p>

    <p>Regards,</p>
	<p>Bookmark Monster Team.</p>
    `, link)

	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: "hey@bookmarkmonster.xyz",
				Name:  "Bookmark Monster Team",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: m.Email,
				},
			},
			Subject:  "Confirm your email address",
			TextPart: body,
			HTMLPart: body,
		},
	}

	messages := mailjet.MessagesV31{Info: messagesInfo}

	_, err = client.SendMailV31(&messages)
	if err != nil {
		return fmt.Errorf("could not send email verification email: %v", err)
	}

	return nil
}
