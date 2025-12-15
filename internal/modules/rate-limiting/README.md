# Bounded Frequency Runner - 频率限制执行器

一个通用的、功能丰富的频率限制执行器，用于控制函数的执行频率，避免短时间内重复执行。

## 特性

- ✅ **灵活的时间控制**：支持自定义最小和最大执行间隔
- ✅ **多种触发方式**：手动触发、定时触发、阻塞式触发
- ✅ **错误处理和重试**：支持自动重试机制，带指数退避
- ✅ **优雅关闭**：基于 context 的优雅停止机制
- ✅ **并发安全**：使用互斥锁保护关键操作
- ✅ **可观测性**：内置执行统计和自定义日志支持
- ✅ **高度可配置**：通过 Options 模式灵活配置各项参数

## 快速开始

### 基本用法

```go
package main

import (
    "context"
    "fmt"
    "time"
    "your-project/internal/modules/rate-limiting/bounded_frequency"
)

func main() {
    // 创建执行器
    runner := bounded_frequency.NewBoundedFrequencyRunner(
        func(ctx context.Context) error {
            fmt.Println("执行业务逻辑")
            return nil
        },
    )

    ctx := context.Background()
    
    // 启动执行器
    go runner.Start(ctx)
    
    // 手动触发执行
    runner.Run()
    
    // 等待一段时间
    time.Sleep(5 * time.Second)
    
    // 停止执行器
    runner.Stop()
}
```

## 配置选项

### 可用选项

| 选项 | 说明 | 默认值 |
|-----|------|--------|
| `WithMinInterval(duration)` | 设置最小执行间隔 | 1秒 |
| `WithMaxInterval(duration)` | 设置最大执行间隔（定时触发） | 等于MinInterval |
| `WithBufferSize(size)` | 设置触发信号缓冲区大小 | 1 |
| `WithLogFunc(func)` | 设置自定义日志函数 | nil（不输出） |
| `WithAutoRun(bool)` | 是否启用定时自动执行 | true |
| `WithRetry(bool, maxRetries)` | 是否启用重试及最大重试次数 | false, 3 |

### 自定义配置示例

```go
runner := bounded_frequency.NewBoundedFrequencyRunner(
    func(ctx context.Context) error {
        // 你的业务逻辑
        return nil
    },
    bounded_frequency.WithMinInterval(500*time.Millisecond),  // 最小间隔 500ms
    bounded_frequency.WithMaxInterval(5*time.Second),         // 最多 5 秒自动触发一次
    bounded_frequency.WithBufferSize(10),                     // 缓冲区大小 10
    bounded_frequency.WithAutoRun(true),                      // 启用自动运行
    bounded_frequency.WithRetry(true, 3),                     // 失败时重试 3 次
    bounded_frequency.WithLogFunc(myLogFunc),                 // 自定义日志
)
```

## 使用场景

### 1. 配置热重载

避免频繁重新加载配置文件，确保配置更新有合理的间隔：

```go
reloadRunner := bounded_frequency.NewBoundedFrequencyRunner(
    func(ctx context.Context) error {
        // 重新加载配置
        return loadConfig()
    },
    bounded_frequency.WithMinInterval(2*time.Second),   // 最少 2 秒才重载一次
    bounded_frequency.WithMaxInterval(30*time.Second),  // 最多 30 秒自动检查一次
    bounded_frequency.WithBufferSize(1),
)

ctx := context.Background()
go reloadRunner.Start(ctx)

// 在文件监听器中触发
fileWatcher.OnChange(func() {
    reloadRunner.Run() // 不会阻塞
})
```

### 2. 缓存同步

控制缓存与数据源的同步频率：

```go
syncRunner := bounded_frequency.NewBoundedFrequencyRunner(
    func(ctx context.Context) error {
        return syncCacheWithDatabase()
    },
    bounded_frequency.WithMinInterval(1*time.Second),
    bounded_frequency.WithMaxInterval(10*time.Second),
    bounded_frequency.WithRetry(true, 5), // 失败时重试
)
```

### 3. 日志刷新

批量写入日志，减少磁盘 I/O：

```go
flushRunner := bounded_frequency.NewBoundedFrequencyRunner(
    func(ctx context.Context) error {
        return flushLogBuffer()
    },
    bounded_frequency.WithMinInterval(100*time.Millisecond),
    bounded_frequency.WithMaxInterval(1*time.Second),
    bounded_frequency.WithAutoRun(true),
)
```

