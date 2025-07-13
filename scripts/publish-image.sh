#!/bin/bash

# 发布镜像的脚本
# 用于构建并发布 hello-go 项目的 Docker 镜像

set -e

# =============================================================================
# 配置变量
# =============================================================================

# 镜像基本信息
IMAGE_NAME="hello-go"
DOCKERFILE_PATH="dockerfiles/Dockerfile"
BUILD_CONTEXT="."

# Docker Registry 配置 (可根据需要修改)
# REGISTRY_URL="registry.cn-shanghai.aliyuncs.com"  # 阿里云镜像仓库
# REGISTRY_NAMESPACE="your-namespac e"               # 命名空间
REGISTRY_URL="docker.io"                           # Docker Hub (默认)
REGISTRY_NAMESPACE="henryzhuhr"                     # Docker Hub 用户名

# 版本标签配置
VERSION_TAG=${VERSION:-"latest"}
GIT_COMMIT=$(git rev-parse --short HEAD 2>/dev/null || echo "unknown")
BUILD_TIME=$(date -u +"%Y%m%d-%H%M%S")

# 完整镜像名称
FULL_IMAGE_NAME="${REGISTRY_URL}/${REGISTRY_NAMESPACE}/${IMAGE_NAME}"

# =============================================================================
# 工具函数
# =============================================================================

# 打印带颜色的信息
print_info()    { echo -e "\033[32m[INFO]\033[0m $1"; }
print_warning() { echo -e "\033[33m[WARN]\033[0m $1"; }
print_error()   { echo -e "\033[31m[ERROR]\033[0m $1"; }

# 检查必要的工具
check_requirements() {
    print_info "检查必要工具..."
    
    if ! command -v docker &> /dev/null; then
        print_error "Docker 未安装或未在 PATH 中"
        exit 1
    fi
    
    if ! command -v git &> /dev/null; then
        print_warning "Git 未安装，将使用默认版本标签"
    fi
    
    print_info "✓ 工具检查完成"
}

# 检查 Dockerfile 是否存在
check_dockerfile() {
    print_info "检查 Dockerfile..."
    
    if [ ! -f "$DOCKERFILE_PATH" ]; then
        print_error "Dockerfile 不存在: $DOCKERFILE_PATH"
        exit 1
    fi
    
    print_info "✓ Dockerfile 检查完成: $DOCKERFILE_PATH"
}

# 构建 Docker 镜像
build_image() {
    print_info "开始构建 Docker 镜像..."
    print_info "镜像名称: $FULL_IMAGE_NAME"
    print_info "构建上下文: $BUILD_CONTEXT"
    print_info "Dockerfile: $DOCKERFILE_PATH"
    print_info "version tag: ${FULL_IMAGE_NAME}:${VERSION_TAG}"
    print_info " commit tag: ${FULL_IMAGE_NAME}:${GIT_COMMIT}"
    print_info "  build tag: ${FULL_IMAGE_NAME}:${BUILD_TIME}"
    
    # 构建镜像并添加多个标签
    docker build \
        -f "$DOCKERFILE_PATH" \
        -t "${FULL_IMAGE_NAME}:${VERSION_TAG}" \
        -t "${FULL_IMAGE_NAME}:${GIT_COMMIT}" \
        -t "${FULL_IMAGE_NAME}:${BUILD_TIME}" \
        --build-arg BUILD_TIME="$BUILD_TIME" \
        --build-arg GIT_COMMIT="$GIT_COMMIT" \
        "$BUILD_CONTEXT"
    
    print_info "✓ Docker 镜像构建完成"
}

# 显示镜像信息
show_image_info() {
    print_info "镜像构建信息:"
    echo "  镜像名称: ${FULL_IMAGE_NAME}"
    echo "  版本标签: ${VERSION_TAG}, ${GIT_COMMIT}, ${BUILD_TIME}"
    echo "  构建时间: $(date)"
    echo "  Git提交: ${GIT_COMMIT}"
    
    # 显示镜像大小
    IMAGE_SIZE=$(docker images "${FULL_IMAGE_NAME}:${VERSION_TAG}" --format "table {{.Size}}" | tail -n 1)
    echo "  镜像大小: ${IMAGE_SIZE}"
}

