/*
decoder.go
Created at 21.06.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package reddit

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Reddit only returns an HTML page as a response for error. So decoding the error with the default JSONDecoder will always fail and result in being unable to check for the error
// Instead we set our own decoder where we only decode the JSON response if the statuscode is between 200 and 299
type redditDecoder struct{}

// Decode decodes the Response Body if statuscode valid into the value pointed to by v.
func (r redditDecoder) Decode(resp *http.Response, v interface{}) error {
	if code := resp.StatusCode; 200 <= code && code <= 299 {
		return json.NewDecoder(resp.Body).Decode(v)
	}

	if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
		// Else decode some json so that the ErrorDetail != nil
		e := fmt.Sprintf(`{"Message": "%s", "error": %d}`, "Failed request", resp.StatusCode)
		return json.Unmarshal([]byte(e), v)
	}

	return nil
}
