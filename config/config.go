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
	SocialAccounts struct {
		Github   OAuth `json:"github"`
		Dribbble OAuth `json:"dribbble"`
		Spotify  OAuth `json:"spotify"`
		Reddit   OAuth `json:"reddit"`
		Facebook OAuth `json:"facebook"`
		Twitter  OAuth `json:"twitter"`
		Tumblr   OAuth `json:"tumblr"`
	}

	OAuth struct {
		ConsumerKey    string `json:"consumer_key"`
		ConsumerSecret string `json:"consumer_secret"`
	}

	OAuth1Config struct {
		OAuthConfig OAuth `json:"oauth_config"`
	}

	OAuth2Config struct {
		OAuthConfig OAuth `json:"oauth_config"`
	}
)

var Accounts *SocialAccounts

func LoadConfig(path string) error {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	return json.Unmarshal(content, &Accounts)
}
