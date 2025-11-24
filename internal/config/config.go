package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port    string
	DBHost  string
	DBPort  string
	DBUser  string
	DBPass  string
	DBName  string
	SSLMode string
}

var Cfg Config
var JwtToken []byte
var JwtExpire int

func Load() {
	_ = godotenv.Load()

	Cfg = Config{
		Port:    getEnv("PORT", "8080"),
		DBHost:  getEnv("DB_HOST", "localhost"),
		DBPort:  getEnv("DB_PORT", "5432"),
		DBUser:  getEnv("DB_USER", "user"),
		DBPass:  getEnv("DB_PASS", ""),
		DBName:  getEnv("DB_NAME", "crud"),
		SSLMode: getEnv("SSL_MODE", "disable"),
	}

	JwtToken = []byte(getEnv("JWT_TOKEN", ""))
	JwtExpire, _ = strconv.Atoi(getEnv("JWT_EXPIRE", "1"))
}

func getEnv(key string, def string) string {
	val := os.Getenv(key)
	if val == "" {
		return def
	}

	return val
}
