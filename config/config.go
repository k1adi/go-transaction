package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type DBConfig struct {
	Host, Port, Name, User, Password, Driver string
}

type APIConfig struct {
	APIHost, APIPort string
}

type Config struct {
	APIConfig
	DBConfig
	FileConfig
	TokenConfig
}

type FileConfig struct {
	FilePath string
}

type TokenConfig struct {
	ApplicationName     string
	JwtSignatureKey     []byte
	JwtSigningMethod    *jwt.SigningMethodHMAC
	AccessTokenLifeTime time.Duration
}

func (c Config) ReadConfig() Config {
	c.DBConfig = DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Driver:   os.Getenv("DB_DRIVER"),
	}

	c.APIConfig = APIConfig{
		APIHost: os.Getenv("API_HOST"),
		APIPort: os.Getenv("API_PORT"),
	}

	c.FileConfig = FileConfig{
		FilePath: os.Getenv("FILE_PATH"),
	}

	appTokenExpire, err := strconv.Atoi(os.Getenv("APP_TOKEN_EXPIRE"))
	if err != nil {
		log.Printf("%v \n", err.Error())
	}

	accessTokenLifeTime := time.Duration(appTokenExpire) * time.Minute

	c.TokenConfig = TokenConfig{
		ApplicationName:     os.Getenv("APP_TOKEN_NAME"),
		JwtSignatureKey:     []byte(os.Getenv("APP_TOKEN_KEY")),
		JwtSigningMethod:    jwt.SigningMethodHS256,
		AccessTokenLifeTime: accessTokenLifeTime,
	}

	if c.DBConfig.Host == "" || c.DBConfig.Port == "" || c.DBConfig.Name == "" || c.DBConfig.User == "" || c.DBConfig.Password == "" || c.DBConfig.Driver == "" || c.APIConfig.APIHost == "" || c.APIConfig.APIPort == "" {
		log.Println("missing required enivronment variables")
	}

	return c
}

func NewConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	config := &Config{}
	return config.ReadConfig()
}
