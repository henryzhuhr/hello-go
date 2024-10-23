type=[]int---
outline: [3,6]
date: 2024-09-03
---
## 数据结构

### 数组 array

#### 数组的声明和初始化

定义一个数组，数组的长度是固定的，不能动态改变。
```go
var array1 [3]int        // 定义一个长度为3的数组
array1 = [3]int{1, 2, 3} // 初始化数组
// [1 2 3]
```

也可以在定义的时候初始化
```go
array2 := [3]int{}             // 定义数组 [0 0 0]
array3 := [3]int{1, 2, 3}      // 初始化数组 [1 2 3]
array4 := [...]int{1, 2, 3}    // 初始化数组，不指定长度 [1 2 3]
array5 := [5]int{1: 10, 3: 30} // 指定下标初始化 [0 10 0 30 0]
```

还可以通过 `new` 函数创建一个数组指针
```go
nums := new([5]int) // new创建数组指针
// *nums = [0 0 0 0 0]
```

上述的方式，都可以分配一片固定大小的内存


#### 数组的分割

数组的分割，可以通过 `:` 来实现
```go
array := [5]int{1, 2, 3, 4, 5}
s1 := array[1:3] // 下标范围 [1, 3) [2, 3]
s2 := array[:3]  // 下标范围 [0, 3) [0, 1, 2]
s3 := array[1:]  // 下标范围 [1, 5) [1, 2, 3, 4]
s4 := array[:]   // 下标范围 [0, 5) [0, 1, 2, 3, 4]
// s1 类型是：[]int 
```
切割后的数组，会变为切片 slice 类型

分割得到切片后，对切片的修改会影响到原数组，这是因为切片是对原数组的引用
```go
s1[0] = 10 
// array = [1 10 3 4 5]
```


### 切片 slice

#### 创建切片



数组一旦创建就不能更改其长度和类型，而切片就不同，切片可以按需自动增长和缩小，增长一般使用内置的 append 函数来实现，而缩小则是通过对切片再次进行切割来实现。

切片是对数组的引用，是一个动态的数组

创建切片的方式有两种，一种是使用 `make` 函数，另一种是使用切片字面量

```go
//          make([]类型, 长度, 容量)
var slice = make([]int, 5,    5)

// 切片的长度需要小于等于容量
s := make([]int, 3) // 创建一个长度为3的切片（容量和长度相同）
s := make([]int, 3, 5) // 创建一个长度为3，容量为5的切片
s := make([]int, 0, 0) // 创建一个空切片
```

声明并初始化切片时，可以指定所有的元素，也可以只初始化部分元素，此时需要指定要初始化的元素索引


```go
// 初始化全部元素
s := []int{1, 2, 3}
var s = []int{1, 2, 3}

// 初始化部分元素，索引为1、2、5的元素
s := []int{1:1, 2:6, 5:10}
// [0 1 6 0 0 10]
```

可以使用 `len(s)` 和 `cap(s)` 来获取切片的长度和容量

通过 `var nums []int` 这种方式声明的切片，默认值为nil，所以不会为其分配内存
```go
var slice1 []int        // 定义一个切片
// []int{} len=0 cap=0
```

空切片和 nil 切片是不同的，空切片是一个长度为0的切片，而 nil 切片是一个没有底层数组的切片，这两种情况赋值的时候会报错

空切片
```go
var slice = []int{}
t.Log(slice, len(slice), cap(slice), slice == nil)
// 给切片赋值
slice[0] = 1 // panic: runtime error: index out of range [0] with length 0
```

nil 切片
```go
var slice []int
t.Log(slice, len(slice), cap(slice), slice == nil)

// 给切片赋值
slice[0] = 1 // panic: runtime error: index out of range [0] with length 0
```

切片的底层实现仍然是数组，是一个指向数组的指针，所以切片是引用类型

#### 切片操作 Append

`append` 函数可以用来向切片中添加元素，函数签名如下
```go
func append(slice []Type, elems ...Type) []Type
```
可以使用的方法如下
```go
slice := []int{1}
elem1, elem2 := 2, 3
anotherSlice := []int{4, 5}
slice = append(slice, elem1, elem2)          // 添加元素, [1 2 3]
slice = append(slice, anotherSlice...)       // 添加切片, [1 2 3 4 5]
```

```go
slice := append([]byte("hello "), "world"...) // 添加字符串, "hello world"
```

如果切片的容量不够，会重新分配内存，将原来的元素复制到新的内存中

<!-- 在 golang1.18 版本更新之前网上大多数的文章都是这样描述slice的扩容策略的： 当原 slice 容量小于 1024 的时候，新 slice 容量变成原来的 2 倍；原 slice 容量超过 1024，新 slice 容量变成原来的1.25倍。 在1.18版本更新之后，slice的扩容策略变为了： 当原slice容量(oldcap)小于256的时候，新slice(newcap)容量为原来的2倍；原slice容量超过256，新slice容量newcap = oldcap+(oldcap+3*256)/4 -->


#### 切片的插入元素

切片的插入元素，可以通过 `append` 函数来实现

```go
slice := []int{2, 3}
// 在头部插入元素 使用 ... 运算符来辅助解构切片
slice = append([]int{0, 1}, slice...) // [0 1 2 3]
// 在尾部插入元素
slice = append(slice, 4, 5) // [0 1 2 3 4 5]
```

