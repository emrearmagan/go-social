/*
subscriber.go
Created at 19.06.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package twitch

import (
	"github.com/emrearmagan/go-social/oauth/oauth2"
	"github.com/emrearmagan/go-social/social"
)

const (
	SubscriberPath = "/helix/subscriptions/"
)

// SubscriberService provides methods for information about the authenticated users subscribers
type SubscriberService struct {
	oauth2 *oauth2.OAuth2
}

// newSubscriberService returns a new Twitch SubscriberService.
func newSubscriberService(oauth2 *oauth2.OAuth2) *SubscriberService {
	return &SubscriberService{
		oauth2: oauth2,
	}
}

// BroadcasterSubscriptions  Gets a list of users that subscribe to the specified broadcaster.
// https://dev.twitch.tv/docs/api/reference#get-broadcaster-subscriptions
// Required scopes: channel:read:subscriptions
func (s *SubscriberService) BroadcasterSubscriptions(params SubscriberParams) (*Subscribers, error) {
	subs := new(Subscribers)
	apiError := new(APIError)

	err := s.oauth2.Get(SubscriberPath, subs, apiError, params)
	return subs, social.CheckError(err)
}

// SubscriberParams are the params for BroadcasterSubscriptions.
type SubscriberParams struct {
	// BroadCasterId. ID of the broadcaster. Must match the User ID in the Bearer token. Required
	BroadCasterId string `url:"broadcaster_id"`

	//Optional params
	// UserId Filters the list to include only the specified subscribers. To specify more than one subscriber, include this parameter for each subscriber. For example, &user_id=1234&user_id=5678. You may specify a maximum of 100 subscribers.
	UserId string `url:"user_id,omitempty"`
	// After. Cursor for forward pagination: tells the server where to start fetching the next set of results in a multi-page response. This applies only to queries without user_id. If a user_id is specified, it supersedes any cursor/offset combinations. The cursor value specified here is from the pagination response field of a prior query.
	After string `url:"after,omitempty"`
	// First Maximum number of objects to return. Maximum: 100. Default: 20.
	First string `url:"first,omitempty"`
}

type Subscribers struct {
	Broadcaster []Broadcaster `json:"data"`
	Pagination  struct {
		Cursor string `json:"cursor"`
	} `json:"pagination"`
	Total  int `json:"total"`
	Points int `json:"points"`
}

type Broadcaster struct {
	BroadcasterId    string `json:"broadcaster_id"`
	BroadcasterLogin string `json:"broadcaster_login"`
	BroadcasterName  string `json:"broadcaster_name"`
	GifterId         string `json:"gifter_id"`
	GifterLogin      string `json:"gifter_login"`
	GifterName       string `json:"gifter_name"`
	IsGift           bool   `json:"is_gift"`
	Tier             string `json:"tier"`
	PlanName         string `json:"plan_name"`
	UserId           string `json:"user_id"`
	UserName         string `json:"user_name"`
	UserLogin        string `json:"user_login"`
}