### 4. 仅手动触发模式

禁用定时触发，完全由外部控制：

```go
manualRunner := bounded_frequency.NewBoundedFrequencyRunner(
    func(ctx context.Context) error {
        return processEvent()
    },
    bounded_frequency.WithMinInterval(500*time.Millisecond),
    bounded_frequency.WithAutoRun(false), // 禁用自动触发
)

// 只在特定事件发生时触发
onEvent(func() {
    manualRunner.Run()
})
```

## API 参考

### NewBoundedFrequencyRunner

```go
func NewBoundedFrequencyRunner(fn ExecuteFunc, options ...Option) *BoundedFrequencyRunner
```

创建新的频率限制执行器。

- `fn`: 要执行的函数，签名为 `func(ctx context.Context) error`
- `options`: 可选的配置选项

### Run

```go
func (r *BoundedFrequencyRunner) Run()
```

手动触发执行（非阻塞）。如果缓冲区已满，信号会被丢弃。

### RunBlocking

```go
func (r *BoundedFrequencyRunner) RunBlocking(ctx context.Context) error
```

手动触发执行（阻塞版本）。会阻塞直到信号被写入或上下文取消。

### Start

```go
func (r *BoundedFrequencyRunner) Start(ctx context.Context) error
```

启动执行器循环（阻塞）。通常在 goroutine 中运行。

### Stop

```go
func (r *BoundedFrequencyRunner) Stop()
```

停止执行器。可以安全地多次调用。

### GetStats

```go
func (r *BoundedFrequencyRunner) GetStats() (executeCount, failCount int64, lastExecute time.Time)
```

获取执行统计信息。

## 自定义日志

提供自己的日志函数来集成现有的日志系统：

```go
import "log"

logFunc := func(level bounded_frequency.LogLevel, msg string, args ...interface{}) {
    levelStr := "INFO"
    switch level {
    case bounded_frequency.LogLevelDebug:
        levelStr = "DEBUG"
    case bounded_frequency.LogLevelWarn:
        levelStr = "WARN"
    case bounded_frequency.LogLevelError:
        levelStr = "ERROR"
    }
    log.Printf("[%s] %s %v", levelStr, msg, args)
}

runner := bounded_frequency.NewBoundedFrequencyRunner(
    yourFunc,
    bounded_frequency.WithLogFunc(logFunc),
)
```

## 错误处理和重试

启用重试机制来处理临时性错误：

```go
runner := bounded_frequency.NewBoundedFrequencyRunner(
    func(ctx context.Context) error {
        // 可能失败的操作
        return riskyOperation()
    },
    bounded_frequency.WithRetry(true, 5), // 失败时重试，最多 5 次
    bounded_frequency.WithLogFunc(logFunc),
)
```

重试策略：

- 使用简单的指数退避：100ms * 重试次数
- 只有返回 error 时才会重试
- 达到最大重试次数后停止重试

## 并发安全

所有公共方法都是并发安全的，可以从多个 goroutine 安全调用：

```go
// 从多个地方触发
go runner.Run()
go runner.Run()
go runner.Run()

// 只有满足最小间隔的触发会实际执行
```

## 最佳实践

1. **选择合适的间隔**：根据业务特性选择 MinInterval 和 MaxInterval
2. **合理的缓冲区大小**：BufferSize=1 适合防抖场景，更大的值适合突发流量
3. **使用 context 控制生命周期**：通过 context 实现优雅关闭
4. **监控执行统计**：定期调用 GetStats() 了解执行情况
5. **自定义日志**：集成到你的日志系统中，便于问题排查

## 性能考虑

- 使用 channel 和 timer 实现，性能开销极小
- 互斥锁只在实际执行时持有，不影响触发操作
- 支持高并发场景下的频繁触发

## 对比原版

| 特性 | 原版 | 优化版 |
|-----|------|--------|
| 时间间隔配置 | 硬编码 1 秒 | 完全可配置 |
| 日志输出 | fmt.Println | 可自定义日志函数 |
| 错误处理 | 不支持 | 支持错误返回和重试 |
| 优雅关闭 | 不支持 | 基于 context 和 Stop() |
| 统计信息 | 无 | 提供详细统计 |
| 配置方式 | 构造函数参数 | Options 模式 |
| 触发方式 | 仅非阻塞 | 支持阻塞和非阻塞 |
| 自动运行 | 总是开启 | 可配置开关 |

## License

该代码遵循项目的整体 License。
