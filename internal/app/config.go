package app

import (
	"log"

	"github.com/spf13/viper"
)

type Configuration struct {
	ServerPort string `mapstructure:"SERVER_PORT"`

	CompanyName string `mapstructure:"COMPANY_NAME"`

	MaxUploadSize int    `mapstructure:"MAX_UPLOAD_SIZE"`
	FileFormats   string `mapstructure:"FILE_FORMATS"`

	EmailHost     string `mapstructure:"EMAIL_SERVER"`
	EmailPort     int    `mapstructure:"EMAIL_PORT"`
	EmailAddress  string `mapstructure:"EMAIL_ADDRESS"`
	EmailPassword string `mapstructure:"EMAIL_PASSWORD"`
}

// Reading configuration from file or environment variables.
func NewConfig() (*Configuration, error) {
	cfg := &Configuration{}

	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
		return nil, err
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
