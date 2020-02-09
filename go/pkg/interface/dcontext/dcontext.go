package dcontext

import (
	"context"
	"errors"
	"l-semi-chat/pkg/domain"
	"l-semi-chat/pkg/domain/logger"
)

type key string

const (
	userIDKey key = "userID"
)

// SetUserID contextにuserIDを設定する
func SetUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userIDKey, userID)
}

// GetUserIDFromContext contextからuserIDの取得
func GetUserIDFromContext(ctx context.Context) (string, error) {
	var userID string
	var err error
	if ctx.Value(userIDKey) != nil {
		userID = ctx.Value(userIDKey).(string)
	}
	if userID == "" {
		logger.Warn("context: userID is empty")
		err = domain.BadRequest(errors.New("userID is empty"))
	}
	return userID, err
}
