
# 🤖 Ask: Your AI Terminal Expert | 你的 AI 终端指令专家

<p align="center">
  <img src="https://img.shields.io/github/v/release/shanhai1024/aiAsk?style=flat-square" alt="Release">
  <img src="https://img.shields.io/github/license/shanhai1024/aiAsk?style=flat-square" alt="License">
  <img src="https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat-square&logo=go" alt="Go Version">
</p>

[English](#english) | [中文](#中文)

---

## English

**Ask** is a minimalist AI-powered CLI assistant designed to boost your terminal productivity. It automatically detects your OS environment (macOS, Linux, Windows) and generates the precise shell commands you need using LLMs.

### ✨ Features
- **OS Aware**: Automatically detects your OS, Shell, and CWD.
- **Multi-Provider**: Supports DeepSeek, ChatGPT, Gemini, and Custom (Ollama).
- **Interactive UI**: Select, copy, or execute commands directly within a TUI.
- **I18n**: Automatically switches UI language based on your system locale.
- **One-Line Install**: Get up and running in seconds.

### 🚀 Quick Start

#### One-Line Installation (macOS/Linux)
```bash
curl -fsSL [https://raw.githubusercontent.com/shanhai1024/aiAsk/main/install.sh](https://raw.githubusercontent.com/shanhai1024/aiAsk/main/install.sh) | bash

```

#### First-Run Setup

Simply run `ask` followed by a query. If it's your first time, the setup wizard will appear:

```bash
ask "check disk space"

```

#### Manual Configuration

```bash
ask set init    # Restart the setup wizard
ask set key     # Update API Key

```

---

## 中文

**Ask** 是一款极简的 AI 驱动命令行助手，旨在提升你的终端生产力。它能自动识别系统环境（macOS, Linux, Windows），并利用大模型生成你所需的精确 Shell 指令。

### ✨ 项目特性

* **环境感知**：自动识别操作系统、Shell 类型及当前工作目录。
* **多模型支持**：预设支持 DeepSeek、ChatGPT、Gemini 以及自定义模型（如 Ollama）。
* **交互式界面**：在 TUI 界面中直接选择、复制或执行生成的命令。
* **多语言支持**：根据系统语言环境自动切换中英文界面。
* **极简安装**：支持一行命令安装，无环境依赖。

### 🚀 快速开始

#### 一键安装 (macOS/Linux)

```bash
curl -fsSL [https://raw.githubusercontent.com/shanhai1024/aiAsk/main/install.sh](https://raw.githubusercontent.com/shanhai1024/aiAsk/main/install.sh) | bash

```

#### 首次运行

直接输入 `ask` 加上你的需求。如果是首次运行，会自动弹出配置向导：

```bash
ask "查看磁盘空间"

```

#### 手动配置

```bash
ask set init    # 重新运行初始化向导
ask set key     # 单独更新 API Key

```

---

## 🛠️ Configuration / 配置说明

| Provider | URL | Model |
| --- | --- | --- |
| **DeepSeek** | `https://api.deepseek.com/v1/chat/completions` | `deepseek-chat` |
| **OpenAI** | `https://api.openai.com/v1/chat/completions` | `gpt-4o-mini` |
| **Ollama** | `http://localhost:11434/v1/chat/completions` | `llama3 / qwen2.5` |

---

## 🤝 Contributing

Feel free to open issues or submit pull requests to help make **Ask** even better!

## 📄 License

[MIT](https://www.google.com/search?q=LICENSE)



