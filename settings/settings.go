package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Version      string `mapstructure:"version"`
	Port         int    `mapstructure:"port"`
	Mode         string `mapstructure:"mode"`
	StartTime    string `mapstructure:"start_time"`
	MachineID    int64  `mapstructure:"machine_id"`
	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	FileName   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	DBName       string `mapstructure:"dbname"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	DBName   int    `mapstructure:"dbname"`
	Password string `mapstructure:"password"`
	PoolSize int    `mapstructure:"pool_size"`
}

//func Init() (err error) {
//	//  设置配置文件信息
//	viper.SetConfigName("config")
//	viper.SetConfigType("yaml")
//	viper.AddConfigPath(".")
//	// 读取配置文件内容
//	err = viper.ReadInConfig()
//	if err != nil {
//		fmt.Printf("viper.ReadInConfig failed, err:%v \n", err)
//		return
//		panic(fmt.Errorf("Fatal error config file: %s \n", err))
//	}
//	// 监测配置文件是否修改
//	viper.WatchConfig()
//	viper.OnConfigChange(func(e fsnotify.Event) {
//		fmt.Println("Config file Changed ...")
//	})
//	return nil
//}

func Init() (err error) {
	//  设置配置文件信息
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	// 读取配置文件内容
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("viper.ReadInConfig failed, err:%v \n", err)
		return
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	// Unmarshal config
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed, err%v\n", err)
	}

	// 监测配置文件是否修改
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file Changed ...")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal failed again, err%v\n", err)
		}
	})
	return
}
