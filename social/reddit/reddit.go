/*
reddit.go
Created at 08.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package reddit

import (
	"go-social/social/oauth/oauth2"
)

const (
	Base = "https://oauth.reddit.com"

	RefreshHost = "https://www.reddit.com"
	RefreshPath = "/api/v1/access_token"

	authorizationPrefix = "bearer " // trailing space is required
)

type Client struct {
	User *UserService
}

// NewClient returns a new Dribbble Client.
func NewClient(oauth *oauth2.OAuth2, userAgent string) *Client {
	oauth = oauth.NewClient(oauth.Client().Base(Base))
	return &Client{
		User: newUserService(oauth.New(), userAgent),
	}
}

func get(oauth2 *oauth2.OAuth2, path string, resp interface{}, apiError *APIError, params interface{}, userAgent string) error {
	client := oauth2.Client().AddQuery(params).Get(path)

	req, err := oauth2.AddHeader(authorizationPrefix)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Content-Type", "application/json")

	httpResp, err := client.Do(req, resp, &apiError)
	if httpResp != nil {
		if code := httpResp.StatusCode; code >= 300 {
			apiError.ErrorCode = httpResp.StatusCode
			return apiError
		}
	}

	//TODO: Not checking relevant error since reddit does not provide a proper error
	return nil
}
