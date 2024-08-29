package utils

import (
	"SinalizeB/constants"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"log"
	"os"
	"strconv"
)

var Defaults = map[string]interface{}{
	constants.ApiPort:      "your_port",
	constants.PostgresHost: "your_host",
	constants.PostgresPort: "5432",
	constants.PostgresUser: "your_user",
	constants.PostgresPass: "your_pass",
	constants.PostgresName: "your_name",
	constants.DatabaseDSN:  "your_database",
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("error to load .env: %v", err)
	}

	for key := range Defaults {
		if value := os.Getenv(key); value != "" {
			Defaults[key] = value
		}
	}
}

func EnvString(key string) string {
	valueInterface, ok := Defaults[key]
	if !ok {
		zap.L().Fatal("missing env", zap.Error(ErrMissingEnv), zap.String("env", key))
	}

	value, ok := valueInterface.(string)
	if !ok {
		zap.L().Fatal("wrong type", zap.Error(ErrWrongEnvType), zap.String("env", key))
	}

	return value
}

func EnvInt(key string) int64 {
	valueStr := os.Getenv(key)
	value, err := strconv.ParseInt(valueStr, 10, 64)
	if valueStr == "" || err != nil {
		if err != nil && valueStr != "" {
			zap.L().Panic("failed to convert integer from env", zap.Error(err))
		}

		valueInterface, ok := Defaults[key]
		if !ok {
			zap.L().Fatal("missing env", zap.Error(ErrMissingEnv), zap.String("env", key))
		}

		valueInt, ok := valueInterface.(int)
		if !ok {
			zap.L().Fatal("wrong type", zap.Error(ErrWrongEnvType), zap.String("env", key))
		}
		value = int64(valueInt)
	}

	//zap.L().Debug("returned env", zap.String("key", key), zap.Int64("value", value))
	return value
}
