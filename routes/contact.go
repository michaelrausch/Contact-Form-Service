package routes

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/michaelrausch/Contact-Form-Service/lib"
)

// Handler for contact form messages
//
// Route: /
// Method: POST
// Example Request:
// {
//    "Name": "Joe Bloggs"
//    "Email": "example@email.com"
//    "Message": "Hello, World!"
//    "Destination": "personalwebsite"
// }
func ContactMessage(c *fiber.Ctx) error {
	m := new(lib.ContactMessage)
	configFilePath := os.Args[2]

	config, err := lib.ReadConf(configFilePath)

	slackBot := lib.SlackBot{
		ApiToken:  config.Slack.ApiKey,
		ChannelID: config.Slack.ChannelID,
	}

	// There was an error loading or parsing config.yaml
	// Return an internal server error
	if err != nil {
		go slackBot.SendStatusMessage("FATAL", config.ServerID, "Failed To Load Config")
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	if err := c.BodyParser(m); err != nil {
		go slackBot.SendStatusMessage("WARN", config.ServerID, "Failed to decode")
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}

	// Failed to decode body as JSON, or validation failed
	if m.Validate() == false {
		go slackBot.SendStatusMessage("WARN", config.ServerID, "Validation Failed")
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}

	m.Sanitize()

	err, dest := config.GetDestinationById(m.Destination)

	// Failed to find a destination with the given ID
	if err != nil {
		go slackBot.SendStatusMessage("WARN", config.ServerID, "Bad Destination "+dest.Name)
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}

	mail := lib.MailMessage{
		Dest:    dest,
		Message: *m,
	}

	go slackBot.SendContactFormMessage(*m, dest.ID)
	err = mail.Send(config.Mailjet)

	// There was an error sending the message
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")

	}

	return c.Status(fiber.StatusOK).SendString("OK")
}
