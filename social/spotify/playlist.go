/*
playlist.go
Created at 09.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package spotify

import (
	"github.com/emrearmagan/go-social/oauth/oauth2"
	"github.com/emrearmagan/go-social/social"
)

const (
	PlaylistPath = "/v1/me/playlists"
)

// PlaylistService provides methods information about the users playlist
type PlaylistService struct {
	oauth2 *oauth2.OAuth2
}

// newAccountService returns a new Reddit UserService.
func newPlaylistService(oauth2 *oauth2.OAuth2) *PlaylistService {
	return &PlaylistService{
		oauth2: oauth2,
	}
}

// UserPlaylists returns the playlists for the authenticated user.
// https://developer.spotify.com/console/get-current-user/
func (p *PlaylistService) UserPlaylists(params *UserPlaylistParams) (*Playlist, error) {
	playlist := new(Playlist)
	apiError := new(APIError)

	err := p.oauth2.Get(PlaylistPath, playlist, apiError, params)
	return playlist, social.CheckError(err)
}

// UserPlaylistParams are the params for the Playlist endpoint.
type UserPlaylistParams struct {
	Limit  int `url:"limit,omitempty"`  // The maximum number of playlists to return. Default: 20. Minimum: 1. Maximum: 50.’
	offset int `url:"offset,omitempty"` // The index of the first playlist to return. Default: 0 (the first object). Maximum offset: 100.000. Use with limit to get the next set of playlists.’
}

type Playlist struct {
	Href  string `json:"href"`
	Items []struct {
		Collaborative bool   `json:"collaborative"`
		Description   string `json:"description"`
		ExternalUrls  struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Href   string `json:"href"`
		ID     string `json:"id"`
		Images []struct {
			Height int    `json:"height"`
			URL    string `json:"url"`
			Width  int    `json:"width"`
		} `json:"images"`
		Name  string `json:"name"`
		Owner struct {
			DisplayName  string `json:"display_name"`
			ExternalUrls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href string `json:"href"`
			ID   string `json:"id"`
			Type string `json:"type"`
			URI  string `json:"uri"`
		} `json:"owner"`
		PrimaryColor interface{} `json:"primary_color"`
		Public       bool        `json:"public"`
		SnapshotID   string      `json:"snapshot_id"`
		Tracks       struct {
			Href  string `json:"href"`
			Total int    `json:"total"`
		} `json:"tracks"`
		Type string `json:"type"`
		URI  string `json:"uri"`
	} `json:"items"`
	Limit    int         `json:"limit"`
	Next     interface{} `json:"next"`
	Offset   int         `json:"offset"`
	Previous interface{} `json:"previous"`
	Total    int         `json:"total"`
}
