package main

import (
	"net/http"
	"net/http/httputil"

	"github.com/aki-0421/transport.wasm"
	"github.com/go-chi/chi/v5"
	"github.com/syumai/workers"
)

func addUserAgent(r *http.Request) error {
	r.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/110.0")
	return nil
}

func main() {
	r := chi.NewRouter()

	t, err := transport.New(
		transport.WithModifyRequestFunc(addUserAgent),
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
