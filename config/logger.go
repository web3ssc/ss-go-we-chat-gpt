package config

import (
	"github.com/spf13/viper"
	"github.com/web3ssc/ss-go-common/log"
)

type LoggerConfig struct {
	Writers        string
	LoggerLevel    string
	LoggerFile     string
	LogFormatText  bool
	RollingPolicy  string
	LogRotateDate  int
	LogRotateSize  int
	LogBackupCount int
}

func LoadLoggerConfig(viper *viper.Viper) *LoggerConfig {
	cfg := &LoggerConfig{
		Writers:        viper.GetString("writers"),
		LoggerLevel:    viper.GetString("logger_level"),
		LoggerFile:     viper.GetString("logger_file"),
		LogFormatText:  viper.GetBool("log_format_text"),
		RollingPolicy:  viper.GetString("rolling_policy"),
		LogRotateDate:  viper.GetInt("log_rotate_date"),
		LogRotateSize:  viper.GetInt("log_rotate_size"),
		LogBackupCount: viper.GetInt("log_backup_count"),
	}
	return cfg
}

func (cfg *LoggerConfig) InitLogger() {
	passLagerCfg := log.PassLagerCfg{
		Writers:        cfg.Writers,
		LoggerLevel:    cfg.LoggerLevel,
		LoggerFile:     cfg.LoggerFile,
		LogFormatText:  cfg.LogFormatText,
		RollingPolicy:  cfg.RollingPolicy,
		LogRotateDate:  cfg.LogRotateDate,
		LogRotateSize:  cfg.LogRotateSize,
		LogBackupCount: cfg.LogBackupCount,
	}
	if err := log.InitWithConfig(&passLagerCfg); err != nil {
		panic(any("InitLogger error"))
	}
}
