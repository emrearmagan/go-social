/*
oauth2.go
Created at 08.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package oauth2

import (
	"context"
	"github.com/emrearmagan/go-social/models/errors"
	"github.com/emrearmagan/go-social/oauth"
	"github.com/emrearmagan/go-social/social"
	"github.com/emrearmagan/go-social/social/client"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type OAuth2 struct {
	ctx         context.Context
	credentials *oauth.Credentials
	client      *client.HttpClient
	token       *Token
	signer      Signer
}

func NewOAuth(ctx context.Context, c *oauth.Credentials, token *Token, cl *client.HttpClient) *OAuth2 {
	return &OAuth2{
		ctx:         ctx,
		credentials: c,
		token:       token,
		client:      cl,
		signer: BearerSigner{
			ConsumerKey:    c.ConsumerKey,
			ConsumerSecret: c.ConsumerSecret,
		},
	}
}

// New return a copy of OAuth2 object
func (a *OAuth2) New() *OAuth2 {
	return &OAuth2{
		ctx:         a.ctx,
		credentials: a.credentials,
		token:       a.token,
		client:      a.client,
		signer:      a.signer,
	}
}

// NewClient return a new OAuth2 with a new given client
func (a *OAuth2) NewClient(client *client.HttpClient) *OAuth2 {
	return &OAuth2{
		ctx:         a.ctx,
		credentials: a.credentials,
		token:       a.token,
		client:      client,
		signer:      a.signer,
	}
}

// Basic return a new OAuth2 with a basic authentication
func (a *OAuth2) Basic() *OAuth2 {
	return &OAuth2{
		ctx:         a.ctx,
		credentials: a.credentials,
		token:       a.token,
		client:      a.client,
		signer:      BasicSigner{ConsumerKey: a.credentials.ConsumerKey, ConsumerSecret: a.credentials.ConsumerSecret},
	}
}

// Signer sets a Signer for signing the oauth requests
func (a *OAuth2) Signer(s Signer) *OAuth2 {
	return &OAuth2{
		ctx:         a.ctx,
		credentials: a.credentials,
		token:       a.token,
		client:      a.client,
		signer:      s,
	}
}

func (a *OAuth2) Get(path string, resp interface{}, apiError social.ApiErrors, params interface{}) error {
	cl := a.client.AddQuery(params).Get(path)
	req, err := cl.Request()
	if err != nil {
		return err
	}

	req = req.WithContext(a.ctx)
	for k, v := range a.signer.OAuthParams(a.token.Token) {
		req.Header.Set(k, v)
	}

	httpResp, err := cl.Do(req, resp, apiError.ErrorDetail())
	if httpResp != nil {
		if code := httpResp.StatusCode; code >= 300 {
			apiError.SetStatus(httpResp.StatusCode)
		}
	}

	return social.RelevantError(err, apiError)
}

func (a *OAuth2) RefreshToken(refreshBase string, path string, resp interface{}, apiError social.ApiErrors) error {
	// Using a new http cl, so we don't mess up the base path for other requests.
	// Since some APIs use a different base path for refreshing tokens, like reddit
	cl := a.client.New().Base(refreshBase).Post(path)
	req, err := cl.Request()
	if err != nil {
		return err
	}
	req = req.WithContext(a.ctx)
	req, err = a.signRequest(req)
	if err != nil {
		return err
	}

	httpResp, err := cl.Do(req, resp, apiError.ErrorDetail())
	if httpResp != nil {
		if code := httpResp.StatusCode; code >= 300 {
			apiError.SetStatus(httpResp.StatusCode)
		}
	}

	return social.RelevantError(err, apiError)
}

func (a *OAuth2) RevokeToken(revokeBase string, path string, resp interface{}, apiError social.ApiErrors) error {
	cl := a.client.New().Base(revokeBase).Post(path)
	req, err := cl.Request()
	if err != nil {
		return err
	}

	req = req.WithContext(a.ctx)

	httpResp, err := cl.Do(req, resp, apiError.ErrorDetail())
	if httpResp != nil {
		if code := httpResp.StatusCode; code >= 300 {
			apiError.SetStatus(httpResp.StatusCode)
		}
	}
	return social.RelevantError(err, apiError)
}

func (a *OAuth2) signRequest(req *http.Request) (*http.Request, error) {
	if a.credentials.ConsumerKey == "" {
		return nil, errors.New(errors.ErrBadAuthenticationData, "OAuth2: provide valid credentials")
	}

	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", a.token.RefreshToken)

	req.Body = ioutil.NopCloser(strings.NewReader(data.Encode()))

	for k, v := range a.signer.AuthSigningParams() {
		req.Header.Set(k, v)
	}

	req.Header.Add(ContentLengthHeaderName, strconv.Itoa(len(data.Encode())))

	return req, nil
}

func (a *OAuth2) Client() *client.HttpClient {
	return a.client
}

func (a *OAuth2) Token() *Token {
	return a.token
}

func (a *OAuth2) UpdateToken(token *Token) {
	a.token = token
}

func (a *OAuth2) Credentials() oauth.Credentials {
	return *a.credentials
}
