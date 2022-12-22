package bootstrap

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go_code/gintest/bootstrap/gdb"
	"go_code/gintest/bootstrap/glog"
	"go_code/gintest/bootstrap/gredis"
	"log"
)

var Config = new(LoadConfig)

type LoadConfig struct {
	Name                string `mapstructure:"name"`
	Mode                string `mapstructure:"mode"`
	Version             string `mapstructure:"version"`
	StartTime           string `mapstructure:"start_time"`
	Port                int    `mapstructure:"port"`
	MachineId           int64  `mapstructure:"machine_id"`
	*glog.LoggerConfig  `mapstructure:"log"`
	*gdb.MySQLConfig    `mapstructure:"mysql"`
	*gredis.RedisConfig `mapstructure:"redis"`
}

func Initialize(path string) {
	// 初始化配置信息
	initConfig(path)
	// 初始化日志
	glog.InitLogger(Config.LoggerConfig, Config.Mode)
	// 初始化 MySQL
	gdb.InitMysql(Config.MySQLConfig)
	// 初始化 Redis
	gredis.InitRedis(Config.RedisConfig)
}

func initConfig(path string) {
	viper.SetConfigFile(path)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("读取配置信息错误, %s\n", err))
	}
	if err = viper.Unmarshal(Config); err != nil {
		panic(fmt.Errorf("反序列化配置信息错误, %s\n", err))
	}

	// 监控配置文件更新
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Println("配置文件发生改变", in.Name)
		if err = viper.Unmarshal(Config); err != nil {
			panic(fmt.Errorf("重载配置错误: %s\n", err))
		}
	})
	viper.WatchConfig()
}
