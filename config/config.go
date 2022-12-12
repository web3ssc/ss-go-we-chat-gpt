package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/web3ssc/ss-go-common/log"
	"os"
)

var (
	Cfg      *Config
	Nickname string
	Token    string
)

type Config struct {
	Logger *LoggerConfig
}

func Init(cfgName string) {
	setConfig(cfgName)
	Cfg = loadConfig()
	initConfig(Cfg)
	watchConfig()
}

func setConfig(cfgName string) {
	if cfgName != "" {
		viper.SetConfigFile(cfgName)
	} else {
		viper.AddConfigPath("profile")
		viper.SetConfigName("profile")
	}
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		panic(any("initConfig error"))
	}
}

func loadConfig() *Config {
	cfg := &Config{
		Logger: LoadLoggerConfig(viper.Sub("logger")),
	}
	return cfg
}

func initConfig(cfg *Config) {
	cfg.Logger.InitLogger()
}


func SetValue(nickname, token string) {
	Nickname = nickname
	Token = token
	if nickname == ""  {
		log.Errorf(nil, "nickname empty")
		os.Exit(0)
	}
	if token == "" {
		log.Errorf(nil, "token empty")
		os.Exit(0)
	}
	log.Infof("nickname:%v, token:%v", nickname, token)
}



func watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Infof("Config file changed: %s", e.Name)
	})
}
