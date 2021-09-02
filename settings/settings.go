package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Name  string `mapstructure:"name"`
	Mode  string `mapstructure:"mode"`
	Port  int    `mapstructure:"port"`
	Log   *LogConfig
	Mysql *MysqlConfig
	Redis *RedisConfig
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}
type MysqlConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	Dbname       string `mapstructure:"dbname"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Db       int    `mapstructure:"db"`
}

func Init() (err error) {
	viper.SetConfigFile("config.yaml")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("init settings failed, err:%v\n", err)
		return
	}
	// 将读取的配置信息保存至全局变量 Conf
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("Unmarshal settings failed, err:%v\n", err)
		return err
	}

	// 监控配置文件变化
	viper.WatchConfig()

	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("配置文件更改成功")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("Unmarshal settings failed, err:%v\n", err)
			return
		}
	})
	return
}
