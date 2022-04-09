/*
credentials.go
Created at 08.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package oauth

// Credentials represents an OAuth consumer's (api) key and (api) secret
type Credentials struct {
	// Consumer Key (API key)
	ConsumerKey string `json:"consumer_key"`
	// Consumer Secret (API shared secret)
	ConsumerSecret string `json:"consumer_secret"`
}

// NewCredentials returns a new Credentials with the given consumer key and secret.
func NewCredentials(consumerKey, consumerSecret string) *Credentials {
	return &Credentials{
		ConsumerKey:    consumerKey,
		ConsumerSecret: consumerSecret,
	}
}
