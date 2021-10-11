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

	// There was an error loading or parsing config.yaml
	// Return an internal server error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	if err := c.BodyParser(m); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}

	// Failed to decode body as JSON, or validation failed
	if m.Validate() == false {
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}

	err, dest := config.GetDestinationById(m.Destination)

	// Failed to find a destination with the given ID
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}

	mail := lib.MailMessage{
		Dest:    dest,
		Message: *m,
	}

	err = mail.Send(config.Mailjet)

	// There was an error sending the message
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")

	}

	return c.Status(fiber.StatusOK).SendString("OK")
}
