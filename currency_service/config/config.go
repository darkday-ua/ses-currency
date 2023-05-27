package config

import(
	"fmt"
//	"os"
	"context"
	"github.com/sethvargo/go-envconfig"

)

var	packageVersion string = "0.0.1"


type AppConfig struct {
	DB_PATH string `env:"DB_PATH,default=/tmp/currency_service_users.db"`
	APP_PORT string `env:"APP_PORT,default=8080"`
	SMTP_HOST string `env:"SMTP_HOST"`
	SMTP_PORT string `env:"SMTP_PORT"`
	SMTP_USER string `env:"SMTP_USER"`
	SMTP_PASS string `env:"SMTP_PASS"`
	SMTP_FROM string `env:"SMTP_FROM"`
	LIVE_COIN_WATCH_API_KEY string `env:"LIVE_COIN_WATCH_API_KEY"`}

var	Config AppConfig

func init() {

	ctx := context.Background()

	envconfig.Process(ctx, &Config)
	
	fmt.Printf("..config package version %s\n%+v\n", packageVersion,Config)
}

