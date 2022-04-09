## go-social (Development)

![Commit](https://img.shields.io/github/last-commit/emrearmagan/go-social)
![MIT](https://img.shields.io/github/license/mashape/apistatus.svg)
----
<br>
go-social is a Go client library for the various social media APIs.

### Supported APIs
<p>

[![Twitter](https://img.shields.io/badge/-Twitter-FFFFFF?style=flat&logo=twitter)](https://developer.twitter.com/en/docs/twitter-api)
[![Dribbble](https://img.shields.io/badge/-Dribbble-FFFFFF?style=flat&logo=dribbble)](https://developer.dribbble.com/v2/)
[![Github](https://img.shields.io/badge/-Github-FFFFFF?style=flat&logo=github&logoColor=black)](https://docs.github.com/en/rest)
[![Reddit](https://img.shields.io/badge/-Reddit-FFFFFF?style=flat&logo=reddit)](https://www.reddit.com/dev/api)
[![Spotify](https://img.shields.io/badge/-Spotify-FFFFFF?style=flat&logo=spotify)](https://developer.spotify.com)
<!--[![Tumblr](https://img.shields.io/badge/-Tumblr-FFFFFF?style=flat&logo=tumblr&logoColor=black)]()-->
<!--[![Facebook](https://img.shields.io/badge/-Facebook-FFFFFF?style=flat&logo=facebook)]()-->
<!--[![Instagram](https://img.shields.io/badge/-Instagram-FFFFFF?style=flat&logo=instagram)]()-->

</p>

### Endpoints
- Twitter API 
  - User Credentials
  - Follower IDs
  - Following IDs
- Dribbble API:
  - User Credentials
  - User Shots
- Github API:
    - User Credentials
    - User Shots
    - Follower IDs
    - Following IDs
- Reddit API
  - User Credentials
  - Refresh token
- Spotify API
  - User Credentials
  - Followed Artists
  - Refresh token
  - User Playlist


## Usage

---
### Account configuration
Initialize the required configuration for each account.
```go
cred := oauth.NewCredentials("CONSUMER_KEY", "CONSUMER_SECRET")
token := oauth2.NewToken("ACCESS_TOKEN", "REFRESH_TOKEN")
config := oauth2.NewOAuth(context.TODO(), cred, token)
```
You can also provide a config file to load your credentials and token. See [Config](./config/config_example.json) for an example.
```go
// pass config file
flag.StringVar(&ConfigPath, "c", "./config/config_example.json", "Specified the config file for running server. Default is the \"config_example\" in the config directory.")
flag.Parse() 

//load config
accounts, err := config.LoadConfig(ConfigPath)
if err != nil {
    log.Fatal(err.Error())
}
    
conf := oauth2.NewOAuth(context.TODO(), &accounts.Spotify.Credentials, &accounts.Spotify.Token)
```
### Access API
Afterwards each social media package provides a Client with a corresponding service for accessing the API. 
```go
spotify := spotify.NewClient(conf)

//Use UserService for User related API calls
u, err := spotify.User.UserCredentials()
fmt.Printf("User credentials: %v \n\n", u)

//Use PlaylistService for Playlist related API calls
p, _ := spotify.Playlist.UserPlaylists(&spotify.UserPlaylistParams{
        Limit: 10,
})
fmt.Printf("User Playlist: %v \n\n", p)
```


### Go-Social User Response
Each Package also provides a method for generalized credentials response which provides basic information about the user:
```go
type SocialUser struct {
    Username     string `json:"username"`
    Name         string `json:"name"`
    UserId       string `json:"user_id"`
    Verified     bool   `json:"verified"` // Flag to indicate if a user is verified or uses or pro version
    ContentCount int64  `json:"contentCount"`
    AvatarUrl    string `json:"avatar_url"`
    Followers    int    `json:"followers"`
    Following    *int   `json:"following"` // Can be nil, since some APIs do not provide/have this
    Url          string `json:"url"`
}
```
Simply call client.GoSocialUser()
```go
// Returns a SocialUser struct
s, _ := spotify.GoSocialUser()
fmt.Printf("go-social user: %v \n", s)
```

### Error Handling
Each API Error code is mapped to models.Error structs which will provide additional information.

```go
_, err = spotify.User.UserCredentials()
    if err != nil {
        // Logging the error
        if e, ok := err.(errors.SocialError); ok {
            switch e.Errors {
            case errors.ErrBadRequest:
                // Something bad happened
            case errors.ErrUnauthorized, errors.ErrBadAuthenticationData, errors.ErrInvalidOrExpiredToken:
                // Request was unauthorized
            case errors.ErrRateLimit:
                // Rate limit exceeded. Try later again.
            case errors.ErrApiError:
                // Some API Error happened. See err.Error() for further information
                fmt.Println(err.Error())
            case errors.ErrNotModified:
                // The requested resource has not been modified since the previous transmission
            default:
                // Some other error
            }
        }
        return
    }
```
The previous snippet would print the following if an unauthorized/invalid spotify user attempted to retrieve user credentials:

```
Spotify: 401 - The access token expired
```

### Refreshing a Token
Most access tokens typically have a limited lifetime such as the `Spotify` API. Once they expire clients can use the refresh token to `refresh` the access token.
Calling the refresh Method automatically sets the new access token so that further API calls become valid again.
```go
if _, err := spotify.User.UserCredentials(); err != nil {
    //Access token expired
    if err == models.ErrUnauthorized {
        newToken, _ := spotify.Account.RefreshToken()
        fmt.Printf("Refreshed Token: %v \n\n", newToken)
    }
}

// Access token updated, do request with the updated token
user,  := spotify.User.UserCredentials()
```
### Custom API calls

The go-social library comes with some standard api calls and structures for like User Credentials etc., but you are not required to use them.
Create your own OAuth object and make request.

```go
// Struct for the api response
type CustomStruct struct {
    CustomID        string `json:"id"`
    CustomUsername  string `json:"username"`
}

type ErrorStruct struct {
    Code    int     `json:"code"`
    Message string `json:"message"`
}

cred := oauth.NewCredentials("CONSUMER_KEY", "CONSUMER_SECRET")
token := oauth1.NewToken("TOKEN", "TOKEN_SECRET")
auth := oauth1.NewOAuth(context.TODO(), cred, token)

// Either use the build in http client or create your own
httpClient := social.NewClient().Base("https://api.somesite.com/v2/")
// Set the http client for further requests
auth = auth.NewClient(httpClient)
// Initialize your own response and error struct
resp := new(CustomStruct)
error := new(ErrorStruct)

// Make the request. Request will be automatically signed using the default HMAC Signer.
err := auth.Get("/me/user", resp, error, nil)
if err != nil {
    log.Fatal(err.Error())
}
// do something with `resp`
```
## Installation
Run

    go get github.com/emrearmagan/go-social

Include in your source:

    import "github.com/kurrik/twittergo"