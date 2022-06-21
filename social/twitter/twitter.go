/*
twitter.go
Created at 09.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package twitter

import (
	"context"
	"fmt"
	"github.com/emrearmagan/go-social/models"
	"github.com/emrearmagan/go-social/oauth"
	"github.com/emrearmagan/go-social/oauth/oauth1"
	"github.com/emrearmagan/go-social/social/client"
)

type Client struct {
	User     *UserService
	Follower *FollowerService
}

const (
	Base = "https://api.twitter.com/"
)

// NewClient returns a new Twitter Client.
func NewClient(ctx context.Context, c *oauth.Credentials, token *oauth1.Token) *Client {
	cl := client.NewHttpClient().Base(Base)
	auther := oauth1.NewOAuth(ctx, c, token, cl)

	return &Client{
		User:     newUserService(auther),
		Follower: newFollowerService(auther),
	}
}

func (r *Client) GoSocialUser() (*models.SocialUser, error) {
	u, err := r.User.UserCredentials(nil)
	if err != nil {
		return nil, err
	}

	goSocial := models.SocialUser{
		Username:     u.ScreenName,
		Name:         u.Name,
		UserId:       fmt.Sprintf("%v", u.ID),
		Verified:     u.Verified,
		ContentCount: int64(u.StatusCount),
		AvatarUrl:    u.ProfileImageURL,
		Followers:    u.FollowersCount,
		Following:    &u.FriendsCount,
		Url:          fmt.Sprintf("https://twitter.com/%s", u.ScreenName),
	}

	return &goSocial, nil
}
