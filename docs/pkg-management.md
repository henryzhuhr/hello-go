---
outline: deep
---
# Go Modules 包管理

[Go Modules](https://go.dev/blog/using-go-modules) 是 Go 1.11 版本后引入的包管理工具，它可以帮助我们更好的管理第三方包，解决了以往的 GOPATH 机制下，包管理不方便的问题，同时也解决了包版本管理的问题。


创建 modules 的命令
```bash
go mod init github.com/<username>/<projectname>
```



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


拉取私有仓库

Go 1.13 版本后，拉取私有仓库的方式发生了变化，需要配置环境变量，否则会报错

```shell
export GOPRIVATE=github.com/<username>
```
同时由于 `go get` 默认使用的是 https 协议，如果你的私有仓库是 ssh 协议，需要配置 `~/.gitconfig` 文件，添加如下配置

```shell
git config --global url."git@git.xxx.com".insteadOf "https://git.xxx.com/"
```