
# 🤖 Ask (aiAsk)

[English](#english) | [中文](#chinese)

---

<a name="english"></a>

## English

> **Your AI Terminal Expert** —— Context-aware, safety-first, and command-ready.

[![Release](https://img.shields.io/github/v/release/shanhai1024/aiAsk)](https://github.com/shanhai1024/aiAsk/releases)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go)](https://go.dev)

`ask` is a minimalist CLI tool designed for developers. It translates natural language descriptions into executable terminal commands while being fully aware of your environment.

### ✨ Features
- **Environment Awareness**: Automatically identifies OS (macOS/Linux/Windows), shell type, and current directory structure.
- **Safety First**: Built-in risk engine that intercepts dangerous commands like `rm -rf` for confirmation.
- **Multi-Platform**: Support for macOS, Linux, and Windows via GoReleaser.
- **i18n**: Automatically switches between English and Chinese based on system settings.

### 📦 Installation

#### Homebrew (macOS/Linux)
```bash
brew tap shanhai1024/ask
brew install ask
```

#### One-liner (Linux/macOS)

If you don't use Homebrew, use our installation script:

```bash
curl -fsSL [https://raw.githubusercontent.com/shanhai1024/aiAsk/main/install.sh](https://raw.githubusercontent.com/shanhai1024/aiAsk/main/install.sh) | bash
```

### 🚀 Quick Start

Run any query to trigger the setup guide:

```bash
ask "how to find port 8080 process"
```

---

<a name="chinese"></a>

## 中文

> **你的终端 AI 专家** —— 自动生成指令、感知系统环境、守护操作安全。

`ask` 是一个专为程序员设计的极简命令行工具。它能将你的自然语言描述直接转化为可执行的终端指令，让你不再受困于复杂的参数记忆。

### ✨ 功能特性

* **环境感知**：自动识别操作系统、Shell 类型以及当前目录的文件结构。
* **安全拦截**：内置风险检测引擎，针对 `rm -rf`、`format` 等危险指令强制弹窗确认。
* **全平台支持**：通过 GoReleaser 提供 macOS、Linux 和 Windows 的原生支持。
* **国际化**：根据系统语言环境自动切换中英文界面。

### 📦 安装方式

#### 使用 Homebrew (macOS/Linux)

```bash
brew tap shanhai1024/ask
brew install ask
```

#### 一键安装脚本 (Linux/macOS)

如果你不使用 Homebrew，可以执行以下命令直接安装：

```bash
curl -fsSL [https://raw.githubusercontent.com/shanhai1024/aiAsk/main/install.sh](https://raw.githubusercontent.com/shanhai1024/aiAsk/main/install.sh) | bash
```

### 🚀 快速上手

运行任意指令即可触发初始化向导：

```bash
ask "查找占用 8080 端口的进程"
```

---

## 📄 License

本项目采用 [MIT License](https://www.google.com/search?q=LICENSE) 协议。
This project is licensed under the [MIT License](https://www.google.com/search?q=LICENSE).

