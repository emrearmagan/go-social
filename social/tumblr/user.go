/*
user.go
Created at 10.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package tumblr

import (
	"github.com/emrearmagan/go-social/oauth/oauth1"
	"github.com/emrearmagan/go-social/social"
)

const (
	UserPath = "/v2/user/info"
)

// UserService provides methods for user credentials
type UserService struct {
	oauth1 *oauth1.OAuth1
}

// newUserService returns a new Tumblr UserService.
func newUserService(oauth1 *oauth1.OAuth1) *UserService {
	return &UserService{
		oauth1: oauth1,
	}
}

// UserCredentials returns the authorized user if credentials are valid and returns an error otherwise.
// https://www.tumblr.com/docs/en/api/v2#userinfo--get-a-users-information
func (u *UserService) UserCredentials() (*User, error) {
	user := new(User)
	apiError := new(APIError)

	err := u.oauth1.Get(UserPath, user, apiError, nil)
	return user, social.CheckError(err)
}

type User struct {
	Meta struct {
		Status int    `json:"status"`
		Msg    string `json:"msg"`
	} `json:"meta"`
	Response struct {
		User struct {
			Name              string `json:"name"`
			Likes             int    `json:"likes"`
			Following         int    `json:"following"`
			DefaultPostFormat string `json:"default_post_format"`
			Blogs             []struct {
				Admin          bool   `json:"admin"`
				Ask            bool   `json:"ask"`
				AskAnon        bool   `json:"ask_anon"`
				AskPageTitle   string `json:"ask_page_title"`
				AsksAllowMedia bool   `json:"asks_allow_media"`
				Avatar         []struct {
					Width  int    `json:"width"`
					Height int    `json:"height"`
					URL    string `json:"url"`
				} `json:"avatar"`
				CanChat                  bool   `json:"can_chat"`
				CanSendFanMail           bool   `json:"can_send_fan_mail"`
				CanSubscribe             bool   `json:"can_subscribe"`
				Description              string `json:"description"`
				Drafts                   int    `json:"drafts"`
				Facebook                 string `json:"facebook"`
				FacebookOpengraphEnabled string `json:"facebook_opengraph_enabled"`
				Followed                 bool   `json:"followed"`
				Followers                int    `json:"followers"`
				IsBlockedFromPrimary     bool   `json:"is_blocked_from_primary"`
				IsNsfw                   bool   `json:"is_nsfw"`
				Likes                    int    `json:"likes"`
				Messages                 int    `json:"messages"`
				Name                     string `json:"name"`
				Posts                    int    `json:"posts"`
				Primary                  bool   `json:"primary"`
				Queue                    int    `json:"queue"`
				ShareLikes               bool   `json:"share_likes"`
				Subscribed               bool   `json:"subscribed"`
				Theme                    struct {
					AvatarShape        string `json:"avatar_shape"`
					BackgroundColor    string `json:"background_color"`
					BodyFont           string `json:"body_font"`
					HeaderBounds       string `json:"header_bounds"`
					HeaderImage        string `json:"header_image"`
					HeaderImageFocused string `json:"header_image_focused"`
					HeaderImagePoster  string `json:"header_image_poster"`
					HeaderImageScaled  string `json:"header_image_scaled"`
					HeaderStretch      bool   `json:"header_stretch"`
					LinkColor          string `json:"link_color"`
					ShowAvatar         bool   `json:"show_avatar"`
					ShowDescription    bool   `json:"show_description"`
					ShowHeaderImage    bool   `json:"show_header_image"`
					ShowTitle          bool   `json:"show_title"`
					TitleColor         string `json:"title_color"`
					TitleFont          string `json:"title_font"`
					TitleFontWeight    string `json:"title_font_weight"`
				} `json:"theme"`
				Title          string `json:"title"`
				TotalPosts     int    `json:"total_posts"`
				Tweet          string `json:"tweet"`
				TwitterEnabled bool   `json:"twitter_enabled"`
				TwitterSend    bool   `json:"twitter_send"`
				Type           string `json:"type"`
				Updated        int    `json:"updated"`
				URL            string `json:"url"`
				UUID           string `json:"uuid"`
			} `json:"blogs"`
		} `json:"user"`
	} `json:"response"`
}
