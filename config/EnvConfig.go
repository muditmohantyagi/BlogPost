package config

import (
	"log"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func GetEnvWithKey(key string, defaultValue string) string {
	keyVal, found := syscall.Getenv(key)
	if !found {
		syscall.Setenv(key, defaultValue)
		return defaultValue
	}
	return keyVal
}

func loadEnv() {
	err := godotenv.Overload(".env")
	if err != nil {
		log.Fatal("Error loading .env file ", err)
		return
	}
}

func init() {
	loadEnv()
	if GetEnvWithKey("APP_ENVIRONMENT", "dev") == "pro" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
}
