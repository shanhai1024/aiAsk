package shell

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
)

// CheckRisk 评估指令风险等级 (0: 安全, 1: 危险)
// 增加了对复合指令（如 find -delete）和敏感路径的深度校验
func CheckRisk(command string) int {
	cmdLower := strings.ToLower(strings.TrimSpace(command))
	if cmdLower == "" {
		return 0
	}

	// 1. 危险二进制命令名 (精确匹配首个单词)
	// 移除了 ls，增加了磁盘管理类工具
	dangerousBins := []string{
		"rm", "mkfs", "dd", "shutdown", "reboot", "format",
		"fdisk", "parted", "mkswap", "shred",
	}
	parts := strings.Fields(cmdLower)
	if len(parts) > 0 {
		firstWord := parts[0]
		for _, bin := range dangerousBins {
			if firstWord == bin {
				return 1
			}
		}
	}

	// 2. 高级风险模式校验 (使用正则表达式)
	riskPatterns := []string{
		// 递归删除类：拦截 find . -delete 或 find -exec rm 等变体
		`find\s+.*-(delete|exec\s+rm)`,

		// 递归强制删除：拦截 rm -rf 或 rm -fr
		`rm\s+.*-r?f`,

		// 敏感系统路径保护：拦截涉及根目录、配置目录、引导目录的危险操作
		`/\s*(etc|bin|sbin|var|usr|root|boot|dev)(\s|/|$)`,

		// 磁盘物理操作：拦截直接向物理磁盘写入数据
		`>\s*/dev/sd[a-z]`,

		// 危险网络执行：拦截 curl/wget 获取内容并直接交给 Shell 执行的行为
		`\|\s*(bash|sh|zsh|py|python|perl|php)`,

		// 破坏性脚本：防御 Fork 炸弹（耗尽系统资源）
		`:\(\)\{\s*:\|:&\s*\};:`,

		// 权限/所有权滥用：拦截对根路径进行权限全开的操作
		`chmod\s+.*777`,
		`chown\s+.*root`,
	}

	for _, pattern := range riskPatterns {
		matched, _ := regexp.MatchString("(?i)"+pattern, cmdLower)
		if matched {
			return 1
		}
	}

	return 0
}

// CheckBin 检查命令在当前系统中是否存在
func CheckBin(command string) bool {
	fields := strings.Fields(command)
	if len(fields) == 0 {
		return false
	}
	bin := fields[0]
	_, err := exec.LookPath(bin)
	return err == nil
}

// GetInstallCmd 根据操作系统提供对应的包管理安装指令
func GetInstallCmd(command string) string {
	fields := strings.Fields(command)
	if len(fields) == 0 {
		return ""
	}
	bin := fields[0]
	if runtime.GOOS == "darwin" {
		return fmt.Sprintf("brew install %s", bin)
	}
	if runtime.GOOS == "linux" {
		return fmt.Sprintf("sudo apt install %s", bin)
	}
	return ""
}

// Execute 执行终端指令
// 支持 Windows (cmd /C) 和类 Unix 系统 (读取 $SHELL 环境变量)
func Execute(command string) {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", command)
	} else {
		shell := os.Getenv("SHELL")
		if shell == "" {
			shell = "/bin/bash"
		}
		cmd = exec.Command(shell, "-c", command)
	}
	// 绑定标准输入输出，确保交互式命令正常工作
	cmd.Stdout, cmd.Stderr, cmd.Stdin = os.Stdout, os.Stderr, os.Stdin
	cmd.Run()
}
