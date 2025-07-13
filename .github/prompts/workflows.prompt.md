# Github Actions 工作流自动构建和推送 Docker 镜像

需求：通过每次提交或发布时自动构建并推送 Docker 镜像到 GitHub Container Registry (GHCR)，支持多版本 Go 和多架构。

工作流文件定义在 `.github/workflows/build-and-push-image.yaml`

Dockerfile 定义在 `dockerfiles/Dockerfile`，里面设置了一些自定义参数，你需要参考 workflows 文件如何输入参数进行构建