package log

import "go.uber.org/zap"

var Logger *zap.Logger

func InitLogger() (err error) {
	Logger, err = zap.NewProduction()
	if err != nil {
		return err
	}
	return nil
}
