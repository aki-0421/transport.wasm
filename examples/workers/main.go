package main

import (
	"io"
	"net/http"

	"github.com/aki-0421/transport.wasm"
	"github.com/go-chi/chi/v5"
	"github.com/syumai/workers"
)

func main() {
	r := chi.NewRouter()

	t, err := transport.New()
	if err != nil {
		panic(err)
	}
	cli := &http.Client{
		Transport: t,
	}

	r.HandleFunc("/*", func(w http.ResponseWriter, r *http.Request) {
		resp, err := cli.Get("https://api.github.com/repos/aki-0421/transport.wasm")
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}

		// Read response body
		io.Copy(w, resp.Body)
	})

	workers.Serve(r)
}
