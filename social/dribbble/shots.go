/*
shots.go
Created at 08.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package dribbble

import (
	"github.com/emrearmagan/go-social/oauth/oauth2"
	"github.com/emrearmagan/go-social/social"
	"time"
)

const (
	ShotsPath = "/v2/user/shots"
)

// ShotService provides a method for account credential verification.
type ShotService struct {
	oauth2 *oauth2.OAuth2
}

// newUserService returns a new AccountService.
func newShotService(oauth2 *oauth2.OAuth2) *ShotService {
	return &ShotService{
		oauth2: oauth2,
	}
}

// DribbbleShots returns all shots for the authenticated user.
// See: https://developer.dribbble.com/v2/shots/#list-shots for more information
func (s *ShotService) DribbbleShots() (*Shots, error) {
	shots := new(Shots)
	apiError := new(APIError)

	err := s.oauth2.Get(ShotsPath, shots, apiError, nil)

	return shots, social.CheckError(err)
}

// Shots represents the Shots of the Dribbble User
// https://developer.dribbble.com/v2/shots/
type Shots []struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	Images      struct {
		Hidpi  interface{} `json:"hidpi"`
		Normal string      `json:"normal"`
		Teaser string      `json:"teaser"`
	} `json:"images"`
	PublishedAt time.Time `json:"published_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	HTMLURL     string    `json:"html_url"`
	Animated    bool      `json:"animated"`
	Tags        []string  `json:"tags"`
	Attachments []struct {
		ID           int       `json:"id"`
		URL          string    `json:"url"`
		ThumbnailURL string    `json:"thumbnail_url"`
		Size         int       `json:"size"`
		ContentType  string    `json:"content_type"`
		CreatedAt    time.Time `json:"created_at"`
	} `json:"attachments"`
	Projects []struct {
		ID          int       `json:"id"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		ShotsCount  int       `json:"shots_count"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	} `json:"projects"`
	Team struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Login     string `json:"login"`
		HTMLURL   string `json:"html_url"`
		AvatarURL string `json:"avatar_url"`
		Bio       string `json:"bio"`
		Location  string `json:"location"`
		Links     struct {
			Web     string `json:"web"`
			Twitter string `json:"twitter"`
		} `json:"links"`
		Type      string    `json:"type"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"team"`
	Video struct {
		ID               int       `json:"id"`
		Duration         int       `json:"duration"`
		VideoFileName    string    `json:"video_file_name"`
		VideoFileSize    int       `json:"video_file_size"`
		Width            int       `json:"width"`
		Height           int       `json:"height"`
		Silent           bool      `json:"silent"`
		CreatedAt        time.Time `json:"created_at"`
		UpdatedAt        time.Time `json:"updated_at"`
		URL              string    `json:"url"`
		SmallPreviewURL  string    `json:"small_preview_url"`
		LargePreviewURL  string    `json:"large_preview_url"`
		XlargePreviewURL string    `json:"xlarge_preview_url"`
	} `json:"video"`
	LowProfile bool `json:"low_profile"`
}
