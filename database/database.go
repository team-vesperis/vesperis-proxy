package database

import (
	"go.uber.org/zap"
)

func InitializeDatabase(log *zap.SugaredLogger) {
	logger = log
	logger.Info("Initializing databases...")

	initializeMysql(logger)
	initializeRedis(logger)
}
