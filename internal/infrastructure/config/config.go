package config

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Config struct {
	App    AppConfig
	Server ServerConfig
}

type AppConfig struct {
	Debug bool
}

type ServerConfig struct {
	Host string
	Port int
}

func Load() (Config, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.AddConfigPath("./config")

	v.SetEnvPrefix("APP")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	v.SetDefault("app.debug", false)
	v.SetDefault("server.host", "0.0.0.0")
	v.SetDefault("server.port", 8080)

	if err := v.ReadInConfig(); err != nil {
		var notFound viper.ConfigFileNotFoundError
		if !errors.As(err, &notFound) {
			return Config{}, err
		}
	}

	return Config{
		App: AppConfig{
			Debug: v.GetBool("app.debug"),
		},
		Server: ServerConfig{
			Host: v.GetString("server.host"),
			Port: v.GetInt("server.port"),
		},
	}, nil
}

func (c ServerConfig) Address() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

func (c ServerConfig) ZapFields() []zap.Field {
	return []zap.Field{
		zap.String("host", c.Host),
		zap.Int("port", c.Port),
	}
}

func (c ServerConfig) ErrorField(err error) zap.Field {
	return zap.Error(err)
}
