package middleware

import (
	"context"
	"net/http"

	"l-semi-chat/pkg/domain"
	"l-semi-chat/pkg/domain/logger"
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
			logger.Warn("middleware: ", err)
			response.HttpError(w, domain.BadRequest(err))
			return
		}

		token, err := auth.VerifyToken(cookie.Value)
		if err != nil {
			response.HttpError(w, domain.BadRequest(err))
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		userID := claims["user_id"].(string)
		id := claims["id"].(string)

		ctx = dcontext.SetUserID(ctx, userID)
		ctx = dcontext.SetID(ctx, id)

		nextFunc(w, r.WithContext(ctx))
	}
}
