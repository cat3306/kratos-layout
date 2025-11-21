package service

import (
	v1 "github.com/go-kratos/kratos-layout/api/server/v1"
	"github.com/go-kratos/kratos-layout/internal/middleware"
	"github.com/go-kratos/kratos-layout/internal/thirdmodule"
	"github.com/go-kratos/kratos/v2/log"
)

type Service struct {
	v1.UnimplementedServerServer
	thirdModule *thirdmodule.Module
	logger      *log.Helper
}

func NewService(thirdModule *thirdmodule.Module, logger log.Logger) *Service {
	return &Service{thirdModule: thirdModule, logger: log.NewHelper(
		log.With(logger,
			"metadata",
			middleware.MetadataLog(map[string]bool{
				middleware.RequestIdMetaKey: true,
			})))}
}
