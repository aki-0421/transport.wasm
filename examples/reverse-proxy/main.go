package main

import (
	"net/http"
	"net/http/httputil"

	"github.com/aki-0421/transport.wasm"
	"github.com/go-chi/chi/v5"
	"github.com/syumai/workers"
)

func addUserAgent(req *http.Request) error {
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/110.0")
	return nil
}

func noCache(res *http.Response) error {
	res.Header.Del("etag")
	res.Header.Set("cache-control", "no-cache")
	return nil
}

func main() {
	r := chi.NewRouter()

	t, err := transport.New(
		transport.WithModifyRequestFunc(addUserAgent),
		transport.WithModifyResponseFunc(noCache),
	)
	if err != nil {
		panic(err)
	}

	r.HandleFunc("/*", func(w http.ResponseWriter, r *http.Request) {
		proxy := httputil.ReverseProxy{
			Transport: t,
			Director: func(r *http.Request) {
				r.URL.Host = "api.github.com"
			},
		}

		proxy.ServeHTTP(w, r)
	})

	workers.Serve(r)
}
