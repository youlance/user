package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBAddress     string `mapstructure:"DB_ADDRESS"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPass        string `mapstructure:"DB_PASS"`
	DBName        string `mapstructure:"DB_NAME"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
