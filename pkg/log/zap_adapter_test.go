package log

import (
	"bytes"
	"strings"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// TestNewZapAdapter 测试创建 zap 适配器
func TestNewZapAdapter(t *testing.T) {
	zapLogger, err := zap.NewDevelopment()
	if err != nil {
		t.Fatalf("Failed to create zap logger: %v", err)
	}

	logger := NewZapAdapter(zapLogger)
	if logger == nil {
		t.Fatal("NewZapAdapter returned nil")
	}

	// 验证返回的是 zapAdapter 类型
	if _, ok := logger.(*zapAdapter); !ok {
		t.Error("NewZapAdapter did not return *zapAdapter")
	}
}

// TestNewDefaultZapAdapter_Development 测试创建开发模式的默认适配器
func TestNewDefaultZapAdapter_Development(t *testing.T) {
	logger, err := NewDefaultZapAdapter(true)
	if err != nil {
		t.Fatalf("Failed to create development zap adapter: %v", err)
	}

	if logger == nil {
		t.Fatal("NewDefaultZapAdapter returned nil")
	}
}

// TestNewDefaultZapAdapter_Production 测试创建生产模式的默认适配器
func TestNewDefaultZapAdapter_Production(t *testing.T) {
	logger, err := NewDefaultZapAdapter(false)
	if err != nil {
		t.Fatalf("Failed to create production zap adapter: %v", err)
	}

	if logger == nil {
		t.Fatal("NewDefaultZapAdapter returned nil")
	}
}

// TestZapAdapter_LogLevels 测试所有日志级别
func TestZapAdapter_LogLevels(t *testing.T) {
	// 创建一个带有自定义输出的 zap logger 用于捕获日志
	var buf bytes.Buffer
	encoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		EncodeLevel: zapcore.LowercaseLevelEncoder,
	})
	core := zapcore.NewCore(encoder, zapcore.AddSync(&buf), zapcore.DebugLevel)
	zapLogger := zap.New(core)

	logger := NewZapAdapter(zapLogger)

	tests := []struct {
		name     string
		logFunc  func(string)
		message  string
		expected string
	}{
		{
			name:     "Debug",
			logFunc:  logger.Debug,
			message:  "debug message",
			expected: "debug",
		},
		{
			name:     "Info",
			logFunc:  logger.Info,
			message:  "info message",
			expected: "info",
		},
		{
			name:     "Warn",
			logFunc:  logger.Warn,
			message:  "warn message",
			expected: "warn",
		},
		{
			name:     "Error",
			logFunc:  logger.Error,
			message:  "error message",
			expected: "error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf.Reset()
			tt.logFunc(tt.message)

			output := buf.String()
			if !strings.Contains(output, tt.expected) {
				t.Errorf("Expected log level %s, got: %s", tt.expected, output)
			}
			if !strings.Contains(output, tt.message) {
				t.Errorf("Expected message %s, got: %s", tt.message, output)
			}
		})
	}
}

// TestZapAdapter_FormattedLogs 测试格式化日志方法
func TestZapAdapter_FormattedLogs(t *testing.T) {
	var buf bytes.Buffer
	encoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		EncodeLevel: zapcore.LowercaseLevelEncoder,
	})
	core := zapcore.NewCore(encoder, zapcore.AddSync(&buf), zapcore.DebugLevel)
	zapLogger := zap.New(core)

	logger := NewZapAdapter(zapLogger)

	tests := []struct {
		name     string
		logFunc  func(string, ...interface{})
		format   string
		args     []interface{}
		expected string
	}{
		{
			name:     "Debugf",
			logFunc:  logger.Debugf,
			format:   "user %s logged in with id %d",
			args:     []interface{}{"alice", 123},
			expected: "user alice logged in with id 123",
		},
		{
			name:     "Infof",
			logFunc:  logger.Infof,
			format:   "processing %d items",
			args:     []interface{}{42},
			expected: "processing 42 items",
		},
		{
			name:     "Warnf",
			logFunc:  logger.Warnf,
			format:   "high load: %.1f%%",
			args:     []interface{}{85.5},
			expected: "high load: 85.5%",
		},
		{
			name:     "Errorf",
			logFunc:  logger.Errorf,
			format:   "failed to connect to %s",
			args:     []interface{}{"database"},
			expected: "failed to connect to database",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf.Reset()
			tt.logFunc(tt.format, tt.args...)

			output := buf.String()
			if !strings.Contains(output, tt.expected) {
				t.Errorf("Expected message '%s', got: %s", tt.expected, output)
			}
		})
	}
}
