package env

import (
	"github.com/caarlos0/env/v8"
	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Environment    string `env:"APP_ENV,unset"                       envDefault:"debug"`
	Port           int16  `env:"APP_PORT,unset"                      envDefault:"5000"`
	RedisPort      int16  `env:"REDIS_PORT,unset"                    envDefault:"5000"`
	Google_API_Key string `env:"GOOGLE_API_KEY,unset"`
}

func LoadConfig() *Config {
	cfg := new(Config)
	if err := env.Parse(cfg); err != nil {
		panic(err)
	}
	return cfg
}
