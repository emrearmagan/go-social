/*
follower.go
Created at 09.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package spotify

import (
	"github.com/emrearmagan/go-social/oauth/oauth2"
	"github.com/emrearmagan/go-social/social"
)

const (
	FollowerPath = "/v1/me/following"
)

// FollowerService provides methods for users followers
type FollowerService struct {
	oauth2 *oauth2.OAuth2
}

// newFollowerService returns a new Spotify FollowerService
func newFollowerService(oauth2 *oauth2.OAuth2) *FollowerService {
	return &FollowerService{
		oauth2: oauth2,
	}
}

// Following returns information about the authenticated users followers.
//TODO: not really testes
func (f *FollowerService) Following(params *FollowingParams) (*FollowingResponse, error) {
	followed := new(FollowingResponse)
	apiError := new(APIError)

	err := f.oauth2.Get(FollowerPath, followed, apiError, params)
	return followed, social.CheckError(err)
}

type FollowingParams struct {
	Type  string `url:"type,omitempty"`  // The ID type: currently only artist is supported.
	After int    `url:"after,omitempty"` // The last artist ID retrieved from the previous request.
	Limit int    `url:"limit,omitempty"` // The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50.
}

type FollowingResponse struct {
	Artists struct {
		Items []struct {
			ExternalUrls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Followers struct {
				Href  interface{} `json:"href"`
				Total int         `json:"total"`
			} `json:"followers"`
			Genres []string `json:"genres"`
			Href   string   `json:"href"`
			ID     string   `json:"id"`
			Images []struct {
				Height int    `json:"height"`
				URL    string `json:"url"`
				Width  int    `json:"width"`
			} `json:"images"`
			Name       string `json:"name"`
			Popularity int    `json:"popularity"`
			Type       string `json:"type"`
			URI        string `json:"uri"`
		} `json:"items"`
		Next    interface{} `json:"next"`
		Total   int         `json:"total"`
		Cursors struct {
			After interface{} `json:"after"`
		} `json:"cursors"`
		Limit int    `json:"limit"`
		Href  string `json:"href"`
	} `json:"artists"`
}
