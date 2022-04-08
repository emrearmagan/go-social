/*
follower.go
Created at 08.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package github

import (
	"go-social/social"
	"go-social/social/oauth/oauth2"
)

const (
	FollowerPath  = "/user/followers"
	FollowingPath = "/user/following"
)

// FollowerService provides methods for follower information
type FollowerService struct {
	oauth2 *oauth2.OAuth2
}

// newFollowerService returns a new GitHub FollowerService.
func newFollowerService(oauth2 *oauth2.OAuth2) *FollowerService {
	return &FollowerService{
		oauth2: oauth2,
	}
}

// FollowerIds returns the ids of the follower for the authenticated user.
// https://docs.github.com/en/rest/reference/users#list-followers-of-the-authenticated-user
func (f *FollowerService) FollowerIds(cursor int64, max *int) (*FollowersIdResponse, error) {
	followers := new(FollowersIdResponse)
	apiError := new(APIError)

	params := UserFollowerIdParams{
		PerPage: max,
		Page:    int(cursor),
	}
	err := get(f.oauth2, FollowerPath, &followers, apiError, params)

	return followers, social.CheckError(err)
}

// FollowingIds returns the ids of the following for the authenticated user.
// https://docs.github.com/en/rest/reference/users#list-the-people-the-authenticated-user-follows
func (f *FollowerService) FollowingIds(cursor int64, max *int) (*FollowersIdResponse, error) {
	followers := new(FollowersIdResponse)
	apiError := new(APIError)

	params := UserFollowerIdParams{
		PerPage: max,
		Page:    int(cursor),
	}
	err := get(f.oauth2, FollowingPath, &followers, apiError, params)

	return followers, social.CheckError(err)
}

// UserFollowerIdParams are the parameters for FollowerIds and FollowingIds
type UserFollowerIdParams struct {
	PerPage *int `url:"per_page,omitempty"` // PerPage per page (max 100), Default: 30
	Page    int  `url:"page,omitempty"`     // Page number of the results to fetch, Default: 1
}
