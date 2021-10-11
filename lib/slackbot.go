package lib

import (
	"github.com/slack-go/slack"
)

type SlackBot struct {
	ApiToken  string
	ChannelID string
}

func (bot *SlackBot) SendMessage(message string) error {
	api := slack.New(bot.ApiToken)

	_, _, err := api.PostMessage(bot.ChannelID, slack.MsgOptionText(message, false), slack.MsgOptionAsUser(false))

	return err
}

func (bot *SlackBot) SendContactFormMessage(message ContactMessage, destination string) error {
	body := "New Message for " + destination + "\n\nFrom: " + message.Email + "(" + message.Name + ")\n\n" + message.Message
	return bot.SendMessage(body)
}

func (bot *SlackBot) SendStatusMessage(level string, instance string, message string) error {
	return bot.SendMessage("[" + level + "] (" + instance + ") " + message)
}
