package storage

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"time"
)

type Database struct {
	DB                    *sqlx.DB
	host                  string
	port                  int
	user                  string
	password              string
	name                  string
	maxOpenConnections    int
	maxIdleConnections    int
	connectionMaxLifetime time.Duration
	connectionMaxIdletime time.Duration
	log                   *zap.Logger
}

type NewDatabaseOptions struct {
	Host                  string
	Port                  int
	User                  string
	Password              string
	Name                  string
	MaxOpenConnections    int
	MaxIdleConnections    int
	ConnectionMaxLifetime time.Duration
	ConnectionMaxIdletime time.Duration
	Log                   *zap.Logger
}

func NewDatabase(opts NewDatabaseOptions) *Database {
	if opts.Log == nil {
		opts.Log = zap.NewNop()
	}

	return &Database{
		host:                  opts.Host,
		port:                  opts.Port,
		user:                  opts.User,
		password:              opts.Password,
		name:                  opts.Name,
		maxOpenConnections:    opts.MaxOpenConnections,
		maxIdleConnections:    opts.MaxIdleConnections,
		connectionMaxLifetime: opts.ConnectionMaxLifetime,
		connectionMaxIdletime: opts.ConnectionMaxIdletime,
		log:                   opts.Log,
	}
}

func (d *Database) Connect() error {
	d.log.Info("Connecting to database", zap.String("url", d.createDataSourceName(false)))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error

	d.DB, err := sqlx.ConnectContext(ctx, "pgx", d.createDataSourceName(true))
	if err != nil{
		return err
	}

	d.log.Debug("setting connection pool options",
		zap.Int("max open connection", d.maxOpenConnections),
		zap.Int("max idle connections", d.maxIdleConnections),
		zap.Duration("connection max lifetime", d.connectionMaxLifetime),
		zap.Duration("connection max idle time", d.connectionMaxIdletime),
	)
	d.DB.SetMaxOpenConns(d.maxOpenConnections)
	d.DB.SetMaxIdleConns(d.maxIdleConnections)
	d.DB.SetConnMaxLifetime(d.connectionMaxLifetime)
	d.DB.SetConnMaxIdleTime(d.connectionMaxIdleTime)

	return nil
}

func (d *Database) createDataSourceName(withPassword bool) string {
	password := d.password

	if !withPassword {
		password = "XXX"
	}

	return fmt.Sprintf("postgresql://%v:%v@%v:%v/%v?sslmode=disable", d.user, password, d.host, d.port, d.name)
}

func (d *Database) Ping(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	if err := d.DB.PingContext(ctx); err != nil {
		return err
	}

	_, err := d.DB.ExecContext(ctx, `select 1`)
	return err
}
