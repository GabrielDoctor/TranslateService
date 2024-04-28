package config

import (
	logger "backend/logging"

	"github.com/joho/godotenv"
)

func InitConfig() {
	godotenv.Load(".env")
	logger.InitLogger()
}
