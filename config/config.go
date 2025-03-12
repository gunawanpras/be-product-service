package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

var (
	conf *Config
)

// option defines configuration option
type option struct {
	configFolder string
	configFile   string
	configType   string
}

// Init initializes `config` from the default config file.
// use `WithConfigFile` to specify the location of the config file
func Init(opts ...Option) error {
	opt := &option{
		configFolder: getDefaultConfigFolder(),
		configFile:   getDefaultConfigFile(),
		configType:   getDefaultConfigType(),
	}

	for _, optFunc := range opts {
		optFunc(opt)
	}

	// Config File Path
	viper.AddConfigPath(opt.configFolder)
	// Config File Name
	viper.SetConfigName(opt.configFile)
	// Config File Type
	viper.SetConfigType(opt.configType)
	viper.AutomaticEnv()

	// Application config --> config/config.yaml
	err := viper.ReadInConfig()
	if err != nil {
		log.Panic("failed to initialize config:", err)
	}

	if err = viper.Unmarshal(&conf); err != nil {
		log.Panic("failed to initialize config:", err)
	}

	return nil
}

// Option define an option for config package
type Option func(*option)

// WithConfigFolder set `config` to use the given config folder
func WithConfigFolder(configFolder string) Option {
	return func(opt *option) {
		opt.configFolder = configFolder
	}
}

// WithConfigFile set `config` to use the given config file
func WithConfigFile(configFile string) Option {
	return func(opt *option) {
		env := os.Getenv("GO_ENV")
		if env == "" {
			opt.configFile = configFile
			return
		}

		opt.configFile = configFile + "." + env
	}
}

// WithConfigType set `config` to use the given config type
func WithConfigType(configType string) Option {
	return func(opt *option) {
		opt.configType = configType
	}
}

// getDefaultConfigFolder get default config folder.
func getDefaultConfigFolder() string {
	configPath := "./"

	return configPath
}

// getDefaultConfigFile get default config file.
func getDefaultConfigFile() string {
	env := os.Getenv("GO_ENV")
	if env == "" {
		return "config"
	}

	return "config." + env
}

// getDefaultConfigType get default config type.
func getDefaultConfigType() string {
	return "yaml"
}

// Get config
func Get() *Config {
	if conf == nil {
		conf = &Config{}
	}
	return conf
}