# 推送镜像到仓库
push_image() {
    if [ "$SKIP_PUSH" = "true" ]; then
        print_warning "跳过镜像推送 (SKIP_PUSH=true)"
        return 0
    fi
    
    print_info "开始推送镜像到仓库..."
    
    # 检查是否已登录 Docker Registry
    if ! docker info | grep -q "Username"; then
        print_warning "请先登录 Docker Registry:"
        echo "  docker login ${REGISTRY_URL}"
        
        read -p "是否现在登录? (y/N): " -n 1 -r
        echo
        if [[ $REPLY =~ ^[Yy]$ ]]; then
            docker login "${REGISTRY_URL}"
        else
            print_warning "跳过镜像推送"
            return 0
        fi
    fi
    
    # 推送所有标签
    print_info "推送标签: ${VERSION_TAG}"
    docker push "${FULL_IMAGE_NAME}:${VERSION_TAG}"
    
    print_info "推送标签: ${GIT_COMMIT}"
    docker push "${FULL_IMAGE_NAME}:${GIT_COMMIT}"
    
    print_info "推送标签: ${BUILD_TIME}"
    docker push "${FULL_IMAGE_NAME}:${BUILD_TIME}"
    
    print_info "✓ 镜像推送完成"
}

# 清理本地镜像 (可选)
cleanup_images() {
    if [ "$CLEANUP" = "true" ]; then
        print_info "清理本地镜像..."
        
        # 只保留 latest 标签，删除其他标签
        docker rmi "${FULL_IMAGE_NAME}:${GIT_COMMIT}" 2>/dev/null || true
        docker rmi "${FULL_IMAGE_NAME}:${BUILD_TIME}" 2>/dev/null || true
        
        print_info "✓ 镜像清理完成"
    fi
}

# 显示使用帮助
show_help() {
    cat << EOF
    用法: $0 [选项]

    选项:
        -v, --version VERSION    指定版本标签 (默认: latest)
        -r, --registry URL       指定 Docker Registry URL (默认: docker.io)
        -n, --namespace NAME     指定命名空间 (默认: henryzhuhr)
        --skip-push             跳过推送到仓库
        --cleanup               构建后清理本地镜像
        -h, --help              显示此帮助信息

    环境变量:
        VERSION                 版本标签
        REGISTRY_URL            Docker Registry URL
        REGISTRY_NAMESPACE      命名空间
        SKIP_PUSH              跳过推送 (true/false)
        CLEANUP                清理镜像 (true/false)

    示例:
        $0                                    # 使用默认配置
        $0 -v v1.0.0                         # 指定版本
        $0 --skip-push                       # 只构建不推送
        $0 -r registry.cn-shanghai.aliyuncs.com -n my-namespace  # 使用阿里云仓库

EOF
}

# =============================================================================
# 主函数
# =============================================================================

main() {
    print_info "开始执行 Docker 镜像发布流程..."
    print_info "项目: hello-go"
    print_info "分支: $(git branch --show-current 2>/dev/null || echo 'unknown')"
    
    # 执行构建流程
    check_requirements
    check_dockerfile
    build_image
    show_image_info
    push_image
    cleanup_images
    
    print_info "🎉 Docker 镜像发布完成!"
    print_info "镜像地址: ${FULL_IMAGE_NAME}:${VERSION_TAG}"
}

# =============================================================================
# 参数解析
# =============================================================================

# 解析命令行参数
while [[ $# -gt 0 ]]; do
    case $1 in
        -v|--version)
            VERSION_TAG="$2"
            shift 2
            ;;
        -r|--registry)
            REGISTRY_URL="$2"
            shift 2
            ;;
        -n|--namespace)
            REGISTRY_NAMESPACE="$2"
            shift 2
            ;;
        --skip-push)
            SKIP_PUSH="true"
            shift
            ;;
        --cleanup)
            CLEANUP="true"
            shift
            ;;
        -h|--help)
            show_help
            exit 0
            ;;
        *)
            print_error "未知参数: $1"
            show_help
            exit 1
            ;;
    esac
done

# 更新完整镜像名称
FULL_IMAGE_NAME="${REGISTRY_URL}/${REGISTRY_NAMESPACE}/${IMAGE_NAME}"

# 执行主函数
main
