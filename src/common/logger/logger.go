package logger

import (
	"os"

	"github.com/google/logger"
)

func InitLogger() {
	logger.Init("DefaultLogger", false, true, os.Stderr)
}
