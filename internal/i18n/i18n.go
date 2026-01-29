package i18n

import (
	"os"
	"strings"
)

type Language string

const (
	ZH Language = "zh"
	EN Language = "en"
)

var CurrentLang Language

func Detect() {
	// 优先级：LANG > LC_ALL > 默认英文
	lang := os.Getenv("LANG")
	if lang == "" {
		lang = os.Getenv("LC_ALL")
	}

	if strings.HasPrefix(strings.ToLower(lang), "zh") {
		CurrentLang = ZH
	} else {
		CurrentLang = EN
	}
}

var texts = map[string]map[Language]string{
	// Root
	"root_short":   {ZH: "🤖 Ask: 你的 AI 终端指令专家", EN: "🤖 Ask: Your AI Terminal Expert"},
	"root_long":    {ZH: "极简 AI 命令行助手，自动识别系统环境并生成指令。", EN: "Minimalist AI CLI tool that detects your OS environment."},
	"root_example": {ZH: "  ask \"查看磁盘空间\"\n  ask set init", EN: "  ask \"check disk space\"\n  ask set init"},

	// Groups
	"group_start":  {ZH: "开始使用:", EN: "Start:"},
	"group_config": {ZH: "配置管理:", EN: "Config:"},
	"group_help":   {ZH: "获取帮助:", EN: "Help:"},

	// Common
	"usage":       {ZH: "用法", EN: "Usage"},
	"requirement": {ZH: "需求描述", EN: "requirement"},
	"commands":    {ZH: "可用命令", EN: "Commands"},
	"example":     {ZH: "示例", EN: "Example"},
	"done_msg":    {ZH: "\n✅ 完成。按回车键返回终端...", EN: "\n✅ Done. Press Enter to return..."},

	// Settings
	"set_short":      {ZH: "配置管理", EN: "Settings"},
	"set_init_short": {ZH: "初始化配置向导", EN: "Initialize Setup Guide"},
	"risk_warning":   {ZH: "警告：该命令具有潜在风险！", EN: "WARNING: Potential risk detected!"},
	"auth_err":       {ZH: "身份验证失败：请检查 API Key 或余额。", EN: "Auth Error: Please check API Key or balance."},

	// UI
	"ai_thinking": {ZH: "AI 正在分析方案...", EN: "AI Thinking..."},
	"ui_title":    {ZH: "选择命令 (c:复制 | Enter:执行)", EN: "Select (Enter:Run | c:Copy)"},
	"copied_msg":  {ZH: "已成功复制到剪贴板。", EN: "Copied to clipboard."},

	// Setup Guide
	"setup_welcome": {ZH: "\n🚀 欢迎使用 Ask 配置向导", EN: "\n🚀 Welcome to Ask Setup"},
	"setup_select":  {ZH: "请选择 AI 服务商:", EN: "Select AI Provider:"},
	"setup_input":   {ZH: "\n请输入编号 (1-4): ", EN: "\nEnter choice (1-4): "},
	"setup_url":     {ZH: "输入 API URL: ", EN: "Enter API URL: "},
	"setup_model":   {ZH: "输入模型名称: ", EN: "Enter Model Name: "},
	"setup_key":     {ZH: "请输入 API Key: ", EN: "Enter API Key: "},
	"setup_invalid": {ZH: "⚠️ 无效选择，默认使用 DeepSeek。", EN: "⚠️ Invalid choice, using DeepSeek."},
	"setup_done":    {ZH: "\n✅ 配置已保存！", EN: "\n✅ Setup complete!"},
}

func T(key string) string {
	if m, ok := texts[key]; ok {
		return m[CurrentLang]
	}
	return key
}
