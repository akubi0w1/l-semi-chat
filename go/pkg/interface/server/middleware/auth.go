package middleware

import (
	"context"
	"net/http"

	"l-semi-chat/pkg/interface/auth"
	"l-semi-chat/pkg/interface/dcontext"
	"l-semi-chat/pkg/interface/server/response"

	"github.com/dgrijalva/jwt-go"
)

// Authorized JWTの検証を行う
func Authorized(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		if ctx == nil {
			ctx = context.Background()
		}

		cookie, err := r.Cookie("x-token")
		if err != nil {
			response.BadRequest(w, err.Error())
			return
		}

		token, err := auth.VerifyToken(cookie.Value)
		if err != nil {
			response.BadRequest(w, err.Error())
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		userID := claims["user_id"].(string)

		ctx = dcontext.SetUserID(ctx, userID)

		nextFunc(w, r.WithContext(ctx))
	}
}
