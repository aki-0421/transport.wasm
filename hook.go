package transport

import (
	"net/http"
)

// BeforeRequestFunc executes immediately before sending HTTP Request
// You can pass the options to the Fetch API here
type BeforeRequestFunc func(*Transport)

// ModifyRequestFunc rewrites http.Request
type ModifyRequestFunc func(*http.Request) error

// ModifyResponseFunc rewrites http.Response
type ModifyResponseFunc func(*http.Response) error
