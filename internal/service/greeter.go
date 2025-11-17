package service

import (
	"context"

	v1 "github.com/go-kratos/kratos-layout/api/helloworld/v1"
	"github.com/go-kratos/kratos-layout/internal/thirdmodule"
	"github.com/go-kratos/kratos/v2/log"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer
	thirdModule *thirdmodule.Module
	logger      log.Logger
}

// NewGreeterService new a greeter service.
func NewGreeterService(thirdModule *thirdmodule.Module, logger log.Logger) *GreeterService {
	return &GreeterService{thirdModule: thirdModule, logger: logger}
}

// SayHello implements helloworld.GreeterServer.
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	//TODO logic here
	//s.logger.WithContext(ctx).Infof("SayHello Received: %v", in.Name)
	return &v1.HelloReply{Message: "Hello " + in.Name}, nil
}
