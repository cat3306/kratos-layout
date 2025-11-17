package middleware

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
)

func extractError(err error) (log.Level, string) {
	if err != nil {
		return log.LevelError, fmt.Sprintf("%+v", err)
	}
	return log.LevelInfo, ""
}

func Log(logger log.Logger) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (reply any, err error) {
			var (
				code          int32
				reason        string
				transportKind string
				operation     string
			)
			startTime := time.Now()
			if info, ok := transport.FromServerContext(ctx); ok {
				transportKind = info.Kind().String()
				operation = info.Operation()
			}
			reply, err = handler(ctx, req)
			if se := errors.FromError(err); se != nil {
				code = se.Code
				reason = se.Reason
			}
			formatMetadata := func() string {
				metadata := extractMetadata(ctx)
				str := strings.Builder{}
				cnt := 0
				for k, v := range metadata {
					str.WriteString(fmt.Sprintf("%s:%s", k, v))
					cnt++
					if cnt < len(metadata) {
						str.WriteString(" ")
					}
				}
				return str.String()
			}
			level, stack := extractError(err)
			log.NewHelper(log.WithContext(ctx, logger)).Log(level,
				"proto", transportKind,
				"path", operation,
				"args", extractArgs(req),
				"code", code,
				"reason", reason,
				"stack", stack,
				"latency", time.Since(startTime),
				"metadata", formatMetadata(),
			)
			return
		}
	}
}

func extractArgs(req any) string {
	if redacter, ok := req.(Redacter); ok {
		return redacter.Redact()
	}
	if stringer, ok := req.(fmt.Stringer); ok {
		return stringer.String()
	}
	return fmt.Sprintf("%+v", req)
}

type Redacter interface {
	Redact() string
}
