package log

import (
	"flag"
	"testing"
)

// 初始化 glog 的 flag
func init() {
	flag.Set("logtostderr", "true")
	flag.Set("v", "2")
}

// TestNewGlogAdapter 测试创建 glog 适配器
func TestNewGlogAdapter(t *testing.T) {
	logger := NewGlogAdapter()
	if logger == nil {
		t.Fatal("NewGlogAdapter returned nil")
	}

	// 验证返回的是 glogAdapter 类型
	if _, ok := logger.(*glogAdapter); !ok {
		t.Error("NewGlogAdapter did not return *glogAdapter")
	}
}



// TestGlogAdapter_LogLevels 测试各种日志级别
// 注意：由于 glog 会输出到 stderr，这个测试主要验证不会 panic
func TestGlogAdapter_LogLevels(t *testing.T) {
	logger := NewGlogAdapter()

	// 捕获可能的 panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Logging caused panic: %v", r)
		}
	}()

	// 测试各个日志级别（除了 Fatal，因为它会退出程序）
	logger.Debug("debug message")
	logger.Info("info message")
	logger.Warn("warn message")
	logger.Error("error message")
}

// TestGlogAdapter_FormattedLogs 测试格式化日志方法
func TestGlogAdapter_FormattedLogs(t *testing.T) {
	logger := NewGlogAdapter()

	// 捕获可能的 panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Formatted logging caused panic: %v", r)
		}
	}()

	logger.Debugf("debug: user %s with id %d", "alice", 123)
	logger.Infof("info: processing %d items", 42)
	logger.Warnf("warn: high load %.2f%%", 85.5)
	logger.Errorf("error: failed to connect to %s", "database")
}

// BenchmarkGlogAdapter_Info 基准测试 Info 方法
func BenchmarkGlogAdapter_Info(b *testing.B) {
	logger := NewGlogAdapter()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("benchmark message")
	}
}

// BenchmarkGlogAdapter_Infof 基准测试 Infof 方法
func BenchmarkGlogAdapter_Infof(b *testing.B) {
	logger := NewGlogAdapter()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Infof("benchmark message: iteration %d", i)
	}
}
