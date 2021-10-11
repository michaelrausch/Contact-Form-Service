// Copyright 2021 Michael Rausch. All Rights Reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/michaelrausch/Contact-Form-Service/routes"
)

func main() {
	if len(os.Args) != 3 || os.Args[1] != "-c" {
		fmt.Println("Usage: api -c config.yaml")
		return
	}

	app := fiber.New()

	app.Post("/", routes.ContactMessage)

	log.Fatal(app.Listen(":8080"))
}
