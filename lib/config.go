// Copyright 2021 Michael Rausch. All Rights Reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lib

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Destination struct {
	ID    string // A unique ID to identify this destination
	Name  string // The name of this destination
	Email string // The email where the messages will be sent
}

type MailjetConfig struct {
	Privatekey string // The API keys provided by MailJet
	Publickey  string // ^
	From       string // This is the email address that messages will be sent from
	Name       string // This is the name that messages will be sent from
	Subject    string // This is the subject that all messages will have
}

type ConfigFile struct {
	Destinations []Destination
	Mailjet      MailjetConfig
}

// Gets the destination for a given ID
func (cfg *ConfigFile) GetDestinationById(id string) (error, Destination) {
	for _, dest := range cfg.Destinations {
		if dest.ID == id {
			return nil, dest
		}
	}

	return fmt.Errorf("Failed to find destination"), Destination{}
}

// Read the configuration file from disk
func ReadConf(filename string) (*ConfigFile, error) {
	buf, err := ioutil.ReadFile(filename)

	// Error reading the file
	if err != nil {
		return nil, err
	}

	c := &ConfigFile{}
	err = yaml.Unmarshal(buf, c)

	// Error parsing the YAML file
	if err != nil {
		return nil, err
	}

	return c, nil
}
