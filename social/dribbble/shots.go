/*
shots.go
Created at 08.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package dribbble

import (
	"go-social/social"
	"go-social/social/oauth/oauth2"
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

	err := get(s.oauth2, ShotsPath, shots, apiError, nil)
	return shots, social.CheckError(err)
}
