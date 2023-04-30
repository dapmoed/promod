package pm_logger

import "go.uber.org/zap"

type ZapLogger struct {
	logger *zap.Logger
}

func NewZapLogger() *ZapLogger {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	return &ZapLogger{
		logger: logger,
	}
}

func (z *ZapLogger) Info(msg string, fields ...interface{}) {
	zapFields := make([]zap.Field, 0)
	for _, f := range fields {
		if z, ok := f.(zap.Field); ok {
			zapFields = append(zapFields, z)
		}
	}
	z.logger.Info(msg, zapFields...)
}
func (z *ZapLogger) Error(msg string, fields ...interface{}) {
	zapFields := make([]zap.Field, 0)
	for _, f := range fields {
		if z, ok := f.(zap.Field); ok {
			zapFields = append(zapFields, z)
		}
	}
	z.logger.Error(msg, zapFields...)
}
