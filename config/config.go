package config

import (
	"github.com/spf13/viper"
	"github.com/webhook-issue-manager/model"
)

func Config(file string) *model.Config {
	var config model.Config
	vi := viper.New()
	vi.SetConfigFile(file)
	vi.ReadInConfig()
	config.Port = vi.GetInt("port")
	config.Hostname = vi.GetString("hostname")
	config.User = vi.GetString("postgres_user")
	config.Password = vi.GetInt("postgres_password")
	config.Database = vi.GetString("postgres_database")

	return &config
}
