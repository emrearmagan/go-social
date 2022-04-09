/*
shots.go
Created at 08.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package dribbble

import (
	"go-social/social"
	"go-social/social/oauth/oauth2"
	"time"
)

const (
	ShotsPath = "/v2/user/shots"
)

// ShotService provides a method for account credential verification.
type ShotService struct {
	oauth2 *oauth2.OAuth2
}

// newUserService returns a new AccountService.
func newShotService(oauth2 *oauth2.OAuth2) *ShotService {
	return &ShotService{
		oauth2: oauth2,
	}
}

// DribbbleShots returns all shots for the authenticated user.
// See: https://developer.dribbble.com/v2/shots/#list-shots for more information
func (s *ShotService) DribbbleShots() (*Shots, error) {
	shots := new(Shots)
	apiError := new(APIError)

	err := s.oauth2.Get(UserPath, shots, apiError, nil)

	return shots, social.CheckError(err)
}

// Shots represents the Shots of the Dribbble User
// https://developer.dribbble.com/v2/shots/
type Shots []struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	Images      struct {
		Hidpi  interface{} `json:"hidpi"`
		Normal string      `json:"normal"`
		Teaser string      `json:"teaser"`
	} `json:"images"`
	PublishedAt time.Time `json:"published_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	HTMLURL     string    `json:"html_url"`
}
