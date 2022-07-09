/*
channel.go
Created at 09.07.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package youtube

import (
	"github.com/emrearmagan/go-social/oauth/oauth2"
	"github.com/emrearmagan/go-social/social"
	"time"
)

const (
	ChannelPath = "/youtube/v3/channels"
)

// ChannelService provides methods for channel information
type ChannelService struct {
	oauth2 *oauth2.OAuth2
}

// newChannelService returns a new YouTube ChannelService.
func newChannelService(oauth2 *oauth2.OAuth2) *ChannelService {
	return &ChannelService{
		oauth2: oauth2,
	}
}

// Channel returns the channel of the authorized user if credentials are valid and returns an error otherwise.
// Required scopes: https://www.googleapis.com/auth/youtube.readonly
func (c *ChannelService) Channel(params *ChannelPartParams) (*ChannelResp, error) {
	cl := new(ChannelResp)
	apiError := new(APIError)

	err := c.oauth2.Get(ChannelPath, cl, apiError, params)
	return cl, social.CheckError(err)
}

// ChannelPartParams are the params for Channel information.
type ChannelPartParams struct {
	// The part parameter specifies a comma-separated list of one or more channel resource properties that the API response will include.
	Part string `url:"part,omitempty"` // The following list contains the part names that you can include in the parameter value: auditDetails, brandingSettings, contentDetails, contentOwnerDetails, id, localizations, snippet, statistics, status, topicDetails
	// The forUsername parameter specifies a YouTube username, thereby requesting the channel associated with that username.
	ForUsername string `url:"forUsername,omitempty"`
	// The id parameter specifies a comma-separated list of the YouTube channel ID(s) for the resource(s) that are being retrieved. In a channel resource, the id property specifies the channel's YouTube channel ID.
	Id string `url:"id,omitempty"`
	// This parameter can only be used in a properly authorized request. Note: This parameter is intended exclusively for YouTube content partners.
	//Set this parameter's value to true to instruct the API to only return channels managed by the content owner that the onBehalfOfContentOwner parameter specifies. The user must be authenticated as a CMS account linked to the specified content owner and onBehalfOfContentOwner must be provided.
	ManagedByMe bool `url:"managedByMe,omitempty"`
	//This parameter can only be used in a properly authorized request. Set this parameter's value to true to instruct the API to only return channels owned by the authenticated user.
	Mine bool `url:"mine,omitempty"`

	//Optional params
	//The maxResults parameter specifies the maximum number of items that should be returned in the result set. Acceptable values are 0 to 50, inclusive. The default value is 5.
	MaxResults uint `url:"maxResults,omitempty"`
	//The pageToken parameter identifies a specific page in the result set that should be returned. In an API response, the nextPageToken and prevPageToken properties identify other pages that could be retrieved.
	PageToken uint `url:"pageToken,omitempty"`
}

type ChannelResp struct {
	Kind          string `json:"kind"`
	Etag          string `json:"etag"`
	NextPageToken string `json:"nextPageToken,omitempty"`
	PrevPageToken string `json:"prevPageToken,omitempty"`
	PageInfo      struct {
		TotalResults   int `json:"totalResults"`
		ResultsPerPage int `json:"resultsPerPage"`
	} `json:"pageInfo"`
	Items []struct {
		Kind    string `json:"kind"`
		Etag    string `json:"etag"`
		Id      string `json:"id"`
		Snippet struct {
			Title       string    `json:"title"`
			Description string    `json:"description"`
			PublishedAt time.Time `json:"publishedAt"`
			Thumbnails  struct {
				Default struct {
					Url    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"default"`
				Medium struct {
					Url    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"medium"`
				High struct {
					Url    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"high"`
			} `json:"thumbnails"`
			Localized struct {
				Title       string `json:"title"`
				Description string `json:"description"`
			} `json:"localized"`
		} `json:"snippet"`
		ContentDetails struct {
			RelatedPlaylists struct {
				Likes   string `json:"likes"`
				Uploads string `json:"uploads"`
			} `json:"relatedPlaylists"`
		} `json:"contentDetails"`
		Statistics struct {
			ViewCount             string `json:"viewCount"`
			SubscriberCount       string `json:"subscriberCount"`
			HiddenSubscriberCount bool   `json:"hiddenSubscriberCount"`
			VideoCount            string `json:"videoCount"`
		} `json:"statistics"`
		Status struct {
			PrivacyStatus     string `json:"privacyStatus"`
			IsLinked          bool   `json:"isLinked"`
			LongUploadsStatus string `json:"longUploadsStatus"`
		} `json:"status"`
		BrandingSettings struct {
			Channel struct {
				Title string `json:"title"`
			} `json:"channel"`
		} `json:"brandingSettings"`
		ContentOwnerDetails struct {
		} `json:"contentOwnerDetails"`
	} `json:"items"`
}
