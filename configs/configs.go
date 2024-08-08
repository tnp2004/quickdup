package configs

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	Server   *Server
	Database *Database
	Auth     *Auth
}

type Server struct {
	Host         string
	Port         string
	AllowOrigins []string
	BodyLimit    string
	Timeout      uint16
}

type Database struct {
	Host     string
	Port     string
	Name     string
	Username string
	Password string
	Schema   string
	Sslmode  string
}

type Auth struct {
	Jwt        *Jwt
	Cloudinary *Cloudinary
}

type Jwt struct {
	AccessSecretKey            string
	RefreshSecretKey           string
	AccessTokenExpireDuration  uint
	RefreshTokenExpireDuration uint
}

type Cloudinary struct {
	ApiKey    string
	ApiSecret string
}

var (
	once           sync.Once
	ConfigInstance *Config
)

func NewConfig() *Config {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./configs")

		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				log.Fatal("config file not found")
			} else {
				log.Fatalf("error read config. Err: %s", err.Error())
			}
		}

		var config *Config
		if err := viper.UnmarshalExact(&config); err != nil {
			log.Fatalf("error to decode into struct. Err: %s", err.Error())
		}

		ConfigInstance = config
	})

	return ConfigInstance
}
