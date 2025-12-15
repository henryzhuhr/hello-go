package log_test

import (
	"go.uber.org/zap"

	"github.com/henryzhuhr/hello-go/pkg/log"
)

// ExampleNewZapAdapter_basic 演示基本用法
func ExampleNewZapAdapter_basic() {
	zapLogger, _ := zap.NewDevelopment()
	logger := log.NewZapAdapter(zapLogger)

	// 简单的日志消息
	logger.Info("user logged in")

	// 格式化的日志消息
	logger.Infof("user %s logged in with id %d", "alice", 123)
}

// ExampleNewZapAdapter_allLevels 演示所有日志级别
func ExampleNewZapAdapter_allLevels() {
	zapLogger, _ := zap.NewDevelopment()
	logger := log.NewZapAdapter(zapLogger)

	// 不同级别的简单日志
	logger.Debug("debug message")
	logger.Info("info message")
	logger.Warn("warn message")
	logger.Error("error message")

	// 格式化版本
	logger.Debugf("debug: user %s", "alice")
	logger.Infof("info: processing %d items", 42)
	logger.Warnf("warn: high load %.2f%%", 85.5)
	logger.Errorf("error: failed to connect to %s", "database")
}

// ExampleNewDefaultZapAdapter 演示使用默认配置
func ExampleNewDefaultZapAdapter() {
	// 开发模式
	devLogger, _ := log.NewDefaultZapAdapter(true)
	devLogger.Debug("this is debug info")

	// 生产模式
	prodLogger, _ := log.NewDefaultZapAdapter(false)
	prodLogger.Info("application started")
}

// ExampleNewGlogAdapter_basic 演示 glog 适配器的基本用法
func ExampleNewGlogAdapter_basic() {
	logger := log.NewGlogAdapter()

	// 简单的日志消息
	logger.Info("user logged in")

	// 格式化的日志消息
	logger.Infof("user %s logged in with id %d", "alice", 123)
}

// ExampleNewGlogAdapter_allLevels 演示所有日志级别
func ExampleNewGlogAdapter_allLevels() {
	logger := log.NewGlogAdapter()

	// 不同级别的简单日志
	logger.Debug("debug message")
	logger.Info("info message")
	logger.Warn("warn message")
	logger.Error("error message")

	// 格式化版本
	logger.Debugf("debug: user %s", "alice")
	logger.Infof("info: processing %d items", 42)
	logger.Warnf("warn: high load %.2f%%", 85.5)
	logger.Errorf("error: failed to connect to %s", "database")
}

// ExampleNewGlogAdapter_comparison 演示 glog 和 zap 的对比使用
func ExampleNewGlogAdapter_comparison() {
	// glog 适配器 - 适合简单场景，配置通过命令行标志
	glogLogger := log.NewGlogAdapter()
	glogLogger.Info("glog message")

	// zap 适配器 - 适合需要高性能和复杂配置的场景
	zapLogger, _ := zap.NewProduction()
	zapAdapterLogger := log.NewZapAdapter(zapLogger)
	zapAdapterLogger.Info("zap message")

	// 两者实现相同的接口，可以无缝切换
	var logger log.Logger
	logger = glogLogger // 或者 zapAdapterLogger
	logger.Info("unified interface")
}
