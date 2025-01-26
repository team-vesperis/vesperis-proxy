package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var getLogger *zap.SugaredLogger

const logDir = "logs"
const maxLogFiles = 20

func CreateLogger() {
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		panic(fmt.Sprintf("Failed to create log directory: %v", err))
	}

	logFileName := fmt.Sprintf("proxy_%s.log", time.Now().Format("2006-01-02_15-04-05"))
	logFilePath := filepath.Join(logDir, logFileName)
	file, err := os.Create(logFilePath)
	if err != nil {
		panic(fmt.Sprintf("Failed to create log file: %v", err))
	}

	manageLogFiles()

	config := zap.NewProductionConfig()

	consoleEncoderConfig := zap.NewDevelopmentEncoderConfig()
	consoleEncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	consoleEncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(consoleEncoderConfig)

	jsonEncoderConfig := config.EncoderConfig
	jsonEncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.UnixDate)
	jsonEncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	jsonEncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	jsonEncoder := zapcore.NewJSONEncoder(jsonEncoderConfig)

	consoleCore := zapcore.NewCore(consoleEncoder, zapcore.AddSync(zapcore.Lock(os.Stdout)), zapcore.InfoLevel)
	fileCore := zapcore.NewCore(jsonEncoder, zapcore.AddSync(zapcore.Lock(file)), zapcore.InfoLevel)

	core := zapcore.NewTee(consoleCore, fileCore)

	logger := zap.New(core, zap.AddCaller())
	sugar := logger.Sugar()
	defer logger.Sync()

	getLogger = sugar
}

func GetLogger() *zap.SugaredLogger {
	return getLogger
}

func manageLogFiles() {
	files, err := os.ReadDir(logDir)
	if err != nil {
		panic(fmt.Sprintf("Failed to read log directory: %v", err))
	}

	if len(files) > maxLogFiles {
		fileInfos := make([]os.FileInfo, len(files))
		for i, file := range files {
			info, err := file.Info()
			if err != nil {
				panic(fmt.Sprintf("Failed to get file info: %v", err))
			}
			fileInfos[i] = info
		}

		sort.Slice(fileInfos, func(i, j int) bool {
			return fileInfos[i].ModTime().Before(fileInfos[j].ModTime())
		})

		for i := 0; i < len(fileInfos)-maxLogFiles; i++ {
			os.Remove(filepath.Join(logDir, fileInfos[i].Name()))
		}
	}
}
