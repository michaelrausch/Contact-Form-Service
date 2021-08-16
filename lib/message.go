// Copyright 2021 Michael Rausch. All Rights Reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lib

import (
	"regexp"
	"strings"
)

type ContactMessage struct {
	Name        string
	Email       string
	Message     string
	Destination string
}

// Validate that all fields have been submitted
// and that the email address is valid
//
// Returns: true if validation passes
func (msg *ContactMessage) Validate() bool {
	emailRegex := regexp.MustCompile(".+@.+\\..+")
	emailRegexResult := emailRegex.Match([]byte(msg.Email))

	if msg.Name == "" || msg.Email == "" || msg.Destination == "" {
		return false
	}

	if emailRegexResult == false {
		return false
	}

	if strings.TrimSpace(msg.Message) == "" {
		return false
	}

	return true
}
