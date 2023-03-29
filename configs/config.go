package configs

import (
	"log"

	"github.com/go-chi/jwtauth/v5"
	"github.com/spf13/viper"
)

type config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
	JWTSecret     string `mapstructure:"JWT_SECRET"`
	JwtExpiresIn  int    `mapstructure:"JWT_EXPIRES_IN"`
	TokenAuth     *jwtauth.JWTAuth
}

func LoadConfig() *config {
	var cfg *config

	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("error while reading config file %s", err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)
	return cfg
}
