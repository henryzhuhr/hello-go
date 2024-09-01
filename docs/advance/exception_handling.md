# 异常处理

## Go 的异常类型

Go 语言中的异常处理机制与其他语言有很大的不同。Go 语言中的异常处理机制是通过 `panic` 和 `recover` 来实现的。

Go 的异常主要有以下几种类型：
- `error`: 部分流程出错，需要处理
- `panic`: 很严重的问题，程序应该在处理完问题后立即退出
- `fatal`: 非常致命的问题，程序应该立即退出

> 准确的来说，Go 并没有异常，更多的是通过错误来体现，同样的，Go 中也并没有 `try-catch-finally` 这种语句，Go 创始人希望能够将错误可控，他们不希望干什么事情都需要嵌套一堆 `try-catch` ，所以大多数情况会将其作为函数的返回值来返回

这里有两篇Go团队关于错误处理的文章，感兴趣可以看看

## 参考资料

[1][2] 是Go团队关于错误处理的文章

- [1] [Errors are values - Rob Pike](https://go.dev/blog/errors-are-values)
- [2] [Error handling and Go - Andrew Gerrand](https://go.dev/blog/error-handling-and-go)
- [3] [Golang深入浅出之-Go语言 defer、panic、recover：异常处理机制-腾讯云开发者社区-腾讯云](https://cloud.tencent.com/developer/article/2412169)
- [4] [一套优雅的 Go 错误问题解决方案-腾讯云开发者社区-腾讯云](https://cloud.tencent.com/developer/article/1885094)