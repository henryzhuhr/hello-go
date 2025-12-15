/*
定义了一个适配器，将 glog 日志记录器适配为通用日志接口 Logger 的实现。
glog 是 Google 的日志库，提供了分级日志和命令行标志控制功能。
*/
package log

import (
	"github.com/golang/glog"
)

// glogAdapter 是 glog 日志记录器的适配器实现
type glogAdapter struct{}

// GlogAdapterOption 是用于配置 glogAdapter 的函数选项
type GlogAdapterOption func(*glogAdapter)

// NewGlogAdapter 创建一个新的 glog 适配器
// opts: 可选的配置选项
func NewGlogAdapter(opts ...GlogAdapterOption) Logger {
	adapter := &glogAdapter{}

	for _, opt := range opts {
		opt(adapter)
	}

	return adapter
}

// Debug implements [Logger.Debug]
func (g *glogAdapter) Debug(msg string) {
	if glog.V(2) {
		glog.InfoDepth(1, msg)
	}
}

// Debugf implements [Logger.Debugf]
func (g *glogAdapter) Debugf(format string, args ...interface{}) {
	if glog.V(2) {
		glog.InfoDepthf(1, format, args...)
	}
}

// Info implements [Logger.Info]
func (g *glogAdapter) Info(msg string) {
	glog.InfoDepth(1, msg)
}

// Infof implements [Logger.Infof]
func (g *glogAdapter) Infof(format string, args ...interface{}) {
	glog.InfoDepthf(1, format, args...)
}

// Warn implements [Logger.Warn]
func (g *glogAdapter) Warn(msg string) {
	glog.WarningDepth(1, msg)
}

// Warnf implements [Logger.Warnf]
func (g *glogAdapter) Warnf(format string, args ...interface{}) {
	glog.WarningDepthf(1, format, args...)
}

// Error implements [Logger.Error]
func (g *glogAdapter) Error(msg string) {
	glog.ErrorDepth(1, msg)
}

// Errorf implements [Logger.Errorf]
func (g *glogAdapter) Errorf(format string, args ...interface{}) {
	glog.ErrorDepthf(1, format, args...)
}

// Fatal implements [Logger.Fatal]
func (g *glogAdapter) Fatal(msg string) {
	glog.FatalDepth(1, msg)
}

// Fatalf implements [Logger.Fatalf]
func (g *glogAdapter) Fatalf(format string, args ...interface{}) {
	glog.FatalDepthf(1, format, args...)
}
