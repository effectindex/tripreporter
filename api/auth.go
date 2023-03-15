package api

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/cristalhq/jwt/v4"
	"github.com/effectindex/tripreporter/models"
	"github.com/effectindex/tripreporter/types"
	"github.com/effectindex/tripreporter/util"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func AuthMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			refreshToken, err := r.Cookie(util.CookieRefreshToken)
			if err != nil {
				ctx.Context.Logger.Debugw("Failed to get refresh token", zap.Error(err))
				ctx.Handle(w, r, MsgForbidden)
				return
			}

			jwtToken, _ := r.Cookie(util.CookieJwtToken)
			sessionClaims, _ := AccountIDFromToken(jwtToken) // This will actually verify that our jwtToken cookie is valid and not expired

			// If we need to generate a new access token
			if sessionClaims == nil || sessionClaims.Account.UUID == uuid.Nil {
				// First make sure we have a valid refresh token
				account, err := (&models.Account{Context: ctx.Context}).FromRefreshToken(refreshToken)
				if err != nil || &account.ID == nil {
					ctx.Context.Logger.Debugw("Failed to validate refresh token", zap.Error(err))
					ctx.Handle(w, r, MsgForbidden)
					return
				}

				// Refresh token is okay, build the access token
				expiryTime := time.Now().Add(time.Minute * 15) // TODO: Change this once we've implemented refreshing
				claims := &models.SessionClaims{
					RegisteredClaims: jwt.RegisteredClaims{
						Audience:  jwt.Audience([]string{"account"}),
						IssuedAt:  jwt.NewNumericDate(time.Now()),
						ExpiresAt: jwt.NewNumericDate(expiryTime),
					},
					Account: uuid.NullUUID{
						UUID:  account.ID,
						Valid: true,
					},
				}
				sessionClaims = claims // Update so we can use it outside before serving

				// Build the claims and set a new cookie to refresh the access token
				token, err := ctx.JwtBuilder.Build(claims)
				if err != nil {
					ctx.Logger.Warnw("Failed to create access token", zap.Error(err))
				}

				SetAuthCookie(w, util.CookieJwtToken, token.String(), expiryTime)

				ctx.Logger.Debugw("Successfully refreshed access token", "account", account.ID, "path", r.URL.Path)
			}

			// Set SessionClaims as the context value
			rCtx := r.Context()
			values := &models.ContextValues{SessionClaims: sessionClaims}
			rCtx = context.WithValue(rCtx, models.ContextValuesKey, values)
			r = r.WithContext(rCtx)

			ctx.Logger.Debugw("Successfully authenticated user", "account", sessionClaims.Account.UUID, "path", r.URL.Path)
			next.ServeHTTP(w, r)
		})
	}
}

func AccountIDFromToken(cookie *http.Cookie) (*models.SessionClaims, error) {
	if cookie == nil || len(cookie.Value) == 0 {
		ctx.Logger.Debugw("Failed to get JWT cookie because nil")
		return nil, nil
	}

	// Create a HMAC verifier
	verifier, err := jwt.NewVerifierHS(jwt.HS512, ctx.JwtKey)
	if err != nil {
		ctx.Logger.Debugw("Failed to make JWT verifier", zap.Error(err))
		return nil, err
	}

	// Decode the token and verify the signature
	token, err := jwt.Parse([]byte(cookie.Value), verifier)
	if err != nil {
		ctx.Logger.Debugw("Failed to verify JWT contents", zap.Error(err))
		return nil, err
	}

	// get Registered claims
	var claims models.SessionClaims
	err = json.Unmarshal(token.Claims(), &claims)
	if err != nil {
		ctx.Logger.Debugw("Failed to get token claims", zap.Error(err))
		return nil, err
	}

	// We need an account ID to reference
	if !claims.Account.Valid {
		ctx.Logger.Debugw("Failed to get token account ID because it is nil")
		return nil, types.ErrorSessionClaimNotValid
	}

	// Check if audience is a normal account, and if claims are currently valid
	if !claims.IsForAudience("account") || !claims.IsValidAt(time.Now()) {
		ctx.Logger.Debugw("Failed to verify claims audience or expiry", "registered claims", claims)
		return nil, types.ErrorSessionClaimNotValid
	}

	// Token has been validated, give it back
	return &claims, nil
}

func SetAuthCookie(w http.ResponseWriter, name string, token string, expiry time.Time) {
	http.SetCookie(w, &http.Cookie{
		Name:     name,
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Expires:  expiry,
		SameSite: http.SameSiteLaxMode,
		Secure:   true,
	})
}

func DeleteAuthCookies(w http.ResponseWriter, names ...string) {
	for _, name := range names {
		http.SetCookie(w, &http.Cookie{
			Name: name,
			Path: "/",
		})
	}
}
