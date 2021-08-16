// Copyright 2021 Michael Rausch. All Rights Reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lib

import (
	mailjet "github.com/mailjet/mailjet-apiv3-go"
)

type MailMessage struct {
	Message ContactMessage
	Dest    Destination
}

// Get the message body formatted as an HTML string
func (msg *MailMessage) GetHtmlPart() string {
	messageFrom := msg.Message.Email
	messageFromName := msg.Message.Name
	message := msg.Message.Message

	body := "<b>New contact form message from " + messageFromName + " (" + messageFrom + ")" + "</b>"
	body += "<br/><br/>"
	body += message
	body += "<br/><br/><i>Please reply to: " + messageFrom + "</i>"

	return body
}

// Get the message body formatted as a plain-text string
func (msg *MailMessage) GetTextPart() string {
	messageFrom := msg.Message.Email
	messageFromName := msg.Message.Name
	message := msg.Message.Message

	body := "New Contact Form Message From " + messageFromName + " (" + messageFrom + ")" + "\n\n"
	body += message
	body += "\n\nPlease reply to: " + messageFrom

	return body
}

// Send the message
//
// config: The MailJet configuration, this should include the API keys
//         See the definition of MailjetConfig in mail.go
//
func (msg *MailMessage) Send(config MailjetConfig) error {
	mailjetClient := mailjet.NewMailjetClient(config.Publickey, config.Privatekey)

	messageInfo := []mailjet.InfoMessagesV31{
		mailjet.InfoMessagesV31{
			From: &mailjet.RecipientV31{
				Email: config.From,
				Name:  config.Name,
			},
			ReplyTo: &mailjet.RecipientV31{
				Email: msg.Message.Email,
				Name:  msg.Message.Name,
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: msg.Dest.Email,
					Name:  msg.Dest.Name,
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
