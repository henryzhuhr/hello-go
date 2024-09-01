---
outline: deep
---
# Go 开发环境

## 安装

### 安装预编译版本

### 从源码编译

参考 [Installing Go from source](https://go.dev/doc/install/source)


## 配置
安装 Go 后，配置 GOPATH， `export GOPATH="$HOME/project/goproject"` ，开发时，代码存放在 `$GOPATH/src` 下， 工程经过 `go build`、`go install` 或`go get` 等指令后，会将下载的第三方包源代码文件放在 `$GOPATH/src` 目录下， 产生的二进制可执行文件放在 `$GOPATH/bin` 目录下，生成的中间缓存文件会被保存在 `$GOPATH/pkg` 下


配置代理
```bash
export GOPROXY="https://goproxy.cn,direct"
```


## VSCode 配置
### 断点调试
vscode 还有一项很强大的功能就是断点调试,结合 delve 可以很好的进行 Go 代码调试

```shell
go get -v -u github.com/peterh/liner github.com/derekparker/delve/cmd/dlv
brew install go-delve/delve/delve # mac 可以使用 brew 安装
```

如果有问题再来一遍:
```shell
go get -v -u github.com/peterh/liner github.com/derekparker/delve/cmd/dlv
```

注意:修改"dlv-cert"证书, 选择"显示简介"->"信任"->"代码签名" 修改为: 始终信任

打开首选项-工作区设置,配置launch.json:
```json
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "main.go",
            "type": "go",
            "request": "launch",
            "mode": "debug",
            "remotePath": "",
            "port": 2345,
            "host": "127.0.0.1",
            "program": "${workspaceRoot}", // 工作空间路径
            "env": {},
            "args": [],
            "showLog": true
        }
    ]
}
```
