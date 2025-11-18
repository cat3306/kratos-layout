package middleware

import (
	"context"
	"net"
	"strings"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/grpc/peer"
)

const (
	IpMetaKey = "ip"
)

func HttpClientIp() middleware.Middleware {
	return func(h middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (any, error) {
			meta := extractMetadata(ctx)
			meta[IpMetaKey] = httpTransportClientIP(ctx)
			ctx = context.WithValue(ctx, MetadataKey{}, meta)
			return h(ctx, req)
		}
	}
}

func GrpcClientIp() middleware.Middleware {
	return func(h middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (any, error) {
			meta := extractMetadata(ctx)
			meta[IpMetaKey] = grpcTransportClientIP(ctx)
			ctx = context.WithValue(ctx, MetadataKey{}, meta)
			return h(ctx, req)
		}
	}
}

func grpcTransportClientIP(ctx context.Context) string {
	p, ok := peer.FromContext(ctx)
	if !ok {
		return ""
	}
	remoteIP := net.ParseIP(remoteIP(p.Addr.String()))
	if remoteIP == nil {
		return ""
	}
	return remoteIP.String()
}

func httpTransportClientIP(ctx context.Context) string {
	request, ok := http.RequestFromServerContext(ctx)
	if !ok {
		return ""
	}
	ipKeys := []string{"X-Forwarded-For", "X-Real-Ip"}
	for _, k := range ipKeys {
		if ip := request.Header.Get(k); ip != "" {
			if ip, valid := validateHeader(ip); valid {
				return ip
			}
		}
	}
	remoteIP := net.ParseIP(remoteIP(request.RemoteAddr))
	if remoteIP == nil {
		return ""
	}
	return remoteIP.String()
}
func remoteIP(remoteAddr string) string {
	ip, _, err := net.SplitHostPort(strings.TrimSpace(remoteAddr))
	if err != nil {
		return ""
	}
	return ip
}
func validateHeader(header string) (clientIP string, valid bool) {
	if header == "" {
		return "", false
	}
	items := strings.Split(header, ",")
	for i := len(items) - 1; i >= 0; i-- {
		ipStr := strings.TrimSpace(items[i])
		ip := net.ParseIP(ipStr)
		if ip == nil {
			break
		}
		if i == 0 {
			return ipStr, true
		}
	}
	return "", false
}
