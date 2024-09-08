package go_struct

// Go 语言不是一种 “传统” 的面向对象编程语言：它里面没有类和继承的概念。
// 但是 Go 语言里有非常灵活的 接口 概念，通过它可以实现很多面向对象的特性
type User struct {
	id    int    // 首字母小写，私有字段，外部包不可以访问
	Name  string // 首字母大写，公有字段，外部包可以访问
	Phone string `json:"phone"` // 使用反引号定义 tag，可以在运行时通过反射的机制读取出来，常用于序列化和反序列化
	// 如果是小写 struct field author has json tag but is not exported (https://pkg.go.dev/golang.org/x/tools/go/analysis/passes/structtag)
}

// Go语言的结构体没有构造函数，我们可以自己实现。 例如，下方的代码就实现了一个person的构造函数。
// 因为struct是值类型，如果结构体比较复杂的话，值拷贝性能开销会比较大，所以该构造函数返回的是结构体指针类型。
func NewUser(id int, name string, phone string) *User {
	var newUser *User
	newUser = &User{id: id, Name: name, Phone: phone}
	return newUser
}

type UserOptions func(*User)

// 如果不定义 UserOptions ，那就要写成
// func WithId(id int) func(*User) {
func WithId(id int) UserOptions {
	return func(u *User) {
		u.id = id
	}
}

func WithName(name string) UserOptions { return func(u *User) { u.Name = name } }

func WithPhone(phone string) UserOptions { return func(u *User) { u.Phone = phone } }

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

func NewUserWithNameWithOptions(name string, options ...UserOptions) *User {
	user := &User{Name: name}
	for _, option := range options {
		option(user)
	}
	return user
}

func (u User) GetId() int      { return u.id }
func (u User) TrySetId(id int) { u.id = id } // unused write to field id

// (u *User) 是接收者，表示这个方法是 User 类型的方法
// 使用 * 作为 User 的接收者，表示这个方法是 User 的指针类型方法，只有指针类型的接收者才能修改结构体的值
func (u *User) SetId(id int) { u.id = id }
