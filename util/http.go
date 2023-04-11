// SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>
//
// SPDX-License-Identifier: OSL-3.0

package util

import (
	"net/http/httputil"
	"net/url"

	"go.uber.org/zap"
)

const (
	CookieRefreshToken = "refresh-token"
	CookieJwtToken     = "access-token"
)

// NewProxy takes target host and creates a reverse proxy
func NewProxy(target string, logger *zap.SugaredLogger) *httputil.ReverseProxy {
	u, err := url.Parse(target)
	if err != nil {
		logger.Panicf("failed to parse target: %v", err) // likely malformed addr
	}

	return httputil.NewSingleHostReverseProxy(u)
}
