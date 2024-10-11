---
outline: deep
---

## VSCode 配置

### 插件和依赖

VSCode 需要安装如下插件，可以在命令面板(`Ctrl+Shift+P`)中搜索 `Go: Install/Update Tools` 安装

- [`gopls`](https://github.com/golang/tools/blob/master/gopls/README.md#installation)
- [`gotests`](https://github.com/cweill/gotests)
- [`gomodifytags`](https://github.com/fatih/gomodifytags)
- [`impl`](https://github.com/josharian/impl)
- [`goplay`](https://github.com/haya14busa/goplay)
- [`dlv`](https://github.com/go-delve/delve)
- [`staticcheck`](https://github.com/dominikh/go-tools)

如果存在网络问题，可以手动使用下面命令手动安装，或者参照各个工具的手动安装手册，下载源码编译安装
```bash
go install golang.org/x/tools/gopls@latest
go get -u  github.com/cweill/gotests/...
go install github.com/fatih/gomodifytags@latest
go install github.com/josharian/impl@latest
# go get -u  github.com/haya14busa/goplay            # as a Library
go get -u  github.com/haya14busa/goplay/cmd/goplay # as a command line tool
go install github.com/go-delve/delve/cmd/dlv@latest
go install honnef.co/go/tools/cmd/staticcheck@latest
```

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

