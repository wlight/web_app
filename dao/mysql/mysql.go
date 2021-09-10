package mysql

import (
	"fmt"
	"web_app/settings"

	"go.uber.org/zap"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func Init(conf *settings.MysqlConfig) (err error) {
	// DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Dbname)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("mysql open failed", zap.Error(err))
		return
	}
	err = db.Ping()
	if err != nil {
		zap.L().Error("mysql connect failed", zap.Error(err))
		return
	}
	db.SetMaxOpenConns(conf.MaxOpenConns)
	db.SetMaxIdleConns(conf.MaxIdleConns)

	return
}

func Close() {
	_ = db.Close()
}
