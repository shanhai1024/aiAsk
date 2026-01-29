#!/bin/bash
set -e

# 配置信息
REPO="shanhai1024/aiAsk"
BINARY_NAME="ask"
INSTALL_PATH="/usr/local/bin"

echo "正在探测系统环境..."

# 1. 识别操作系统
OS="$(uname -s | tr '[:upper:]' '[:lower:]')"
case "${OS}" in
    darwin*)  OS='darwin' ;;
    linux*)   OS='linux' ;;
    *)        echo "❌ 暂不支持该系统: ${OS}"; exit 1 ;;
esac

# 2. 识别架构
ARCH="$(uname -m)"
case "${ARCH}" in
    x86_64) ARCH='amd64' ;;
    arm64|aarch64) ARCH='arm64' ;;
    *)      echo "❌ 暂不支持该架构: ${ARCH}"; exit 1 ;;
esac

# 3. 获取最新版本号
echo "正在从 GitHub 获取最新版本..."
LATEST_TAG=$(curl -s "https://api.github.com/repos/${REPO}/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')

if [ -z "$LATEST_TAG" ]; then
    echo "❌ 无法获取最新版本，请检查网络或仓库是否已有 Release。"
    exit 1
fi

# 4. 构建下载链接 (匹配 GoReleaser 的命名规则)
# 例如: aiAsk_v1.0.0_darwin_arm64.tar.gz
DOWNLOAD_URL="https://github.com/${REPO}/releases/download/${LATEST_TAG}/aiAsk_${LATEST_TAG#v}_${OS}_${ARCH}.tar.gz"

echo "正在下载版本 ${LATEST_TAG}..."
curl -L "$DOWNLOAD_URL" -o "${BINARY_NAME}.tar.gz"

# 5. 解压并安装
echo "正在安装到 ${INSTALL_PATH}..."
tar -xzf "${BINARY_NAME}.tar.gz"
sudo mv "${BINARY_NAME}" "${INSTALL_PATH}/"
sudo chmod +x "${INSTALL_PATH}/${BINARY_NAME}"

# 6. 清理
rm "${BINARY_NAME}.tar.gz"

echo "-------------------------------------------"
echo "✅ 安装成功！请输入 'ask' 开始使用。"
echo "如果不生效，请运行 'ask set init' 进行初始化。"