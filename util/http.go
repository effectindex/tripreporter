package util

import (
	"go.uber.org/zap"
	"net/http/httputil"
	"net/url"
)

// NewProxy takes target host and creates a reverse proxy
func NewProxy(target string, logger *zap.SugaredLogger) *httputil.ReverseProxy {
	u, err := url.Parse(target)
	if err != nil {
		logger.Panicf("failed to parse target: %v", err) // likely malformed addr
	}

	return httputil.NewSingleHostReverseProxy(u)
}
