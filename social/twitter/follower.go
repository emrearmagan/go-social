/*
follower.go
Created at 09.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package twitter

import (
	"github.com/emrearmagan/go-social/oauth/oauth1"
	"github.com/emrearmagan/go-social/social"
)

const (
	FollowerIdsPath  = "/1.1/followers/ids.json"
	FollowingIdsPath = "/1.1/friends/ids.json"
)

// FollowerService provides methods for information about the followers of the authenticated user
type FollowerService struct {
	oauth1 *oauth1.OAuth1
}

// newUserService returns a new Twitter UserService.
func newFollowerService(oauth1 *oauth1.OAuth1) *FollowerService {
	return &FollowerService{
		oauth1: oauth1,
	}
}

// FollowerIDs returns a cursored collection of Users following the authorized user.
// https://developer.twitter.com/en/docs/twitter-api/v1/accounts-and-users/follow-search-get-users/api-reference/get-followers-ids
func (f *FollowerService) FollowerIDs(params *FollowerIDParams) (*UserFollowerIDs, error) {
	ids := new(UserFollowerIDs)
	apiError := new(APIError)

	err := f.oauth1.Get(FollowerIdsPath, ids, apiError, params)
	return ids, social.CheckError(err)
}

// FollowingIDs returns a cursored collection of Users the authorized user is following.
// https://developer.twitter.com/en/docs/twitter-api/v1/accounts-and-users/follow-search-get-users/api-reference/get-friends-ids
func (f *FollowerService) FollowingIDs(params *FollowerIDParams) (*UserFollowerIDs, error) {
	ids := new(UserFollowerIDs)
	apiError := new(APIError)

	err := f.oauth1.Get(FollowingIdsPath, ids, apiError, params)
	return ids, social.CheckError(err)
}

// FollowerIDParams are the parameters for IDs
type FollowerIDParams struct {
	UserID     int64  `url:"user_id,omitempty"`
	ScreenName string `url:"screen_name,omitempty"`
	Cursor     int64  `url:"cursor,omitempty"`
	Count      *int   `url:"count,omitempty"`
}

// UserFollowerIDs is a cursored collection of follower ids.
type UserFollowerIDs struct {
	IDs               []int64     `json:"ids"`
	NextCursor        int64       `json:"next_cursor"`
	NextCursorStr     string      `json:"next_cursor_str"`
	PreviousCursor    int64       `json:"previous_cursor"`
	PreviousCursorStr string      `json:"previous_cursor_str"`
	TotalCount        interface{} `json:"total_count"`
}
