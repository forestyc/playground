package db

import (
	"database/sql"
	"errors"
	driver "github.com/go-sql-driver/mysql"
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

func (mdb *Mysql) Close() error {
	return mdb.sqlDb.Close()
}

func (mdb *Mysql) DBError(err error) (int, string) {
	if err != nil {
		var mysqlErr *driver.MySQLError
		if ok := errors.As(err, &mysqlErr); ok {
			return int(mysqlErr.Number), mysqlErr.Message
		} else {
			return 0, err.Error()
		}
	} else {
		return 0, ""
	}
}
