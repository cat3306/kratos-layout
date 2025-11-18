package middleware

import (
	"context"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/google/uuid"
)

type (
	MetadataKey struct{}
)

const (
	RequestIdMetaKey = "requestID"
)

func Metadata() middleware.Middleware {
	return func(h middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (any, error) {
			md := extractMetadata(ctx)
			md[RequestIdMetaKey] = generateRequestID()
			ctx = context.WithValue(ctx, MetadataKey{}, md)
			return h(ctx, req)
		}
	}
}
func generateRequestID() string {
	return uuid.New().String()
}
func extractMetadata(ctx context.Context) map[string]string {
	md, ok := ctx.Value(MetadataKey{}).(map[string]string)
	if !ok {
		return map[string]string{}
	}
	return md
}
