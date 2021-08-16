// Copyright 2021 Michael Rausch. All Rights Reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/michaelrausch/Contact-Form-Service/lib"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", ContactHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}

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
func ContactHandler(w http.ResponseWriter, r *http.Request) {
	var m lib.ContactMessage

	config, err := lib.ReadConf("conf.yaml")

	// There was an error loading or parsing config.yaml
	// Return an internal server error
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Internal Server Error")
		return
	}

	err = json.NewDecoder(r.Body).Decode(&m)

	// Failed to decode body as JSON, or validation failed
	if err != nil || m.Validate() == false {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Bad Request")
		return
	}

	err, dest := config.GetDestinationById(m.Destination)

	// Failed to find a destination with the given ID
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Bad Request")
		return
	}

	mail := lib.MailMessage{
		Dest:    dest,
		Message: m,
	}

	err = mail.Send(config.Mailjet)

	// There was an error sending the message
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		fmt.Fprintf(w, "Failed To Send Message")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}
