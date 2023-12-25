package xcontext

import (
	"context"
)

type userIdCtxKey string

const userIdCtxKeyValue userIdCtxKey = "userId"

type requestIdCtxKey string

const requestIdCtxKeyValue requestIdCtxKey = "requestId"

func WithUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userIdCtxKeyValue, userID)
}

func GetUserID(ctx context.Context) string {
	userID, ok := ctx.Value(userIdCtxKeyValue).(string)
	if !ok {
		return ""
	}
	return userID
}

func WithRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, requestIdCtxKeyValue, requestID)
}

func GetRequestID(ctx context.Context) string {
	value := ctx.Value(requestIdCtxKeyValue)
	if value == nil {
		return ""
	}
	requestID, ok := value.(string)
	if !ok {
		return ""
	}
	return requestID
}
