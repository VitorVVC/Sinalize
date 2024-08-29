package services

import (
	"SinalizeB/constants"
	"SinalizeB/utils"
	"fmt"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var postgresInstance *gorm.DB

func Postgres() *gorm.DB {
	if postgresInstance == nil {
		connectionURL := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
			utils.EnvString(constants.PostgresHost),
			utils.EnvString(constants.PostgresUser),
			utils.EnvString(constants.PostgresPass),
			utils.EnvString(constants.PostgresName),
			utils.EnvInt(constants.PostgresPort),
			"require",
		)

		db, err := gorm.Open(postgres.Open(connectionURL))
		if err != nil {
			zap.L().Panic("failed to connect to postgres", zap.Error(err))
		}

		postgresInstance = db
	}

	return postgresInstance
}