在中间插入元素比较麻烦，需要如下操作
```go
// 在中间插入元素
index, value := 2, 10
slice = append(slice[:index], append([]int{value}, slice[index:]...)...)
// [0 1 10 2 3 4 5]
```

一方面，需要通过将原来的数组切割成两个部分 `slice[:index]` 和 `slice[index:]`，然后将插入的元素插入到中间（例如插入后半部分的头部），即
```go
append([]int{value}, slice[index:]...)
```
然后在组合两个切片，得到最终的切片

<!-- ```go
var slice = []int{1, 2, 3, 4, 5} // [1 2 3 4 5]
// 解构 slice 元素 到一个新的空切片中，
// 因为空切片会基于新数组创建的,不会和 slice 共享底层数组
// 所以影响到 slice 的切片
var newSlice = append([]int{}, slice...)
``` -->

#### 切片的删除元素

```go
var slice = []int{1, 2, 3, 4, 5, 6}
// 从切片首部删除
slice = slice[1:]
​
// 从切片尾部删除
slice = slice[:len(slice) - 2]
​
// 从切片中间删除, 如从索引为i，删除2个元素(i+2)
slice = append(slice[:1], slice[3:]...)
```

#### 切片的复制

使用 `copy` 函数可以将一个切片的内容复制到另一个切片中，函数签名如下
```go
func copy(dst, src []Type) int
```
`copy` 函数会返回复制的元素个数，复制的元素个数是两个切片长度的最小值

```go
slice1 := []int{1, 2, 3}
slice2 := make([]int, 2)
copy(slice2, slice1)
// slice2 = [1 2]
```

#### 切片作为函数参数

切片作为函数参数时，是引用传递，所以在函数内部对切片的修改会影响到原切片

```go
func modify(slice []int) { slice[0] = 10 }
func main() {
    slice := []int{1, 2, 3}
    modify(slice)
    fmt.Println(slice) // [10 2 3]
}
```

如果不希望修改原切片，可以在函数内部创建一个新的切片，然后返回

```go
func modify(slice []int) []int {
    newSlice := append([]int{}, slice...)
    newSlice[0] = 10
    return newSlice
}
```

### 映射表 map

**映射表**数据结构实现通常有两种，**哈希表(hash table)**和**搜索树(search tree)**，前者是无序的，后者是有序的。

在 Go 中，map 的实现是基于**哈希桶(也是一种哈希表)**，所以也是无序的

#### 创建 map

map 定义为
```go
map[keyType]valueType{}
```

在Go中，map 的键类型必须是可比较的，比如 `string`, `int` 是可比较的，而 `[]int` 是不可比较的，也就无法作为map的键。

创建 map 的方式有两种，一种是使用 map 字面量，另一种是使用 `make` 函数

使用字面量
```go
mp := map[string]int{"a": 0, "b": 1}
```

使用 `make` 函数
```go
mp := make(map[string]int, 10) // 创建一个容量为10的 map
mp := make(map[string][]int, 10) // 创建一个容量为10的 map
```

`map` 是引用类型，零值或未初始化的map可以访问，但是无法存放元素，所以必须要为其分配内存。


#### 访问 map

访问 map 的方式和访问数组类似，通过键来访问值

```go
val = map["b"] // 1
val = map["c"] // 0
```

可以看到，如果访问不存在的键，会返回值类型的零值，这里是 `int` 类型的零值 `0`。而访问的时候，实际上会返回两个值，第二个值是一个布尔值，表示是否存在这个键

```go
val, exist := map["b"] // 1, true
val, exist := map["c"] // 0, false

if val, exist := map["b"]; exist {
    fmt.Println(val)
}
```

#### 修改 map

修改 map 的方式和访问 map 类似，通过键来修改值，如果键不存在，会添加一个新的键值对，也就是添加元素

```go
map["b"] = 10 // map[a:0 b:10]
map["c"] = 20 // map[a:0 b:10 c:20]
```

删除 map 中的元素，使用 `delete` 函数，函数签名如下
```go
func delete(m map[Type]Type1, key Type)
```
使用则如下
```go
delete(map, "b")
```

#### 遍历 map

遍历 map 使用 `range` 关键字，遍历 map 时，每次迭代会返回两个值，一个是键，一个是值
```go
for key, value := range map {
    fmt.Println(key, value)
}
```

#### 清空 map

Go 1.21 之前，没有提供清空 map 的方法，只能对每一个键进行删除操作

```go
for k, _ := range map {
    delete(mp, k)
}
```

Go 1.21 之后，可以使用 `clear` 来清空 map，函数签名如下
```go
func clear[T ~[]Type | ~map[Type]Type1](t T)
```
使用则如下
```go
clear(mp)
```

#### 使用 map 实现集合 Set 

Go 语言中没有提供集合 Set 的数据结构。由于 map 的键是唯一的，所以可以用 map 来实现一个简单的集合

```go
set := make(map[int]struct{})  // 创建一个集合
set[1] = struct{}{}            // 添加元素
set[3] = struct{}{}            // 添加元素
set[7] = struct{}{}            // 添加元素
// map[1:{} 3:{} 7:{}]
```

并且一个空结构体 `struct{}{}` 不占用内存，所以可以用来作为 map 的值，这样可以节省内存

#### map 的并发安全性

Go 语言中的 map 是非线程安全的，如果多个 goroutine 并发访问 map，可能会导致数据竞争，所以在并发环境下，需要使用 `sync.Map` 来保证 map 的并发安全性