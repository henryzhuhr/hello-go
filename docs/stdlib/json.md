---
outline: deep
---

## Json 解析库

JSON 作为一种数据格式，核心动作就是两个：**序列化，反序列化**。
- **序列化**，是把一个 **Go 对象**转化为 **JSON 格式的字符串（或字节序列）**
- **反序列化**，则相反，把 **JSON 格式**的数据转化成 **Go 对象**

> 对象是一个广义的概念，不单指结构体对象，包括 slice、map 类型数据也支持 JSON 的序列化。

### Go 语言内置的 Json 解析库

Go 语言内置的 Json 解析库是 `encoding/json`，它提供了一些函数和结构体，可以用来解析 Json 数据。


#### 读取和写入 Json 文件

```go
// 创建一个文件
file, err := os.Create("test.json")
// file, err := os.OpenFile("test.json", os.O_CREATE|os.O_RDWR, 0666)

// 创建一个json编码器
encoder := json.NewEncoder(file)

// 创建一个map
jsonData := map[string]interface{}{
    "name":  "zhangsan",
    "age":   18,
    "email": "xxx@xxx.com",
}

// 编码map数据到文件中
err = encoder.Encode(jsonData)
```

### 参考资料

- [后端 - Golang 操作 JSON 时容易踩的 7 个坑 - 程序员小屋 - SegmentFault 思否](https://segmentfault.com/a/1190000044876262)