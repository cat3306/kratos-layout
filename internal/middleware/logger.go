package middleware

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
)

var (
	LogMetadataKeys = map[string]string{
		ClientIpMetaKey:  "clientip",
		RequestIdMetaKey: "reqid",
	}
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
			level, stack := extractError(err)
			log.NewHelper(log.WithContext(ctx, logger)).Log(level,
				"proto", transportKind,
				"path", operation,
				"args", extractArgs(req),
				"code", code,
				"reason", reason,
				"stack", stack,
				"latency", time.Since(startTime),
				"metadata", MetadataLog(nil)(ctx),
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

func MetadataLog(set map[string]bool) log.Valuer {
	return func(ctx context.Context) any {

		formatMetadata := func() string {
			if md, ok := metadata.FromServerContext(ctx); ok {
				str := strings.Builder{}
				cnt := 0
				plen := len(md)
				if len(set) != 0 {
					plen = len(set)
				}

				for k, v := range md {
					if len(v) == 0 {
						continue
					}
					if set != nil && !set[k] {
						continue
					}
					k = LogMetadataKeys[k]
					fmt.Fprintf(&str, "%s:%s", k, v[0])
					cnt++
					if cnt < plen {
						str.WriteString(" ")
					}
				}
				return str.String()
			}
			return ""
		}
		return formatMetadata()
	}
}
