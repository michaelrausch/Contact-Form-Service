package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Destination struct {
	ID    string
	Name  string
	Email string
}

type MailjetConfig struct {
	Privatekey string
	Publickey  string
	From       string
	Name       string
	Subject    string
}

type ConfigFile struct {
	Destinations []Destination
	Mailjet      MailjetConfig
}

func (cfg *ConfigFile) GetDestinationById(id string) (error, Destination) {
	for _, dest := range cfg.Destinations {
		if dest.ID == id {
			return nil, dest
		}
	}

	return fmt.Errorf("Failed to find destination"), Destination{}
}

func readConf(filename string) (*ConfigFile, error) {
	buf, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	c := &ConfigFile{}
	err = yaml.Unmarshal(buf, c)

	if err != nil {
		return nil, err
	}

	return c, nil
}
