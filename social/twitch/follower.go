/*
follower.go
Created at 19.06.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package twitch

import (
	"github.com/emrearmagan/go-social/oauth/oauth2"
	"github.com/emrearmagan/go-social/social"
	"time"
)

const (
	FollowerPath = "/helix/users/follows/"
)

// FollowerService provides methods for information about the authenticated users followers
type FollowerService struct {
	oauth2 *oauth2.OAuth2
}

// newFollowerService returns a new Twitch FollowerService.
func newFollowerService(oauth2 *oauth2.OAuth2) *FollowerService {
	return &FollowerService{
		oauth2: oauth2,
	}
}

// GetFollower  Gets information on follow relationships between two Twitch users. This can return information like “who is qotrok following,” “who is following qotrok,” or “is user X following user Y.” Information returned is sorted in order, most recent follow first.
// https://dev.twitch.tv/docs/api/reference#get-users-follows
// Required scopes: -
func (s *FollowerService) GetFollower(params FollowerParams) (*FollowerResp, error) {
	subs := new(FollowerResp)
	apiError := new(APIError)

	err := s.oauth2.Get(FollowerPath, subs, apiError, params)
	return subs, social.CheckError(err)
}

// FollowerParams are the params for GetFollower.
// At minimum, from_id or to_id must be provided for a query to be valid.
type FollowerParams struct {
	// FromId. user id. The request returns information about users who are being followed by the FromId user.
	FromId string `url:"from_id"`
	// ToId. user ID. The request returns information about users who are following the ToId user.
	ToId string `url:"to_id"`

	//Optional params
	// After. Cursor for forward pagination: tells the server where to start fetching the next set of results, in a multi-page response. The cursor value specified here is from the pagination response field of a prior query.
	After string `url:"after,omitempty"`
	// First Maximum number of objects to return. Maximum: 100. Default: 20.
	First int `url:"first,omitempty"`
}

type FollowerResp struct {
	Total      int    `json:"total"`
	Data       []Data `json:"data"`
	Pagination struct {
		Cursor string `json:"cursor"`
	} `json:"pagination"`
}

type Data struct {
	FromId     string    `json:"from_id"`
	FromLogin  string    `json:"from_login"`
	FromName   string    `json:"from_name"`
	ToId       string    `json:"to_id"`
	ToName     string    `json:"to_name"`
	FollowedAt time.Time `json:"followed_at"`
}
