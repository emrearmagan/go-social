/*
api-error.go
Created at 08.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package social

type ApiErrors interface {
	Error() string
	ErrorDetail() interface{}
	Empty() bool
	Status() int
	ReturnErrorResponse() error
	SetStatus(code int)
}

// RelevantError returns any non-nil http-related error if any. If the decoded apiError is non-zero
// the apiError is returned. Otherwise, no errors occurred, returns nil.
func RelevantError(httpError error, apiError ApiErrors) error {
	if httpError != nil {
		return httpError
	}

	if apiError.Empty() {
		return nil
	}

	return apiError
}

func CheckError(err error) error {
	switch e := err.(type) {
	case ApiErrors:
		return e.ReturnErrorResponse()
	case error:
		return err
	}
	return nil
}
