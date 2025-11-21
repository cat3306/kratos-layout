package service

import (
	"context"

	v1 "github.com/go-kratos/kratos-layout/api/server/v1"
)

func (s *Service) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	//TODO logic here
	//s.thirdModule.RDB().Set(ctx, "key", "value", 0)
	//s.thirdModule.Ent().Foo.Query().Limit(1).Count(ctx)
	s.logger.WithContext(ctx).Infof("SayHello Received: %v", in.Name)
	return &v1.HelloReply{Message: "Hello " + in.Name}, nil
}
