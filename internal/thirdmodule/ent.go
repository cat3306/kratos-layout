package thirdmodule

import (
	"time"

	"entgo.io/ent/dialect/sql"

	"github.com/go-kratos/kratos-layout/internal/conf"
	"github.com/go-kratos/kratos-layout/internal/models/ent"
	_ "github.com/go-sql-driver/mysql"
	// _ "github.com/mattn/go-sqlite3"
	// _ "github.com/lib/pq"
)

func initEnt(config *conf.Data_Database) (*ent.Client, error) {
	db, err := sql.Open(config.Driver, config.Source)
	if err != nil {
		return nil, err
	}
	if config.MaxOpenConns > 0 {
		db.DB().SetMaxOpenConns(int(config.MaxOpenConns))
	}

	if config.MaxIdleConns > 0 {
		db.DB().SetMaxIdleConns(int(config.MaxOpenConns))
	}
	if config.ConnMaxIdleTimeSeconds > 0 {
		db.DB().SetConnMaxIdleTime(time.Duration(config.ConnMaxIdleTimeSeconds) * time.Second)
	}
	if config.ConnMaxLifetimeSeconds > 0 {
		db.DB().SetConnMaxLifetime(time.Duration(config.ConnMaxLifetimeSeconds) * time.Second)
	}
	client := ent.NewClient(ent.Driver(db))
	return client, err
}
