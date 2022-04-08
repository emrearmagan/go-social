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
	APIHost   = "https://api.dribbble.com/"
	UserPath  = "/v2/user/"
	ShotsPath = "/v2/user/shots"

	AuthorizationHeaderName = "Authorization"
	authorizationPrefix     = "Bearer " // trailing space is required
)

func get(token string, path string, resp interface{}, apiError *APIError) error {
	client := social.NewClient().Base(APIHost).Get(path)

	req, err := client.Request()
	if err != nil {
		return err
	}

	header := []string{authorizationPrefix, token}
	req.Header.Set(AuthorizationHeaderName, strings.Join(header, ""))

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
func DribbbleShots(config config.OAuth2Config) (*Shots, error) {
	shots := new(Shots)
	apiError := new(APIError)

	err := get(config.Token, ShotsPath, &shots, apiError)
	if err != nil {
		return nil, err
	}

	return shots, social.CheckError(err)
}

// UserCredentials returns the user credentials for the authenticated user.
// https://developer.dribbble.com/v2/user/
func UserCredentials(config config.OAuth2Config) (*User, error) {
	user := new(User)
	apiError := new(APIError)

	err := get(config.Token, UserPath, &user, apiError)
	if err != nil {
		return nil, err
	}

	return user, social.CheckError(err)
}

func GoSocialUser(config config.OAuth2Config) (*models.SocialUser, error) {
	s, err := DribbbleShots(config)
	if err != nil {
		return nil, social.CheckError(err)
	}

	u, err := UserCredentials(config)
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
