
package pkg

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

const (
	maxConn           = 50
	healthCheckPeriod = 3 * time.Minute
	maxConnIdleTime   = 1 * time.Minute
	maxConnLifetime   = 3 * time.Minute
	minConns          = 10
	lazyConnect       = false
)

// NewPgxConn pool
func NewPgxConn(env *Env) (*pgxpool.Pool, error) {

	username := env.DBUser
	password := env.DBPass
	host := env.DBHost
	port := env.DBPort
	dbname := env.DBName
	ctx := context.Background()
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		username,
		password,
		dbname,
	)

	poolCfg, err := pgxpool.ParseConfig(dataSourceName)
	if err != nil {
		return nil, err
	}

	poolCfg.MaxConns = maxConn
	poolCfg.HealthCheckPeriod = healthCheckPeriod
	poolCfg.MaxConnIdleTime = maxConnIdleTime
	poolCfg.MaxConnLifetime = maxConnLifetime
	poolCfg.MinConns = minConns
	poolCfg.LazyConnect = lazyConnect

	connPool, err := pgxpool.ConnectConfig(ctx, poolCfg)
	if err != nil {
		return nil, errors.Wrap(err, "pgx.ConnectConfig")
	}

	fmt.Println("success connection")

	return connPool, nil
}



func Close(p *pgxpool.Pool) {
	if p != nil {
		p.Close()
	}
}
