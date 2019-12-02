package config

import "github.com/spf13/viper"

type Config struct {
	Development bool
	Server      ServerConfig
	Logger      LoggerConfig
}

type ConfigLoader struct {
	FileName string
	FilePath string
}

func (l *ConfigLoader) Load() (*Config, error) {
	var C *Config

	viper.SetConfigName(l.FileName)
	viper.AddConfigPath(l.FilePath)
	if err := viper.ReadInConfig(); err != nil {
		return C, err
	}

	if err := viper.Unmarshal(&C); err != nil {
		return C, err
	}
	return C, nil
}
