package thirdservice

import (
	"github.com/go-kratos/kratos-layout/internal/conf"
	"github.com/go-kratos/kratos-layout/internal/middleware"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewService)

type ThirdService struct {
	logger *log.Helper
	config *conf.ThirdService
}

func NewService(config *conf.ThirdService, logger log.Logger) *ThirdService {
	return &ThirdService{logger: log.NewHelper(
		log.With(logger,
			"metadata",
			middleware.MetadataLog(map[string]bool{
				middleware.RequestIdMetaKey: true,
			})))}
}
