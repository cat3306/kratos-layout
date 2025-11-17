package middleware

import (
	"context"

	"github.com/go-kratos/kratos/v2/middleware"
)

type (
	MetadataKey struct{}
)

func Metadata() middleware.Middleware {
	return func(h middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (any, error) {
			md := make(map[string]string)
			ctx = context.WithValue(ctx, MetadataKey{}, md)
			return h(ctx, req)
		}
	}
}

func extractMetadata(ctx context.Context) map[string]string {
	md, ok := ctx.Value(MetadataKey{}).(map[string]string)
	if !ok {
		return nil
	}
	return md
}
