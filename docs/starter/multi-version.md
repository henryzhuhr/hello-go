---
outline: deep
---
# Go 多版本管理

## 官方多版本管理方案

官方版本管理方案参考 [_Managing Go installations_](https://go.dev/doc/manage-install)

安装新版本 Go 例如

```shell
# 下载 Go 1.20 版本
go install golang.org/dl/go1.20@latest
# go: downloading golang.org/dl v0.0.0-20240813161640-304e16060ce9
# 如果可以搜到版本，则继续下载

go1.20 download
# Unpacking ~/sdk/go1.20/go1.20.darwin-arm64.tar.gz ...
# Success. You may now run 'go1.20'
```

随后就可以使用 `go1.20` 命令来使用 Go 1.20 版本，例如查看版本和环境变量确认是否切换成功

```shell
go1.20 version
# go version go1.20 darwin/arm64

go1.20 env GOROOT      
# ~/sdk/go1.20
```


## GVM

### 安装 GVM

[GVM (Go Version Manager)](https://github.com/moovweb/gvm) 是一个 Go 语言版本管理工具，可以方便的安装和管理多个版本的 Go 语言

::: code-group

```shell [自动安装]
bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
# Cloning from https://github.com/moovweb/gvm.git to ~/.gvm
# macOS detected. User shell is: /bin/zsh
# Created profile for existing install of Go at '/opt/homebrew/Cellar/go/1.22.2/libexec'
# Installed GVM v1.0.22

# Please restart your terminal session or to get started right away run
#  `source ~/.gvm/scripts/gvm`
```

```shell [手动安装]
# 科学上网下载该脚本
curl -o gvm-installer https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer 
chmod +x gvm-installer
./gvm-installer
```

::: 

估计安装提示，重启终端或者执行 `source ~/.gvm/scripts/gvm` 使得 gvm 生效。

激活脚本已经写入 `~/.zshrc`(`~/.bashrc`) 文件中，启动终端会自动激活 gvm
```shell
[[ -s "$HOME/.gvm/scripts/gvm" ]] && source "$HOME/.gvm/scripts/gvm"
```

### 使用 gvm 管理 Go 版本

查看可用的 go 版本
```shell
gvm listall
# gvm gos (available)
#    go1
#    go1.0.1
#    ...
#    go1.20
#    go1.21.0
#    ...
```

安装指定版本
```shell
gvm install go1.20
```

切换到指定版本，或者设置默认版本
```shell
gvm use go1.20
gvm use go1.20 --default  # 设置默认版本
```

查看当前使用的版本（两个命令都可以查询）
```shell
go version
# go version go1.20 darwin/arm64

gvm list
# gvm gos
# => go1.20
#    system
```

卸载指定版本
```shell
gvm uninstall go1.20
gvm use system  # 切换到系统版本，否则 go 会无法使用
```
