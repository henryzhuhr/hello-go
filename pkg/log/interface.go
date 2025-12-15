/*
通用的日志接口，核心目标是解耦日志实现与业务逻辑，便于在不同环境（开发、测试、生产）或不同日志库（如 logrus、zap、标准库 log 等）之间切换

- 避免全局日志变量：通过构造函数或依赖注入传递 Logger，便于单元测试的依赖注入

- 上下文支持：可扩展接口支持 LoggerFromContext(ctx)，便于在分布式系统或复杂调用链中传递请求级上下文信息，并将其自动嵌入日志中，从而实现 可追踪、可关联、结构化 的日志记录
*/
package log

import "context"

// Logger 定义了通用的日志接口
type Logger interface {
	// Debug 记录调试级别的日志信息
	Debug(msg string)

	// Debugf 记录格式化的调试级别日志
	Debugf(format string, args ...interface{})

	// Info 记录信息级别的日志
	Info(msg string)

	// Infof 记录格式化的信息级别日志
	Infof(format string, args ...interface{})

	// Warn 记录警告级别的日志
	Warn(msg string)

	// Warnf 记录格式化的警告级别日志
	Warnf(format string, args ...interface{})

	// Error 记录错误级别的日志
	Error(msg string)

	// Errorf 记录格式化的错误级别日志
	Errorf(format string, args ...interface{})

	// Fatal 记录致命错误级别的日志，之后会终止程序
	Fatal(msg string)

	// Fatalf 记录格式化的致命错误级别日志，之后会终止程序
	Fatalf(format string, args ...interface{})
}

// contextKey 用于在 context 中存储和获取 Logger 的键类型
type contextKey struct{}

var loggerKey = contextKey{}

// NewContext 将 Logger 存入 context 中
func NewContext(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

// FromContext 从 context 中获取 Logger，如果 context 中没有 Logger，返回 nil
func FromContext(ctx context.Context) Logger {
	if logger, ok := ctx.Value(loggerKey).(Logger); ok {
		return logger
	}
	return nil
}
