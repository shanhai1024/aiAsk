package shell

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// CheckRisk 评估指令风险等级 (0: 安全, 1: 危险)
func CheckRisk(command string) int {
	cmdLower := strings.ToLower(strings.TrimSpace(command))
	parts := strings.Fields(cmdLower)
	if len(parts) == 0 {
		return 0
	}

	// 危险二进制命令名 (精确匹配第一个单词)
	dangerousBins := []string{
		"rm", "ls", "mkfs", "dd", "shutdown", "reboot", "format", "mkfs.ext4",
	}
	firstWord := parts[0]
	for _, bin := range dangerousBins {
		if firstWord == bin {
			return 1
		}
	}

	// 危险语法特征 (全文本包含匹配)
	specialThreats := []string{
		"> /dev/",
		":(){ :|:& };:",
		"mv / ",
		"chmod -r 777 /",
	}
	for _, threat := range specialThreats {
		if strings.Contains(cmdLower, threat) {
			return 1
		}
	}

	return 0
}

// CheckBin 检查命令是否存在
func CheckBin(command string) bool {
	bin := strings.Fields(command)[0]
	_, err := exec.LookPath(bin)
	return err == nil
}

// GetInstallCmd 根据系统给出安装建议
func GetInstallCmd(command string) string {
	bin := strings.Fields(command)[0]
	if runtime.GOOS == "darwin" {
		return fmt.Sprintf("brew install %s", bin)
	}
	if runtime.GOOS == "linux" {
		return fmt.Sprintf("sudo apt install %s", bin)
	}
	return ""
}

// Execute 执行命令
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
	cmd.Stdout, cmd.Stderr, cmd.Stdin = os.Stdout, os.Stderr, os.Stdin
	cmd.Run()
}
