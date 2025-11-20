package thirdmodule

import (
	"github.com/go-kratos/kratos-layout/internal/conf"
	"github.com/go-kratos/kratos-layout/internal/models/ent"
	"github.com/redis/go-redis/v9"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewModule)

// Data .
type Module struct {
	logger *log.Helper
	config *conf.Data
	ent    *ent.Client
	rdb    *redis.Client
}

// NewData .
func NewModule(c *conf.Data, logger log.Logger) (*Module, func(), error) {
	entClient, err := initEnt(c.GetDatabase())
	if err != nil {
		panic(err)
	}
	rdb, err := initRedis(c.GetRedis())
	if err != nil {
		panic(err)
	}
	cleanup := func() {
		entClient.Close()
		rdb.Close()
	}

	return &Module{
		logger: log.NewHelper(logger),
		config: c,
		ent:    entClient,
		rdb:    rdb,
	}, cleanup, nil
}

func (m *Module) Ent() *ent.Client {
	return m.ent
}

func (m *Module) RDB() *redis.Client {
	return m.rdb
}
