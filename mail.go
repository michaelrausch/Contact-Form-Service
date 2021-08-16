package main

import (
	mailjet "github.com/mailjet/mailjet-apiv3-go"
)

type MailMessage struct {
	message ContactMessage
	dest    Destination
}

func (msg *MailMessage) GetHtmlPart() string {
	messageFrom := msg.message.Email
	messageFromName := msg.message.Name
	message := msg.message.Message

	body := "<b>New contact form message from " + messageFromName + " (" + messageFrom + ")" + "</b>"
	body += "<br/><br/>"
	body += message
	body += "<br/><br/><i>Please reply to: " + messageFrom + "</i>"

	return body
}

func (msg *MailMessage) GetTextPart() string {
	messageFrom := msg.message.Email
	messageFromName := msg.message.Name
	message := msg.message.Message

	body := "New Contact Form Message From " + messageFromName + " (" + messageFrom + ")" + "\n\n"
	body += message
	body += "\n\nPlease reply to: " + messageFrom

	return body
}

func (msg *MailMessage) Send(config MailjetConfig) error {

	mailjetClient := mailjet.NewMailjetClient(config.Publickey, config.Privatekey)

	messageInfo := []mailjet.InfoMessagesV31{
		mailjet.InfoMessagesV31{
			From: &mailjet.RecipientV31{
				Email: config.From,
				Name:  config.Name,
			},
			ReplyTo: &mailjet.RecipientV31{
				Email: msg.message.Email,
				Name:  msg.message.Name,
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: msg.dest.Email,
					Name:  msg.dest.Name,
				},
			},
			Subject:  config.Subject,
			TextPart: msg.GetTextPart(),
			HTMLPart: msg.GetHtmlPart(),
		},
	}

	messages := mailjet.MessagesV31{Info: messageInfo}
	_, err := mailjetClient.SendMailV31(&messages)

	if err != nil {
		return err
	}

	return nil
}
