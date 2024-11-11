package config

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	ClientID     string
	ClientSecret string
	TenantID     string
	RedirectURL  string

	AllowedOrigins string

	ServerPort string

	TokenSecret    string
	TokenApi       string
	TokenExpiresIn time.Duration
	TokenMaxAge    int
	DataInicio     time.Time
	ProductMode    string

	MailBox string
}

func LoadConfig(path string) (Config, error) {
	var config Config
	productMode := os.Getenv("PRODUCT_MODE")
	switch productMode {
	case "prod":
		durationStr := "600m"
		// Convertendo a string para time.Duration
		duration, err := time.ParseDuration(durationStr)
		if err != nil {
			fmt.Println("Erro ao parsear duração:", err)
		}
		config = Config{
			ProductMode:    os.Getenv("PRODUCT_MODE"),
			ClientID:       os.Getenv("CLIENT_ID"),
			ClientSecret:   os.Getenv("CLIENT_SECRET"),
			TenantID:       os.Getenv("TENANT_ID"),
			RedirectURL:    os.Getenv("REDIRECT_URL"),
			ServerPort:     os.Getenv("SERVER_PORT"),
			TokenSecret:    os.Getenv("TOKEN_SECRET"),
			TokenApi:       os.Getenv("TOKEN_API"),
			AllowedOrigins: os.Getenv("ALLOWED_ORIGINS"),
			MailBox:        os.Getenv("MAIL_BOX"),
			TokenExpiresIn: duration,
			TokenMaxAge:    60,
		}
	case "hmlg":
		config = Config{
			ProductMode:    os.Getenv("PRODUCT_MODE"),
			ClientID:       os.Getenv("CLIENT_ID"),
			ClientSecret:   os.Getenv("CLIENT_SECRET"),
			TenantID:       os.Getenv("TENANT_ID"),
			RedirectURL:    os.Getenv("REDIRECT_URL"),
			ServerPort:     os.Getenv("SERVER_PORT"),
			TokenSecret:    os.Getenv("TOKEN_SECRET"),
			TokenApi:       os.Getenv("TOKEN_API"),
			AllowedOrigins: os.Getenv("ALLOWED_ORIGINS"),
			MailBox:        os.Getenv("MAIL_BOX"),
		}
	case "local":
		config = Config{
			ProductMode:    os.Getenv("PRODUCT_MODE"),
			ClientID:       os.Getenv("CLIENT_ID"),
			ClientSecret:   os.Getenv("CLIENT_SECRET"),
			TenantID:       os.Getenv("TENANT_ID"),
			RedirectURL:    os.Getenv("REDIRECT_URL"),
			ServerPort:     os.Getenv("SERVER_PORT"),
			TokenSecret:    os.Getenv("TOKEN_SECRET"),
			TokenApi:       os.Getenv("TOKEN_API"),
			AllowedOrigins: os.Getenv("ALLOWED_ORIGINS"),
			MailBox:        os.Getenv("MAIL_BOX"),
		}
	default:
		viper.AddConfigPath(path)
		viper.SetConfigType("env")
		viper.SetConfigName("app")
		viper.AutomaticEnv()
		err := viper.ReadInConfig()
		if err != nil {
			return config, err
		}
		err = viper.Unmarshal(&config)
		if err != nil {
			return config, err
		}
	}
	return config, nil
}
