package gdb

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go_code/gintest/bootstrap/glog"
	"time"
)

var DB *sqlx.DB

type MySQLConfig struct {
	Host        string `mapstructure:"host"`
	User        string `mapstructure:"user"`
	Password    string `mapstructure:"password"`
	Dbname      string `mapstructure:"dbname"`
	Port        int    `mapstructure:"port"`
	MaxOpenConn int    `mapstructure:"max_open_conn"`
	MaxLifetime int    `mapstructure:"max_lifetime"`
	MaxIdleConn int    `mapstructure:"max_idle_conn"`
}

func InitMysql(config *MySQLConfig)  {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Dbname,
	)

	var err error
	DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		glog.SL.Error("连接 MySQL 错误，", err)
		panic(err)
	}
	DB.SetMaxOpenConns(config.MaxOpenConn)
	DB.SetConnMaxLifetime(time.Second * time.Duration(config.MaxLifetime))
	DB.SetMaxIdleConns(config.MaxIdleConn)
}

func Close() {
	defer func(DB *sqlx.DB) {
		_ = DB.Close()
	}(DB)
}