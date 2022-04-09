/*
spotify.go
Created at 09.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package spotify

import (
	"github.com/emrearmagan/go-social/models"
	"github.com/emrearmagan/go-social/oauth/oauth2"
	"strings"
)

const (
	Base                = "https://api.spotify.com/"
	AuthorizationPrefix = "Bearer " // trailing space is required
)

type Client struct {
	Account  *AccountService
	User     *UserService
	Playlist *PlaylistService
	Follower *FollowerService
}

// NewClient returns a new Spotify Client.
func NewClient(oauth *oauth2.OAuth2) *Client {
	oauth = oauth.NewClient(oauth.Client().Base(Base))
	oauth.AuthorizationPrefix = AuthorizationPrefix
	return &Client{
		Account:  newAccountService(oauth),
		User:     newUserService(oauth),
		Playlist: newPlaylistService(oauth),
		Follower: newFollowerService(oauth),
	}
}

func (s *Client) GoSocialUser() (*models.SocialUser, error) {
	p, err := s.Playlist.UserPlaylists(nil)
	if err != nil {
		return nil, err
	}

	f, err := s.Follower.Following(&FollowingParams{
		Type: "artist",
	})
	if err != nil {
		return nil, err
	}

	u, err := s.User.UserCredentials()
	if err != nil {
		return nil, err
	}

	avatarUrl := ""
	if len(u.Images) > 0 {
		avatarUrl = u.Images[0].URL
	}
	goSocial := models.SocialUser{
		Username:     u.DisplayName,
		Name:         u.DisplayName,
		UserId:       u.ID,
		Verified:     strings.ToLower(u.Product) == "premium",
		ContentCount: int64(p.Total),
		Following:    &f.Artists.Total,
		AvatarUrl:    avatarUrl,
		Followers:    u.Followers.Total,
		Url:          u.ExternalUrls.Spotify,
	}

	return &goSocial, nil
}
