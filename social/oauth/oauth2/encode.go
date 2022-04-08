/*
encode.go
Created at 08.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package oauth2

import "encoding/base64"

func base64Enc(cs string) string {
	return base64.StdEncoding.EncodeToString([]byte(cs))
}
