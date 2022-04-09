/*
config.go
Created at 08.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package config

import (
	"encoding/json"
	"io/ioutil"
)

type (
	Config struct {
		Github   OAuth2Config `json:"github"`
		Dribbble OAuth2Config `json:"dribbble"`
		Spotify  OAuth2Config `json:"spotify"`
		Reddit   OAuth2Config `json:"reddit"`
		Facebook OAuth2Config `json:"facebook"`
		Twitter  OAuth1Config `json:"twitter"`
		Tumblr   OAuth1Config `json:"tumblr"`
	}

	OAuth1Config struct {
		ConsumerKey    string `json:"consumer_key"`
		ConsumerSecret string `json:"consumer_secret"`
		Token          string `json:"access_token"`
		TokenSecret    string `json:"token_secret"`
	}

	OAuth2Config struct {
		ConsumerKey    string `json:"consumer_key"`
		ConsumerSecret string `json:"consumer_secret"`
		Token          string `json:"access_token"`
		RefreshToken   string `json:"refresh_token"`
	}
)

func LoadConfig(path string) (*Config, error) {
	var accounts *Config

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(content, &accounts); err != nil {
		return nil, err
	}

	return accounts, nil
}
