package middleware

import (
	"context"

	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/google/uuid"
)

const (
	RequestIdMetaKey = "x-md-global-reqid"
)

func RequestId() middleware.Middleware {
	return func(h middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (any, error) {
			if md, ok := metadata.FromServerContext(ctx); ok {
				md.Set(RequestIdMetaKey, generateRequestID())
			}
			return h(ctx, req)
		}
	}
}
func generateRequestID() string {
	return uuid.New().String()
}
