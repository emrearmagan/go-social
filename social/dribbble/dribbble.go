/*
dribbble.go
Created at 08.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package dribbble

import (
	"go-social/models"
	"go-social/social"
	"go-social/social/oauth/oauth2"
	"strconv"
)

const (
	apiHost             = "https://api.dribbble.com/"
	authorizationPrefix = "Bearer " // trailing space is required
)

type Client struct {
	User  *UserService
	Shots *ShotService
}

// NewClient returns a new Dribbble Client.
func NewClient(oauth *oauth2.OAuth2) *Client {
	oauth = oauth.NewClient(oauth.Client().Base(apiHost))
	return &Client{
		User:  newUserService(oauth.New()),
		Shots: newShotService(oauth.New()),
	}
}

func (d *Client) GoSocialUser() (*models.SocialUser, error) {
	s, err := d.Shots.DribbbleShots()
	if err != nil {
		return nil, social.CheckError(err)
	}

	u, err := d.User.UserCredentials()
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

func get(oauth2 *oauth2.OAuth2, path string, resp interface{}, apiError *APIError, params interface{}) error {
	client := oauth2.Client().AddQuery(params).Get(path)

	req, err := oauth2.AddHeader(authorizationPrefix)
	if err != nil {
		return err
	}

	httpResp, err := client.Do(req, resp, &apiError.Errors)
	if httpResp != nil {
		if code := httpResp.StatusCode; code >= 300 {
			apiError.StatusCode = httpResp.StatusCode
		}
	}

	return social.RelevantError(err, apiError)
}
