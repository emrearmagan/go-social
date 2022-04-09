/*
oauth1.go
Created at 08.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package oauth1

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/emrearmagan/go-social/oauth"
	"github.com/emrearmagan/go-social/social"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	// AuthorizationHeaderName constant is the header key name for the
	// oauth1.0a based authorization header
	AuthorizationHeaderName   = "Authorization"
	nonceLength               = 16
	alphaNumericChars         = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	authorizationPrefix       = "OAuth " // trailing space is required
	oauthConsumerKeyParam     = "oauth_consumer_key"
	oauthNonceParam           = "oauth_nonce"
	oauthSignatureParam       = "oauth_signature"
	oauthTokenParam           = "oauth_token"
	oauthSignatureMethodParam = "oauth_signature_method"
	oauthTimestampParam       = "oauth_timestamp"
	oauthVersionParam         = "oauth_version"
	defaultOauthVersion       = "1.0"
)

type OAuth1 struct {
	ctx         context.Context
	credentials *oauth.Credentials
	token       *Token
	client      *social.HttpClient
	// OAuth1 signer (defaults is HMAC-SHA1)
	signer Signer
}

func NewOAuth(ctx context.Context, c *oauth.Credentials, token *Token) *OAuth1 {
	return &OAuth1{
		ctx:         ctx,
		credentials: c,
		token:       token,
		client:      social.NewClient(),
	}
}

func (a *OAuth1) NewClient(client *social.HttpClient) *OAuth1 {
	return &OAuth1{
		ctx:         a.ctx,
		credentials: a.credentials,
		token:       a.token,
		client:      client,
	}
}

func (a *OAuth1) Get(path string, resp interface{}, apiError social.Errors, params interface{}) error {
	client := a.client.AddQuery(params).Get(path)

	req, err := a.client.Request()
	if err != nil {
		return err
	}
	req = req.WithContext(a.ctx)

	if err := a.SignRequest(req); err != nil {
		return err
	}
	httpResp, err := client.Do(req, resp, apiError.ErrorDetail())
	if httpResp != nil {
		if code := httpResp.StatusCode; code >= 300 {
			apiError.SetStatus(httpResp.StatusCode)
		}
	}
	return social.RelevantError(err, apiError)
}

// oauthParams returns the OAuth request parameters for the given credentials,
// method, URL and application params. See
// http://tools.ietf.org/html/rfc5849#section-3.4 for more information about
// signatures.
func (a *OAuth1) oAuthParams(req *http.Request) map[string]string {
	params := map[string]string{
		oauthConsumerKeyParam:     a.credentials.ConsumerKey,
		oauthTokenParam:           a.token.Token,
		oauthNonceParam:           getNonce(),
		oauthSignatureMethodParam: a.Signer().Name(),
		oauthTimestampParam:       strconv.FormatInt(time.Now().Unix(), 10), //"1318622958",
		oauthVersionParam:         defaultOauthVersion,
	}

	/// add request query
	for key, value := range req.URL.Query() {
		params[key] = value[0]
	}

	return params
}

func (a *OAuth1) Client() *social.HttpClient {
	return a.client
}

func (a *OAuth1) SignRequest(req *http.Request) error {
	if a.credentials.ConsumerKey == "" || a.credentials.ConsumerSecret == "" {
		return errors.New("OAuth1: provide valid credentials")
	}
	if a.token.Token == "" || a.token.TokenSecret == "" {
		return errors.New("OAuth1: provide valid token")
	}

	oauthParams := a.oAuthParams(req)

	//Signature Base
	signatureBase := signatureBase(req, oauthParams)

	//Sign
	signature, err := a.Signer().Sign(a.token.TokenSecret, signatureBase)
	if err != nil {
		return errors.New("OAuth1: failed to sign message")
	}

	oauthParams[oauthSignatureParam] = signature
	req.Header.Set(AuthorizationHeaderName, authHeaderValue(oauthParams))

	return nil
}

// authHeaderValue formats OAuth parameters according to RFC 5849 3.5.1. OAuth
// params are percent encoded, sorted by key (for testability), and joined by
// "=" into pairs. Pairs are joined with a ", " comma separator into a header
// string.
// The given OAuth params should include the "oauth_signature" key.
func authHeaderValue(oauthParams map[string]string) string {
	pairs := sortParameters(encodeParameters(oauthParams), `%s="%s"`)
	return authorizationPrefix + strings.Join(pairs, ",")
}

func (a *OAuth1) Signer() Signer {
	if a.signer != nil {
		return a.signer
	}

	return &HMACSigner{ConsumerSecret: a.credentials.ConsumerSecret}
}

func signatureBase(req *http.Request, params map[string]string) string {
	method := strings.ToUpper(req.Method)
	baseURL := baseURI(req)
	paramString := strings.Join(sortParameters(encodeParameters(params), "%s=%s"), "&")

	//siganture base string consutrcted accoirding to 3.4.1
	baseParts := []string{method, PercentEncode(baseURL), PercentEncode(paramString)}
	return strings.Join(baseParts, "&")
}

// baseURI returns the base string URI of a request according to RFC 5849
// 3.4.1.2. The scheme and host are lowercased, the port is dropped if it
// is 80 or 443, and the path minus query parameters is included.
func baseURI(req *http.Request) string {
	scheme := strings.ToLower(req.URL.Scheme)
	host := strings.ToLower(req.URL.Host)
	if hostPort := strings.Split(host, ":"); len(hostPort) == 2 && (hostPort[1] == "80" || hostPort[1] == "443") {
		host = hostPort[0]
	}

	path := req.URL.Path
	if path != "" {
		path = req.URL.EscapedPath()
	}

	//query := req.URL.RawQuery
	//println(fmt.Sprintf("%v://%v%v?%v", scheme, host, path, query))
	//return fmt.Sprintf("%v://%v%v?%v", scheme, host, path, query)
	return fmt.Sprintf("%v://%v%v", scheme, host, path)
}

// encodeParameters percent encodes parameter keys and values according to
// RFC5849 3.6 and RFC3986 2.1 and returns a new map.
func encodeParameters(params map[string]string) map[string]string {
	encoded := map[string]string{}
	for key, value := range params {
		encoded[PercentEncode(key)] = PercentEncode(value)
	}
	return encoded
}

// sortParameters sorts parameters by key and returns a slice of key/value
// pairs formatted with the given format string (e.g. "%s=%s").
func sortParameters(params map[string]string, format string) []string {
	// sort by key
	keys := make([]string, len(params))
	i := 0
	for key := range params {
		keys[i] = key
		i++
	}
	sort.Strings(keys)
	// parameter join
	pairs := make([]string, len(params))
	for i, key := range keys {
		pairs[i] = fmt.Sprintf(format, key, params[key])
	}
	return pairs
}

// The getNonce generates a random string for replay protection as per
// https://tools.ietf.org/html/rfc5849#section-3.3
func getNonce() (nonce string) {
	randomVal := make([]byte, nonceLength)
	_, _ = rand.Read(randomVal)

	var length = len(alphaNumericChars)
	for _, v := range randomVal {
		nonce += string(alphaNumericChars[int(v)%length])
	}
	return
}
