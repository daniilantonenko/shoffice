package app

import (
	"github.com/spf13/viper"
)

type Configuration struct {
	CompanyName   string `mapstructure:"COMPANY_NAME"`
	MaxUploadSize int    `mapstructure:"MAX_UPLOAD_SIZE"`
	ServerPort    string `mapstructure:"SERVER_PORT"`
	ConfigMail    ConfigurationEmail
}

type ConfigurationEmail struct {
	EmailServer   string `mapstructure:"EMAIL_SERVER"`
	EmailPort     int    `mapstructure:"EMAIL_PORT"`
	EmailAddress  string `mapstructure:"EMAIL_ADDRESS"`
	EmailPassword string `mapstructure:"EMAIL_PASSWORD"`
}

type Config struct {
	EmailServer   string
	EmailPort     int
	FromEmail     string
	FromPass      string
	CompanyName   string
	FileFormats   []string
	MaxUploadSize int64
	Mode          string
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Configuration, err error) {
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
