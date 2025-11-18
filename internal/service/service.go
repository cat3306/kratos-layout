package service

import (
	"context"

	v1 "github.com/go-kratos/kratos-layout/api/server/v1"
	"github.com/go-kratos/kratos-layout/internal/thirdmodule"
	"github.com/go-kratos/kratos/v2/log"
)

// GreeterService is a greeter service.
type Service struct {
	v1.UnimplementedServerServer
	thirdModule *thirdmodule.Module
	logger      log.Logger
}

// NewGreeterService new a greeter service.
func NewService(thirdModule *thirdmodule.Module, logger log.Logger) *Service {
	return &Service{thirdModule: thirdModule, logger: logger}
}

// SayHello implements helloworld.GreeterServer.
func (s *Service) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	//TODO logic here
	//s.logger.WithContext(ctx).Infof("SayHello Received: %v", in.Name)
	return &v1.HelloReply{Message: "Hello " + in.Name}, nil
}
