---
outline: # 展示的目录层级范围 deep=[3,6]
  - "3"
  - "6"
date: 2024-09-03
---

## 结构体 Struct

### Go 的结构体

结构体可以存储一组不同类型的数据，是一种复合类型。Go抛弃了类与继承，同时也抛弃了构造方法，刻意弱化了面向对象的功能，Go并非是一个传统OOP的语言，但是Go依旧有着OOP的影子，通过结构体和方法也可以模拟出一个类。下面是一个简单的结构体的例子。

### 结构体的定义

```go
type structName struct {
    // 结构体字段
    field1 type1 
    Field2 type2
    Field2 type2 `tag`
    // ...
}
```
- `structName`：结构体的名称
- `field/Field`：结构体的字段名
  - 首字母小写，私有字段，外部包不可以访问
  - 首字母大写，公有字段，外部包可以访问
- `type`：结构体的字段类型
- `` `tag` ``：字段标签，用反引号(`` ` ``)包括起来，可以在运行时通过反射的机制读取出来，常用于序列化和反序列化


例如定义一个结构体 `User`
```go
type User struct {
    id    int    // 首字母小写，私有字段，外部包不可以访问
    Name  string // 首字母大写，公有字段，外部包可以访问
    Phone string `json:"phone"` // 使用反引号定义 tag，可以在运行时通过反射的机制读取出来，常用于序列化和反序列化
    // 如果是小写 struct field author has json tag but is not exported (https://pkg.go.dev/golang.org/x/tools/go/analysis/passes/structtag)
}
```

空结构体也是一种结构体，它不包含任何字段，也不会占用内存，一般可以用于 `map` 的值类型，例如 `map[string]struct{}`，也可以作为通道的元素类型，例如 `chan struct{}`

```go
type Empty struct {}
fmt.Println(unsafe.Sizeof(Empty{})) // 计算结构体的大小
// 0
```

### 结构体实例化

Go 语言的结构体没有构造函数，我们可以自己实现。例如，下方的代码就实现了一个 `User` 的“构造函数”：

```go
func NewUser(id int, name string, phone string) *User {
    var newUser *User
    newUser = &User{id: id, Name: name, Phone: phone}
    // 也可以使用 `newUser:=&User{}` 语法糖定义
    return newUser
}
```

`newUser` 是一个 `*User` 类型，实例化 `User{id: id, Name: name, Phone: phone}` 的过程中是可以直接给私有字段赋值的，内部包可以访问，但是外部包就无法访问了

因为 `struct` 是值类型，如果结构体比较复杂的话，值拷贝性能开销会比较大，所以该构造函数返回的是结构体指针类型

因此通过该方法就可以实例化一个结构体
```go
func TestStruct(t *testing.T) {
    user := NewUser(0, "J", "123")
}
```

### 函数选项模式

Functional Options Pattern


声明一个 `UserOptions` 类型， 它接受一个 `*User` 类型的参数，它必须是指针，因为我们要在闭包中对 `User` 赋值。

```go
type UserOptions func(*User)
```


```go
// 如果不定义 UserOptions ，那就要写成
// func WithId(id int) func(*User) {
func WithId(id int) UserOptions {
    return func(u *User) {
        u.id = id
    }
}

// 另外两个类似，不展开
func WithName(name string) UserOptions { return func(u *User) { u.Name = name } }
func WithPhone(phone string) UserOptions { return func(u *User) { u.Phone = phone } }
```

```go
func NewUserWithOptions(options ...UserOptions) *User {
    // 优先应用 options
    user := &User{}
    for _, option := range options {
        option(user)
    }

    // 默认值处理
    if user.id < 0 {
        user.id = 0
    }
    // ...
    return user
}
```

随后初始化的时候，就可以根据需求进行

```go
func TestStruct(t *testing.T) {
    user1 := NewUserWithOptions(
        WithId(2),
        WithName("B"),
        WithPhone("345"),
    )
    user2 := NewUserWithOptions(
        WithName("B"),
    )
    fmt.Println("user1=", *user1)
    fmt.Println("user2=", *user2)
}
// user1= {2 B 345}
// user2= {0 B }
```

另外，也可以指定必须初始化某些参数，其他参数是可选的
```go
func NewUserWithNameWithOptions(name string, options ...UserOptions) *User {
    user := &User{Name: name}
    for _, option := range options {
        option(user)
    }
    return user
}
func TestStruct(t *testing.T) {
    user3 := NewUserWithNameWithOptions("D", WithPhone("000"))
    fmt.Println("user3=", *user3)
}
// user3= {0 D 000}
```


### 结构体的方法

`User` 的 `id` 字段是一个私有字段，因此外部无法直接修改，可以在实例化（如上）设置，但是随后的修改只能通过如下的方式

### 值接收者

定义两个方法，`GetId()` 获取字段 `id`，`TrySetId(id int)` 尝试修改字段 `id`

```go
func (u User) GetId() int { return u.id }
func (u User) TrySetId(id int) { u.id = id } // unused write to field id
```

```go
func TestStruct(t *testing.T) {
    user := NewUser(0, "J", "123")
    fmt.Println("user id =", user.GetId())
    user.TrySetId(10)
    fmt.Println("user id =", user.GetId())
}
// user id = 0
// user id = 0
```

### 指针接收者

```go
func (u *User) SetId(id int) { u.id = id }
```

`(u *User)` 是接收者，表示这个方法是 `User` 类型的方法

使用 `*` 作为 `User` 的接收者，表示这个方法是 `User` 的指针类型方法，只有指针类型的接收者才能修改结构体的值

```go
func TestStruct(t *testing.T) {
    user := NewUser(0, "J", "123")
    fmt.Println("user id =", user.GetId())
    user.SetId(20)
    fmt.Println("user id =", user.GetId())
}
// user id = 0
// user id = 20
```


### 结构体的标签


结构体标签是一种元编程的形式，结合反射可以做出很多奇妙的功能，格式如下

```go
`key1:"val1" key2:"val2"`
```

标签是一种键值对的形式，使用空格进行分隔。结构体标签的容错性很低，如果没能按照正确的格式书写结构体，那么将会导致无法正常读取，但是在编译时却不会有任何的报错，下方是一个使用示例。

```go
type Programmer struct {
    Name     string `json:"name"`
    Age      int `yaml:"age"`
    Job      string `toml:"job"`
    Language []string `properties:"language"`
}
```

结构体标签最广泛的应用就是在各种序列化格式中的别名定义，标签的使用需要结合反射才能完整发挥出其功能。

```go
func TestStruct(t *testing.T) {
    user := NewUser(0, "J", "123")
    fmt.Println("Get tag of user:", reflect.TypeOf(user.Phone), reflect.ValueOf(user.Phone))
}
// Get tag of user: string 123
```


### 组合

Go 语言中没有继承的概念，但是可以通过组合的方式实现类似继承的功能


例如有一个 `Person` 结构体，包含 `name` 和 `age` 两个字段，以及 `Eat()` 和 `Walk()` 两个方法

```go
type Person struct {
    name string
    age  int
}
func (p *Person) Eat() { fmt.Println("Person  Eat") }
func (p *Person) Walk() { fmt.Println("Person  Walk") }
```

`Student` 结构体组合了 `Person` 结构体，同时也拥有了 `Person` 的属性和方法，`Student` 结构体新增了 `school` 字段和 `Study()` 方法

```go
type Student struct {
    p      Person // Student 组合了 Person，Student 也拥有了 Person 的属性和方法
    school string // Student 自己的属性
}
// Eat Student 重写了 Eat 方法
func (s *Student) Eat() { fmt.Println("Student Eat") }
// Fly Student 新增了 Study 方法
func (s *Student) Study() { fmt.Println("Student Walk") }
```

测试输出如下
```go
func TestStructA(t *testing.T) {
    person := Person{"Tom", 20}
    person.Eat()
    person.Walk()

    student := Student{Person{"Jack", 19}, "MIT"}
    student.Eat()
    student.Study()
}
// Person  Eat
// Person  Walk
// Student Eat
// Student Walk
```


此外，Go 语言还支持匿名字段，匿名字段可以像组合一样使用，但是匿名字段不需要指定字段名，只需要指定字段的类型即可

```go
type Teacher struct {
    Person // 匿名字段，Teacher 继承了 Person 的属性和方法
    school string
}
```

初始化的时候，可以直接使用 `Person` 的字段和方法

```go
teacher := Teacher{
    Person: Person{"Tom", 20},
    school: "MIT",
}
```

### 结构体指针

对于结构体，不需要解引用就可以直接使用字段，例如

```go
p := &Person{"Tom", 20}
p.name = "Jerry"
p.age = 30
```

这是因为，在编译器会自动将 `p.name` 转换为 `(*p).name`，这种语法糖使得我们可以直接使用 `p.name` 来访问结构体的字段，而不需要每次都写 `(*p).name`

### 标签

Go 的结构体可以通过 `` `tag` `` 来定义字段的元信息（用反引号(`` ` ``)包括起来，多个字段间用空格分割

