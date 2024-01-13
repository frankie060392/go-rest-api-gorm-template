package bootstrap

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost                 string        `mapstructure:"POSTGRES_HOST"`
	DBUserName             string        `mapstructure:"POSTGRES_USER"`
	DBUserPassword         string        `mapstructure:"POSTGRES_PASSWORD"`
	DBName                 string        `mapstructure:"POSTGRES_DB"`
	DBPort                 string        `mapstructure:"POSTGRES_PORT"`
	ServerAddress          string        `mapstructure:"SERVER_ADDRESS"`
	AccessTokenPrivateKey  string        `mapstructure:"ACCESS_TOKEN_PRIVATE_KEY"`
	AccessTokenPublicKey   string        `mapstructure:"ACCESS_TOKEN_PUBLIC_KEY"`
	RefreshTokenPrivateKey string        `mapstructure:"REFRESH_TOKEN_PRIVATE_KEY"`
	RefreshTokenPublicKey  string        `mapstructure:"REFRESH_TOKEN_PUBLIC_KEY"`
	AccessTokenExpiresIn   time.Duration `mapstructure:"ACCESS_TOKEN_EXPIRED_IN"`
	RefreshTokenExpiresIn  time.Duration `mapstructure:"REFRESH_TOKEN_EXPIRED_IN"`
	AccessTokenMaxAge      int           `mapstructure:"ACCESS_TOKEN_MAXAGE"`
	RefreshTokenMaxAge     int           `mapstructure:"REFRESH_TOKEN_MAXAGE"`
	AwsAccessKeyId         string        `mapstructure:"AWS_ACCESS_KEY_ID"`
	AwsSecretKey           string        `mapstructure:"AWS_SECRET_ACCESS_KEY"`
	AwsS3Region            string        `mapstructure:"AWS_S3_REGION"`
	AwsS3Bucket            string        `mapstructure:"AWS_S3_BUCKET"`
	TelegramUrl            string        `mapstructure:"TELEGRAM_BOT_URL"`
	TelegramChatId         string        `mapstructure:"TELEGRAM_BOT_CHAT_ID"`
}

func LoadConfig(path string) *Config {
	config := Config{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}
	return &config
}
