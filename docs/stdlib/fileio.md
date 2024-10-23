

# 文件处理

文件处理的标准库油：
- `os`，负责OS文件系统交互的具体实现
- `io`，读写IO的抽象层
- `fs`，文件系统的抽象层

## 文件

### 打开文件

使用 `os.OpenFile` 函数打开文件 ，返回一个 `os.File` 对象，函数签名如下：
```go
func OpenFile(name string, flag int, perm os.FileMode) (*File, error)
```

`name` 是文件路径，`flag` 是打开文件的模式，`perm` 是文件权限。


`src/os/file.go` 文件中定义了 `flag` 的取值：
```go
// Flags to OpenFile wrapping those of the underlying system. Not all
// flags may be implemented on a given system.
const (
	// 必须指定 O_RDONLY, O_WRONLY, O_RDWR 之一
	O_RDONLY int = syscall.O_RDONLY // read-only  只读
	O_WRONLY int = syscall.O_WRONLY // write-only 只写
	O_RDWR   int = syscall.O_RDWR   // read-write 读写
    // 其他的值可以用于控制行为
	O_APPEND int = syscall.O_APPEND // 附加写模式 (在文件末尾写)
	O_CREATE int = syscall.O_CREAT  // 如果不存在，则创建一个新文件。
	O_EXCL   int = syscall.O_EXCL   // 与 O_CREATE 一起使用时，文件必须不存在。
	O_SYNC   int = syscall.O_SYNC   // 以同步IO的方式打开文件
	O_TRUNC  int = syscall.O_TRUNC  // 当打开的时候截断可写的文件
)
```
`flag` 可以是 `O_RDONLY`、`O_WRONLY`、`O_RDWR` 之一，也可以和其他的标志位进行或运算，例如 `O_CREATE|O_WRONLY` 表示如果文件不存在则创建一个新文件并以只写模式打开。

`perm` 是文件权限，`src/os/types.go` 文件中定义了 `os.FileMode` 类型：
```go
// The defined file mode bits are the most significant bits of the [FileMode].
// The nine least-significant bits are the standard Unix rwxrwxrwx permissions.
// The values of these bits should be considered part of the public API and
// may be used in wire protocols or disk representations: they must not be
// changed, although new bits might be added.
const (
	// The single letters are the abbreviations
	// used by the String method's formatting.
	ModeDir        = fs.ModeDir        // d: is a directory
	ModeAppend     = fs.ModeAppend     // a: append-only
	ModeExclusive  = fs.ModeExclusive  // l: exclusive use
	ModeTemporary  = fs.ModeTemporary  // T: temporary file; Plan 9 only
	ModeSymlink    = fs.ModeSymlink    // L: symbolic link
	ModeDevice     = fs.ModeDevice     // D: device file
	ModeNamedPipe  = fs.ModeNamedPipe  // p: named pipe (FIFO)
	ModeSocket     = fs.ModeSocket     // S: Unix domain socket
	ModeSetuid     = fs.ModeSetuid     // u: setuid
	ModeSetgid     = fs.ModeSetgid     // g: setgid
	ModeCharDevice = fs.ModeCharDevice // c: Unix character device, when ModeDevice is set
	ModeSticky     = fs.ModeSticky     // t: sticky
	ModeIrregular  = fs.ModeIrregular  // ?: non-regular file; nothing else is known about this file

	// Mask for the type bits. For regular files, none will be set.
	ModeType = fs.ModeType

	ModePerm = fs.ModePerm // Unix permission bits, 0o777
)
```