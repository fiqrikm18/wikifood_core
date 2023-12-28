package config

type AppConfigOpt func(opt *ConfigOpt)

type ConfigOpt struct {
	FilePath string
}

type AppConfig struct{}

// NewAppConfig return all configuration from application
func NewAppConfig(opts ...AppConfigOpt) *AppConfig {
	return &AppConfig{}
}

// WithFilePath return function that inject value file path to the configuration
func WithFilePath(path string) AppConfigOpt {
	return func(opt *ConfigOpt) {
		opt.FilePath = path
	}
}
