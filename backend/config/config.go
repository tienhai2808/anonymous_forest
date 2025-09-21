package config

import (
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Port      int    `mapstructure:"port"`
		ApiPrefix string `mapstructure:"api_prefix"`
		Cors      struct {
			AllowOrigins        string `mapstructure:"allow_origins"`
			AllowMethods        string `mapstructure:"allow_methods"`
			AllowHeaders        string `mapstructure:"allow_headers"`
			AllowCredentials    bool   `mapstructure:"allow_credentials"`
			AllowFiles          bool   `mapstructure:"allow_files"`
			AllowPrivateNetwork bool   `mapstructure:"allow_private_network"`
			MaxAge              int    `mapstructure:"max_age"`
			ExposeHeaders       string `mapstructure:"expose_headers"`
		} `mapstructure:"cors"`
		Http struct {
			Prefork      bool          `mapstructure:"prefork"`
			WriteTimeout time.Duration `mapstructure:"write_timeout"`
			ReadTimeout  time.Duration `mapstructure:"read_timeout"`
			IdleTimeout  time.Duration `mapstructure:"idle_timeout"`
			BodyLimit    int           `mapstructure:"body_limit"`
		} `mapstructure:"http"`
		ClientToken    string        `mapstructure:"client_token"`
		TokenExpiresIn time.Duration `mapstructure:"token_expires_in"`
		SecureCookie   bool          `mapstructure:"secure_cookie"`
		HttpCookie     bool          `mapstructure:"http_cookie"`
	} `mapstructure:"app"`

	Database struct {
		DBUri  string `mapstructure:"mongo_uri"`
		DBName string `mapstructure:"mongo_db"`
	} `mapstructure:"database"`

	Cache struct {
		CAddr string `mapstructure:"redis_addr"`
		CDb   int    `mapstructure:"redis_db"`
	} `mapstructure:"cache"`
}

func LoadConfig() (*Config, error) {
	cfgFile := os.Getenv("CONFIG_FILE")
	if cfgFile == "" {
		cfgFile = "config/config.yaml"
	}
	viper.SetConfigFile(cfgFile)
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
