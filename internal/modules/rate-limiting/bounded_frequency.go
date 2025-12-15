package bounded_frequency

import (
	"context"
	"errors"
	"sync"
	"time"
)

// LogLevel 日志级别
type LogLevel int

const (
	LogLevelDebug LogLevel = iota
	LogLevelInfo
	LogLevelWarn
	LogLevelError
)

// LogFunc 日志函数类型
type LogFunc func(level LogLevel, msg string, args ...interface{})

// Options 配置选项
type Options struct {
	// MinInterval 最小执行间隔，默认 1 秒
	MinInterval time.Duration
	// MaxInterval 最大执行间隔（定时器触发），默认等于 MinInterval
	MaxInterval time.Duration
	// BufferSize 触发信号缓冲区大小，默认 1
	BufferSize int
	// LogFunc 自定义日志函数，默认为 nil（不输出日志）
	LogFunc LogFunc
	// EnableAutoRun 是否启用定时自动执行，默认 true
	EnableAutoRun bool
	// RetryOnError 执行失败时是否重试，默认 false
	RetryOnError bool
	// MaxRetries 最大重试次数，默认 3
	MaxRetries int
}

// Option 配置函数
type Option func(*Options)

// WithMinInterval 设置最小执行间隔
func WithMinInterval(interval time.Duration) Option {
	return func(o *Options) {
		o.MinInterval = interval
	}
}

// WithMaxInterval 设置最大执行间隔（定时触发）
func WithMaxInterval(interval time.Duration) Option {
	return func(o *Options) {
		o.MaxInterval = interval
	}
}

// WithBufferSize 设置触发信号缓冲区大小
func WithBufferSize(size int) Option {
	return func(o *Options) {
		o.BufferSize = size
	}
}

// WithLogFunc 设置自定义日志函数
func WithLogFunc(logFunc LogFunc) Option {
	return func(o *Options) {
		o.LogFunc = logFunc
	}
}

// WithAutoRun 设置是否启用定时自动执行
func WithAutoRun(enable bool) Option {
	return func(o *Options) {
		o.EnableAutoRun = enable
	}
}

// WithRetry 设置是否在执行失败时重试
func WithRetry(enable bool, maxRetries int) Option {
	return func(o *Options) {
		o.RetryOnError = enable
		o.MaxRetries = maxRetries
	}
}

// defaultOptions 默认配置
func defaultOptions() Options {
	return Options{
		MinInterval:   time.Second,
		MaxInterval:   time.Second,
		BufferSize:    1,
		LogFunc:       nil,
		EnableAutoRun: true,
		RetryOnError:  false,
		MaxRetries:    3,
	}
}

// ExecuteFunc 执行函数类型，返回 error 以支持错误处理
type ExecuteFunc func(ctx context.Context) error

// BoundedFrequencyRunner 频率限制执行器
type BoundedFrequencyRunner struct {
	mu sync.Mutex

	// 主动触发信号通道
	trigger chan struct{}

	// 定时器
	timer *time.Timer

	// 执行函数
	fn ExecuteFunc

	// 配置选项
	opts Options

	// 停止信号
	stopCh chan struct{}
	stopOnce sync.Once

	// 执行统计
	executeCount int64
	failCount    int64
	lastExecute  time.Time
}

// NewBoundedFrequencyRunner 创建新的频率限制执行器
func NewBoundedFrequencyRunner(fn ExecuteFunc, options ...Option) *BoundedFrequencyRunner {
	opts := defaultOptions()
	for _, opt := range options {
		opt(&opts)
	}

	// 如果 MaxInterval 未设置，则使用 MinInterval
	if opts.MaxInterval == 0 || opts.MaxInterval < opts.MinInterval {
		opts.MaxInterval = opts.MinInterval
	}

	runner := &BoundedFrequencyRunner{
		trigger: make(chan struct{}, opts.BufferSize),
		fn:      fn,
		opts:    opts,
		stopCh:  make(chan struct{}),
		timer:   time.NewTimer(opts.MaxInterval),
	}

	// 如果不启用自动运行，停止定时器
	if !opts.EnableAutoRun {
		runner.timer.Stop()
	}

	return runner
}

