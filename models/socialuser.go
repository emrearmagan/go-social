/*
socialuser.go
Created at 08.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package models

type SocialUser struct {
	Username     string `json:"username"`
	Name         string `json:"name"`
	UserId       string `json:"user_id"`
	Verified     bool   `json:"verified"` // Flag to indicate if a user is verified or uses or pro version
	ContentCount int64  `json:"contentCount"`
	AvatarUrl    string `json:"avatar_url"`
	Followers    int    `json:"followers"`
	Following    *int   `json:"following"`
	Url          string `json:"url"`
}
