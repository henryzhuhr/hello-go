---
outline: [3,6]
date: 2024-10-10
---

## 包导入

### 基本导入
Go 中导入包可以使用 `import` 关键字。导入的包可以是标准库的包，也可以是第三方包，也可以是你自己写的包。

单行导入：
```go
import "fmt"
import "math"
```

多行导入：
```go
import (
    "fmt"
    "math"
)
```

### 别名导入
如果有同名冲突的包，可以使用别名来解决：
```go
import (
    "crypto/rand"
    mrand "math/rand" // 将名称替换为mrand避免冲突
)
```

如果导入的包名字很长，可以使用别名来简化：
```go
import hwtm "helloworldtestmodule"
```

### 点操作

在导入包时，可以使用 `.` 操作符来简化调用包中的函数。例如导入 `fmt` 包后，可以直接调用 `fmt.Println` 函数，而不需要写 `fmt.`。
```go
import . "fmt"
func main() {
    Println("Hello, World!")
}
```
> 但是不推荐使用这种方式，因为会使代码变得难以阅读，也容易产生冲突


### 空白导入

如果导入的包没有使用，编译器会报错。但是有时候我们只是想导入包，而不使用包中的函数（有可能执行一些初始化任务），可以使用 `_` 来代替包名，这样就不会报错了。
```go
import _ "github.com/go-sql-driver/mysql"
```

### 导入路径

Go 语言的包是通过导入路径来区分的。导入路径是唯一的，不同的包不能有相同的导入路径。导入路径可以是相对路径，也可以是绝对路径。绝对路径一般是指 `github.com/username/projectname` 这种形式。

### 包的初始化

Go 语言中包的初始化是自动的，当导入包时，会自动执行包中的 `init` 函数。`init` 函数没有参数，也没有返回值。`init` 函数在包中可以有多个，执行顺序是按照导入包的顺序执行的。

```go
package main
import "fmt"
func init() { fmt.Println("init function") }
func main() { fmt.Println("main function") }
// 输出：
// init function
// main function
```

- `init` 优先于 `main` 函数执行。
- `init` 函数在包中只会执行一次，即使包被导入多次，`init` 函数也只会执行一次。
- `init` 函数是不能被调用的，只能在包中自动执行。
- `init` 函数的执行顺序是按照导入包的顺序执行的，例如：`import "a"; import "b";`，则 `a` 包的 `init` 函数会先执行，然后是 `b` 包的 `init` 函数。
- 如果一个包中有多个 `init` 函数，执行顺序是按照代码中的顺序执行的。
- 如果一个包中有多个文件，每个文件中都有 `init` 函数，执行顺序是按照文件名的字典序执行的。
