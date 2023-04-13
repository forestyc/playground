package gorm

import (
	"context"
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type MysqlConf struct {
	Dsn              string
	MaxOpen          int
	MaxIdle          int
	MaxIdleTime      int
	OperationTimeout int
}

type Mysql struct {
	config MysqlConf
	db     *gorm.DB
	sqlDb  *sql.DB
}

func NewMysql(config MysqlConf) Mysql {
	var err error
	var conn Mysql
	conn.config = config
	if conn.db, err = gorm.Open(mysql.Open(config.Dsn), &gorm.Config{}); err != nil {
		panic(err)
	} else {
		conn.sqlDb, err = conn.db.DB()
		conn.sqlDb.SetMaxIdleConns(config.MaxOpen)
		conn.sqlDb.SetMaxOpenConns(config.MaxIdle)
		conn.sqlDb.SetConnMaxIdleTime(time.Duration(config.MaxIdleTime) * time.Second)
		return conn
	}
}

func (mysql Mysql) Session() (*gorm.DB, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(mysql.config.OperationTimeout)*time.Second)
	return mysql.db.Session(&gorm.Session{
		PrepareStmt: true,
		NewDB:       true,
		Context:     ctx,
	}), cancel
}

func (mysql Mysql) Close() {
	mysql.sqlDb.Close()
}
