package db

import (
	"context"
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Config struct {
	Dsn              string `mapstructure:"dsn"`
	MaxOpen          int    `mapstructure:"max-open"`
	IdleConns        int    `mapstructure:"idle-conns"`
	MaxIdleTime      int    `mapstructure:"idle-timeout"`
	OperationTimeout int    `mapstructure:"operation-timeout"`
}

type Mysql struct {
	config Config
	db     *gorm.DB
	sqlDb  *sql.DB
}

func NewMysql(config Config) *Mysql {
	var err error
	var conn Mysql
	conn.config = config
	if conn.db, err = gorm.Open(mysql.Open(config.Dsn), &gorm.Config{}); err != nil {
		panic(err)
	} else {
		conn.sqlDb, err = conn.db.DB()
		conn.sqlDb.SetMaxIdleConns(config.IdleConns)
		conn.sqlDb.SetMaxOpenConns(config.MaxOpen)
		conn.sqlDb.SetConnMaxIdleTime(time.Duration(config.MaxIdleTime) * time.Second)
		return &conn
	}
}

func (mysql *Mysql) Session() (*sql.DB, error) {
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(mysql.config.OperationTimeout)*time.Second)
	db := mysql.db.Session(&gorm.Session{
		PrepareStmt: true,
		NewDB:       true,
		Context:     ctx,
	})
	return db.DB()
}

func (mysql *Mysql) Close() error {
	return mysql.sqlDb.Close()
}
