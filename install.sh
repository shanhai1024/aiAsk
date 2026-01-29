#!/bin/bash
set -e

# 配置信息
REPO="shanhai1024/aiAsk"
BINARY_NAME="ask"
INSTALL_PATH="/usr/local/bin"

# 颜色定义
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m' # 无颜色

echo -e "${GREEN}正在探测系统环境...${NC}"

# 1. 识别操作系统
OS="$(uname -s | tr '[:upper:]' '[:lower:]')"
case "${OS}" in
    darwin*)  OS='darwin' ;;
    linux*)   OS='linux' ;;
    msys*|cygwin*|mingw*) OS='windows' ;; # 增加对 Windows Git Bash 的支持
    *)        echo -e "${RED}❌ 暂不支持该系统: ${OS}${NC}"; exit 1 ;;
esac

# 2. 识别架构
ARCH="$(uname -m)"
case "${ARCH}" in
    x86_64) ARCH='amd64' ;;
    arm64|aarch64) ARCH='arm64' ;;
    i386|i686) ARCH='386' ;;
    *)      echo -e "${RED}❌ 暂不支持该架构: ${ARCH}${NC}"; exit 1 ;;
esac

# 3. 获取最新版本号
echo "正在从 GitHub 获取最新版本..."
LATEST_TAG=$(curl -s "https://api.github.com/repos/${REPO}/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')

if [ -z "$LATEST_TAG" ]; then
    echo -e "${RED}❌ 无法获取最新版本，请确认仓库已有 Release 且网络通畅。${NC}"
    exit 1
fi

# 4. 构建下载链接 (严格匹配 GoReleaser 默认格式)
# 去掉 v 前缀（如果 Release 标签带 v）
VERSION_NUM=${LATEST_TAG#v}
# 这里的命名必须和 .goreleaser.yaml 里的 name_template 保持完全一致
# 假设模板是: {{ .ProjectName }}_{{ .Version }}_{{ .Os }}_{{ .Arch }}
DOWNLOAD_URL="https://github.com/${REPO}/releases/download/${LATEST_TAG}/aiAsk_${VERSION_NUM}_${OS}_${ARCH}.tar.gz"

# 如果是 Windows，后缀改为 .zip
if [ "$OS" = "windows" ]; then
    DOWNLOAD_URL="${DOWNLOAD_URL%.tar.gz}.zip"
fi

echo -e "正在下载版本 ${GREEN}${LATEST_TAG}${NC}..."
curl -L "$DOWNLOAD_URL" -o "${BINARY_NAME}_download"

# 5. 解压并安装
echo -e "正在安装到 ${INSTALL_PATH}..."
if [ "$OS" = "windows" ]; then
    unzip -o "${BINARY_NAME}_download"
    mv "${BINARY_NAME}.exe" "${INSTALL_PATH}/"
else
    tar -xzf "${BINARY_NAME}_download"
    # 增加 sudo 检查，如果目录不可写则报错提示
    if [ ! -w "$INSTALL_PATH" ]; then
        echo "权限不足，尝试使用 sudo 移动文件..."
        sudo mv "${BINARY_NAME}" "${INSTALL_PATH}/"
    else
        mv "${BINARY_NAME}" "${INSTALL_PATH}/"
    fi
    sudo chmod +x "${INSTALL_PATH}/${BINARY_NAME}"
fi

# 6. 清理
rm "${BINARY_NAME}_download"
# 清理解压出的多余文件（如 README, LICENSE 等）
rm -f CHANGELOG.md LICENSE README.md checksums.txt 2>/dev/null || true

echo "-------------------------------------------"
echo -e "${GREEN}✅ 安装成功！${NC}请输入 '${BINARY_NAME}' 开始使用。"