```go
`key1:"val1" key2:"val2"`
```
> 标签的需要严格按照格式书写，否则将无法正常读取，但是在编译时却不会有任何的报错

一个结构体的字段可以有多个标签，例如

```go
type Programmer struct {
    Name     string   `type:"string" json:"name"`
    Age      int      `type:"int" yaml:"age" `
    Language []string `properties:"language"`
}
```

标签可以在运行时通过反射的机制读取出来，常见的用途是序列化和反序列化，例如 JSON、YAML、TOML、Properties 等格式

```go
func TestTag(t *testing.T) {
    programmer := &Programmer{"Tom", 20, []string{"Go", "Python"}}
    fmt.Println("Programmer:")
    fmt.Println("  Tag   =", reflect.TypeOf(*programmer).Field(0).Tag)
    fmt.Println("  Value =", programmer.Name)
    fmt.Println("  Tag   =", reflect.TypeOf(*programmer).Field(1).Tag)
    fmt.Println("  Value =", programmer.Age)
}
// Programmer:
//   Tag   = type:"string" json:"name"
//   Value = Tom
//   Tag   = type:"int" yaml:"age"
//   Value = 20
```

### 结构体的内存对齐

结构体的内存对齐是指，结构体中的字段在内存中的存储顺序和对齐方式，Go 语言中的结构体内存对齐遵循以下规则：
- 结构体中每个字段的首地址是字段大小的整数倍
- 结构体的大小是结构体中最大字段大小的整数倍
- 结构体的大小是字段大小的整数倍

