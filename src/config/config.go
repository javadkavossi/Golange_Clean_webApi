package config

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Redis    RedisConfig
	Password PasswordConfig
	Cors     CorsConfig
	Logger   LoggerConfig
	Otp      OtpConfig
	JWT      JWTConfig
}

type ServerConfig struct {
	Port         string
	InternalPort string
	ExternalPort string
	RunMode      string
}

type LoggerConfig struct {
	FilePath string
	Encoding string
	Level    string
	Logger   string
}

type PostgresConfig struct {
	Host            string
	Port            string
	User            string
	Password        string
	DbName          string
	SSLMode         string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

type RedisConfig struct {
	Host               string
	Port               string
	Password           string
	Db                 string
	DialTimeout        time.Duration
	ReadTimeout        time.Duration
	WriteTimeout       time.Duration
	IdleCheckFrequency time.Duration
	PoolSize           int
	PoolTimeout        time.Duration
}

type PasswordConfig struct {
	IncludeChars     bool
	IncludeDigits    bool
	MinLength        int
	MaxLength        int
	IncludeUppercase bool
	IncludeLowercase bool
}

type CorsConfig struct {
	AllowOrigins string
}

type OtpConfig struct {
	ExpireTime time.Duration
	Digits     int
	Limiter    time.Duration
}

type JWTConfig struct {
	AccessTokenExpireDuration  time.Duration
	RefreshTokenExpireDuration time.Duration
	Secret                     string
	RefreshSecret              string
}


func GetConfig() *Config {
	cfgPath := getConfigPath(os.Getenv("APP_ENV"))
	v, err := LoadConfig(cfgPath, "yml")
	if err != nil {
		log.Fatalf("Error in Load Config : %v", err)
	}
	cfg, err := ParsConfig(v)
	if err != nil {
		log.Fatalf("Error in pars Config : %v", err)
	}
	return cfg
}

func ParsConfig(v *viper.Viper) (*Config, error) {
	var cfg Config
	err := v.Unmarshal(&cfg)

	if err != nil {
		log.Printf("unable to parse config: %v", err)
		return nil, err
	}
	return &cfg, nil
}

func LoadConfig(filename string, filetype string) (*viper.Viper, error) {
	V := viper.New()
	V.SetConfigType(filetype)
	V.SetConfigName(filename)
	V.AddConfigPath(".")
	V.AutomaticEnv()

	err := V.ReadInConfig()
	if err != nil {
		log.Printf("unable to read config: %v", err)
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found ! ")
		}
		return nil, err
	}
	return V, nil
}

func getConfigPath(env string) string {
	if env == "docker" {
		return "config/config-docker"
	} else if env == "production" {
		return "config/config-production"
	} else {
		return "../config/config-development"
	}
}
