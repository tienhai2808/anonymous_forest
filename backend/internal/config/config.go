package config

import (
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
		CAddr     string `mapstructure:"redis_addr"`
		CPassword string `mapstructure:"redis_password"`
	} `mapstructure:"cache"`
}

func LoadConfig() (*Config, error) {
	viper.AutomaticEnv()
	viper.BindEnv("database.mongo_uri", "MONGO_URI")
	viper.BindEnv("database.mongo_db", "MONGO_DB")
	viper.BindEnv("cache.redis_addr", "REDIS_ADDR")
	viper.BindEnv("cache.redis_password", "REDIS_PASSWORD")

	viper.BindEnv("app.port", "APP_PORT")
	viper.BindEnv("app.api_prefix", "APP_API_PREFIX")
	viper.BindEnv("app.client_token", "APP_CLIENT_TOKEN")
	viper.BindEnv("app.token_expires_in", "APP_TOKEN_EXPIRES_IN")
	viper.BindEnv("app.secure_cookie", "APP_SECURE_COOKIE")
	viper.BindEnv("app.http_cookie", "APP_HTTP_COOKIE")

	viper.BindEnv("app.cors.allow_origins", "APP_CORS_ALLOW_ORIGINS")
	viper.BindEnv("app.cors.allow_methods", "APP_CORS_ALLOW_METHODS")
	viper.BindEnv("app.cors.allow_headers", "APP_CORS_ALLOW_HEADERS")
	viper.BindEnv("app.cors.allow_credentials", "APP_CORS_ALLOW_CREDENTIALS")
	viper.BindEnv("app.cors.allow_files", "APP_CORS_ALLOW_FILES")
	viper.BindEnv("app.cors.allow_private_network", "APP_CORS_ALLOW_PRIVATE_NET_WORK")
	viper.BindEnv("app.cors.max_age", "APP_CORS_MAX_AGE")
	viper.BindEnv("app.cors.expose_headers", "APP_CORS_EXPOSE_HEADERS")

	viper.BindEnv("app.http.prefork", "APP_HTTP_PREFORK")
	viper.BindEnv("app.http.write_timeout", "APP_HTTP_WRITE_TIMEOUT")
	viper.BindEnv("app.http.read_timeout", "APP_HTTP_READ_TIMEOUT")
	viper.BindEnv("app.http.idle_timeout", "APP_HTTP_IDLE_TIMEOUT")
	viper.BindEnv("app.http.body_limit", "APP_HTTP_BODY_LIMIT")

	viper.SetConfigType("yml")
	viper.SetConfigName("config")
	viper.AddConfigPath("./configs")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, err
		}
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
