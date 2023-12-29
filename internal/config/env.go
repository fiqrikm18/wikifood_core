package config

import (
	"os"
	"path"

	"github.com/spf13/viper"
)

type AppConfigOpt func(opt *ConfigOpt)

type ConfigOpt struct {
	FilePath string
}

type AppConfig struct {
	Port     string   `mapstructure:"port"`
	Database DBConfig `mapstructure:"database"`
}

type DBConfig struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Post     string `mapstructure:"port"`
	Name     string `mapstructure:"db_name"`
}

// NewAppConfig return all configuration from application
func NewAppConfig(opts ...AppConfigOpt) (*AppConfig, error) {
	options := ConfigOpt{}
	conf := AppConfig{}
	for _, opt := range opts {
		opt(&options)
	}

	workDir, _ := os.Getwd()
	confPath := path.Join(workDir, "../../config/")
	if len(opts) > 0 && options.FilePath != "" {
		confPath = options.FilePath
	}

	_, err := os.Stat(confPath + "/config.yaml")
	if err != nil {
		return nil, err
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(confPath)

	_ = viper.ReadInConfig()
	_ = viper.Unmarshal(&conf)

	return &conf, nil
}

// WithFilePath return function that inject value file path to the configuration
func WithFilePath(path string) AppConfigOpt {
	return func(opt *ConfigOpt) {
		opt.FilePath = path
	}
}
