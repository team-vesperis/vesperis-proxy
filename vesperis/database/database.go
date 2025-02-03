package database

import (
	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

func InitializeDatabase(log *zap.SugaredLogger) {
	logger = log
	logger.Info("Initializing databases...")

	initializeMysql()
	initializeRedis()
}

func CloseDatabase() {
	closeMySQL()
	closeRedis()
}
