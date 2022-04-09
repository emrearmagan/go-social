/*
signer.go
Created at 08.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package oauth2

import "net/http"

type Signer interface {
	SignRequest(req *http.Request) error
}