内存对齐是为了提高 CPU 的读写效率，因为 CPU 读取内存的时候是按照字节对齐的，如果结构体的字段没有对齐，那么 CPU 就需要多次读取内存，效率就会降低，而不同架构的 CPU 对齐方式可能不同，32 位 CPU 一般是 4 字节对齐，64 位 CPU 一般是 8 字节对齐

```go
type Num struct {
    A int64   // 8 byte
    B int32   // 4 byte
    C int16   // 2 byte
    D int8    // 1 byte
    E int32   // 4 byte
}	
```

内存布局如下，可以看出，`D` 后补了 1 byte，`E` 后补了 4 byte，以保证内存对齐，因此总共占用 24 byte

```
|       8       |       8       |       8       |
| | | | | | | | | | | | | | | | | | | | | | | | |
|       8       |   4   | 2 |1|1|   4   |   4   |
|       A       |   B   | C |D| |   E   |       |
|                    24 byte                    |
```

```go
type Num struct {
	A int8
	B int64
	C int8
}
```

内存模型如下，可以看出，`A` 后补了 7 byte，`C` 后补了 7 byte，以保证内存对齐，因此总共占用 24 byte，这是会产生浪费的
```
|       8       |       8       |       8       |
| | | | | | | | | | | | | | | | | | | | | | | | |
|1|             |       8       |1|             |
|       A       |       B       |       C       |
|                    24 byte                    |
```

如果将 `Num` 结构体的字段顺序调整一下，可以减少内存的浪费

```go
type Num struct {
    B int64
    A int8
    C int8
}
```
这样一来内存布局就变成如下，总共占用 16 byte
```
|       8       |       8       |
| | | | | | | | | | | | | | | | |
|       8       |1|1|           |
|       B       |A|C|           |
|            16 byte            |
```

> 结构体的内存对齐是编译器决定的，不同的编译器可能会有不同的内存对齐方式，可以通过 `unsafe.Alignof` 函数获取字段的对齐方式

```go
func TestAlign(t *testing.T) {
    num := Num{}
    fmt.Println("Alignof A =", unsafe.Alignof(num.A))
    fmt.Println("Alignof B =", unsafe.Alignof(num.B))
    fmt.Println("Alignof C =", unsafe.Alignof(num.C))
}
// Alignof A = 1
// Alignof B = 8
// Alignof C = 1
```

但是编码过程中不太需要关心内存对齐，编译器会自动处理，只有在需要优化内存的时候才需要关心内存对齐， [`go-tools`](https://github.com/dominikh/go-tools?tab=readme-ov-file) 和 [`betteralign`](https://github.com/dkorunic/betteralign) 等工具提供了结构体重新排列的功能，可以帮助我们优化内存对齐