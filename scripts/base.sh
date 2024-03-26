# 初始化:
go mod init [module 名称]

# 检测和清理依赖:
go mod tidy


# 安装指定包:
go get -v github.com/go-ego/gse@v0.60.0-rc4.2



# 更新依赖
go get -u

# 更新指定包依赖:
go get -u github.com/go-ego/gse

# 指定版本:
go get -u github/com/go-ego/gse@v0.60.0-rc4.2


go mod init  # 初始化 go.mod
go mod tidy  # 更新依赖文件
go mod download  # 下载依赖文件

go mod vendor  # 将依赖转移至本地的 vendor 文件
go mod edit  # 手动修改依赖文件
go mod graph  # 打印依赖图
go mod verify  # 校验依赖