/*
config.go
Created at 08.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package config

import (
	"encoding/json"
	"go-social/social/oauth"
	"go-social/social/oauth/oauth1"
	"go-social/social/oauth/oauth2"
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

	OAuthCredentials struct {
		Credentials oauth.Credentials `json:"credentials"`
	}

	OAuth1Config struct {
		Credentials oauth.Credentials `json:"credentials"`
		Token       oauth1.Token      `json:"token"`
	}

	OAuth2Config struct {
		Credentials oauth.Credentials `json:"credentials"`
		Token       oauth2.Token      `json:"token"`
	}
)

var Accounts *Config

func LoadConfig(path string) error {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	return json.Unmarshal(content, &Accounts)
}
