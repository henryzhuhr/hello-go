# Hello VuePress


## 开发环境

安装 Go 后，配置 GOPATH， `export GOPATH="$HOME/project/goproject"` ，开发时，代码存放在 `$GOPATH/src` 下， 工程经过 `go build`、`go install` 或`go get` 等指令后，会将下载的第三方包源代码文件放在 `$GOPATH/src` 目录下， 产生的二进制可执行文件放在 `$GOPATH/bin` 目录下，生成的中间缓存文件会被保存在 `$GOPATH/pkg` 下


配置代理
```bash
export GOPROXY="https://goproxy.cn,direct"
```


## Go Modules 包管理

[Go Modules](https://go.dev/blog/using-go-modules) 是 Go 1.11 版本后引入的包管理工具，它可以帮助我们更好的管理第三方包，解决了以往的 GOPATH 机制下，包管理不方便的问题，同时也解决了包版本管理的问题。


创建 modules 的命令
```bash
go mod init github.com/<username>/<projectname>
```



vscode 还有一项很强大的功能就是断点调试,结合 delve 可以很好的进行 Go 代码调试

```shell
go get -v -u github.com/peterh/liner github.com/derekparker/delve/cmd/dlv

brew install go-delve/delve/delve (mac可选)
```

如果有问题再来一遍:
```shell
go get -v -u github.com/peterh/liner github.com/derekparker/delve/cmd/dlv
```

注意:修改"dlv-cert"证书, 选择"显示简介"->"信任"->"代码签名" 修改为: 始终信任

打开首选项-工作区设置,配置launch.json:
```
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
```shell