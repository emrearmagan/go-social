/*
dribbble.go
Created at 08.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package dribbble

import (
	"go-social/config"
	"go-social/models"
	"go-social/social"
	"strconv"
	"strings"
)

const (
	apiHost   = "https://api.dribbble.com/"
	userPath  = "/v2/user/"
	shotsPath = "/v2/user/shots"

	authorizationHeaderName = "Authorization"
	authorizationPrefix     = "Bearer " // trailing space is required
)

type Client struct {
	httpClient *social.HttpClient
	Config     config.OAuth2Config
}

// NewClient returns a new Dribbble Client.
func NewClient(config config.OAuth2Config) *Client {
	httpClient := social.NewClient().Base(apiHost)
	return &Client{
		httpClient: httpClient,
		Config:     config,
	}
}

func (d *Client) get(path string, resp interface{}, apiError *APIError) error {
	client := d.httpClient.Get(path)

	req, err := client.Request()
	if err != nil {
		return err
	}

	header := []string{authorizationPrefix, d.Config.Token}
	req.Header.Set(authorizationHeaderName, strings.Join(header, ""))

	httpResp, err := client.Do(req, resp, &apiError.Errors)
	if httpResp != nil {
		if code := httpResp.StatusCode; code >= 300 {
			apiError.StatusCode = httpResp.StatusCode
		}
	}

	return social.RelevantError(err, apiError)
}

// DribbbleShots returns all shots for the authenticated user.
// See: https://developer.dribbble.com/v2/shots/#list-shots for more information
func (d *Client) DribbbleShots() (*Shots, error) {
	shots := new(Shots)
	apiError := new(APIError)

	err := d.get(shotsPath, &shots, apiError)
	if err != nil {
		return nil, err
	}

	return shots, social.CheckError(err)
}

// UserCredentials returns the user credentials for the authenticated user.
// https://developer.dribbble.com/v2/user/
func (d *Client) UserCredentials() (*User, error) {
	user := new(User)
	apiError := new(APIError)

	err := d.get(userPath, &user, apiError)
	if err != nil {
		return nil, err
	}

	return user, social.CheckError(err)
}

func (d *Client) GoSocialUser() (*models.SocialUser, error) {
	s, err := d.DribbbleShots()
	if err != nil {
		return nil, social.CheckError(err)
	}

	u, err := d.UserCredentials()
	if err != nil {
		return nil, social.CheckError(err)
	}

	goSocial := models.SocialUser{
		Username:     u.Login,
		Name:         u.Name,
		UserId:       strconv.Itoa(u.ID),
		ContentCount: int64(len(*s)),
		Verified:     u.Pro,
		AvatarUrl:    u.AvatarURL,
		Followers:    u.FollowersCount,
		Url:          u.HTMLURL,
	}

	return &goSocial, nil
}
