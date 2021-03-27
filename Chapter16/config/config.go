package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/i-coder-robot/book_final_code/Chapter16/MyLog"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Name string
}

func Init(name string) error {
	c := Config{
		Name: name,
	}

	// 初始化配置文件
	if err := c.initConfig(); err != nil {
		MyLog.Log.Error(err)
		return err
	}

	// 监控配置文件变化并热加载程序
	c.watchConfig()

	return nil
}

func (c *Config) initConfig() error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name) // 如果指定了配置文件，则解析指定的配置文件
	} else {
		//viper.AddConfigPath("./Chapter16/conf/") //goland debug的配置
		viper.AddConfigPath("./conf/") //命令行的配置
		viper.SetConfigName("app_config")
	}
	viper.SetConfigType("yaml")                  // 设置配置文件格式为YAML
	if err := viper.ReadInConfig(); err != nil { // viper解析配置文件
		return err
	}

	return nil
}

// 监控配置文件变化并热加载程序
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed: %s", e.Name)
	})
}
