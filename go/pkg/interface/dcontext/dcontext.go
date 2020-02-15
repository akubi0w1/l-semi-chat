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
	idKey     key = "id"
)

// SetUserID contextにuserIDを設定する
func SetUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userIDKey, userID)
}

// SetID contextにuserのIDを設定する
func SetID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, idKey, id)
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

// GetIDFromContext contextからIDの取得
func GetIDFromContext(ctx context.Context) (string, error) {
	var id string
	var err error
	if ctx.Value(idKey) != nil {
		id = ctx.Value(idKey).(string)
	}
	if id == "" {
		logger.Warn("context: id is empty")
		err = domain.BadRequest(errors.New("ID is empty"))
	}
	return id, err
}
