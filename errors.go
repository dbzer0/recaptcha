package recaptcha

import (
	"errors"
)

// https://developers.google.com/recaptcha/docs/verify#error_code_reference

var (
	// missing-input-secret	The secret parameter is missing.
	ErrMissingInputSecret = errors.New("the secret parameter is missing")
	// invalid-input-secret	The secret parameter is invalid or malformed.
	ErrMissingInputResponse = errors.New("the response parameter is missing")
	// invalid-input-response	The response parameter is invalid or malformed.
	ErrInvalidInputResponse = errors.New("the response parameter is invalid or malformed")
	// bad-request	The request is invalid or malformed
	ErrBadRequest = errors.New("the request is invalid or malformed")
	// timeout-or-duplicate	The response is no longer valid: either is too old or has been used previously.
	ErrTimeoutOrDuplicate = errors.New("the response is no longer valid: either is too old or has been used previously")
	
	ErrUnknown = errors.New("unknown error")
)

