package log

import (
	"context"
	"testing"
)

// TestNewContext 测试将 Logger 存入 context
func TestNewContext(t *testing.T) {
	// 创建一个 mock logger（这里用 nil 模拟，实际使用时会是真实的 logger）
	var mockLogger Logger = &mockLoggerImpl{}

	ctx := context.Background()
	newCtx := NewContext(ctx, mockLogger)

	// 验证 context 不为 nil
	if newCtx == nil {
		t.Fatal("NewContext returned nil context")
	}

	// 验证可以从 context 中取出 logger
	retrieved := FromContext(newCtx)
	if retrieved == nil {
		t.Fatal("FromContext returned nil logger")
	}

	if retrieved != mockLogger {
		t.Error("Retrieved logger is not the same as the original logger")
	}
}

// TestFromContext_WithLogger 测试从包含 Logger 的 context 中获取
func TestFromContext_WithLogger(t *testing.T) {
	mockLogger := &mockLoggerImpl{}
	ctx := NewContext(context.Background(), mockLogger)

	retrieved := FromContext(ctx)
	if retrieved == nil {
		t.Fatal("FromContext returned nil for context with logger")
	}

	if retrieved != mockLogger {
		t.Error("Retrieved logger does not match the original logger")
	}
}

// TestFromContext_WithoutLogger 测试从不包含 Logger 的 context 中获取
func TestFromContext_WithoutLogger(t *testing.T) {
	ctx := context.Background()

	retrieved := FromContext(ctx)
	if retrieved != nil {
		t.Error("FromContext should return nil for context without logger")
	}
}

// TestFromContext_WithWrongType 测试 context 中存储了错误类型的值
func TestFromContext_WithWrongType(t *testing.T) {
	ctx := context.WithValue(context.Background(), loggerKey, "not a logger")

	retrieved := FromContext(ctx)
	if retrieved != nil {
		t.Error("FromContext should return nil when context value is not a Logger")
	}
}

// TestContextChaining 测试 context 的链式传递
func TestContextChaining(t *testing.T) {
	logger1 := &mockLoggerImpl{name: "logger1"}
	logger2 := &mockLoggerImpl{name: "logger2"}

	ctx1 := NewContext(context.Background(), logger1)
	ctx2 := NewContext(ctx1, logger2)

	// ctx2 应该覆盖 ctx1 的 logger
	retrieved := FromContext(ctx2)
	if retrieved == nil {
		t.Fatal("FromContext returned nil")
	}

	if mock, ok := retrieved.(*mockLoggerImpl); ok {
		if mock.name != "logger2" {
			t.Errorf("Expected logger2, got %s", mock.name)
		}
	} else {
		t.Error("Retrieved logger is not mockLoggerImpl")
	}

	// ctx1 应该仍然保持原来的 logger
	retrieved1 := FromContext(ctx1)
	if retrieved1 == nil {
		t.Fatal("FromContext returned nil for ctx1")
	}

	if mock, ok := retrieved1.(*mockLoggerImpl); ok {
		if mock.name != "logger1" {
			t.Errorf("Expected logger1, got %s", mock.name)
		}
	} else {
		t.Error("Retrieved logger from ctx1 is not mockLoggerImpl")
	}
}

// mockLoggerImpl 是一个用于测试的 Logger 实现
type mockLoggerImpl struct {
	name string
}

// Debug implements [Logger.Debug]
func (m *mockLoggerImpl) Debug(msg string) {}

// Debugf implements [Logger.Debugf]
func (m *mockLoggerImpl) Debugf(format string, args ...interface{}) {}

// Info implements [Logger.Info]
func (m *mockLoggerImpl) Info(msg string) {}

// Infof implements [Logger.Infof]
func (m *mockLoggerImpl) Infof(format string, args ...interface{}) {}

// Warn implements [Logger.Warn]
func (m *mockLoggerImpl) Warn(msg string) {}

// Warnf implements [Logger.Warnf]
func (m *mockLoggerImpl) Warnf(format string, args ...interface{}) {}

// Error implements [Logger.Error]
func (m *mockLoggerImpl) Error(msg string) {}

// Errorf implements [Logger.Errorf]
func (m *mockLoggerImpl) Errorf(format string, args ...interface{}) {}

// Fatal implements [Logger.Fatal]
func (m *mockLoggerImpl) Fatal(msg string) {}

// Fatalf implements [Logger.Fatalf]
func (m *mockLoggerImpl) Fatalf(format string, args ...interface{}) {}
