/*
following.go
Created at 21.06.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package github

import (
	"github.com/emrearmagan/go-social/oauth/oauth2"
	"github.com/emrearmagan/go-social/social"
)

const (
	FollowingPath = "/user/following"
)

// FollowingService provides methods for following information
type FollowingService struct {
	oauth2 *oauth2.OAuth2
}

// newFollowingService returns a new GitHub FollowingService.
func newFollowingService(oauth2 *oauth2.OAuth2) *FollowingService {
	return &FollowingService{
		oauth2: oauth2,
	}
}

// FollowingIds returns the ids of the following for the authenticated user.
// https://docs.github.com/en/rest/reference/users#list-the-people-the-authenticated-user-follows
func (f *FollowingService) FollowingIds(params *UserFollowerIdParams) (*FollowersIdResponse, error) {
	following := new(FollowersIdResponse)
	apiError := new(APIError)

	err := f.oauth2.Get(FollowingPath, following, apiError, params)
	return following, social.CheckError(err)
}
