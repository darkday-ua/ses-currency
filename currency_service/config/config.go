package config

import(
	"fmt"
//	"os"
	"context"
	"github.com/sethvargo/go-envconfig"
	"github.com/joho/godotenv"

)

var	packageVersion string = "0.0.1"


type AppConfig struct {
	DB_PATH string `env:"DB_PATH,default=/tmp/currency_service_users.db"`
	APP_PORT string `env:"APP_PORT,default=8080"`
	SMTP_HOST string `env:"SMTP_HOST"`
	SMTP_PORT string `env:"SMTP_PORT"`
	SMTP_USER string `env:"SMTP_USER"`
	SMTP_PASSWORD string `env:"SMTP_PASSWORD"`
	SMTP_FROM string `env:"SMTP_FROM"`
	LIVE_COIN_WATCH_API_KEY string `env:"LIVE_COIN_WATCH_API_KEY"`}

var	Config AppConfig

func init() {
	// for local dev or can be injected in docker container
	godotenv.Load(".env")
	ctx := context.Background()

	envconfig.Process(ctx, &Config)
	
	fmt.Printf("..config package version %s\n%+v\n", packageVersion,Config)
}

