package thirdmodule

import (
	"github.com/go-kratos/kratos-layout/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewModule)

// Data .
type Module struct {
	logger *log.Helper
	// TODO wrapped database client
}

// NewData .
func NewModule(c *conf.Data, logger log.Logger) (*Module, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Module{
		logger: log.NewHelper(logger),
	}, cleanup, nil
}
