// Copyright 2021 Michael Rausch. All Rights Reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/michaelrausch/Contact-Form-Service/lib"
	"github.com/michaelrausch/Contact-Form-Service/routes"
)

func sendSlackServerUpdate(config lib.ConfigFile) {
	slackBot := lib.SlackBot{
		ApiToken:  config.Slack.ApiKey,
		ChannelID: config.Slack.ChannelID,
	}

	slackBot.SendStatusMessage("INFO", config.ServerID, "Server Online")
}

func main() {
	if len(os.Args) != 3 || os.Args[1] != "-c" {
		fmt.Println("Usage: api -c config.yaml")
		return
	}

	configFilePath := os.Args[2]
	config, err := lib.ReadConf(configFilePath)

	if err != nil {
		fmt.Println("Failed To Read Config")
		return
	}

	go sendSlackServerUpdate(*config)

	app := fiber.New()

	app.Post("/", routes.ContactMessage)

	log.Fatal(app.Listen(":8080"))
}
