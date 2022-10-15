package util

import (
	"log"
	"net/http/httputil"
	"net/url"
)

// NewProxy takes target host and creates a reverse proxy
func NewProxy(target string) *httputil.ReverseProxy {
	u, err := url.Parse(target)
	if err != nil {
		log.Fatalf("error making reverse proxy: %v\n", err) // likely malformed addr
		return nil
	}

	return httputil.NewSingleHostReverseProxy(u)
}
