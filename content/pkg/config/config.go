package config

import (
	"log"
	"os"

	"content"

	"github.com/spf13/viper"
)

type Config struct {
	AppPort    string `mapstructure:"APP_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBName     string `mapstructure:"DB_NAME"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`

	UserAccessToken  string `mapstructure:"USER_ACCESS_TOKEN"`
	UserRefreshToken string `mapstructure:"USER_REFRESH_TOKEN"`

	ContributorAccessToken  string `mapstructure:"CONTRIBUTOR_ACCESS_TOKEN"`
	ContributorRefreshToken string `mapstructure:"CONTRIBUTOR_REFRESH_TOKEN"`

	AccessTokenExpiryHour  int `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour int `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`

	AwsBucket          string `mapstructure:"AWS_BUCKET"`
	AwsRegion          string `mapstructure:"AWS_REGION"`
	AwsAccessKey       string `mapstructure:"AWS_ACCESS_KEY"`
	AwsSecretAccessKey string `mapstructure:"AWS_SECRET_ACCESS_KEY"`
}

var cfg Config

func InitConfig() Config {

	// Set the current working directory to the directory containing the .env file
	if err := os.Chdir(content.ImageCurrentWorkingDirectory()); err != nil {
		log.Fatal("Error setting current working directory:", err)
	}

	// viper.AddConfigPath("../")
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err.Error())
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatal(err.Error())
	}
	return cfg
}

func GetConfig() Config {
	return cfg
}
