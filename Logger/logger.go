package Logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var zapLogger *zap.Logger

/*
PrintErrorLogLevel1 : 위험 수준 낮음
PrintErrorLogLevel2 : 위험 수준 중간
PrintErrorLogLevel3 : 위험 수준 높음
PrintErrorLogLevel4 : 위험 수준 매우 높음
*/

func init() {
	var err error

	config := zap.NewProductionConfig()
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.StacktraceKey = ""
	config.EncoderConfig = encoderConfig

	zapLogger, _ = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

func PrintErrorLogLevel1(err error) {
	zapLogger.Log(6, err.Error())
}

func PrintErrorLogLevel2(err error) {
	zapLogger.Log(7, err.Error())
}

func PrintErrorLogLevel3(err error) {
	zapLogger.Log(8, err.Error())
}

func PrintErrorLogLevel4(err error) {
	zapLogger.Log(9, err.Error())
}
