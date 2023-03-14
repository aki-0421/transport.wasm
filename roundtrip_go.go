//go:build !js || !wasm

package transport

import (
	"net/http"
)

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Modify http.request
	if t.modifyRequest != nil {
		err := t.modifyRequest(req)
		if err != nil {
			return nil, err
		}
	}
	// Exec beforeRequest hook
	if t.beforeRequest != nil {
		t.beforeRequest(t)
	}
	res, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	// Modify http.response
	if t.modifyResponse != nil {
		err := t.modifyResponse(res)
		if err != nil {
			return nil, err
		}
	}

	return res, nil
}
