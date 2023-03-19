package transport

import (
	"net/http"
)

type Transport struct {
	http.RoundTripper

	// Mode present an option to the Fetch API
	Mode FetchMode
	// Redirect present an option to the Fetch API
	Redirect FetchRedirect

	// beforeRequest executes immediately before sending HTTP Request. You can pass the options to the Fetch API here
	beforeRequest BeforeRequestFunc

	// modifyRequest rewrites http.Request
	modifyRequest ModifyRequestFunc
	// modifyResponse rewrites http.Response
	modifyResponse ModifyResponseFunc
}

func (t *Transport) SetMode(mode FetchMode) {
	t.Mode = mode
}

func (t *Transport) SetRedirect(redirect FetchRedirect) {
	t.Redirect = redirect
}

// New returns Transport struct
func New(opts ...Option) *Transport {
	t := &Transport{
		Mode:     FetchModeSameOrigin,
		Redirect: FetchRedirectFollow,
	}
	// Use options
	for _, opt := range opts {
		opt(t)
	}

	return t
}
