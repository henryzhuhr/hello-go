/*
定义了一个适配器，将 zap 日志记录器适配为通用日志接口 Logger 的实现，便于在代码中使用统一的日志接口，同时利用 zap 提供的日志功能。
*/
package log

import (
	"go.uber.org/zap"
)

// zapAdapter 是 zap 日志记录器的适配器实现
type zapAdapter struct {
	logger *zap.SugaredLogger
}

// ZapAdapterOption 是用于配置 zapAdapter 的函数选项
type ZapAdapterOption func(*zapAdapter)

// NewZapAdapter 创建一个新的 zap 适配器
// zapLogger: zap 的原始日志记录器
// opts: 可选的配置选项
func NewZapAdapter(zapLogger *zap.Logger, opts ...ZapAdapterOption) Logger {
	adapter := &zapAdapter{
		logger: zapLogger.Sugar(),
	}

	for _, opt := range opts {
		opt(adapter)
	}

	return adapter
}

// NewDefaultZapAdapter 创建一个使用默认配置的 zap 适配器
// development: 是否为开发模式，开发模式会输出更多调试信息
// opts: 可选的配置选项
func NewDefaultZapAdapter(development bool, opts ...ZapAdapterOption) (Logger, error) {
	var logger *zap.Logger
	var err error

	if development {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}

	if err != nil {
		return nil, err
	}

	return NewZapAdapter(logger, opts...), nil
}

// Debug implements [Logger.Debug]
func (z *zapAdapter) Debug(msg string) {
	z.logger.Debug(msg)
}

// Debugf implements [Logger.Debugf]
func (z *zapAdapter) Debugf(format string, args ...interface{}) {
	z.logger.Debugf(format, args...)
}

// Info implements [Logger.Info]
func (z *zapAdapter) Info(msg string) {
	z.logger.Info(msg)
}

// Infof implements [Logger.Infof]
func (z *zapAdapter) Infof(format string, args ...interface{}) {
	z.logger.Infof(format, args...)
}

// Warn implements [Logger.Warn]
func (z *zapAdapter) Warn(msg string) {
	z.logger.Warn(msg)
}

// Warnf implements [Logger.Warnf]
func (z *zapAdapter) Warnf(format string, args ...interface{}) {
	z.logger.Warnf(format, args...)
}

// Error implements [Logger.Error]
func (z *zapAdapter) Error(msg string) {
	z.logger.Error(msg)
}

// Errorf implements [Logger.Errorf]
func (z *zapAdapter) Errorf(format string, args ...interface{}) {
	z.logger.Errorf(format, args...)
}

// Fatal implements [Logger.Fatal]
func (z *zapAdapter) Fatal(msg string) {
	z.logger.Fatal(msg)
}

// Fatalf implements [Logger.Fatalf]
func (z *zapAdapter) Fatalf(format string, args ...interface{}) {
	z.logger.Fatalf(format, args...)
}
