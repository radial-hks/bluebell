package mysql

import (
	"bluebell/settings"
	"fmt"

	"go.uber.org/zap"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func Init(cfg *settings.MySQLConfig) (err error) {
	//dsn := "root:123456@(172.17.0.2:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName,
	)
	//
	DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("Connect Mysql Failed", zap.Error(err))
	}
	DB.SetMaxOpenConns(cfg.MaxOpenConns)
	DB.SetMaxIdleConns(cfg.MaxIdleConns)
	//DB.SetMaxOpenConns(20)
	//DB.SetMaxIdleConns(10)
	return
}

func Close() {
	_ = DB.Close()
	zap.L().Info("Mysql Close")
}
