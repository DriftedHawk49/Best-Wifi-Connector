package log

import (
	"fmt"
	"os"
	"strings"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	LOGFILENAME string = "bwcservice.log"
)

// if level is empty string then log level INFO is set, if DEBUG is sent then DEBUG is set.
func InitLogger(level string) *zap.Logger {
	logFile, err := os.CreateTemp("", LOGFILENAME)
	if err != nil {
		panic(fmt.Sprintf("unable to create log file, err : %s, exiting.", err.Error()))
	}

	logLevel := zap.NewAtomicLevelAt(zapcore.InfoLevel)
	if strings.EqualFold(level, "DEBUG") {
		logLevel = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	}

	stdout := zapcore.AddSync(os.Stdout)

	file := zapcore.AddSync(&lumberjack.Logger{
		Filename: logFile.Name(),
	})

	prodCfg := zap.NewProductionEncoderConfig()
	prodCfg.TimeKey = "timestamp"
	prodCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	devCfg := zap.NewDevelopmentEncoderConfig()
	devCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder

	cEncoder := zapcore.NewConsoleEncoder(devCfg)
	fEncoder := zapcore.NewJSONEncoder(prodCfg)

	core := zapcore.NewTee(
		zapcore.NewCore(cEncoder, stdout, logLevel),
		zapcore.NewCore(fEncoder, file, logLevel),
	)

	return zap.New(core)
}
