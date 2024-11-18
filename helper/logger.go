package helper

import (
	"context"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func initLogger(ctx context.Context) {
	// 创建一个日志记录器
	Logger = logrus.New()

	// 设置日志格式为 JSON
	Logger.SetFormatter(&logrus.JSONFormatter{})

	// 设置日志级别
	Logger.SetLevel(logrus.DebugLevel)

	// 记录一条信息日志
	Logger.Info("This is an info message.")

	// 记录一条带有字段的日志
	Logger.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	// 记录一条错误日志
	Logger.Error("This is an error message.")
}
