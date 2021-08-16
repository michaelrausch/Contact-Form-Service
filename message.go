package main

import (
	"regexp"
	"strings"
)

var rxEmail = regexp.MustCompile(".+@.+\\..+")

type ContactMessage struct {
	Name        string
	Email       string
	Message     string
	Destination string
}

func (msg *ContactMessage) Validate() bool {
	emailRegex := rxEmail.Match([]byte(msg.Email))

	if msg.Name == "" || msg.Email == "" || msg.Destination == "" || msg.Message == "" {
		return false
	}

	if emailRegex == false {
		return false
	}

	if strings.TrimSpace(msg.Message) == "" {
		return false
	}

	return true
}