// Run 手动触发执行
// 如果缓冲区已满，根据 BufferSize 决定是否阻塞或丢弃
func (r *BoundedFrequencyRunner) Run() {
	select {
	case r.trigger <- struct{}{}:
		r.log(LogLevelDebug, "触发信号已写入队列")
	default:
		r.log(LogLevelWarn, "触发队列已满，信号被丢弃")
	}
}

// RunBlocking 手动触发执行（阻塞版本）
// 会阻塞直到信号被写入
func (r *BoundedFrequencyRunner) RunBlocking(ctx context.Context) error {
	select {
	case r.trigger <- struct{}{}:
		r.log(LogLevelDebug, "触发信号已写入队列")
		return nil
	case <-ctx.Done():
		return ctx.Err()
	case <-r.stopCh:
		return errors.New("runner已停止")
	}
}

// Start 启动执行器循环（阻塞）
func (r *BoundedFrequencyRunner) Start(ctx context.Context) error {
	if r.opts.EnableAutoRun {
		r.timer.Reset(r.opts.MaxInterval)
	}

	r.log(LogLevelInfo, "执行器已启动", "minInterval", r.opts.MinInterval, "maxInterval", r.opts.MaxInterval)

	for {
		select {
		case <-ctx.Done():
			r.log(LogLevelInfo, "收到上下文取消信号，执行器停止")
			return ctx.Err()
		case <-r.stopCh:
			r.log(LogLevelInfo, "执行器已停止")
			return nil
		case <-r.trigger:
			r.log(LogLevelDebug, "收到手动触发信号")
			r.tryExecute(ctx, "manual")
		case <-r.timer.C:
			if r.opts.EnableAutoRun {
				r.log(LogLevelDebug, "定时器触发")
				r.tryExecute(ctx, "timer")
			}
		}
	}
}

// Stop 停止执行器
func (r *BoundedFrequencyRunner) Stop() {
	r.stopOnce.Do(func() {
		close(r.stopCh)
		r.timer.Stop()
		r.log(LogLevelInfo, "执行器停止信号已发送")
	})
}

// tryExecute 尝试执行函数
func (r *BoundedFrequencyRunner) tryExecute(ctx context.Context, source string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// 检查是否满足最小间隔要求
	if !r.lastExecute.IsZero() {
		elapsed := time.Since(r.lastExecute)
		if elapsed < r.opts.MinInterval {
			remaining := r.opts.MinInterval - elapsed
			r.log(LogLevelDebug, "未达到最小执行间隔", "remaining", remaining, "source", source)
			// 重置定时器以确保下次执行
			if r.opts.EnableAutoRun {
				r.timer.Reset(remaining)
			}
			return
		}
	}

	// 执行函数
	r.executeCount++
	r.lastExecute = time.Now()

	var err error
	retries := 0
	maxRetries := 1 // 默认执行一次

	if r.opts.RetryOnError {
		maxRetries = r.opts.MaxRetries + 1
	}

	for retries < maxRetries {
		err = r.fn(ctx)
		if err == nil {
			r.log(LogLevelInfo, "函数执行成功", "source", source, "executeCount", r.executeCount)
			break
		}

		retries++
		r.failCount++

		if retries < maxRetries {
			r.log(LogLevelWarn, "函数执行失败，准备重试", "error", err, "retry", retries, "maxRetries", r.opts.MaxRetries)
			// 简单的退避策略
			time.Sleep(time.Millisecond * 100 * time.Duration(retries))
		} else {
			r.log(LogLevelError, "函数执行失败", "error", err, "source", source, "retries", retries-1)
		}
	}

	// 重置定时器
	if r.opts.EnableAutoRun {
		r.timer.Reset(r.opts.MaxInterval)
	}
}

// GetStats 获取执行统计信息
func (r *BoundedFrequencyRunner) GetStats() (executeCount, failCount int64, lastExecute time.Time) {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.executeCount, r.failCount, r.lastExecute
}

// log 内部日志函数
func (r *BoundedFrequencyRunner) log(level LogLevel, msg string, args ...interface{}) {
	if r.opts.LogFunc != nil {
		r.opts.LogFunc(level, msg, args...)
	}
}

