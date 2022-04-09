/*
config.go
Created at 08.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package config

import (
	"encoding/json"
	"github.com/emrearmagan/go-social/oauth"
	"github.com/emrearmagan/go-social/oauth/oauth1"
	"github.com/emrearmagan/go-social/oauth/oauth2"
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
		Token       oauth1.Token      `json:"token"`
		Credentials oauth.Credentials `json:"credentials"`
	}

	OAuth2Config struct {
		Token       oauth2.Token      `json:"token"`
		Credentials oauth.Credentials `json:"credentials"`
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
