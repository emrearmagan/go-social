/*
playlist.go
Created at 09.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package spotify

import (
	"go-social/social"
	"go-social/social/oauth/oauth2"
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

// UserPlaylistParams are the params for the Playlist endpoint. Currently not used
type UserPlaylistParams struct {
	Limit  int `url:"limit,omitempty"`  // The maximum number of playlists to return. Default: 20. Minimum: 1. Maximum: 50.’
	offset int `url:"offset,omitempty"` // The index of the first playlist to return. Default: 0 (the first object). Maximum offset: 100.000. Use with limit to get the next set of playlists.’
}
