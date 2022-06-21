/*
github.go
Created at 08.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package github

import (
	"context"
	"fmt"
	"github.com/emrearmagan/go-social/models"
	"github.com/emrearmagan/go-social/oauth"
	"github.com/emrearmagan/go-social/oauth/oauth2"
	"github.com/emrearmagan/go-social/social/client"
	"strconv"
	"strings"
)

const (
	Base = "https://api.github.com/"
)

type Client struct {
	User      *UserService
	Follower  *FollowerService
	Following *FollowingService
}

// https://docs.github.com/en/rest/overview/resources-in-the-rest-api#user-agent-required

// NewClient returns a new GitHub Client.
// All API requests MUST include a valid User-Agent header. Requests with no User-Agent header will be rejected.
// It is requested to use the username or the application name
// See for more: https://docs.github.com/en/rest/overview/resources-in-the-rest-api#user-agent-required
func NewClient(ctx context.Context, c *oauth.Credentials, token *oauth2.Token, useragent *string) *Client {
	cl := client.NewHttpClient().Base(Base)
	if useragent != nil {
		cl.Add("User-Agent", *useragent)
	}

	auther := oauth2.NewOAuth(ctx, c, token, cl).Signer(GithubSigner{
		ConsumerKey:    c.ConsumerKey,
		ConsumerSecret: c.ConsumerSecret,
	})
	//oauth.AuthorizationPrefix = AuthorizationPrefix //TODO: Need different authorization header
	return &Client{
		User:      newUserService(auther),
		Follower:  newFollowerService(auther),
		Following: newFollowingService(auther),
	}
}

func (g *Client) GoSocialUser() (*models.SocialUser, error) {
	u, err := g.User.UserCredentials()
	if err != nil {
		return nil, err
	}

	pro := strings.ToLower(u.Plan.Name) == "pro"
	repos := u.PublicRepos + u.TotalPrivateRepos
	goSocial := models.SocialUser{
		Username:     u.Login,
		Name:         u.Name,
		UserId:       strconv.Itoa(u.ID),
		ContentCount: int64(repos),
		Verified:     pro,
		AvatarUrl:    u.AvatarURL,
		Followers:    u.Followers,
		Following:    &u.Following,
		Url:          u.HTMLURL,
	}

	return &goSocial, nil
}

// AuthorizationPrefix
// GitHub recommends sending OAuth tokens using the Authorization header.
// Read more about here: https://docs.github.com/en/rest/overview/resources-in-the-rest-api#oauth2-token-sent-in-a-header
const AuthorizationPrefix = "token " //trailing space required

// A GithubSigner signs request with an basic header prefix
type GithubSigner struct {
	ConsumerKey    string
	ConsumerSecret string
}

func (b GithubSigner) AuthSigningParams() map[string]string {
	return map[string]string{
		"Authorization": b.authorizationHeaderValue(),
		"Content-Type":  "application/x-www-form-urlencoded",
	}
}

func (b GithubSigner) OAuthParams(token string) map[string]string {
	header := []string{"token ", token}

	return map[string]string{
		"Authorization": strings.Join(header, ""),
		"Content-Type":  "application/json",
	}
}

func (b GithubSigner) authorizationHeaderValue() string {
	return AuthorizationPrefix + oauth2.Base64Enc(fmt.Sprintf("%s:%s", b.ConsumerKey, b.ConsumerSecret))
}
