/*
search.go
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
	SearchPath = "/youtube/v3/search"
)

// SearchService provides methods for searching information about a video, channel or playlist that matches the search params
type SearchService struct {
	oauth2 *oauth2.OAuth2
}

// newChannelService returns a new YouTube ChannelService.
func newSearchService(oauth2 *oauth2.OAuth2) *SearchService {
	return &SearchService{
		oauth2: oauth2,
	}
}

// Search returns the search result of the defined params
// Required scopes: https://www.googleapis.com/auth/youtube.readonly
func (c *SearchService) Search(params *SearchParams) (*SearchResp, error) {
	cl := new(SearchResp)
	apiError := new(APIError)

	err := c.oauth2.Get(SearchPath, cl, apiError, params)
	return cl, social.CheckError(err)
}

// SearchParams are the params for Searching the YouTube API.
type SearchParams struct {
	// ID. Multiple user IDs can be specified. Limit: 100. Optional
	Part string `url:"part"` // The part parameter specifies a comma-separated list of one or more search resource properties that the API response will include. Set the parameter value to snippet.
	// This parameter can only be used in a properly authorized request, and it is intended exclusively for YouTube content partners.
	// The forContentOwner parameter restricts the search to only retrieve videos owned by the content owner identified by the onBehalfOfContentOwner parameter. If forContentOwner is set to true, the request must also meet these requirements:
	ForContentOwner string `url:"forContentOwner,omitempty"`
	// This parameter can only be used in a properly authorized request. The forDeveloper parameter restricts the search to only retrieve videos uploaded via the developer's application or website. The API server uses the request's authorization credentials to identify the developer. The forDeveloper parameter can be used in conjunction with optional search parameters like the q parameter.
	ForDeveloper string `url:"forDeveloper,omitempty"`
	// This parameter can only be used in a properly authorized request. The forMine parameter restricts the search to only retrieve videos owned by the authenticated user.
	// If you set this parameter to true, then the type parameter's value must also be set to video. In addition, none of the following other parameters can be set in the same request: videoDefinition, videoDimension, videoDuration, videoLicense, videoEmbeddable, videoSyndicated, videoType.
	ForMine bool `url:"forMine,omitempty"`
	// The relatedToVideoId parameter retrieves a list of videos that are related to the video that the parameter value identifies. The parameter value must be set to a YouTube video ID and, if you are using this parameter, the type parameter must be set to video.
	// Note: that if the relatedToVideoId parameter is set, the only other supported parameters are part, maxResults, pageToken, regionCode, relevanceLanguage, safeSearch, type (which must be set to video), and fields.
	RelatedToVideoId string `url:"relatedToVideoId,omitempty"`

	//Optional params

	//The maxResults parameter specifies the maximum number of items that should be returned in the result set. Acceptable values are 0 to 50, inclusive. The default value is 5.
	MaxResults uint `url:"maxResults,omitempty"`
	// The order parameter specifies the method that will be used to order resources in the API response. The default value is relevance.
	// Acceptable values are: date, rating, relevance, title, videoCount, viewCount
	Order string `url:"order,omitempty"`
	// The pageToken parameter identifies a specific page in the result set that should be returned. In an API response, the nextPageToken and prevPageToken properties identify other pages that could be retrieved.
	PageToken string `url:"pageToken,omitempty"`
	// The type parameter restricts a search query to only retrieve a particular type of resource. The value is a comma-separated list of resource types. The default value is video,channel,playlist.
	// Acceptable values are: channel, playlist, video
	Type string `url:"type,omitempty"`
}

type SearchResp struct {
	Kind          string `json:"kind"`
	Etag          string `json:"etag"`
	RegionCode    string `json:"regionCode,omitempty"`
	NextPageToken string `json:"nextPageToken,omitempty"`
	PrevPageToken string `json:"prevPageToken,omitempty"`
	PageInfo      struct {
		TotalResults   int `json:"totalResults"`
		ResultsPerPage int `json:"resultsPerPage"`
	} `json:"pageInfo"`
	Items []struct {
		Kind string `json:"kind"`
		Etag string `json:"etag"`
		Id   struct {
			Kind      string `json:"kind"`
			ChannelId string `json:"channelId"`
		} `json:"id"`
		Snippet struct {
			PublishedAt time.Time `json:"publishedAt"`
			ChannelId   string    `json:"channelId"`
			Title       string    `json:"title"`
			Description string    `json:"description"`
			Thumbnails  struct {
				Default struct {
					Url string `json:"url"`
				} `json:"default"`
				Medium struct {
					Url string `json:"url"`
				} `json:"medium"`
				High struct {
					Url string `json:"url"`
				} `json:"high"`
			} `json:"thumbnails"`
			ChannelTitle         string    `json:"channelTitle"`
			LiveBroadcastContent string    `json:"liveBroadcastContent"`
			PublishTime          time.Time `json:"publishTime"`
		} `json:"snippet"`
	} `json:"items"`
}
